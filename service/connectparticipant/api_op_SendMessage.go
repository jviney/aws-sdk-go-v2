// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connectparticipant

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type SendMessageInput struct {
	_ struct{} `type:"structure"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string `type:"string" idempotencyToken:"true"`

	// The authentication token associated with the connection.
	//
	// ConnectionToken is a required field
	ConnectionToken *string `location:"header" locationName:"X-Amz-Bearer" min:"1" type:"string" required:"true"`

	// The content of the message.
	//
	// Content is a required field
	Content *string `min:"1" type:"string" required:"true"`

	// The type of the content. Supported types are text/plain.
	//
	// ContentType is a required field
	ContentType *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s SendMessageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendMessageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SendMessageInput"}

	if s.ConnectionToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionToken"))
	}
	if s.ConnectionToken != nil && len(*s.ConnectionToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConnectionToken", 1))
	}

	if s.Content == nil {
		invalidParams.Add(aws.NewErrParamRequired("Content"))
	}
	if s.Content != nil && len(*s.Content) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Content", 1))
	}

	if s.ContentType == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContentType"))
	}
	if s.ContentType != nil && len(*s.ContentType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContentType", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SendMessageInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	var ClientToken string
	if s.ClientToken != nil {
		ClientToken = *s.ClientToken
	} else {
		ClientToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ClientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Content != nil {
		v := *s.Content

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Content", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ContentType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConnectionToken != nil {
		v := *s.ConnectionToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Amz-Bearer", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type SendMessageOutput struct {
	_ struct{} `type:"structure"`

	// The time when the message was sent.
	//
	// It's specified in ISO 8601 format: yyyy-MM-ddThh:mm:ss.SSSZ. For example,
	// 2019-11-08T02:41:28.172Z.
	AbsoluteTime *string `min:"1" type:"string"`

	// The ID of the message.
	Id *string `min:"1" type:"string"`
}

// String returns the string representation
func (s SendMessageOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SendMessageOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AbsoluteTime != nil {
		v := *s.AbsoluteTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AbsoluteTime", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opSendMessage = "SendMessage"

// SendMessageRequest returns a request value for making API operation for
// Amazon Connect Participant Service.
//
// Sends a message. Note that ConnectionToken is used for invoking this API
// instead of ParticipantToken.
//
//    // Example sending a request using SendMessageRequest.
//    req := client.SendMessageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/connectparticipant-2018-09-07/SendMessage
func (c *Client) SendMessageRequest(input *SendMessageInput) SendMessageRequest {
	op := &aws.Operation{
		Name:       opSendMessage,
		HTTPMethod: "POST",
		HTTPPath:   "/participant/message",
	}

	if input == nil {
		input = &SendMessageInput{}
	}

	req := c.newRequest(op, input, &SendMessageOutput{})

	return SendMessageRequest{Request: req, Input: input, Copy: c.SendMessageRequest}
}

// SendMessageRequest is the request type for the
// SendMessage API operation.
type SendMessageRequest struct {
	*aws.Request
	Input *SendMessageInput
	Copy  func(*SendMessageInput) SendMessageRequest
}

// Send marshals and sends the SendMessage API request.
func (r SendMessageRequest) Send(ctx context.Context) (*SendMessageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SendMessageResponse{
		SendMessageOutput: r.Request.Data.(*SendMessageOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SendMessageResponse is the response type for the
// SendMessage API operation.
type SendMessageResponse struct {
	*SendMessageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SendMessage request.
func (r *SendMessageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
