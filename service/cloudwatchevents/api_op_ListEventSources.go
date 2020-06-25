// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchevents

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListEventSourcesInput struct {
	_ struct{} `type:"structure"`

	// Specifying this limits the number of results returned by this operation.
	// The operation also returns a NextToken which you can use in a subsequent
	// operation to retrieve the next set of results.
	Limit *int64 `min:"1" type:"integer"`

	// Specifying this limits the results to only those partner event sources with
	// names that start with the specified prefix.
	NamePrefix *string `min:"1" type:"string"`

	// The token returned by a previous call to retrieve the next set of results.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListEventSourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListEventSourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListEventSourcesInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.NamePrefix != nil && len(*s.NamePrefix) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NamePrefix", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListEventSourcesOutput struct {
	_ struct{} `type:"structure"`

	// The list of event sources.
	EventSources []EventSource `type:"list"`

	// A token you can use in a subsequent operation to retrieve the next set of
	// results.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListEventSourcesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListEventSources = "ListEventSources"

// ListEventSourcesRequest returns a request value for making API operation for
// Amazon CloudWatch Events.
//
// You can use this to see all the partner event sources that have been shared
// with your AWS account. For more information about partner event sources,
// see CreateEventBus.
//
//    // Example sending a request using ListEventSourcesRequest.
//    req := client.ListEventSourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/events-2015-10-07/ListEventSources
func (c *Client) ListEventSourcesRequest(input *ListEventSourcesInput) ListEventSourcesRequest {
	op := &aws.Operation{
		Name:       opListEventSources,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListEventSourcesInput{}
	}

	req := c.newRequest(op, input, &ListEventSourcesOutput{})

	return ListEventSourcesRequest{Request: req, Input: input, Copy: c.ListEventSourcesRequest}
}

// ListEventSourcesRequest is the request type for the
// ListEventSources API operation.
type ListEventSourcesRequest struct {
	*aws.Request
	Input *ListEventSourcesInput
	Copy  func(*ListEventSourcesInput) ListEventSourcesRequest
}

// Send marshals and sends the ListEventSources API request.
func (r ListEventSourcesRequest) Send(ctx context.Context) (*ListEventSourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListEventSourcesResponse{
		ListEventSourcesOutput: r.Request.Data.(*ListEventSourcesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListEventSourcesResponse is the response type for the
// ListEventSources API operation.
type ListEventSourcesResponse struct {
	*ListEventSourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListEventSources request.
func (r *ListEventSourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
