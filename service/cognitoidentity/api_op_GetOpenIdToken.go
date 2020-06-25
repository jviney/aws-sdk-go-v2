// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentity

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Input to the GetOpenIdToken action.
type GetOpenIdTokenInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier in the format REGION:GUID.
	//
	// IdentityId is a required field
	IdentityId *string `min:"1" type:"string" required:"true"`

	// A set of optional name-value pairs that map provider names to provider tokens.
	// When using graph.facebook.com and www.amazon.com, supply the access_token
	// returned from the provider's authflow. For accounts.google.com, an Amazon
	// Cognito user pool provider, or any other OpenId Connect provider, always
	// include the id_token.
	Logins map[string]string `type:"map"`
}

// String returns the string representation
func (s GetOpenIdTokenInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetOpenIdTokenInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetOpenIdTokenInput"}

	if s.IdentityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityId"))
	}
	if s.IdentityId != nil && len(*s.IdentityId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Returned in response to a successful GetOpenIdToken request.
type GetOpenIdTokenOutput struct {
	_ struct{} `type:"structure"`

	// A unique identifier in the format REGION:GUID. Note that the IdentityId returned
	// may not match the one passed on input.
	IdentityId *string `min:"1" type:"string"`

	// An OpenID token, valid for 10 minutes.
	Token *string `type:"string"`
}

// String returns the string representation
func (s GetOpenIdTokenOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetOpenIdToken = "GetOpenIdToken"

// GetOpenIdTokenRequest returns a request value for making API operation for
// Amazon Cognito Identity.
//
// Gets an OpenID token, using a known Cognito ID. This known Cognito ID is
// returned by GetId. You can optionally add additional logins for the identity.
// Supplying multiple logins creates an implicit link.
//
// The OpenId token is valid for 10 minutes.
//
// This is a public API. You do not need any credentials to call this API.
//
//    // Example sending a request using GetOpenIdTokenRequest.
//    req := client.GetOpenIdTokenRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-identity-2014-06-30/GetOpenIdToken
func (c *Client) GetOpenIdTokenRequest(input *GetOpenIdTokenInput) GetOpenIdTokenRequest {
	op := &aws.Operation{
		Name:       opGetOpenIdToken,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetOpenIdTokenInput{}
	}

	req := c.newRequest(op, input, &GetOpenIdTokenOutput{})
	req.Config.Credentials = aws.AnonymousCredentials

	return GetOpenIdTokenRequest{Request: req, Input: input, Copy: c.GetOpenIdTokenRequest}
}

// GetOpenIdTokenRequest is the request type for the
// GetOpenIdToken API operation.
type GetOpenIdTokenRequest struct {
	*aws.Request
	Input *GetOpenIdTokenInput
	Copy  func(*GetOpenIdTokenInput) GetOpenIdTokenRequest
}

// Send marshals and sends the GetOpenIdToken API request.
func (r GetOpenIdTokenRequest) Send(ctx context.Context) (*GetOpenIdTokenResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetOpenIdTokenResponse{
		GetOpenIdTokenOutput: r.Request.Data.(*GetOpenIdTokenOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetOpenIdTokenResponse is the response type for the
// GetOpenIdToken API operation.
type GetOpenIdTokenResponse struct {
	*GetOpenIdTokenOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetOpenIdToken request.
func (r *GetOpenIdTokenResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
