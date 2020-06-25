// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewaymanagementapi

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetConnectionInput struct {
	_ struct{} `type:"structure"`

	// ConnectionId is a required field
	ConnectionId *string `location:"uri" locationName:"connectionId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetConnectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetConnectionInput"}

	if s.ConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetConnectionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ConnectionId != nil {
		v := *s.ConnectionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "connectionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetConnectionOutput struct {
	_ struct{} `type:"structure"`

	ConnectedAt *time.Time `locationName:"connectedAt" type:"timestamp" timestampFormat:"iso8601"`

	Identity *Identity `locationName:"identity" type:"structure"`

	LastActiveAt *time.Time `locationName:"lastActiveAt" type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s GetConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetConnectionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ConnectedAt != nil {
		v := *s.ConnectedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "connectedAt",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.Identity != nil {
		v := s.Identity

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "identity", v, metadata)
	}
	if s.LastActiveAt != nil {
		v := *s.LastActiveAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastActiveAt",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	return nil
}

const opGetConnection = "GetConnection"

// GetConnectionRequest returns a request value for making API operation for
// AmazonApiGatewayManagementApi.
//
// Get information about the connection with the provided id.
//
//    // Example sending a request using GetConnectionRequest.
//    req := client.GetConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewaymanagementapi-2018-11-29/GetConnection
func (c *Client) GetConnectionRequest(input *GetConnectionInput) GetConnectionRequest {
	op := &aws.Operation{
		Name:       opGetConnection,
		HTTPMethod: "GET",
		HTTPPath:   "/@connections/{connectionId}",
	}

	if input == nil {
		input = &GetConnectionInput{}
	}

	req := c.newRequest(op, input, &GetConnectionOutput{})

	return GetConnectionRequest{Request: req, Input: input, Copy: c.GetConnectionRequest}
}

// GetConnectionRequest is the request type for the
// GetConnection API operation.
type GetConnectionRequest struct {
	*aws.Request
	Input *GetConnectionInput
	Copy  func(*GetConnectionInput) GetConnectionRequest
}

// Send marshals and sends the GetConnection API request.
func (r GetConnectionRequest) Send(ctx context.Context) (*GetConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetConnectionResponse{
		GetConnectionOutput: r.Request.Data.(*GetConnectionOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetConnectionResponse is the response type for the
// GetConnection API operation.
type GetConnectionResponse struct {
	*GetConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetConnection request.
func (r *GetConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
