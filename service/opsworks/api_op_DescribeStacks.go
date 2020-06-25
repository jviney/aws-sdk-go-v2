// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeStacksInput struct {
	_ struct{} `type:"structure"`

	// An array of stack IDs that specify the stacks to be described. If you omit
	// this parameter, DescribeStacks returns a description of every stack.
	StackIds []string `type:"list"`
}

// String returns the string representation
func (s DescribeStacksInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the response to a DescribeStacks request.
type DescribeStacksOutput struct {
	_ struct{} `type:"structure"`

	// An array of Stack objects that describe the stacks.
	Stacks []Stack `type:"list"`
}

// String returns the string representation
func (s DescribeStacksOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeStacks = "DescribeStacks"

// DescribeStacksRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Requests a description of one or more stacks.
//
// Required Permissions: To use this action, an IAM user must have a Show, Deploy,
// or Manage permissions level for the stack, or an attached policy that explicitly
// grants permissions. For more information about user permissions, see Managing
// User Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using DescribeStacksRequest.
//    req := client.DescribeStacksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeStacks
func (c *Client) DescribeStacksRequest(input *DescribeStacksInput) DescribeStacksRequest {
	op := &aws.Operation{
		Name:       opDescribeStacks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeStacksInput{}
	}

	req := c.newRequest(op, input, &DescribeStacksOutput{})

	return DescribeStacksRequest{Request: req, Input: input, Copy: c.DescribeStacksRequest}
}

// DescribeStacksRequest is the request type for the
// DescribeStacks API operation.
type DescribeStacksRequest struct {
	*aws.Request
	Input *DescribeStacksInput
	Copy  func(*DescribeStacksInput) DescribeStacksRequest
}

// Send marshals and sends the DescribeStacks API request.
func (r DescribeStacksRequest) Send(ctx context.Context) (*DescribeStacksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStacksResponse{
		DescribeStacksOutput: r.Request.Data.(*DescribeStacksOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeStacksResponse is the response type for the
// DescribeStacks API operation.
type DescribeStacksResponse struct {
	*DescribeStacksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStacks request.
func (r *DescribeStacksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
