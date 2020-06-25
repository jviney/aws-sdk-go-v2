// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type PutVoiceConnectorTerminationInput struct {
	_ struct{} `type:"structure"`

	// The termination setting details to add.
	//
	// Termination is a required field
	Termination *Termination `type:"structure" required:"true"`

	// The Amazon Chime Voice Connector ID.
	//
	// VoiceConnectorId is a required field
	VoiceConnectorId *string `location:"uri" locationName:"voiceConnectorId" type:"string" required:"true"`
}

// String returns the string representation
func (s PutVoiceConnectorTerminationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutVoiceConnectorTerminationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutVoiceConnectorTerminationInput"}

	if s.Termination == nil {
		invalidParams.Add(aws.NewErrParamRequired("Termination"))
	}

	if s.VoiceConnectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorId"))
	}
	if s.Termination != nil {
		if err := s.Termination.Validate(); err != nil {
			invalidParams.AddNested("Termination", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutVoiceConnectorTerminationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Termination != nil {
		v := s.Termination

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Termination", v, metadata)
	}
	if s.VoiceConnectorId != nil {
		v := *s.VoiceConnectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "voiceConnectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type PutVoiceConnectorTerminationOutput struct {
	_ struct{} `type:"structure"`

	// The updated termination setting details.
	Termination *Termination `type:"structure"`
}

// String returns the string representation
func (s PutVoiceConnectorTerminationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutVoiceConnectorTerminationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Termination != nil {
		v := s.Termination

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Termination", v, metadata)
	}
	return nil
}

const opPutVoiceConnectorTermination = "PutVoiceConnectorTermination"

// PutVoiceConnectorTerminationRequest returns a request value for making API operation for
// Amazon Chime.
//
// Adds termination settings for the specified Amazon Chime Voice Connector.
//
//    // Example sending a request using PutVoiceConnectorTerminationRequest.
//    req := client.PutVoiceConnectorTerminationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/PutVoiceConnectorTermination
func (c *Client) PutVoiceConnectorTerminationRequest(input *PutVoiceConnectorTerminationInput) PutVoiceConnectorTerminationRequest {
	op := &aws.Operation{
		Name:       opPutVoiceConnectorTermination,
		HTTPMethod: "PUT",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}/termination",
	}

	if input == nil {
		input = &PutVoiceConnectorTerminationInput{}
	}

	req := c.newRequest(op, input, &PutVoiceConnectorTerminationOutput{})

	return PutVoiceConnectorTerminationRequest{Request: req, Input: input, Copy: c.PutVoiceConnectorTerminationRequest}
}

// PutVoiceConnectorTerminationRequest is the request type for the
// PutVoiceConnectorTermination API operation.
type PutVoiceConnectorTerminationRequest struct {
	*aws.Request
	Input *PutVoiceConnectorTerminationInput
	Copy  func(*PutVoiceConnectorTerminationInput) PutVoiceConnectorTerminationRequest
}

// Send marshals and sends the PutVoiceConnectorTermination API request.
func (r PutVoiceConnectorTerminationRequest) Send(ctx context.Context) (*PutVoiceConnectorTerminationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutVoiceConnectorTerminationResponse{
		PutVoiceConnectorTerminationOutput: r.Request.Data.(*PutVoiceConnectorTerminationOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutVoiceConnectorTerminationResponse is the response type for the
// PutVoiceConnectorTermination API operation.
type PutVoiceConnectorTerminationResponse struct {
	*PutVoiceConnectorTerminationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutVoiceConnectorTermination request.
func (r *PutVoiceConnectorTerminationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
