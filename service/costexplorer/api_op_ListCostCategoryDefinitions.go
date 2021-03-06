// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costexplorer

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListCostCategoryDefinitionsInput struct {
	_ struct{} `type:"structure"`

	// The date when the Cost Category was effective.
	EffectiveOn *string `min:"20" type:"string"`

	// The number of entries a paginated response contains.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token to retrieve the next set of results. Amazon Web Services provides
	// the token when the response from a previous call has more results than the
	// maximum page size.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListCostCategoryDefinitionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListCostCategoryDefinitionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListCostCategoryDefinitionsInput"}
	if s.EffectiveOn != nil && len(*s.EffectiveOn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("EffectiveOn", 20))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListCostCategoryDefinitionsOutput struct {
	_ struct{} `type:"structure"`

	// A reference to a Cost Category containing enough information to identify
	// the Cost Category.
	CostCategoryReferences []CostCategoryReference `type:"list"`

	// The token to retrieve the next set of results. Amazon Web Services provides
	// the token when the response from a previous call has more results than the
	// maximum page size.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListCostCategoryDefinitionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListCostCategoryDefinitions = "ListCostCategoryDefinitions"

// ListCostCategoryDefinitionsRequest returns a request value for making API operation for
// AWS Cost Explorer Service.
//
// Returns the name, ARN, NumberOfRules and effective dates of all Cost Categories
// defined in the account. You have the option to use EffectiveOn to return
// a list of Cost Categories that were active on a specific date. If there is
// no EffectiveOn specified, you’ll see Cost Categories that are effective
// on the current date. If Cost Category is still effective, EffectiveEnd is
// omitted in the response. ListCostCategoryDefinitions supports pagination.
// The request can have a MaxResults range up to 100.
//
//    // Example sending a request using ListCostCategoryDefinitionsRequest.
//    req := client.ListCostCategoryDefinitionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/ListCostCategoryDefinitions
func (c *Client) ListCostCategoryDefinitionsRequest(input *ListCostCategoryDefinitionsInput) ListCostCategoryDefinitionsRequest {
	op := &aws.Operation{
		Name:       opListCostCategoryDefinitions,
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
		input = &ListCostCategoryDefinitionsInput{}
	}

	req := c.newRequest(op, input, &ListCostCategoryDefinitionsOutput{})

	return ListCostCategoryDefinitionsRequest{Request: req, Input: input, Copy: c.ListCostCategoryDefinitionsRequest}
}

// ListCostCategoryDefinitionsRequest is the request type for the
// ListCostCategoryDefinitions API operation.
type ListCostCategoryDefinitionsRequest struct {
	*aws.Request
	Input *ListCostCategoryDefinitionsInput
	Copy  func(*ListCostCategoryDefinitionsInput) ListCostCategoryDefinitionsRequest
}

// Send marshals and sends the ListCostCategoryDefinitions API request.
func (r ListCostCategoryDefinitionsRequest) Send(ctx context.Context) (*ListCostCategoryDefinitionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListCostCategoryDefinitionsResponse{
		ListCostCategoryDefinitionsOutput: r.Request.Data.(*ListCostCategoryDefinitionsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListCostCategoryDefinitionsRequestPaginator returns a paginator for ListCostCategoryDefinitions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListCostCategoryDefinitionsRequest(input)
//   p := costexplorer.NewListCostCategoryDefinitionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListCostCategoryDefinitionsPaginator(req ListCostCategoryDefinitionsRequest) ListCostCategoryDefinitionsPaginator {
	return ListCostCategoryDefinitionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListCostCategoryDefinitionsInput
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

// ListCostCategoryDefinitionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListCostCategoryDefinitionsPaginator struct {
	aws.Pager
}

func (p *ListCostCategoryDefinitionsPaginator) CurrentPage() *ListCostCategoryDefinitionsOutput {
	return p.Pager.CurrentPage().(*ListCostCategoryDefinitionsOutput)
}

// ListCostCategoryDefinitionsResponse is the response type for the
// ListCostCategoryDefinitions API operation.
type ListCostCategoryDefinitionsResponse struct {
	*ListCostCategoryDefinitionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListCostCategoryDefinitions request.
func (r *ListCostCategoryDefinitionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
