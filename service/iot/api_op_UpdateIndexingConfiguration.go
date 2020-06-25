// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type UpdateIndexingConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Thing group indexing configuration.
	ThingGroupIndexingConfiguration *ThingGroupIndexingConfiguration `locationName:"thingGroupIndexingConfiguration" type:"structure"`

	// Thing indexing configuration.
	ThingIndexingConfiguration *ThingIndexingConfiguration `locationName:"thingIndexingConfiguration" type:"structure"`
}

// String returns the string representation
func (s UpdateIndexingConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateIndexingConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateIndexingConfigurationInput"}
	if s.ThingGroupIndexingConfiguration != nil {
		if err := s.ThingGroupIndexingConfiguration.Validate(); err != nil {
			invalidParams.AddNested("ThingGroupIndexingConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.ThingIndexingConfiguration != nil {
		if err := s.ThingIndexingConfiguration.Validate(); err != nil {
			invalidParams.AddNested("ThingIndexingConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateIndexingConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ThingGroupIndexingConfiguration != nil {
		v := s.ThingGroupIndexingConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "thingGroupIndexingConfiguration", v, metadata)
	}
	if s.ThingIndexingConfiguration != nil {
		v := s.ThingIndexingConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "thingIndexingConfiguration", v, metadata)
	}
	return nil
}

type UpdateIndexingConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateIndexingConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateIndexingConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateIndexingConfiguration = "UpdateIndexingConfiguration"

// UpdateIndexingConfigurationRequest returns a request value for making API operation for
// AWS IoT.
//
// Updates the search configuration.
//
//    // Example sending a request using UpdateIndexingConfigurationRequest.
//    req := client.UpdateIndexingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateIndexingConfigurationRequest(input *UpdateIndexingConfigurationInput) UpdateIndexingConfigurationRequest {
	op := &aws.Operation{
		Name:       opUpdateIndexingConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/indexing/config",
	}

	if input == nil {
		input = &UpdateIndexingConfigurationInput{}
	}

	req := c.newRequest(op, input, &UpdateIndexingConfigurationOutput{})

	return UpdateIndexingConfigurationRequest{Request: req, Input: input, Copy: c.UpdateIndexingConfigurationRequest}
}

// UpdateIndexingConfigurationRequest is the request type for the
// UpdateIndexingConfiguration API operation.
type UpdateIndexingConfigurationRequest struct {
	*aws.Request
	Input *UpdateIndexingConfigurationInput
	Copy  func(*UpdateIndexingConfigurationInput) UpdateIndexingConfigurationRequest
}

// Send marshals and sends the UpdateIndexingConfiguration API request.
func (r UpdateIndexingConfigurationRequest) Send(ctx context.Context) (*UpdateIndexingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateIndexingConfigurationResponse{
		UpdateIndexingConfigurationOutput: r.Request.Data.(*UpdateIndexingConfigurationOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateIndexingConfigurationResponse is the response type for the
// UpdateIndexingConfiguration API operation.
type UpdateIndexingConfigurationResponse struct {
	*UpdateIndexingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateIndexingConfiguration request.
func (r *UpdateIndexingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
