// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListDashboardsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to be returned per paginated request.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token to be used for the next set of paginated results.
	NextToken *string `location:"querystring" locationName:"nextToken" min:"1" type:"string"`

	// The ID of the project.
	//
	// ProjectId is a required field
	ProjectId *string `location:"querystring" locationName:"projectId" min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s ListDashboardsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDashboardsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDashboardsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.ProjectId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectId"))
	}
	if s.ProjectId != nil && len(*s.ProjectId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDashboardsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProjectId != nil {
		v := *s.ProjectId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "projectId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListDashboardsOutput struct {
	_ struct{} `type:"structure"`

	// A list that summarizes each dashboard in the project.
	//
	// DashboardSummaries is a required field
	DashboardSummaries []DashboardSummary `locationName:"dashboardSummaries" type:"list" required:"true"`

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListDashboardsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDashboardsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DashboardSummaries != nil {
		v := s.DashboardSummaries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "dashboardSummaries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListDashboards = "ListDashboards"

// ListDashboardsRequest returns a request value for making API operation for
// AWS IoT SiteWise.
//
// Retrieves a paginated list of dashboards for an AWS IoT SiteWise Monitor
// project.
//
//    // Example sending a request using ListDashboardsRequest.
//    req := client.ListDashboardsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotsitewise-2019-12-02/ListDashboards
func (c *Client) ListDashboardsRequest(input *ListDashboardsInput) ListDashboardsRequest {
	op := &aws.Operation{
		Name:       opListDashboards,
		HTTPMethod: "GET",
		HTTPPath:   "/dashboards",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDashboardsInput{}
	}

	req := c.newRequest(op, input, &ListDashboardsOutput{})
	req.Handlers.Build.PushBackNamed(protocol.NewHostPrefixHandler("monitor.", nil))
	req.Handlers.Build.PushBackNamed(protocol.ValidateEndpointHostHandler)

	return ListDashboardsRequest{Request: req, Input: input, Copy: c.ListDashboardsRequest}
}

// ListDashboardsRequest is the request type for the
// ListDashboards API operation.
type ListDashboardsRequest struct {
	*aws.Request
	Input *ListDashboardsInput
	Copy  func(*ListDashboardsInput) ListDashboardsRequest
}

// Send marshals and sends the ListDashboards API request.
func (r ListDashboardsRequest) Send(ctx context.Context) (*ListDashboardsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDashboardsResponse{
		ListDashboardsOutput: r.Request.Data.(*ListDashboardsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDashboardsRequestPaginator returns a paginator for ListDashboards.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDashboardsRequest(input)
//   p := iotsitewise.NewListDashboardsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDashboardsPaginator(req ListDashboardsRequest) ListDashboardsPaginator {
	return ListDashboardsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDashboardsInput
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

// ListDashboardsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDashboardsPaginator struct {
	aws.Pager
}

func (p *ListDashboardsPaginator) CurrentPage() *ListDashboardsOutput {
	return p.Pager.CurrentPage().(*ListDashboardsOutput)
}

// ListDashboardsResponse is the response type for the
// ListDashboards API operation.
type ListDashboardsResponse struct {
	*ListDashboardsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDashboards request.
func (r *ListDashboardsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
