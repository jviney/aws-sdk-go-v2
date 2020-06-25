// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointsmsvoice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListConfigurationSetsInput struct {
	_ struct{} `type:"structure"`

	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`

	PageSize *string `location:"querystring" locationName:"PageSize" type:"string"`
}

// String returns the string representation
func (s ListConfigurationSetsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListConfigurationSetsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PageSize != nil {
		v := *s.PageSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "PageSize", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An object that contains information about the configuration sets for your
// account in the current region.
type ListConfigurationSetsOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains a list of configuration sets for your account in
	// the current region.
	ConfigurationSets []string `type:"list"`

	// A token returned from a previous call to ListConfigurationSets to indicate
	// the position in the list of configuration sets.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListConfigurationSetsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListConfigurationSetsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ConfigurationSets != nil {
		v := s.ConfigurationSets

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ConfigurationSets", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListConfigurationSets = "ListConfigurationSets"

// ListConfigurationSetsRequest returns a request value for making API operation for
// Amazon Pinpoint SMS and Voice Service.
//
// List all of the configuration sets associated with your Amazon Pinpoint account
// in the current region.
//
//    // Example sending a request using ListConfigurationSetsRequest.
//    req := client.ListConfigurationSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-sms-voice-2018-09-05/ListConfigurationSets
func (c *Client) ListConfigurationSetsRequest(input *ListConfigurationSetsInput) ListConfigurationSetsRequest {
	op := &aws.Operation{
		Name:       opListConfigurationSets,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/sms-voice/configuration-sets",
	}

	if input == nil {
		input = &ListConfigurationSetsInput{}
	}

	req := c.newRequest(op, input, &ListConfigurationSetsOutput{})

	return ListConfigurationSetsRequest{Request: req, Input: input, Copy: c.ListConfigurationSetsRequest}
}

// ListConfigurationSetsRequest is the request type for the
// ListConfigurationSets API operation.
type ListConfigurationSetsRequest struct {
	*aws.Request
	Input *ListConfigurationSetsInput
	Copy  func(*ListConfigurationSetsInput) ListConfigurationSetsRequest
}

// Send marshals and sends the ListConfigurationSets API request.
func (r ListConfigurationSetsRequest) Send(ctx context.Context) (*ListConfigurationSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListConfigurationSetsResponse{
		ListConfigurationSetsOutput: r.Request.Data.(*ListConfigurationSetsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListConfigurationSetsResponse is the response type for the
// ListConfigurationSets API operation.
type ListConfigurationSetsResponse struct {
	*ListConfigurationSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListConfigurationSets request.
func (r *ListConfigurationSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
