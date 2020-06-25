// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type PutPermissionPolicyInput struct {
	_ struct{} `type:"structure"`

	// The policy to attach to the specified RuleGroup.
	//
	// Policy is a required field
	Policy *string `min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the RuleGroup to which you want to attach
	// the policy.
	//
	// ResourceArn is a required field
	ResourceArn *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutPermissionPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutPermissionPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutPermissionPolicyInput"}

	if s.Policy == nil {
		invalidParams.Add(aws.NewErrParamRequired("Policy"))
	}
	if s.Policy != nil && len(*s.Policy) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Policy", 1))
	}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}
	if s.ResourceArn != nil && len(*s.ResourceArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutPermissionPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutPermissionPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutPermissionPolicy = "PutPermissionPolicy"

// PutPermissionPolicyRequest returns a request value for making API operation for
// AWS WAF.
//
//
// This is AWS WAF Classic documentation. For more information, see AWS WAF
// Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the AWS
// WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// With the latest version, AWS WAF has a single set of endpoints for regional
// and global use.
//
// Attaches an IAM policy to the specified resource. The only supported use
// for this action is to share a RuleGroup across accounts.
//
// The PutPermissionPolicy is subject to the following restrictions:
//
//    * You can attach only one policy with each PutPermissionPolicy request.
//
//    * The policy must include an Effect, Action and Principal.
//
//    * Effect must specify Allow.
//
//    * The Action in the policy must be waf:UpdateWebACL, waf-regional:UpdateWebACL,
//    waf:GetRuleGroup and waf-regional:GetRuleGroup . Any extra or wildcard
//    actions in the policy will be rejected.
//
//    * The policy cannot include a Resource parameter.
//
//    * The ARN in the request must be a valid WAF RuleGroup ARN and the RuleGroup
//    must exist in the same region.
//
//    * The user making the request must be the owner of the RuleGroup.
//
//    * Your policy must be composed using IAM Policy version 2012-10-17.
//
// For more information, see IAM Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html).
//
// An example of a valid policy parameter is shown in the Examples section below.
//
//    // Example sending a request using PutPermissionPolicyRequest.
//    req := client.PutPermissionPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/PutPermissionPolicy
func (c *Client) PutPermissionPolicyRequest(input *PutPermissionPolicyInput) PutPermissionPolicyRequest {
	op := &aws.Operation{
		Name:       opPutPermissionPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutPermissionPolicyInput{}
	}

	req := c.newRequest(op, input, &PutPermissionPolicyOutput{})

	return PutPermissionPolicyRequest{Request: req, Input: input, Copy: c.PutPermissionPolicyRequest}
}

// PutPermissionPolicyRequest is the request type for the
// PutPermissionPolicy API operation.
type PutPermissionPolicyRequest struct {
	*aws.Request
	Input *PutPermissionPolicyInput
	Copy  func(*PutPermissionPolicyInput) PutPermissionPolicyRequest
}

// Send marshals and sends the PutPermissionPolicy API request.
func (r PutPermissionPolicyRequest) Send(ctx context.Context) (*PutPermissionPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutPermissionPolicyResponse{
		PutPermissionPolicyOutput: r.Request.Data.(*PutPermissionPolicyOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutPermissionPolicyResponse is the response type for the
// PutPermissionPolicy API operation.
type PutPermissionPolicyResponse struct {
	*PutPermissionPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutPermissionPolicy request.
func (r *PutPermissionPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
