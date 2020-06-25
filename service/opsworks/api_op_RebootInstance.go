// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type RebootInstanceInput struct {
	_ struct{} `type:"structure"`

	// The instance ID.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RebootInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RebootInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RebootInstanceInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RebootInstanceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RebootInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opRebootInstance = "RebootInstance"

// RebootInstanceRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Reboots a specified instance. For more information, see Starting, Stopping,
// and Rebooting Instances (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-starting.html).
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using RebootInstanceRequest.
//    req := client.RebootInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/RebootInstance
func (c *Client) RebootInstanceRequest(input *RebootInstanceInput) RebootInstanceRequest {
	op := &aws.Operation{
		Name:       opRebootInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RebootInstanceInput{}
	}

	req := c.newRequest(op, input, &RebootInstanceOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return RebootInstanceRequest{Request: req, Input: input, Copy: c.RebootInstanceRequest}
}

// RebootInstanceRequest is the request type for the
// RebootInstance API operation.
type RebootInstanceRequest struct {
	*aws.Request
	Input *RebootInstanceInput
	Copy  func(*RebootInstanceInput) RebootInstanceRequest
}

// Send marshals and sends the RebootInstance API request.
func (r RebootInstanceRequest) Send(ctx context.Context) (*RebootInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RebootInstanceResponse{
		RebootInstanceOutput: r.Request.Data.(*RebootInstanceOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RebootInstanceResponse is the response type for the
// RebootInstance API operation.
type RebootInstanceResponse struct {
	*RebootInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RebootInstance request.
func (r *RebootInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
