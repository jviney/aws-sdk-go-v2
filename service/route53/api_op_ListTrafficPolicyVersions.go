// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// A complex type that contains the information about the request to list your
// traffic policies.
type ListTrafficPolicyVersionsInput struct {
	_ struct{} `type:"structure"`

	// Specify the value of Id of the traffic policy for which you want to list
	// all versions.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" min:"1" type:"string" required:"true"`

	// The maximum number of traffic policy versions that you want Amazon Route
	// 53 to include in the response body for this request. If the specified traffic
	// policy has more than MaxItems versions, the value of IsTruncated in the response
	// is true, and the value of the TrafficPolicyVersionMarker element is the ID
	// of the first version that Route 53 will return if you submit another request.
	MaxItems *string `location:"querystring" locationName:"maxitems" type:"string"`

	// For your first request to ListTrafficPolicyVersions, don't include the TrafficPolicyVersionMarker
	// parameter.
	//
	// If you have more traffic policy versions than the value of MaxItems, ListTrafficPolicyVersions
	// returns only the first group of MaxItems versions. To get more traffic policy
	// versions, submit another ListTrafficPolicyVersions request. For the value
	// of TrafficPolicyVersionMarker, specify the value of TrafficPolicyVersionMarker
	// in the previous response.
	TrafficPolicyVersionMarker *string `location:"querystring" locationName:"trafficpolicyversion" type:"string"`
}

// String returns the string representation
func (s ListTrafficPolicyVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTrafficPolicyVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTrafficPolicyVersionsInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTrafficPolicyVersionsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxitems", protocol.StringValue(v), metadata)
	}
	if s.TrafficPolicyVersionMarker != nil {
		v := *s.TrafficPolicyVersionMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "trafficpolicyversion", protocol.StringValue(v), metadata)
	}
	return nil
}

// A complex type that contains the response information for the request.
type ListTrafficPolicyVersionsOutput struct {
	_ struct{} `type:"structure"`

	// A flag that indicates whether there are more traffic policies to be listed.
	// If the response was truncated, you can get the next group of traffic policies
	// by submitting another ListTrafficPolicyVersions request and specifying the
	// value of NextMarker in the marker parameter.
	//
	// IsTruncated is a required field
	IsTruncated *bool `type:"boolean" required:"true"`

	// The value that you specified for the maxitems parameter in the ListTrafficPolicyVersions
	// request that produced the current response.
	//
	// MaxItems is a required field
	MaxItems *string `type:"string" required:"true"`

	// A list that contains one TrafficPolicy element for each traffic policy version
	// that is associated with the specified traffic policy.
	//
	// TrafficPolicies is a required field
	TrafficPolicies []TrafficPolicy `locationNameList:"TrafficPolicy" type:"list" required:"true"`

	// If IsTruncated is true, the value of TrafficPolicyVersionMarker identifies
	// the first traffic policy that Amazon Route 53 will return if you submit another
	// request. Call ListTrafficPolicyVersions again and specify the value of TrafficPolicyVersionMarker
	// in the TrafficPolicyVersionMarker request parameter.
	//
	// This element is present only if IsTruncated is true.
	//
	// TrafficPolicyVersionMarker is a required field
	TrafficPolicyVersionMarker *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListTrafficPolicyVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTrafficPolicyVersionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.IsTruncated != nil {
		v := *s.IsTruncated

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IsTruncated", protocol.BoolValue(v), metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxItems", protocol.StringValue(v), metadata)
	}
	if s.TrafficPolicies != nil {
		v := s.TrafficPolicies

		metadata := protocol.Metadata{ListLocationName: "TrafficPolicy"}
		ls0 := e.List(protocol.BodyTarget, "TrafficPolicies", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.TrafficPolicyVersionMarker != nil {
		v := *s.TrafficPolicyVersionMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TrafficPolicyVersionMarker", protocol.StringValue(v), metadata)
	}
	return nil
}

const opListTrafficPolicyVersions = "ListTrafficPolicyVersions"

// ListTrafficPolicyVersionsRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Gets information about all of the versions for a specified traffic policy.
//
// Traffic policy versions are listed in numerical order by VersionNumber.
//
//    // Example sending a request using ListTrafficPolicyVersionsRequest.
//    req := client.ListTrafficPolicyVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/ListTrafficPolicyVersions
func (c *Client) ListTrafficPolicyVersionsRequest(input *ListTrafficPolicyVersionsInput) ListTrafficPolicyVersionsRequest {
	op := &aws.Operation{
		Name:       opListTrafficPolicyVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/2013-04-01/trafficpolicies/{Id}/versions",
	}

	if input == nil {
		input = &ListTrafficPolicyVersionsInput{}
	}

	req := c.newRequest(op, input, &ListTrafficPolicyVersionsOutput{})

	return ListTrafficPolicyVersionsRequest{Request: req, Input: input, Copy: c.ListTrafficPolicyVersionsRequest}
}

// ListTrafficPolicyVersionsRequest is the request type for the
// ListTrafficPolicyVersions API operation.
type ListTrafficPolicyVersionsRequest struct {
	*aws.Request
	Input *ListTrafficPolicyVersionsInput
	Copy  func(*ListTrafficPolicyVersionsInput) ListTrafficPolicyVersionsRequest
}

// Send marshals and sends the ListTrafficPolicyVersions API request.
func (r ListTrafficPolicyVersionsRequest) Send(ctx context.Context) (*ListTrafficPolicyVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTrafficPolicyVersionsResponse{
		ListTrafficPolicyVersionsOutput: r.Request.Data.(*ListTrafficPolicyVersionsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTrafficPolicyVersionsResponse is the response type for the
// ListTrafficPolicyVersions API operation.
type ListTrafficPolicyVersionsResponse struct {
	*ListTrafficPolicyVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTrafficPolicyVersions request.
func (r *ListTrafficPolicyVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
