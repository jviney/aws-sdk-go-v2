// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListContactFlowsInput struct {
	_ struct{} `type:"structure"`

	// The type of contact flow.
	ContactFlowTypes []ContactFlowType `location:"querystring" locationName:"contactFlowTypes" type:"list"`

	// The identifier of the Amazon Connect instance.
	//
	// InstanceId is a required field
	InstanceId *string `location:"uri" locationName:"InstanceId" min:"1" type:"string" required:"true"`

	// The maximimum number of results to return per page.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListContactFlowsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListContactFlowsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListContactFlowsInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListContactFlowsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.InstanceId != nil {
		v := *s.InstanceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "InstanceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ContactFlowTypes != nil {
		v := s.ContactFlowTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.QueryTarget, "contactFlowTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
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

type ListContactFlowsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the contact flows.
	ContactFlowSummaryList []ContactFlowSummary `type:"list"`

	// If there are additional results, this is the token for the next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListContactFlowsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListContactFlowsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ContactFlowSummaryList != nil {
		v := s.ContactFlowSummaryList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ContactFlowSummaryList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListContactFlows = "ListContactFlows"

// ListContactFlowsRequest returns a request value for making API operation for
// Amazon Connect Service.
//
// Provides information about the contact flows for the specified Amazon Connect
// instance.
//
//    // Example sending a request using ListContactFlowsRequest.
//    req := client.ListContactFlowsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/ListContactFlows
func (c *Client) ListContactFlowsRequest(input *ListContactFlowsInput) ListContactFlowsRequest {
	op := &aws.Operation{
		Name:       opListContactFlows,
		HTTPMethod: "GET",
		HTTPPath:   "/contact-flows-summary/{InstanceId}",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListContactFlowsInput{}
	}

	req := c.newRequest(op, input, &ListContactFlowsOutput{})

	return ListContactFlowsRequest{Request: req, Input: input, Copy: c.ListContactFlowsRequest}
}

// ListContactFlowsRequest is the request type for the
// ListContactFlows API operation.
type ListContactFlowsRequest struct {
	*aws.Request
	Input *ListContactFlowsInput
	Copy  func(*ListContactFlowsInput) ListContactFlowsRequest
}

// Send marshals and sends the ListContactFlows API request.
func (r ListContactFlowsRequest) Send(ctx context.Context) (*ListContactFlowsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListContactFlowsResponse{
		ListContactFlowsOutput: r.Request.Data.(*ListContactFlowsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListContactFlowsRequestPaginator returns a paginator for ListContactFlows.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListContactFlowsRequest(input)
//   p := connect.NewListContactFlowsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListContactFlowsPaginator(req ListContactFlowsRequest) ListContactFlowsPaginator {
	return ListContactFlowsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListContactFlowsInput
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

// ListContactFlowsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListContactFlowsPaginator struct {
	aws.Pager
}

func (p *ListContactFlowsPaginator) CurrentPage() *ListContactFlowsOutput {
	return p.Pager.CurrentPage().(*ListContactFlowsOutput)
}

// ListContactFlowsResponse is the response type for the
// ListContactFlows API operation.
type ListContactFlowsResponse struct {
	*ListContactFlowsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListContactFlows request.
func (r *ListContactFlowsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
