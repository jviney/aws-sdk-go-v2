// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type PutCodeBindingInput struct {
	_ struct{} `type:"structure"`

	// Language is a required field
	Language *string `location:"uri" locationName:"language" type:"string" required:"true"`

	// RegistryName is a required field
	RegistryName *string `location:"uri" locationName:"registryName" type:"string" required:"true"`

	// SchemaName is a required field
	SchemaName *string `location:"uri" locationName:"schemaName" type:"string" required:"true"`

	SchemaVersion *string `location:"querystring" locationName:"schemaVersion" type:"string"`
}

// String returns the string representation
func (s PutCodeBindingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutCodeBindingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutCodeBindingInput"}

	if s.Language == nil {
		invalidParams.Add(aws.NewErrParamRequired("Language"))
	}

	if s.RegistryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RegistryName"))
	}

	if s.SchemaName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutCodeBindingInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Language != nil {
		v := *s.Language

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "language", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RegistryName != nil {
		v := *s.RegistryName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "registryName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaName != nil {
		v := *s.SchemaName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "schemaName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaVersion != nil {
		v := *s.SchemaVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "schemaVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type PutCodeBindingOutput struct {
	_ struct{} `type:"structure"`

	CreationDate *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	LastModified *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	SchemaVersion *string `type:"string"`

	Status CodeGenerationStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s PutCodeBindingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutCodeBindingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreationDate != nil {
		v := *s.CreationDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationDate",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.LastModified != nil {
		v := *s.LastModified

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastModified",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.SchemaVersion != nil {
		v := *s.SchemaVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SchemaVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opPutCodeBinding = "PutCodeBinding"

// PutCodeBindingRequest returns a request value for making API operation for
// Schemas.
//
// Put code binding URI
//
//    // Example sending a request using PutCodeBindingRequest.
//    req := client.PutCodeBindingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/schemas-2019-12-02/PutCodeBinding
func (c *Client) PutCodeBindingRequest(input *PutCodeBindingInput) PutCodeBindingRequest {
	op := &aws.Operation{
		Name:       opPutCodeBinding,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/registries/name/{registryName}/schemas/name/{schemaName}/language/{language}",
	}

	if input == nil {
		input = &PutCodeBindingInput{}
	}

	req := c.newRequest(op, input, &PutCodeBindingOutput{})

	return PutCodeBindingRequest{Request: req, Input: input, Copy: c.PutCodeBindingRequest}
}

// PutCodeBindingRequest is the request type for the
// PutCodeBinding API operation.
type PutCodeBindingRequest struct {
	*aws.Request
	Input *PutCodeBindingInput
	Copy  func(*PutCodeBindingInput) PutCodeBindingRequest
}

// Send marshals and sends the PutCodeBinding API request.
func (r PutCodeBindingRequest) Send(ctx context.Context) (*PutCodeBindingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutCodeBindingResponse{
		PutCodeBindingOutput: r.Request.Data.(*PutCodeBindingOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutCodeBindingResponse is the response type for the
// PutCodeBinding API operation.
type PutCodeBindingResponse struct {
	*PutCodeBindingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutCodeBinding request.
func (r *PutCodeBindingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
