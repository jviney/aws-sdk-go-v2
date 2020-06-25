// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot1clickprojects

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type AssociateDeviceWithPlacementInput struct {
	_ struct{} `type:"structure"`

	// The ID of the physical device to be associated with the given placement in
	// the project. Note that a mandatory 4 character prefix is required for all
	// deviceId values.
	//
	// DeviceId is a required field
	DeviceId *string `locationName:"deviceId" min:"1" type:"string" required:"true"`

	// The device template name to associate with the device ID.
	//
	// DeviceTemplateName is a required field
	DeviceTemplateName *string `location:"uri" locationName:"deviceTemplateName" min:"1" type:"string" required:"true"`

	// The name of the placement in which to associate the device.
	//
	// PlacementName is a required field
	PlacementName *string `location:"uri" locationName:"placementName" min:"1" type:"string" required:"true"`

	// The name of the project containing the placement in which to associate the
	// device.
	//
	// ProjectName is a required field
	ProjectName *string `location:"uri" locationName:"projectName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateDeviceWithPlacementInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateDeviceWithPlacementInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateDeviceWithPlacementInput"}

	if s.DeviceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceId"))
	}
	if s.DeviceId != nil && len(*s.DeviceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeviceId", 1))
	}

	if s.DeviceTemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceTemplateName"))
	}
	if s.DeviceTemplateName != nil && len(*s.DeviceTemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeviceTemplateName", 1))
	}

	if s.PlacementName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlacementName"))
	}
	if s.PlacementName != nil && len(*s.PlacementName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PlacementName", 1))
	}

	if s.ProjectName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectName"))
	}
	if s.ProjectName != nil && len(*s.ProjectName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociateDeviceWithPlacementInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DeviceId != nil {
		v := *s.DeviceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deviceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DeviceTemplateName != nil {
		v := *s.DeviceTemplateName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "deviceTemplateName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PlacementName != nil {
		v := *s.PlacementName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "placementName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProjectName != nil {
		v := *s.ProjectName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "projectName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type AssociateDeviceWithPlacementOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateDeviceWithPlacementOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociateDeviceWithPlacementOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opAssociateDeviceWithPlacement = "AssociateDeviceWithPlacement"

// AssociateDeviceWithPlacementRequest returns a request value for making API operation for
// AWS IoT 1-Click Projects Service.
//
// Associates a physical device with a placement.
//
//    // Example sending a request using AssociateDeviceWithPlacementRequest.
//    req := client.AssociateDeviceWithPlacementRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/AssociateDeviceWithPlacement
func (c *Client) AssociateDeviceWithPlacementRequest(input *AssociateDeviceWithPlacementInput) AssociateDeviceWithPlacementRequest {
	op := &aws.Operation{
		Name:       opAssociateDeviceWithPlacement,
		HTTPMethod: "PUT",
		HTTPPath:   "/projects/{projectName}/placements/{placementName}/devices/{deviceTemplateName}",
	}

	if input == nil {
		input = &AssociateDeviceWithPlacementInput{}
	}

	req := c.newRequest(op, input, &AssociateDeviceWithPlacementOutput{})

	return AssociateDeviceWithPlacementRequest{Request: req, Input: input, Copy: c.AssociateDeviceWithPlacementRequest}
}

// AssociateDeviceWithPlacementRequest is the request type for the
// AssociateDeviceWithPlacement API operation.
type AssociateDeviceWithPlacementRequest struct {
	*aws.Request
	Input *AssociateDeviceWithPlacementInput
	Copy  func(*AssociateDeviceWithPlacementInput) AssociateDeviceWithPlacementRequest
}

// Send marshals and sends the AssociateDeviceWithPlacement API request.
func (r AssociateDeviceWithPlacementRequest) Send(ctx context.Context) (*AssociateDeviceWithPlacementResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateDeviceWithPlacementResponse{
		AssociateDeviceWithPlacementOutput: r.Request.Data.(*AssociateDeviceWithPlacementOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateDeviceWithPlacementResponse is the response type for the
// AssociateDeviceWithPlacement API operation.
type AssociateDeviceWithPlacementResponse struct {
	*AssociateDeviceWithPlacementOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateDeviceWithPlacement request.
func (r *AssociateDeviceWithPlacementResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
