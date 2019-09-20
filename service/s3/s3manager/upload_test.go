package s3manager_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting/unit"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/s3testing"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
)

const respBody = `<?xml version="1.0" encoding="UTF-8"?>
<CompleteMultipartUploadOutput>
   <Location>mockValue</Location>
   <Bucket>mockValue</Bucket>
   <Key>mockValue</Key>
   <ETag>mockValue</ETag>
</CompleteMultipartUploadOutput>`

var emptyList []string

func val(i interface{}, s string) interface{} {
	v, err := awsutil.ValuesAtPath(i, s)
	if err != nil || len(v) == 0 {
		return nil
	}
	if _, ok := v[0].(io.Reader); ok {
		return v[0]
	}

	if rv := reflect.ValueOf(v[0]); rv.Kind() == reflect.Ptr {
		return rv.Elem().Interface()
	}

	return v[0]
}

func contains(src []string, s string) bool {
	for _, v := range src {
		if s == v {
			return true
		}
	}
	return false
}

func loggingSvc(ignoreOps []string) (*s3.Client, *[]string, *[]interface{}) {
	var m sync.Mutex
	partNum := 0
	names := []string{}
	params := []interface{}{}
	svc := s3.New(unit.Config())
	svc.Handlers.Unmarshal.Clear()
	svc.Handlers.UnmarshalMeta.Clear()
	svc.Handlers.UnmarshalError.Clear()
	svc.Handlers.Send.Clear()
	svc.Handlers.Send.PushBack(func(r *aws.Request) {
		m.Lock()
		defer m.Unlock()

		if !contains(ignoreOps, r.Operation.Name) {
			names = append(names, r.Operation.Name)
			params = append(params, r.Params)
		}

		r.HTTPResponse = &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(respBody))),
		}

		switch data := r.Data.(type) {
		case *s3.CreateMultipartUploadOutput:
			data.UploadId = aws.String("UPLOAD-ID")
		case *s3.UploadPartOutput:
			partNum++
			data.ETag = aws.String(fmt.Sprintf("ETAG%d", partNum))
		case *s3.CompleteMultipartUploadOutput:
			data.Location = aws.String("https://location")
			data.VersionId = aws.String("VERSION-ID")
		case *s3.PutObjectOutput:
			data.VersionId = aws.String("VERSION-ID")
		}
	})

	return svc, &names, &params
}

func buflen(i interface{}) int {
	r := i.(io.Reader)
	b, _ := ioutil.ReadAll(r)
	return len(b)
}

func TestUploadOrderMulti(t *testing.T) {
	s, ops, args := loggingSvc(emptyList)
	u := s3manager.NewUploaderWithClient(s)

	resp, err := u.Upload(&s3manager.UploadInput{
		Bucket:               aws.String("Bucket"),
		Key:                  aws.String("Key - value"),
		Body:                 bytes.NewReader(buf12MB),
		ServerSideEncryption: s3.ServerSideEncryptionAwsKms,
		SSEKMSKeyId:          aws.String("KmsId"),
		ContentType:          aws.String("content/type"),
	})

	if err != nil {
		t.Errorf("Expected no error but received %v", err)
	}

	expected := []string{"CreateMultipartUpload", "UploadPart", "UploadPart", "UploadPart", "CompleteMultipartUpload"}
	if !reflect.DeepEqual(expected, *ops) {
		t.Errorf("Expected %v, but received %v", expected, *ops)
	}

	if e, a := `https://endpoint/Bucket/Key%20-%20value`, resp.Location; e != a {
		t.Errorf("Expected %q, but received %q", e, a)
	}

	if "UPLOAD-ID" != resp.UploadID {
		t.Errorf("Expected %q, but received %q", "UPLOAD-ID", resp.UploadID)
	}

	if "VERSION-ID" != *resp.VersionID {
		t.Errorf("Expected %q, but received %q", "VERSION-ID", *resp.VersionID)
	}

	// Validate input values

	// UploadPart
	for i := 1; i < 5; i++ {
		v := val((*args)[i], "UploadId")
		if "UPLOAD-ID" != v {
			t.Errorf("Expected %q, but received %q", "UPLOAD-ID", v)
		}
	}

	// CompleteMultipartUpload
	v := val((*args)[4], "UploadId")
	if "UPLOAD-ID" != v {
		t.Errorf("Expected %q, but received %q", "UPLOAD-ID", v)
	}

	for i := 0; i < 3; i++ {
		e := val((*args)[4], fmt.Sprintf("MultipartUpload.Parts[%d].PartNumber", i))
		if int64(i+1) != e.(int64) {
			t.Errorf("Expected %d, but received %d", i+1, e)
		}
	}

	vals := []string{
		val((*args)[4], "MultipartUpload.Parts[0].ETag").(string),
		val((*args)[4], "MultipartUpload.Parts[1].ETag").(string),
		val((*args)[4], "MultipartUpload.Parts[2].ETag").(string),
	}

	for _, a := range vals {
		if matched, err := regexp.MatchString(`^ETAG\d+$`, a); !matched || err != nil {
			t.Errorf("Failed regexp expression `^ETAG\\d+$`")
		}
	}

	// Custom headers
	e := val((*args)[0], "ServerSideEncryption")
	if e.(s3.ServerSideEncryption) != s3.ServerSideEncryptionAwsKms {
		t.Errorf("Expected %q, but received %q", "aws:kms", e)
	}

	e = val((*args)[0], "SSEKMSKeyId")
	if e != "KmsId" {
		t.Errorf("Expected %q, but received %q", "KmsId", e)
	}

	e = val((*args)[0], "ContentType")
	if e != "content/type" {
		t.Errorf("Expected %q, but received %q", "content/type", e)
	}
}

