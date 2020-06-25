// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeregisterResourceInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the resource that you want to deregister.
	//
	// ResourceArn is a required field
	ResourceArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeregisterResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeregisterResourceInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeregisterResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeregisterResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeregisterResource = "DeregisterResource"

// DeregisterResourceRequest returns a request value for making API operation for
// AWS Lake Formation.
//
// Deregisters the resource as managed by the Data Catalog.
//
// When you deregister a path, Lake Formation removes the path from the inline
// policy attached to your service-linked role.
//
//    // Example sending a request using DeregisterResourceRequest.
//    req := client.DeregisterResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lakeformation-2017-03-31/DeregisterResource
func (c *Client) DeregisterResourceRequest(input *DeregisterResourceInput) DeregisterResourceRequest {
	op := &aws.Operation{
		Name:       opDeregisterResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeregisterResourceInput{}
	}

	req := c.newRequest(op, input, &DeregisterResourceOutput{})

	return DeregisterResourceRequest{Request: req, Input: input, Copy: c.DeregisterResourceRequest}
}

// DeregisterResourceRequest is the request type for the
// DeregisterResource API operation.
type DeregisterResourceRequest struct {
	*aws.Request
	Input *DeregisterResourceInput
	Copy  func(*DeregisterResourceInput) DeregisterResourceRequest
}

// Send marshals and sends the DeregisterResource API request.
func (r DeregisterResourceRequest) Send(ctx context.Context) (*DeregisterResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterResourceResponse{
		DeregisterResourceOutput: r.Request.Data.(*DeregisterResourceOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterResourceResponse is the response type for the
// DeregisterResource API operation.
type DeregisterResourceResponse struct {
	*DeregisterResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterResource request.
func (r *DeregisterResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
