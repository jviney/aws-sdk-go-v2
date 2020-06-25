// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type StopTrainingEntityRecognizerInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that identifies the entity recognizer currently
	// being trained.
	//
	// EntityRecognizerArn is a required field
	EntityRecognizerArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StopTrainingEntityRecognizerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopTrainingEntityRecognizerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopTrainingEntityRecognizerInput"}

	if s.EntityRecognizerArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntityRecognizerArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopTrainingEntityRecognizerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopTrainingEntityRecognizerOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopTrainingEntityRecognizer = "StopTrainingEntityRecognizer"

// StopTrainingEntityRecognizerRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Stops an entity recognizer training job while in progress.
//
// If the training job state is TRAINING, the job is marked for termination
// and put into the STOP_REQUESTED state. If the training job completes before
// it can be stopped, it is put into the TRAINED; otherwise the training job
// is stopped and putted into the STOPPED state and the service sends back an
// HTTP 200 response with an empty HTTP body.
//
//    // Example sending a request using StopTrainingEntityRecognizerRequest.
//    req := client.StopTrainingEntityRecognizerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/StopTrainingEntityRecognizer
func (c *Client) StopTrainingEntityRecognizerRequest(input *StopTrainingEntityRecognizerInput) StopTrainingEntityRecognizerRequest {
	op := &aws.Operation{
		Name:       opStopTrainingEntityRecognizer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopTrainingEntityRecognizerInput{}
	}

	req := c.newRequest(op, input, &StopTrainingEntityRecognizerOutput{})

	return StopTrainingEntityRecognizerRequest{Request: req, Input: input, Copy: c.StopTrainingEntityRecognizerRequest}
}

// StopTrainingEntityRecognizerRequest is the request type for the
// StopTrainingEntityRecognizer API operation.
type StopTrainingEntityRecognizerRequest struct {
	*aws.Request
	Input *StopTrainingEntityRecognizerInput
	Copy  func(*StopTrainingEntityRecognizerInput) StopTrainingEntityRecognizerRequest
}

// Send marshals and sends the StopTrainingEntityRecognizer API request.
func (r StopTrainingEntityRecognizerRequest) Send(ctx context.Context) (*StopTrainingEntityRecognizerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopTrainingEntityRecognizerResponse{
		StopTrainingEntityRecognizerOutput: r.Request.Data.(*StopTrainingEntityRecognizerOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopTrainingEntityRecognizerResponse is the response type for the
// StopTrainingEntityRecognizer API operation.
type StopTrainingEntityRecognizerResponse struct {
	*StopTrainingEntityRecognizerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopTrainingEntityRecognizer request.
func (r *StopTrainingEntityRecognizerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
