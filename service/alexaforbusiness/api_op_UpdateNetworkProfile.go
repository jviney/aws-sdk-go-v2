// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type UpdateNetworkProfileInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the Private Certificate Authority (PCA) created in AWS Certificate
	// Manager (ACM). This is used to issue certificates to the devices.
	CertificateAuthorityArn *string `type:"string"`

	// The current password of the Wi-Fi network.
	CurrentPassword *string `min:"5" type:"string" sensitive:"true"`

	// Detailed information about a device's network profile.
	Description *string `type:"string"`

	// The ARN of the network profile associated with a device.
	//
	// NetworkProfileArn is a required field
	NetworkProfileArn *string `type:"string" required:"true"`

	// The name of the network profile associated with a device.
	NetworkProfileName *string `min:"1" type:"string"`

	// The next, or subsequent, password of the Wi-Fi network. This password is
	// asynchronously transmitted to the device and is used when the password of
	// the network changes to NextPassword.
	NextPassword *string `type:"string" sensitive:"true"`

	// The root certificate(s) of your authentication server that will be installed
	// on your devices and used to trust your authentication server during EAP negotiation.
	TrustAnchors []string `min:"1" type:"list"`
}

// String returns the string representation
func (s UpdateNetworkProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateNetworkProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateNetworkProfileInput"}
	if s.CurrentPassword != nil && len(*s.CurrentPassword) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("CurrentPassword", 5))
	}

	if s.NetworkProfileArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkProfileArn"))
	}
	if s.NetworkProfileName != nil && len(*s.NetworkProfileName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NetworkProfileName", 1))
	}
	if s.TrustAnchors != nil && len(s.TrustAnchors) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TrustAnchors", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateNetworkProfileOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateNetworkProfileOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateNetworkProfile = "UpdateNetworkProfile"

// UpdateNetworkProfileRequest returns a request value for making API operation for
// Alexa For Business.
//
// Updates a network profile by the network profile ARN.
//
//    // Example sending a request using UpdateNetworkProfileRequest.
//    req := client.UpdateNetworkProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/UpdateNetworkProfile
func (c *Client) UpdateNetworkProfileRequest(input *UpdateNetworkProfileInput) UpdateNetworkProfileRequest {
	op := &aws.Operation{
		Name:       opUpdateNetworkProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateNetworkProfileInput{}
	}

	req := c.newRequest(op, input, &UpdateNetworkProfileOutput{})

	return UpdateNetworkProfileRequest{Request: req, Input: input, Copy: c.UpdateNetworkProfileRequest}
}

// UpdateNetworkProfileRequest is the request type for the
// UpdateNetworkProfile API operation.
type UpdateNetworkProfileRequest struct {
	*aws.Request
	Input *UpdateNetworkProfileInput
	Copy  func(*UpdateNetworkProfileInput) UpdateNetworkProfileRequest
}

// Send marshals and sends the UpdateNetworkProfile API request.
func (r UpdateNetworkProfileRequest) Send(ctx context.Context) (*UpdateNetworkProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateNetworkProfileResponse{
		UpdateNetworkProfileOutput: r.Request.Data.(*UpdateNetworkProfileOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateNetworkProfileResponse is the response type for the
// UpdateNetworkProfile API operation.
type UpdateNetworkProfileResponse struct {
	*UpdateNetworkProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateNetworkProfile request.
func (r *UpdateNetworkProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
