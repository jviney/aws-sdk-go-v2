// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ModifyScheduledActionInput struct {
	_ struct{} `type:"structure"`

	// A modified enable flag of the scheduled action. If true, the scheduled action
	// is active. If false, the scheduled action is disabled.
	Enable *bool `type:"boolean"`

	// A modified end time of the scheduled action. For more information about this
	// parameter, see ScheduledAction.
	EndTime *time.Time `type:"timestamp"`

	// A different IAM role to assume to run the target action. For more information
	// about this parameter, see ScheduledAction.
	IamRole *string `type:"string"`

	// A modified schedule in either at( ) or cron( ) format. For more information
	// about this parameter, see ScheduledAction.
	Schedule *string `type:"string"`

	// A modified description of the scheduled action.
	ScheduledActionDescription *string `type:"string"`

	// The name of the scheduled action to modify.
	//
	// ScheduledActionName is a required field
	ScheduledActionName *string `type:"string" required:"true"`

	// A modified start time of the scheduled action. For more information about
	// this parameter, see ScheduledAction.
	StartTime *time.Time `type:"timestamp"`

	// A modified JSON format of the scheduled action. For more information about
	// this parameter, see ScheduledAction.
	TargetAction *ScheduledActionType `type:"structure"`
}

// String returns the string representation
func (s ModifyScheduledActionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyScheduledActionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyScheduledActionInput"}

	if s.ScheduledActionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScheduledActionName"))
	}
	if s.TargetAction != nil {
		if err := s.TargetAction.Validate(); err != nil {
			invalidParams.AddNested("TargetAction", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes a scheduled action. You can use a scheduled action to trigger some
// Amazon Redshift API operations on a schedule. For information about which
// API operations can be scheduled, see ScheduledActionType.
type ModifyScheduledActionOutput struct {
	_ struct{} `type:"structure"`

	// The end time in UTC when the schedule is no longer active. After this time,
	// the scheduled action does not trigger.
	EndTime *time.Time `type:"timestamp"`

	// The IAM role to assume to run the scheduled action. This IAM role must have
	// permission to run the Amazon Redshift API operation in the scheduled action.
	// This IAM role must allow the Amazon Redshift scheduler (Principal scheduler.redshift.amazonaws.com)
	// to assume permissions on your behalf. For more information about the IAM
	// role to use with the Amazon Redshift scheduler, see Using Identity-Based
	// Policies for Amazon Redshift (https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-identity-based.html)
	// in the Amazon Redshift Cluster Management Guide.
	IamRole *string `type:"string"`

	// List of times when the scheduled action will run.
	NextInvocations []time.Time `locationNameList:"ScheduledActionTime" type:"list"`

	// The schedule for a one-time (at format) or recurring (cron format) scheduled
	// action. Schedule invocations must be separated by at least one hour.
	//
	// Format of at expressions is "at(yyyy-mm-ddThh:mm:ss)". For example, "at(2016-03-04T17:27:00)".
	//
	// Format of cron expressions is "cron(Minutes Hours Day-of-month Month Day-of-week
	// Year)". For example, "cron(0 10 ? * MON *)". For more information, see Cron
	// Expressions (https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions)
	// in the Amazon CloudWatch Events User Guide.
	Schedule *string `type:"string"`

	// The description of the scheduled action.
	ScheduledActionDescription *string `type:"string"`

	// The name of the scheduled action.
	ScheduledActionName *string `type:"string"`

	// The start time in UTC when the schedule is active. Before this time, the
	// scheduled action does not trigger.
	StartTime *time.Time `type:"timestamp"`

	// The state of the scheduled action. For example, DISABLED.
	State ScheduledActionState `type:"string" enum:"true"`

	// A JSON format string of the Amazon Redshift API operation with input parameters.
	//
	// "{\"ResizeCluster\":{\"NodeType\":\"ds2.8xlarge\",\"ClusterIdentifier\":\"my-test-cluster\",\"NumberOfNodes\":3}}".
	TargetAction *ScheduledActionType `type:"structure"`
}

// String returns the string representation
func (s ModifyScheduledActionOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyScheduledAction = "ModifyScheduledAction"

// ModifyScheduledActionRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Modifies a scheduled action.
//
//    // Example sending a request using ModifyScheduledActionRequest.
//    req := client.ModifyScheduledActionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/ModifyScheduledAction
func (c *Client) ModifyScheduledActionRequest(input *ModifyScheduledActionInput) ModifyScheduledActionRequest {
	op := &aws.Operation{
		Name:       opModifyScheduledAction,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyScheduledActionInput{}
	}

	req := c.newRequest(op, input, &ModifyScheduledActionOutput{})

	return ModifyScheduledActionRequest{Request: req, Input: input, Copy: c.ModifyScheduledActionRequest}
}

// ModifyScheduledActionRequest is the request type for the
// ModifyScheduledAction API operation.
type ModifyScheduledActionRequest struct {
	*aws.Request
	Input *ModifyScheduledActionInput
	Copy  func(*ModifyScheduledActionInput) ModifyScheduledActionRequest
}

// Send marshals and sends the ModifyScheduledAction API request.
func (r ModifyScheduledActionRequest) Send(ctx context.Context) (*ModifyScheduledActionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyScheduledActionResponse{
		ModifyScheduledActionOutput: r.Request.Data.(*ModifyScheduledActionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyScheduledActionResponse is the response type for the
// ModifyScheduledAction API operation.
type ModifyScheduledActionResponse struct {
	*ModifyScheduledActionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyScheduledAction request.
func (r *ModifyScheduledActionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
