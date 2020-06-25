// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworkscm

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type RestoreServerInput struct {
	_ struct{} `type:"structure"`

	// The ID of the backup that you want to use to restore a server.
	//
	// BackupId is a required field
	BackupId *string `type:"string" required:"true"`

	// The type of instance to restore. Valid values must be specified in the following
	// format: ^([cm][34]|t2).* For example, m5.large. Valid values are m5.large,
	// r5.xlarge, and r5.2xlarge. If you do not specify this parameter, RestoreServer
	// uses the instance type from the specified backup.
	InstanceType *string `type:"string"`

	// The name of the key pair to set on the new EC2 instance. This can be helpful
	// if the administrator no longer has the SSH key.
	KeyPair *string `type:"string"`

	// The name of the server that you want to restore.
	//
	// ServerName is a required field
	ServerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RestoreServerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreServerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestoreServerInput"}

	if s.BackupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupId"))
	}

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

type RestoreServerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RestoreServerOutput) String() string {
	return awsutil.Prettify(s)
}

const opRestoreServer = "RestoreServer"

// RestoreServerRequest returns a request value for making API operation for
// AWS OpsWorks CM.
//
// Restores a backup to a server that is in a CONNECTION_LOST, HEALTHY, RUNNING,
// UNHEALTHY, or TERMINATED state. When you run RestoreServer, the server's
// EC2 instance is deleted, and a new EC2 instance is configured. RestoreServer
// maintains the existing server endpoint, so configuration management of the
// server's client devices (nodes) should continue to work.
//
// Restoring from a backup is performed by creating a new EC2 instance. If restoration
// is successful, and the server is in a HEALTHY state, AWS OpsWorks CM switches
// traffic over to the new instance. After restoration is finished, the old
// EC2 instance is maintained in a Running or Stopped state, but is eventually
// terminated.
//
// This operation is asynchronous.
//
// An InvalidStateException is thrown when the server is not in a valid state.
// A ResourceNotFoundException is thrown when the server does not exist. A ValidationException
// is raised when parameters of the request are not valid.
//
//    // Example sending a request using RestoreServerRequest.
//    req := client.RestoreServerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/RestoreServer
func (c *Client) RestoreServerRequest(input *RestoreServerInput) RestoreServerRequest {
	op := &aws.Operation{
		Name:       opRestoreServer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestoreServerInput{}
	}

	req := c.newRequest(op, input, &RestoreServerOutput{})

	return RestoreServerRequest{Request: req, Input: input, Copy: c.RestoreServerRequest}
}

// RestoreServerRequest is the request type for the
// RestoreServer API operation.
type RestoreServerRequest struct {
	*aws.Request
	Input *RestoreServerInput
	Copy  func(*RestoreServerInput) RestoreServerRequest
}

// Send marshals and sends the RestoreServer API request.
func (r RestoreServerRequest) Send(ctx context.Context) (*RestoreServerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreServerResponse{
		RestoreServerOutput: r.Request.Data.(*RestoreServerOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreServerResponse is the response type for the
// RestoreServer API operation.
type RestoreServerResponse struct {
	*RestoreServerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreServer request.
func (r *RestoreServerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
