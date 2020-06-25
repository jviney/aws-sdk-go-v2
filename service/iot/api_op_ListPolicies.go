// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// The input for the ListPolicies operation.
type ListPoliciesInput struct {
	_ struct{} `type:"structure"`

	// Specifies the order for results. If true, the results are returned in ascending
	// creation order.
	AscendingOrder *bool `location:"querystring" locationName:"isAscendingOrder" type:"boolean"`

	// The marker for the next set of results.
	Marker *string `location:"querystring" locationName:"marker" type:"string"`

	// The result page size.
	PageSize *int64 `location:"querystring" locationName:"pageSize" min:"1" type:"integer"`
}

// String returns the string representation
func (s ListPoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPoliciesInput"}
	if s.PageSize != nil && *s.PageSize < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("PageSize", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPoliciesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AscendingOrder != nil {
		v := *s.AscendingOrder

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "isAscendingOrder", protocol.BoolValue(v), metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PageSize != nil {
		v := *s.PageSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "pageSize", protocol.Int64Value(v), metadata)
	}
	return nil
}

// The output from the ListPolicies operation.
type ListPoliciesOutput struct {
	_ struct{} `type:"structure"`

	// The marker for the next set of results, or null if there are no additional
	// results.
	NextMarker *string `locationName:"nextMarker" type:"string"`

	// The descriptions of the policies.
	Policies []Policy `locationName:"policies" type:"list"`
}

// String returns the string representation
func (s ListPoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPoliciesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextMarker != nil {
		v := *s.NextMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextMarker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Policies != nil {
		v := s.Policies

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "policies", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListPolicies = "ListPolicies"

// ListPoliciesRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists your policies.
//
//    // Example sending a request using ListPoliciesRequest.
//    req := client.ListPoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListPoliciesRequest(input *ListPoliciesInput) ListPoliciesRequest {
	op := &aws.Operation{
		Name:       opListPolicies,
		HTTPMethod: "GET",
		HTTPPath:   "/policies",
	}

	if input == nil {
		input = &ListPoliciesInput{}
	}

	req := c.newRequest(op, input, &ListPoliciesOutput{})

	return ListPoliciesRequest{Request: req, Input: input, Copy: c.ListPoliciesRequest}
}

// ListPoliciesRequest is the request type for the
// ListPolicies API operation.
type ListPoliciesRequest struct {
	*aws.Request
	Input *ListPoliciesInput
	Copy  func(*ListPoliciesInput) ListPoliciesRequest
}

// Send marshals and sends the ListPolicies API request.
func (r ListPoliciesRequest) Send(ctx context.Context) (*ListPoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPoliciesResponse{
		ListPoliciesOutput: r.Request.Data.(*ListPoliciesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListPoliciesResponse is the response type for the
// ListPolicies API operation.
type ListPoliciesResponse struct {
	*ListPoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPolicies request.
func (r *ListPoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
