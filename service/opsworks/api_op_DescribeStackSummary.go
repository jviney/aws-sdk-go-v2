// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeStackSummaryInput struct {
	_ struct{} `type:"structure"`

	// The stack ID.
	//
	// StackId is a required field
	StackId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeStackSummaryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStackSummaryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStackSummaryInput"}

	if s.StackId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a DescribeStackSummary request.
type DescribeStackSummaryOutput struct {
	_ struct{} `type:"structure"`

	// A StackSummary object that contains the results.
	StackSummary *StackSummary `type:"structure"`
}

// String returns the string representation
func (s DescribeStackSummaryOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeStackSummary = "DescribeStackSummary"

// DescribeStackSummaryRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Describes the number of layers and apps in a specified stack, and the number
// of instances in each state, such as running_setup or online.
//
// Required Permissions: To use this action, an IAM user must have a Show, Deploy,
// or Manage permissions level for the stack, or an attached policy that explicitly
// grants permissions. For more information about user permissions, see Managing
// User Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using DescribeStackSummaryRequest.
//    req := client.DescribeStackSummaryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeStackSummary
func (c *Client) DescribeStackSummaryRequest(input *DescribeStackSummaryInput) DescribeStackSummaryRequest {
	op := &aws.Operation{
		Name:       opDescribeStackSummary,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeStackSummaryInput{}
	}

	req := c.newRequest(op, input, &DescribeStackSummaryOutput{})

	return DescribeStackSummaryRequest{Request: req, Input: input, Copy: c.DescribeStackSummaryRequest}
}

// DescribeStackSummaryRequest is the request type for the
// DescribeStackSummary API operation.
type DescribeStackSummaryRequest struct {
	*aws.Request
	Input *DescribeStackSummaryInput
	Copy  func(*DescribeStackSummaryInput) DescribeStackSummaryRequest
}

// Send marshals and sends the DescribeStackSummary API request.
func (r DescribeStackSummaryRequest) Send(ctx context.Context) (*DescribeStackSummaryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStackSummaryResponse{
		DescribeStackSummaryOutput: r.Request.Data.(*DescribeStackSummaryOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeStackSummaryResponse is the response type for the
// DescribeStackSummary API operation.
type DescribeStackSummaryResponse struct {
	*DescribeStackSummaryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStackSummary request.
func (r *DescribeStackSummaryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
