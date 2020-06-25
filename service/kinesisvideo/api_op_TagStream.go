// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideo

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type TagStreamInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the resource that you want to add the tag
	// or tags to.
	StreamARN *string `min:"1" type:"string"`

	// The name of the stream that you want to add the tag or tags to.
	StreamName *string `min:"1" type:"string"`

	// A list of tags to associate with the specified stream. Each tag is a key-value
	// pair (the value is optional).
	//
	// Tags is a required field
	Tags map[string]string `min:"1" type:"map" required:"true"`
}

// String returns the string representation
func (s TagStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagStreamInput"}
	if s.StreamARN != nil && len(*s.StreamARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamARN", 1))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TagStreamInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.StreamARN != nil {
		v := *s.StreamARN

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StreamARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StreamName != nil {
		v := *s.StreamName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StreamName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

type TagStreamOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TagStreamOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TagStreamOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opTagStream = "TagStream"

// TagStreamRequest returns a request value for making API operation for
// Amazon Kinesis Video Streams.
//
// Adds one or more tags to a stream. A tag is a key-value pair (the value is
// optional) that you can define and assign to AWS resources. If you specify
// a tag that already exists, the tag value is replaced with the value that
// you specify in the request. For more information, see Using Cost Allocation
// Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
// in the AWS Billing and Cost Management User Guide.
//
// You must provide either the StreamName or the StreamARN.
//
// This operation requires permission for the KinesisVideo:TagStream action.
//
// Kinesis video streams support up to 50 tags.
//
//    // Example sending a request using TagStreamRequest.
//    req := client.TagStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisvideo-2017-09-30/TagStream
func (c *Client) TagStreamRequest(input *TagStreamInput) TagStreamRequest {
	op := &aws.Operation{
		Name:       opTagStream,
		HTTPMethod: "POST",
		HTTPPath:   "/tagStream",
	}

	if input == nil {
		input = &TagStreamInput{}
	}

	req := c.newRequest(op, input, &TagStreamOutput{})

	return TagStreamRequest{Request: req, Input: input, Copy: c.TagStreamRequest}
}

// TagStreamRequest is the request type for the
// TagStream API operation.
type TagStreamRequest struct {
	*aws.Request
	Input *TagStreamInput
	Copy  func(*TagStreamInput) TagStreamRequest
}

// Send marshals and sends the TagStream API request.
func (r TagStreamRequest) Send(ctx context.Context) (*TagStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TagStreamResponse{
		TagStreamOutput: r.Request.Data.(*TagStreamOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TagStreamResponse is the response type for the
// TagStream API operation.
type TagStreamResponse struct {
	*TagStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TagStream request.
func (r *TagStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
