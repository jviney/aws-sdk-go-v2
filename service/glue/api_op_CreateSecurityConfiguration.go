// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateSecurityConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The encryption configuration for the new security configuration.
	//
	// EncryptionConfiguration is a required field
	EncryptionConfiguration *EncryptionConfiguration `type:"structure" required:"true"`

	// The name for the new security configuration.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateSecurityConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSecurityConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSecurityConfigurationInput"}

	if s.EncryptionConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("EncryptionConfiguration"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateSecurityConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The time at which the new security configuration was created.
	CreatedTimestamp *time.Time `type:"timestamp"`

	// The name assigned to the new security configuration.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateSecurityConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateSecurityConfiguration = "CreateSecurityConfiguration"

// CreateSecurityConfigurationRequest returns a request value for making API operation for
// AWS Glue.
//
// Creates a new security configuration. A security configuration is a set of
// security properties that can be used by AWS Glue. You can use a security
// configuration to encrypt data at rest. For information about using security
// configurations in AWS Glue, see Encrypting Data Written by Crawlers, Jobs,
// and Development Endpoints (https://docs.aws.amazon.com/glue/latest/dg/encryption-security-configuration.html).
//
//    // Example sending a request using CreateSecurityConfigurationRequest.
//    req := client.CreateSecurityConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/CreateSecurityConfiguration
func (c *Client) CreateSecurityConfigurationRequest(input *CreateSecurityConfigurationInput) CreateSecurityConfigurationRequest {
	op := &aws.Operation{
		Name:       opCreateSecurityConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSecurityConfigurationInput{}
	}

	req := c.newRequest(op, input, &CreateSecurityConfigurationOutput{})

	return CreateSecurityConfigurationRequest{Request: req, Input: input, Copy: c.CreateSecurityConfigurationRequest}
}

// CreateSecurityConfigurationRequest is the request type for the
// CreateSecurityConfiguration API operation.
type CreateSecurityConfigurationRequest struct {
	*aws.Request
	Input *CreateSecurityConfigurationInput
	Copy  func(*CreateSecurityConfigurationInput) CreateSecurityConfigurationRequest
}

// Send marshals and sends the CreateSecurityConfiguration API request.
func (r CreateSecurityConfigurationRequest) Send(ctx context.Context) (*CreateSecurityConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSecurityConfigurationResponse{
		CreateSecurityConfigurationOutput: r.Request.Data.(*CreateSecurityConfigurationOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSecurityConfigurationResponse is the response type for the
// CreateSecurityConfiguration API operation.
type CreateSecurityConfigurationResponse struct {
	*CreateSecurityConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSecurityConfiguration request.
func (r *CreateSecurityConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
