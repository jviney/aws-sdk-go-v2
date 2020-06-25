// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package translate

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListTerminologiesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of custom terminologies returned per list request.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the result of the request to ListTerminologies was truncated, include
	// the NextToken to fetch the next group of custom terminologies.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListTerminologiesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTerminologiesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTerminologiesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListTerminologiesOutput struct {
	_ struct{} `type:"structure"`

	// If the response to the ListTerminologies was truncated, the NextToken fetches
	// the next group of custom terminologies.
	NextToken *string `type:"string"`

	// The properties list of the custom terminologies returned on the list request.
	TerminologyPropertiesList []TerminologyProperties `type:"list"`
}

// String returns the string representation
func (s ListTerminologiesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTerminologies = "ListTerminologies"

// ListTerminologiesRequest returns a request value for making API operation for
// Amazon Translate.
//
// Provides a list of custom terminologies associated with your account.
//
//    // Example sending a request using ListTerminologiesRequest.
//    req := client.ListTerminologiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/translate-2017-07-01/ListTerminologies
func (c *Client) ListTerminologiesRequest(input *ListTerminologiesInput) ListTerminologiesRequest {
	op := &aws.Operation{
		Name:       opListTerminologies,
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
		input = &ListTerminologiesInput{}
	}

	req := c.newRequest(op, input, &ListTerminologiesOutput{})

	return ListTerminologiesRequest{Request: req, Input: input, Copy: c.ListTerminologiesRequest}
}

// ListTerminologiesRequest is the request type for the
// ListTerminologies API operation.
type ListTerminologiesRequest struct {
	*aws.Request
	Input *ListTerminologiesInput
	Copy  func(*ListTerminologiesInput) ListTerminologiesRequest
}

// Send marshals and sends the ListTerminologies API request.
func (r ListTerminologiesRequest) Send(ctx context.Context) (*ListTerminologiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTerminologiesResponse{
		ListTerminologiesOutput: r.Request.Data.(*ListTerminologiesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTerminologiesRequestPaginator returns a paginator for ListTerminologies.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTerminologiesRequest(input)
//   p := translate.NewListTerminologiesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTerminologiesPaginator(req ListTerminologiesRequest) ListTerminologiesPaginator {
	return ListTerminologiesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListTerminologiesInput
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

// ListTerminologiesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTerminologiesPaginator struct {
	aws.Pager
}

func (p *ListTerminologiesPaginator) CurrentPage() *ListTerminologiesOutput {
	return p.Pager.CurrentPage().(*ListTerminologiesOutput)
}

// ListTerminologiesResponse is the response type for the
// ListTerminologies API operation.
type ListTerminologiesResponse struct {
	*ListTerminologiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTerminologies request.
func (r *ListTerminologiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
