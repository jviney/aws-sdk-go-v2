// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// The input for the ExecuteChangeSet action.
type ExecuteChangeSetInput struct {
	_ struct{} `type:"structure"`

	// The name or ARN of the change set that you want use to update the specified
	// stack.
	//
	// ChangeSetName is a required field
	ChangeSetName *string `min:"1" type:"string" required:"true"`

	// A unique identifier for this ExecuteChangeSet request. Specify this token
	// if you plan to retry requests so that AWS CloudFormation knows that you're
	// not attempting to execute a change set to update a stack with the same name.
	// You might retry ExecuteChangeSet requests to ensure that AWS CloudFormation
	// successfully received them.
	ClientRequestToken *string `min:"1" type:"string"`

	// If you specified the name of a change set, specify the stack name or ID (ARN)
	// that is associated with the change set you want to execute.
	StackName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ExecuteChangeSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ExecuteChangeSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ExecuteChangeSetInput"}

	if s.ChangeSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeSetName"))
	}
	if s.ChangeSetName != nil && len(*s.ChangeSetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeSetName", 1))
	}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}
	if s.StackName != nil && len(*s.StackName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StackName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output for the ExecuteChangeSet action.
type ExecuteChangeSetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ExecuteChangeSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opExecuteChangeSet = "ExecuteChangeSet"

// ExecuteChangeSetRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Updates a stack using the input information that was provided when the specified
// change set was created. After the call successfully completes, AWS CloudFormation
// starts updating the stack. Use the DescribeStacks action to view the status
// of the update.
//
// When you execute a change set, AWS CloudFormation deletes all other change
// sets associated with the stack because they aren't valid for the updated
// stack.
//
// If a stack policy is associated with the stack, AWS CloudFormation enforces
// the policy during the update. You can't specify a temporary stack policy
// that overrides the current policy.
//
//    // Example sending a request using ExecuteChangeSetRequest.
//    req := client.ExecuteChangeSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ExecuteChangeSet
func (c *Client) ExecuteChangeSetRequest(input *ExecuteChangeSetInput) ExecuteChangeSetRequest {
	op := &aws.Operation{
		Name:       opExecuteChangeSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ExecuteChangeSetInput{}
	}

	req := c.newRequest(op, input, &ExecuteChangeSetOutput{})

	return ExecuteChangeSetRequest{Request: req, Input: input, Copy: c.ExecuteChangeSetRequest}
}

// ExecuteChangeSetRequest is the request type for the
// ExecuteChangeSet API operation.
type ExecuteChangeSetRequest struct {
	*aws.Request
	Input *ExecuteChangeSetInput
	Copy  func(*ExecuteChangeSetInput) ExecuteChangeSetRequest
}

// Send marshals and sends the ExecuteChangeSet API request.
func (r ExecuteChangeSetRequest) Send(ctx context.Context) (*ExecuteChangeSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ExecuteChangeSetResponse{
		ExecuteChangeSetOutput: r.Request.Data.(*ExecuteChangeSetOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ExecuteChangeSetResponse is the response type for the
// ExecuteChangeSet API operation.
type ExecuteChangeSetResponse struct {
	*ExecuteChangeSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ExecuteChangeSet request.
func (r *ExecuteChangeSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
