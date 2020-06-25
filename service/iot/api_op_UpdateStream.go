// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type UpdateStreamInput struct {
	_ struct{} `type:"structure"`

	// The description of the stream.
	Description *string `locationName:"description" type:"string"`

	// The files associated with the stream.
	Files []StreamFile `locationName:"files" min:"1" type:"list"`

	// An IAM role that allows the IoT service principal assumes to access your
	// S3 files.
	RoleArn *string `locationName:"roleArn" min:"20" type:"string"`

	// The stream ID.
	//
	// StreamId is a required field
	StreamId *string `location:"uri" locationName:"streamId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateStreamInput"}
	if s.Files != nil && len(s.Files) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Files", 1))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 20))
	}

	if s.StreamId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamId"))
	}
	if s.StreamId != nil && len(*s.StreamId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamId", 1))
	}
	if s.Files != nil {
		for i, v := range s.Files {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Files", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateStreamInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Files != nil {
		v := s.Files

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "files", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.RoleArn != nil {
		v := *s.RoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "roleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StreamId != nil {
		v := *s.StreamId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "streamId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateStreamOutput struct {
	_ struct{} `type:"structure"`

	// A description of the stream.
	Description *string `locationName:"description" type:"string"`

	// The stream ARN.
	StreamArn *string `locationName:"streamArn" type:"string"`

	// The stream ID.
	StreamId *string `locationName:"streamId" min:"1" type:"string"`

	// The stream version.
	StreamVersion *int64 `locationName:"streamVersion" type:"integer"`
}

// String returns the string representation
func (s UpdateStreamOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateStreamOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StreamArn != nil {
		v := *s.StreamArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "streamArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StreamId != nil {
		v := *s.StreamId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "streamId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StreamVersion != nil {
		v := *s.StreamVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "streamVersion", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opUpdateStream = "UpdateStream"

// UpdateStreamRequest returns a request value for making API operation for
// AWS IoT.
//
// Updates an existing stream. The stream version will be incremented by one.
//
//    // Example sending a request using UpdateStreamRequest.
//    req := client.UpdateStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateStreamRequest(input *UpdateStreamInput) UpdateStreamRequest {
	op := &aws.Operation{
		Name:       opUpdateStream,
		HTTPMethod: "PUT",
		HTTPPath:   "/streams/{streamId}",
	}

	if input == nil {
		input = &UpdateStreamInput{}
	}

	req := c.newRequest(op, input, &UpdateStreamOutput{})

	return UpdateStreamRequest{Request: req, Input: input, Copy: c.UpdateStreamRequest}
}

// UpdateStreamRequest is the request type for the
// UpdateStream API operation.
type UpdateStreamRequest struct {
	*aws.Request
	Input *UpdateStreamInput
	Copy  func(*UpdateStreamInput) UpdateStreamRequest
}

// Send marshals and sends the UpdateStream API request.
func (r UpdateStreamRequest) Send(ctx context.Context) (*UpdateStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateStreamResponse{
		UpdateStreamOutput: r.Request.Data.(*UpdateStreamOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateStreamResponse is the response type for the
// UpdateStream API operation.
type UpdateStreamResponse struct {
	*UpdateStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateStream request.
func (r *UpdateStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
