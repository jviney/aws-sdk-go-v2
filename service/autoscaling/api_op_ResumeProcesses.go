// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/query"
)

type ResumeProcessesInput struct {
	_ struct{} `type:"structure"`

	// The name of the Auto Scaling group.
	//
	// AutoScalingGroupName is a required field
	AutoScalingGroupName *string `min:"1" type:"string" required:"true"`

	// One or more of the following processes. If you omit this parameter, all processes
	// are specified.
	//
	//    * Launch
	//
	//    * Terminate
	//
	//    * HealthCheck
	//
	//    * ReplaceUnhealthy
	//
	//    * AZRebalance
	//
	//    * AlarmNotification
	//
	//    * ScheduledActions
	//
	//    * AddToLoadBalancer
	ScalingProcesses []string `type:"list"`
}

// String returns the string representation
func (s ResumeProcessesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResumeProcessesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResumeProcessesInput"}

	if s.AutoScalingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutoScalingGroupName"))
	}
	if s.AutoScalingGroupName != nil && len(*s.AutoScalingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoScalingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ResumeProcessesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ResumeProcessesOutput) String() string {
	return awsutil.Prettify(s)
}

const opResumeProcesses = "ResumeProcesses"

// ResumeProcessesRequest returns a request value for making API operation for
// Auto Scaling.
//
// Resumes the specified suspended automatic scaling processes, or all suspended
// process, for the specified Auto Scaling group.
//
// For more information, see Suspending and Resuming Scaling Processes (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-suspend-resume-processes.html)
// in the Amazon EC2 Auto Scaling User Guide.
//
//    // Example sending a request using ResumeProcessesRequest.
//    req := client.ResumeProcessesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/ResumeProcesses
func (c *Client) ResumeProcessesRequest(input *ResumeProcessesInput) ResumeProcessesRequest {
	op := &aws.Operation{
		Name:       opResumeProcesses,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResumeProcessesInput{}
	}

	req := c.newRequest(op, input, &ResumeProcessesOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return ResumeProcessesRequest{Request: req, Input: input, Copy: c.ResumeProcessesRequest}
}

// ResumeProcessesRequest is the request type for the
// ResumeProcesses API operation.
type ResumeProcessesRequest struct {
	*aws.Request
	Input *ResumeProcessesInput
	Copy  func(*ResumeProcessesInput) ResumeProcessesRequest
}

// Send marshals and sends the ResumeProcesses API request.
func (r ResumeProcessesRequest) Send(ctx context.Context) (*ResumeProcessesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResumeProcessesResponse{
		ResumeProcessesOutput: r.Request.Data.(*ResumeProcessesOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResumeProcessesResponse is the response type for the
// ResumeProcesses API operation.
type ResumeProcessesResponse struct {
	*ResumeProcessesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResumeProcesses request.
func (r *ResumeProcessesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
