// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type StartSchemaCreationInput struct {
	_ struct{} `type:"structure"`

	// The API ID.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// The schema definition, in GraphQL schema language format.
	//
	// Definition is automatically base64 encoded/decoded by the SDK.
	//
	// Definition is a required field
	Definition []byte `locationName:"definition" type:"blob" required:"true"`
}

// String returns the string representation
func (s StartSchemaCreationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartSchemaCreationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartSchemaCreationInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.Definition == nil {
		invalidParams.Add(aws.NewErrParamRequired("Definition"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartSchemaCreationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Definition != nil {
		v := s.Definition

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "definition", protocol.QuotedValue{ValueMarshaler: protocol.BytesValue(v)}, metadata)
	}
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type StartSchemaCreationOutput struct {
	_ struct{} `type:"structure"`

	// The current state of the schema (PROCESSING, FAILED, SUCCESS, or NOT_APPLICABLE).
	// When the schema is in the ACTIVE state, you can add data.
	Status SchemaStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s StartSchemaCreationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartSchemaCreationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opStartSchemaCreation = "StartSchemaCreation"

// StartSchemaCreationRequest returns a request value for making API operation for
// AWS AppSync.
//
// Adds a new schema to your GraphQL API.
//
// This operation is asynchronous. Use to determine when it has completed.
//
//    // Example sending a request using StartSchemaCreationRequest.
//    req := client.StartSchemaCreationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/StartSchemaCreation
func (c *Client) StartSchemaCreationRequest(input *StartSchemaCreationInput) StartSchemaCreationRequest {
	op := &aws.Operation{
		Name:       opStartSchemaCreation,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apis/{apiId}/schemacreation",
	}

	if input == nil {
		input = &StartSchemaCreationInput{}
	}

	req := c.newRequest(op, input, &StartSchemaCreationOutput{})

	return StartSchemaCreationRequest{Request: req, Input: input, Copy: c.StartSchemaCreationRequest}
}

// StartSchemaCreationRequest is the request type for the
// StartSchemaCreation API operation.
type StartSchemaCreationRequest struct {
	*aws.Request
	Input *StartSchemaCreationInput
	Copy  func(*StartSchemaCreationInput) StartSchemaCreationRequest
}

// Send marshals and sends the StartSchemaCreation API request.
func (r StartSchemaCreationRequest) Send(ctx context.Context) (*StartSchemaCreationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartSchemaCreationResponse{
		StartSchemaCreationOutput: r.Request.Data.(*StartSchemaCreationOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartSchemaCreationResponse is the response type for the
// StartSchemaCreation API operation.
type StartSchemaCreationResponse struct {
	*StartSchemaCreationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartSchemaCreation request.
func (r *StartSchemaCreationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
