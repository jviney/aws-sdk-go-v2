// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeFleetsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The filters.
	//
	//    * activity-status - The progress of the EC2 Fleet ( error | pending-fulfillment
	//    | pending-termination | fulfilled).
	//
	//    * excess-capacity-termination-policy - Indicates whether to terminate
	//    running instances if the target capacity is decreased below the current
	//    EC2 Fleet size (true | false).
	//
	//    * fleet-state - The state of the EC2 Fleet (submitted | active | deleted
	//    | failed | deleted-running | deleted-terminating | modifying).
	//
	//    * replace-unhealthy-instances - Indicates whether EC2 Fleet should replace
	//    unhealthy instances (true | false).
	//
	//    * type - The type of request (instant | request | maintain).
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The ID of the EC2 Fleets.
	FleetIds []string `locationName:"FleetId" type:"list"`

	// The maximum number of results to return in a single call. Specify a value
	// between 1 and 1000. The default value is 1000. To retrieve the remaining
	// results, make another call with the returned NextToken value.
	MaxResults *int64 `type:"integer"`

	// The token for the next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeFleetsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeFleetsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the EC2 Fleets.
	Fleets []FleetData `locationName:"fleetSet" locationNameList:"item" type:"list"`

	// The token for the next set of results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeFleetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeFleets = "DescribeFleets"

// DescribeFleetsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified EC2 Fleets or all of your EC2 Fleets.
//
//    // Example sending a request using DescribeFleetsRequest.
//    req := client.DescribeFleetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeFleets
func (c *Client) DescribeFleetsRequest(input *DescribeFleetsInput) DescribeFleetsRequest {
	op := &aws.Operation{
		Name:       opDescribeFleets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeFleetsInput{}
	}

	req := c.newRequest(op, input, &DescribeFleetsOutput{})

	return DescribeFleetsRequest{Request: req, Input: input, Copy: c.DescribeFleetsRequest}
}

// DescribeFleetsRequest is the request type for the
// DescribeFleets API operation.
type DescribeFleetsRequest struct {
	*aws.Request
	Input *DescribeFleetsInput
	Copy  func(*DescribeFleetsInput) DescribeFleetsRequest
}

// Send marshals and sends the DescribeFleets API request.
func (r DescribeFleetsRequest) Send(ctx context.Context) (*DescribeFleetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFleetsResponse{
		DescribeFleetsOutput: r.Request.Data.(*DescribeFleetsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeFleetsRequestPaginator returns a paginator for DescribeFleets.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeFleetsRequest(input)
//   p := ec2.NewDescribeFleetsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeFleetsPaginator(req DescribeFleetsRequest) DescribeFleetsPaginator {
	return DescribeFleetsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeFleetsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeFleetsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeFleetsPaginator struct {
	aws.Pager
}

func (p *DescribeFleetsPaginator) CurrentPage() *DescribeFleetsOutput {
	return p.Pager.CurrentPage().(*DescribeFleetsOutput)
}

// DescribeFleetsResponse is the response type for the
// DescribeFleets API operation.
type DescribeFleetsResponse struct {
	*DescribeFleetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFleets request.
func (r *DescribeFleetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
