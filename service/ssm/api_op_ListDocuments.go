// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListDocumentsInput struct {
	_ struct{} `type:"structure"`

	// This data type is deprecated. Instead, use Filters.
	DocumentFilterList []DocumentFilter `min:"1" type:"list"`

	// One or more DocumentKeyValuesFilter objects. Use a filter to return a more
	// specific list of results. For keys, you can specify one or more key-value
	// pair tags that have been applied to a document. Other valid keys include
	// Owner, Name, PlatformTypes, DocumentType, and TargetType. For example, to
	// return documents you own use Key=Owner,Values=Self. To specify a custom key-value
	// pair, use the format Key=tag:tagName,Values=valueName.
	Filters []DocumentKeyValuesFilter `type:"list"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDocumentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDocumentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDocumentsInput"}
	if s.DocumentFilterList != nil && len(s.DocumentFilterList) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DocumentFilterList", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.DocumentFilterList != nil {
		for i, v := range s.DocumentFilterList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "DocumentFilterList", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListDocumentsOutput struct {
	_ struct{} `type:"structure"`

	// The names of the Systems Manager documents.
	DocumentIdentifiers []DocumentIdentifier `type:"list"`

	// The token to use when requesting the next set of items. If there are no additional
	// items to return, the string is empty.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDocumentsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDocuments = "ListDocuments"

// ListDocumentsRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Returns all Systems Manager (SSM) documents in the current AWS account and
// Region. You can limit the results of this request by using a filter.
//
//    // Example sending a request using ListDocumentsRequest.
//    req := client.ListDocumentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/ListDocuments
func (c *Client) ListDocumentsRequest(input *ListDocumentsInput) ListDocumentsRequest {
	op := &aws.Operation{
		Name:       opListDocuments,
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
		input = &ListDocumentsInput{}
	}

	req := c.newRequest(op, input, &ListDocumentsOutput{})

	return ListDocumentsRequest{Request: req, Input: input, Copy: c.ListDocumentsRequest}
}

// ListDocumentsRequest is the request type for the
// ListDocuments API operation.
type ListDocumentsRequest struct {
	*aws.Request
	Input *ListDocumentsInput
	Copy  func(*ListDocumentsInput) ListDocumentsRequest
}

// Send marshals and sends the ListDocuments API request.
func (r ListDocumentsRequest) Send(ctx context.Context) (*ListDocumentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDocumentsResponse{
		ListDocumentsOutput: r.Request.Data.(*ListDocumentsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDocumentsRequestPaginator returns a paginator for ListDocuments.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDocumentsRequest(input)
//   p := ssm.NewListDocumentsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDocumentsPaginator(req ListDocumentsRequest) ListDocumentsPaginator {
	return ListDocumentsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDocumentsInput
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

// ListDocumentsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDocumentsPaginator struct {
	aws.Pager
}

func (p *ListDocumentsPaginator) CurrentPage() *ListDocumentsOutput {
	return p.Pager.CurrentPage().(*ListDocumentsOutput)
}

// ListDocumentsResponse is the response type for the
// ListDocuments API operation.
type ListDocumentsResponse struct {
	*ListDocumentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDocuments request.
func (r *ListDocumentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
