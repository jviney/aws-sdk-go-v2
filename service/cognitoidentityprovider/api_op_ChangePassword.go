// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to change a user password.
type ChangePasswordInput struct {
	_ struct{} `type:"structure"`

	// The access token.
	//
	// AccessToken is a required field
	AccessToken *string `type:"string" required:"true" sensitive:"true"`

	// The old password.
	//
	// PreviousPassword is a required field
	PreviousPassword *string `min:"6" type:"string" required:"true" sensitive:"true"`

	// The new password.
	//
	// ProposedPassword is a required field
	ProposedPassword *string `min:"6" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s ChangePasswordInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ChangePasswordInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ChangePasswordInput"}

	if s.AccessToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessToken"))
	}

	if s.PreviousPassword == nil {
		invalidParams.Add(aws.NewErrParamRequired("PreviousPassword"))
	}
	if s.PreviousPassword != nil && len(*s.PreviousPassword) < 6 {
		invalidParams.Add(aws.NewErrParamMinLen("PreviousPassword", 6))
	}

	if s.ProposedPassword == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProposedPassword"))
	}
	if s.ProposedPassword != nil && len(*s.ProposedPassword) < 6 {
		invalidParams.Add(aws.NewErrParamMinLen("ProposedPassword", 6))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The response from the server to the change password request.
type ChangePasswordOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ChangePasswordOutput) String() string {
	return awsutil.Prettify(s)
}

const opChangePassword = "ChangePassword"

// ChangePasswordRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Changes the password for a specified user in a user pool.
//
//    // Example sending a request using ChangePasswordRequest.
//    req := client.ChangePasswordRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ChangePassword
func (c *Client) ChangePasswordRequest(input *ChangePasswordInput) ChangePasswordRequest {
	op := &aws.Operation{
		Name:       opChangePassword,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ChangePasswordInput{}
	}

	req := c.newRequest(op, input, &ChangePasswordOutput{})
	req.Config.Credentials = aws.AnonymousCredentials

	return ChangePasswordRequest{Request: req, Input: input, Copy: c.ChangePasswordRequest}
}

// ChangePasswordRequest is the request type for the
// ChangePassword API operation.
type ChangePasswordRequest struct {
	*aws.Request
	Input *ChangePasswordInput
	Copy  func(*ChangePasswordInput) ChangePasswordRequest
}

// Send marshals and sends the ChangePassword API request.
func (r ChangePasswordRequest) Send(ctx context.Context) (*ChangePasswordResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ChangePasswordResponse{
		ChangePasswordOutput: r.Request.Data.(*ChangePasswordOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ChangePasswordResponse is the response type for the
// ChangePassword API operation.
type ChangePasswordResponse struct {
	*ChangePasswordOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ChangePassword request.
func (r *ChangePasswordResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
