// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmailmessageflow

import (
	"context"
	"io"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetRawMessageContentInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the email message to retrieve.
	//
	// MessageId is a required field
	MessageId *string `location:"uri" locationName:"messageId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRawMessageContentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRawMessageContentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRawMessageContentInput"}

	if s.MessageId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MessageId"))
	}
	if s.MessageId != nil && len(*s.MessageId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MessageId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetRawMessageContentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MessageId != nil {
		v := *s.MessageId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "messageId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetRawMessageContentOutput struct {
	_ struct{} `type:"structure" payload:"MessageContent"`

	// The raw content of the email message, in MIME format.
	//
	// MessageContent is a required field
	MessageContent io.ReadCloser `locationName:"messageContent" type:"blob" required:"true"`
}

// String returns the string representation
func (s GetRawMessageContentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetRawMessageContentOutput) MarshalFields(e protocol.FieldEncoder) error {
	// Skipping MessageContent Output type's body not valid.
	return nil
}

const opGetRawMessageContent = "GetRawMessageContent"

// GetRawMessageContentRequest returns a request value for making API operation for
// Amazon WorkMail Message Flow.
//
// Retrieves the raw content of an in-transit email message, in MIME format.
//
//    // Example sending a request using GetRawMessageContentRequest.
//    req := client.GetRawMessageContentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmailmessageflow-2019-05-01/GetRawMessageContent
func (c *Client) GetRawMessageContentRequest(input *GetRawMessageContentInput) GetRawMessageContentRequest {
	op := &aws.Operation{
		Name:       opGetRawMessageContent,
		HTTPMethod: "GET",
		HTTPPath:   "/messages/{messageId}",
	}

	if input == nil {
		input = &GetRawMessageContentInput{}
	}

	req := c.newRequest(op, input, &GetRawMessageContentOutput{})

	return GetRawMessageContentRequest{Request: req, Input: input, Copy: c.GetRawMessageContentRequest}
}

// GetRawMessageContentRequest is the request type for the
// GetRawMessageContent API operation.
type GetRawMessageContentRequest struct {
	*aws.Request
	Input *GetRawMessageContentInput
	Copy  func(*GetRawMessageContentInput) GetRawMessageContentRequest
}

// Send marshals and sends the GetRawMessageContent API request.
func (r GetRawMessageContentRequest) Send(ctx context.Context) (*GetRawMessageContentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRawMessageContentResponse{
		GetRawMessageContentOutput: r.Request.Data.(*GetRawMessageContentOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRawMessageContentResponse is the response type for the
// GetRawMessageContent API operation.
type GetRawMessageContentResponse struct {
	*GetRawMessageContentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRawMessageContent request.
func (r *GetRawMessageContentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
