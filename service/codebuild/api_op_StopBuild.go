// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type StopBuildInput struct {
	_ struct{} `type:"structure"`

	// The ID of the build.
	//
	// Id is a required field
	Id *string `locationName:"id" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopBuildInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopBuildInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopBuildInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopBuildOutput struct {
	_ struct{} `type:"structure"`

	// Information about the build.
	Build *Build `locationName:"build" type:"structure"`
}

// String returns the string representation
func (s StopBuildOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopBuild = "StopBuild"

// StopBuildRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// Attempts to stop running a build.
//
//    // Example sending a request using StopBuildRequest.
//    req := client.StopBuildRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/StopBuild
func (c *Client) StopBuildRequest(input *StopBuildInput) StopBuildRequest {
	op := &aws.Operation{
		Name:       opStopBuild,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopBuildInput{}
	}

	req := c.newRequest(op, input, &StopBuildOutput{})

	return StopBuildRequest{Request: req, Input: input, Copy: c.StopBuildRequest}
}

// StopBuildRequest is the request type for the
// StopBuild API operation.
type StopBuildRequest struct {
	*aws.Request
	Input *StopBuildInput
	Copy  func(*StopBuildInput) StopBuildRequest
}

// Send marshals and sends the StopBuild API request.
func (r StopBuildRequest) Send(ctx context.Context) (*StopBuildResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopBuildResponse{
		StopBuildOutput: r.Request.Data.(*StopBuildOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopBuildResponse is the response type for the
// StopBuild API operation.
type StopBuildResponse struct {
	*StopBuildOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopBuild request.
func (r *StopBuildResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
