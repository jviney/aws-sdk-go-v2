// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListTemplateVersionsInput struct {
	_ struct{} `type:"structure"`

	NextToken *string `location:"querystring" locationName:"next-token" type:"string"`

	PageSize *string `location:"querystring" locationName:"page-size" type:"string"`

	// TemplateName is a required field
	TemplateName *string `location:"uri" locationName:"template-name" type:"string" required:"true"`

	// TemplateType is a required field
	TemplateType *string `location:"uri" locationName:"template-type" type:"string" required:"true"`
}

// String returns the string representation
func (s ListTemplateVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTemplateVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTemplateVersionsInput"}

	if s.TemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateName"))
	}

	if s.TemplateType == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTemplateVersionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.TemplateName != nil {
		v := *s.TemplateName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "template-name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateType != nil {
		v := *s.TemplateType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "template-type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "next-token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PageSize != nil {
		v := *s.PageSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "page-size", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListTemplateVersionsOutput struct {
	_ struct{} `type:"structure" payload:"TemplateVersionsResponse"`

	// Provides information about all the versions of a specific message template.
	//
	// TemplateVersionsResponse is a required field
	TemplateVersionsResponse *TemplateVersionsResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s ListTemplateVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTemplateVersionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.TemplateVersionsResponse != nil {
		v := s.TemplateVersionsResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "TemplateVersionsResponse", v, metadata)
	}
	return nil
}

const opListTemplateVersions = "ListTemplateVersions"

// ListTemplateVersionsRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about all the versions of a specific message template.
//
//    // Example sending a request using ListTemplateVersionsRequest.
//    req := client.ListTemplateVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/ListTemplateVersions
func (c *Client) ListTemplateVersionsRequest(input *ListTemplateVersionsInput) ListTemplateVersionsRequest {
	op := &aws.Operation{
		Name:       opListTemplateVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/templates/{template-name}/{template-type}/versions",
	}

	if input == nil {
		input = &ListTemplateVersionsInput{}
	}

	req := c.newRequest(op, input, &ListTemplateVersionsOutput{})

	return ListTemplateVersionsRequest{Request: req, Input: input, Copy: c.ListTemplateVersionsRequest}
}

// ListTemplateVersionsRequest is the request type for the
// ListTemplateVersions API operation.
type ListTemplateVersionsRequest struct {
	*aws.Request
	Input *ListTemplateVersionsInput
	Copy  func(*ListTemplateVersionsInput) ListTemplateVersionsRequest
}

// Send marshals and sends the ListTemplateVersions API request.
func (r ListTemplateVersionsRequest) Send(ctx context.Context) (*ListTemplateVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTemplateVersionsResponse{
		ListTemplateVersionsOutput: r.Request.Data.(*ListTemplateVersionsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTemplateVersionsResponse is the response type for the
// ListTemplateVersions API operation.
type ListTemplateVersionsResponse struct {
	*ListTemplateVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTemplateVersions request.
func (r *ListTemplateVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
