// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediatailor

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListPlaybackConfigurationsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `location:"querystring" locationName:"MaxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s ListPlaybackConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPlaybackConfigurationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPlaybackConfigurationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPlaybackConfigurationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListPlaybackConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	// Array of playback configurations. This might be all the available configurations
	// or a subset, depending on the settings that you provide and the total number
	// of configurations stored.
	Items []PlaybackConfiguration `type:"list"`

	// Pagination token returned by the GET list request when results exceed the
	// maximum allowed. Use the token to fetch the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListPlaybackConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPlaybackConfigurationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Items != nil {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Items", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
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

const opListPlaybackConfigurations = "ListPlaybackConfigurations"

// ListPlaybackConfigurationsRequest returns a request value for making API operation for
// AWS MediaTailor.
//
// Returns a list of the playback configurations defined in AWS Elemental MediaTailor.
// You can specify a maximum number of configurations to return at a time. The
// default maximum is 50. Results are returned in pagefuls. If MediaTailor has
// more configurations than the specified maximum, it provides parameters in
// the response that you can use to retrieve the next pageful.
//
//    // Example sending a request using ListPlaybackConfigurationsRequest.
//    req := client.ListPlaybackConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/ListPlaybackConfigurations
func (c *Client) ListPlaybackConfigurationsRequest(input *ListPlaybackConfigurationsInput) ListPlaybackConfigurationsRequest {
	op := &aws.Operation{
		Name:       opListPlaybackConfigurations,
		HTTPMethod: "GET",
		HTTPPath:   "/playbackConfigurations",
	}

	if input == nil {
		input = &ListPlaybackConfigurationsInput{}
	}

	req := c.newRequest(op, input, &ListPlaybackConfigurationsOutput{})

	return ListPlaybackConfigurationsRequest{Request: req, Input: input, Copy: c.ListPlaybackConfigurationsRequest}
}

// ListPlaybackConfigurationsRequest is the request type for the
// ListPlaybackConfigurations API operation.
type ListPlaybackConfigurationsRequest struct {
	*aws.Request
	Input *ListPlaybackConfigurationsInput
	Copy  func(*ListPlaybackConfigurationsInput) ListPlaybackConfigurationsRequest
}

// Send marshals and sends the ListPlaybackConfigurations API request.
func (r ListPlaybackConfigurationsRequest) Send(ctx context.Context) (*ListPlaybackConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPlaybackConfigurationsResponse{
		ListPlaybackConfigurationsOutput: r.Request.Data.(*ListPlaybackConfigurationsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListPlaybackConfigurationsResponse is the response type for the
// ListPlaybackConfigurations API operation.
type ListPlaybackConfigurationsResponse struct {
	*ListPlaybackConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPlaybackConfigurations request.
func (r *ListPlaybackConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
