// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// This input determines which bootstrap actions to retrieve.
type ListBootstrapActionsInput struct {
	_ struct{} `type:"structure"`

	// The cluster identifier for the bootstrap actions to list.
	//
	// ClusterId is a required field
	ClusterId *string `type:"string" required:"true"`

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s ListBootstrapActionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBootstrapActionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListBootstrapActionsInput"}

	if s.ClusterId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// This output contains the bootstrap actions detail.
type ListBootstrapActionsOutput struct {
	_ struct{} `type:"structure"`

	// The bootstrap actions associated with the cluster.
	BootstrapActions []Command `type:"list"`

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s ListBootstrapActionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListBootstrapActions = "ListBootstrapActions"

// ListBootstrapActionsRequest returns a request value for making API operation for
// Amazon Elastic MapReduce.
//
// Provides information about the bootstrap actions associated with a cluster.
//
//    // Example sending a request using ListBootstrapActionsRequest.
//    req := client.ListBootstrapActionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/ListBootstrapActions
func (c *Client) ListBootstrapActionsRequest(input *ListBootstrapActionsInput) ListBootstrapActionsRequest {
	op := &aws.Operation{
		Name:       opListBootstrapActions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListBootstrapActionsInput{}
	}

	req := c.newRequest(op, input, &ListBootstrapActionsOutput{})

	return ListBootstrapActionsRequest{Request: req, Input: input, Copy: c.ListBootstrapActionsRequest}
}

// ListBootstrapActionsRequest is the request type for the
// ListBootstrapActions API operation.
type ListBootstrapActionsRequest struct {
	*aws.Request
	Input *ListBootstrapActionsInput
	Copy  func(*ListBootstrapActionsInput) ListBootstrapActionsRequest
}

// Send marshals and sends the ListBootstrapActions API request.
func (r ListBootstrapActionsRequest) Send(ctx context.Context) (*ListBootstrapActionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListBootstrapActionsResponse{
		ListBootstrapActionsOutput: r.Request.Data.(*ListBootstrapActionsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListBootstrapActionsRequestPaginator returns a paginator for ListBootstrapActions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListBootstrapActionsRequest(input)
//   p := emr.NewListBootstrapActionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListBootstrapActionsPaginator(req ListBootstrapActionsRequest) ListBootstrapActionsPaginator {
	return ListBootstrapActionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListBootstrapActionsInput
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

// ListBootstrapActionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListBootstrapActionsPaginator struct {
	aws.Pager
}

func (p *ListBootstrapActionsPaginator) CurrentPage() *ListBootstrapActionsOutput {
	return p.Pager.CurrentPage().(*ListBootstrapActionsOutput)
}

// ListBootstrapActionsResponse is the response type for the
// ListBootstrapActions API operation.
type ListBootstrapActionsResponse struct {
	*ListBootstrapActionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListBootstrapActions request.
func (r *ListBootstrapActionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
