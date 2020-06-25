// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateVirtualMFADeviceInput struct {
	_ struct{} `type:"structure"`

	// The path for the virtual MFA device. For more information about paths, see
	// IAM Identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html)
	// in the IAM User Guide.
	//
	// This parameter is optional. If it is not included, it defaults to a slash
	// (/).
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of either a forward slash (/) by itself
	// or a string that must begin and end with forward slashes. In addition, it
	// can contain any ASCII character from the ! (\u0021) through the DEL character
	// (\u007F), including most punctuation characters, digits, and upper and lowercased
	// letters.
	Path *string `min:"1" type:"string"`

	// The name of the virtual MFA device. Use with path to uniquely identify a
	// virtual MFA device.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// VirtualMFADeviceName is a required field
	VirtualMFADeviceName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateVirtualMFADeviceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVirtualMFADeviceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVirtualMFADeviceInput"}
	if s.Path != nil && len(*s.Path) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Path", 1))
	}

	if s.VirtualMFADeviceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualMFADeviceName"))
	}
	if s.VirtualMFADeviceName != nil && len(*s.VirtualMFADeviceName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VirtualMFADeviceName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful CreateVirtualMFADevice request.
type CreateVirtualMFADeviceOutput struct {
	_ struct{} `type:"structure"`

	// A structure containing details about the new virtual MFA device.
	//
	// VirtualMFADevice is a required field
	VirtualMFADevice *VirtualMFADevice `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateVirtualMFADeviceOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateVirtualMFADevice = "CreateVirtualMFADevice"

// CreateVirtualMFADeviceRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Creates a new virtual MFA device for the AWS account. After creating the
// virtual MFA, use EnableMFADevice to attach the MFA device to an IAM user.
// For more information about creating and working with virtual MFA devices,
// go to Using a Virtual MFA Device (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_VirtualMFA.html)
// in the IAM User Guide.
//
// For information about limits on the number of MFA devices you can create,
// see Limitations on Entities (https://docs.aws.amazon.com/IAM/latest/UserGuide/LimitationsOnEntities.html)
// in the IAM User Guide.
//
// The seed information contained in the QR code and the Base32 string should
// be treated like any other secret access information. In other words, protect
// the seed information as you would your AWS access keys or your passwords.
// After you provision your virtual device, you should ensure that the information
// is destroyed following secure procedures.
//
//    // Example sending a request using CreateVirtualMFADeviceRequest.
//    req := client.CreateVirtualMFADeviceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/CreateVirtualMFADevice
func (c *Client) CreateVirtualMFADeviceRequest(input *CreateVirtualMFADeviceInput) CreateVirtualMFADeviceRequest {
	op := &aws.Operation{
		Name:       opCreateVirtualMFADevice,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVirtualMFADeviceInput{}
	}

	req := c.newRequest(op, input, &CreateVirtualMFADeviceOutput{})

	return CreateVirtualMFADeviceRequest{Request: req, Input: input, Copy: c.CreateVirtualMFADeviceRequest}
}

// CreateVirtualMFADeviceRequest is the request type for the
// CreateVirtualMFADevice API operation.
type CreateVirtualMFADeviceRequest struct {
	*aws.Request
	Input *CreateVirtualMFADeviceInput
	Copy  func(*CreateVirtualMFADeviceInput) CreateVirtualMFADeviceRequest
}

// Send marshals and sends the CreateVirtualMFADevice API request.
func (r CreateVirtualMFADeviceRequest) Send(ctx context.Context) (*CreateVirtualMFADeviceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVirtualMFADeviceResponse{
		CreateVirtualMFADeviceOutput: r.Request.Data.(*CreateVirtualMFADeviceOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVirtualMFADeviceResponse is the response type for the
// CreateVirtualMFADevice API operation.
type CreateVirtualMFADeviceResponse struct {
	*CreateVirtualMFADeviceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVirtualMFADevice request.
func (r *CreateVirtualMFADeviceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
