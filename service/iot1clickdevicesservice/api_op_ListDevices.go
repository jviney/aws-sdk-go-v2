// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot1clickdevicesservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListDevicesInput struct {
	_ struct{} `type:"structure"`

	DeviceType *string `location:"querystring" locationName:"deviceType" type:"string"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDevicesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDevicesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDevicesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDevicesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DeviceType != nil {
		v := *s.DeviceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "deviceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	return nil
}

type ListDevicesOutput struct {
	_ struct{} `type:"structure"`

	// A list of devices.
	Devices []DeviceDescription `locationName:"devices" type:"list"`

	// The token to retrieve the next set of results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDevicesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDevicesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Devices != nil {
		v := s.Devices

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "devices", metadata)
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

const opListDevices = "ListDevices"

// ListDevicesRequest returns a request value for making API operation for
// AWS IoT 1-Click Devices Service.
//
// Lists the 1-Click compatible devices associated with your AWS account.
//
//    // Example sending a request using ListDevicesRequest.
//    req := client.ListDevicesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devices-2018-05-14/ListDevices
func (c *Client) ListDevicesRequest(input *ListDevicesInput) ListDevicesRequest {
	op := &aws.Operation{
		Name:       opListDevices,
		HTTPMethod: "GET",
		HTTPPath:   "/devices",
	}

	if input == nil {
		input = &ListDevicesInput{}
	}

	req := c.newRequest(op, input, &ListDevicesOutput{})

	return ListDevicesRequest{Request: req, Input: input, Copy: c.ListDevicesRequest}
}

// ListDevicesRequest is the request type for the
// ListDevices API operation.
type ListDevicesRequest struct {
	*aws.Request
	Input *ListDevicesInput
	Copy  func(*ListDevicesInput) ListDevicesRequest
}

// Send marshals and sends the ListDevices API request.
func (r ListDevicesRequest) Send(ctx context.Context) (*ListDevicesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDevicesResponse{
		ListDevicesOutput: r.Request.Data.(*ListDevicesOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDevicesResponse is the response type for the
// ListDevices API operation.
type ListDevicesResponse struct {
	*ListDevicesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDevices request.
func (r *ListDevicesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
