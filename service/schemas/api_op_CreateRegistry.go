// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type CreateRegistryInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	// RegistryName is a required field
	RegistryName *string `location:"uri" locationName:"registryName" type:"string" required:"true"`

	// Key-value pairs associated with a resource.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateRegistryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRegistryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRegistryInput"}

	if s.RegistryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RegistryName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateRegistryInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.RegistryName != nil {
		v := *s.RegistryName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "registryName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateRegistryOutput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	RegistryArn *string `type:"string"`

	RegistryName *string `type:"string"`

	// Key-value pairs associated with a resource.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateRegistryOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateRegistryOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RegistryArn != nil {
		v := *s.RegistryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RegistryArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RegistryName != nil {
		v := *s.RegistryName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RegistryName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opCreateRegistry = "CreateRegistry"

// CreateRegistryRequest returns a request value for making API operation for
// Schemas.
//
// Creates a registry.
//
//    // Example sending a request using CreateRegistryRequest.
//    req := client.CreateRegistryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/schemas-2019-12-02/CreateRegistry
func (c *Client) CreateRegistryRequest(input *CreateRegistryInput) CreateRegistryRequest {
	op := &aws.Operation{
		Name:       opCreateRegistry,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/registries/name/{registryName}",
	}

	if input == nil {
		input = &CreateRegistryInput{}
	}

	req := c.newRequest(op, input, &CreateRegistryOutput{})

	return CreateRegistryRequest{Request: req, Input: input, Copy: c.CreateRegistryRequest}
}

// CreateRegistryRequest is the request type for the
// CreateRegistry API operation.
type CreateRegistryRequest struct {
	*aws.Request
	Input *CreateRegistryInput
	Copy  func(*CreateRegistryInput) CreateRegistryRequest
}

// Send marshals and sends the CreateRegistry API request.
func (r CreateRegistryRequest) Send(ctx context.Context) (*CreateRegistryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRegistryResponse{
		CreateRegistryOutput: r.Request.Data.(*CreateRegistryOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRegistryResponse is the response type for the
// CreateRegistry API operation.
type CreateRegistryResponse struct {
	*CreateRegistryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRegistry request.
func (r *CreateRegistryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
