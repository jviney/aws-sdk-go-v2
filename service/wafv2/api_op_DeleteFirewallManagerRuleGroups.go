// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafv2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteFirewallManagerRuleGroupsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the web ACL.
	//
	// WebACLArn is a required field
	WebACLArn *string `min:"20" type:"string" required:"true"`

	// A token used for optimistic locking. AWS WAF returns a token to your get
	// and list requests, to mark the state of the entity at the time of the request.
	// To make changes to the entity associated with the token, you provide the
	// token to operations like update and delete. AWS WAF uses the token to ensure
	// that no changes have been made to the entity since you last retrieved it.
	// If a change has been made, the update fails with a WAFOptimisticLockException.
	// If this happens, perform another get, and use the new token returned by that
	// operation.
	//
	// WebACLLockToken is a required field
	WebACLLockToken *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteFirewallManagerRuleGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFirewallManagerRuleGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFirewallManagerRuleGroupsInput"}

	if s.WebACLArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("WebACLArn"))
	}
	if s.WebACLArn != nil && len(*s.WebACLArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("WebACLArn", 20))
	}

	if s.WebACLLockToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("WebACLLockToken"))
	}
	if s.WebACLLockToken != nil && len(*s.WebACLLockToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WebACLLockToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteFirewallManagerRuleGroupsOutput struct {
	_ struct{} `type:"structure"`

	// A token used for optimistic locking. AWS WAF returns a token to your get
	// and list requests, to mark the state of the entity at the time of the request.
	// To make changes to the entity associated with the token, you provide the
	// token to operations like update and delete. AWS WAF uses the token to ensure
	// that no changes have been made to the entity since you last retrieved it.
	// If a change has been made, the update fails with a WAFOptimisticLockException.
	// If this happens, perform another get, and use the new token returned by that
	// operation.
	NextWebACLLockToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteFirewallManagerRuleGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteFirewallManagerRuleGroups = "DeleteFirewallManagerRuleGroups"

// DeleteFirewallManagerRuleGroupsRequest returns a request value for making API operation for
// AWS WAFV2.
//
// Deletes all rule groups that are managed by AWS Firewall Manager for the
// specified web ACL.
//
// You can only use this if ManagedByFirewallManager is false in the specified
// WebACL.
//
//    // Example sending a request using DeleteFirewallManagerRuleGroupsRequest.
//    req := client.DeleteFirewallManagerRuleGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/wafv2-2019-07-29/DeleteFirewallManagerRuleGroups
func (c *Client) DeleteFirewallManagerRuleGroupsRequest(input *DeleteFirewallManagerRuleGroupsInput) DeleteFirewallManagerRuleGroupsRequest {
	op := &aws.Operation{
		Name:       opDeleteFirewallManagerRuleGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteFirewallManagerRuleGroupsInput{}
	}

	req := c.newRequest(op, input, &DeleteFirewallManagerRuleGroupsOutput{})

	return DeleteFirewallManagerRuleGroupsRequest{Request: req, Input: input, Copy: c.DeleteFirewallManagerRuleGroupsRequest}
}

// DeleteFirewallManagerRuleGroupsRequest is the request type for the
// DeleteFirewallManagerRuleGroups API operation.
type DeleteFirewallManagerRuleGroupsRequest struct {
	*aws.Request
	Input *DeleteFirewallManagerRuleGroupsInput
	Copy  func(*DeleteFirewallManagerRuleGroupsInput) DeleteFirewallManagerRuleGroupsRequest
}

// Send marshals and sends the DeleteFirewallManagerRuleGroups API request.
func (r DeleteFirewallManagerRuleGroupsRequest) Send(ctx context.Context) (*DeleteFirewallManagerRuleGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteFirewallManagerRuleGroupsResponse{
		DeleteFirewallManagerRuleGroupsOutput: r.Request.Data.(*DeleteFirewallManagerRuleGroupsOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteFirewallManagerRuleGroupsResponse is the response type for the
// DeleteFirewallManagerRuleGroups API operation.
type DeleteFirewallManagerRuleGroupsResponse struct {
	*DeleteFirewallManagerRuleGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteFirewallManagerRuleGroups request.
func (r *DeleteFirewallManagerRuleGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
