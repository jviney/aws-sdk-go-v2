// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetPullRequestOverrideStateInput struct {
	_ struct{} `type:"structure"`

	// The ID of the pull request for which you want to get information about whether
	// approval rules have been set aside (overridden).
	//
	// PullRequestId is a required field
	PullRequestId *string `locationName:"pullRequestId" type:"string" required:"true"`

	// The system-generated ID of the revision for the pull request. To retrieve
	// the most recent revision ID, use GetPullRequest.
	//
	// RevisionId is a required field
	RevisionId *string `locationName:"revisionId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPullRequestOverrideStateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPullRequestOverrideStateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPullRequestOverrideStateInput"}

	if s.PullRequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PullRequestId"))
	}

	if s.RevisionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RevisionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetPullRequestOverrideStateOutput struct {
	_ struct{} `type:"structure"`

	// A Boolean value that indicates whether a pull request has had its rules set
	// aside (TRUE) or whether all approval rules still apply (FALSE).
	Overridden *bool `locationName:"overridden" type:"boolean"`

	// The Amazon Resource Name (ARN) of the user or identity that overrode the
	// rules and their requirements for the pull request.
	Overrider *string `locationName:"overrider" type:"string"`
}

// String returns the string representation
func (s GetPullRequestOverrideStateOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetPullRequestOverrideState = "GetPullRequestOverrideState"

// GetPullRequestOverrideStateRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Returns information about whether approval rules have been set aside (overridden)
// for a pull request, and if so, the Amazon Resource Name (ARN) of the user
// or identity that overrode the rules and their requirements for the pull request.
//
//    // Example sending a request using GetPullRequestOverrideStateRequest.
//    req := client.GetPullRequestOverrideStateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetPullRequestOverrideState
func (c *Client) GetPullRequestOverrideStateRequest(input *GetPullRequestOverrideStateInput) GetPullRequestOverrideStateRequest {
	op := &aws.Operation{
		Name:       opGetPullRequestOverrideState,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetPullRequestOverrideStateInput{}
	}

	req := c.newRequest(op, input, &GetPullRequestOverrideStateOutput{})

	return GetPullRequestOverrideStateRequest{Request: req, Input: input, Copy: c.GetPullRequestOverrideStateRequest}
}

// GetPullRequestOverrideStateRequest is the request type for the
// GetPullRequestOverrideState API operation.
type GetPullRequestOverrideStateRequest struct {
	*aws.Request
	Input *GetPullRequestOverrideStateInput
	Copy  func(*GetPullRequestOverrideStateInput) GetPullRequestOverrideStateRequest
}

// Send marshals and sends the GetPullRequestOverrideState API request.
func (r GetPullRequestOverrideStateRequest) Send(ctx context.Context) (*GetPullRequestOverrideStateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPullRequestOverrideStateResponse{
		GetPullRequestOverrideStateOutput: r.Request.Data.(*GetPullRequestOverrideStateOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPullRequestOverrideStateResponse is the response type for the
// GetPullRequestOverrideState API operation.
type GetPullRequestOverrideStateResponse struct {
	*GetPullRequestOverrideStateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPullRequestOverrideState request.
func (r *GetPullRequestOverrideStateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
