// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestarnotifications

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListEventTypesInput struct {
	_ struct{} `type:"structure"`

	// The filters to use to return information by service or resource type.
	Filters []ListEventTypesFilter `type:"list"`

	// A non-negative integer used to limit the number of returned results. The
	// default number is 50. The maximum number of results that can be returned
	// is 100.
	MaxResults *int64 `min:"1" type:"integer"`

	// An enumeration token that, when provided in a request, returns the next batch
	// of the results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListEventTypesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListEventTypesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListEventTypesInput"}
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

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListEventTypesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Filters != nil {
		v := s.Filters

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Filters", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListEventTypesOutput struct {
	_ struct{} `type:"structure"`

	// Information about each event, including service name, resource type, event
	// ID, and event name.
	EventTypes []EventTypeSummary `type:"list"`

	// An enumeration token that can be used in a request to return the next batch
	// of the results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListEventTypesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListEventTypesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.EventTypes != nil {
		v := s.EventTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "EventTypes", metadata)
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

const opListEventTypes = "ListEventTypes"

// ListEventTypesRequest returns a request value for making API operation for
// AWS CodeStar Notifications.
//
// Returns information about the event types available for configuring notifications.
//
//    // Example sending a request using ListEventTypesRequest.
//    req := client.ListEventTypesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-notifications-2019-10-15/ListEventTypes
func (c *Client) ListEventTypesRequest(input *ListEventTypesInput) ListEventTypesRequest {
	op := &aws.Operation{
		Name:       opListEventTypes,
		HTTPMethod: "POST",
		HTTPPath:   "/listEventTypes",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListEventTypesInput{}
	}

	req := c.newRequest(op, input, &ListEventTypesOutput{})

	return ListEventTypesRequest{Request: req, Input: input, Copy: c.ListEventTypesRequest}
}

// ListEventTypesRequest is the request type for the
// ListEventTypes API operation.
type ListEventTypesRequest struct {
	*aws.Request
	Input *ListEventTypesInput
	Copy  func(*ListEventTypesInput) ListEventTypesRequest
}

// Send marshals and sends the ListEventTypes API request.
func (r ListEventTypesRequest) Send(ctx context.Context) (*ListEventTypesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListEventTypesResponse{
		ListEventTypesOutput: r.Request.Data.(*ListEventTypesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListEventTypesRequestPaginator returns a paginator for ListEventTypes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListEventTypesRequest(input)
//   p := codestarnotifications.NewListEventTypesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListEventTypesPaginator(req ListEventTypesRequest) ListEventTypesPaginator {
	return ListEventTypesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListEventTypesInput
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

// ListEventTypesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListEventTypesPaginator struct {
	aws.Pager
}

func (p *ListEventTypesPaginator) CurrentPage() *ListEventTypesOutput {
	return p.Pager.CurrentPage().(*ListEventTypesOutput)
}

// ListEventTypesResponse is the response type for the
// ListEventTypes API operation.
type ListEventTypesResponse struct {
	*ListEventTypesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListEventTypes request.
func (r *ListEventTypesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
