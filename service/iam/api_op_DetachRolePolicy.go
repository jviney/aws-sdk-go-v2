// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/query"
)

type DetachRolePolicyInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the IAM policy you want to detach.
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	//
	// PolicyArn is a required field
	PolicyArn *string `min:"20" type:"string" required:"true"`

	// The name (friendly name, not ARN) of the IAM role to detach the policy from.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// RoleName is a required field
	RoleName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DetachRolePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachRolePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetachRolePolicyInput"}

	if s.PolicyArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyArn"))
	}
	if s.PolicyArn != nil && len(*s.PolicyArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyArn", 20))
	}

	if s.RoleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleName"))
	}
	if s.RoleName != nil && len(*s.RoleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DetachRolePolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DetachRolePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDetachRolePolicy = "DetachRolePolicy"

// DetachRolePolicyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Removes the specified managed policy from the specified role.
//
// A role can also have inline policies embedded with it. To delete an inline
// policy, use the DeleteRolePolicy API. For information about policies, see
// Managed Policies and Inline Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
//
//    // Example sending a request using DetachRolePolicyRequest.
//    req := client.DetachRolePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DetachRolePolicy
func (c *Client) DetachRolePolicyRequest(input *DetachRolePolicyInput) DetachRolePolicyRequest {
	op := &aws.Operation{
		Name:       opDetachRolePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachRolePolicyInput{}
	}

	req := c.newRequest(op, input, &DetachRolePolicyOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DetachRolePolicyRequest{Request: req, Input: input, Copy: c.DetachRolePolicyRequest}
}

// DetachRolePolicyRequest is the request type for the
// DetachRolePolicy API operation.
type DetachRolePolicyRequest struct {
	*aws.Request
	Input *DetachRolePolicyInput
	Copy  func(*DetachRolePolicyInput) DetachRolePolicyRequest
}

// Send marshals and sends the DetachRolePolicy API request.
func (r DetachRolePolicyRequest) Send(ctx context.Context) (*DetachRolePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachRolePolicyResponse{
		DetachRolePolicyOutput: r.Request.Data.(*DetachRolePolicyOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachRolePolicyResponse is the response type for the
// DetachRolePolicy API operation.
type DetachRolePolicyResponse struct {
	*DetachRolePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachRolePolicy request.
func (r *DetachRolePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
