// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type UpdateVoiceTemplateInput struct {
	_ struct{} `type:"structure" payload:"VoiceTemplateRequest"`

	CreateNewVersion *bool `location:"querystring" locationName:"create-new-version" type:"boolean"`

	// TemplateName is a required field
	TemplateName *string `location:"uri" locationName:"template-name" type:"string" required:"true"`

	Version *string `location:"querystring" locationName:"version" type:"string"`

	// Specifies the content and settings for a message template that can be used
	// in messages that are sent through the voice channel.
	//
	// VoiceTemplateRequest is a required field
	VoiceTemplateRequest *VoiceTemplateRequest `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateVoiceTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateVoiceTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateVoiceTemplateInput"}

	if s.TemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateName"))
	}

	if s.VoiceTemplateRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceTemplateRequest"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateVoiceTemplateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.TemplateName != nil {
		v := *s.TemplateName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "template-name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VoiceTemplateRequest != nil {
		v := s.VoiceTemplateRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "VoiceTemplateRequest", v, metadata)
	}
	if s.CreateNewVersion != nil {
		v := *s.CreateNewVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "create-new-version", protocol.BoolValue(v), metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateVoiceTemplateOutput struct {
	_ struct{} `type:"structure" payload:"MessageBody"`

	// Provides information about an API request or response.
	//
	// MessageBody is a required field
	MessageBody *MessageBody `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateVoiceTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateVoiceTemplateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.MessageBody != nil {
		v := s.MessageBody

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "MessageBody", v, metadata)
	}
	return nil
}

const opUpdateVoiceTemplate = "UpdateVoiceTemplate"

// UpdateVoiceTemplateRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Updates an existing message template for messages that are sent through the
// voice channel.
//
//    // Example sending a request using UpdateVoiceTemplateRequest.
//    req := client.UpdateVoiceTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateVoiceTemplate
func (c *Client) UpdateVoiceTemplateRequest(input *UpdateVoiceTemplateInput) UpdateVoiceTemplateRequest {
	op := &aws.Operation{
		Name:       opUpdateVoiceTemplate,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/templates/{template-name}/voice",
	}

	if input == nil {
		input = &UpdateVoiceTemplateInput{}
	}

	req := c.newRequest(op, input, &UpdateVoiceTemplateOutput{})

	return UpdateVoiceTemplateRequest{Request: req, Input: input, Copy: c.UpdateVoiceTemplateRequest}
}

// UpdateVoiceTemplateRequest is the request type for the
// UpdateVoiceTemplate API operation.
type UpdateVoiceTemplateRequest struct {
	*aws.Request
	Input *UpdateVoiceTemplateInput
	Copy  func(*UpdateVoiceTemplateInput) UpdateVoiceTemplateRequest
}

// Send marshals and sends the UpdateVoiceTemplate API request.
func (r UpdateVoiceTemplateRequest) Send(ctx context.Context) (*UpdateVoiceTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateVoiceTemplateResponse{
		UpdateVoiceTemplateOutput: r.Request.Data.(*UpdateVoiceTemplateOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateVoiceTemplateResponse is the response type for the
// UpdateVoiceTemplate API operation.
type UpdateVoiceTemplateResponse struct {
	*UpdateVoiceTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateVoiceTemplate request.
func (r *UpdateVoiceTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
