// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type EnablePolicyTypeInput struct {
	_ struct{} `type:"structure"`

	// The policy type that you want to enable.
	//
	// PolicyType is a required field
	PolicyType PolicyType `type:"string" required:"true" enum:"true"`

	// The unique identifier (ID) of the root in which you want to enable a policy
	// type. You can get the ID from the ListRoots operation.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a root ID string
	// requires "r-" followed by from 4 to 32 lowercase letters or digits.
	//
	// RootId is a required field
	RootId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s EnablePolicyTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnablePolicyTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EnablePolicyTypeInput"}
	if len(s.PolicyType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("PolicyType"))
	}

	if s.RootId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RootId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type EnablePolicyTypeOutput struct {
	_ struct{} `type:"structure"`

	// A structure that shows the root with the updated list of enabled policy types.
	Root *Root `type:"structure"`
}

// String returns the string representation
func (s EnablePolicyTypeOutput) String() string {
	return awsutil.Prettify(s)
}

const opEnablePolicyType = "EnablePolicyType"

// EnablePolicyTypeRequest returns a request value for making API operation for
// AWS Organizations.
//
// Enables a policy type in a root. After you enable a policy type in a root,
// you can attach policies of that type to the root, any organizational unit
// (OU), or account in that root. You can undo this by using the DisablePolicyType
// operation.
//
// This is an asynchronous request that AWS performs in the background. AWS
// recommends that you first use ListRoots to see the status of policy types
// for a specified root, and then use this operation.
//
// This operation can be called only from the organization's master account.
//
// You can enable a policy type in a root only if that policy type is available
// in the organization. To view the status of available policy types in the
// organization, use DescribeOrganization.
//
//    // Example sending a request using EnablePolicyTypeRequest.
//    req := client.EnablePolicyTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/EnablePolicyType
func (c *Client) EnablePolicyTypeRequest(input *EnablePolicyTypeInput) EnablePolicyTypeRequest {
	op := &aws.Operation{
		Name:       opEnablePolicyType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnablePolicyTypeInput{}
	}

	req := c.newRequest(op, input, &EnablePolicyTypeOutput{})

	return EnablePolicyTypeRequest{Request: req, Input: input, Copy: c.EnablePolicyTypeRequest}
}

// EnablePolicyTypeRequest is the request type for the
// EnablePolicyType API operation.
type EnablePolicyTypeRequest struct {
	*aws.Request
	Input *EnablePolicyTypeInput
	Copy  func(*EnablePolicyTypeInput) EnablePolicyTypeRequest
}

// Send marshals and sends the EnablePolicyType API request.
func (r EnablePolicyTypeRequest) Send(ctx context.Context) (*EnablePolicyTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnablePolicyTypeResponse{
		EnablePolicyTypeOutput: r.Request.Data.(*EnablePolicyTypeOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnablePolicyTypeResponse is the response type for the
// EnablePolicyType API operation.
type EnablePolicyTypeResponse struct {
	*EnablePolicyTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnablePolicyType request.
func (r *EnablePolicyTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
