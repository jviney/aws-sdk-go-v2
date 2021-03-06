package simpledb_test

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/aws/awserr"
	"github.com/jviney/aws-sdk-go-v2/internal/awstesting/unit"
	"github.com/jviney/aws-sdk-go-v2/service/simpledb"
)

var statusCodeErrorTests = []struct {
	scode   int
	status  string
	code    string
	message string
}{
	{301, "Moved Permanently", "MovedPermanently", "Moved Permanently"},
	{403, "Forbidden", "Forbidden", "Forbidden"},
	{400, "Bad Request", "BadRequest", "Bad Request"},
	{404, "Not Found", "NotFound", "Not Found"},
	{500, "Internal Error", "InternalError", "Internal Error"},
}

func TestStatusCodeError(t *testing.T) {
	for _, test := range statusCodeErrorTests {
		s := simpledb.New(unit.Config())
		s.Handlers.Send.Clear()
		s.Handlers.Send.PushBack(func(r *aws.Request) {
			body := ioutil.NopCloser(bytes.NewReader([]byte{}))
			r.HTTPResponse = &http.Response{
				ContentLength: 0,
				StatusCode:    test.scode,
				Status:        test.status,
				Body:          body,
			}
		})
		req := s.CreateDomainRequest(&simpledb.CreateDomainInput{
			DomainName: aws.String("test-domain"),
		})
		_, err := req.Send(context.Background())

		if err == nil {
			t.Fatalf("expect error, got nil")
		}

		var v awserr.Error
		if !errors.As(err, &v) {
			t.Fatalf("expect API error, got %v", err)
		}
		if e, a := test.code, v.Code(); e != a {
			t.Errorf("expect %v, got %v", e, a)
		}
		if e, a := test.message, v.Message(); e != a {
			t.Errorf("expect %v, got %v", e, a)
		}
	}
}

var responseErrorTests = []struct {
	scode     int
	status    string
	code      string
	message   string
	requestID string
	errors    []struct {
		code    string
		message string
	}
}{
	{
		scode:     400,
		status:    "Bad Request",
		code:      "MissingError",
		message:   "missing error code in SimpleDB XML error response",
		requestID: "101",
		errors:    []struct{ code, message string }{},
	},
	{
		scode:     403,
		status:    "Forbidden",
		code:      "AuthFailure",
		message:   "AWS was not able to validate the provided access keys.",
		requestID: "1111",
		errors: []struct{ code, message string }{
			{"AuthFailure", "AWS was not able to validate the provided access keys."},
		},
	},
	{
		scode:     500,
		status:    "Internal Error",
		code:      "MissingParameter",
		message:   "Message #1",
		requestID: "8756",
		errors: []struct{ code, message string }{
			{"MissingParameter", "Message #1"},
			{"InternalError", "Message #2"},
		},
	},
}

func TestResponseError(t *testing.T) {
	for _, test := range responseErrorTests {
		s := simpledb.New(unit.Config())
		s.Handlers.Send.Clear()
		s.Handlers.Send.PushBack(func(r *aws.Request) {
			xml := createXMLResponse(test.requestID, test.errors)
			body := ioutil.NopCloser(bytes.NewReader([]byte(xml)))
			r.HTTPResponse = &http.Response{
				ContentLength: int64(len(xml)),
				StatusCode:    test.scode,
				Status:        test.status,
				Body:          body,
			}
		})
		req := s.CreateDomainRequest(&simpledb.CreateDomainInput{
			DomainName: aws.String("test-domain"),
		})
		_, err := req.Send(context.Background())

		if err == nil {
			t.Fatalf("expect error, got none")
		}

		var v awserr.Error
		if !errors.As(err, &v) {
			t.Fatalf("expect awserr error, got %v", err)
		}
		if e, a := test.code, v.Code(); e != a {
			t.Errorf("expect %v, got %v", e, a)
		}
		if e, a := test.message, v.Message(); e != a {
			t.Errorf("expect %v, got %v", e, a)
		}
		if len(test.errors) > 0 {
			var vv awserr.RequestFailure
			if !errors.As(err, &vv) {
				t.Fatalf("expect RequestFailure error, got %v", err)
			}
			if e, a := test.requestID, vv.RequestID(); e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
			if e, a := test.scode, vv.StatusCode(); e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
		}
	}
}

// createXMLResponse constructs an XML string that has one or more error messages in it.
func createXMLResponse(requestID string, errors []struct{ code, message string }) []byte {
	var buf bytes.Buffer
	buf.WriteString(`<?xml version="1.0"?><Response><Errors>`)
	for _, e := range errors {
		buf.WriteString(`<Error><Code>`)
		buf.WriteString(e.code)
		buf.WriteString(`</Code><Message>`)
		buf.WriteString(e.message)
		buf.WriteString(`</Message></Error>`)
	}
	buf.WriteString(`</Errors><RequestID>`)
	buf.WriteString(requestID)
	buf.WriteString(`</RequestID></Response>`)
	return buf.Bytes()
}
