// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/query"
)

type UpdateUserInput struct {
	_ struct{} `type:"structure"`

	// New path for the IAM user. Include this parameter only if you're changing
	// the user's path.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of either a forward slash (/) by itself
	// or a string that must begin and end with forward slashes. In addition, it
	// can contain any ASCII character from the ! (\u0021) through the DEL character
	// (\u007F), including most punctuation characters, digits, and upper and lowercased
	// letters.
	NewPath *string `min:"1" type:"string"`

	// New name for the user. Include this parameter only if you're changing the
	// user's name.
	//
	// IAM user, group, role, and policy names must be unique within the account.
	// Names are not distinguished by case. For example, you cannot create resources
	// named both "MyResource" and "myresource".
	NewUserName *string `min:"1" type:"string"`

	// Name of the user to update. If you're changing the name of the user, this
	// is the original user name.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateUserInput"}
	if s.NewPath != nil && len(*s.NewPath) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NewPath", 1))
	}
	if s.NewUserName != nil && len(*s.NewUserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NewUserName", 1))
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

type UpdateUserOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateUserOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateUser = "UpdateUser"

// UpdateUserRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Updates the name and/or the path of the specified IAM user.
//
// You should understand the implications of changing an IAM user's path or
// name. For more information, see Renaming an IAM User (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_manage.html#id_users_renaming)
// and Renaming an IAM Group (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_groups_manage_rename.html)
// in the IAM User Guide.
//
// To change a user name, the requester must have appropriate permissions on
// both the source object and the target object. For example, to change Bob
// to Robert, the entity making the request must have permission on Bob and
// Robert, or must have permission on all (*). For more information about permissions,
// see Permissions and Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/PermissionsAndPolicies.html).
//
//    // Example sending a request using UpdateUserRequest.
//    req := client.UpdateUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateUser
func (c *Client) UpdateUserRequest(input *UpdateUserInput) UpdateUserRequest {
	op := &aws.Operation{
		Name:       opUpdateUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateUserInput{}
	}

	req := c.newRequest(op, input, &UpdateUserOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return UpdateUserRequest{Request: req, Input: input, Copy: c.UpdateUserRequest}
}

// UpdateUserRequest is the request type for the
// UpdateUser API operation.
type UpdateUserRequest struct {
	*aws.Request
	Input *UpdateUserInput
	Copy  func(*UpdateUserInput) UpdateUserRequest
}

// Send marshals and sends the UpdateUser API request.
func (r UpdateUserRequest) Send(ctx context.Context) (*UpdateUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateUserResponse{
		UpdateUserOutput: r.Request.Data.(*UpdateUserOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateUserResponse is the response type for the
// UpdateUser API operation.
type UpdateUserResponse struct {
	*UpdateUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateUser request.
func (r *UpdateUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
