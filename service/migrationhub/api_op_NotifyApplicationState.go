// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type NotifyApplicationStateInput struct {
	_ struct{} `type:"structure"`

	// The configurationId in Application Discovery Service that uniquely identifies
	// the grouped application.
	//
	// ApplicationId is a required field
	ApplicationId *string `min:"1" type:"string" required:"true"`

	// Optional boolean flag to indicate whether any effect should take place. Used
	// to test if the caller has permission to make the call.
	DryRun *bool `type:"boolean"`

	// Status of the application - Not Started, In-Progress, Complete.
	//
	// Status is a required field
	Status ApplicationStatus `type:"string" required:"true" enum:"true"`

	// The timestamp when the application state changed.
	UpdateDateTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s NotifyApplicationStateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *NotifyApplicationStateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "NotifyApplicationStateInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}
	if s.ApplicationId != nil && len(*s.ApplicationId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationId", 1))
	}
	if len(s.Status) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Status"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type NotifyApplicationStateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s NotifyApplicationStateOutput) String() string {
	return awsutil.Prettify(s)
}

const opNotifyApplicationState = "NotifyApplicationState"

// NotifyApplicationStateRequest returns a request value for making API operation for
// AWS Migration Hub.
//
// Sets the migration state of an application. For a given application identified
// by the value passed to ApplicationId, its status is set or updated by passing
// one of three values to Status: NOT_STARTED | IN_PROGRESS | COMPLETED.
//
//    // Example sending a request using NotifyApplicationStateRequest.
//    req := client.NotifyApplicationStateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/NotifyApplicationState
func (c *Client) NotifyApplicationStateRequest(input *NotifyApplicationStateInput) NotifyApplicationStateRequest {
	op := &aws.Operation{
		Name:       opNotifyApplicationState,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &NotifyApplicationStateInput{}
	}

	req := c.newRequest(op, input, &NotifyApplicationStateOutput{})

	return NotifyApplicationStateRequest{Request: req, Input: input, Copy: c.NotifyApplicationStateRequest}
}

// NotifyApplicationStateRequest is the request type for the
// NotifyApplicationState API operation.
type NotifyApplicationStateRequest struct {
	*aws.Request
	Input *NotifyApplicationStateInput
	Copy  func(*NotifyApplicationStateInput) NotifyApplicationStateRequest
}

// Send marshals and sends the NotifyApplicationState API request.
func (r NotifyApplicationStateRequest) Send(ctx context.Context) (*NotifyApplicationStateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NotifyApplicationStateResponse{
		NotifyApplicationStateOutput: r.Request.Data.(*NotifyApplicationStateOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NotifyApplicationStateResponse is the response type for the
// NotifyApplicationState API operation.
type NotifyApplicationStateResponse struct {
	*NotifyApplicationStateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NotifyApplicationState request.
func (r *NotifyApplicationStateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
