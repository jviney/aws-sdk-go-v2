// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package qldb

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListLedgersInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in a single ListLedgers request.
	// (The actual number of results returned might be fewer.)
	MaxResults *int64 `location:"querystring" locationName:"max_results" min:"1" type:"integer"`

	// A pagination token, indicating that you want to retrieve the next page of
	// results. If you received a value for NextToken in the response from a previous
	// ListLedgers call, then you should use that value as input here.
	NextToken *string `location:"querystring" locationName:"next_token" min:"4" type:"string"`
}

// String returns the string representation
func (s ListLedgersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListLedgersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListLedgersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListLedgersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "max_results", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "next_token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListLedgersOutput struct {
	_ struct{} `type:"structure"`

	// The array of ledger summaries that are associated with the current AWS account
	// and Region.
	Ledgers []LedgerSummary `type:"list"`

	// A pagination token, indicating whether there are more results available:
	//
	//    * If NextToken is empty, then the last page of results has been processed
	//    and there are no more results to be retrieved.
	//
	//    * If NextToken is not empty, then there are more results available. To
	//    retrieve the next page of results, use the value of NextToken in a subsequent
	//    ListLedgers call.
	NextToken *string `min:"4" type:"string"`
}

// String returns the string representation
func (s ListLedgersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListLedgersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Ledgers != nil {
		v := s.Ledgers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Ledgers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListLedgers = "ListLedgers"

// ListLedgersRequest returns a request value for making API operation for
// Amazon QLDB.
//
// Returns an array of ledger summaries that are associated with the current
// AWS account and Region.
//
// This action returns a maximum of 100 items and is paginated so that you can
// retrieve all the items by calling ListLedgers multiple times.
//
//    // Example sending a request using ListLedgersRequest.
//    req := client.ListLedgersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/qldb-2019-01-02/ListLedgers
func (c *Client) ListLedgersRequest(input *ListLedgersInput) ListLedgersRequest {
	op := &aws.Operation{
		Name:       opListLedgers,
		HTTPMethod: "GET",
		HTTPPath:   "/ledgers",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListLedgersInput{}
	}

	req := c.newRequest(op, input, &ListLedgersOutput{})

	return ListLedgersRequest{Request: req, Input: input, Copy: c.ListLedgersRequest}
}

// ListLedgersRequest is the request type for the
// ListLedgers API operation.
type ListLedgersRequest struct {
	*aws.Request
	Input *ListLedgersInput
	Copy  func(*ListLedgersInput) ListLedgersRequest
}

// Send marshals and sends the ListLedgers API request.
func (r ListLedgersRequest) Send(ctx context.Context) (*ListLedgersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListLedgersResponse{
		ListLedgersOutput: r.Request.Data.(*ListLedgersOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListLedgersRequestPaginator returns a paginator for ListLedgers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListLedgersRequest(input)
//   p := qldb.NewListLedgersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListLedgersPaginator(req ListLedgersRequest) ListLedgersPaginator {
	return ListLedgersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListLedgersInput
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

// ListLedgersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListLedgersPaginator struct {
	aws.Pager
}

func (p *ListLedgersPaginator) CurrentPage() *ListLedgersOutput {
	return p.Pager.CurrentPage().(*ListLedgersOutput)
}

// ListLedgersResponse is the response type for the
// ListLedgers API operation.
type ListLedgersResponse struct {
	*ListLedgersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListLedgers request.
func (r *ListLedgersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
