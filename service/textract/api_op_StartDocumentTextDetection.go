// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package textract

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type StartDocumentTextDetectionInput struct {
	_ struct{} `type:"structure"`

	// The idempotent token that's used to identify the start request. If you use
	// the same token with multiple StartDocumentTextDetection requests, the same
	// JobId is returned. Use ClientRequestToken to prevent the same job from being
	// accidentally started more than once. For more information, see Calling Amazon
	// Textract Asynchronous Operations (https://docs.aws.amazon.com/textract/latest/dg/api-async.html).
	ClientRequestToken *string `min:"1" type:"string"`

	// The location of the document to be processed.
	//
	// DocumentLocation is a required field
	DocumentLocation *DocumentLocation `type:"structure" required:"true"`

	// An identifier that you specify that's included in the completion notification
	// published to the Amazon SNS topic. For example, you can use JobTag to identify
	// the type of document that the completion notification corresponds to (such
	// as a tax form or a receipt).
	JobTag *string `min:"1" type:"string"`

	// The Amazon SNS topic ARN that you want Amazon Textract to publish the completion
	// status of the operation to.
	NotificationChannel *NotificationChannel `type:"structure"`
}

// String returns the string representation
func (s StartDocumentTextDetectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartDocumentTextDetectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartDocumentTextDetectionInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.DocumentLocation == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentLocation"))
	}
	if s.JobTag != nil && len(*s.JobTag) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobTag", 1))
	}
	if s.DocumentLocation != nil {
		if err := s.DocumentLocation.Validate(); err != nil {
			invalidParams.AddNested("DocumentLocation", err.(aws.ErrInvalidParams))
		}
	}
	if s.NotificationChannel != nil {
		if err := s.NotificationChannel.Validate(); err != nil {
			invalidParams.AddNested("NotificationChannel", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartDocumentTextDetectionOutput struct {
	_ struct{} `type:"structure"`

	// The identifier of the text detection job for the document. Use JobId to identify
	// the job in a subsequent call to GetDocumentTextDetection. A JobId value is
	// only valid for 7 days.
	JobId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StartDocumentTextDetectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartDocumentTextDetection = "StartDocumentTextDetection"

// StartDocumentTextDetectionRequest returns a request value for making API operation for
// Amazon Textract.
//
// Starts the asynchronous detection of text in a document. Amazon Textract
// can detect lines of text and the words that make up a line of text.
//
// StartDocumentTextDetection can analyze text in documents that are in JPEG,
// PNG, and PDF format. The documents are stored in an Amazon S3 bucket. Use
// DocumentLocation to specify the bucket name and file name of the document.
//
// StartTextDetection returns a job identifier (JobId) that you use to get the
// results of the operation. When text detection is finished, Amazon Textract
// publishes a completion status to the Amazon Simple Notification Service (Amazon
// SNS) topic that you specify in NotificationChannel. To get the results of
// the text detection operation, first check that the status value published
// to the Amazon SNS topic is SUCCEEDED. If so, call GetDocumentTextDetection,
// and pass the job identifier (JobId) from the initial call to StartDocumentTextDetection.
//
// For more information, see Document Text Detection (https://docs.aws.amazon.com/textract/latest/dg/how-it-works-detecting.html).
//
//    // Example sending a request using StartDocumentTextDetectionRequest.
//    req := client.StartDocumentTextDetectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/textract-2018-06-27/StartDocumentTextDetection
func (c *Client) StartDocumentTextDetectionRequest(input *StartDocumentTextDetectionInput) StartDocumentTextDetectionRequest {
	op := &aws.Operation{
		Name:       opStartDocumentTextDetection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartDocumentTextDetectionInput{}
	}

	req := c.newRequest(op, input, &StartDocumentTextDetectionOutput{})

	return StartDocumentTextDetectionRequest{Request: req, Input: input, Copy: c.StartDocumentTextDetectionRequest}
}

// StartDocumentTextDetectionRequest is the request type for the
// StartDocumentTextDetection API operation.
type StartDocumentTextDetectionRequest struct {
	*aws.Request
	Input *StartDocumentTextDetectionInput
	Copy  func(*StartDocumentTextDetectionInput) StartDocumentTextDetectionRequest
}

// Send marshals and sends the StartDocumentTextDetection API request.
func (r StartDocumentTextDetectionRequest) Send(ctx context.Context) (*StartDocumentTextDetectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartDocumentTextDetectionResponse{
		StartDocumentTextDetectionOutput: r.Request.Data.(*StartDocumentTextDetectionOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartDocumentTextDetectionResponse is the response type for the
// StartDocumentTextDetection API operation.
type StartDocumentTextDetectionResponse struct {
	*StartDocumentTextDetectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartDocumentTextDetection request.
func (r *StartDocumentTextDetectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
