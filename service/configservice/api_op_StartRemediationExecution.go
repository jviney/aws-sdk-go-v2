// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type StartRemediationExecutionInput struct {
	_ struct{} `type:"structure"`

	// The list of names of AWS Config rules that you want to run remediation execution
	// for.
	//
	// ConfigRuleName is a required field
	ConfigRuleName *string `min:"1" type:"string" required:"true"`

	// A list of resource keys to be processed with the current request. Each element
	// in the list consists of the resource type and resource ID.
	//
	// ResourceKeys is a required field
	ResourceKeys []ResourceKey `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s StartRemediationExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartRemediationExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartRemediationExecutionInput"}

	if s.ConfigRuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigRuleName"))
	}
	if s.ConfigRuleName != nil && len(*s.ConfigRuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConfigRuleName", 1))
	}

	if s.ResourceKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceKeys"))
	}
	if s.ResourceKeys != nil && len(s.ResourceKeys) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceKeys", 1))
	}
	if s.ResourceKeys != nil {
		for i, v := range s.ResourceKeys {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ResourceKeys", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartRemediationExecutionOutput struct {
	_ struct{} `type:"structure"`

	// For resources that have failed to start execution, the API returns a resource
	// key object.
	FailedItems []ResourceKey `min:"1" type:"list"`

	// Returns a failure message. For example, the resource is already compliant.
	FailureMessage *string `type:"string"`
}

// String returns the string representation
func (s StartRemediationExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartRemediationExecution = "StartRemediationExecution"

// StartRemediationExecutionRequest returns a request value for making API operation for
// AWS Config.
//
// Runs an on-demand remediation for the specified AWS Config rules against
// the last known remediation configuration. It runs an execution against the
// current state of your resources. Remediation execution is asynchronous.
//
// You can specify up to 100 resource keys per request. An existing StartRemediationExecution
// call for the specified resource keys must complete before you can call the
// API again.
//
//    // Example sending a request using StartRemediationExecutionRequest.
//    req := client.StartRemediationExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/StartRemediationExecution
func (c *Client) StartRemediationExecutionRequest(input *StartRemediationExecutionInput) StartRemediationExecutionRequest {
	op := &aws.Operation{
		Name:       opStartRemediationExecution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartRemediationExecutionInput{}
	}

	req := c.newRequest(op, input, &StartRemediationExecutionOutput{})

	return StartRemediationExecutionRequest{Request: req, Input: input, Copy: c.StartRemediationExecutionRequest}
}

// StartRemediationExecutionRequest is the request type for the
// StartRemediationExecution API operation.
type StartRemediationExecutionRequest struct {
	*aws.Request
	Input *StartRemediationExecutionInput
	Copy  func(*StartRemediationExecutionInput) StartRemediationExecutionRequest
}

// Send marshals and sends the StartRemediationExecution API request.
func (r StartRemediationExecutionRequest) Send(ctx context.Context) (*StartRemediationExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartRemediationExecutionResponse{
		StartRemediationExecutionOutput: r.Request.Data.(*StartRemediationExecutionOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartRemediationExecutionResponse is the response type for the
// StartRemediationExecution API operation.
type StartRemediationExecutionResponse struct {
	*StartRemediationExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartRemediationExecution request.
func (r *StartRemediationExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
