// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// Request to list information about a collection of resources.
type GetResourcesInput struct {
	_ struct{} `type:"structure"`

	// A query parameter used to retrieve the specified resources embedded in the
	// returned Resources resource in the response. This embed parameter value is
	// a list of comma-separated strings. Currently, the request supports only retrieval
	// of the embedded Method resources this way. The query parameter value must
	// be a single-valued list and contain the "methods" string. For example, GET
	// /restapis/{restapi_id}/resources?embed=methods.
	Embed []string `location:"querystring" locationName:"embed" type:"list"`

	// The maximum number of returned results per page. The default value is 25
	// and the maximum value is 500.
	Limit *int64 `location:"querystring" locationName:"limit" type:"integer"`

	// The current pagination position in the paged result set.
	Position *string `location:"querystring" locationName:"position" type:"string"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetResourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetResourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetResourcesInput"}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetResourcesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Embed != nil {
		v := s.Embed

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.QueryTarget, "embed", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.Position != nil {
		v := *s.Position

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "position", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents a collection of Resource resources.
//
// Create an API (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
type GetResourcesOutput struct {
	_ struct{} `type:"structure"`

	// The current page of elements from this collection.
	Items []Resource `locationName:"item" type:"list"`

	Position *string `locationName:"position" type:"string"`
}

// String returns the string representation
func (s GetResourcesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetResourcesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Items != nil {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "item", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Position != nil {
		v := *s.Position

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "position", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetResources = "GetResources"

// GetResourcesRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Lists information about a collection of Resource resources.
//
//    // Example sending a request using GetResourcesRequest.
//    req := client.GetResourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetResourcesRequest(input *GetResourcesInput) GetResourcesRequest {
	op := &aws.Operation{
		Name:       opGetResources,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis/{restapi_id}/resources",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"position"},
			OutputTokens:    []string{"position"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetResourcesInput{}
	}

	req := c.newRequest(op, input, &GetResourcesOutput{})

	return GetResourcesRequest{Request: req, Input: input, Copy: c.GetResourcesRequest}
}

// GetResourcesRequest is the request type for the
// GetResources API operation.
type GetResourcesRequest struct {
	*aws.Request
	Input *GetResourcesInput
	Copy  func(*GetResourcesInput) GetResourcesRequest
}

// Send marshals and sends the GetResources API request.
func (r GetResourcesRequest) Send(ctx context.Context) (*GetResourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetResourcesResponse{
		GetResourcesOutput: r.Request.Data.(*GetResourcesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetResourcesRequestPaginator returns a paginator for GetResources.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetResourcesRequest(input)
//   p := apigateway.NewGetResourcesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetResourcesPaginator(req GetResourcesRequest) GetResourcesPaginator {
	return GetResourcesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetResourcesInput
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

// GetResourcesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetResourcesPaginator struct {
	aws.Pager
}

func (p *GetResourcesPaginator) CurrentPage() *GetResourcesOutput {
	return p.Pager.CurrentPage().(*GetResourcesOutput)
}

// GetResourcesResponse is the response type for the
// GetResources API operation.
type GetResourcesResponse struct {
	*GetResourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetResources request.
func (r *GetResourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
