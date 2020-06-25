// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeEffectivePolicyInput struct {
	_ struct{} `type:"structure"`

	// The type of policy that you want information about.
	//
	// PolicyType is a required field
	PolicyType EffectivePolicyType `type:"string" required:"true" enum:"true"`

	// When you're signed in as the master account, specify the ID of the account
	// that you want details about. Specifying an organization root or OU as the
	// target is not supported.
	TargetId *string `type:"string"`
}

// String returns the string representation
func (s DescribeEffectivePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEffectivePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEffectivePolicyInput"}
	if len(s.PolicyType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("PolicyType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeEffectivePolicyOutput struct {
	_ struct{} `type:"structure"`

	// The contents of the effective policy.
	EffectivePolicy *EffectivePolicy `type:"structure"`
}

// String returns the string representation
func (s DescribeEffectivePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEffectivePolicy = "DescribeEffectivePolicy"

// DescribeEffectivePolicyRequest returns a request value for making API operation for
// AWS Organizations.
//
// Returns the contents of the effective tag policy for the account. The effective
// tag policy is the aggregation of any tag policies the account inherits, plus
// any policy directly that is attached to the account.
//
// This action returns information on tag policies only.
//
// For more information on policy inheritance, see How Policy Inheritance Works
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies-inheritance.html)
// in the AWS Organizations User Guide.
//
// This operation can be called only from the organization's master account
// or by a member account that is a delegated administrator for an AWS service.
//
//    // Example sending a request using DescribeEffectivePolicyRequest.
//    req := client.DescribeEffectivePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/DescribeEffectivePolicy
func (c *Client) DescribeEffectivePolicyRequest(input *DescribeEffectivePolicyInput) DescribeEffectivePolicyRequest {
	op := &aws.Operation{
		Name:       opDescribeEffectivePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEffectivePolicyInput{}
	}

	req := c.newRequest(op, input, &DescribeEffectivePolicyOutput{})

	return DescribeEffectivePolicyRequest{Request: req, Input: input, Copy: c.DescribeEffectivePolicyRequest}
}

// DescribeEffectivePolicyRequest is the request type for the
// DescribeEffectivePolicy API operation.
type DescribeEffectivePolicyRequest struct {
	*aws.Request
	Input *DescribeEffectivePolicyInput
	Copy  func(*DescribeEffectivePolicyInput) DescribeEffectivePolicyRequest
}

// Send marshals and sends the DescribeEffectivePolicy API request.
func (r DescribeEffectivePolicyRequest) Send(ctx context.Context) (*DescribeEffectivePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEffectivePolicyResponse{
		DescribeEffectivePolicyOutput: r.Request.Data.(*DescribeEffectivePolicyOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEffectivePolicyResponse is the response type for the
// DescribeEffectivePolicy API operation.
type DescribeEffectivePolicyResponse struct {
	*DescribeEffectivePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEffectivePolicy request.
func (r *DescribeEffectivePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
