// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListRobotsInput struct {
	_ struct{} `type:"structure"`

	// Optional filters to limit results.
	//
	// The filter names status and fleetName are supported. When filtering, you
	// must use the complete value of the filtered item. You can use up to three
	// filters, but they must be for the same named item. For example, if you are
	// looking for items with the status Registered or the status Available.
	Filters []Filter `locationName:"filters" min:"1" type:"list"`

	// When this parameter is used, ListRobots only returns maxResults results in
	// a single page along with a nextToken response element. The remaining results
	// of the initial request can be seen by sending another ListRobots request
	// with the returned nextToken value. This value can be between 1 and 200. If
	// this parameter is not used, then ListRobots returns up to 200 results and
	// a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListRobots request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListRobotsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRobotsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRobotsInput"}
	if s.Filters != nil && len(s.Filters) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Filters", 1))
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

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListRobotsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Filters != nil {
		v := s.Filters

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "filters", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListRobotsOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken value to include in a future ListRobots request. When the results
	// of a ListRobot request exceed maxResults, this value can be used to retrieve
	// the next page of results. This value is null when there are no more results
	// to return.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// A list of robots that meet the criteria of the request.
	Robots []Robot `locationName:"robots" type:"list"`
}

// String returns the string representation
func (s ListRobotsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListRobotsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Robots != nil {
		v := s.Robots

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "robots", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListRobots = "ListRobots"

// ListRobotsRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Returns a list of robots. You can optionally provide filters to retrieve
// specific robots.
//
//    // Example sending a request using ListRobotsRequest.
//    req := client.ListRobotsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/ListRobots
func (c *Client) ListRobotsRequest(input *ListRobotsInput) ListRobotsRequest {
	op := &aws.Operation{
		Name:       opListRobots,
		HTTPMethod: "POST",
		HTTPPath:   "/listRobots",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListRobotsInput{}
	}

	req := c.newRequest(op, input, &ListRobotsOutput{})

	return ListRobotsRequest{Request: req, Input: input, Copy: c.ListRobotsRequest}
}

// ListRobotsRequest is the request type for the
// ListRobots API operation.
type ListRobotsRequest struct {
	*aws.Request
	Input *ListRobotsInput
	Copy  func(*ListRobotsInput) ListRobotsRequest
}

// Send marshals and sends the ListRobots API request.
func (r ListRobotsRequest) Send(ctx context.Context) (*ListRobotsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRobotsResponse{
		ListRobotsOutput: r.Request.Data.(*ListRobotsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListRobotsRequestPaginator returns a paginator for ListRobots.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListRobotsRequest(input)
//   p := robomaker.NewListRobotsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListRobotsPaginator(req ListRobotsRequest) ListRobotsPaginator {
	return ListRobotsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListRobotsInput
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

// ListRobotsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListRobotsPaginator struct {
	aws.Pager
}

func (p *ListRobotsPaginator) CurrentPage() *ListRobotsOutput {
	return p.Pager.CurrentPage().(*ListRobotsOutput)
}

// ListRobotsResponse is the response type for the
// ListRobots API operation.
type ListRobotsResponse struct {
	*ListRobotsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRobots request.
func (r *ListRobotsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
