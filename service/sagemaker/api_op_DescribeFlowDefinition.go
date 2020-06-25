// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeFlowDefinitionInput struct {
	_ struct{} `type:"structure"`

	// The name of the flow definition.
	//
	// FlowDefinitionName is a required field
	FlowDefinitionName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeFlowDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeFlowDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeFlowDefinitionInput"}

	if s.FlowDefinitionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FlowDefinitionName"))
	}
	if s.FlowDefinitionName != nil && len(*s.FlowDefinitionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FlowDefinitionName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeFlowDefinitionOutput struct {
	_ struct{} `type:"structure"`

	// The timestamp when the flow definition was created.
	//
	// CreationTime is a required field
	CreationTime *time.Time `type:"timestamp" required:"true"`

	FailureReason *string `type:"string"`

	// The Amazon Resource Name (ARN) of the flow defintion.
	//
	// FlowDefinitionArn is a required field
	FlowDefinitionArn *string `type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the flow definition.
	//
	// FlowDefinitionName is a required field
	FlowDefinitionName *string `min:"1" type:"string" required:"true"`

	// The status of the flow definition. Valid values are listed below.
	//
	// FlowDefinitionStatus is a required field
	FlowDefinitionStatus FlowDefinitionStatus `type:"string" required:"true" enum:"true"`

	// An object containing information about what triggers a human review workflow.
	HumanLoopActivationConfig *HumanLoopActivationConfig `type:"structure"`

	// An object containing information about who works on the task, the workforce
	// task price, and other task details.
	//
	// HumanLoopConfig is a required field
	HumanLoopConfig *HumanLoopConfig `type:"structure" required:"true"`

	// Container for configuring the source of human task requests. Used to specify
	// if Amazon Rekognition or Amazon Textract is used as an integration source.
	HumanLoopRequestSource *HumanLoopRequestSource `type:"structure"`

	// An object containing information about the output file.
	//
	// OutputConfig is a required field
	OutputConfig *FlowDefinitionOutputConfig `type:"structure" required:"true"`

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management
	// (IAM) execution role for the flow definition.
	//
	// RoleArn is a required field
	RoleArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeFlowDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeFlowDefinition = "DescribeFlowDefinition"

// DescribeFlowDefinitionRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Returns information about the specified flow definition.
//
//    // Example sending a request using DescribeFlowDefinitionRequest.
//    req := client.DescribeFlowDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeFlowDefinition
func (c *Client) DescribeFlowDefinitionRequest(input *DescribeFlowDefinitionInput) DescribeFlowDefinitionRequest {
	op := &aws.Operation{
		Name:       opDescribeFlowDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeFlowDefinitionInput{}
	}

	req := c.newRequest(op, input, &DescribeFlowDefinitionOutput{})

	return DescribeFlowDefinitionRequest{Request: req, Input: input, Copy: c.DescribeFlowDefinitionRequest}
}

// DescribeFlowDefinitionRequest is the request type for the
// DescribeFlowDefinition API operation.
type DescribeFlowDefinitionRequest struct {
	*aws.Request
	Input *DescribeFlowDefinitionInput
	Copy  func(*DescribeFlowDefinitionInput) DescribeFlowDefinitionRequest
}

// Send marshals and sends the DescribeFlowDefinition API request.
func (r DescribeFlowDefinitionRequest) Send(ctx context.Context) (*DescribeFlowDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFlowDefinitionResponse{
		DescribeFlowDefinitionOutput: r.Request.Data.(*DescribeFlowDefinitionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeFlowDefinitionResponse is the response type for the
// DescribeFlowDefinition API operation.
type DescribeFlowDefinitionResponse struct {
	*DescribeFlowDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFlowDefinition request.
func (r *DescribeFlowDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