func TestUploadOrderMultiDifferentPartSize(t *testing.T) {
	s, ops, args := loggingSvc(emptyList)
	mgr := s3manager.NewUploaderWithClient(s, func(u *s3manager.Uploader) {
		u.PartSize = 1024 * 1024 * 7
		u.Concurrency = 1
	})
	_, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   bytes.NewReader(buf12MB),
	})

	if err != nil {
		t.Errorf("Expected no error but received %v", err)
	}

	vals := []string{"CreateMultipartUpload", "UploadPart", "UploadPart", "CompleteMultipartUpload"}
	if !reflect.DeepEqual(vals, *ops) {
		t.Errorf("Expected %v, but received %v", vals, *ops)
	}

	// Part lengths
	if len := buflen(val((*args)[1], "Body")); 1024*1024*7 != len {
		t.Errorf("Expected %d, but received %d", 1024*1024*7, len)
	}
	if len := buflen(val((*args)[2], "Body")); 1024*1024*5 != len {
		t.Errorf("Expected %d, but received %d", 1024*1024*5, len)
	}
}

func TestUploadIncreasePartSize(t *testing.T) {
	s, ops, args := loggingSvc(emptyList)
	mgr := s3manager.NewUploaderWithClient(s, func(u *s3manager.Uploader) {
		u.Concurrency = 1
		u.MaxUploadParts = 2
	})
	_, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   bytes.NewReader(buf12MB),
	})

	if err != nil {
		t.Errorf("Expected no error but received %v", err)
	}

	if int64(s3manager.DefaultDownloadPartSize) != mgr.PartSize {
		t.Errorf("Expected %d, but received %d", s3manager.DefaultDownloadPartSize, mgr.PartSize)
	}

	vals := []string{"CreateMultipartUpload", "UploadPart", "UploadPart", "CompleteMultipartUpload"}
	if !reflect.DeepEqual(vals, *ops) {
		t.Errorf("Expected %v, but received %v", vals, *ops)
	}

	// Part lengths
	if len := buflen(val((*args)[1], "Body")); (1024*1024*6)+1 != len {
		t.Errorf("Expected %d, but received %d", (1024*1024*6)+1, len)
	}

	if len := buflen(val((*args)[2], "Body")); (1024*1024*6)-1 != len {
		t.Errorf("Expected %d, but received %d", (1024*1024*6)-1, len)
	}
}

func TestUploadFailIfPartSizeTooSmall(t *testing.T) {
	mgr := s3manager.NewUploader(unit.Config(), func(u *s3manager.Uploader) {
		u.PartSize = 5
	})
	resp, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   bytes.NewReader(buf12MB),
	})

	if resp != nil {
		t.Errorf("Expected response to be nil, but received %v", resp)
	}

	if err == nil {
		t.Errorf("Expected error, but received nil")
	}

	aerr := err.(awserr.Error)
	if e, a := "ConfigError", aerr.Code(); e != a {
		t.Errorf("Expected %q, but received %q", e, a)
	}

	if e, a := "part size must be at least", aerr.Message(); !strings.Contains(a, e) {
		t.Errorf("expect %v to be in %v", e, a)
	}
}

type negativeReader struct {
}

func (nr *negativeReader) Read(_ []byte) (int, error) {
	return -1, io.ErrUnexpectedEOF
}

func TestUploadReaderReturnsNegative(t *testing.T) {
	s, _, _ := loggingSvc(emptyList)
	mgr := s3manager.NewUploaderWithClient(s, func(u *s3manager.Uploader) {
		u.Concurrency = 1
	})

	// should not panic
	_, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   &negativeReader{},
	})
	if err == nil {
		t.Error("Expected error, but received none")
	}

	if e, a := io.ErrUnexpectedEOF, err; !errors.Is(a, e) {
		t.Fatalf("expect %v error, got %v", e, a)
	}
}

func TestUploadOrderSingle(t *testing.T) {
	s, ops, args := loggingSvc(emptyList)
	mgr := s3manager.NewUploaderWithClient(s)
	resp, err := mgr.Upload(&s3manager.UploadInput{
		Bucket:               aws.String("Bucket"),
		Key:                  aws.String("Key - value"),
		Body:                 bytes.NewReader(buf2MB),
		ServerSideEncryption: s3.ServerSideEncryptionAwsKms,
		SSEKMSKeyId:          aws.String("KmsId"),
		ContentType:          aws.String("content/type"),
	})

	if err != nil {
		t.Errorf("Expected no error but received %v", err)
	}

	if vals := []string{"PutObject"}; !reflect.DeepEqual(vals, *ops) {
		t.Errorf("Expected %v, but received %v", vals, *ops)
	}

	if e, a := `https://endpoint/Bucket/Key%20-%20value`, resp.Location; e != a {
		t.Errorf("Expected %q, but received %q", e, a)
	}

	if e := "VERSION-ID"; e != *resp.VersionID {
		t.Errorf("Expected %q, but received %q", e, *resp.VersionID)
	}

	if len(resp.UploadID) > 0 {
		t.Errorf("Expected empty string, but received %q", resp.UploadID)
	}

	if e, a := s3.ServerSideEncryptionAwsKms, val((*args)[0], "ServerSideEncryption").(s3.ServerSideEncryption); e != a {
		t.Errorf("Expected %q, but received %q", e, a)
	}

	if e, a := "KmsId", val((*args)[0], "SSEKMSKeyId").(string); e != a {
		t.Errorf("Expected %q, but received %q", e, a)
	}

	if e, a := "content/type", val((*args)[0], "ContentType").(string); e != a {
		t.Errorf("Expected %q, but received %q", e, a)
	}
}

