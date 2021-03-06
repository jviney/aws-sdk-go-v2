// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/query"
)

type EnableMFADeviceInput struct {
	_ struct{} `type:"structure"`

	// An authentication code emitted by the device.
	//
	// The format for this parameter is a string of six digits.
	//
	// Submit your request immediately after generating the authentication codes.
	// If you generate the codes and then wait too long to submit the request, the
	// MFA device successfully associates with the user but the MFA device becomes
	// out of sync. This happens because time-based one-time passwords (TOTP) expire
	// after a short period of time. If this happens, you can resync the device
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_mfa_sync.html).
	//
	// AuthenticationCode1 is a required field
	AuthenticationCode1 *string `min:"6" type:"string" required:"true"`

	// A subsequent authentication code emitted by the device.
	//
	// The format for this parameter is a string of six digits.
	//
	// Submit your request immediately after generating the authentication codes.
	// If you generate the codes and then wait too long to submit the request, the
	// MFA device successfully associates with the user but the MFA device becomes
	// out of sync. This happens because time-based one-time passwords (TOTP) expire
	// after a short period of time. If this happens, you can resync the device
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_mfa_sync.html).
	//
	// AuthenticationCode2 is a required field
	AuthenticationCode2 *string `min:"6" type:"string" required:"true"`

	// The serial number that uniquely identifies the MFA device. For virtual MFA
	// devices, the serial number is the device ARN.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: =,.@:/-
	//
	// SerialNumber is a required field
	SerialNumber *string `min:"9" type:"string" required:"true"`

	// The name of the IAM user for whom you want to enable the MFA device.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s EnableMFADeviceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableMFADeviceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EnableMFADeviceInput"}

	if s.AuthenticationCode1 == nil {
		invalidParams.Add(aws.NewErrParamRequired("AuthenticationCode1"))
	}
	if s.AuthenticationCode1 != nil && len(*s.AuthenticationCode1) < 6 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationCode1", 6))
	}

	if s.AuthenticationCode2 == nil {
		invalidParams.Add(aws.NewErrParamRequired("AuthenticationCode2"))
	}
	if s.AuthenticationCode2 != nil && len(*s.AuthenticationCode2) < 6 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationCode2", 6))
	}

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

type EnableMFADeviceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EnableMFADeviceOutput) String() string {
	return awsutil.Prettify(s)
}

const opEnableMFADevice = "EnableMFADevice"

// EnableMFADeviceRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Enables the specified MFA device and associates it with the specified IAM
// user. When enabled, the MFA device is required for every subsequent login
// by the IAM user associated with the device.
//
//    // Example sending a request using EnableMFADeviceRequest.
//    req := client.EnableMFADeviceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/EnableMFADevice
func (c *Client) EnableMFADeviceRequest(input *EnableMFADeviceInput) EnableMFADeviceRequest {
	op := &aws.Operation{
		Name:       opEnableMFADevice,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableMFADeviceInput{}
	}

	req := c.newRequest(op, input, &EnableMFADeviceOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return EnableMFADeviceRequest{Request: req, Input: input, Copy: c.EnableMFADeviceRequest}
}

// EnableMFADeviceRequest is the request type for the
// EnableMFADevice API operation.
type EnableMFADeviceRequest struct {
	*aws.Request
	Input *EnableMFADeviceInput
	Copy  func(*EnableMFADeviceInput) EnableMFADeviceRequest
}

// Send marshals and sends the EnableMFADevice API request.
func (r EnableMFADeviceRequest) Send(ctx context.Context) (*EnableMFADeviceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableMFADeviceResponse{
		EnableMFADeviceOutput: r.Request.Data.(*EnableMFADeviceOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableMFADeviceResponse is the response type for the
// EnableMFADevice API operation.
type EnableMFADeviceResponse struct {
	*EnableMFADeviceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableMFADevice request.
func (r *EnableMFADeviceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
