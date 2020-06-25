// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetDocumentPathInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string" sensitive:"true"`

	// The ID of the document.
	//
	// DocumentId is a required field
	DocumentId *string `location:"uri" locationName:"DocumentId" min:"1" type:"string" required:"true"`

	// A comma-separated list of values. Specify NAME to include the names of the
	// parent folders.
	Fields *string `location:"querystring" locationName:"fields" min:"1" type:"string"`

	// The maximum number of levels in the hierarchy to return.
	Limit *int64 `location:"querystring" locationName:"limit" min:"1" type:"integer"`

	// This value is not supported.
	Marker *string `location:"querystring" locationName:"marker" min:"1" type:"string"`
}

// String returns the string representation
func (s GetDocumentPathInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDocumentPathInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDocumentPathInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}

	if s.DocumentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentId"))
	}
	if s.DocumentId != nil && len(*s.DocumentId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DocumentId", 1))
	}
	if s.Fields != nil && len(*s.Fields) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Fields", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDocumentPathInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AuthenticationToken != nil {
		v := *s.AuthenticationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Authentication", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DocumentId != nil {
		v := *s.DocumentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DocumentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Fields != nil {
		v := *s.Fields

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "fields", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetDocumentPathOutput struct {
	_ struct{} `type:"structure"`

	// The path information.
	Path *ResourcePath `type:"structure"`
}

// String returns the string representation
func (s GetDocumentPathOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDocumentPathOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Path != nil {
		v := s.Path

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Path", v, metadata)
	}
	return nil
}

const opGetDocumentPath = "GetDocumentPath"

// GetDocumentPathRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Retrieves the path information (the hierarchy from the root folder) for the
// requested document.
//
// By default, Amazon WorkDocs returns a maximum of 100 levels upwards from
// the requested document and only includes the IDs of the parent folders in
// the path. You can limit the maximum number of levels. You can also request
// the names of the parent folders.
//
//    // Example sending a request using GetDocumentPathRequest.
//    req := client.GetDocumentPathRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/GetDocumentPath
func (c *Client) GetDocumentPathRequest(input *GetDocumentPathInput) GetDocumentPathRequest {
	op := &aws.Operation{
		Name:       opGetDocumentPath,
		HTTPMethod: "GET",
		HTTPPath:   "/api/v1/documents/{DocumentId}/path",
	}

	if input == nil {
		input = &GetDocumentPathInput{}
	}

	req := c.newRequest(op, input, &GetDocumentPathOutput{})

	return GetDocumentPathRequest{Request: req, Input: input, Copy: c.GetDocumentPathRequest}
}

// GetDocumentPathRequest is the request type for the
// GetDocumentPath API operation.
type GetDocumentPathRequest struct {
	*aws.Request
	Input *GetDocumentPathInput
	Copy  func(*GetDocumentPathInput) GetDocumentPathRequest
}

// Send marshals and sends the GetDocumentPath API request.
func (r GetDocumentPathRequest) Send(ctx context.Context) (*GetDocumentPathResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDocumentPathResponse{
		GetDocumentPathOutput: r.Request.Data.(*GetDocumentPathOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDocumentPathResponse is the response type for the
// GetDocumentPath API operation.
type GetDocumentPathResponse struct {
	*GetDocumentPathOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDocumentPath request.
func (r *GetDocumentPathResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
