// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationinsights

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListApplicationsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token to request the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListApplicationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListApplicationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListApplicationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListApplicationsOutput struct {
	_ struct{} `type:"structure"`

	// The list of applications.
	ApplicationInfoList []ApplicationInfo `type:"list"`

	// The token used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListApplicationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListApplications = "ListApplications"

// ListApplicationsRequest returns a request value for making API operation for
// Amazon CloudWatch Application Insights.
//
// Lists the IDs of the applications that you are monitoring.
//
//    // Example sending a request using ListApplicationsRequest.
//    req := client.ListApplicationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-insights-2018-11-25/ListApplications
func (c *Client) ListApplicationsRequest(input *ListApplicationsInput) ListApplicationsRequest {
	op := &aws.Operation{
		Name:       opListApplications,
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
		input = &ListApplicationsInput{}
	}

	req := c.newRequest(op, input, &ListApplicationsOutput{})

	return ListApplicationsRequest{Request: req, Input: input, Copy: c.ListApplicationsRequest}
}

// ListApplicationsRequest is the request type for the
// ListApplications API operation.
type ListApplicationsRequest struct {
	*aws.Request
	Input *ListApplicationsInput
	Copy  func(*ListApplicationsInput) ListApplicationsRequest
}

// Send marshals and sends the ListApplications API request.
func (r ListApplicationsRequest) Send(ctx context.Context) (*ListApplicationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListApplicationsResponse{
		ListApplicationsOutput: r.Request.Data.(*ListApplicationsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListApplicationsRequestPaginator returns a paginator for ListApplications.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListApplicationsRequest(input)
//   p := applicationinsights.NewListApplicationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListApplicationsPaginator(req ListApplicationsRequest) ListApplicationsPaginator {
	return ListApplicationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListApplicationsInput
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

// ListApplicationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListApplicationsPaginator struct {
	aws.Pager
}

func (p *ListApplicationsPaginator) CurrentPage() *ListApplicationsOutput {
	return p.Pager.CurrentPage().(*ListApplicationsOutput)
}

// ListApplicationsResponse is the response type for the
// ListApplications API operation.
type ListApplicationsResponse struct {
	*ListApplicationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListApplications request.
func (r *ListApplicationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
