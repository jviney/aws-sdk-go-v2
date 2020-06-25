// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DisableUserInput struct {
	_ struct{} `type:"structure"`

	// The authentication type for the user. You must specify USERPOOL.
	//
	// AuthenticationType is a required field
	AuthenticationType AuthenticationType `type:"string" required:"true" enum:"true"`

	// The email address of the user.
	//
	// Users' email addresses are case-sensitive.
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s DisableUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisableUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisableUserInput"}
	if len(s.AuthenticationType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("AuthenticationType"))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisableUserOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisableUserOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisableUser = "DisableUser"

// DisableUserRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Disables the specified user in the user pool. Users can't sign in to AppStream
// 2.0 until they are re-enabled. This action does not delete the user.
//
//    // Example sending a request using DisableUserRequest.
//    req := client.DisableUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DisableUser
func (c *Client) DisableUserRequest(input *DisableUserInput) DisableUserRequest {
	op := &aws.Operation{
		Name:       opDisableUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisableUserInput{}
	}

	req := c.newRequest(op, input, &DisableUserOutput{})

	return DisableUserRequest{Request: req, Input: input, Copy: c.DisableUserRequest}
}

// DisableUserRequest is the request type for the
// DisableUser API operation.
type DisableUserRequest struct {
	*aws.Request
	Input *DisableUserInput
	Copy  func(*DisableUserInput) DisableUserRequest
}

// Send marshals and sends the DisableUser API request.
func (r DisableUserRequest) Send(ctx context.Context) (*DisableUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisableUserResponse{
		DisableUserOutput: r.Request.Data.(*DisableUserOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisableUserResponse is the response type for the
// DisableUser API operation.
type DisableUserResponse struct {
	*DisableUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisableUser request.
func (r *DisableUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
