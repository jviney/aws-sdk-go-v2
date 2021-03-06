// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Input for ListSubscriptions action.
type ListSubscriptionsInput struct {
	_ struct{} `type:"structure"`

	// Token returned by the previous ListSubscriptions request.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListSubscriptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Response for ListSubscriptions action
type ListSubscriptionsOutput struct {
	_ struct{} `type:"structure"`

	// Token to pass along to the next ListSubscriptions request. This element is
	// returned if there are more subscriptions to retrieve.
	NextToken *string `type:"string"`

	// A list of subscriptions.
	Subscriptions []Subscription `type:"list"`
}

// String returns the string representation
func (s ListSubscriptionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListSubscriptions = "ListSubscriptions"

// ListSubscriptionsRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Returns a list of the requester's subscriptions. Each call returns a limited
// list of subscriptions, up to 100. If there are more subscriptions, a NextToken
// is also returned. Use the NextToken parameter in a new ListSubscriptions
// call to get further results.
//
// This action is throttled at 30 transactions per second (TPS).
//
//    // Example sending a request using ListSubscriptionsRequest.
//    req := client.ListSubscriptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/ListSubscriptions
func (c *Client) ListSubscriptionsRequest(input *ListSubscriptionsInput) ListSubscriptionsRequest {
	op := &aws.Operation{
		Name:       opListSubscriptions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListSubscriptionsInput{}
	}

	req := c.newRequest(op, input, &ListSubscriptionsOutput{})

	return ListSubscriptionsRequest{Request: req, Input: input, Copy: c.ListSubscriptionsRequest}
}

// ListSubscriptionsRequest is the request type for the
// ListSubscriptions API operation.
type ListSubscriptionsRequest struct {
	*aws.Request
	Input *ListSubscriptionsInput
	Copy  func(*ListSubscriptionsInput) ListSubscriptionsRequest
}

// Send marshals and sends the ListSubscriptions API request.
func (r ListSubscriptionsRequest) Send(ctx context.Context) (*ListSubscriptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSubscriptionsResponse{
		ListSubscriptionsOutput: r.Request.Data.(*ListSubscriptionsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListSubscriptionsRequestPaginator returns a paginator for ListSubscriptions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListSubscriptionsRequest(input)
//   p := sns.NewListSubscriptionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListSubscriptionsPaginator(req ListSubscriptionsRequest) ListSubscriptionsPaginator {
	return ListSubscriptionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListSubscriptionsInput
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

// ListSubscriptionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListSubscriptionsPaginator struct {
	aws.Pager
}

func (p *ListSubscriptionsPaginator) CurrentPage() *ListSubscriptionsOutput {
	return p.Pager.CurrentPage().(*ListSubscriptionsOutput)
}

// ListSubscriptionsResponse is the response type for the
// ListSubscriptions API operation.
type ListSubscriptionsResponse struct {
	*ListSubscriptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSubscriptions request.
func (r *ListSubscriptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
