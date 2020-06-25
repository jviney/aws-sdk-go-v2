// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideo

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type UpdateSignalingChannelInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the signaling channel that you want to
	// update.
	//
	// ChannelARN is a required field
	ChannelARN *string `min:"1" type:"string" required:"true"`

	// The current version of the signaling channel that you want to update.
	//
	// CurrentVersion is a required field
	CurrentVersion *string `min:"1" type:"string" required:"true"`

	// The structure containing the configuration for the SINGLE_MASTER type of
	// the signaling channel that you want to update.
	SingleMasterConfiguration *SingleMasterConfiguration `type:"structure"`
}

// String returns the string representation
func (s UpdateSignalingChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSignalingChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSignalingChannelInput"}

	if s.ChannelARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChannelARN"))
	}
	if s.ChannelARN != nil && len(*s.ChannelARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChannelARN", 1))
	}

	if s.CurrentVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("CurrentVersion"))
	}
	if s.CurrentVersion != nil && len(*s.CurrentVersion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CurrentVersion", 1))
	}
	if s.SingleMasterConfiguration != nil {
		if err := s.SingleMasterConfiguration.Validate(); err != nil {
			invalidParams.AddNested("SingleMasterConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateSignalingChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ChannelARN != nil {
		v := *s.ChannelARN

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ChannelARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CurrentVersion != nil {
		v := *s.CurrentVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CurrentVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SingleMasterConfiguration != nil {
		v := s.SingleMasterConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SingleMasterConfiguration", v, metadata)
	}
	return nil
}

type UpdateSignalingChannelOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateSignalingChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateSignalingChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateSignalingChannel = "UpdateSignalingChannel"

// UpdateSignalingChannelRequest returns a request value for making API operation for
// Amazon Kinesis Video Streams.
//
// Updates the existing signaling channel. This is an asynchronous operation
// and takes time to complete.
//
// If the MessageTtlSeconds value is updated (either increased or reduced),
// it only applies to new messages sent via this channel after it's been updated.
// Existing messages are still expired as per the previous MessageTtlSeconds
// value.
//
//    // Example sending a request using UpdateSignalingChannelRequest.
//    req := client.UpdateSignalingChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisvideo-2017-09-30/UpdateSignalingChannel
func (c *Client) UpdateSignalingChannelRequest(input *UpdateSignalingChannelInput) UpdateSignalingChannelRequest {
	op := &aws.Operation{
		Name:       opUpdateSignalingChannel,
		HTTPMethod: "POST",
		HTTPPath:   "/updateSignalingChannel",
	}

	if input == nil {
		input = &UpdateSignalingChannelInput{}
	}

	req := c.newRequest(op, input, &UpdateSignalingChannelOutput{})

	return UpdateSignalingChannelRequest{Request: req, Input: input, Copy: c.UpdateSignalingChannelRequest}
}

// UpdateSignalingChannelRequest is the request type for the
// UpdateSignalingChannel API operation.
type UpdateSignalingChannelRequest struct {
	*aws.Request
	Input *UpdateSignalingChannelInput
	Copy  func(*UpdateSignalingChannelInput) UpdateSignalingChannelRequest
}

// Send marshals and sends the UpdateSignalingChannel API request.
func (r UpdateSignalingChannelRequest) Send(ctx context.Context) (*UpdateSignalingChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSignalingChannelResponse{
		UpdateSignalingChannelOutput: r.Request.Data.(*UpdateSignalingChannelOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSignalingChannelResponse is the response type for the
// UpdateSignalingChannel API operation.
type UpdateSignalingChannelResponse struct {
	*UpdateSignalingChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSignalingChannel request.
func (r *UpdateSignalingChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
