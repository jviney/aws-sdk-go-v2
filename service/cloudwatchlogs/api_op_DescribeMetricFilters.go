// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeMetricFiltersInput struct {
	_ struct{} `type:"structure"`

	// The prefix to match.
	FilterNamePrefix *string `locationName:"filterNamePrefix" min:"1" type:"string"`

	// The maximum number of items returned. If you don't specify a value, the default
	// is up to 50 items.
	Limit *int64 `locationName:"limit" min:"1" type:"integer"`

	// The name of the log group.
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string"`

	// Filters results to include only those with the specified metric name. If
	// you include this parameter in your request, you must also include the metricNamespace
	// parameter.
	MetricName *string `locationName:"metricName" type:"string"`

	// Filters results to include only those in the specified namespace. If you
	// include this parameter in your request, you must also include the metricName
	// parameter.
	MetricNamespace *string `locationName:"metricNamespace" type:"string"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeMetricFiltersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeMetricFiltersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeMetricFiltersInput"}
	if s.FilterNamePrefix != nil && len(*s.FilterNamePrefix) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FilterNamePrefix", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.LogGroupName != nil && len(*s.LogGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogGroupName", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeMetricFiltersOutput struct {
	_ struct{} `type:"structure"`

	// The metric filters.
	MetricFilters []MetricFilter `locationName:"metricFilters" type:"list"`

	// The token for the next set of items to return. The token expires after 24
	// hours.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeMetricFiltersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeMetricFilters = "DescribeMetricFilters"

// DescribeMetricFiltersRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Lists the specified metric filters. You can list all the metric filters or
// filter the results by log name, prefix, metric name, or metric namespace.
// The results are ASCII-sorted by filter name.
//
//    // Example sending a request using DescribeMetricFiltersRequest.
//    req := client.DescribeMetricFiltersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DescribeMetricFilters
func (c *Client) DescribeMetricFiltersRequest(input *DescribeMetricFiltersInput) DescribeMetricFiltersRequest {
	op := &aws.Operation{
		Name:       opDescribeMetricFilters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeMetricFiltersInput{}
	}

	req := c.newRequest(op, input, &DescribeMetricFiltersOutput{})

	return DescribeMetricFiltersRequest{Request: req, Input: input, Copy: c.DescribeMetricFiltersRequest}
}

// DescribeMetricFiltersRequest is the request type for the
// DescribeMetricFilters API operation.
type DescribeMetricFiltersRequest struct {
	*aws.Request
	Input *DescribeMetricFiltersInput
	Copy  func(*DescribeMetricFiltersInput) DescribeMetricFiltersRequest
}

// Send marshals and sends the DescribeMetricFilters API request.
func (r DescribeMetricFiltersRequest) Send(ctx context.Context) (*DescribeMetricFiltersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeMetricFiltersResponse{
		DescribeMetricFiltersOutput: r.Request.Data.(*DescribeMetricFiltersOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeMetricFiltersRequestPaginator returns a paginator for DescribeMetricFilters.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeMetricFiltersRequest(input)
//   p := cloudwatchlogs.NewDescribeMetricFiltersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeMetricFiltersPaginator(req DescribeMetricFiltersRequest) DescribeMetricFiltersPaginator {
	return DescribeMetricFiltersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeMetricFiltersInput
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

// DescribeMetricFiltersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeMetricFiltersPaginator struct {
	aws.Pager
}

func (p *DescribeMetricFiltersPaginator) CurrentPage() *DescribeMetricFiltersOutput {
	return p.Pager.CurrentPage().(*DescribeMetricFiltersOutput)
}

// DescribeMetricFiltersResponse is the response type for the
// DescribeMetricFilters API operation.
type DescribeMetricFiltersResponse struct {
	*DescribeMetricFiltersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeMetricFilters request.
func (r *DescribeMetricFiltersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
