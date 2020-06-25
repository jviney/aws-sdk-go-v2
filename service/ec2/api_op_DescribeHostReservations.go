// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeHostReservationsInput struct {
	_ struct{} `type:"structure"`

	// The filters.
	//
	//    * instance-family - The instance family (for example, m4).
	//
	//    * payment-option - The payment option (NoUpfront | PartialUpfront | AllUpfront).
	//
	//    * state - The state of the reservation (payment-pending | payment-failed
	//    | active | retired).
	//
	//    * tag:<key> - The key/value combination of a tag assigned to the resource.
	//    Use the tag key in the filter name and the tag value as the filter value.
	//    For example, to find all resources that have a tag with the key Owner
	//    and the value TeamA, specify tag:Owner for the filter name and TeamA for
	//    the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. Use this filter
	//    to find all resources assigned a tag with a specific key, regardless of
	//    the tag value.
	Filter []Filter `locationNameList:"Filter" type:"list"`

	// The host reservation IDs.
	HostReservationIdSet []string `locationNameList:"item" type:"list"`

	// The maximum number of results to return for the request in a single page.
	// The remaining results can be seen by sending another request with the returned
	// nextToken value. This value can be between 5 and 500. If maxResults is given
	// a larger value than 500, you receive an error.
	MaxResults *int64 `type:"integer"`

	// The token to use to retrieve the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeHostReservationsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeHostReservationsOutput struct {
	_ struct{} `type:"structure"`

	// Details about the reservation's configuration.
	HostReservationSet []HostReservation `locationName:"hostReservationSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeHostReservationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeHostReservations = "DescribeHostReservations"

// DescribeHostReservationsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes reservations that are associated with Dedicated Hosts in your account.
//
//    // Example sending a request using DescribeHostReservationsRequest.
//    req := client.DescribeHostReservationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeHostReservations
func (c *Client) DescribeHostReservationsRequest(input *DescribeHostReservationsInput) DescribeHostReservationsRequest {
	op := &aws.Operation{
		Name:       opDescribeHostReservations,
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
		input = &DescribeHostReservationsInput{}
	}

	req := c.newRequest(op, input, &DescribeHostReservationsOutput{})

	return DescribeHostReservationsRequest{Request: req, Input: input, Copy: c.DescribeHostReservationsRequest}
}

// DescribeHostReservationsRequest is the request type for the
// DescribeHostReservations API operation.
type DescribeHostReservationsRequest struct {
	*aws.Request
	Input *DescribeHostReservationsInput
	Copy  func(*DescribeHostReservationsInput) DescribeHostReservationsRequest
}

// Send marshals and sends the DescribeHostReservations API request.
func (r DescribeHostReservationsRequest) Send(ctx context.Context) (*DescribeHostReservationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeHostReservationsResponse{
		DescribeHostReservationsOutput: r.Request.Data.(*DescribeHostReservationsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeHostReservationsRequestPaginator returns a paginator for DescribeHostReservations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeHostReservationsRequest(input)
//   p := ec2.NewDescribeHostReservationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeHostReservationsPaginator(req DescribeHostReservationsRequest) DescribeHostReservationsPaginator {
	return DescribeHostReservationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeHostReservationsInput
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

// DescribeHostReservationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeHostReservationsPaginator struct {
	aws.Pager
}

func (p *DescribeHostReservationsPaginator) CurrentPage() *DescribeHostReservationsOutput {
	return p.Pager.CurrentPage().(*DescribeHostReservationsOutput)
}

// DescribeHostReservationsResponse is the response type for the
// DescribeHostReservations API operation.
type DescribeHostReservationsResponse struct {
	*DescribeHostReservationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeHostReservations request.
func (r *DescribeHostReservationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
