// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package forecast

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListForecastExportJobsInput struct {
	_ struct{} `type:"structure"`

	// An array of filters. For each filter, you provide a condition and a match
	// statement. The condition is either IS or IS_NOT, which specifies whether
	// to include or exclude the forecast export jobs that match the statement from
	// the list, respectively. The match statement consists of a key and a value.
	//
	// Filter properties
	//
	//    * Condition - The condition to apply. Valid values are IS and IS_NOT.
	//    To include the forecast export jobs that match the statement, specify
	//    IS. To exclude matching forecast export jobs, specify IS_NOT.
	//
	//    * Key - The name of the parameter to filter on. Valid values are ForecastArn
	//    and Status.
	//
	//    * Value - The value to match.
	//
	// For example, to list all jobs that export a forecast named electricityforecast,
	// specify the following filter:
	//
	// "Filters": [ { "Condition": "IS", "Key": "ForecastArn", "Value": "arn:aws:forecast:us-west-2:<acct-id>:forecast/electricityforecast"
	// } ]
	Filters []Filter `type:"list"`

	// The number of items to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the result of the previous request was truncated, the response includes
	// a NextToken. To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListForecastExportJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListForecastExportJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListForecastExportJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
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

type ListForecastExportJobsOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that summarize each export job's properties.
	ForecastExportJobs []ForecastExportJobSummary `type:"list"`

	// If the response is truncated, Amazon Forecast returns this token. To retrieve
	// the next set of results, use the token in the next request.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListForecastExportJobsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListForecastExportJobs = "ListForecastExportJobs"

// ListForecastExportJobsRequest returns a request value for making API operation for
// Amazon Forecast Service.
//
// Returns a list of forecast export jobs created using the CreateForecastExportJob
// operation. For each forecast export job, this operation returns a summary
// of its properties, including its Amazon Resource Name (ARN). To retrieve
// the complete set of properties, use the ARN with the DescribeForecastExportJob
// operation. You can filter the list using an array of Filter objects.
//
//    // Example sending a request using ListForecastExportJobsRequest.
//    req := client.ListForecastExportJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/forecast-2018-06-26/ListForecastExportJobs
func (c *Client) ListForecastExportJobsRequest(input *ListForecastExportJobsInput) ListForecastExportJobsRequest {
	op := &aws.Operation{
		Name:       opListForecastExportJobs,
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
		input = &ListForecastExportJobsInput{}
	}

	req := c.newRequest(op, input, &ListForecastExportJobsOutput{})

	return ListForecastExportJobsRequest{Request: req, Input: input, Copy: c.ListForecastExportJobsRequest}
}

// ListForecastExportJobsRequest is the request type for the
// ListForecastExportJobs API operation.
type ListForecastExportJobsRequest struct {
	*aws.Request
	Input *ListForecastExportJobsInput
	Copy  func(*ListForecastExportJobsInput) ListForecastExportJobsRequest
}

// Send marshals and sends the ListForecastExportJobs API request.
func (r ListForecastExportJobsRequest) Send(ctx context.Context) (*ListForecastExportJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListForecastExportJobsResponse{
		ListForecastExportJobsOutput: r.Request.Data.(*ListForecastExportJobsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListForecastExportJobsRequestPaginator returns a paginator for ListForecastExportJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListForecastExportJobsRequest(input)
//   p := forecast.NewListForecastExportJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListForecastExportJobsPaginator(req ListForecastExportJobsRequest) ListForecastExportJobsPaginator {
	return ListForecastExportJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListForecastExportJobsInput
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

// ListForecastExportJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListForecastExportJobsPaginator struct {
	aws.Pager
}

func (p *ListForecastExportJobsPaginator) CurrentPage() *ListForecastExportJobsOutput {
	return p.Pager.CurrentPage().(*ListForecastExportJobsOutput)
}

// ListForecastExportJobsResponse is the response type for the
// ListForecastExportJobs API operation.
type ListForecastExportJobsResponse struct {
	*ListForecastExportJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListForecastExportJobs request.
func (r *ListForecastExportJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
