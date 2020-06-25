// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeEventSubscriptionsInput struct {
	_ struct{} `type:"structure"`

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeEventSubscriptions request
	// exceed the value specified in MaxRecords, AWS returns a value in the Marker
	// field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the Marker parameter and retrying
	// the request.
	Marker *string `type:"string"`

	// The maximum number of response records to return in each call. If the number
	// of remaining response records exceeds the specified MaxRecords value, a value
	// is returned in a marker field of the response. You can retrieve the next
	// set of records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The name of the Amazon Redshift event notification subscription to be described.
	SubscriptionName *string `type:"string"`

	// A tag key or keys for which you want to return all matching event notification
	// subscriptions that are associated with the specified key or keys. For example,
	// suppose that you have subscriptions that are tagged with keys called owner
	// and environment. If you specify both of these tag keys in the request, Amazon
	// Redshift returns a response with the subscriptions that have either or both
	// of these tag keys associated with them.
	TagKeys []string `locationNameList:"TagKey" type:"list"`

	// A tag value or values for which you want to return all matching event notification
	// subscriptions that are associated with the specified tag value or values.
	// For example, suppose that you have subscriptions that are tagged with values
	// called admin and test. If you specify both of these tag values in the request,
	// Amazon Redshift returns a response with the subscriptions that have either
	// or both of these tag values associated with them.
	TagValues []string `locationNameList:"TagValue" type:"list"`
}

// String returns the string representation
func (s DescribeEventSubscriptionsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeEventSubscriptionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of event subscriptions.
	EventSubscriptionsList []EventSubscription `locationNameList:"EventSubscription" type:"list"`

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeEventSubscriptionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEventSubscriptions = "DescribeEventSubscriptions"

// DescribeEventSubscriptionsRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Lists descriptions of all the Amazon Redshift event notification subscriptions
// for a customer account. If you specify a subscription name, lists the description
// for that subscription.
//
// If you specify both tag keys and tag values in the same request, Amazon Redshift
// returns all event notification subscriptions that match any combination of
// the specified keys and values. For example, if you have owner and environment
// for tag keys, and admin and test for tag values, all subscriptions that have
// any combination of those values are returned.
//
// If both tag keys and values are omitted from the request, subscriptions are
// returned regardless of whether they have tag keys or values associated with
// them.
//
//    // Example sending a request using DescribeEventSubscriptionsRequest.
//    req := client.DescribeEventSubscriptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeEventSubscriptions
func (c *Client) DescribeEventSubscriptionsRequest(input *DescribeEventSubscriptionsInput) DescribeEventSubscriptionsRequest {
	op := &aws.Operation{
		Name:       opDescribeEventSubscriptions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeEventSubscriptionsInput{}
	}

	req := c.newRequest(op, input, &DescribeEventSubscriptionsOutput{})

	return DescribeEventSubscriptionsRequest{Request: req, Input: input, Copy: c.DescribeEventSubscriptionsRequest}
}

// DescribeEventSubscriptionsRequest is the request type for the
// DescribeEventSubscriptions API operation.
type DescribeEventSubscriptionsRequest struct {
	*aws.Request
	Input *DescribeEventSubscriptionsInput
	Copy  func(*DescribeEventSubscriptionsInput) DescribeEventSubscriptionsRequest
}

// Send marshals and sends the DescribeEventSubscriptions API request.
func (r DescribeEventSubscriptionsRequest) Send(ctx context.Context) (*DescribeEventSubscriptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEventSubscriptionsResponse{
		DescribeEventSubscriptionsOutput: r.Request.Data.(*DescribeEventSubscriptionsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeEventSubscriptionsRequestPaginator returns a paginator for DescribeEventSubscriptions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeEventSubscriptionsRequest(input)
//   p := redshift.NewDescribeEventSubscriptionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeEventSubscriptionsPaginator(req DescribeEventSubscriptionsRequest) DescribeEventSubscriptionsPaginator {
	return DescribeEventSubscriptionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeEventSubscriptionsInput
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

// DescribeEventSubscriptionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeEventSubscriptionsPaginator struct {
	aws.Pager
}

func (p *DescribeEventSubscriptionsPaginator) CurrentPage() *DescribeEventSubscriptionsOutput {
	return p.Pager.CurrentPage().(*DescribeEventSubscriptionsOutput)
}

// DescribeEventSubscriptionsResponse is the response type for the
// DescribeEventSubscriptions API operation.
type DescribeEventSubscriptionsResponse struct {
	*DescribeEventSubscriptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEventSubscriptions request.
func (r *DescribeEventSubscriptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
