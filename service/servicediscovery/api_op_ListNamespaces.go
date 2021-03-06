// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListNamespacesInput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains specifications for the namespaces that you want
	// to list.
	//
	// If you specify more than one filter, a namespace must match all filters to
	// be returned by ListNamespaces.
	Filters []NamespaceFilter `type:"list"`

	// The maximum number of namespaces that you want AWS Cloud Map to return in
	// the response to a ListNamespaces request. If you don't specify a value for
	// MaxResults, AWS Cloud Map returns up to 100 namespaces.
	MaxResults *int64 `min:"1" type:"integer"`

	// For the first ListNamespaces request, omit this value.
	//
	// If the response contains NextToken, submit another ListNamespaces request
	// to get the next group of results. Specify the value of NextToken from the
	// previous response in the next request.
	//
	// AWS Cloud Map gets MaxResults namespaces and then filters them based on the
	// specified criteria. It's possible that no namespaces in the first MaxResults
	// namespaces matched the specified criteria but that subsequent groups of MaxResults
	// namespaces do contain namespaces that match the criteria.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListNamespacesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListNamespacesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListNamespacesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
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

type ListNamespacesOutput struct {
	_ struct{} `type:"structure"`

	// An array that contains one NamespaceSummary object for each namespace that
	// matches the specified filter criteria.
	Namespaces []NamespaceSummary `type:"list"`

	// If the response contains NextToken, submit another ListNamespaces request
	// to get the next group of results. Specify the value of NextToken from the
	// previous response in the next request.
	//
	// AWS Cloud Map gets MaxResults namespaces and then filters them based on the
	// specified criteria. It's possible that no namespaces in the first MaxResults
	// namespaces matched the specified criteria but that subsequent groups of MaxResults
	// namespaces do contain namespaces that match the criteria.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListNamespacesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListNamespaces = "ListNamespaces"

// ListNamespacesRequest returns a request value for making API operation for
// AWS Cloud Map.
//
// Lists summary information about the namespaces that were created by the current
// AWS account.
//
//    // Example sending a request using ListNamespacesRequest.
//    req := client.ListNamespacesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/ListNamespaces
func (c *Client) ListNamespacesRequest(input *ListNamespacesInput) ListNamespacesRequest {
	op := &aws.Operation{
		Name:       opListNamespaces,
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
		input = &ListNamespacesInput{}
	}

	req := c.newRequest(op, input, &ListNamespacesOutput{})

	return ListNamespacesRequest{Request: req, Input: input, Copy: c.ListNamespacesRequest}
}

// ListNamespacesRequest is the request type for the
// ListNamespaces API operation.
type ListNamespacesRequest struct {
	*aws.Request
	Input *ListNamespacesInput
	Copy  func(*ListNamespacesInput) ListNamespacesRequest
}

// Send marshals and sends the ListNamespaces API request.
func (r ListNamespacesRequest) Send(ctx context.Context) (*ListNamespacesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListNamespacesResponse{
		ListNamespacesOutput: r.Request.Data.(*ListNamespacesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListNamespacesRequestPaginator returns a paginator for ListNamespaces.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListNamespacesRequest(input)
//   p := servicediscovery.NewListNamespacesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListNamespacesPaginator(req ListNamespacesRequest) ListNamespacesPaginator {
	return ListNamespacesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListNamespacesInput
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

// ListNamespacesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListNamespacesPaginator struct {
	aws.Pager
}

func (p *ListNamespacesPaginator) CurrentPage() *ListNamespacesOutput {
	return p.Pager.CurrentPage().(*ListNamespacesOutput)
}

// ListNamespacesResponse is the response type for the
// ListNamespaces API operation.
type ListNamespacesResponse struct {
	*ListNamespacesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListNamespaces request.
func (r *ListNamespacesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
