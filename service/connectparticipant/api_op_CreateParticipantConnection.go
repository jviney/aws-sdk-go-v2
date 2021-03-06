// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connectparticipant

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type CreateParticipantConnectionInput struct {
	_ struct{} `type:"structure"`

	// Participant Token as obtained from StartChatContact (https://docs.aws.amazon.com/connect/latest/APIReference/API_StartChatContactResponse.html)
	// API response.
	//
	// ParticipantToken is a required field
	ParticipantToken *string `location:"header" locationName:"X-Amz-Bearer" min:"1" type:"string" required:"true"`

	// Type of connection information required.
	//
	// Type is a required field
	Type []ConnectionType `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s CreateParticipantConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateParticipantConnectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateParticipantConnectionInput"}

	if s.ParticipantToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParticipantToken"))
	}
	if s.ParticipantToken != nil && len(*s.ParticipantToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ParticipantToken", 1))
	}

	if s.Type == nil {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}
	if s.Type != nil && len(s.Type) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Type", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateParticipantConnectionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Type != nil {
		v := s.Type

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Type", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.ParticipantToken != nil {
		v := *s.ParticipantToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Amz-Bearer", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateParticipantConnectionOutput struct {
	_ struct{} `type:"structure"`

	// Creates the participant's connection credentials. The authentication token
	// associated with the participant's connection.
	ConnectionCredentials *ConnectionCredentials `type:"structure"`

	// Creates the participant's websocket connection.
	Websocket *Websocket `type:"structure"`
}

// String returns the string representation
func (s CreateParticipantConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateParticipantConnectionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ConnectionCredentials != nil {
		v := s.ConnectionCredentials

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ConnectionCredentials", v, metadata)
	}
	if s.Websocket != nil {
		v := s.Websocket

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Websocket", v, metadata)
	}
	return nil
}

const opCreateParticipantConnection = "CreateParticipantConnection"

// CreateParticipantConnectionRequest returns a request value for making API operation for
// Amazon Connect Participant Service.
//
// Creates the participant's connection. Note that ParticipantToken is used
// for invoking this API instead of ConnectionToken.
//
// The participant token is valid for the lifetime of the participant – until
// the they are part of a contact.
//
// The response URL for WEBSOCKET Type has a connect expiry timeout of 100s.
// Clients must manually connect to the returned websocket URL and subscribe
// to the desired topic.
//
// For chat, you need to publish the following on the established websocket
// connection:
//
// {"topic":"aws/subscribe","content":{"topics":["aws/chat"]}}
//
// Upon websocket URL expiry, as specified in the response ConnectionExpiry
// parameter, clients need to call this API again to obtain a new websocket
// URL and perform the same steps as before.
//
//    // Example sending a request using CreateParticipantConnectionRequest.
//    req := client.CreateParticipantConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/connectparticipant-2018-09-07/CreateParticipantConnection
func (c *Client) CreateParticipantConnectionRequest(input *CreateParticipantConnectionInput) CreateParticipantConnectionRequest {
	op := &aws.Operation{
		Name:       opCreateParticipantConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/participant/connection",
	}

	if input == nil {
		input = &CreateParticipantConnectionInput{}
	}

	req := c.newRequest(op, input, &CreateParticipantConnectionOutput{})

	return CreateParticipantConnectionRequest{Request: req, Input: input, Copy: c.CreateParticipantConnectionRequest}
}

// CreateParticipantConnectionRequest is the request type for the
// CreateParticipantConnection API operation.
type CreateParticipantConnectionRequest struct {
	*aws.Request
	Input *CreateParticipantConnectionInput
	Copy  func(*CreateParticipantConnectionInput) CreateParticipantConnectionRequest
}

// Send marshals and sends the CreateParticipantConnection API request.
func (r CreateParticipantConnectionRequest) Send(ctx context.Context) (*CreateParticipantConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateParticipantConnectionResponse{
		CreateParticipantConnectionOutput: r.Request.Data.(*CreateParticipantConnectionOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateParticipantConnectionResponse is the response type for the
// CreateParticipantConnection API operation.
type CreateParticipantConnectionResponse struct {
	*CreateParticipantConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateParticipantConnection request.
func (r *CreateParticipantConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
