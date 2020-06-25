// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeParametersInput struct {
	_ struct{} `type:"structure"`

	// This data type is deprecated. Instead, use ParameterFilters.
	Filters []ParametersFilter `type:"list"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`

	// Filters to limit the request results.
	ParameterFilters []ParameterStringFilter `type:"list"`
}

// String returns the string representation
func (s DescribeParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeParametersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeParametersInput"}
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
	if s.ParameterFilters != nil {
		for i, v := range s.ParameterFilters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ParameterFilters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeParametersOutput struct {
	_ struct{} `type:"structure"`

	// The token to use when requesting the next set of items.
	NextToken *string `type:"string"`

	// Parameters returned by the request.
	Parameters []ParameterMetadata `type:"list"`
}

// String returns the string representation
func (s DescribeParametersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeParameters = "DescribeParameters"

// DescribeParametersRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Get information about a parameter.
//
// Request results are returned on a best-effort basis. If you specify MaxResults
// in the request, the response includes information up to the limit specified.
// The number of items returned, however, can be between zero and the value
// of MaxResults. If the service reaches an internal limit while processing
// the results, it stops the operation and returns the matching values up to
// that point and a NextToken. You can specify the NextToken in a subsequent
// call to get the next set of results.
//
//    // Example sending a request using DescribeParametersRequest.
//    req := client.DescribeParametersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeParameters
func (c *Client) DescribeParametersRequest(input *DescribeParametersInput) DescribeParametersRequest {
	op := &aws.Operation{
		Name:       opDescribeParameters,
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
		input = &DescribeParametersInput{}
	}

	req := c.newRequest(op, input, &DescribeParametersOutput{})

	return DescribeParametersRequest{Request: req, Input: input, Copy: c.DescribeParametersRequest}
}

// DescribeParametersRequest is the request type for the
// DescribeParameters API operation.
type DescribeParametersRequest struct {
	*aws.Request
	Input *DescribeParametersInput
	Copy  func(*DescribeParametersInput) DescribeParametersRequest
}

// Send marshals and sends the DescribeParameters API request.
func (r DescribeParametersRequest) Send(ctx context.Context) (*DescribeParametersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeParametersResponse{
		DescribeParametersOutput: r.Request.Data.(*DescribeParametersOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeParametersRequestPaginator returns a paginator for DescribeParameters.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeParametersRequest(input)
//   p := ssm.NewDescribeParametersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeParametersPaginator(req DescribeParametersRequest) DescribeParametersPaginator {
	return DescribeParametersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeParametersInput
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

// DescribeParametersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeParametersPaginator struct {
	aws.Pager
}

func (p *DescribeParametersPaginator) CurrentPage() *DescribeParametersOutput {
	return p.Pager.CurrentPage().(*DescribeParametersOutput)
}

// DescribeParametersResponse is the response type for the
// DescribeParameters API operation.
type DescribeParametersResponse struct {
	*DescribeParametersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeParameters request.
func (r *DescribeParametersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
