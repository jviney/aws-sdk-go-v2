// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListActiveViolationsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return at one time.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The name of the Device Defender security profile for which violations are
	// listed.
	SecurityProfileName *string `location:"querystring" locationName:"securityProfileName" min:"1" type:"string"`

	// The name of the thing whose active violations are listed.
	ThingName *string `location:"querystring" locationName:"thingName" min:"1" type:"string"`
}

// String returns the string representation
func (s ListActiveViolationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListActiveViolationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListActiveViolationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.SecurityProfileName != nil && len(*s.SecurityProfileName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecurityProfileName", 1))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListActiveViolationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SecurityProfileName != nil {
		v := *s.SecurityProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "securityProfileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "thingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListActiveViolationsOutput struct {
	_ struct{} `type:"structure"`

	// The list of active violations.
	ActiveViolations []ActiveViolation `locationName:"activeViolations" type:"list"`

	// A token that can be used to retrieve the next set of results, or null if
	// there are no additional results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListActiveViolationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListActiveViolationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ActiveViolations != nil {
		v := s.ActiveViolations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "activeViolations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListActiveViolations = "ListActiveViolations"

// ListActiveViolationsRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists the active violations for a given Device Defender security profile.
//
//    // Example sending a request using ListActiveViolationsRequest.
//    req := client.ListActiveViolationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListActiveViolationsRequest(input *ListActiveViolationsInput) ListActiveViolationsRequest {
	op := &aws.Operation{
		Name:       opListActiveViolations,
		HTTPMethod: "GET",
		HTTPPath:   "/active-violations",
	}

	if input == nil {
		input = &ListActiveViolationsInput{}
	}

	req := c.newRequest(op, input, &ListActiveViolationsOutput{})

	return ListActiveViolationsRequest{Request: req, Input: input, Copy: c.ListActiveViolationsRequest}
}

// ListActiveViolationsRequest is the request type for the
// ListActiveViolations API operation.
type ListActiveViolationsRequest struct {
	*aws.Request
	Input *ListActiveViolationsInput
	Copy  func(*ListActiveViolationsInput) ListActiveViolationsRequest
}

// Send marshals and sends the ListActiveViolations API request.
func (r ListActiveViolationsRequest) Send(ctx context.Context) (*ListActiveViolationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListActiveViolationsResponse{
		ListActiveViolationsOutput: r.Request.Data.(*ListActiveViolationsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListActiveViolationsResponse is the response type for the
// ListActiveViolations API operation.
type ListActiveViolationsResponse struct {
	*ListActiveViolationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListActiveViolations request.
func (r *ListActiveViolationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
