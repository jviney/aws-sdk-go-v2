// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transfer

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteUserInput struct {
	_ struct{} `type:"structure"`

	// A system-assigned unique identifier for a file transfer protocol-enabled
	// server instance that has the user assigned to it.
	//
	// ServerId is a required field
	ServerId *string `min:"19" type:"string" required:"true"`

	// A unique string that identifies a user that is being deleted from a file
	// transfer protocol-enabled server.
	//
	// UserName is a required field
	UserName *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteUserInput"}

	if s.ServerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerId"))
	}
	if s.ServerId != nil && len(*s.ServerId) < 19 {
		invalidParams.Add(aws.NewErrParamMinLen("ServerId", 19))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteUserOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteUserOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteUser = "DeleteUser"

// DeleteUserRequest returns a request value for making API operation for
// AWS Transfer Family.
//
// Deletes the user belonging to a file transfer protocol-enabled server you
// specify.
//
// No response returns from this operation.
//
// When you delete a user from a server, the user's information is lost.
//
//    // Example sending a request using DeleteUserRequest.
//    req := client.DeleteUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/DeleteUser
func (c *Client) DeleteUserRequest(input *DeleteUserInput) DeleteUserRequest {
	op := &aws.Operation{
		Name:       opDeleteUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteUserInput{}
	}

	req := c.newRequest(op, input, &DeleteUserOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteUserRequest{Request: req, Input: input, Copy: c.DeleteUserRequest}
}

// DeleteUserRequest is the request type for the
// DeleteUser API operation.
type DeleteUserRequest struct {
	*aws.Request
	Input *DeleteUserInput
	Copy  func(*DeleteUserInput) DeleteUserRequest
}

// Send marshals and sends the DeleteUser API request.
func (r DeleteUserRequest) Send(ctx context.Context) (*DeleteUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteUserResponse{
		DeleteUserOutput: r.Request.Data.(*DeleteUserOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteUserResponse is the response type for the
// DeleteUser API operation.
type DeleteUserResponse struct {
	*DeleteUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteUser request.
func (r *DeleteUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
