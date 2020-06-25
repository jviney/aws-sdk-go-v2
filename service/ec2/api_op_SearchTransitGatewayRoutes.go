// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type SearchTransitGatewayRoutesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters. The possible values are:
	//
	//    * attachment.transit-gateway-attachment-id- The id of the transit gateway
	//    attachment.
	//
	//    * attachment.resource-id - The resource id of the transit gateway attachment.
	//
	//    * attachment.resource-type - The attachment resource type (vpc | vpn).
	//
	//    * route-search.exact-match - The exact match of the specified filter.
	//
	//    * route-search.longest-prefix-match - The longest prefix that matches
	//    the route.
	//
	//    * route-search.subnet-of-match - The routes with a subnet that match the
	//    specified CIDR filter.
	//
	//    * route-search.supernet-of-match - The routes with a CIDR that encompass
	//    the CIDR filter. For example, if you have 10.0.1.0/29 and 10.0.1.0/31
	//    routes in your route table and you specify supernet-of-match as 10.0.1.0/30,
	//    then the result returns 10.0.1.0/29.
	//
	//    * state - The state of the route (active | blackhole).
	//
	//    * type - The type of route (propagated | static).
	//
	// Filters is a required field
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list" required:"true"`

	// The maximum number of routes to return.
	MaxResults *int64 `min:"5" type:"integer"`

	// The ID of the transit gateway route table.
	//
	// TransitGatewayRouteTableId is a required field
	TransitGatewayRouteTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SearchTransitGatewayRoutesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SearchTransitGatewayRoutesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SearchTransitGatewayRoutesInput"}

	if s.Filters == nil {
		invalidParams.Add(aws.NewErrParamRequired("Filters"))
	}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if s.TransitGatewayRouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayRouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SearchTransitGatewayRoutesOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether there are additional routes available.
	AdditionalRoutesAvailable *bool `locationName:"additionalRoutesAvailable" type:"boolean"`

	// Information about the routes.
	Routes []TransitGatewayRoute `locationName:"routeSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s SearchTransitGatewayRoutesOutput) String() string {
	return awsutil.Prettify(s)
}

const opSearchTransitGatewayRoutes = "SearchTransitGatewayRoutes"

// SearchTransitGatewayRoutesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Searches for routes in the specified transit gateway route table.
//
//    // Example sending a request using SearchTransitGatewayRoutesRequest.
//    req := client.SearchTransitGatewayRoutesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/SearchTransitGatewayRoutes
func (c *Client) SearchTransitGatewayRoutesRequest(input *SearchTransitGatewayRoutesInput) SearchTransitGatewayRoutesRequest {
	op := &aws.Operation{
		Name:       opSearchTransitGatewayRoutes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SearchTransitGatewayRoutesInput{}
	}

	req := c.newRequest(op, input, &SearchTransitGatewayRoutesOutput{})

	return SearchTransitGatewayRoutesRequest{Request: req, Input: input, Copy: c.SearchTransitGatewayRoutesRequest}
}

// SearchTransitGatewayRoutesRequest is the request type for the
// SearchTransitGatewayRoutes API operation.
type SearchTransitGatewayRoutesRequest struct {
	*aws.Request
	Input *SearchTransitGatewayRoutesInput
	Copy  func(*SearchTransitGatewayRoutesInput) SearchTransitGatewayRoutesRequest
}

// Send marshals and sends the SearchTransitGatewayRoutes API request.
func (r SearchTransitGatewayRoutesRequest) Send(ctx context.Context) (*SearchTransitGatewayRoutesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SearchTransitGatewayRoutesResponse{
		SearchTransitGatewayRoutesOutput: r.Request.Data.(*SearchTransitGatewayRoutesOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SearchTransitGatewayRoutesResponse is the response type for the
// SearchTransitGatewayRoutes API operation.
type SearchTransitGatewayRoutesResponse struct {
	*SearchTransitGatewayRoutesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SearchTransitGatewayRoutes request.
func (r *SearchTransitGatewayRoutesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