func TestUploadOrderSingleFailure(t *testing.T) {
	s, ops, _ := loggingSvc(emptyList)
	s.Handlers.Send.PushBack(func(r *aws.Request) {
		r.HTTPResponse.StatusCode = 400
	})
	mgr := s3manager.NewUploaderWithClient(s)
	resp, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   bytes.NewReader(buf2MB),
	})

	if err == nil {
		t.Error("Expected error, but receievd nil")
	}

	if vals := []string{"PutObject"}; !reflect.DeepEqual(vals, *ops) {
		t.Errorf("Expected %v, but received %v", vals, *ops)
	}

	if resp != nil {
		t.Errorf("Expected response to be nil, but received %v", resp)
	}
}

func TestUploadOrderZero(t *testing.T) {
	s, ops, args := loggingSvc(emptyList)
	mgr := s3manager.NewUploaderWithClient(s)
	resp, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   bytes.NewReader(make([]byte, 0)),
	})

	if err != nil {
		t.Errorf("Expected no error but received %v", err)
	}

	if vals := []string{"PutObject"}; !reflect.DeepEqual(vals, *ops) {
		t.Errorf("Expected %v, but received %v", vals, *ops)
	}

	if len(resp.Location) == 0 {
		t.Error("Expected Location to not be empty")
	}

	if len(resp.UploadID) > 0 {
		t.Errorf("Expected empty string, but received %q", resp.UploadID)
	}

	if e, a := 0, buflen(val((*args)[0], "Body")); e != a {
		t.Errorf("Expected %d, but received %d", e, a)
	}
}

func TestUploadOrderMultiFailure(t *testing.T) {
	s, ops, _ := loggingSvc(emptyList)
	s.Handlers.Send.PushBack(func(r *aws.Request) {
		switch t := r.Data.(type) {
		case *s3.UploadPartOutput:
			if *t.ETag == "ETAG2" {
				r.HTTPResponse.StatusCode = 400
			}
		}
	})

	mgr := s3manager.NewUploaderWithClient(s, func(u *s3manager.Uploader) {
		u.Concurrency = 1
	})
	_, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   bytes.NewReader(buf12MB),
	})

	if err == nil {
		t.Error("Expected error, but receievd nil")
	}

	if e, a := []string{"CreateMultipartUpload", "UploadPart", "UploadPart", "AbortMultipartUpload"}, *ops; !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %v, but received %v", e, a)
	}
}

func TestUploadOrderMultiFailureOnComplete(t *testing.T) {
	s, ops, _ := loggingSvc(emptyList)
	s.Handlers.Send.PushBack(func(r *aws.Request) {
		switch r.Data.(type) {
		case *s3.CompleteMultipartUploadOutput:
			r.HTTPResponse.StatusCode = 400
		}
	})

	mgr := s3manager.NewUploaderWithClient(s, func(u *s3manager.Uploader) {
		u.Concurrency = 1
	})
	_, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   bytes.NewReader(buf12MB),
	})

	if err == nil {
		t.Error("Expected error, but receievd nil")
	}

	if e, a := []string{"CreateMultipartUpload", "UploadPart", "UploadPart",
		"UploadPart", "CompleteMultipartUpload", "AbortMultipartUpload"}, *ops; !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %v, but received %v", e, a)
	}
}

func TestUploadOrderMultiFailureOnCreate(t *testing.T) {
	s, ops, _ := loggingSvc(emptyList)
	s.Handlers.Send.PushBack(func(r *aws.Request) {
		switch r.Data.(type) {
		case *s3.CreateMultipartUploadOutput:
			r.HTTPResponse.StatusCode = 400
		}
	})

	mgr := s3manager.NewUploaderWithClient(s)
	_, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   bytes.NewReader(make([]byte, 1024*1024*12)),
	})

	if err == nil {
		t.Error("Expected error, but receievd nil")
	}

	if e, a := []string{"CreateMultipartUpload"}, *ops; !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %v, but received %v", e, a)
	}
}

func TestUploadOrderMultiFailureLeaveParts(t *testing.T) {
	s, ops, _ := loggingSvc(emptyList)
	s.Handlers.Send.PushBack(func(r *aws.Request) {
		switch data := r.Data.(type) {
		case *s3.UploadPartOutput:
			if *data.ETag == "ETAG2" {
				r.HTTPResponse.StatusCode = 400
			}
		}
	})

	mgr := s3manager.NewUploaderWithClient(s, func(u *s3manager.Uploader) {
		u.Concurrency = 1
		u.LeavePartsOnError = true
	})
	_, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   bytes.NewReader(make([]byte, 1024*1024*12)),
	})

	if err == nil {
		t.Error("Expected error, but receievd nil")
	}

	if e, a := []string{"CreateMultipartUpload", "UploadPart", "UploadPart"}, *ops; !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %v, but received %v", e, a)
	}
}

type failreader struct {
	times     int
	failCount int
}

func (f *failreader) Read(b []byte) (int, error) {
	f.failCount++
	if f.failCount >= f.times {
		return 0, fmt.Errorf("random failure")
	}
	return len(b), nil
}

