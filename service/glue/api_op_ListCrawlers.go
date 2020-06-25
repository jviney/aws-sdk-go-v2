// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListCrawlersInput struct {
	_ struct{} `type:"structure"`

	// The maximum size of a list to return.
	MaxResults *int64 `min:"1" type:"integer"`

	// A continuation token, if this is a continuation request.
	NextToken *string `type:"string"`

	// Specifies to return only these tagged resources.
	Tags map[string]string `type:"map"`
}

// String returns the string representation
func (s ListCrawlersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListCrawlersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListCrawlersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListCrawlersOutput struct {
	_ struct{} `type:"structure"`

	// The names of all crawlers in the account, or the crawlers with the specified
	// tags.
	CrawlerNames []string `type:"list"`

	// A continuation token, if the returned list does not contain the last metric
	// available.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListCrawlersOutput) String() string {
	return awsutil.Prettify(s)
}

const opListCrawlers = "ListCrawlers"

// ListCrawlersRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves the names of all crawler resources in this AWS account, or the
// resources with the specified tag. This operation allows you to see which
// resources are available in your account, and their names.
//
// This operation takes the optional Tags field, which you can use as a filter
// on the response so that tagged resources can be retrieved as a group. If
// you choose to use tags filtering, only resources with the tag are retrieved.
//
//    // Example sending a request using ListCrawlersRequest.
//    req := client.ListCrawlersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/ListCrawlers
func (c *Client) ListCrawlersRequest(input *ListCrawlersInput) ListCrawlersRequest {
	op := &aws.Operation{
		Name:       opListCrawlers,
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
		input = &ListCrawlersInput{}
	}

	req := c.newRequest(op, input, &ListCrawlersOutput{})

	return ListCrawlersRequest{Request: req, Input: input, Copy: c.ListCrawlersRequest}
}

// ListCrawlersRequest is the request type for the
// ListCrawlers API operation.
type ListCrawlersRequest struct {
	*aws.Request
	Input *ListCrawlersInput
	Copy  func(*ListCrawlersInput) ListCrawlersRequest
}

// Send marshals and sends the ListCrawlers API request.
func (r ListCrawlersRequest) Send(ctx context.Context) (*ListCrawlersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListCrawlersResponse{
		ListCrawlersOutput: r.Request.Data.(*ListCrawlersOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListCrawlersRequestPaginator returns a paginator for ListCrawlers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListCrawlersRequest(input)
//   p := glue.NewListCrawlersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListCrawlersPaginator(req ListCrawlersRequest) ListCrawlersPaginator {
	return ListCrawlersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListCrawlersInput
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

// ListCrawlersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListCrawlersPaginator struct {
	aws.Pager
}

func (p *ListCrawlersPaginator) CurrentPage() *ListCrawlersOutput {
	return p.Pager.CurrentPage().(*ListCrawlersOutput)
}

// ListCrawlersResponse is the response type for the
// ListCrawlers API operation.
type ListCrawlersResponse struct {
	*ListCrawlersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListCrawlers request.
func (r *ListCrawlersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
