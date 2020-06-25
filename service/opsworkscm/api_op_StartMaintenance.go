// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworkscm

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type StartMaintenanceInput struct {
	_ struct{} `type:"structure"`

	// Engine attributes that are specific to the server on which you want to run
	// maintenance.
	//
	// Attributes accepted in a StartMaintenance request for Chef
	//
	//    * CHEF_MAJOR_UPGRADE: If a Chef Automate server is eligible for upgrade
	//    to Chef Automate 2, add this engine attribute to a StartMaintenance request
	//    and set the value to true to upgrade the server to Chef Automate 2. For
	//    more information, see Upgrade an AWS OpsWorks for Chef Automate Server
	//    to Chef Automate 2 (https://docs.aws.amazon.com/opsworks/latest/userguide/opscm-a2upgrade.html).
	EngineAttributes []EngineAttribute `type:"list"`

	// The name of the server on which to run maintenance.
	//
	// ServerName is a required field
	ServerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartMaintenanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartMaintenanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartMaintenanceInput"}

	if s.ServerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerName"))
	}
	if s.ServerName != nil && len(*s.ServerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartMaintenanceOutput struct {
	_ struct{} `type:"structure"`

	// Contains the response to a StartMaintenance request.
	Server *Server `type:"structure"`
}

// String returns the string representation
func (s StartMaintenanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartMaintenance = "StartMaintenance"

// StartMaintenanceRequest returns a request value for making API operation for
// AWS OpsWorks CM.
//
// Manually starts server maintenance. This command can be useful if an earlier
// maintenance attempt failed, and the underlying cause of maintenance failure
// has been resolved. The server is in an UNDER_MAINTENANCE state while maintenance
// is in progress.
//
// Maintenance can only be started on servers in HEALTHY and UNHEALTHY states.
// Otherwise, an InvalidStateException is thrown. A ResourceNotFoundException
// is thrown when the server does not exist. A ValidationException is raised
// when parameters of the request are not valid.
//
//    // Example sending a request using StartMaintenanceRequest.
//    req := client.StartMaintenanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/StartMaintenance
func (c *Client) StartMaintenanceRequest(input *StartMaintenanceInput) StartMaintenanceRequest {
	op := &aws.Operation{
		Name:       opStartMaintenance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartMaintenanceInput{}
	}

	req := c.newRequest(op, input, &StartMaintenanceOutput{})

	return StartMaintenanceRequest{Request: req, Input: input, Copy: c.StartMaintenanceRequest}
}

// StartMaintenanceRequest is the request type for the
// StartMaintenance API operation.
type StartMaintenanceRequest struct {
	*aws.Request
	Input *StartMaintenanceInput
	Copy  func(*StartMaintenanceInput) StartMaintenanceRequest
}

// Send marshals and sends the StartMaintenance API request.
func (r StartMaintenanceRequest) Send(ctx context.Context) (*StartMaintenanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartMaintenanceResponse{
		StartMaintenanceOutput: r.Request.Data.(*StartMaintenanceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartMaintenanceResponse is the response type for the
// StartMaintenance API operation.
type StartMaintenanceResponse struct {
	*StartMaintenanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartMaintenance request.
func (r *StartMaintenanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
