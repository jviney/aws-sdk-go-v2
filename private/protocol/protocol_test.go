package protocol_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awstesting"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/query"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/rest"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/restxml"
)

func xmlData(set bool, b []byte, size, delta int) {
	const openingTags = "<B><A>"
	const closingTags = "</A></B>"
	if !set {
		copy(b, []byte(openingTags))
	}
	if size == 0 {
		copy(b[delta-len(closingTags):], []byte(closingTags))
	}
}

func jsonData(set bool, b []byte, size, delta int) {
	if !set {
		copy(b, []byte("{\"A\": \""))
	}
	if size == 0 {
		copy(b[delta-len("\"}"):], []byte("\"}"))
	}
}

func buildNewRequest(data interface{}) *aws.Request {
	v := url.Values{}
	v.Set("test", "TEST")
	v.Add("test1", "TEST1")

	req := &aws.Request{
		HTTPRequest: &http.Request{
			Header: make(http.Header),
			Body:   &awstesting.ReadCloser{Size: 2048},
			URL: &url.URL{
				RawQuery: v.Encode(),
			},
		},
		Params: &struct {
			LocationName string `locationName:"test"`
		}{
			"Test",
		},
		Endpoint: aws.Endpoint{URL: "test"},
		Metadata: aws.Metadata{
			ServiceName:   "test",
			TargetPrefix:  "test",
			JSONVersion:   "test",
			APIVersion:    "test",
			SigningName:   "test",
			SigningRegion: "test",
		},
		Operation: &aws.Operation{
			Name: "test",
		},
	}
	req.HTTPResponse = &http.Response{
		Body: &awstesting.ReadCloser{Size: 2048},
		Header: http.Header{
			"X-Amzn-Requestid": []string{"1"},
		},
		StatusCode: http.StatusOK,
	}

	if data == nil {
		data = &struct {
			_            struct{} `type:"structure"`
			LocationName *string  `locationName:"testName"`
			Location     *string  `location:"statusCode"`
			A            *string  `type:"string"`
		}{}
	}

	req.Data = data

	return req
}

type expected struct {
	dataType  int
	closed    bool
	size      int
	errExists bool
}

const (
	jsonType = iota
	xmlType
)

func checkForLeak(data interface{}, build, fn func(*aws.Request), t *testing.T, result expected) {
	req := buildNewRequest(data)
	reader := req.HTTPResponse.Body.(*awstesting.ReadCloser)
	switch result.dataType {
	case jsonType:
		reader.FillData = jsonData
	case xmlType:
		reader.FillData = xmlData
	}
	build(req)
	fn(req)

	if result.errExists {
		if req.Error == nil {
			t.Fatalf("expect error, got none")
		}
	} else {
		if req.Error != nil {
			t.Fatalf("expect no error, got %v", req.Error)
		}
	}

	if e, a := reader.Closed, result.closed; e != a {
		t.Errorf("expect %v closed, got %v", e, a)
	}
	if e, a := reader.Size, result.size; e != a {
		t.Errorf("expect %v size, got %v", e, a)
	}
}

func TestJSONRpc(t *testing.T) {
	checkForLeak(nil, jsonrpc.Build, jsonrpc.Unmarshal, t, expected{jsonType, true, 0, false})
	checkForLeak(nil, jsonrpc.Build, jsonrpc.UnmarshalMeta, t, expected{jsonType, false, 2048, false})
	checkForLeak(nil, jsonrpc.Build, jsonrpc.UnmarshalError, t, expected{jsonType, true, 0, true})
}

func TestQuery(t *testing.T) {
	checkForLeak(nil, query.Build, query.Unmarshal, t, expected{jsonType, true, 0, false})
	checkForLeak(nil, query.Build, query.UnmarshalMeta, t, expected{jsonType, false, 2048, false})
	checkForLeak(nil, query.Build, query.UnmarshalError, t, expected{jsonType, true, 0, true})
}

func TestRest(t *testing.T) {
	// case 1: Payload io.ReadSeeker
	checkForLeak(nil, rest.Build, rest.Unmarshal, t, expected{jsonType, false, 2048, false})
	checkForLeak(nil, query.Build, query.UnmarshalMeta, t, expected{jsonType, false, 2048, false})

	// case 2: Payload *string
	// should close the body
	dataStr := struct {
		_            struct{} `type:"structure" payload:"Payload"`
		LocationName *string  `locationName:"testName"`
		Location     *string  `location:"statusCode"`
		A            *string  `type:"string"`
		Payload      *string  `locationName:"payload" type:"blob" required:"true"`
	}{}
	checkForLeak(&dataStr, rest.Build, rest.Unmarshal, t, expected{jsonType, true, 0, false})
	checkForLeak(&dataStr, query.Build, query.UnmarshalMeta, t, expected{jsonType, false, 2048, false})

	// case 3: Payload []byte
	// should close the body
	dataBytes := struct {
		_            struct{} `type:"structure" payload:"Payload"`
		LocationName *string  `locationName:"testName"`
		Location     *string  `location:"statusCode"`
		A            *string  `type:"string"`
		Payload      []byte   `locationName:"payload" type:"blob" required:"true"`
	}{}
	checkForLeak(&dataBytes, rest.Build, rest.Unmarshal, t, expected{jsonType, true, 0, false})
	checkForLeak(&dataBytes, query.Build, query.UnmarshalMeta, t, expected{jsonType, false, 2048, false})

	// case 4: Payload unsupported type
	// should close the body
	dataUnsupported := struct {
		_            struct{} `type:"structure" payload:"Payload"`
		LocationName *string  `locationName:"testName"`
		Location     *string  `location:"statusCode"`
		A            *string  `type:"string"`
		Payload      string   `locationName:"payload" type:"blob" required:"true"`
	}{}
	checkForLeak(&dataUnsupported, rest.Build, rest.Unmarshal, t, expected{jsonType, true, 0, true})
	checkForLeak(&dataUnsupported, query.Build, query.UnmarshalMeta, t, expected{jsonType, false, 2048, false})
}

func TestRestJSON(t *testing.T) {
	checkForLeak(nil, restjson.Build, restjson.Unmarshal, t, expected{jsonType, true, 0, false})
	checkForLeak(nil, restjson.Build, restjson.UnmarshalMeta, t, expected{jsonType, false, 2048, false})
	checkForLeak(nil, restjson.Build, restjson.UnmarshalError, t, expected{jsonType, true, 0, true})
}

func TestRestXML(t *testing.T) {
	checkForLeak(nil, restxml.Build, restxml.Unmarshal, t, expected{xmlType, true, 0, false})
	checkForLeak(nil, restxml.Build, restxml.UnmarshalMeta, t, expected{xmlType, false, 2048, false})
	checkForLeak(nil, restxml.Build, restxml.UnmarshalError, t, expected{xmlType, true, 0, true})
}

func TestXML(t *testing.T) {
	checkForLeak(nil, ec2query.Build, ec2query.Unmarshal, t, expected{jsonType, true, 0, false})
	checkForLeak(nil, ec2query.Build, ec2query.UnmarshalMeta, t, expected{jsonType, false, 2048, false})
	checkForLeak(nil, ec2query.Build, ec2query.UnmarshalError, t, expected{jsonType, true, 0, true})
}

func TestProtocol(t *testing.T) {
	checkForLeak(nil, restxml.Build, protocol.UnmarshalDiscardBody, t, expected{xmlType, true, 0, false})
}
