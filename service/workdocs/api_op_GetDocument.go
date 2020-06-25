// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetDocumentInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string" sensitive:"true"`

	// The ID of the document.
	//
	// DocumentId is a required field
	DocumentId *string `location:"uri" locationName:"DocumentId" min:"1" type:"string" required:"true"`

	// Set this to TRUE to include custom metadata in the response.
	IncludeCustomMetadata *bool `location:"querystring" locationName:"includeCustomMetadata" type:"boolean"`
}

// String returns the string representation
func (s GetDocumentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDocumentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDocumentInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}

	if s.DocumentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentId"))
	}
	if s.DocumentId != nil && len(*s.DocumentId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DocumentId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDocumentInput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.IncludeCustomMetadata != nil {
		v := *s.IncludeCustomMetadata

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "includeCustomMetadata", protocol.BoolValue(v), metadata)
	}
	return nil
}

type GetDocumentOutput struct {
	_ struct{} `type:"structure"`

	// The custom metadata on the document.
	CustomMetadata map[string]string `min:"1" type:"map"`

	// The metadata details of the document.
	Metadata *DocumentMetadata `type:"structure"`
}

// String returns the string representation
func (s GetDocumentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDocumentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CustomMetadata != nil {
		v := s.CustomMetadata

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "CustomMetadata", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Metadata != nil {
		v := s.Metadata

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Metadata", v, metadata)
	}
	return nil
}

const opGetDocument = "GetDocument"

// GetDocumentRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Retrieves details of a document.
//
//    // Example sending a request using GetDocumentRequest.
//    req := client.GetDocumentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/GetDocument
func (c *Client) GetDocumentRequest(input *GetDocumentInput) GetDocumentRequest {
	op := &aws.Operation{
		Name:       opGetDocument,
		HTTPMethod: "GET",
		HTTPPath:   "/api/v1/documents/{DocumentId}",
	}

	if input == nil {
		input = &GetDocumentInput{}
	}

	req := c.newRequest(op, input, &GetDocumentOutput{})

	return GetDocumentRequest{Request: req, Input: input, Copy: c.GetDocumentRequest}
}

// GetDocumentRequest is the request type for the
// GetDocument API operation.
type GetDocumentRequest struct {
	*aws.Request
	Input *GetDocumentInput
	Copy  func(*GetDocumentInput) GetDocumentRequest
}

// Send marshals and sends the GetDocument API request.
func (r GetDocumentRequest) Send(ctx context.Context) (*GetDocumentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDocumentResponse{
		GetDocumentOutput: r.Request.Data.(*GetDocumentOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDocumentResponse is the response type for the
// GetDocument API operation.
type GetDocumentResponse struct {
	*GetDocumentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDocument request.
func (r *GetDocumentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
