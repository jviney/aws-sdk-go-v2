// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackage

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type RotateIngestEndpointCredentialsInput struct {
	_ struct{} `type:"structure"`

	// Id is a required field
	Id *string `location:"uri" locationName:"id" type:"string" required:"true"`

	// IngestEndpointId is a required field
	IngestEndpointId *string `location:"uri" locationName:"ingest_endpoint_id" type:"string" required:"true"`
}

// String returns the string representation
func (s RotateIngestEndpointCredentialsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RotateIngestEndpointCredentialsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RotateIngestEndpointCredentialsInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if s.IngestEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IngestEndpointId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RotateIngestEndpointCredentialsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IngestEndpointId != nil {
		v := *s.IngestEndpointId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ingest_endpoint_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type RotateIngestEndpointCredentialsOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `locationName:"arn" type:"string"`

	Description *string `locationName:"description" type:"string"`

	// An HTTP Live Streaming (HLS) ingest resource configuration.
	HlsIngest *HlsIngest `locationName:"hlsIngest" type:"structure"`

	Id *string `locationName:"id" type:"string"`

	// A collection of tags associated with a resource
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s RotateIngestEndpointCredentialsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RotateIngestEndpointCredentialsOutput) MarshalFields(e protocol.FieldEncoder) error {
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

const opRotateIngestEndpointCredentials = "RotateIngestEndpointCredentials"

// RotateIngestEndpointCredentialsRequest returns a request value for making API operation for
// AWS Elemental MediaPackage.
//
// Rotate the IngestEndpoint's username and password, as specified by the IngestEndpoint's
// id.
//
//    // Example sending a request using RotateIngestEndpointCredentialsRequest.
//    req := client.RotateIngestEndpointCredentialsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-2017-10-12/RotateIngestEndpointCredentials
func (c *Client) RotateIngestEndpointCredentialsRequest(input *RotateIngestEndpointCredentialsInput) RotateIngestEndpointCredentialsRequest {
	op := &aws.Operation{
		Name:       opRotateIngestEndpointCredentials,
		HTTPMethod: "PUT",
		HTTPPath:   "/channels/{id}/ingest_endpoints/{ingest_endpoint_id}/credentials",
	}

	if input == nil {
		input = &RotateIngestEndpointCredentialsInput{}
	}

	req := c.newRequest(op, input, &RotateIngestEndpointCredentialsOutput{})

	return RotateIngestEndpointCredentialsRequest{Request: req, Input: input, Copy: c.RotateIngestEndpointCredentialsRequest}
}

// RotateIngestEndpointCredentialsRequest is the request type for the
// RotateIngestEndpointCredentials API operation.
type RotateIngestEndpointCredentialsRequest struct {
	*aws.Request
	Input *RotateIngestEndpointCredentialsInput
	Copy  func(*RotateIngestEndpointCredentialsInput) RotateIngestEndpointCredentialsRequest
}

// Send marshals and sends the RotateIngestEndpointCredentials API request.
func (r RotateIngestEndpointCredentialsRequest) Send(ctx context.Context) (*RotateIngestEndpointCredentialsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RotateIngestEndpointCredentialsResponse{
		RotateIngestEndpointCredentialsOutput: r.Request.Data.(*RotateIngestEndpointCredentialsOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RotateIngestEndpointCredentialsResponse is the response type for the
// RotateIngestEndpointCredentials API operation.
type RotateIngestEndpointCredentialsResponse struct {
	*RotateIngestEndpointCredentialsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RotateIngestEndpointCredentials request.
func (r *RotateIngestEndpointCredentialsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
