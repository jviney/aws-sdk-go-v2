// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type UpdateTemplateActiveVersionInput struct {
	_ struct{} `type:"structure" payload:"TemplateActiveVersionRequest"`

	// Specifies which version of a message template to use as the active version
	// of the template.
	//
	// TemplateActiveVersionRequest is a required field
	TemplateActiveVersionRequest *TemplateActiveVersionRequest `type:"structure" required:"true"`

	// TemplateName is a required field
	TemplateName *string `location:"uri" locationName:"template-name" type:"string" required:"true"`

	// TemplateType is a required field
	TemplateType *string `location:"uri" locationName:"template-type" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateTemplateActiveVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTemplateActiveVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTemplateActiveVersionInput"}

	if s.TemplateActiveVersionRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateActiveVersionRequest"))
	}

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
func (s UpdateTemplateActiveVersionInput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.TemplateActiveVersionRequest != nil {
		v := s.TemplateActiveVersionRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "TemplateActiveVersionRequest", v, metadata)
	}
	return nil
}

type UpdateTemplateActiveVersionOutput struct {
	_ struct{} `type:"structure" payload:"MessageBody"`

	// Provides information about an API request or response.
	//
	// MessageBody is a required field
	MessageBody *MessageBody `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateTemplateActiveVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateTemplateActiveVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.MessageBody != nil {
		v := s.MessageBody

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "MessageBody", v, metadata)
	}
	return nil
}

const opUpdateTemplateActiveVersion = "UpdateTemplateActiveVersion"

// UpdateTemplateActiveVersionRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Changes the status of a specific version of a message template to active.
//
//    // Example sending a request using UpdateTemplateActiveVersionRequest.
//    req := client.UpdateTemplateActiveVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateTemplateActiveVersion
func (c *Client) UpdateTemplateActiveVersionRequest(input *UpdateTemplateActiveVersionInput) UpdateTemplateActiveVersionRequest {
	op := &aws.Operation{
		Name:       opUpdateTemplateActiveVersion,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/templates/{template-name}/{template-type}/active-version",
	}

	if input == nil {
		input = &UpdateTemplateActiveVersionInput{}
	}

	req := c.newRequest(op, input, &UpdateTemplateActiveVersionOutput{})

	return UpdateTemplateActiveVersionRequest{Request: req, Input: input, Copy: c.UpdateTemplateActiveVersionRequest}
}

// UpdateTemplateActiveVersionRequest is the request type for the
// UpdateTemplateActiveVersion API operation.
type UpdateTemplateActiveVersionRequest struct {
	*aws.Request
	Input *UpdateTemplateActiveVersionInput
	Copy  func(*UpdateTemplateActiveVersionInput) UpdateTemplateActiveVersionRequest
}

// Send marshals and sends the UpdateTemplateActiveVersion API request.
func (r UpdateTemplateActiveVersionRequest) Send(ctx context.Context) (*UpdateTemplateActiveVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTemplateActiveVersionResponse{
		UpdateTemplateActiveVersionOutput: r.Request.Data.(*UpdateTemplateActiveVersionOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTemplateActiveVersionResponse is the response type for the
// UpdateTemplateActiveVersion API operation.
type UpdateTemplateActiveVersionResponse struct {
	*UpdateTemplateActiveVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTemplateActiveVersion request.
func (r *UpdateTemplateActiveVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