func TestUploadOrderReadFail1(t *testing.T) {
	s, ops, _ := loggingSvc(emptyList)
	mgr := s3manager.NewUploaderWithClient(s)
	_, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   &failreader{times: 1},
	})

	var aerr awserr.Error
	if !errors.As(err, &aerr) {
		t.Fatalf("expect %T error, got %v", aerr, err)
	}

	if e, a := "ReadRequestBody", aerr.Code(); e != a {
		t.Errorf("expect %q code, got %q", e, a)
	}

	origErr := errors.Unwrap(aerr)
	if origErr == nil {
		t.Fatalf("expect wrapped error, was none")
	}
	if e, a := "random failure", origErr.Error(); !strings.Contains(a, e) {
		t.Errorf("expected %q, but received %q", e, a)
	}
	if e, a := []string{}, *ops; !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %v, but received %v", e, a)
	}
}

func TestUploadOrderReadFail2(t *testing.T) {
	s, ops, _ := loggingSvc([]string{"UploadPart"})
	mgr := s3manager.NewUploaderWithClient(s, func(u *s3manager.Uploader) {
		u.Concurrency = 1
	})
	_, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   &failreader{times: 2},
	})

	var aerr awserr.Error
	if !errors.As(err, &aerr) {
		t.Fatalf("expect %T error, got %v", aerr, err)
	}

	if e, a := "MultipartUpload", aerr.Code(); e != a {
		t.Errorf("Expected %q, but received %q", e, a)
	}

	origErr := errors.Unwrap(aerr)
	if origErr == nil {
		t.Fatalf("expect wrapped error, was none")
	}
	if e, a := "ReadRequestBody", origErr.(awserr.Error).Code(); e != a {
		t.Errorf("Expected %q, but received %q", e, a)
	}

	if e, a := "random failure", origErr.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %q error, got %q", e, a)
	}

	if e, a := []string{"CreateMultipartUpload", "AbortMultipartUpload"}, *ops; !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %v, but receievd %v", e, a)
	}
}

type sizedReader struct {
	size int
	cur  int
	err  error
}

func (s *sizedReader) Read(p []byte) (n int, err error) {
	if s.cur >= s.size {
		if s.err == nil {
			s.err = io.EOF
		}
		return 0, s.err
	}

	n = len(p)
	s.cur += len(p)
	if s.cur > s.size {
		n -= s.cur - s.size
	}

	return n, err
}

func TestUploadOrderMultiBufferedReader(t *testing.T) {
	s, ops, args := loggingSvc(emptyList)
	mgr := s3manager.NewUploaderWithClient(s)
	_, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   &sizedReader{size: 1024 * 1024 * 12},
	})

	if err != nil {
		t.Errorf("Expected no error but received %v", err)
	}

	if e, a := []string{"CreateMultipartUpload", "UploadPart", "UploadPart", "UploadPart", "CompleteMultipartUpload"}, *ops; !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %v, but receievd %v", e, a)
	}

	// Part lengths
	parts := []int{
		buflen(val((*args)[1], "Body")),
		buflen(val((*args)[2], "Body")),
		buflen(val((*args)[3], "Body")),
	}
	sort.Ints(parts)

	if e, a := []int{1024 * 1024 * 2, 1024 * 1024 * 5, 1024 * 1024 * 5}, parts; !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %v, but receievd %v", e, a)
	}
}

func TestUploadOrderMultiBufferedReaderPartial(t *testing.T) {
	s, ops, args := loggingSvc(emptyList)
	mgr := s3manager.NewUploaderWithClient(s)
	_, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   &sizedReader{size: 1024 * 1024 * 12, err: io.EOF},
	})

	if err != nil {
		t.Errorf("Expected no error but received %v", err)
	}

	if e, a := []string{"CreateMultipartUpload", "UploadPart", "UploadPart", "UploadPart", "CompleteMultipartUpload"}, *ops; !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %v, but receievd %v", e, a)
	}

	// Part lengths
	parts := []int{
		buflen(val((*args)[1], "Body")),
		buflen(val((*args)[2], "Body")),
		buflen(val((*args)[3], "Body")),
	}
	sort.Ints(parts)

	if e, a := []int{1024 * 1024 * 2, 1024 * 1024 * 5, 1024 * 1024 * 5}, parts; !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %v, but receievd %v", e, a)
	}
}

// TestUploadOrderMultiBufferedReaderEOF tests the edge case where the
// file size is the same as part size.
func TestUploadOrderMultiBufferedReaderEOF(t *testing.T) {
	s, ops, args := loggingSvc(emptyList)
	mgr := s3manager.NewUploaderWithClient(s)
	_, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   &sizedReader{size: 1024 * 1024 * 10, err: io.EOF},
	})

	if err != nil {
		t.Errorf("Expected no error but received %v", err)
	}

	if e, a := []string{"CreateMultipartUpload", "UploadPart", "UploadPart", "CompleteMultipartUpload"}, *ops; !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %v, but receievd %v", e, a)
	}

	// Part lengths
	parts := []int{
		buflen(val((*args)[1], "Body")),
		buflen(val((*args)[2], "Body")),
	}
	sort.Ints(parts)

	if e, a := []int{1024 * 1024 * 5, 1024 * 1024 * 5}, parts; !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %v, but receievd %v", e, a)
	}
}

