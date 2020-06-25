// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteVoiceConnectorTerminationCredentialsInput struct {
	_ struct{} `type:"structure"`

	// The RFC2617 compliant username associated with the SIP credentials, in US-ASCII
	// format.
	Usernames []string `type:"list"`

	// The Amazon Chime Voice Connector ID.
	//
	// VoiceConnectorId is a required field
	VoiceConnectorId *string `location:"uri" locationName:"voiceConnectorId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVoiceConnectorTerminationCredentialsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVoiceConnectorTerminationCredentialsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVoiceConnectorTerminationCredentialsInput"}

	if s.VoiceConnectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVoiceConnectorTerminationCredentialsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Usernames != nil {
		v := s.Usernames

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Usernames", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.VoiceConnectorId != nil {
		v := *s.VoiceConnectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "voiceConnectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteVoiceConnectorTerminationCredentialsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteVoiceConnectorTerminationCredentialsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVoiceConnectorTerminationCredentialsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteVoiceConnectorTerminationCredentials = "DeleteVoiceConnectorTerminationCredentials"

// DeleteVoiceConnectorTerminationCredentialsRequest returns a request value for making API operation for
// Amazon Chime.
//
// Deletes the specified SIP credentials used by your equipment to authenticate
// during call termination.
//
//    // Example sending a request using DeleteVoiceConnectorTerminationCredentialsRequest.
//    req := client.DeleteVoiceConnectorTerminationCredentialsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DeleteVoiceConnectorTerminationCredentials
func (c *Client) DeleteVoiceConnectorTerminationCredentialsRequest(input *DeleteVoiceConnectorTerminationCredentialsInput) DeleteVoiceConnectorTerminationCredentialsRequest {
	op := &aws.Operation{
		Name:       opDeleteVoiceConnectorTerminationCredentials,
		HTTPMethod: "POST",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}/termination/credentials?operation=delete",
	}

	if input == nil {
		input = &DeleteVoiceConnectorTerminationCredentialsInput{}
	}

	req := c.newRequest(op, input, &DeleteVoiceConnectorTerminationCredentialsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteVoiceConnectorTerminationCredentialsRequest{Request: req, Input: input, Copy: c.DeleteVoiceConnectorTerminationCredentialsRequest}
}

// DeleteVoiceConnectorTerminationCredentialsRequest is the request type for the
// DeleteVoiceConnectorTerminationCredentials API operation.
type DeleteVoiceConnectorTerminationCredentialsRequest struct {
	*aws.Request
	Input *DeleteVoiceConnectorTerminationCredentialsInput
	Copy  func(*DeleteVoiceConnectorTerminationCredentialsInput) DeleteVoiceConnectorTerminationCredentialsRequest
}

// Send marshals and sends the DeleteVoiceConnectorTerminationCredentials API request.
func (r DeleteVoiceConnectorTerminationCredentialsRequest) Send(ctx context.Context) (*DeleteVoiceConnectorTerminationCredentialsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVoiceConnectorTerminationCredentialsResponse{
		DeleteVoiceConnectorTerminationCredentialsOutput: r.Request.Data.(*DeleteVoiceConnectorTerminationCredentialsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVoiceConnectorTerminationCredentialsResponse is the response type for the
// DeleteVoiceConnectorTerminationCredentials API operation.
type DeleteVoiceConnectorTerminationCredentialsResponse struct {
	*DeleteVoiceConnectorTerminationCredentialsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVoiceConnectorTerminationCredentials request.
func (r *DeleteVoiceConnectorTerminationCredentialsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
