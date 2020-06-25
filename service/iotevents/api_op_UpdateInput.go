// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotevents

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type UpdateInputInput struct {
	_ struct{} `type:"structure"`

	// The definition of the input.
	//
	// InputDefinition is a required field
	InputDefinition *InputDefinition `locationName:"inputDefinition" type:"structure" required:"true"`

	// A brief description of the input.
	InputDescription *string `locationName:"inputDescription" type:"string"`

	// The name of the input you want to update.
	//
	// InputName is a required field
	InputName *string `location:"uri" locationName:"inputName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateInputInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateInputInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateInputInput"}

	if s.InputDefinition == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputDefinition"))
	}

	if s.InputName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputName"))
	}
	if s.InputName != nil && len(*s.InputName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InputName", 1))
	}
	if s.InputDefinition != nil {
		if err := s.InputDefinition.Validate(); err != nil {
			invalidParams.AddNested("InputDefinition", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateInputInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.InputDefinition != nil {
		v := s.InputDefinition

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "inputDefinition", v, metadata)
	}
	if s.InputDescription != nil {
		v := *s.InputDescription

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "inputDescription", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.InputName != nil {
		v := *s.InputName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "inputName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateInputOutput struct {
	_ struct{} `type:"structure"`

	// Information about the configuration of the input.
	InputConfiguration *InputConfiguration `locationName:"inputConfiguration" type:"structure"`
}

// String returns the string representation
func (s UpdateInputOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateInputOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.InputConfiguration != nil {
		v := s.InputConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "inputConfiguration", v, metadata)
	}
	return nil
}

const opUpdateInput = "UpdateInput"

// UpdateInputRequest returns a request value for making API operation for
// AWS IoT Events.
//
// Updates an input.
//
//    // Example sending a request using UpdateInputRequest.
//    req := client.UpdateInputRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-2018-07-27/UpdateInput
func (c *Client) UpdateInputRequest(input *UpdateInputInput) UpdateInputRequest {
	op := &aws.Operation{
		Name:       opUpdateInput,
		HTTPMethod: "PUT",
		HTTPPath:   "/inputs/{inputName}",
	}

	if input == nil {
		input = &UpdateInputInput{}
	}

	req := c.newRequest(op, input, &UpdateInputOutput{})

	return UpdateInputRequest{Request: req, Input: input, Copy: c.UpdateInputRequest}
}

// UpdateInputRequest is the request type for the
// UpdateInput API operation.
type UpdateInputRequest struct {
	*aws.Request
	Input *UpdateInputInput
	Copy  func(*UpdateInputInput) UpdateInputRequest
}

// Send marshals and sends the UpdateInput API request.
func (r UpdateInputRequest) Send(ctx context.Context) (*UpdateInputResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateInputResponse{
		UpdateInputOutput: r.Request.Data.(*UpdateInputOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateInputResponse is the response type for the
// UpdateInput API operation.
type UpdateInputResponse struct {
	*UpdateInputOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateInput request.
func (r *UpdateInputResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
