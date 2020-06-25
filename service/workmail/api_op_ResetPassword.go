// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ResetPasswordInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the organization that contains the user for which the password
	// is reset.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`

	// The new password for the user.
	//
	// Password is a required field
	Password *string `type:"string" required:"true" sensitive:"true"`

	// The identifier of the user for whom the password is reset.
	//
	// UserId is a required field
	UserId *string `min:"12" type:"string" required:"true"`
}

// String returns the string representation
func (s ResetPasswordInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResetPasswordInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResetPasswordInput"}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if s.Password == nil {
		invalidParams.Add(aws.NewErrParamRequired("Password"))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}
	if s.UserId != nil && len(*s.UserId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("UserId", 12))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ResetPasswordOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ResetPasswordOutput) String() string {
	return awsutil.Prettify(s)
}

const opResetPassword = "ResetPassword"

// ResetPasswordRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Allows the administrator to reset the password for a user.
//
//    // Example sending a request using ResetPasswordRequest.
//    req := client.ResetPasswordRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/ResetPassword
func (c *Client) ResetPasswordRequest(input *ResetPasswordInput) ResetPasswordRequest {
	op := &aws.Operation{
		Name:       opResetPassword,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResetPasswordInput{}
	}

	req := c.newRequest(op, input, &ResetPasswordOutput{})

	return ResetPasswordRequest{Request: req, Input: input, Copy: c.ResetPasswordRequest}
}

// ResetPasswordRequest is the request type for the
// ResetPassword API operation.
type ResetPasswordRequest struct {
	*aws.Request
	Input *ResetPasswordInput
	Copy  func(*ResetPasswordInput) ResetPasswordRequest
}

// Send marshals and sends the ResetPassword API request.
func (r ResetPasswordRequest) Send(ctx context.Context) (*ResetPasswordResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResetPasswordResponse{
		ResetPasswordOutput: r.Request.Data.(*ResetPasswordOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResetPasswordResponse is the response type for the
// ResetPassword API operation.
type ResetPasswordResponse struct {
	*ResetPasswordOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResetPassword request.
func (r *ResetPasswordResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
