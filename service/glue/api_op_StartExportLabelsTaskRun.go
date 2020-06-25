// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type StartExportLabelsTaskRunInput struct {
	_ struct{} `type:"structure"`

	// The Amazon S3 path where you export the labels.
	//
	// OutputS3Path is a required field
	OutputS3Path *string `type:"string" required:"true"`

	// The unique identifier of the machine learning transform.
	//
	// TransformId is a required field
	TransformId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartExportLabelsTaskRunInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartExportLabelsTaskRunInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartExportLabelsTaskRunInput"}

	if s.OutputS3Path == nil {
		invalidParams.Add(aws.NewErrParamRequired("OutputS3Path"))
	}

	if s.TransformId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransformId"))
	}
	if s.TransformId != nil && len(*s.TransformId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TransformId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartExportLabelsTaskRunOutput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the task run.
	TaskRunId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StartExportLabelsTaskRunOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartExportLabelsTaskRun = "StartExportLabelsTaskRun"

// StartExportLabelsTaskRunRequest returns a request value for making API operation for
// AWS Glue.
//
// Begins an asynchronous task to export all labeled data for a particular transform.
// This task is the only label-related API call that is not part of the typical
// active learning workflow. You typically use StartExportLabelsTaskRun when
// you want to work with all of your existing labels at the same time, such
// as when you want to remove or change labels that were previously submitted
// as truth. This API operation accepts the TransformId whose labels you want
// to export and an Amazon Simple Storage Service (Amazon S3) path to export
// the labels to. The operation returns a TaskRunId. You can check on the status
// of your task run by calling the GetMLTaskRun API.
//
//    // Example sending a request using StartExportLabelsTaskRunRequest.
//    req := client.StartExportLabelsTaskRunRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/StartExportLabelsTaskRun
func (c *Client) StartExportLabelsTaskRunRequest(input *StartExportLabelsTaskRunInput) StartExportLabelsTaskRunRequest {
	op := &aws.Operation{
		Name:       opStartExportLabelsTaskRun,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartExportLabelsTaskRunInput{}
	}

	req := c.newRequest(op, input, &StartExportLabelsTaskRunOutput{})

	return StartExportLabelsTaskRunRequest{Request: req, Input: input, Copy: c.StartExportLabelsTaskRunRequest}
}

// StartExportLabelsTaskRunRequest is the request type for the
// StartExportLabelsTaskRun API operation.
type StartExportLabelsTaskRunRequest struct {
	*aws.Request
	Input *StartExportLabelsTaskRunInput
	Copy  func(*StartExportLabelsTaskRunInput) StartExportLabelsTaskRunRequest
}

// Send marshals and sends the StartExportLabelsTaskRun API request.
func (r StartExportLabelsTaskRunRequest) Send(ctx context.Context) (*StartExportLabelsTaskRunResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartExportLabelsTaskRunResponse{
		StartExportLabelsTaskRunOutput: r.Request.Data.(*StartExportLabelsTaskRunOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartExportLabelsTaskRunResponse is the response type for the
// StartExportLabelsTaskRun API operation.
type StartExportLabelsTaskRunResponse struct {
	*StartExportLabelsTaskRunOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartExportLabelsTaskRun request.
func (r *StartExportLabelsTaskRunResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
