// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListAlgorithmsInput struct {
	_ struct{} `type:"structure"`

	// A filter that returns only algorithms created after the specified time (timestamp).
	CreationTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only algorithms created before the specified time (timestamp).
	CreationTimeBefore *time.Time `type:"timestamp"`

	// The maximum number of algorithms to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A string in the algorithm name. This filter returns only algorithms whose
	// name contains the specified string.
	NameContains *string `type:"string"`

	// If the response to a previous ListAlgorithms request was truncated, the response
	// includes a NextToken. To retrieve the next set of algorithms, use the token
	// in the next request.
	NextToken *string `type:"string"`

	// The parameter by which to sort the results. The default is CreationTime.
	SortBy AlgorithmSortBy `type:"string" enum:"true"`

	// The sort order for the results. The default is Ascending.
	SortOrder SortOrder `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListAlgorithmsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAlgorithmsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAlgorithmsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListAlgorithmsOutput struct {
	_ struct{} `type:"structure"`

	// >An array of AlgorithmSummary objects, each of which lists an algorithm.
	//
	// AlgorithmSummaryList is a required field
	AlgorithmSummaryList []AlgorithmSummary `type:"list" required:"true"`

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of algorithms, use it in the subsequent request.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListAlgorithmsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListAlgorithms = "ListAlgorithms"

// ListAlgorithmsRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Lists the machine learning algorithms that have been created.
//
//    // Example sending a request using ListAlgorithmsRequest.
//    req := client.ListAlgorithmsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListAlgorithms
func (c *Client) ListAlgorithmsRequest(input *ListAlgorithmsInput) ListAlgorithmsRequest {
	op := &aws.Operation{
		Name:       opListAlgorithms,
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
		input = &ListAlgorithmsInput{}
	}

	req := c.newRequest(op, input, &ListAlgorithmsOutput{})

	return ListAlgorithmsRequest{Request: req, Input: input, Copy: c.ListAlgorithmsRequest}
}

// ListAlgorithmsRequest is the request type for the
// ListAlgorithms API operation.
type ListAlgorithmsRequest struct {
	*aws.Request
	Input *ListAlgorithmsInput
	Copy  func(*ListAlgorithmsInput) ListAlgorithmsRequest
}

// Send marshals and sends the ListAlgorithms API request.
func (r ListAlgorithmsRequest) Send(ctx context.Context) (*ListAlgorithmsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAlgorithmsResponse{
		ListAlgorithmsOutput: r.Request.Data.(*ListAlgorithmsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListAlgorithmsRequestPaginator returns a paginator for ListAlgorithms.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListAlgorithmsRequest(input)
//   p := sagemaker.NewListAlgorithmsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListAlgorithmsPaginator(req ListAlgorithmsRequest) ListAlgorithmsPaginator {
	return ListAlgorithmsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListAlgorithmsInput
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

// ListAlgorithmsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListAlgorithmsPaginator struct {
	aws.Pager
}

func (p *ListAlgorithmsPaginator) CurrentPage() *ListAlgorithmsOutput {
	return p.Pager.CurrentPage().(*ListAlgorithmsOutput)
}

// ListAlgorithmsResponse is the response type for the
// ListAlgorithms API operation.
type ListAlgorithmsResponse struct {
	*ListAlgorithmsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAlgorithms request.
func (r *ListAlgorithmsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
