// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListMeetingTagsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime SDK meeting ID.
	//
	// MeetingId is a required field
	MeetingId *string `location:"uri" locationName:"meetingId" type:"string" required:"true"`
}

// String returns the string representation
func (s ListMeetingTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListMeetingTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListMeetingTagsInput"}

	if s.MeetingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeetingId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListMeetingTagsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MeetingId != nil {
		v := *s.MeetingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meetingId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListMeetingTagsOutput struct {
	_ struct{} `type:"structure"`

	// A list of tag key-value pairs.
	Tags []Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s ListMeetingTagsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListMeetingTagsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListMeetingTags = "ListMeetingTags"

// ListMeetingTagsRequest returns a request value for making API operation for
// Amazon Chime.
//
// Lists the tags applied to an Amazon Chime SDK meeting resource.
//
//    // Example sending a request using ListMeetingTagsRequest.
//    req := client.ListMeetingTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/ListMeetingTags
func (c *Client) ListMeetingTagsRequest(input *ListMeetingTagsInput) ListMeetingTagsRequest {
	op := &aws.Operation{
		Name:       opListMeetingTags,
		HTTPMethod: "GET",
		HTTPPath:   "/meetings/{meetingId}/tags",
	}

	if input == nil {
		input = &ListMeetingTagsInput{}
	}

	req := c.newRequest(op, input, &ListMeetingTagsOutput{})

	return ListMeetingTagsRequest{Request: req, Input: input, Copy: c.ListMeetingTagsRequest}
}

// ListMeetingTagsRequest is the request type for the
// ListMeetingTags API operation.
type ListMeetingTagsRequest struct {
	*aws.Request
	Input *ListMeetingTagsInput
	Copy  func(*ListMeetingTagsInput) ListMeetingTagsRequest
}

// Send marshals and sends the ListMeetingTags API request.
func (r ListMeetingTagsRequest) Send(ctx context.Context) (*ListMeetingTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListMeetingTagsResponse{
		ListMeetingTagsOutput: r.Request.Data.(*ListMeetingTagsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListMeetingTagsResponse is the response type for the
// ListMeetingTags API operation.
type ListMeetingTagsResponse struct {
	*ListMeetingTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListMeetingTags request.
func (r *ListMeetingTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
