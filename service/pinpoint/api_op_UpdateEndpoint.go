// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type UpdateEndpointInput struct {
	_ struct{} `type:"structure" payload:"EndpointRequest"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// EndpointId is a required field
	EndpointId *string `location:"uri" locationName:"endpoint-id" type:"string" required:"true"`

	// Specifies the channel type and other settings for an endpoint.
	//
	// EndpointRequest is a required field
	EndpointRequest *EndpointRequest `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateEndpointInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.EndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointId"))
	}

	if s.EndpointRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointRequest"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateEndpointInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EndpointId != nil {
		v := *s.EndpointId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "endpoint-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EndpointRequest != nil {
		v := s.EndpointRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "EndpointRequest", v, metadata)
	}
	return nil
}

type UpdateEndpointOutput struct {
	_ struct{} `type:"structure" payload:"MessageBody"`

	// Provides information about an API request or response.
	//
	// MessageBody is a required field
	MessageBody *MessageBody `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateEndpointOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.MessageBody != nil {
		v := s.MessageBody

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "MessageBody", v, metadata)
	}
	return nil
}

const opUpdateEndpoint = "UpdateEndpoint"

// UpdateEndpointRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Creates a new endpoint for an application or updates the settings and attributes
// of an existing endpoint for an application. You can also use this operation
// to define custom attributes for an endpoint. If an update includes one or
// more values for a custom attribute, Amazon Pinpoint replaces (overwrites)
// any existing values with the new values.
//
//    // Example sending a request using UpdateEndpointRequest.
//    req := client.UpdateEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateEndpoint
func (c *Client) UpdateEndpointRequest(input *UpdateEndpointInput) UpdateEndpointRequest {
	op := &aws.Operation{
		Name:       opUpdateEndpoint,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/apps/{application-id}/endpoints/{endpoint-id}",
	}

	if input == nil {
		input = &UpdateEndpointInput{}
	}

	req := c.newRequest(op, input, &UpdateEndpointOutput{})

	return UpdateEndpointRequest{Request: req, Input: input, Copy: c.UpdateEndpointRequest}
}

// UpdateEndpointRequest is the request type for the
// UpdateEndpoint API operation.
type UpdateEndpointRequest struct {
	*aws.Request
	Input *UpdateEndpointInput
	Copy  func(*UpdateEndpointInput) UpdateEndpointRequest
}

// Send marshals and sends the UpdateEndpoint API request.
func (r UpdateEndpointRequest) Send(ctx context.Context) (*UpdateEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateEndpointResponse{
		UpdateEndpointOutput: r.Request.Data.(*UpdateEndpointOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateEndpointResponse is the response type for the
// UpdateEndpoint API operation.
type UpdateEndpointResponse struct {
	*UpdateEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateEndpoint request.
func (r *UpdateEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
