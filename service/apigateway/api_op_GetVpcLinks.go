// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// Gets the VpcLinks collection under the caller's account in a selected region.
type GetVpcLinksInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of returned results per page. The default value is 25
	// and the maximum value is 500.
	Limit *int64 `location:"querystring" locationName:"limit" type:"integer"`

	// The current pagination position in the paged result set.
	Position *string `location:"querystring" locationName:"position" type:"string"`
}

// String returns the string representation
func (s GetVpcLinksInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetVpcLinksInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

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

// The collection of VPC links under the caller's account in a region.
//
// Getting Started with Private Integrations (https://docs.aws.amazon.com/apigateway/latest/developerguide/getting-started-with-private-integration.html),
// Set up Private Integrations (https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-private-integration.html)
type GetVpcLinksOutput struct {
	_ struct{} `type:"structure"`

	// The current page of elements from this collection.
	Items []VpcLink `locationName:"item" type:"list"`

	Position *string `locationName:"position" type:"string"`
}

// String returns the string representation
func (s GetVpcLinksOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetVpcLinksOutput) MarshalFields(e protocol.FieldEncoder) error {
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

const opGetVpcLinks = "GetVpcLinks"

// GetVpcLinksRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Gets the VpcLinks collection under the caller's account in a selected region.
//
//    // Example sending a request using GetVpcLinksRequest.
//    req := client.GetVpcLinksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetVpcLinksRequest(input *GetVpcLinksInput) GetVpcLinksRequest {
	op := &aws.Operation{
		Name:       opGetVpcLinks,
		HTTPMethod: "GET",
		HTTPPath:   "/vpclinks",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"position"},
			OutputTokens:    []string{"position"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetVpcLinksInput{}
	}

	req := c.newRequest(op, input, &GetVpcLinksOutput{})

	return GetVpcLinksRequest{Request: req, Input: input, Copy: c.GetVpcLinksRequest}
}

// GetVpcLinksRequest is the request type for the
// GetVpcLinks API operation.
type GetVpcLinksRequest struct {
	*aws.Request
	Input *GetVpcLinksInput
	Copy  func(*GetVpcLinksInput) GetVpcLinksRequest
}

// Send marshals and sends the GetVpcLinks API request.
func (r GetVpcLinksRequest) Send(ctx context.Context) (*GetVpcLinksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetVpcLinksResponse{
		GetVpcLinksOutput: r.Request.Data.(*GetVpcLinksOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetVpcLinksRequestPaginator returns a paginator for GetVpcLinks.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetVpcLinksRequest(input)
//   p := apigateway.NewGetVpcLinksRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetVpcLinksPaginator(req GetVpcLinksRequest) GetVpcLinksPaginator {
	return GetVpcLinksPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetVpcLinksInput
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

// GetVpcLinksPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetVpcLinksPaginator struct {
	aws.Pager
}

func (p *GetVpcLinksPaginator) CurrentPage() *GetVpcLinksOutput {
	return p.Pager.CurrentPage().(*GetVpcLinksOutput)
}

// GetVpcLinksResponse is the response type for the
// GetVpcLinks API operation.
type GetVpcLinksResponse struct {
	*GetVpcLinksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetVpcLinks request.
func (r *GetVpcLinksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
