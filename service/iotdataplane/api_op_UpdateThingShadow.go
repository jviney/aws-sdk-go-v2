// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotdataplane

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// The input for the UpdateThingShadow operation.
type UpdateThingShadowInput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// The state information, in JSON format.
	//
	// Payload is a required field
	Payload []byte `locationName:"payload" type:"blob" required:"true"`

	// The name of the thing.
	//
	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateThingShadowInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateThingShadowInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateThingShadowInput"}

	if s.Payload == nil {
		invalidParams.Add(aws.NewErrParamRequired("Payload"))
	}

	if s.ThingName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateThingShadowInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Payload != nil {
		v := s.Payload

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "payload", protocol.BytesStream(v), metadata)
	}
	return nil
}

// The output from the UpdateThingShadow operation.
type UpdateThingShadowOutput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// The state information, in JSON format.
	Payload []byte `locationName:"payload" type:"blob"`
}

// String returns the string representation
func (s UpdateThingShadowOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateThingShadowOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Payload != nil {
		v := s.Payload

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "payload", protocol.BytesStream(v), metadata)
	}
	return nil
}

const opUpdateThingShadow = "UpdateThingShadow"

// UpdateThingShadowRequest returns a request value for making API operation for
// AWS IoT Data Plane.
//
// Updates the thing shadow for the specified thing.
//
// For more information, see UpdateThingShadow (http://docs.aws.amazon.com/iot/latest/developerguide/API_UpdateThingShadow.html)
// in the AWS IoT Developer Guide.
//
//    // Example sending a request using UpdateThingShadowRequest.
//    req := client.UpdateThingShadowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateThingShadowRequest(input *UpdateThingShadowInput) UpdateThingShadowRequest {
	op := &aws.Operation{
		Name:       opUpdateThingShadow,
		HTTPMethod: "POST",
		HTTPPath:   "/things/{thingName}/shadow",
	}

	if input == nil {
		input = &UpdateThingShadowInput{}
	}

	req := c.newRequest(op, input, &UpdateThingShadowOutput{})

	return UpdateThingShadowRequest{Request: req, Input: input, Copy: c.UpdateThingShadowRequest}
}

// UpdateThingShadowRequest is the request type for the
// UpdateThingShadow API operation.
type UpdateThingShadowRequest struct {
	*aws.Request
	Input *UpdateThingShadowInput
	Copy  func(*UpdateThingShadowInput) UpdateThingShadowRequest
}

// Send marshals and sends the UpdateThingShadow API request.
func (r UpdateThingShadowRequest) Send(ctx context.Context) (*UpdateThingShadowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateThingShadowResponse{
		UpdateThingShadowOutput: r.Request.Data.(*UpdateThingShadowOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateThingShadowResponse is the response type for the
// UpdateThingShadow API operation.
type UpdateThingShadowResponse struct {
	*UpdateThingShadowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateThingShadow request.
func (r *UpdateThingShadowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
