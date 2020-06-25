// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeStackDriftDetectionStatusInput struct {
	_ struct{} `type:"structure"`

	// The ID of the drift detection results of this operation.
	//
	// AWS CloudFormation generates new results, with a new drift detection ID,
	// each time this operation is run. However, the number of drift results AWS
	// CloudFormation retains for any given stack, and for how long, may vary.
	//
	// StackDriftDetectionId is a required field
	StackDriftDetectionId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeStackDriftDetectionStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStackDriftDetectionStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStackDriftDetectionStatusInput"}

	if s.StackDriftDetectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackDriftDetectionId"))
	}
	if s.StackDriftDetectionId != nil && len(*s.StackDriftDetectionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StackDriftDetectionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeStackDriftDetectionStatusOutput struct {
	_ struct{} `type:"structure"`

	// The status of the stack drift detection operation.
	//
	//    * DETECTION_COMPLETE: The stack drift detection operation has successfully
	//    completed for all resources in the stack that support drift detection.
	//    (Resources that do not currently support stack detection remain unchecked.)
	//    If you specified logical resource IDs for AWS CloudFormation to use as
	//    a filter for the stack drift detection operation, only the resources with
	//    those logical IDs are checked for drift.
	//
	//    * DETECTION_FAILED: The stack drift detection operation has failed for
	//    at least one resource in the stack. Results will be available for resources
	//    on which AWS CloudFormation successfully completed drift detection.
	//
	//    * DETECTION_IN_PROGRESS: The stack drift detection operation is currently
	//    in progress.
	//
	// DetectionStatus is a required field
	DetectionStatus StackDriftDetectionStatus `type:"string" required:"true" enum:"true"`

	// The reason the stack drift detection operation has its current status.
	DetectionStatusReason *string `type:"string"`

	// Total number of stack resources that have drifted. This is NULL until the
	// drift detection operation reaches a status of DETECTION_COMPLETE. This value
	// will be 0 for stacks whose drift status is IN_SYNC.
	DriftedStackResourceCount *int64 `type:"integer"`

	// The ID of the drift detection results of this operation.
	//
	// AWS CloudFormation generates new results, with a new drift detection ID,
	// each time this operation is run. However, the number of reports AWS CloudFormation
	// retains for any given stack, and for how long, may vary.
	//
	// StackDriftDetectionId is a required field
	StackDriftDetectionId *string `min:"1" type:"string" required:"true"`

	// Status of the stack's actual configuration compared to its expected configuration.
	//
	//    * DRIFTED: The stack differs from its expected template configuration.
	//    A stack is considered to have drifted if one or more of its resources
	//    have drifted.
	//
	//    * NOT_CHECKED: AWS CloudFormation has not checked if the stack differs
	//    from its expected template configuration.
	//
	//    * IN_SYNC: The stack's actual configuration matches its expected template
	//    configuration.
	//
	//    * UNKNOWN: This value is reserved for future use.
	StackDriftStatus StackDriftStatus `type:"string" enum:"true"`

	// The ID of the stack.
	//
	// StackId is a required field
	StackId *string `type:"string" required:"true"`

	// Time at which the stack drift detection operation was initiated.
	//
	// Timestamp is a required field
	Timestamp *time.Time `type:"timestamp" required:"true"`
}

// String returns the string representation
func (s DescribeStackDriftDetectionStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeStackDriftDetectionStatus = "DescribeStackDriftDetectionStatus"

// DescribeStackDriftDetectionStatusRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns information about a stack drift detection operation. A stack drift
// detection operation detects whether a stack's actual configuration differs,
// or has drifted, from it's expected configuration, as defined in the stack
// template and any values specified as template parameters. A stack is considered
// to have drifted if one or more of its resources have drifted. For more information
// on stack and resource drift, see Detecting Unregulated Configuration Changes
// to Stacks and Resources (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift.html).
//
// Use DetectStackDrift to initiate a stack drift detection operation. DetectStackDrift
// returns a StackDriftDetectionId you can use to monitor the progress of the
// operation using DescribeStackDriftDetectionStatus. Once the drift detection
// operation has completed, use DescribeStackResourceDrifts to return drift
// information about the stack and its resources.
//
//    // Example sending a request using DescribeStackDriftDetectionStatusRequest.
//    req := client.DescribeStackDriftDetectionStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeStackDriftDetectionStatus
func (c *Client) DescribeStackDriftDetectionStatusRequest(input *DescribeStackDriftDetectionStatusInput) DescribeStackDriftDetectionStatusRequest {
	op := &aws.Operation{
		Name:       opDescribeStackDriftDetectionStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeStackDriftDetectionStatusInput{}
	}

	req := c.newRequest(op, input, &DescribeStackDriftDetectionStatusOutput{})

	return DescribeStackDriftDetectionStatusRequest{Request: req, Input: input, Copy: c.DescribeStackDriftDetectionStatusRequest}
}

// DescribeStackDriftDetectionStatusRequest is the request type for the
// DescribeStackDriftDetectionStatus API operation.
type DescribeStackDriftDetectionStatusRequest struct {
	*aws.Request
	Input *DescribeStackDriftDetectionStatusInput
	Copy  func(*DescribeStackDriftDetectionStatusInput) DescribeStackDriftDetectionStatusRequest
}

// Send marshals and sends the DescribeStackDriftDetectionStatus API request.
func (r DescribeStackDriftDetectionStatusRequest) Send(ctx context.Context) (*DescribeStackDriftDetectionStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStackDriftDetectionStatusResponse{
		DescribeStackDriftDetectionStatusOutput: r.Request.Data.(*DescribeStackDriftDetectionStatusOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeStackDriftDetectionStatusResponse is the response type for the
// DescribeStackDriftDetectionStatus API operation.
type DescribeStackDriftDetectionStatusResponse struct {
	*DescribeStackDriftDetectionStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStackDriftDetectionStatus request.
func (r *DescribeStackDriftDetectionStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
