// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type CancelAuditTaskInput struct {
	_ struct{} `type:"structure"`

	// The ID of the audit you want to cancel. You can only cancel an audit that
	// is "IN_PROGRESS".
	//
	// TaskId is a required field
	TaskId *string `location:"uri" locationName:"taskId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CancelAuditTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelAuditTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelAuditTaskInput"}

	if s.TaskId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskId"))
	}
	if s.TaskId != nil && len(*s.TaskId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TaskId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelAuditTaskInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.TaskId != nil {
		v := *s.TaskId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "taskId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CancelAuditTaskOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CancelAuditTaskOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelAuditTaskOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCancelAuditTask = "CancelAuditTask"

// CancelAuditTaskRequest returns a request value for making API operation for
// AWS IoT.
//
// Cancels an audit that is in progress. The audit can be either scheduled or
// on-demand. If the audit is not in progress, an "InvalidRequestException"
// occurs.
//
//    // Example sending a request using CancelAuditTaskRequest.
//    req := client.CancelAuditTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CancelAuditTaskRequest(input *CancelAuditTaskInput) CancelAuditTaskRequest {
	op := &aws.Operation{
		Name:       opCancelAuditTask,
		HTTPMethod: "PUT",
		HTTPPath:   "/audit/tasks/{taskId}/cancel",
	}

	if input == nil {
		input = &CancelAuditTaskInput{}
	}

	req := c.newRequest(op, input, &CancelAuditTaskOutput{})

	return CancelAuditTaskRequest{Request: req, Input: input, Copy: c.CancelAuditTaskRequest}
}

// CancelAuditTaskRequest is the request type for the
// CancelAuditTask API operation.
type CancelAuditTaskRequest struct {
	*aws.Request
	Input *CancelAuditTaskInput
	Copy  func(*CancelAuditTaskInput) CancelAuditTaskRequest
}

// Send marshals and sends the CancelAuditTask API request.
func (r CancelAuditTaskRequest) Send(ctx context.Context) (*CancelAuditTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelAuditTaskResponse{
		CancelAuditTaskOutput: r.Request.Data.(*CancelAuditTaskOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelAuditTaskResponse is the response type for the
// CancelAuditTask API operation.
type CancelAuditTaskResponse struct {
	*CancelAuditTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelAuditTask request.
func (r *CancelAuditTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
