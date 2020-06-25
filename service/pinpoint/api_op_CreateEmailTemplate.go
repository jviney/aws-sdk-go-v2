// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type CreateEmailTemplateInput struct {
	_ struct{} `type:"structure" payload:"EmailTemplateRequest"`

	// Specifies the content and settings for a message template that can be used
	// in messages that are sent through the email channel.
	//
	// EmailTemplateRequest is a required field
	EmailTemplateRequest *EmailTemplateRequest `type:"structure" required:"true"`

	// TemplateName is a required field
	TemplateName *string `location:"uri" locationName:"template-name" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateEmailTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateEmailTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateEmailTemplateInput"}

	if s.EmailTemplateRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("EmailTemplateRequest"))
	}

	if s.TemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateEmailTemplateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.TemplateName != nil {
		v := *s.TemplateName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "template-name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EmailTemplateRequest != nil {
		v := s.EmailTemplateRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "EmailTemplateRequest", v, metadata)
	}
	return nil
}

type CreateEmailTemplateOutput struct {
	_ struct{} `type:"structure" payload:"CreateTemplateMessageBody"`

	// Provides information about a request to create a message template.
	//
	// CreateTemplateMessageBody is a required field
	CreateTemplateMessageBody *CreateTemplateMessageBody `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateEmailTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateEmailTemplateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreateTemplateMessageBody != nil {
		v := s.CreateTemplateMessageBody

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "CreateTemplateMessageBody", v, metadata)
	}
	return nil
}

const opCreateEmailTemplate = "CreateEmailTemplate"

// CreateEmailTemplateRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Creates a message template for messages that are sent through the email channel.
//
//    // Example sending a request using CreateEmailTemplateRequest.
//    req := client.CreateEmailTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/CreateEmailTemplate
func (c *Client) CreateEmailTemplateRequest(input *CreateEmailTemplateInput) CreateEmailTemplateRequest {
	op := &aws.Operation{
		Name:       opCreateEmailTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/templates/{template-name}/email",
	}

	if input == nil {
		input = &CreateEmailTemplateInput{}
	}

	req := c.newRequest(op, input, &CreateEmailTemplateOutput{})

	return CreateEmailTemplateRequest{Request: req, Input: input, Copy: c.CreateEmailTemplateRequest}
}

// CreateEmailTemplateRequest is the request type for the
// CreateEmailTemplate API operation.
type CreateEmailTemplateRequest struct {
	*aws.Request
	Input *CreateEmailTemplateInput
	Copy  func(*CreateEmailTemplateInput) CreateEmailTemplateRequest
}

// Send marshals and sends the CreateEmailTemplate API request.
func (r CreateEmailTemplateRequest) Send(ctx context.Context) (*CreateEmailTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateEmailTemplateResponse{
		CreateEmailTemplateOutput: r.Request.Data.(*CreateEmailTemplateOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateEmailTemplateResponse is the response type for the
// CreateEmailTemplate API operation.
type CreateEmailTemplateResponse struct {
	*CreateEmailTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateEmailTemplate request.
func (r *CreateEmailTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
