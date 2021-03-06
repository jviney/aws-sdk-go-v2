// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListAliasesInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the entity for which to list the aliases.
	//
	// EntityId is a required field
	EntityId *string `min:"12" type:"string" required:"true"`

	// The maximum number of results to return in a single call.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token to use to retrieve the next page of results. The first call does
	// not contain any tokens.
	NextToken *string `min:"1" type:"string"`

	// The identifier for the organization under which the entity exists.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListAliasesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAliasesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAliasesInput"}

	if s.EntityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntityId"))
	}
	if s.EntityId != nil && len(*s.EntityId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityId", 12))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListAliasesOutput struct {
	_ struct{} `type:"structure"`

	// The entity's paginated aliases.
	Aliases []string `type:"list"`

	// The token to use to retrieve the next page of results. The value is "null"
	// when there are no more results to return.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListAliasesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListAliases = "ListAliases"

// ListAliasesRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Creates a paginated call to list the aliases associated with a given entity.
//
//    // Example sending a request using ListAliasesRequest.
//    req := client.ListAliasesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/ListAliases
func (c *Client) ListAliasesRequest(input *ListAliasesInput) ListAliasesRequest {
	op := &aws.Operation{
		Name:       opListAliases,
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
		input = &ListAliasesInput{}
	}

	req := c.newRequest(op, input, &ListAliasesOutput{})

	return ListAliasesRequest{Request: req, Input: input, Copy: c.ListAliasesRequest}
}

// ListAliasesRequest is the request type for the
// ListAliases API operation.
type ListAliasesRequest struct {
	*aws.Request
	Input *ListAliasesInput
	Copy  func(*ListAliasesInput) ListAliasesRequest
}

// Send marshals and sends the ListAliases API request.
func (r ListAliasesRequest) Send(ctx context.Context) (*ListAliasesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAliasesResponse{
		ListAliasesOutput: r.Request.Data.(*ListAliasesOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListAliasesRequestPaginator returns a paginator for ListAliases.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListAliasesRequest(input)
//   p := workmail.NewListAliasesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListAliasesPaginator(req ListAliasesRequest) ListAliasesPaginator {
	return ListAliasesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListAliasesInput
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

// ListAliasesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListAliasesPaginator struct {
	aws.Pager
}

func (p *ListAliasesPaginator) CurrentPage() *ListAliasesOutput {
	return p.Pager.CurrentPage().(*ListAliasesOutput)
}

// ListAliasesResponse is the response type for the
// ListAliases API operation.
type ListAliasesResponse struct {
	*ListAliasesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAliases request.
func (r *ListAliasesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
