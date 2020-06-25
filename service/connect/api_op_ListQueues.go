// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListQueuesInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the Amazon Connect instance.
	//
	// InstanceId is a required field
	InstanceId *string `location:"uri" locationName:"InstanceId" min:"1" type:"string" required:"true"`

	// The maximimum number of results to return per page.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The type of queue.
	QueueTypes []QueueType `location:"querystring" locationName:"queueTypes" type:"list"`
}

// String returns the string representation
func (s ListQueuesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListQueuesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListQueuesInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListQueuesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.InstanceId != nil {
		v := *s.InstanceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "InstanceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
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
	if s.QueueTypes != nil {
		v := s.QueueTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.QueryTarget, "queueTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

type ListQueuesOutput struct {
	_ struct{} `type:"structure"`

	// If there are additional results, this is the token for the next set of results.
	NextToken *string `type:"string"`

	// Information about the queues.
	QueueSummaryList []QueueSummary `type:"list"`
}

// String returns the string representation
func (s ListQueuesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListQueuesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.QueueSummaryList != nil {
		v := s.QueueSummaryList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "QueueSummaryList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListQueues = "ListQueues"

// ListQueuesRequest returns a request value for making API operation for
// Amazon Connect Service.
//
// Provides information about the queues for the specified Amazon Connect instance.
//
//    // Example sending a request using ListQueuesRequest.
//    req := client.ListQueuesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/ListQueues
func (c *Client) ListQueuesRequest(input *ListQueuesInput) ListQueuesRequest {
	op := &aws.Operation{
		Name:       opListQueues,
		HTTPMethod: "GET",
		HTTPPath:   "/queues-summary/{InstanceId}",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListQueuesInput{}
	}

	req := c.newRequest(op, input, &ListQueuesOutput{})

	return ListQueuesRequest{Request: req, Input: input, Copy: c.ListQueuesRequest}
}

// ListQueuesRequest is the request type for the
// ListQueues API operation.
type ListQueuesRequest struct {
	*aws.Request
	Input *ListQueuesInput
	Copy  func(*ListQueuesInput) ListQueuesRequest
}

// Send marshals and sends the ListQueues API request.
func (r ListQueuesRequest) Send(ctx context.Context) (*ListQueuesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListQueuesResponse{
		ListQueuesOutput: r.Request.Data.(*ListQueuesOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListQueuesRequestPaginator returns a paginator for ListQueues.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListQueuesRequest(input)
//   p := connect.NewListQueuesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListQueuesPaginator(req ListQueuesRequest) ListQueuesPaginator {
	return ListQueuesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListQueuesInput
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

// ListQueuesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListQueuesPaginator struct {
	aws.Pager
}

func (p *ListQueuesPaginator) CurrentPage() *ListQueuesOutput {
	return p.Pager.CurrentPage().(*ListQueuesOutput)
}

// ListQueuesResponse is the response type for the
// ListQueues API operation.
type ListQueuesResponse struct {
	*ListQueuesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListQueues request.
func (r *ListQueuesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
