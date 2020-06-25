// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type UpdatePullRequestTitleInput struct {
	_ struct{} `type:"structure"`

	// The system-generated ID of the pull request. To get this ID, use ListPullRequests.
	//
	// PullRequestId is a required field
	PullRequestId *string `locationName:"pullRequestId" type:"string" required:"true"`

	// The updated title of the pull request. This replaces the existing title.
	//
	// Title is a required field
	Title *string `locationName:"title" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdatePullRequestTitleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdatePullRequestTitleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdatePullRequestTitleInput"}

	if s.PullRequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PullRequestId"))
	}

	if s.Title == nil {
		invalidParams.Add(aws.NewErrParamRequired("Title"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdatePullRequestTitleOutput struct {
	_ struct{} `type:"structure"`

	// Information about the updated pull request.
	//
	// PullRequest is a required field
	PullRequest *PullRequest `locationName:"pullRequest" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdatePullRequestTitleOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdatePullRequestTitle = "UpdatePullRequestTitle"

// UpdatePullRequestTitleRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Replaces the title of a pull request.
//
//    // Example sending a request using UpdatePullRequestTitleRequest.
//    req := client.UpdatePullRequestTitleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/UpdatePullRequestTitle
func (c *Client) UpdatePullRequestTitleRequest(input *UpdatePullRequestTitleInput) UpdatePullRequestTitleRequest {
	op := &aws.Operation{
		Name:       opUpdatePullRequestTitle,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdatePullRequestTitleInput{}
	}

	req := c.newRequest(op, input, &UpdatePullRequestTitleOutput{})

	return UpdatePullRequestTitleRequest{Request: req, Input: input, Copy: c.UpdatePullRequestTitleRequest}
}

// UpdatePullRequestTitleRequest is the request type for the
// UpdatePullRequestTitle API operation.
type UpdatePullRequestTitleRequest struct {
	*aws.Request
	Input *UpdatePullRequestTitleInput
	Copy  func(*UpdatePullRequestTitleInput) UpdatePullRequestTitleRequest
}

// Send marshals and sends the UpdatePullRequestTitle API request.
func (r UpdatePullRequestTitleRequest) Send(ctx context.Context) (*UpdatePullRequestTitleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdatePullRequestTitleResponse{
		UpdatePullRequestTitleOutput: r.Request.Data.(*UpdatePullRequestTitleOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdatePullRequestTitleResponse is the response type for the
// UpdatePullRequestTitle API operation.
type UpdatePullRequestTitleResponse struct {
	*UpdatePullRequestTitleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdatePullRequestTitle request.
func (r *UpdatePullRequestTitleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
