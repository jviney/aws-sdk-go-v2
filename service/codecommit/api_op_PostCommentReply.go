// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type PostCommentReplyInput struct {
	_ struct{} `type:"structure"`

	// A unique, client-generated idempotency token that, when provided in a request,
	// ensures the request cannot be repeated with a changed parameter. If a request
	// is received with the same parameters and a token is included, the request
	// returns information about the initial request that used that token.
	ClientRequestToken *string `locationName:"clientRequestToken" type:"string" idempotencyToken:"true"`

	// The contents of your reply to a comment.
	//
	// Content is a required field
	Content *string `locationName:"content" type:"string" required:"true"`

	// The system-generated ID of the comment to which you want to reply. To get
	// this ID, use GetCommentsForComparedCommit or GetCommentsForPullRequest.
	//
	// InReplyTo is a required field
	InReplyTo *string `locationName:"inReplyTo" type:"string" required:"true"`
}

// String returns the string representation
func (s PostCommentReplyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PostCommentReplyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PostCommentReplyInput"}

	if s.Content == nil {
		invalidParams.Add(aws.NewErrParamRequired("Content"))
	}

	if s.InReplyTo == nil {
		invalidParams.Add(aws.NewErrParamRequired("InReplyTo"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PostCommentReplyOutput struct {
	_ struct{} `type:"structure"`

	// Information about the reply to a comment.
	Comment *Comment `locationName:"comment" type:"structure"`
}

// String returns the string representation
func (s PostCommentReplyOutput) String() string {
	return awsutil.Prettify(s)
}

const opPostCommentReply = "PostCommentReply"

// PostCommentReplyRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Posts a comment in reply to an existing comment on a comparison between commits
// or a pull request.
//
//    // Example sending a request using PostCommentReplyRequest.
//    req := client.PostCommentReplyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/PostCommentReply
func (c *Client) PostCommentReplyRequest(input *PostCommentReplyInput) PostCommentReplyRequest {
	op := &aws.Operation{
		Name:       opPostCommentReply,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PostCommentReplyInput{}
	}

	req := c.newRequest(op, input, &PostCommentReplyOutput{})

	return PostCommentReplyRequest{Request: req, Input: input, Copy: c.PostCommentReplyRequest}
}

// PostCommentReplyRequest is the request type for the
// PostCommentReply API operation.
type PostCommentReplyRequest struct {
	*aws.Request
	Input *PostCommentReplyInput
	Copy  func(*PostCommentReplyInput) PostCommentReplyRequest
}

// Send marshals and sends the PostCommentReply API request.
func (r PostCommentReplyRequest) Send(ctx context.Context) (*PostCommentReplyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PostCommentReplyResponse{
		PostCommentReplyOutput: r.Request.Data.(*PostCommentReplyOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PostCommentReplyResponse is the response type for the
// PostCommentReply API operation.
type PostCommentReplyResponse struct {
	*PostCommentReplyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PostCommentReply request.
func (r *PostCommentReplyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
