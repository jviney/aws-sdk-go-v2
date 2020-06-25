// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteAccessControlRuleInput struct {
	_ struct{} `type:"structure"`

	// The name of the access control rule.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The identifier for the organization.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAccessControlRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAccessControlRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAccessControlRuleInput"}

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

type DeleteAccessControlRuleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAccessControlRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteAccessControlRule = "DeleteAccessControlRule"

// DeleteAccessControlRuleRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Deletes an access control rule for the specified WorkMail organization.
//
//    // Example sending a request using DeleteAccessControlRuleRequest.
//    req := client.DeleteAccessControlRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/DeleteAccessControlRule
func (c *Client) DeleteAccessControlRuleRequest(input *DeleteAccessControlRuleInput) DeleteAccessControlRuleRequest {
	op := &aws.Operation{
		Name:       opDeleteAccessControlRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteAccessControlRuleInput{}
	}

	req := c.newRequest(op, input, &DeleteAccessControlRuleOutput{})

	return DeleteAccessControlRuleRequest{Request: req, Input: input, Copy: c.DeleteAccessControlRuleRequest}
}

// DeleteAccessControlRuleRequest is the request type for the
// DeleteAccessControlRule API operation.
type DeleteAccessControlRuleRequest struct {
	*aws.Request
	Input *DeleteAccessControlRuleInput
	Copy  func(*DeleteAccessControlRuleInput) DeleteAccessControlRuleRequest
}

// Send marshals and sends the DeleteAccessControlRule API request.
func (r DeleteAccessControlRuleRequest) Send(ctx context.Context) (*DeleteAccessControlRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAccessControlRuleResponse{
		DeleteAccessControlRuleOutput: r.Request.Data.(*DeleteAccessControlRuleOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAccessControlRuleResponse is the response type for the
// DeleteAccessControlRule API operation.
type DeleteAccessControlRuleResponse struct {
	*DeleteAccessControlRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAccessControlRule request.
func (r *DeleteAccessControlRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
