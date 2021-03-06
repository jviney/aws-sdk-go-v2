// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package networkmanager

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type UpdateDeviceInput struct {
	_ struct{} `type:"structure"`

	// A description of the device.
	//
	// Length Constraints: Maximum length of 256 characters.
	Description *string `type:"string"`

	// The ID of the device.
	//
	// DeviceId is a required field
	DeviceId *string `location:"uri" locationName:"deviceId" type:"string" required:"true"`

	// The ID of the global network.
	//
	// GlobalNetworkId is a required field
	GlobalNetworkId *string `location:"uri" locationName:"globalNetworkId" type:"string" required:"true"`

	// Describes a location.
	Location *Location `type:"structure"`

	// The model of the device.
	//
	// Length Constraints: Maximum length of 128 characters.
	Model *string `type:"string"`

	// The serial number of the device.
	//
	// Length Constraints: Maximum length of 128 characters.
	SerialNumber *string `type:"string"`

	// The ID of the site.
	SiteId *string `type:"string"`

	// The type of the device.
	Type *string `type:"string"`

	// The vendor of the device.
	//
	// Length Constraints: Maximum length of 128 characters.
	Vendor *string `type:"string"`
}

// String returns the string representation
func (s UpdateDeviceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDeviceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDeviceInput"}

	if s.DeviceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceId"))
	}

	if s.GlobalNetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GlobalNetworkId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDeviceInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Location != nil {
		v := s.Location

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Location", v, metadata)
	}
	if s.Model != nil {
		v := *s.Model

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Model", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SerialNumber != nil {
		v := *s.SerialNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SerialNumber", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SiteId != nil {
		v := *s.SiteId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SiteId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Type != nil {
		v := *s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Vendor != nil {
		v := *s.Vendor

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Vendor", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DeviceId != nil {
		v := *s.DeviceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "deviceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GlobalNetworkId != nil {
		v := *s.GlobalNetworkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "globalNetworkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateDeviceOutput struct {
	_ struct{} `type:"structure"`

	// Information about the device.
	Device *Device `type:"structure"`
}

// String returns the string representation
func (s UpdateDeviceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDeviceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Device != nil {
		v := s.Device

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Device", v, metadata)
	}
	return nil
}

const opUpdateDevice = "UpdateDevice"

// UpdateDeviceRequest returns a request value for making API operation for
// AWS Network Manager.
//
// Updates the details for an existing device. To remove information for any
// of the parameters, specify an empty string.
//
//    // Example sending a request using UpdateDeviceRequest.
//    req := client.UpdateDeviceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/networkmanager-2019-07-05/UpdateDevice
func (c *Client) UpdateDeviceRequest(input *UpdateDeviceInput) UpdateDeviceRequest {
	op := &aws.Operation{
		Name:       opUpdateDevice,
		HTTPMethod: "PATCH",
		HTTPPath:   "/global-networks/{globalNetworkId}/devices/{deviceId}",
	}

	if input == nil {
		input = &UpdateDeviceInput{}
	}

	req := c.newRequest(op, input, &UpdateDeviceOutput{})

	return UpdateDeviceRequest{Request: req, Input: input, Copy: c.UpdateDeviceRequest}
}

// UpdateDeviceRequest is the request type for the
// UpdateDevice API operation.
type UpdateDeviceRequest struct {
	*aws.Request
	Input *UpdateDeviceInput
	Copy  func(*UpdateDeviceInput) UpdateDeviceRequest
}

// Send marshals and sends the UpdateDevice API request.
func (r UpdateDeviceRequest) Send(ctx context.Context) (*UpdateDeviceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDeviceResponse{
		UpdateDeviceOutput: r.Request.Data.(*UpdateDeviceOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDeviceResponse is the response type for the
// UpdateDevice API operation.
type UpdateDeviceResponse struct {
	*UpdateDeviceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDevice request.
func (r *UpdateDeviceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
