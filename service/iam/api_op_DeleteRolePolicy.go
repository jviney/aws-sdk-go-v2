// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/query"
)

type DeleteRolePolicyInput struct {
	_ struct{} `type:"structure"`

	// The name of the inline policy to delete from the specified IAM role.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// PolicyName is a required field
	PolicyName *string `min:"1" type:"string" required:"true"`

	// The name (friendly name, not ARN) identifying the role that the policy is
	// embedded in.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// RoleName is a required field
	RoleName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRolePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRolePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRolePolicyInput"}

	if s.PolicyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyName"))
	}
	if s.PolicyName != nil && len(*s.PolicyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyName", 1))
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

type DeleteRolePolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRolePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteRolePolicy = "DeleteRolePolicy"

// DeleteRolePolicyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Deletes the specified inline policy that is embedded in the specified IAM
// role.
//
// A role can also have managed policies attached to it. To detach a managed
// policy from a role, use DetachRolePolicy. For more information about policies,
// refer to Managed Policies and Inline Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
//
//    // Example sending a request using DeleteRolePolicyRequest.
//    req := client.DeleteRolePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeleteRolePolicy
func (c *Client) DeleteRolePolicyRequest(input *DeleteRolePolicyInput) DeleteRolePolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteRolePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRolePolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteRolePolicyOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteRolePolicyRequest{Request: req, Input: input, Copy: c.DeleteRolePolicyRequest}
}

// DeleteRolePolicyRequest is the request type for the
// DeleteRolePolicy API operation.
type DeleteRolePolicyRequest struct {
	*aws.Request
	Input *DeleteRolePolicyInput
	Copy  func(*DeleteRolePolicyInput) DeleteRolePolicyRequest
}

// Send marshals and sends the DeleteRolePolicy API request.
func (r DeleteRolePolicyRequest) Send(ctx context.Context) (*DeleteRolePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRolePolicyResponse{
		DeleteRolePolicyOutput: r.Request.Data.(*DeleteRolePolicyOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRolePolicyResponse is the response type for the
// DeleteRolePolicy API operation.
type DeleteRolePolicyResponse struct {
	*DeleteRolePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRolePolicy request.
func (r *DeleteRolePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
