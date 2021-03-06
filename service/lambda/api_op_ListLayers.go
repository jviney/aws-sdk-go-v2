// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListLayersInput struct {
	_ struct{} `type:"structure"`

	// A runtime identifier. For example, go1.x.
	CompatibleRuntime Runtime `location:"querystring" locationName:"CompatibleRuntime" type:"string" enum:"true"`

	// A pagination token returned by a previous call.
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// The maximum number of layers to return.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" min:"1" type:"integer"`
}

// String returns the string representation
func (s ListLayersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListLayersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListLayersInput"}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListLayersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if len(s.CompatibleRuntime) > 0 {
		v := s.CompatibleRuntime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "CompatibleRuntime", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxItems", protocol.Int64Value(v), metadata)
	}
	return nil
}

type ListLayersOutput struct {
	_ struct{} `type:"structure"`

	// A list of function layers.
	Layers []LayersListItem `type:"list"`

	// A pagination token returned when the response doesn't contain all layers.
	NextMarker *string `type:"string"`
}

// String returns the string representation
func (s ListLayersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListLayersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Layers != nil {
		v := s.Layers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Layers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextMarker != nil {
		v := *s.NextMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextMarker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListLayers = "ListLayers"

// ListLayersRequest returns a request value for making API operation for
// AWS Lambda.
//
// Lists AWS Lambda layers (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
// and shows information about the latest version of each. Specify a runtime
// identifier (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html)
// to list only layers that indicate that they're compatible with that runtime.
//
//    // Example sending a request using ListLayersRequest.
//    req := client.ListLayersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/ListLayers
func (c *Client) ListLayersRequest(input *ListLayersInput) ListLayersRequest {
	op := &aws.Operation{
		Name:       opListLayers,
		HTTPMethod: "GET",
		HTTPPath:   "/2018-10-31/layers",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"NextMarker"},
			LimitToken:      "MaxItems",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListLayersInput{}
	}

	req := c.newRequest(op, input, &ListLayersOutput{})

	return ListLayersRequest{Request: req, Input: input, Copy: c.ListLayersRequest}
}

// ListLayersRequest is the request type for the
// ListLayers API operation.
type ListLayersRequest struct {
	*aws.Request
	Input *ListLayersInput
	Copy  func(*ListLayersInput) ListLayersRequest
}

// Send marshals and sends the ListLayers API request.
func (r ListLayersRequest) Send(ctx context.Context) (*ListLayersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListLayersResponse{
		ListLayersOutput: r.Request.Data.(*ListLayersOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListLayersRequestPaginator returns a paginator for ListLayers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListLayersRequest(input)
//   p := lambda.NewListLayersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListLayersPaginator(req ListLayersRequest) ListLayersPaginator {
	return ListLayersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListLayersInput
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

// ListLayersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListLayersPaginator struct {
	aws.Pager
}

func (p *ListLayersPaginator) CurrentPage() *ListLayersOutput {
	return p.Pager.CurrentPage().(*ListLayersOutput)
}

// ListLayersResponse is the response type for the
// ListLayers API operation.
type ListLayersResponse struct {
	*ListLayersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListLayers request.
func (r *ListLayersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
