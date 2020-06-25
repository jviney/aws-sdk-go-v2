// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package worklink

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DescribeDeviceInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for a registered user's device.
	//
	// DeviceId is a required field
	DeviceId *string `min:"1" type:"string" required:"true"`

	// The ARN of the fleet.
	//
	// FleetArn is a required field
	FleetArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDeviceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDeviceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDeviceInput"}

	if s.DeviceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceId"))
	}
	if s.DeviceId != nil && len(*s.DeviceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeviceId", 1))
	}

	if s.FleetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetArn"))
	}
	if s.FleetArn != nil && len(*s.FleetArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDeviceInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DeviceId != nil {
		v := *s.DeviceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeviceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FleetArn != nil {
		v := *s.FleetArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FleetArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeDeviceOutput struct {
	_ struct{} `type:"structure"`

	// The date that the device first signed in to Amazon WorkLink.
	FirstAccessedTime *time.Time `type:"timestamp"`

	// The date that the device last accessed Amazon WorkLink.
	LastAccessedTime *time.Time `type:"timestamp"`

	// The manufacturer of the device.
	Manufacturer *string `min:"1" type:"string"`

	// The model of the device.
	Model *string `min:"1" type:"string"`

	// The operating system of the device.
	OperatingSystem *string `min:"1" type:"string"`

	// The operating system version of the device.
	OperatingSystemVersion *string `min:"1" type:"string"`

	// The operating system patch level of the device.
	PatchLevel *string `min:"1" type:"string"`

	// The current state of the device.
	Status DeviceStatus `type:"string" enum:"true"`

	// The user name associated with the device.
	Username *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeDeviceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDeviceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FirstAccessedTime != nil {
		v := *s.FirstAccessedTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FirstAccessedTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.LastAccessedTime != nil {
		v := *s.LastAccessedTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastAccessedTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Manufacturer != nil {
		v := *s.Manufacturer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Manufacturer", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Model != nil {
		v := *s.Model

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Model", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OperatingSystem != nil {
		v := *s.OperatingSystem

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "OperatingSystem", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OperatingSystemVersion != nil {
		v := *s.OperatingSystemVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "OperatingSystemVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PatchLevel != nil {
		v := *s.PatchLevel

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PatchLevel", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Username != nil {
		v := *s.Username

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Username", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeDevice = "DescribeDevice"

// DescribeDeviceRequest returns a request value for making API operation for
// Amazon WorkLink.
//
// Provides information about a user's device.
//
//    // Example sending a request using DescribeDeviceRequest.
//    req := client.DescribeDeviceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/DescribeDevice
func (c *Client) DescribeDeviceRequest(input *DescribeDeviceInput) DescribeDeviceRequest {
	op := &aws.Operation{
		Name:       opDescribeDevice,
		HTTPMethod: "POST",
		HTTPPath:   "/describeDevice",
	}

	if input == nil {
		input = &DescribeDeviceInput{}
	}

	req := c.newRequest(op, input, &DescribeDeviceOutput{})

	return DescribeDeviceRequest{Request: req, Input: input, Copy: c.DescribeDeviceRequest}
}

// DescribeDeviceRequest is the request type for the
// DescribeDevice API operation.
type DescribeDeviceRequest struct {
	*aws.Request
	Input *DescribeDeviceInput
	Copy  func(*DescribeDeviceInput) DescribeDeviceRequest
}

// Send marshals and sends the DescribeDevice API request.
func (r DescribeDeviceRequest) Send(ctx context.Context) (*DescribeDeviceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDeviceResponse{
		DescribeDeviceOutput: r.Request.Data.(*DescribeDeviceOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDeviceResponse is the response type for the
// DescribeDevice API operation.
type DescribeDeviceResponse struct {
	*DescribeDeviceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDevice request.
func (r *DescribeDeviceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
