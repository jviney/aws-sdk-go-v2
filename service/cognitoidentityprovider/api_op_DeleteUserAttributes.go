// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to delete user attributes.
type DeleteUserAttributesInput struct {
	_ struct{} `type:"structure"`

	// The access token used in the request to delete user attributes.
	//
	// AccessToken is a required field
	AccessToken *string `type:"string" required:"true" sensitive:"true"`

	// An array of strings representing the user attribute names you wish to delete.
	//
	// For custom attributes, you must prepend the custom: prefix to the attribute
	// name.
	//
	// UserAttributeNames is a required field
	UserAttributeNames []string `type:"list" required:"true"`
}

// String returns the string representation
func (s DeleteUserAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteUserAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteUserAttributesInput"}

	if s.AccessToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessToken"))
	}

	if s.UserAttributeNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserAttributeNames"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server to delete user attributes.
type DeleteUserAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteUserAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteUserAttributes = "DeleteUserAttributes"

// DeleteUserAttributesRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Deletes the attributes for a user.
//
//    // Example sending a request using DeleteUserAttributesRequest.
//    req := client.DeleteUserAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/DeleteUserAttributes
func (c *Client) DeleteUserAttributesRequest(input *DeleteUserAttributesInput) DeleteUserAttributesRequest {
	op := &aws.Operation{
		Name:       opDeleteUserAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteUserAttributesInput{}
	}

	req := c.newRequest(op, input, &DeleteUserAttributesOutput{})
	req.Config.Credentials = aws.AnonymousCredentials

	return DeleteUserAttributesRequest{Request: req, Input: input, Copy: c.DeleteUserAttributesRequest}
}

// DeleteUserAttributesRequest is the request type for the
// DeleteUserAttributes API operation.
type DeleteUserAttributesRequest struct {
	*aws.Request
	Input *DeleteUserAttributesInput
	Copy  func(*DeleteUserAttributesInput) DeleteUserAttributesRequest
}

// Send marshals and sends the DeleteUserAttributes API request.
func (r DeleteUserAttributesRequest) Send(ctx context.Context) (*DeleteUserAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteUserAttributesResponse{
		DeleteUserAttributesOutput: r.Request.Data.(*DeleteUserAttributesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteUserAttributesResponse is the response type for the
// DeleteUserAttributes API operation.
type DeleteUserAttributesResponse struct {
	*DeleteUserAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteUserAttributes request.
func (r *DeleteUserAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
