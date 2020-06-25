// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// The request to list your distributions.
type ListDistributionsInput struct {
	_ struct{} `type:"structure"`

	// Use this when paginating results to indicate where to begin in your list
	// of distributions. The results include distributions in the list that occur
	// after the marker. To get the next page of results, set the Marker to the
	// value of the NextMarker from the current page's response (which is also the
	// ID of the last distribution on that page).
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// The maximum number of distributions you want in the response body.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" type:"integer"`
}

// String returns the string representation
func (s ListDistributionsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDistributionsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Marker", protocol.StringValue(v), metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxItems", protocol.Int64Value(v), metadata)
	}
	return nil
}

// The returned result of the corresponding request.
type ListDistributionsOutput struct {
	_ struct{} `type:"structure" payload:"DistributionList"`

	// The DistributionList type.
	DistributionList *DistributionList `type:"structure"`
}

// String returns the string representation
func (s ListDistributionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDistributionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DistributionList != nil {
		v := s.DistributionList

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "DistributionList", v, metadata)
	}
	return nil
}

const opListDistributions = "ListDistributions2019_03_26"

// ListDistributionsRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// List CloudFront distributions.
//
//    // Example sending a request using ListDistributionsRequest.
//    req := client.ListDistributionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/ListDistributions
func (c *Client) ListDistributionsRequest(input *ListDistributionsInput) ListDistributionsRequest {
	op := &aws.Operation{
		Name:       opListDistributions,
		HTTPMethod: "GET",
		HTTPPath:   "/2019-03-26/distribution",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"DistributionList.NextMarker"},
			LimitToken:      "MaxItems",
			TruncationToken: "DistributionList.IsTruncated",
		},
	}

	if input == nil {
		input = &ListDistributionsInput{}
	}

	req := c.newRequest(op, input, &ListDistributionsOutput{})

	return ListDistributionsRequest{Request: req, Input: input, Copy: c.ListDistributionsRequest}
}

// ListDistributionsRequest is the request type for the
// ListDistributions API operation.
type ListDistributionsRequest struct {
	*aws.Request
	Input *ListDistributionsInput
	Copy  func(*ListDistributionsInput) ListDistributionsRequest
}

// Send marshals and sends the ListDistributions API request.
func (r ListDistributionsRequest) Send(ctx context.Context) (*ListDistributionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDistributionsResponse{
		ListDistributionsOutput: r.Request.Data.(*ListDistributionsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDistributionsRequestPaginator returns a paginator for ListDistributions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDistributionsRequest(input)
//   p := cloudfront.NewListDistributionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDistributionsPaginator(req ListDistributionsRequest) ListDistributionsPaginator {
	return ListDistributionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDistributionsInput
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

// ListDistributionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDistributionsPaginator struct {
	aws.Pager
}

func (p *ListDistributionsPaginator) CurrentPage() *ListDistributionsOutput {
	return p.Pager.CurrentPage().(*ListDistributionsOutput)
}

// ListDistributionsResponse is the response type for the
// ListDistributions API operation.
type ListDistributionsResponse struct {
	*ListDistributionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDistributions request.
func (r *ListDistributionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
