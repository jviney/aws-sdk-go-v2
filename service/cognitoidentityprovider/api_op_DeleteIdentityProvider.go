// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteIdentityProviderInput struct {
	_ struct{} `type:"structure"`

	// The identity provider name.
	//
	// ProviderName is a required field
	ProviderName *string `min:"1" type:"string" required:"true"`

	// The user pool ID.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteIdentityProviderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteIdentityProviderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteIdentityProviderInput"}

	if s.ProviderName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProviderName"))
	}
	if s.ProviderName != nil && len(*s.ProviderName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProviderName", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteIdentityProviderOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteIdentityProviderOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteIdentityProvider = "DeleteIdentityProvider"

// DeleteIdentityProviderRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Deletes an identity provider for a user pool.
//
//    // Example sending a request using DeleteIdentityProviderRequest.
//    req := client.DeleteIdentityProviderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/DeleteIdentityProvider
func (c *Client) DeleteIdentityProviderRequest(input *DeleteIdentityProviderInput) DeleteIdentityProviderRequest {
	op := &aws.Operation{
		Name:       opDeleteIdentityProvider,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteIdentityProviderInput{}
	}

	req := c.newRequest(op, input, &DeleteIdentityProviderOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteIdentityProviderRequest{Request: req, Input: input, Copy: c.DeleteIdentityProviderRequest}
}

// DeleteIdentityProviderRequest is the request type for the
// DeleteIdentityProvider API operation.
type DeleteIdentityProviderRequest struct {
	*aws.Request
	Input *DeleteIdentityProviderInput
	Copy  func(*DeleteIdentityProviderInput) DeleteIdentityProviderRequest
}

// Send marshals and sends the DeleteIdentityProvider API request.
func (r DeleteIdentityProviderRequest) Send(ctx context.Context) (*DeleteIdentityProviderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteIdentityProviderResponse{
		DeleteIdentityProviderOutput: r.Request.Data.(*DeleteIdentityProviderOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteIdentityProviderResponse is the response type for the
// DeleteIdentityProvider API operation.
type DeleteIdentityProviderResponse struct {
	*DeleteIdentityProviderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteIdentityProvider request.
func (r *DeleteIdentityProviderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
