// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreatePolicyInput struct {
	_ struct{} `type:"structure"`

	// The policy content to add to the new policy. For example, if you create a
	// service control policy (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scp.html)
	// (SCP), this string must be JSON text that specifies the permissions that
	// admins in attached accounts can delegate to their users, groups, and roles.
	// For more information about the SCP syntax, see Service Control Policy Syntax
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_reference_scp-syntax.html)
	// in the AWS Organizations User Guide.
	//
	// Content is a required field
	Content *string `min:"1" type:"string" required:"true"`

	// An optional description to assign to the policy.
	//
	// Description is a required field
	Description *string `type:"string" required:"true"`

	// The friendly name to assign to the policy.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) that is used to validate
	// this parameter is a string of any of the characters in the ASCII character
	// range.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The type of policy to create.
	//
	// In the current release, the only type of policy that you can create is a
	// service control policy (SCP).
	//
	// Type is a required field
	Type PolicyType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreatePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePolicyInput"}

	if s.Content == nil {
		invalidParams.Add(aws.NewErrParamRequired("Content"))
	}
	if s.Content != nil && len(*s.Content) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Content", 1))
	}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreatePolicyOutput struct {
	_ struct{} `type:"structure"`

	// A structure that contains details about the newly created policy.
	Policy *Policy `type:"structure"`
}

// String returns the string representation
func (s CreatePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreatePolicy = "CreatePolicy"

// CreatePolicyRequest returns a request value for making API operation for
// AWS Organizations.
//
// Creates a policy of a specified type that you can attach to a root, an organizational
// unit (OU), or an individual AWS account.
//
// For more information about policies and their use, see Managing Organization
// Policies (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies.html).
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using CreatePolicyRequest.
//    req := client.CreatePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/CreatePolicy
func (c *Client) CreatePolicyRequest(input *CreatePolicyInput) CreatePolicyRequest {
	op := &aws.Operation{
		Name:       opCreatePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreatePolicyInput{}
	}

	req := c.newRequest(op, input, &CreatePolicyOutput{})

	return CreatePolicyRequest{Request: req, Input: input, Copy: c.CreatePolicyRequest}
}

// CreatePolicyRequest is the request type for the
// CreatePolicy API operation.
type CreatePolicyRequest struct {
	*aws.Request
	Input *CreatePolicyInput
	Copy  func(*CreatePolicyInput) CreatePolicyRequest
}

// Send marshals and sends the CreatePolicy API request.
func (r CreatePolicyRequest) Send(ctx context.Context) (*CreatePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePolicyResponse{
		CreatePolicyOutput: r.Request.Data.(*CreatePolicyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePolicyResponse is the response type for the
// CreatePolicy API operation.
type CreatePolicyResponse struct {
	*CreatePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePolicy request.
func (r *CreatePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
