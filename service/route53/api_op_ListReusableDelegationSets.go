// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// A request to get a list of the reusable delegation sets that are associated
// with the current AWS account.
type ListReusableDelegationSetsInput struct {
	_ struct{} `type:"structure"`

	// If the value of IsTruncated in the previous response was true, you have more
	// reusable delegation sets. To get another group, submit another ListReusableDelegationSets
	// request.
	//
	// For the value of marker, specify the value of NextMarker from the previous
	// response, which is the ID of the first reusable delegation set that Amazon
	// Route 53 will return if you submit another request.
	//
	// If the value of IsTruncated in the previous response was false, there are
	// no more reusable delegation sets to get.
	Marker *string `location:"querystring" locationName:"marker" type:"string"`

	// The number of reusable delegation sets that you want Amazon Route 53 to return
	// in the response to this request. If you specify a value greater than 100,
	// Route 53 returns only the first 100 reusable delegation sets.
	MaxItems *string `location:"querystring" locationName:"maxitems" type:"string"`
}

// String returns the string representation
func (s ListReusableDelegationSetsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListReusableDelegationSetsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "marker", protocol.StringValue(v), metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxitems", protocol.StringValue(v), metadata)
	}
	return nil
}

// A complex type that contains information about the reusable delegation sets
// that are associated with the current AWS account.
type ListReusableDelegationSetsOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains one DelegationSet element for each reusable
	// delegation set that was created by the current AWS account.
	//
	// DelegationSets is a required field
	DelegationSets []DelegationSet `locationNameList:"DelegationSet" type:"list" required:"true"`

	// A flag that indicates whether there are more reusable delegation sets to
	// be listed.
	//
	// IsTruncated is a required field
	IsTruncated *bool `type:"boolean" required:"true"`

	// For the second and subsequent calls to ListReusableDelegationSets, Marker
	// is the value that you specified for the marker parameter in the request that
	// produced the current response.
	//
	// Marker is a required field
	Marker *string `type:"string" required:"true"`

	// The value that you specified for the maxitems parameter in the call to ListReusableDelegationSets
	// that produced the current response.
	//
	// MaxItems is a required field
	MaxItems *string `type:"string" required:"true"`

	// If IsTruncated is true, the value of NextMarker identifies the next reusable
	// delegation set that Amazon Route 53 will return if you submit another ListReusableDelegationSets
	// request and specify the value of NextMarker in the marker parameter.
	NextMarker *string `type:"string"`
}

// String returns the string representation
func (s ListReusableDelegationSetsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListReusableDelegationSetsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DelegationSets != nil {
		v := s.DelegationSets

		metadata := protocol.Metadata{ListLocationName: "DelegationSet"}
		ls0 := e.List(protocol.BodyTarget, "DelegationSets", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.IsTruncated != nil {
		v := *s.IsTruncated

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IsTruncated", protocol.BoolValue(v), metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Marker", protocol.StringValue(v), metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxItems", protocol.StringValue(v), metadata)
	}
	if s.NextMarker != nil {
		v := *s.NextMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextMarker", protocol.StringValue(v), metadata)
	}
	return nil
}

const opListReusableDelegationSets = "ListReusableDelegationSets"

// ListReusableDelegationSetsRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Retrieves a list of the reusable delegation sets that are associated with
// the current AWS account.
//
//    // Example sending a request using ListReusableDelegationSetsRequest.
//    req := client.ListReusableDelegationSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/ListReusableDelegationSets
func (c *Client) ListReusableDelegationSetsRequest(input *ListReusableDelegationSetsInput) ListReusableDelegationSetsRequest {
	op := &aws.Operation{
		Name:       opListReusableDelegationSets,
		HTTPMethod: "GET",
		HTTPPath:   "/2013-04-01/delegationset",
	}

	if input == nil {
		input = &ListReusableDelegationSetsInput{}
	}

	req := c.newRequest(op, input, &ListReusableDelegationSetsOutput{})

	return ListReusableDelegationSetsRequest{Request: req, Input: input, Copy: c.ListReusableDelegationSetsRequest}
}

// ListReusableDelegationSetsRequest is the request type for the
// ListReusableDelegationSets API operation.
type ListReusableDelegationSetsRequest struct {
	*aws.Request
	Input *ListReusableDelegationSetsInput
	Copy  func(*ListReusableDelegationSetsInput) ListReusableDelegationSetsRequest
}

// Send marshals and sends the ListReusableDelegationSets API request.
func (r ListReusableDelegationSetsRequest) Send(ctx context.Context) (*ListReusableDelegationSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListReusableDelegationSetsResponse{
		ListReusableDelegationSetsOutput: r.Request.Data.(*ListReusableDelegationSetsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListReusableDelegationSetsResponse is the response type for the
// ListReusableDelegationSets API operation.
type ListReusableDelegationSetsResponse struct {
	*ListReusableDelegationSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListReusableDelegationSets request.
func (r *ListReusableDelegationSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