func TestUploadOrderMultiBufferedReaderExceedTotalParts(t *testing.T) {
	s, ops, _ := loggingSvc([]string{"UploadPart"})
	mgr := s3manager.NewUploaderWithClient(s, func(u *s3manager.Uploader) {
		u.Concurrency = 1
		u.MaxUploadParts = 2
	})
	resp, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   &sizedReader{size: 1024 * 1024 * 12},
	})

	if err == nil {
		t.Error("Expected an error, but received nil")
	}

	if resp != nil {
		t.Errorf("Expected nil, but receievd %v", resp)
	}

	if e, a := []string{"CreateMultipartUpload", "AbortMultipartUpload"}, *ops; !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %v, but receievd %v", e, a)
	}

	aerr := err.(awserr.Error)
	if e, a := "MultipartUpload", aerr.Code(); e != a {
		t.Errorf("Expected %q, but received %q", e, a)
	}

	origErr := errors.Unwrap(aerr)
	if origErr == nil {
		t.Fatalf("expect wrapped error, was none")
	}
	if e, a := "TotalPartsExceeded", origErr.(awserr.Error).Code(); e != a {
		t.Errorf("expect %q, but received %q", e, a)
	}

	if e, a := "configured MaxUploadParts (2)", aerr.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %q, but received %q", e, a)
	}
}

func TestUploadOrderSingleBufferedReader(t *testing.T) {
	s, ops, _ := loggingSvc(emptyList)
	mgr := s3manager.NewUploaderWithClient(s)
	resp, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   &sizedReader{size: 1024 * 1024 * 2},
	})

	if err != nil {
		t.Errorf("Expected no error but received %v", err)
	}

	if e, a := []string{"PutObject"}, *ops; !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %v, but received %v", e, a)
	}

	if len(resp.Location) == 0 {
		t.Error("Expected a value in Location but received empty string")
	}

	if len(resp.UploadID) > 0 {
		t.Errorf("Expected empty string but received %q", resp.UploadID)
	}
}

func TestUploadZeroLenObject(t *testing.T) {
	requestMade := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestMade = true
		w.WriteHeader(http.StatusOK)
	}))

	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL(server.URL)

	mgr := s3manager.NewUploaderWithClient(s3.New(cfg))

	resp, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   strings.NewReader(""),
	})

	if err != nil {
		t.Errorf("Expected no error but received %v", err)
	}
	if !requestMade {
		t.Error("Expected request to have been made, but was not")
	}

	if len(resp.Location) == 0 {
		t.Error("Expected a non-empty string value for Location")
	}

	if len(resp.UploadID) > 0 {
		t.Errorf("Expected empty string, but received %q", resp.UploadID)
	}
}

func TestUploadInputS3PutObjectInputPairity(t *testing.T) {
	matchings := compareStructType(reflect.TypeOf(s3.PutObjectInput{}),
		reflect.TypeOf(s3manager.UploadInput{}))
	aOnly := []string{}
	bOnly := []string{}

	for k, c := range matchings {
		if strings.HasPrefix(k, "Body-") || strings.HasPrefix(k, "ContentLength-") {
			continue
		}
		if c == 1 {
			aOnly = append(aOnly, k)
		} else if c == 2 {
			bOnly = append(bOnly, k)
		}
	}

	if len(aOnly) > 0 {
		t.Errorf("Expected empty array, but received %v", aOnly)
	}

	if len(bOnly) > 0 {
		t.Errorf("Expected empty array, but received %v", bOnly)
	}
}

type testIncompleteReader struct {
	Size int64
	read int64
}

func (r *testIncompleteReader) Read(p []byte) (n int, err error) {
	r.read += int64(len(p))
	if r.read >= r.Size {
		return int(r.read - r.Size), io.ErrUnexpectedEOF
	}
	return len(p), nil
}

func TestUploadUnexpectedEOF(t *testing.T) {
	s, ops, _ := loggingSvc(emptyList)
	mgr := s3manager.NewUploaderWithClient(s, func(u *s3manager.Uploader) {
		u.Concurrency = 1
	})
	_, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body: &testIncompleteReader{
			Size: s3manager.MinUploadPartSize + 1,
		},
	})

	if err == nil {
		t.Error("Expected error, but received none")
	}

	// Ensure upload started.
	if e, a := "CreateMultipartUpload", (*ops)[0]; e != a {
		t.Errorf("Expected %q, but received %q", e, a)
	}

	// Part may or may not be sent due to timing of sending the current part
	// and reading next part in upload manager. Only check for the abort.
	if e, a := "AbortMultipartUpload", (*ops)[len(*ops)-1]; e != a {
		t.Errorf("Expected %q, but received %q", e, a)
	}
}

func compareStructType(a, b reflect.Type) map[string]int {
	if a.Kind() != reflect.Struct || b.Kind() != reflect.Struct {
		panic(fmt.Sprintf("types must both be structs, got %v and %v", a.Kind(), b.Kind()))
	}

	aFields := enumFields(a)
	bFields := enumFields(b)

	matchings := map[string]int{}

	for i := 0; i < len(aFields) || i < len(bFields); i++ {
		if i < len(aFields) {
			name := matchName(aFields[i])
			c := matchings[name]
			matchings[name] = c + 1
		}
		if i < len(bFields) {
			name := matchName(bFields[i])
			c := matchings[name]
			matchings[name] = c + 2
		}
	}

	return matchings
}

func matchName(field reflect.StructField) string {
	return field.Name + "-" + field.Type.String()
}

func enumFields(v reflect.Type) []reflect.StructField {
	fields := []reflect.StructField{}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		// Ignoreing anon fields
		if field.PkgPath != "" {
			// Ignore unexported fields
			continue
		}

		fields = append(fields, field)
	}

	return fields
}

