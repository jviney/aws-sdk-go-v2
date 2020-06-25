// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// Gets a documentation snapshot of an API.
type GetDocumentationVersionInput struct {
	_ struct{} `type:"structure"`

	// [Required] The version identifier of the to-be-retrieved documentation snapshot.
	//
	// DocumentationVersion is a required field
	DocumentationVersion *string `location:"uri" locationName:"doc_version" type:"string" required:"true"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDocumentationVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDocumentationVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDocumentationVersionInput"}

	if s.DocumentationVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentationVersion"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDocumentationVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DocumentationVersion != nil {
		v := *s.DocumentationVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "doc_version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// A snapshot of the documentation of an API.
//
// Publishing API documentation involves creating a documentation version associated
// with an API stage and exporting the versioned documentation to an external
// (e.g., OpenAPI) file.
//
// Documenting an API (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api.html),
// DocumentationPart, DocumentationVersions
type GetDocumentationVersionOutput struct {
	_ struct{} `type:"structure"`

	// The date when the API documentation snapshot is created.
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp"`

	// The description of the API documentation snapshot.
	Description *string `locationName:"description" type:"string"`

	// The version identifier of the API documentation snapshot.
	Version *string `locationName:"version" type:"string"`
}

// String returns the string representation
func (s GetDocumentationVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDocumentationVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetDocumentationVersion = "GetDocumentationVersion"

// GetDocumentationVersionRequest returns a request value for making API operation for
// Amazon API Gateway.
//
//    // Example sending a request using GetDocumentationVersionRequest.
//    req := client.GetDocumentationVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetDocumentationVersionRequest(input *GetDocumentationVersionInput) GetDocumentationVersionRequest {
	op := &aws.Operation{
		Name:       opGetDocumentationVersion,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis/{restapi_id}/documentation/versions/{doc_version}",
	}

	if input == nil {
		input = &GetDocumentationVersionInput{}
	}

	req := c.newRequest(op, input, &GetDocumentationVersionOutput{})

	return GetDocumentationVersionRequest{Request: req, Input: input, Copy: c.GetDocumentationVersionRequest}
}

// GetDocumentationVersionRequest is the request type for the
// GetDocumentationVersion API operation.
type GetDocumentationVersionRequest struct {
	*aws.Request
	Input *GetDocumentationVersionInput
	Copy  func(*GetDocumentationVersionInput) GetDocumentationVersionRequest
}

// Send marshals and sends the GetDocumentationVersion API request.
func (r GetDocumentationVersionRequest) Send(ctx context.Context) (*GetDocumentationVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDocumentationVersionResponse{
		GetDocumentationVersionOutput: r.Request.Data.(*GetDocumentationVersionOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDocumentationVersionResponse is the response type for the
// GetDocumentationVersion API operation.
type GetDocumentationVersionResponse struct {
	*GetDocumentationVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDocumentationVersion request.
func (r *GetDocumentationVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
