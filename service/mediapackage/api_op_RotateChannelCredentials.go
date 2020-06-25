// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackage

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type RotateChannelCredentialsInput struct {
	_ struct{} `deprecated:"true" type:"structure"`

	// Id is a required field
	Id *string `location:"uri" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s RotateChannelCredentialsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RotateChannelCredentialsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RotateChannelCredentialsInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RotateChannelCredentialsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type RotateChannelCredentialsOutput struct {
	_ struct{} `deprecated:"true" type:"structure"`

	Arn *string `locationName:"arn" type:"string"`

	Description *string `locationName:"description" type:"string"`

	// An HTTP Live Streaming (HLS) ingest resource configuration.
	HlsIngest *HlsIngest `locationName:"hlsIngest" type:"structure"`

	Id *string `locationName:"id" type:"string"`

	// A collection of tags associated with a resource
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s RotateChannelCredentialsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RotateChannelCredentialsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.HlsIngest != nil {
		v := s.HlsIngest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "hlsIngest", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opRotateChannelCredentials = "RotateChannelCredentials"

// RotateChannelCredentialsRequest returns a request value for making API operation for
// AWS Elemental MediaPackage.
//
// Changes the Channel's first IngestEndpoint's username and password. WARNING
// - This API is deprecated. Please use RotateIngestEndpointCredentials instead
//
//    // Example sending a request using RotateChannelCredentialsRequest.
//    req := client.RotateChannelCredentialsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-2017-10-12/RotateChannelCredentials
func (c *Client) RotateChannelCredentialsRequest(input *RotateChannelCredentialsInput) RotateChannelCredentialsRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, RotateChannelCredentials, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opRotateChannelCredentials,
		HTTPMethod: "PUT",
		HTTPPath:   "/channels/{id}/credentials",
	}

	if input == nil {
		input = &RotateChannelCredentialsInput{}
	}

	req := c.newRequest(op, input, &RotateChannelCredentialsOutput{})

	return RotateChannelCredentialsRequest{Request: req, Input: input, Copy: c.RotateChannelCredentialsRequest}
}

// RotateChannelCredentialsRequest is the request type for the
// RotateChannelCredentials API operation.
type RotateChannelCredentialsRequest struct {
	*aws.Request
	Input *RotateChannelCredentialsInput
	Copy  func(*RotateChannelCredentialsInput) RotateChannelCredentialsRequest
}

// Send marshals and sends the RotateChannelCredentials API request.
func (r RotateChannelCredentialsRequest) Send(ctx context.Context) (*RotateChannelCredentialsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RotateChannelCredentialsResponse{
		RotateChannelCredentialsOutput: r.Request.Data.(*RotateChannelCredentialsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RotateChannelCredentialsResponse is the response type for the
// RotateChannelCredentials API operation.
type RotateChannelCredentialsResponse struct {
	*RotateChannelCredentialsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RotateChannelCredentials request.
func (r *RotateChannelCredentialsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
