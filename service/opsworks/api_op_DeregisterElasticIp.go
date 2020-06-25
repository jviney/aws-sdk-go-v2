// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeregisterElasticIpInput struct {
	_ struct{} `type:"structure"`

	// The Elastic IP address.
	//
	// ElasticIp is a required field
	ElasticIp *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeregisterElasticIpInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterElasticIpInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeregisterElasticIpInput"}

	if s.ElasticIp == nil {
		invalidParams.Add(aws.NewErrParamRequired("ElasticIp"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeregisterElasticIpOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeregisterElasticIpOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeregisterElasticIp = "DeregisterElasticIp"

// DeregisterElasticIpRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Deregisters a specified Elastic IP address. The address can then be registered
// by another stack. For more information, see Resource Management (https://docs.aws.amazon.com/opsworks/latest/userguide/resources.html).
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using DeregisterElasticIpRequest.
//    req := client.DeregisterElasticIpRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DeregisterElasticIp
func (c *Client) DeregisterElasticIpRequest(input *DeregisterElasticIpInput) DeregisterElasticIpRequest {
	op := &aws.Operation{
		Name:       opDeregisterElasticIp,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeregisterElasticIpInput{}
	}

	req := c.newRequest(op, input, &DeregisterElasticIpOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeregisterElasticIpRequest{Request: req, Input: input, Copy: c.DeregisterElasticIpRequest}
}

// DeregisterElasticIpRequest is the request type for the
// DeregisterElasticIp API operation.
type DeregisterElasticIpRequest struct {
	*aws.Request
	Input *DeregisterElasticIpInput
	Copy  func(*DeregisterElasticIpInput) DeregisterElasticIpRequest
}

// Send marshals and sends the DeregisterElasticIp API request.
func (r DeregisterElasticIpRequest) Send(ctx context.Context) (*DeregisterElasticIpResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterElasticIpResponse{
		DeregisterElasticIpOutput: r.Request.Data.(*DeregisterElasticIpOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterElasticIpResponse is the response type for the
// DeregisterElasticIp API operation.
type DeregisterElasticIpResponse struct {
	*DeregisterElasticIpOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterElasticIp request.
func (r *DeregisterElasticIpResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
