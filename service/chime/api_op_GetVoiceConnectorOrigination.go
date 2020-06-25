// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetVoiceConnectorOriginationInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime Voice Connector ID.
	//
	// VoiceConnectorId is a required field
	VoiceConnectorId *string `location:"uri" locationName:"voiceConnectorId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetVoiceConnectorOriginationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetVoiceConnectorOriginationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetVoiceConnectorOriginationInput"}

	if s.VoiceConnectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetVoiceConnectorOriginationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.VoiceConnectorId != nil {
		v := *s.VoiceConnectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "voiceConnectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetVoiceConnectorOriginationOutput struct {
	_ struct{} `type:"structure"`

	// The origination setting details.
	Origination *Origination `type:"structure"`
}

// String returns the string representation
func (s GetVoiceConnectorOriginationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetVoiceConnectorOriginationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Origination != nil {
		v := s.Origination

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Origination", v, metadata)
	}
	return nil
}

const opGetVoiceConnectorOrigination = "GetVoiceConnectorOrigination"

// GetVoiceConnectorOriginationRequest returns a request value for making API operation for
// Amazon Chime.
//
// Retrieves origination setting details for the specified Amazon Chime Voice
// Connector.
//
//    // Example sending a request using GetVoiceConnectorOriginationRequest.
//    req := client.GetVoiceConnectorOriginationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetVoiceConnectorOrigination
func (c *Client) GetVoiceConnectorOriginationRequest(input *GetVoiceConnectorOriginationInput) GetVoiceConnectorOriginationRequest {
	op := &aws.Operation{
		Name:       opGetVoiceConnectorOrigination,
		HTTPMethod: "GET",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}/origination",
	}

	if input == nil {
		input = &GetVoiceConnectorOriginationInput{}
	}

	req := c.newRequest(op, input, &GetVoiceConnectorOriginationOutput{})

	return GetVoiceConnectorOriginationRequest{Request: req, Input: input, Copy: c.GetVoiceConnectorOriginationRequest}
}

// GetVoiceConnectorOriginationRequest is the request type for the
// GetVoiceConnectorOrigination API operation.
type GetVoiceConnectorOriginationRequest struct {
	*aws.Request
	Input *GetVoiceConnectorOriginationInput
	Copy  func(*GetVoiceConnectorOriginationInput) GetVoiceConnectorOriginationRequest
}

// Send marshals and sends the GetVoiceConnectorOrigination API request.
func (r GetVoiceConnectorOriginationRequest) Send(ctx context.Context) (*GetVoiceConnectorOriginationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetVoiceConnectorOriginationResponse{
		GetVoiceConnectorOriginationOutput: r.Request.Data.(*GetVoiceConnectorOriginationOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetVoiceConnectorOriginationResponse is the response type for the
// GetVoiceConnectorOrigination API operation.
type GetVoiceConnectorOriginationResponse struct {
	*GetVoiceConnectorOriginationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetVoiceConnectorOrigination request.
func (r *GetVoiceConnectorOriginationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
