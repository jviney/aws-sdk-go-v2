// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CancelMLTaskRunInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the task run.
	//
	// TaskRunId is a required field
	TaskRunId *string `min:"1" type:"string" required:"true"`

	// The unique identifier of the machine learning transform.
	//
	// TransformId is a required field
	TransformId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CancelMLTaskRunInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelMLTaskRunInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelMLTaskRunInput"}

	if s.TaskRunId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskRunId"))
	}
	if s.TaskRunId != nil && len(*s.TaskRunId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TaskRunId", 1))
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

type CancelMLTaskRunOutput struct {
	_ struct{} `type:"structure"`

	// The status for this run.
	Status TaskStatusType `type:"string" enum:"true"`

	// The unique identifier for the task run.
	TaskRunId *string `min:"1" type:"string"`

	// The unique identifier of the machine learning transform.
	TransformId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CancelMLTaskRunOutput) String() string {
	return awsutil.Prettify(s)
}

const opCancelMLTaskRun = "CancelMLTaskRun"

// CancelMLTaskRunRequest returns a request value for making API operation for
// AWS Glue.
//
// Cancels (stops) a task run. Machine learning task runs are asynchronous tasks
// that AWS Glue runs on your behalf as part of various machine learning workflows.
// You can cancel a machine learning task run at any time by calling CancelMLTaskRun
// with a task run's parent transform's TransformID and the task run's TaskRunId.
//
//    // Example sending a request using CancelMLTaskRunRequest.
//    req := client.CancelMLTaskRunRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/CancelMLTaskRun
func (c *Client) CancelMLTaskRunRequest(input *CancelMLTaskRunInput) CancelMLTaskRunRequest {
	op := &aws.Operation{
		Name:       opCancelMLTaskRun,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CancelMLTaskRunInput{}
	}

	req := c.newRequest(op, input, &CancelMLTaskRunOutput{})

	return CancelMLTaskRunRequest{Request: req, Input: input, Copy: c.CancelMLTaskRunRequest}
}

// CancelMLTaskRunRequest is the request type for the
// CancelMLTaskRun API operation.
type CancelMLTaskRunRequest struct {
	*aws.Request
	Input *CancelMLTaskRunInput
	Copy  func(*CancelMLTaskRunInput) CancelMLTaskRunRequest
}

// Send marshals and sends the CancelMLTaskRun API request.
func (r CancelMLTaskRunRequest) Send(ctx context.Context) (*CancelMLTaskRunResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelMLTaskRunResponse{
		CancelMLTaskRunOutput: r.Request.Data.(*CancelMLTaskRunOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelMLTaskRunResponse is the response type for the
// CancelMLTaskRun API operation.
type CancelMLTaskRunResponse struct {
	*CancelMLTaskRunOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelMLTaskRun request.
func (r *CancelMLTaskRunResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
