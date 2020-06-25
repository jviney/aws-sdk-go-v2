// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package awsendpointdiscoverytest

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type TestDiscoveryOptionalInput struct {
	_ struct{} `type:"structure"`

	Sdk *string `type:"string"`
}

// String returns the string representation
func (s TestDiscoveryOptionalInput) String() string {
	return awsutil.Prettify(s)
}

type TestDiscoveryOptionalOutput struct {
	_ struct{} `type:"structure"`

	RequestSuccessful *bool `type:"boolean"`
}

// String returns the string representation
func (s TestDiscoveryOptionalOutput) String() string {
	return awsutil.Prettify(s)
}

const opTestDiscoveryOptional = "TestDiscoveryOptional"

// TestDiscoveryOptionalRequest returns a request value for making API operation for
// AwsEndpointDiscoveryTest.
//
//    // Example sending a request using TestDiscoveryOptionalRequest.
//    req := client.TestDiscoveryOptionalRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) TestDiscoveryOptionalRequest(input *TestDiscoveryOptionalInput) TestDiscoveryOptionalRequest {
	op := &aws.Operation{
		Name:       opTestDiscoveryOptional,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TestDiscoveryOptionalInput{}
	}

	req := c.newRequest(op, input, &TestDiscoveryOptionalOutput{})

	if req.Config.EnableEndpointDiscovery {
		de := discovererDescribeEndpoints{
			Client:        c,
			Required:      false,
			EndpointCache: c.endpointCache,
			Params: map[string]*string{
				"op": &req.Operation.Name,
			},
		}

		for k, v := range de.Params {
			if v == nil {
				delete(de.Params, k)
			}
		}

		req.Handlers.Build.PushFrontNamed(aws.NamedHandler{
			Name: "crr.endpointdiscovery",
			Fn:   de.Handler,
		})
	}

	return TestDiscoveryOptionalRequest{Request: req, Input: input, Copy: c.TestDiscoveryOptionalRequest}
}

// TestDiscoveryOptionalRequest is the request type for the
// TestDiscoveryOptional API operation.
type TestDiscoveryOptionalRequest struct {
	*aws.Request
	Input *TestDiscoveryOptionalInput
	Copy  func(*TestDiscoveryOptionalInput) TestDiscoveryOptionalRequest
}

// Send marshals and sends the TestDiscoveryOptional API request.
func (r TestDiscoveryOptionalRequest) Send(ctx context.Context) (*TestDiscoveryOptionalResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TestDiscoveryOptionalResponse{
		TestDiscoveryOptionalOutput: r.Request.Data.(*TestDiscoveryOptionalOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TestDiscoveryOptionalResponse is the response type for the
// TestDiscoveryOptional API operation.
type TestDiscoveryOptionalResponse struct {
	*TestDiscoveryOptionalOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TestDiscoveryOptional request.
func (r *TestDiscoveryOptionalResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
