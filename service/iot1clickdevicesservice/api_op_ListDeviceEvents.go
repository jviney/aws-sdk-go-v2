// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot1clickdevicesservice

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListDeviceEventsInput struct {
	_ struct{} `type:"structure"`

	// DeviceId is a required field
	DeviceId *string `location:"uri" locationName:"deviceId" type:"string" required:"true"`

	// FromTimeStamp is a required field
	FromTimeStamp *time.Time `location:"querystring" locationName:"fromTimeStamp" type:"timestamp" timestampFormat:"iso8601" required:"true"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// ToTimeStamp is a required field
	ToTimeStamp *time.Time `location:"querystring" locationName:"toTimeStamp" type:"timestamp" timestampFormat:"iso8601" required:"true"`
}

// String returns the string representation
func (s ListDeviceEventsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDeviceEventsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDeviceEventsInput"}

	if s.DeviceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceId"))
	}

	if s.FromTimeStamp == nil {
		invalidParams.Add(aws.NewErrParamRequired("FromTimeStamp"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ToTimeStamp == nil {
		invalidParams.Add(aws.NewErrParamRequired("ToTimeStamp"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDeviceEventsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DeviceId != nil {
		v := *s.DeviceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "deviceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FromTimeStamp != nil {
		v := *s.FromTimeStamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "fromTimeStamp",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: false}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ToTimeStamp != nil {
		v := *s.ToTimeStamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "toTimeStamp",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: false}, metadata)
	}
	return nil
}

type ListDeviceEventsOutput struct {
	_ struct{} `type:"structure"`

	Events []DeviceEvent `locationName:"events" type:"list"`

	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDeviceEventsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDeviceEventsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Events != nil {
		v := s.Events

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "events", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListDeviceEvents = "ListDeviceEvents"

// ListDeviceEventsRequest returns a request value for making API operation for
// AWS IoT 1-Click Devices Service.
//
// Using a device ID, returns a DeviceEventsResponse object containing an array
// of events for the device.
//
//    // Example sending a request using ListDeviceEventsRequest.
//    req := client.ListDeviceEventsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devices-2018-05-14/ListDeviceEvents
func (c *Client) ListDeviceEventsRequest(input *ListDeviceEventsInput) ListDeviceEventsRequest {
	op := &aws.Operation{
		Name:       opListDeviceEvents,
		HTTPMethod: "GET",
		HTTPPath:   "/devices/{deviceId}/events",
	}

	if input == nil {
		input = &ListDeviceEventsInput{}
	}

	req := c.newRequest(op, input, &ListDeviceEventsOutput{})

	return ListDeviceEventsRequest{Request: req, Input: input, Copy: c.ListDeviceEventsRequest}
}

// ListDeviceEventsRequest is the request type for the
// ListDeviceEvents API operation.
type ListDeviceEventsRequest struct {
	*aws.Request
	Input *ListDeviceEventsInput
	Copy  func(*ListDeviceEventsInput) ListDeviceEventsRequest
}

// Send marshals and sends the ListDeviceEvents API request.
func (r ListDeviceEventsRequest) Send(ctx context.Context) (*ListDeviceEventsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDeviceEventsResponse{
		ListDeviceEventsOutput: r.Request.Data.(*ListDeviceEventsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDeviceEventsResponse is the response type for the
// ListDeviceEvents API operation.
type ListDeviceEventsResponse struct {
	*ListDeviceEventsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDeviceEvents request.
func (r *ListDeviceEventsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
