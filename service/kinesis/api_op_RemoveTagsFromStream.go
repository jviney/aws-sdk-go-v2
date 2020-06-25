// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Represents the input for RemoveTagsFromStream.
type RemoveTagsFromStreamInput struct {
	_ struct{} `type:"structure"`

	// The name of the stream.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`

	// A list of tag keys. Each corresponding tag is removed from the stream.
	//
	// TagKeys is a required field
	TagKeys []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s RemoveTagsFromStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveTagsFromStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveTagsFromStreamInput"}

	if s.StreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamName"))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}

	if s.TagKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagKeys"))
	}
	if s.TagKeys != nil && len(s.TagKeys) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TagKeys", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RemoveTagsFromStreamOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemoveTagsFromStreamOutput) String() string {
	return awsutil.Prettify(s)
}

const opRemoveTagsFromStream = "RemoveTagsFromStream"

// RemoveTagsFromStreamRequest returns a request value for making API operation for
// Amazon Kinesis.
//
// Removes tags from the specified Kinesis data stream. Removed tags are deleted
// and cannot be recovered after this operation successfully completes.
//
// If you specify a tag that does not exist, it is ignored.
//
// RemoveTagsFromStream has a limit of five transactions per second per account.
//
//    // Example sending a request using RemoveTagsFromStreamRequest.
//    req := client.RemoveTagsFromStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/RemoveTagsFromStream
func (c *Client) RemoveTagsFromStreamRequest(input *RemoveTagsFromStreamInput) RemoveTagsFromStreamRequest {
	op := &aws.Operation{
		Name:       opRemoveTagsFromStream,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveTagsFromStreamInput{}
	}

	req := c.newRequest(op, input, &RemoveTagsFromStreamOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return RemoveTagsFromStreamRequest{Request: req, Input: input, Copy: c.RemoveTagsFromStreamRequest}
}

// RemoveTagsFromStreamRequest is the request type for the
// RemoveTagsFromStream API operation.
type RemoveTagsFromStreamRequest struct {
	*aws.Request
	Input *RemoveTagsFromStreamInput
	Copy  func(*RemoveTagsFromStreamInput) RemoveTagsFromStreamRequest
}

// Send marshals and sends the RemoveTagsFromStream API request.
func (r RemoveTagsFromStreamRequest) Send(ctx context.Context) (*RemoveTagsFromStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveTagsFromStreamResponse{
		RemoveTagsFromStreamOutput: r.Request.Data.(*RemoveTagsFromStreamOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveTagsFromStreamResponse is the response type for the
// RemoveTagsFromStream API operation.
type RemoveTagsFromStreamResponse struct {
	*RemoveTagsFromStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveTagsFromStream request.
func (r *RemoveTagsFromStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
