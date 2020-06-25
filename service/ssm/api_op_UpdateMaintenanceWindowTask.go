// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type UpdateMaintenanceWindowTaskInput struct {
	_ struct{} `type:"structure"`

	// The new task description to specify.
	Description *string `min:"1" type:"string" sensitive:"true"`

	// The new logging location in Amazon S3 to specify.
	//
	// LoggingInfo has been deprecated. To specify an S3 bucket to contain logs,
	// instead use the OutputS3BucketName and OutputS3KeyPrefix options in the TaskInvocationParameters
	// structure. For information about how Systems Manager handles these options
	// for the supported maintenance window task types, see MaintenanceWindowTaskInvocationParameters.
	LoggingInfo *LoggingInfo `type:"structure"`

	// The new MaxConcurrency value you want to specify. MaxConcurrency is the number
	// of targets that are allowed to run this task in parallel.
	MaxConcurrency *string `min:"1" type:"string"`

	// The new MaxErrors value to specify. MaxErrors is the maximum number of errors
	// that are allowed before the task stops being scheduled.
	MaxErrors *string `min:"1" type:"string"`

	// The new task name to specify.
	Name *string `min:"3" type:"string"`

	// The new task priority to specify. The lower the number, the higher the priority.
	// Tasks that have the same priority are scheduled in parallel.
	Priority *int64 `type:"integer"`

	// If True, then all fields that are required by the RegisterTaskWithMaintenanceWndow
	// action are also required for this API request. Optional fields that are not
	// specified are set to null.
	Replace *bool `type:"boolean"`

	// The ARN of the IAM service role for Systems Manager to assume when running
	// a maintenance window task. If you do not specify a service role ARN, Systems
	// Manager uses your account's service-linked role. If no service-linked role
	// for Systems Manager exists in your account, it is created when you run RegisterTaskWithMaintenanceWindow.
	//
	// For more information, see the following topics in the in the AWS Systems
	// Manager User Guide:
	//
	//    * Using service-linked roles for Systems Manager (https://docs.aws.amazon.com/systems-manager/latest/userguide/using-service-linked-roles.html#slr-permissions)
	//
	//    * Should I use a service-linked role or a custom service role to run maintenance
	//    window tasks? (https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-maintenance-permissions.html#maintenance-window-tasks-service-role)
	ServiceRoleArn *string `type:"string"`

	// The targets (either instances or tags) to modify. Instances are specified
	// using Key=instanceids,Values=instanceID_1,instanceID_2. Tags are specified
	// using Key=tag_name,Values=tag_value.
	Targets []Target `type:"list"`

	// The task ARN to modify.
	TaskArn *string `min:"1" type:"string"`

	// The parameters that the task should use during execution. Populate only the
	// fields that match the task type. All other fields should be empty.
	TaskInvocationParameters *MaintenanceWindowTaskInvocationParameters `type:"structure"`

	// The parameters to modify.
	//
	// TaskParameters has been deprecated. To specify parameters to pass to a task
	// when it runs, instead use the Parameters option in the TaskInvocationParameters
	// structure. For information about how Systems Manager handles these options
	// for the supported maintenance window task types, see MaintenanceWindowTaskInvocationParameters.
	//
	// The map has the following format:
	//
	// Key: string, between 1 and 255 characters
	//
	// Value: an array of strings, each string is between 1 and 255 characters
	TaskParameters map[string]MaintenanceWindowTaskParameterValueExpression `type:"map" sensitive:"true"`

	// The maintenance window ID that contains the task to modify.
	//
	// WindowId is a required field
	WindowId *string `min:"20" type:"string" required:"true"`

	// The task ID to modify.
	//
	// WindowTaskId is a required field
	WindowTaskId *string `min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateMaintenanceWindowTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateMaintenanceWindowTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateMaintenanceWindowTaskInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if s.MaxConcurrency != nil && len(*s.MaxConcurrency) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MaxConcurrency", 1))
	}
	if s.MaxErrors != nil && len(*s.MaxErrors) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MaxErrors", 1))
	}
	if s.Name != nil && len(*s.Name) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 3))
	}
	if s.TaskArn != nil && len(*s.TaskArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TaskArn", 1))
	}

	if s.WindowId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WindowId"))
	}
	if s.WindowId != nil && len(*s.WindowId) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("WindowId", 20))
	}

	if s.WindowTaskId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WindowTaskId"))
	}
	if s.WindowTaskId != nil && len(*s.WindowTaskId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("WindowTaskId", 36))
	}
	if s.LoggingInfo != nil {
		if err := s.LoggingInfo.Validate(); err != nil {
			invalidParams.AddNested("LoggingInfo", err.(aws.ErrInvalidParams))
		}
	}
	if s.Targets != nil {
		for i, v := range s.Targets {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Targets", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.TaskInvocationParameters != nil {
		if err := s.TaskInvocationParameters.Validate(); err != nil {
			invalidParams.AddNested("TaskInvocationParameters", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateMaintenanceWindowTaskOutput struct {
	_ struct{} `type:"structure"`

	// The updated task description.
	Description *string `min:"1" type:"string" sensitive:"true"`

	// The updated logging information in Amazon S3.
	//
	// LoggingInfo has been deprecated. To specify an S3 bucket to contain logs,
	// instead use the OutputS3BucketName and OutputS3KeyPrefix options in the TaskInvocationParameters
	// structure. For information about how Systems Manager handles these options
	// for the supported maintenance window task types, see MaintenanceWindowTaskInvocationParameters.
	LoggingInfo *LoggingInfo `type:"structure"`

	// The updated MaxConcurrency value.
	MaxConcurrency *string `min:"1" type:"string"`

	// The updated MaxErrors value.
	MaxErrors *string `min:"1" type:"string"`

	// The updated task name.
	Name *string `min:"3" type:"string"`

	// The updated priority value.
	Priority *int64 `type:"integer"`

	// The ARN of the IAM service role to use to publish Amazon Simple Notification
	// Service (Amazon SNS) notifications for maintenance window Run Command tasks.
	ServiceRoleArn *string `type:"string"`

	// The updated target values.
	Targets []Target `type:"list"`

	// The updated task ARN value.
	TaskArn *string `min:"1" type:"string"`

	// The updated parameter values.
	TaskInvocationParameters *MaintenanceWindowTaskInvocationParameters `type:"structure"`

	// The updated parameter values.
	//
	// TaskParameters has been deprecated. To specify parameters to pass to a task
	// when it runs, instead use the Parameters option in the TaskInvocationParameters
	// structure. For information about how Systems Manager handles these options
	// for the supported maintenance window task types, see MaintenanceWindowTaskInvocationParameters.
	TaskParameters map[string]MaintenanceWindowTaskParameterValueExpression `type:"map" sensitive:"true"`

	// The ID of the maintenance window that was updated.
	WindowId *string `min:"20" type:"string"`

	// The task ID of the maintenance window that was updated.
	WindowTaskId *string `min:"36" type:"string"`
}

// String returns the string representation
func (s UpdateMaintenanceWindowTaskOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateMaintenanceWindowTask = "UpdateMaintenanceWindowTask"

// UpdateMaintenanceWindowTaskRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Modifies a task assigned to a maintenance window. You can't change the task
// type, but you can change the following values:
//
//    * TaskARN. For example, you can change a RUN_COMMAND task from AWS-RunPowerShellScript
//    to AWS-RunShellScript.
//
//    * ServiceRoleArn
//
//    * TaskInvocationParameters
//
//    * Priority
//
//    * MaxConcurrency
//
//    * MaxErrors
//
// If a parameter is null, then the corresponding field is not modified. Also,
// if you set Replace to true, then all fields required by the RegisterTaskWithMaintenanceWindow
// action are required for this request. Optional fields that aren't specified
// are set to null.
//
//    // Example sending a request using UpdateMaintenanceWindowTaskRequest.
//    req := client.UpdateMaintenanceWindowTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/UpdateMaintenanceWindowTask
func (c *Client) UpdateMaintenanceWindowTaskRequest(input *UpdateMaintenanceWindowTaskInput) UpdateMaintenanceWindowTaskRequest {
	op := &aws.Operation{
		Name:       opUpdateMaintenanceWindowTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateMaintenanceWindowTaskInput{}
	}

	req := c.newRequest(op, input, &UpdateMaintenanceWindowTaskOutput{})

	return UpdateMaintenanceWindowTaskRequest{Request: req, Input: input, Copy: c.UpdateMaintenanceWindowTaskRequest}
}

// UpdateMaintenanceWindowTaskRequest is the request type for the
// UpdateMaintenanceWindowTask API operation.
type UpdateMaintenanceWindowTaskRequest struct {
	*aws.Request
	Input *UpdateMaintenanceWindowTaskInput
	Copy  func(*UpdateMaintenanceWindowTaskInput) UpdateMaintenanceWindowTaskRequest
}

// Send marshals and sends the UpdateMaintenanceWindowTask API request.
func (r UpdateMaintenanceWindowTaskRequest) Send(ctx context.Context) (*UpdateMaintenanceWindowTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateMaintenanceWindowTaskResponse{
		UpdateMaintenanceWindowTaskOutput: r.Request.Data.(*UpdateMaintenanceWindowTaskOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateMaintenanceWindowTaskResponse is the response type for the
// UpdateMaintenanceWindowTask API operation.
type UpdateMaintenanceWindowTaskResponse struct {
	*UpdateMaintenanceWindowTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateMaintenanceWindowTask request.
func (r *UpdateMaintenanceWindowTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
