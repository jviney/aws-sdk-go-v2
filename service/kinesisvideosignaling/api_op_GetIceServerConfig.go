// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideosignaling

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetIceServerConfigInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the signaling channel to be used for the peer-to-peer connection
	// between configured peers.
	//
	// ChannelARN is a required field
	ChannelARN *string `min:"1" type:"string" required:"true"`

	// Unique identifier for the viewer. Must be unique within the signaling channel.
	ClientId *string `min:"1" type:"string"`

	// Specifies the desired service. Currently, TURN is the only valid value.
	Service Service `type:"string" enum:"true"`

	// An optional user ID to be associated with the credentials.
	Username *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetIceServerConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetIceServerConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetIceServerConfigInput"}

	if s.ChannelARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChannelARN"))
	}
	if s.ChannelARN != nil && len(*s.ChannelARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChannelARN", 1))
	}
	if s.ClientId != nil && len(*s.ClientId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientId", 1))
	}
	if s.Username != nil && len(*s.Username) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Username", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetIceServerConfigInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ChannelARN != nil {
		v := *s.ChannelARN

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ChannelARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ClientId != nil {
		v := *s.ClientId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ClientId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Service) > 0 {
		v := s.Service

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Service", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Username != nil {
		v := *s.Username

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Username", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetIceServerConfigOutput struct {
	_ struct{} `type:"structure"`

	// The list of ICE server information objects.
	IceServerList []IceServer `type:"list"`
}

// String returns the string representation
func (s GetIceServerConfigOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetIceServerConfigOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.IceServerList != nil {
		v := s.IceServerList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "IceServerList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetIceServerConfig = "GetIceServerConfig"

// GetIceServerConfigRequest returns a request value for making API operation for
// Amazon Kinesis Video Signaling Channels.
//
// Gets the Interactive Connectivity Establishment (ICE) server configuration
// information, including URIs, username, and password which can be used to
// configure the WebRTC connection. The ICE component uses this configuration
// information to setup the WebRTC connection, including authenticating with
// the Traversal Using Relays around NAT (TURN) relay server.
//
// TURN is a protocol that is used to improve the connectivity of peer-to-peer
// applications. By providing a cloud-based relay service, TURN ensures that
// a connection can be established even when one or more peers are incapable
// of a direct peer-to-peer connection. For more information, see A REST API
// For Access To TURN Services (https://tools.ietf.org/html/draft-uberti-rtcweb-turn-rest-00).
//
// You can invoke this API to establish a fallback mechanism in case either
// of the peers is unable to establish a direct peer-to-peer connection over
// a signaling channel. You must specify either a signaling channel ARN or the
// client ID in order to invoke this API.
//
//    // Example sending a request using GetIceServerConfigRequest.
//    req := client.GetIceServerConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-signaling-2019-12-04/GetIceServerConfig
func (c *Client) GetIceServerConfigRequest(input *GetIceServerConfigInput) GetIceServerConfigRequest {
	op := &aws.Operation{
		Name:       opGetIceServerConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/get-ice-server-config",
	}

	if input == nil {
		input = &GetIceServerConfigInput{}
	}

	req := c.newRequest(op, input, &GetIceServerConfigOutput{})

	return GetIceServerConfigRequest{Request: req, Input: input, Copy: c.GetIceServerConfigRequest}
}

// GetIceServerConfigRequest is the request type for the
// GetIceServerConfig API operation.
type GetIceServerConfigRequest struct {
	*aws.Request
	Input *GetIceServerConfigInput
	Copy  func(*GetIceServerConfigInput) GetIceServerConfigRequest
}

// Send marshals and sends the GetIceServerConfig API request.
func (r GetIceServerConfigRequest) Send(ctx context.Context) (*GetIceServerConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetIceServerConfigResponse{
		GetIceServerConfigOutput: r.Request.Data.(*GetIceServerConfigOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetIceServerConfigResponse is the response type for the
// GetIceServerConfig API operation.
type GetIceServerConfigResponse struct {
	*GetIceServerConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetIceServerConfig request.
func (r *GetIceServerConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
