// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DescribeInputDeviceInput struct {
	_ struct{} `type:"structure"`

	// InputDeviceId is a required field
	InputDeviceId *string `location:"uri" locationName:"inputDeviceId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeInputDeviceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInputDeviceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeInputDeviceInput"}

	if s.InputDeviceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputDeviceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeInputDeviceInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.InputDeviceId != nil {
		v := *s.InputDeviceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "inputDeviceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeInputDeviceOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `locationName:"arn" type:"string"`

	// The state of the connection between the input device and AWS.
	ConnectionState InputDeviceConnectionState `locationName:"connectionState" type:"string" enum:"true"`

	// The status of the action to synchronize the device configuration. If you
	// change the configuration of the input device (for example, the maximum bitrate),
	// MediaLive sends the new data to the device. The device might not update itself
	// immediately. SYNCED means the device has updated its configuration. SYNCING
	// means that it has not updated its configuration.
	DeviceSettingsSyncState DeviceSettingsSyncState `locationName:"deviceSettingsSyncState" type:"string" enum:"true"`

	// Settings that describe the active source from the input device, and the video
	// characteristics of that source.
	HdDeviceSettings *InputDeviceHdSettings `locationName:"hdDeviceSettings" type:"structure"`

	Id *string `locationName:"id" type:"string"`

	MacAddress *string `locationName:"macAddress" type:"string"`

	Name *string `locationName:"name" type:"string"`

	// The network settings for the input device.
	NetworkSettings *InputDeviceNetworkSettings `locationName:"networkSettings" type:"structure"`

	SerialNumber *string `locationName:"serialNumber" type:"string"`

	// The type of the input device. For an AWS Elemental Link device that outputs
	// resolutions up to 1080, choose "HD".
	Type InputDeviceType `locationName:"type" type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeInputDeviceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeInputDeviceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ConnectionState) > 0 {
		v := s.ConnectionState

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "connectionState", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.DeviceSettingsSyncState) > 0 {
		v := s.DeviceSettingsSyncState

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deviceSettingsSyncState", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.HdDeviceSettings != nil {
		v := s.HdDeviceSettings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "hdDeviceSettings", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MacAddress != nil {
		v := *s.MacAddress

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "macAddress", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NetworkSettings != nil {
		v := s.NetworkSettings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "networkSettings", v, metadata)
	}
	if s.SerialNumber != nil {
		v := *s.SerialNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "serialNumber", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opDescribeInputDevice = "DescribeInputDevice"

// DescribeInputDeviceRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// Gets the details for the input device
//
//    // Example sending a request using DescribeInputDeviceRequest.
//    req := client.DescribeInputDeviceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/DescribeInputDevice
func (c *Client) DescribeInputDeviceRequest(input *DescribeInputDeviceInput) DescribeInputDeviceRequest {
	op := &aws.Operation{
		Name:       opDescribeInputDevice,
		HTTPMethod: "GET",
		HTTPPath:   "/prod/inputDevices/{inputDeviceId}",
	}

	if input == nil {
		input = &DescribeInputDeviceInput{}
	}

	req := c.newRequest(op, input, &DescribeInputDeviceOutput{})

	return DescribeInputDeviceRequest{Request: req, Input: input, Copy: c.DescribeInputDeviceRequest}
}

// DescribeInputDeviceRequest is the request type for the
// DescribeInputDevice API operation.
type DescribeInputDeviceRequest struct {
	*aws.Request
	Input *DescribeInputDeviceInput
	Copy  func(*DescribeInputDeviceInput) DescribeInputDeviceRequest
}

// Send marshals and sends the DescribeInputDevice API request.
func (r DescribeInputDeviceRequest) Send(ctx context.Context) (*DescribeInputDeviceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInputDeviceResponse{
		DescribeInputDeviceOutput: r.Request.Data.(*DescribeInputDeviceOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeInputDeviceResponse is the response type for the
// DescribeInputDevice API operation.
type DescribeInputDeviceResponse struct {
	*DescribeInputDeviceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInputDevice request.
func (r *DescribeInputDeviceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
