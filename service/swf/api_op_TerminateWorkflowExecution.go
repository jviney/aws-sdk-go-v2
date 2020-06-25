// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type TerminateWorkflowExecutionInput struct {
	_ struct{} `type:"structure"`

	// If set, specifies the policy to use for the child workflow executions of
	// the workflow execution being terminated. This policy overrides the child
	// policy specified for the workflow execution at registration time or when
	// starting the execution.
	//
	// The supported child policies are:
	//
	//    * TERMINATE – The child executions are terminated.
	//
	//    * REQUEST_CANCEL – A request to cancel is attempted for each child execution
	//    by recording a WorkflowExecutionCancelRequested event in its history.
	//    It is up to the decider to take appropriate actions when it receives an
	//    execution history with this event.
	//
	//    * ABANDON – No action is taken. The child executions continue to run.
	//
	// A child policy for this workflow execution must be specified either as a
	// default for the workflow type or through this parameter. If neither this
	// parameter is set nor a default child policy was specified at registration
	// time then a fault is returned.
	ChildPolicy ChildPolicy `locationName:"childPolicy" type:"string" enum:"true"`

	// Details for terminating the workflow execution.
	Details *string `locationName:"details" type:"string"`

	// The domain of the workflow execution to terminate.
	//
	// Domain is a required field
	Domain *string `locationName:"domain" min:"1" type:"string" required:"true"`

	// A descriptive reason for terminating the workflow execution.
	Reason *string `locationName:"reason" type:"string"`

	// The runId of the workflow execution to terminate.
	RunId *string `locationName:"runId" type:"string"`

	// The workflowId of the workflow execution to terminate.
	//
	// WorkflowId is a required field
	WorkflowId *string `locationName:"workflowId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s TerminateWorkflowExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TerminateWorkflowExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TerminateWorkflowExecutionInput"}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 1))
	}

	if s.WorkflowId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkflowId"))
	}
	if s.WorkflowId != nil && len(*s.WorkflowId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WorkflowId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TerminateWorkflowExecutionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TerminateWorkflowExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

const opTerminateWorkflowExecution = "TerminateWorkflowExecution"

// TerminateWorkflowExecutionRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Records a WorkflowExecutionTerminated event and forces closure of the workflow
// execution identified by the given domain, runId, and workflowId. The child
// policy, registered with the workflow type or specified when starting this
// execution, is applied to any open child workflow executions of this workflow
// execution.
//
// If the identified workflow execution was in progress, it is terminated immediately.
//
// If a runId isn't specified, then the WorkflowExecutionTerminated event is
// recorded in the history of the current open workflow with the matching workflowId
// in the domain.
//
// You should consider using RequestCancelWorkflowExecution action instead because
// it allows the workflow to gracefully close while TerminateWorkflowExecution
// doesn't.
//
// Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF resources
// as follows:
//
//    * Use a Resource element with the domain name to limit the action to only
//    specified domains.
//
//    * Use an Action element to allow or deny permission to call this action.
//
//    * You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using TerminateWorkflowExecutionRequest.
//    req := client.TerminateWorkflowExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) TerminateWorkflowExecutionRequest(input *TerminateWorkflowExecutionInput) TerminateWorkflowExecutionRequest {
	op := &aws.Operation{
		Name:       opTerminateWorkflowExecution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TerminateWorkflowExecutionInput{}
	}

	req := c.newRequest(op, input, &TerminateWorkflowExecutionOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return TerminateWorkflowExecutionRequest{Request: req, Input: input, Copy: c.TerminateWorkflowExecutionRequest}
}

// TerminateWorkflowExecutionRequest is the request type for the
// TerminateWorkflowExecution API operation.
type TerminateWorkflowExecutionRequest struct {
	*aws.Request
	Input *TerminateWorkflowExecutionInput
	Copy  func(*TerminateWorkflowExecutionInput) TerminateWorkflowExecutionRequest
}

// Send marshals and sends the TerminateWorkflowExecution API request.
func (r TerminateWorkflowExecutionRequest) Send(ctx context.Context) (*TerminateWorkflowExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TerminateWorkflowExecutionResponse{
		TerminateWorkflowExecutionOutput: r.Request.Data.(*TerminateWorkflowExecutionOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TerminateWorkflowExecutionResponse is the response type for the
// TerminateWorkflowExecution API operation.
type TerminateWorkflowExecutionResponse struct {
	*TerminateWorkflowExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TerminateWorkflowExecution request.
func (r *TerminateWorkflowExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