type fooReaderAt struct{}

func (r *fooReaderAt) Read(p []byte) (n int, err error) {
	return 12, io.EOF
}

func (r *fooReaderAt) ReadAt(p []byte, off int64) (n int, err error) {
	return 12, io.EOF
}

func TestReaderAt(t *testing.T) {
	svc := s3.New(unit.Config())
	svc.Handlers.Unmarshal.Clear()
	svc.Handlers.UnmarshalMeta.Clear()
	svc.Handlers.UnmarshalError.Clear()
	svc.Handlers.Send.Clear()

	contentLen := ""
	svc.Handlers.Send.PushBack(func(r *aws.Request) {
		contentLen = r.HTTPRequest.Header.Get("Content-Length")
		r.HTTPResponse = &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(respBody))),
		}
	})

	mgr := s3manager.NewUploaderWithClient(svc, func(u *s3manager.Uploader) {
		u.Concurrency = 1
	})

	_, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   &fooReaderAt{},
	})

	if err != nil {
		t.Errorf("Expected no error but received %v", err)
	}

	if e, a := "12", contentLen; e != a {
		t.Errorf("Expected %q, but received %q", e, a)
	}
}

func TestSSE(t *testing.T) {
	svc := s3.New(unit.Config())
	svc.Handlers.Unmarshal.Clear()
	svc.Handlers.UnmarshalMeta.Clear()
	svc.Handlers.UnmarshalError.Clear()
	svc.Handlers.ValidateResponse.Clear()
	svc.Handlers.Send.Clear()
	partNum := 0
	mutex := &sync.Mutex{}

	svc.Handlers.Send.PushBack(func(r *aws.Request) {
		mutex.Lock()
		defer mutex.Unlock()
		r.HTTPResponse = &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(respBody))),
		}
		switch data := r.Data.(type) {
		case *s3.CreateMultipartUploadOutput:
			data.UploadId = aws.String("UPLOAD-ID")
		case *s3.UploadPartOutput:
			input := r.Params.(*s3.UploadPartInput)
			if input.SSECustomerAlgorithm == nil {
				t.Fatal("SSECustomerAlgoritm should not be nil")
			}
			if input.SSECustomerKey == nil {
				t.Fatal("SSECustomerKey should not be nil")
			}
			partNum++
			data.ETag = aws.String(fmt.Sprintf("ETAG%d", partNum))
		case *s3.CompleteMultipartUploadOutput:
			data.Location = aws.String("https://location")
			data.VersionId = aws.String("VERSION-ID")
		case *s3.PutObjectOutput:
			data.VersionId = aws.String("VERSION-ID")
		}

	})

	mgr := s3manager.NewUploaderWithClient(svc, func(u *s3manager.Uploader) {
		u.Concurrency = 5
	})

	_, err := mgr.Upload(&s3manager.UploadInput{
		Bucket:               aws.String("Bucket"),
		Key:                  aws.String("Key"),
		SSECustomerAlgorithm: aws.String("AES256"),
		SSECustomerKey:       aws.String("foo"),
		Body:                 bytes.NewBuffer(make([]byte, 1024*1024*10)),
	})

	if err != nil {
		t.Fatal("Expected no error, but received" + err.Error())
	}
}

func TestUploadWithContextCanceled(t *testing.T) {
	u := s3manager.NewUploader(unit.Config())

	params := s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   bytes.NewReader(make([]byte, 0)),
	}

	ctx := &awstesting.FakeContext{DoneCh: make(chan struct{})}
	ctx.Error = fmt.Errorf("context canceled")
	close(ctx.DoneCh)

	_, err := u.UploadWithContext(ctx, &params)
	if err == nil {
		t.Fatalf("expected error, did not get one")
	}

	var rc *aws.RequestCanceledError
	if !errors.As(err, &rc) {
		t.Fatalf("expect %T error, got %v", rc, err)
	}
	if e, a := "canceled", rc.Err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v to be in, %v", e, a)
	}
}

func TestResumeUploadNotFound(t *testing.T) {
	s, ops, _ := loggingSvc(emptyList)
	mgr := s3manager.NewUploaderWithClient(s)
	// s.Handlers.Send.PushBack()
	s.Handlers.Send.PushBack(func(r *aws.Request) {
		switch r.Data.(type) {
		case *s3.ListPartsOutput:
			r.HTTPResponse.StatusCode = 404
		}
	})
	_, err := mgr.ResumeUpload("UPLOAD-ID", &s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   bytes.NewReader(buf12MB),
	})

	if err == nil {
		t.Fatalf("expected error, did not get one")
	}

	aerr := err.(awserr.Error)
	if e, a := "list upload parts failed", aerr.Message(); !strings.Contains(a, e) {
		t.Errorf("expected error message to contain %q, but did not %q", e, a)
	}

	vals := []string{"ListParts"}
	if !reflect.DeepEqual(vals, *ops) {
		t.Errorf("Expected %v, but received %v", vals, *ops)
	}
}

