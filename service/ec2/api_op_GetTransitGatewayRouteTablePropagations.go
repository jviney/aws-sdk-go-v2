// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetTransitGatewayRouteTablePropagationsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters. The possible values are:
	//
	//    * resource-id - The ID of the resource.
	//
	//    * resource-type - The resource type (vpc | vpn).
	//
	//    * transit-gateway-attachment-id - The ID of the attachment.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`

	// The ID of the transit gateway route table.
	//
	// TransitGatewayRouteTableId is a required field
	TransitGatewayRouteTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetTransitGatewayRouteTablePropagationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTransitGatewayRouteTablePropagationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTransitGatewayRouteTablePropagationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if s.TransitGatewayRouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayRouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetTransitGatewayRouteTablePropagationsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the route table propagations.
	TransitGatewayRouteTablePropagations []TransitGatewayRouteTablePropagation `locationName:"transitGatewayRouteTablePropagations" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s GetTransitGatewayRouteTablePropagationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetTransitGatewayRouteTablePropagations = "GetTransitGatewayRouteTablePropagations"

// GetTransitGatewayRouteTablePropagationsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Gets information about the route table propagations for the specified transit
// gateway route table.
//
//    // Example sending a request using GetTransitGatewayRouteTablePropagationsRequest.
//    req := client.GetTransitGatewayRouteTablePropagationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetTransitGatewayRouteTablePropagations
func (c *Client) GetTransitGatewayRouteTablePropagationsRequest(input *GetTransitGatewayRouteTablePropagationsInput) GetTransitGatewayRouteTablePropagationsRequest {
	op := &aws.Operation{
		Name:       opGetTransitGatewayRouteTablePropagations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetTransitGatewayRouteTablePropagationsInput{}
	}

	req := c.newRequest(op, input, &GetTransitGatewayRouteTablePropagationsOutput{})

	return GetTransitGatewayRouteTablePropagationsRequest{Request: req, Input: input, Copy: c.GetTransitGatewayRouteTablePropagationsRequest}
}

// GetTransitGatewayRouteTablePropagationsRequest is the request type for the
// GetTransitGatewayRouteTablePropagations API operation.
type GetTransitGatewayRouteTablePropagationsRequest struct {
	*aws.Request
	Input *GetTransitGatewayRouteTablePropagationsInput
	Copy  func(*GetTransitGatewayRouteTablePropagationsInput) GetTransitGatewayRouteTablePropagationsRequest
}

// Send marshals and sends the GetTransitGatewayRouteTablePropagations API request.
func (r GetTransitGatewayRouteTablePropagationsRequest) Send(ctx context.Context) (*GetTransitGatewayRouteTablePropagationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTransitGatewayRouteTablePropagationsResponse{
		GetTransitGatewayRouteTablePropagationsOutput: r.Request.Data.(*GetTransitGatewayRouteTablePropagationsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetTransitGatewayRouteTablePropagationsRequestPaginator returns a paginator for GetTransitGatewayRouteTablePropagations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetTransitGatewayRouteTablePropagationsRequest(input)
//   p := ec2.NewGetTransitGatewayRouteTablePropagationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetTransitGatewayRouteTablePropagationsPaginator(req GetTransitGatewayRouteTablePropagationsRequest) GetTransitGatewayRouteTablePropagationsPaginator {
	return GetTransitGatewayRouteTablePropagationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetTransitGatewayRouteTablePropagationsInput
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

// GetTransitGatewayRouteTablePropagationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetTransitGatewayRouteTablePropagationsPaginator struct {
	aws.Pager
}

func (p *GetTransitGatewayRouteTablePropagationsPaginator) CurrentPage() *GetTransitGatewayRouteTablePropagationsOutput {
	return p.Pager.CurrentPage().(*GetTransitGatewayRouteTablePropagationsOutput)
}

// GetTransitGatewayRouteTablePropagationsResponse is the response type for the
// GetTransitGatewayRouteTablePropagations API operation.
type GetTransitGatewayRouteTablePropagationsResponse struct {
	*GetTransitGatewayRouteTablePropagationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTransitGatewayRouteTablePropagations request.
func (r *GetTransitGatewayRouteTablePropagationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
