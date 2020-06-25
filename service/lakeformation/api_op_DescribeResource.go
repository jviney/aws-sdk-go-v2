// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeResourceInput struct {
	_ struct{} `type:"structure"`

	// The resource ARN.
	//
	// ResourceArn is a required field
	ResourceArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeResourceInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeResourceOutput struct {
	_ struct{} `type:"structure"`

	// A structure containing information about an AWS Lake Formation resource.
	ResourceInfo *ResourceInfo `type:"structure"`
}

// String returns the string representation
func (s DescribeResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeResource = "DescribeResource"

// DescribeResourceRequest returns a request value for making API operation for
// AWS Lake Formation.
//
// Retrieves the current data access role for the given resource registered
// in AWS Lake Formation.
//
//    // Example sending a request using DescribeResourceRequest.
//    req := client.DescribeResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lakeformation-2017-03-31/DescribeResource
func (c *Client) DescribeResourceRequest(input *DescribeResourceInput) DescribeResourceRequest {
	op := &aws.Operation{
		Name:       opDescribeResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeResourceInput{}
	}

	req := c.newRequest(op, input, &DescribeResourceOutput{})

	return DescribeResourceRequest{Request: req, Input: input, Copy: c.DescribeResourceRequest}
}

// DescribeResourceRequest is the request type for the
// DescribeResource API operation.
type DescribeResourceRequest struct {
	*aws.Request
	Input *DescribeResourceInput
	Copy  func(*DescribeResourceInput) DescribeResourceRequest
}

// Send marshals and sends the DescribeResource API request.
func (r DescribeResourceRequest) Send(ctx context.Context) (*DescribeResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeResourceResponse{
		DescribeResourceOutput: r.Request.Data.(*DescribeResourceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeResourceResponse is the response type for the
// DescribeResource API operation.
type DescribeResourceResponse struct {
	*DescribeResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeResource request.
func (r *DescribeResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
