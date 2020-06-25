// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type PutAccessControlRuleInput struct {
	_ struct{} `type:"structure"`

	// Access protocol actions to include in the rule. Valid values include ActiveSync,
	// AutoDiscover, EWS, IMAP, SMTP, WindowsOutlook, and WebMail.
	Actions []string `type:"list"`

	// The rule description.
	//
	// Description is a required field
	Description *string `type:"string" required:"true"`

	// The rule effect.
	//
	// Effect is a required field
	Effect AccessControlRuleEffect `type:"string" required:"true" enum:"true"`

	// IPv4 CIDR ranges to include in the rule.
	IpRanges []string `type:"list"`

	// The rule name.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// Access protocol actions to exclude from the rule. Valid values include ActiveSync,
	// AutoDiscover, EWS, IMAP, SMTP, WindowsOutlook, and WebMail.
	NotActions []string `type:"list"`

	// IPv4 CIDR ranges to exclude from the rule.
	NotIpRanges []string `type:"list"`

	// User IDs to exclude from the rule.
	NotUserIds []string `type:"list"`

	// The identifier of the organization.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`

	// User IDs to include in the rule.
	UserIds []string `type:"list"`
}

// String returns the string representation
func (s PutAccessControlRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutAccessControlRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutAccessControlRuleInput"}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}
	if len(s.Effect) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Effect"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutAccessControlRuleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutAccessControlRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutAccessControlRule = "PutAccessControlRule"

// PutAccessControlRuleRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Adds a new access control rule for the specified organization. The rule allows
// or denies access to the organization for the specified IPv4 addresses, access
// protocol actions, and user IDs. Adding a new rule with the same name as an
// existing rule replaces the older rule.
//
//    // Example sending a request using PutAccessControlRuleRequest.
//    req := client.PutAccessControlRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/PutAccessControlRule
func (c *Client) PutAccessControlRuleRequest(input *PutAccessControlRuleInput) PutAccessControlRuleRequest {
	op := &aws.Operation{
		Name:       opPutAccessControlRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutAccessControlRuleInput{}
	}

	req := c.newRequest(op, input, &PutAccessControlRuleOutput{})

	return PutAccessControlRuleRequest{Request: req, Input: input, Copy: c.PutAccessControlRuleRequest}
}

// PutAccessControlRuleRequest is the request type for the
// PutAccessControlRule API operation.
type PutAccessControlRuleRequest struct {
	*aws.Request
	Input *PutAccessControlRuleInput
	Copy  func(*PutAccessControlRuleInput) PutAccessControlRuleRequest
}

// Send marshals and sends the PutAccessControlRule API request.
func (r PutAccessControlRuleRequest) Send(ctx context.Context) (*PutAccessControlRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutAccessControlRuleResponse{
		PutAccessControlRuleOutput: r.Request.Data.(*PutAccessControlRuleOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutAccessControlRuleResponse is the response type for the
// PutAccessControlRule API operation.
type PutAccessControlRuleResponse struct {
	*PutAccessControlRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutAccessControlRule request.
func (r *PutAccessControlRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