func TestResumeUploadNoPartsUploaded(t *testing.T) {
	s, ops, _ := loggingSvc(emptyList)
	mgr := s3manager.NewUploaderWithClient(s, func(u *s3manager.Uploader) {
		u.PartSize = 1024 * 1024 * 7
		u.Concurrency = 1
	})
	s.Handlers.Send.PushBack(func(r *aws.Request) {
		switch data := r.Data.(type) {
		case *s3.ListPartsOutput:
			var f bool
			data.IsTruncated = &f
			data.Parts = []s3.Part{}
		}
	})
	_, err := mgr.ResumeUpload("UPLOAD-ID", &s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   bytes.NewReader(buf12MB),
	})

	if err != nil {
		t.Errorf("Expected no error but received %v", err)
	}

	vals := []string{"ListParts", "UploadPart", "UploadPart", "CompleteMultipartUpload"}
	if !reflect.DeepEqual(vals, *ops) {
		t.Errorf("Expected %v, but received %v", vals, *ops)
	}
}

func TestResumeUploadSomePartsUploaded(t *testing.T) {
	s, ops, _ := loggingSvc(emptyList)
	mgr := s3manager.NewUploaderWithClient(s, func(u *s3manager.Uploader) {
		u.Concurrency = 1
	})
	s.Handlers.Send.PushBack(func(r *aws.Request) {
		switch data := r.Data.(type) {
		case *s3.ListPartsOutput:
			var f bool
			data.IsTruncated = &f
			data.Parts = []s3.Part{
				{ETag: aws.String("\"5f363e0e58a95f06cbe9bbc662c5dfb6\""), PartNumber: aws.Int64(1), Size: aws.Int64(1024 * 1024 * 5)},
				{ETag: aws.String("\"5f363e0e58a95f06cbe9bbc662c5dfb6\""), PartNumber: aws.Int64(2), Size: aws.Int64(1024 * 1024 * 5)},
			}
		}
	})
	_, err := mgr.ResumeUpload("UPLOAD-ID", &s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   bytes.NewReader(buf12MB),
	})

	if err != nil {
		t.Errorf("Expected no error but received %v", err)
	}

	vals := []string{"ListParts", "UploadPart", "CompleteMultipartUpload"}
	if !reflect.DeepEqual(vals, *ops) {
		t.Errorf("Expected %v, but received %v", vals, *ops)
	}
}

func TestResumeUploadAllPartsUploaded(t *testing.T) {
	s, ops, _ := loggingSvc(emptyList)
	mgr := s3manager.NewUploaderWithClient(s, func(u *s3manager.Uploader) {
		u.Concurrency = 1
	})
	s.Handlers.Send.PushBack(func(r *aws.Request) {
		switch data := r.Data.(type) {
		case *s3.ListPartsOutput:
			var f bool
			data.IsTruncated = &f
			data.Parts = []s3.Part{
				{ETag: aws.String("\"5f363e0e58a95f06cbe9bbc662c5dfb6\""), PartNumber: aws.Int64(1), Size: aws.Int64(1024 * 1024 * 5)},
				{ETag: aws.String("\"5f363e0e58a95f06cbe9bbc662c5dfb6\""), PartNumber: aws.Int64(2), Size: aws.Int64(1024 * 1024 * 5)},
				{ETag: aws.String("\"b2d1236c286a3c0704224fe4105eca49\""), PartNumber: aws.Int64(3), Size: aws.Int64(1024 * 1024 * 2)},
			}
		}
	})
	_, err := mgr.ResumeUpload("UPLOAD-ID", &s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   bytes.NewReader(buf12MB),
	})

	if err != nil {
		t.Errorf("Expected no error but received %v", err)
	}

	vals := []string{"ListParts", "CompleteMultipartUpload"}
	if !reflect.DeepEqual(vals, *ops) {
		t.Errorf("Expected %v, but received %v", vals, *ops)
	}
}

func TestResumeUploadUsesPreviousPartSize(t *testing.T) {
	s, ops, _ := loggingSvc(emptyList)
	partSize := 1024 * 1024 * 5
	mgr := s3manager.NewUploaderWithClient(s, func(u *s3manager.Uploader) {
		u.PartSize = 1024 * 1024 * 100
		u.Concurrency = 1
	})
	s.Handlers.Send.PushBack(func(r *aws.Request) {
		switch data := r.Data.(type) {
		case *s3.UploadPartInput:
			d, _ := ioutil.ReadAll(data.Body)
			if len(d) > partSize {
				t.Errorf("expected part size %v, but received %v", partSize, len(d))
			}
		case *s3.ListPartsOutput:
			data.IsTruncated = aws.Bool(false)
			data.Parts = []s3.Part{
				{ETag: aws.String("\"5f363e0e58a95f06cbe9bbc662c5dfb6\""), PartNumber: aws.Int64(1), Size: aws.Int64(1024 * 1024 * 5)},
			}
		}
	})
	_, err := mgr.ResumeUpload("UPLOAD-ID", &s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   bytes.NewReader(buf12MB),
	})

	if err != nil {
		t.Errorf("Expected no error but received %v", err)
	}

	vals := []string{"ListParts", "UploadPart", "UploadPart", "CompleteMultipartUpload"}
	if !reflect.DeepEqual(vals, *ops) {
		t.Errorf("Expected %v, but received %v", vals, *ops)
	}
}

