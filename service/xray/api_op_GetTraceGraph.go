// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package xray

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetTraceGraphInput struct {
	_ struct{} `type:"structure"`

	// Pagination token.
	NextToken *string `type:"string"`

	// Trace IDs of requests for which to generate a service graph.
	//
	// TraceIds is a required field
	TraceIds []string `type:"list" required:"true"`
}

// String returns the string representation
func (s GetTraceGraphInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTraceGraphInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTraceGraphInput"}

	if s.TraceIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("TraceIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetTraceGraphInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TraceIds != nil {
		v := s.TraceIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "TraceIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

type GetTraceGraphOutput struct {
	_ struct{} `type:"structure"`

	// Pagination token.
	NextToken *string `type:"string"`

	// The services that have processed one of the specified requests.
	Services []Service `type:"list"`
}

// String returns the string representation
func (s GetTraceGraphOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetTraceGraphOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Services != nil {
		v := s.Services

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Services", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetTraceGraph = "GetTraceGraph"

// GetTraceGraphRequest returns a request value for making API operation for
// AWS X-Ray.
//
// Retrieves a service graph for one or more specific trace IDs.
//
//    // Example sending a request using GetTraceGraphRequest.
//    req := client.GetTraceGraphRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/xray-2016-04-12/GetTraceGraph
func (c *Client) GetTraceGraphRequest(input *GetTraceGraphInput) GetTraceGraphRequest {
	op := &aws.Operation{
		Name:       opGetTraceGraph,
		HTTPMethod: "POST",
		HTTPPath:   "/TraceGraph",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetTraceGraphInput{}
	}

	req := c.newRequest(op, input, &GetTraceGraphOutput{})

	return GetTraceGraphRequest{Request: req, Input: input, Copy: c.GetTraceGraphRequest}
}

// GetTraceGraphRequest is the request type for the
// GetTraceGraph API operation.
type GetTraceGraphRequest struct {
	*aws.Request
	Input *GetTraceGraphInput
	Copy  func(*GetTraceGraphInput) GetTraceGraphRequest
}

// Send marshals and sends the GetTraceGraph API request.
func (r GetTraceGraphRequest) Send(ctx context.Context) (*GetTraceGraphResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTraceGraphResponse{
		GetTraceGraphOutput: r.Request.Data.(*GetTraceGraphOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetTraceGraphRequestPaginator returns a paginator for GetTraceGraph.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetTraceGraphRequest(input)
//   p := xray.NewGetTraceGraphRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetTraceGraphPaginator(req GetTraceGraphRequest) GetTraceGraphPaginator {
	return GetTraceGraphPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetTraceGraphInput
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

// GetTraceGraphPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetTraceGraphPaginator struct {
	aws.Pager
}

func (p *GetTraceGraphPaginator) CurrentPage() *GetTraceGraphOutput {
	return p.Pager.CurrentPage().(*GetTraceGraphOutput)
}

// GetTraceGraphResponse is the response type for the
// GetTraceGraph API operation.
type GetTraceGraphResponse struct {
	*GetTraceGraphOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTraceGraph request.
func (r *GetTraceGraphResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
