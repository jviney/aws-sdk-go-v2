// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListFlowExecutionMessagesInput struct {
	_ struct{} `type:"structure"`

	// The ID of the flow execution.
	//
	// FlowExecutionId is a required field
	FlowExecutionId *string `locationName:"flowExecutionId" type:"string" required:"true"`

	// The maximum number of results to return in the response.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListFlowExecutionMessagesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFlowExecutionMessagesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFlowExecutionMessagesInput"}

	if s.FlowExecutionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FlowExecutionId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListFlowExecutionMessagesOutput struct {
	_ struct{} `type:"structure"`

	// A list of objects that contain information about events in the specified
	// flow execution.
	Messages []FlowExecutionMessage `locationName:"messages" type:"list"`

	// The string to specify as nextToken when you request the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListFlowExecutionMessagesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListFlowExecutionMessages = "ListFlowExecutionMessages"

// ListFlowExecutionMessagesRequest returns a request value for making API operation for
// AWS IoT Things Graph.
//
// Returns a list of objects that contain information about events in a flow
// execution.
//
//    // Example sending a request using ListFlowExecutionMessagesRequest.
//    req := client.ListFlowExecutionMessagesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/ListFlowExecutionMessages
func (c *Client) ListFlowExecutionMessagesRequest(input *ListFlowExecutionMessagesInput) ListFlowExecutionMessagesRequest {
	op := &aws.Operation{
		Name:       opListFlowExecutionMessages,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListFlowExecutionMessagesInput{}
	}

	req := c.newRequest(op, input, &ListFlowExecutionMessagesOutput{})

	return ListFlowExecutionMessagesRequest{Request: req, Input: input, Copy: c.ListFlowExecutionMessagesRequest}
}

// ListFlowExecutionMessagesRequest is the request type for the
// ListFlowExecutionMessages API operation.
type ListFlowExecutionMessagesRequest struct {
	*aws.Request
	Input *ListFlowExecutionMessagesInput
	Copy  func(*ListFlowExecutionMessagesInput) ListFlowExecutionMessagesRequest
}

// Send marshals and sends the ListFlowExecutionMessages API request.
func (r ListFlowExecutionMessagesRequest) Send(ctx context.Context) (*ListFlowExecutionMessagesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFlowExecutionMessagesResponse{
		ListFlowExecutionMessagesOutput: r.Request.Data.(*ListFlowExecutionMessagesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListFlowExecutionMessagesRequestPaginator returns a paginator for ListFlowExecutionMessages.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListFlowExecutionMessagesRequest(input)
//   p := iotthingsgraph.NewListFlowExecutionMessagesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListFlowExecutionMessagesPaginator(req ListFlowExecutionMessagesRequest) ListFlowExecutionMessagesPaginator {
	return ListFlowExecutionMessagesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListFlowExecutionMessagesInput
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

// ListFlowExecutionMessagesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListFlowExecutionMessagesPaginator struct {
	aws.Pager
}

func (p *ListFlowExecutionMessagesPaginator) CurrentPage() *ListFlowExecutionMessagesOutput {
	return p.Pager.CurrentPage().(*ListFlowExecutionMessagesOutput)
}

// ListFlowExecutionMessagesResponse is the response type for the
// ListFlowExecutionMessages API operation.
type ListFlowExecutionMessagesResponse struct {
	*ListFlowExecutionMessagesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFlowExecutionMessages request.
func (r *ListFlowExecutionMessagesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
