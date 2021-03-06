// Package query provides serialization of AWS query requests, and responses.
package query

//go:generate go run -tags codegen ../../../models/protocol_tests/generate.go ../../../models/protocol_tests/input/query.json build_test.go

import (
	"net/url"

	request "github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/aws/awserr"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/query/queryutil"
)

// BuildHandler is a named request handler for building query protocol requests
var BuildHandler = request.NamedHandler{Name: "awssdk.query.Build", Fn: Build}

// Build builds a request for an AWS Query service.
func Build(r *request.Request) {
	body := url.Values{
		"Action":  {r.Operation.Name},
		"Version": {r.Metadata.APIVersion},
	}
	if err := queryutil.Parse(body, r.Params, false); err != nil {
		r.Error = awserr.New("SerializationError", "failed encoding Query request", err)
		return
	}

	if r.ExpireTime == 0 {
		r.HTTPRequest.Method = "POST"
		r.HTTPRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
		r.SetBufferBody([]byte(body.Encode()))
	} else { // This is a pre-signed request
		r.HTTPRequest.Method = "GET"
		r.HTTPRequest.URL.RawQuery = body.Encode()
	}
}
