// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot1clickprojects

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListProjectsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return per request. If not set, a default
	// value of 100 is used.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListProjectsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListProjectsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListProjectsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListProjectsInput) MarshalFields(e protocol.FieldEncoder) error {
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
	return nil
}

type ListProjectsOutput struct {
	_ struct{} `type:"structure"`

	// The token used to retrieve the next set of results - will be effectively
	// empty if there are no further results.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// An object containing the list of projects.
	//
	// Projects is a required field
	Projects []ProjectSummary `locationName:"projects" type:"list" required:"true"`
}

// String returns the string representation
func (s ListProjectsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListProjectsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Projects != nil {
		v := s.Projects

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "projects", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListProjects = "ListProjects"

// ListProjectsRequest returns a request value for making API operation for
// AWS IoT 1-Click Projects Service.
//
// Lists the AWS IoT 1-Click project(s) associated with your AWS account and
// region.
//
//    // Example sending a request using ListProjectsRequest.
//    req := client.ListProjectsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/ListProjects
func (c *Client) ListProjectsRequest(input *ListProjectsInput) ListProjectsRequest {
	op := &aws.Operation{
		Name:       opListProjects,
		HTTPMethod: "GET",
		HTTPPath:   "/projects",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListProjectsInput{}
	}

	req := c.newRequest(op, input, &ListProjectsOutput{})

	return ListProjectsRequest{Request: req, Input: input, Copy: c.ListProjectsRequest}
}

// ListProjectsRequest is the request type for the
// ListProjects API operation.
type ListProjectsRequest struct {
	*aws.Request
	Input *ListProjectsInput
	Copy  func(*ListProjectsInput) ListProjectsRequest
}

// Send marshals and sends the ListProjects API request.
func (r ListProjectsRequest) Send(ctx context.Context) (*ListProjectsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListProjectsResponse{
		ListProjectsOutput: r.Request.Data.(*ListProjectsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListProjectsRequestPaginator returns a paginator for ListProjects.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListProjectsRequest(input)
//   p := iot1clickprojects.NewListProjectsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListProjectsPaginator(req ListProjectsRequest) ListProjectsPaginator {
	return ListProjectsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListProjectsInput
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

// ListProjectsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListProjectsPaginator struct {
	aws.Pager
}

func (p *ListProjectsPaginator) CurrentPage() *ListProjectsOutput {
	return p.Pager.CurrentPage().(*ListProjectsOutput)
}

// ListProjectsResponse is the response type for the
// ListProjects API operation.
type ListProjectsResponse struct {
	*ListProjectsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListProjects request.
func (r *ListProjectsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
