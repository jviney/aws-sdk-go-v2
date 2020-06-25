// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type StartTextDetectionInput struct {
	_ struct{} `type:"structure"`

	// Idempotent token used to identify the start request. If you use the same
	// token with multiple StartTextDetection requests, the same JobId is returned.
	// Use ClientRequestToken to prevent the same job from being accidentaly started
	// more than once.
	ClientRequestToken *string `min:"1" type:"string"`

	// Optional parameters that let you set criteria the text must meet to be included
	// in your response.
	Filters *StartTextDetectionFilters `type:"structure"`

	// An identifier returned in the completion status published by your Amazon
	// Simple Notification Service topic. For example, you can use JobTag to group
	// related jobs and identify them in the completion notification.
	JobTag *string `min:"1" type:"string"`

	// The Amazon Simple Notification Service topic to which Amazon Rekognition
	// publishes the completion status of a video analysis operation. For more information,
	// see api-video.
	NotificationChannel *NotificationChannel `type:"structure"`

	// Video file stored in an Amazon S3 bucket. Amazon Rekognition video start
	// operations such as StartLabelDetection use Video to specify a video for analysis.
	// The supported file formats are .mp4, .mov and .avi.
	//
	// Video is a required field
	Video *Video `type:"structure" required:"true"`
}

// String returns the string representation
func (s StartTextDetectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartTextDetectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartTextDetectionInput"}
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

type StartTextDetectionOutput struct {
	_ struct{} `type:"structure"`

	// Identifier for the text detection job. Use JobId to identify the job in a
	// subsequent call to GetTextDetection.
	JobId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StartTextDetectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartTextDetection = "StartTextDetection"

// StartTextDetectionRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Starts asynchronous detection of text in a stored video.
//
// Amazon Rekognition Video can detect text in a video stored in an Amazon S3
// bucket. Use Video to specify the bucket name and the filename of the video.
// StartTextDetection returns a job identifier (JobId) which you use to get
// the results of the operation. When text detection is finished, Amazon Rekognition
// Video publishes a completion status to the Amazon Simple Notification Service
// topic that you specify in NotificationChannel.
//
// To get the results of the text detection operation, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED. if so, call
// GetTextDetection and pass the job identifier (JobId) from the initial call
// to StartTextDetection.
//
//    // Example sending a request using StartTextDetectionRequest.
//    req := client.StartTextDetectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) StartTextDetectionRequest(input *StartTextDetectionInput) StartTextDetectionRequest {
	op := &aws.Operation{
		Name:       opStartTextDetection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartTextDetectionInput{}
	}

	req := c.newRequest(op, input, &StartTextDetectionOutput{})

	return StartTextDetectionRequest{Request: req, Input: input, Copy: c.StartTextDetectionRequest}
}

// StartTextDetectionRequest is the request type for the
// StartTextDetection API operation.
type StartTextDetectionRequest struct {
	*aws.Request
	Input *StartTextDetectionInput
	Copy  func(*StartTextDetectionInput) StartTextDetectionRequest
}

// Send marshals and sends the StartTextDetection API request.
func (r StartTextDetectionRequest) Send(ctx context.Context) (*StartTextDetectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartTextDetectionResponse{
		StartTextDetectionOutput: r.Request.Data.(*StartTextDetectionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartTextDetectionResponse is the response type for the
// StartTextDetection API operation.
type StartTextDetectionResponse struct {
	*StartTextDetectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartTextDetection request.
func (r *StartTextDetectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
