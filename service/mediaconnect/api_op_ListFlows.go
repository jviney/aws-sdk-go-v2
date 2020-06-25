// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListFlowsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListFlowsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFlowsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFlowsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFlowsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

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
	return nil
}

// The result of a successful ListFlows request. The response includes flow
// summaries and the NextToken to use in a subsequent ListFlows request.
type ListFlowsOutput struct {
	_ struct{} `type:"structure"`

	// A list of flow summaries.
	Flows []ListedFlow `locationName:"flows" type:"list"`

	// The token that identifies which batch of results that you want to see. For
	// example, you submit a ListFlows request with MaxResults set at 5. The service
	// returns the first batch of results (up to 5) and a NextToken value. To see
	// the next batch of results, you can submit the ListFlows request a second
	// time and specify the NextToken value.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListFlowsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFlowsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Flows != nil {
		v := s.Flows

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "flows", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListFlows = "ListFlows"

// ListFlowsRequest returns a request value for making API operation for
// AWS MediaConnect.
//
// Displays a list of flows that are associated with this account. This request
// returns a paginated result.
//
//    // Example sending a request using ListFlowsRequest.
//    req := client.ListFlowsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/ListFlows
func (c *Client) ListFlowsRequest(input *ListFlowsInput) ListFlowsRequest {
	op := &aws.Operation{
		Name:       opListFlows,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/flows",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListFlowsInput{}
	}

	req := c.newRequest(op, input, &ListFlowsOutput{})

	return ListFlowsRequest{Request: req, Input: input, Copy: c.ListFlowsRequest}
}

// ListFlowsRequest is the request type for the
// ListFlows API operation.
type ListFlowsRequest struct {
	*aws.Request
	Input *ListFlowsInput
	Copy  func(*ListFlowsInput) ListFlowsRequest
}

// Send marshals and sends the ListFlows API request.
func (r ListFlowsRequest) Send(ctx context.Context) (*ListFlowsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFlowsResponse{
		ListFlowsOutput: r.Request.Data.(*ListFlowsOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListFlowsRequestPaginator returns a paginator for ListFlows.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListFlowsRequest(input)
//   p := mediaconnect.NewListFlowsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListFlowsPaginator(req ListFlowsRequest) ListFlowsPaginator {
	return ListFlowsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListFlowsInput
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

// ListFlowsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListFlowsPaginator struct {
	aws.Pager
}

func (p *ListFlowsPaginator) CurrentPage() *ListFlowsOutput {
	return p.Pager.CurrentPage().(*ListFlowsOutput)
}

// ListFlowsResponse is the response type for the
// ListFlows API operation.
type ListFlowsResponse struct {
	*ListFlowsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFlows request.
func (r *ListFlowsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