func TestResumeUploadETagMismatch(t *testing.T) {
	s, ops, _ := loggingSvc(emptyList)
	mgr := s3manager.NewUploaderWithClient(s, func(u *s3manager.Uploader) {
		u.PartSize = 1024 * 1024 * 5
		u.Concurrency = 1
	})
	s.Handlers.Send.PushBack(func(r *aws.Request) {
		switch data := r.Data.(type) {
		case *s3.ListPartsOutput:
			var f bool
			data.IsTruncated = &f
			data.Parts = []s3.Part{
				{ETag: aws.String("\"5f363e0e58a95f06cbe9bbc662c5dfb6\""), PartNumber: aws.Int64(1), Size: aws.Int64(1024 * 1024 * 5)},
				{ETag: aws.String("\"different-etag\""), PartNumber: aws.Int64(2), Size: aws.Int64(1)},
			}
		}
	})
	_, err := mgr.ResumeUpload("UPLOAD-ID", &s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   bytes.NewReader(buf12MB),
	})

	if err != nil {
		t.Errorf("Expected no error but received %v", err)
	}

	vals := []string{"ListParts", "UploadPart", "UploadPart", "CompleteMultipartUpload"}
	if !reflect.DeepEqual(vals, *ops) {
		t.Errorf("Expected %v, but received %v", vals, *ops)
	}
}

func TestUploadPartCallbackCancelsUpload(t *testing.T) {
	s, ops, _ := loggingSvc(emptyList)
	mgr := s3manager.NewUploaderWithClient(s, func(u *s3manager.Uploader) {
		u.Concurrency = 1
		u.UploadPartCallback = func(part s3.CompletedPart) bool {
			return *part.PartNumber != 2
		}
	})
	_, err := mgr.Upload(&s3manager.UploadInput{
		Bucket: aws.String("Bucket"),
		Key:    aws.String("Key"),
		Body:   bytes.NewReader(buf12MB),
	})

	if err == nil {
		t.Errorf("Expected error, but received nil")
	}

	aerr := err.(s3manager.MultiUploadFailure)
	if e, a := s3manager.ErrMultiUploadCanceled, aerr.OrigErr(); e != a {
		t.Errorf("Expected %v, but received %v", e, a)
	}

	vals := []string{"CreateMultipartUpload", "UploadPart", "UploadPart", "AbortMultipartUpload"}
	if !reflect.DeepEqual(vals, *ops) {
		t.Errorf("Expected %v, but received %v", vals, *ops)
	}
}

func TestUploadBufferStrategy(t *testing.T) {
	cases := map[string]struct {
		PartSize  int64
		Size      int64
		Strategy  s3manager.ReadSeekerWriteToProvider
		callbacks int
	}{
		"NoBuffer": {
			PartSize: s3manager.DefaultUploadPartSize,
			Strategy: nil,
		},
		"SinglePart": {
			PartSize:  s3manager.DefaultUploadPartSize,
			Size:      s3manager.DefaultUploadPartSize,
			Strategy:  &recordedBufferProvider{size: int(s3manager.DefaultUploadPartSize)},
			callbacks: 1,
		},
		"MultiPart": {
			PartSize:  s3manager.DefaultUploadPartSize,
			Size:      s3manager.DefaultUploadPartSize * 2,
			Strategy:  &recordedBufferProvider{size: int(s3manager.DefaultUploadPartSize)},
			callbacks: 2,
		},
	}

	for name, tCase := range cases {
		t.Run(name, func(t *testing.T) {
			_ = tCase
			cfg := unit.Config()
			svc := s3.New(cfg)
			svc.Handlers.Unmarshal.Clear()
			svc.Handlers.UnmarshalMeta.Clear()
			svc.Handlers.UnmarshalError.Clear()
			svc.Handlers.Send.Clear()

			var etag int64
			svc.Handlers.Send.PushBack(func(r *aws.Request) {
				if r.Body != nil {
					io.Copy(ioutil.Discard, r.Body)
				}

				r.HTTPResponse = &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(respBody))),
				}

				switch data := r.Data.(type) {
				case *s3.CreateMultipartUploadOutput:
					data.UploadId = aws.String("UPLOAD-ID")
				case *s3.UploadPartOutput:
					data.ETag = aws.String(fmt.Sprintf("ETAG%d", atomic.AddInt64(&etag, 1)))
				case *s3.CompleteMultipartUploadOutput:
					data.Location = aws.String("https://location")
					data.VersionId = aws.String("VERSION-ID")
				case *s3.PutObjectOutput:
					data.VersionId = aws.String("VERSION-ID")
				}
			})

			uploader := s3manager.NewUploaderWithClient(svc, func(u *s3manager.Uploader) {
				u.PartSize = tCase.PartSize
				u.BufferProvider = tCase.Strategy
				u.Concurrency = 1
			})

			expected := s3testing.GetTestBytes(int(tCase.Size))
			_, err := uploader.Upload(&s3manager.UploadInput{
				Bucket: aws.String("bucket"),
				Key:    aws.String("key"),
				Body:   bytes.NewReader(expected),
			})
			if err != nil {
				t.Fatalf("failed to upload file: %v", err)
			}

			switch strat := tCase.Strategy.(type) {
			case *recordedBufferProvider:
				if !bytes.Equal(expected, strat.content) {
					t.Errorf("content buffered did not match expected")
				}
				if tCase.callbacks != strat.callbackCount {
					t.Errorf("expected %v, got %v callbacks", tCase.callbacks, strat.callbackCount)
				}
			}
		})
	}
}

type recordedBufferProvider struct {
	content       []byte
	size          int
	callbackCount int
}

func (r *recordedBufferProvider) GetWriteTo(seeker io.ReadSeeker) (s3manager.ReadSeekerWriteTo, func()) {
	b := make([]byte, r.size)
	w := &s3manager.BufferedReadSeekerWriteTo{BufferedReadSeeker: s3manager.NewBufferedReadSeeker(seeker, b)}

	return w, func() {
		r.content = append(r.content, b...)
		r.callbackCount++
	}
}
