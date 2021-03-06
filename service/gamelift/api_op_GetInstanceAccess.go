// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type GetInstanceAccessInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for a fleet that contains the instance you want access
	// to. You can use either the fleet ID or ARN value. The fleet can be in any
	// of the following statuses: ACTIVATING, ACTIVE, or ERROR. Fleets with an ERROR
	// status may be accessible for a short time before they are deleted.
	//
	// FleetId is a required field
	FleetId *string `type:"string" required:"true"`

	// A unique identifier for an instance you want to get access to. You can access
	// an instance in any status.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetInstanceAccessInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetInstanceAccessInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetInstanceAccessInput"}

	if s.FleetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetId"))
	}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type GetInstanceAccessOutput struct {
	_ struct{} `type:"structure"`

	// The connection information for a fleet instance, including IP address and
	// access credentials.
	InstanceAccess *InstanceAccess `type:"structure"`
}

// String returns the string representation
func (s GetInstanceAccessOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetInstanceAccess = "GetInstanceAccess"

// GetInstanceAccessRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Requests remote access to a fleet instance. Remote access is useful for debugging,
// gathering benchmarking data, or observing activity in real time.
//
// To remotely access an instance, you need credentials that match the operating
// system of the instance. For a Windows instance, Amazon GameLift returns a
// user name and password as strings for use with a Windows Remote Desktop client.
// For a Linux instance, Amazon GameLift returns a user name and RSA private
// key, also as strings, for use with an SSH client. The private key must be
// saved in the proper format to a .pem file before using. If you're making
// this request using the AWS CLI, saving the secret can be handled as part
// of the GetInstanceAccess request, as shown in one of the examples for this
// action.
//
// To request access to a specific instance, specify the IDs of both the instance
// and the fleet it belongs to. You can retrieve a fleet's instance IDs by calling
// DescribeInstances. If successful, an InstanceAccess object is returned that
// contains the instance's IP address and a set of credentials.
//
// Learn more
//
// Remotely Access Fleet Instances (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-remote-access.html)
//
// Debug Fleet Issues (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creating-debug.html)
//
// Related operations
//
//    * DescribeInstances
//
//    * GetInstanceAccess
//
//    // Example sending a request using GetInstanceAccessRequest.
//    req := client.GetInstanceAccessRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/GetInstanceAccess
func (c *Client) GetInstanceAccessRequest(input *GetInstanceAccessInput) GetInstanceAccessRequest {
	op := &aws.Operation{
		Name:       opGetInstanceAccess,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetInstanceAccessInput{}
	}

	req := c.newRequest(op, input, &GetInstanceAccessOutput{})

	return GetInstanceAccessRequest{Request: req, Input: input, Copy: c.GetInstanceAccessRequest}
}

// GetInstanceAccessRequest is the request type for the
// GetInstanceAccess API operation.
type GetInstanceAccessRequest struct {
	*aws.Request
	Input *GetInstanceAccessInput
	Copy  func(*GetInstanceAccessInput) GetInstanceAccessRequest
}

// Send marshals and sends the GetInstanceAccess API request.
func (r GetInstanceAccessRequest) Send(ctx context.Context) (*GetInstanceAccessResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetInstanceAccessResponse{
		GetInstanceAccessOutput: r.Request.Data.(*GetInstanceAccessOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetInstanceAccessResponse is the response type for the
// GetInstanceAccess API operation.
type GetInstanceAccessResponse struct {
	*GetInstanceAccessOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetInstanceAccess request.
func (r *GetInstanceAccessResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
