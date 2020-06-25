// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetAppliedSchemaVersionInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the applied schema.
	//
	// SchemaArn is a required field
	SchemaArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetAppliedSchemaVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAppliedSchemaVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAppliedSchemaVersionInput"}

	if s.SchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAppliedSchemaVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SchemaArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetAppliedSchemaVersionOutput struct {
	_ struct{} `type:"structure"`

	// Current applied schema ARN, including the minor version in use if one was
	// provided.
	AppliedSchemaArn *string `type:"string"`
}

// String returns the string representation
func (s GetAppliedSchemaVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAppliedSchemaVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AppliedSchemaArn != nil {
		v := *s.AppliedSchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AppliedSchemaArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetAppliedSchemaVersion = "GetAppliedSchemaVersion"

// GetAppliedSchemaVersionRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Returns current applied schema version ARN, including the minor version in
// use.
//
//    // Example sending a request using GetAppliedSchemaVersionRequest.
//    req := client.GetAppliedSchemaVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/GetAppliedSchemaVersion
func (c *Client) GetAppliedSchemaVersionRequest(input *GetAppliedSchemaVersionInput) GetAppliedSchemaVersionRequest {
	op := &aws.Operation{
		Name:       opGetAppliedSchemaVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/schema/getappliedschema",
	}

	if input == nil {
		input = &GetAppliedSchemaVersionInput{}
	}

	req := c.newRequest(op, input, &GetAppliedSchemaVersionOutput{})

	return GetAppliedSchemaVersionRequest{Request: req, Input: input, Copy: c.GetAppliedSchemaVersionRequest}
}

// GetAppliedSchemaVersionRequest is the request type for the
// GetAppliedSchemaVersion API operation.
type GetAppliedSchemaVersionRequest struct {
	*aws.Request
	Input *GetAppliedSchemaVersionInput
	Copy  func(*GetAppliedSchemaVersionInput) GetAppliedSchemaVersionRequest
}

// Send marshals and sends the GetAppliedSchemaVersion API request.
func (r GetAppliedSchemaVersionRequest) Send(ctx context.Context) (*GetAppliedSchemaVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAppliedSchemaVersionResponse{
		GetAppliedSchemaVersionOutput: r.Request.Data.(*GetAppliedSchemaVersionOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAppliedSchemaVersionResponse is the response type for the
// GetAppliedSchemaVersion API operation.
type GetAppliedSchemaVersionResponse struct {
	*GetAppliedSchemaVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAppliedSchemaVersion request.
func (r *GetAppliedSchemaVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
