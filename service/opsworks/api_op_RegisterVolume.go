// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type RegisterVolumeInput struct {
	_ struct{} `type:"structure"`

	// The Amazon EBS volume ID.
	Ec2VolumeId *string `type:"string"`

	// The stack ID.
	//
	// StackId is a required field
	StackId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RegisterVolumeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterVolumeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterVolumeInput"}

	if s.StackId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a RegisterVolume request.
type RegisterVolumeOutput struct {
	_ struct{} `type:"structure"`

	// The volume ID.
	VolumeId *string `type:"string"`
}

// String returns the string representation
func (s RegisterVolumeOutput) String() string {
	return awsutil.Prettify(s)
}

const opRegisterVolume = "RegisterVolume"

// RegisterVolumeRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Registers an Amazon EBS volume with a specified stack. A volume can be registered
// with only one stack at a time. If the volume is already registered, you must
// first deregister it by calling DeregisterVolume. For more information, see
// Resource Management (https://docs.aws.amazon.com/opsworks/latest/userguide/resources.html).
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using RegisterVolumeRequest.
//    req := client.RegisterVolumeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/RegisterVolume
func (c *Client) RegisterVolumeRequest(input *RegisterVolumeInput) RegisterVolumeRequest {
	op := &aws.Operation{
		Name:       opRegisterVolume,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterVolumeInput{}
	}

	req := c.newRequest(op, input, &RegisterVolumeOutput{})

	return RegisterVolumeRequest{Request: req, Input: input, Copy: c.RegisterVolumeRequest}
}

// RegisterVolumeRequest is the request type for the
// RegisterVolume API operation.
type RegisterVolumeRequest struct {
	*aws.Request
	Input *RegisterVolumeInput
	Copy  func(*RegisterVolumeInput) RegisterVolumeRequest
}

// Send marshals and sends the RegisterVolume API request.
func (r RegisterVolumeRequest) Send(ctx context.Context) (*RegisterVolumeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterVolumeResponse{
		RegisterVolumeOutput: r.Request.Data.(*RegisterVolumeOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterVolumeResponse is the response type for the
// RegisterVolume API operation.
type RegisterVolumeResponse struct {
	*RegisterVolumeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterVolume request.
func (r *RegisterVolumeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
