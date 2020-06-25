// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateIdentityProviderInput struct {
	_ struct{} `type:"structure"`

	// A mapping of identity provider attributes to standard and custom user pool
	// attributes.
	AttributeMapping map[string]string `type:"map"`

	// A list of identity provider identifiers.
	IdpIdentifiers []string `type:"list"`

	// The identity provider details. The following list describes the provider
	// detail keys for each identity provider type.
	//
	//    * For Google, Facebook and Login with Amazon: client_id client_secret
	//    authorize_scopes
	//
	//    * For Sign in with Apple: client_id team_id key_id private_key authorize_scopes
	//
	//    * For OIDC providers: client_id client_secret attributes_request_method
	//    oidc_issuer authorize_scopes authorize_url if not available from discovery
	//    URL specified by oidc_issuer key token_url if not available from discovery
	//    URL specified by oidc_issuer key attributes_url if not available from
	//    discovery URL specified by oidc_issuer key jwks_uri if not available from
	//    discovery URL specified by oidc_issuer key authorize_scopes
	//
	//    * For SAML providers: MetadataFile OR MetadataURL IDPSignout optional
	//
	// ProviderDetails is a required field
	ProviderDetails map[string]string `type:"map" required:"true"`

	// The identity provider name.
	//
	// ProviderName is a required field
	ProviderName *string `min:"1" type:"string" required:"true"`

	// The identity provider type.
	//
	// ProviderType is a required field
	ProviderType IdentityProviderTypeType `type:"string" required:"true" enum:"true"`

	// The user pool ID.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateIdentityProviderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateIdentityProviderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateIdentityProviderInput"}

	if s.ProviderDetails == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProviderDetails"))
	}

	if s.ProviderName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProviderName"))
	}
	if s.ProviderName != nil && len(*s.ProviderName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProviderName", 1))
	}
	if len(s.ProviderType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ProviderType"))
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

type CreateIdentityProviderOutput struct {
	_ struct{} `type:"structure"`

	// The newly created identity provider object.
	//
	// IdentityProvider is a required field
	IdentityProvider *IdentityProviderType `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateIdentityProviderOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateIdentityProvider = "CreateIdentityProvider"

// CreateIdentityProviderRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Creates an identity provider for a user pool.
//
//    // Example sending a request using CreateIdentityProviderRequest.
//    req := client.CreateIdentityProviderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/CreateIdentityProvider
func (c *Client) CreateIdentityProviderRequest(input *CreateIdentityProviderInput) CreateIdentityProviderRequest {
	op := &aws.Operation{
		Name:       opCreateIdentityProvider,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateIdentityProviderInput{}
	}

	req := c.newRequest(op, input, &CreateIdentityProviderOutput{})

	return CreateIdentityProviderRequest{Request: req, Input: input, Copy: c.CreateIdentityProviderRequest}
}

// CreateIdentityProviderRequest is the request type for the
// CreateIdentityProvider API operation.
type CreateIdentityProviderRequest struct {
	*aws.Request
	Input *CreateIdentityProviderInput
	Copy  func(*CreateIdentityProviderInput) CreateIdentityProviderRequest
}

// Send marshals and sends the CreateIdentityProvider API request.
func (r CreateIdentityProviderRequest) Send(ctx context.Context) (*CreateIdentityProviderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateIdentityProviderResponse{
		CreateIdentityProviderOutput: r.Request.Data.(*CreateIdentityProviderOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateIdentityProviderResponse is the response type for the
// CreateIdentityProvider API operation.
type CreateIdentityProviderResponse struct {
	*CreateIdentityProviderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateIdentityProvider request.
func (r *CreateIdentityProviderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
