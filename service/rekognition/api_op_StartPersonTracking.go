// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type StartPersonTrackingInput struct {
	_ struct{} `type:"structure"`

	// Idempotent token used to identify the start request. If you use the same
	// token with multiple StartPersonTracking requests, the same JobId is returned.
	// Use ClientRequestToken to prevent the same job from being accidently started
	// more than once.
	ClientRequestToken *string `min:"1" type:"string"`

	// An identifier you specify that's returned in the completion notification
	// that's published to your Amazon Simple Notification Service topic. For example,
	// you can use JobTag to group related jobs and identify them in the completion
	// notification.
	JobTag *string `min:"1" type:"string"`

	// The Amazon SNS topic ARN you want Amazon Rekognition Video to publish the
	// completion status of the people detection operation to.
	NotificationChannel *NotificationChannel `type:"structure"`

	// The video in which you want to detect people. The video must be stored in
	// an Amazon S3 bucket.
	//
	// Video is a required field
	Video *Video `type:"structure" required:"true"`
}

// String returns the string representation
func (s StartPersonTrackingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartPersonTrackingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartPersonTrackingInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}
	if s.JobTag != nil && len(*s.JobTag) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobTag", 1))
	}

	if s.Video == nil {
		invalidParams.Add(aws.NewErrParamRequired("Video"))
	}
	if s.NotificationChannel != nil {
		if err := s.NotificationChannel.Validate(); err != nil {
			invalidParams.AddNested("NotificationChannel", err.(aws.ErrInvalidParams))
		}
	}
	if s.Video != nil {
		if err := s.Video.Validate(); err != nil {
			invalidParams.AddNested("Video", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartPersonTrackingOutput struct {
	_ struct{} `type:"structure"`

	// The identifier for the person detection job. Use JobId to identify the job
	// in a subsequent call to GetPersonTracking.
	JobId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StartPersonTrackingOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartPersonTracking = "StartPersonTracking"

// StartPersonTrackingRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Starts the asynchronous tracking of a person's path in a stored video.
//
// Amazon Rekognition Video can track the path of people in a video stored in
// an Amazon S3 bucket. Use Video to specify the bucket name and the filename
// of the video. StartPersonTracking returns a job identifier (JobId) which
// you use to get the results of the operation. When label detection is finished,
// Amazon Rekognition publishes a completion status to the Amazon Simple Notification
// Service topic that you specify in NotificationChannel.
//
// To get the results of the person detection operation, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED. If so, call
// GetPersonTracking and pass the job identifier (JobId) from the initial call
// to StartPersonTracking.
//
//    // Example sending a request using StartPersonTrackingRequest.
//    req := client.StartPersonTrackingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) StartPersonTrackingRequest(input *StartPersonTrackingInput) StartPersonTrackingRequest {
	op := &aws.Operation{
		Name:       opStartPersonTracking,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartPersonTrackingInput{}
	}

	req := c.newRequest(op, input, &StartPersonTrackingOutput{})

	return StartPersonTrackingRequest{Request: req, Input: input, Copy: c.StartPersonTrackingRequest}
}

// StartPersonTrackingRequest is the request type for the
// StartPersonTracking API operation.
type StartPersonTrackingRequest struct {
	*aws.Request
	Input *StartPersonTrackingInput
	Copy  func(*StartPersonTrackingInput) StartPersonTrackingRequest
}

// Send marshals and sends the StartPersonTracking API request.
func (r StartPersonTrackingRequest) Send(ctx context.Context) (*StartPersonTrackingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartPersonTrackingResponse{
		StartPersonTrackingOutput: r.Request.Data.(*StartPersonTrackingOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartPersonTrackingResponse is the response type for the
// StartPersonTracking API operation.
type StartPersonTrackingResponse struct {
	*StartPersonTrackingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartPersonTracking request.
func (r *StartPersonTrackingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
