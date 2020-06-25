// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/query"
)

type DeactivateMFADeviceInput struct {
	_ struct{} `type:"structure"`

	// The serial number that uniquely identifies the MFA device. For virtual MFA
	// devices, the serial number is the device ARN.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: =,.@:/-
	//
	// SerialNumber is a required field
	SerialNumber *string `min:"9" type:"string" required:"true"`

	// The name of the user whose MFA device you want to deactivate.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeactivateMFADeviceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeactivateMFADeviceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeactivateMFADeviceInput"}

	if s.SerialNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("SerialNumber"))
	}
	if s.SerialNumber != nil && len(*s.SerialNumber) < 9 {
		invalidParams.Add(aws.NewErrParamMinLen("SerialNumber", 9))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeactivateMFADeviceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeactivateMFADeviceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeactivateMFADevice = "DeactivateMFADevice"

// DeactivateMFADeviceRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Deactivates the specified MFA device and removes it from association with
// the user name for which it was originally enabled.
//
// For more information about creating and working with virtual MFA devices,
// go to Enabling a Virtual Multi-factor Authentication (MFA) Device (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_VirtualMFA.html)
// in the IAM User Guide.
//
//    // Example sending a request using DeactivateMFADeviceRequest.
//    req := client.DeactivateMFADeviceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeactivateMFADevice
func (c *Client) DeactivateMFADeviceRequest(input *DeactivateMFADeviceInput) DeactivateMFADeviceRequest {
	op := &aws.Operation{
		Name:       opDeactivateMFADevice,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeactivateMFADeviceInput{}
	}

	req := c.newRequest(op, input, &DeactivateMFADeviceOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeactivateMFADeviceRequest{Request: req, Input: input, Copy: c.DeactivateMFADeviceRequest}
}

// DeactivateMFADeviceRequest is the request type for the
// DeactivateMFADevice API operation.
type DeactivateMFADeviceRequest struct {
	*aws.Request
	Input *DeactivateMFADeviceInput
	Copy  func(*DeactivateMFADeviceInput) DeactivateMFADeviceRequest
}

// Send marshals and sends the DeactivateMFADevice API request.
func (r DeactivateMFADeviceRequest) Send(ctx context.Context) (*DeactivateMFADeviceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeactivateMFADeviceResponse{
		DeactivateMFADeviceOutput: r.Request.Data.(*DeactivateMFADeviceOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeactivateMFADeviceResponse is the response type for the
// DeactivateMFADevice API operation.
type DeactivateMFADeviceResponse struct {
	*DeactivateMFADeviceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeactivateMFADevice request.
func (r *DeactivateMFADeviceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
