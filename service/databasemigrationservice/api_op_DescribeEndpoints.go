// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeEndpointsInput struct {
	_ struct{} `type:"structure"`

	// Filters applied to the describe action.
	//
	// Valid filter names: endpoint-arn | endpoint-type | endpoint-id | engine-name
	Filters []Filter `type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeEndpointsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEndpointsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEndpointsInput"}
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

type DescribeEndpointsOutput struct {
	_ struct{} `type:"structure"`

	// Endpoint description.
	Endpoints []Endpoint `type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeEndpointsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEndpoints = "DescribeEndpoints"

// DescribeEndpointsRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Returns information about the endpoints for your account in the current region.
//
//    // Example sending a request using DescribeEndpointsRequest.
//    req := client.DescribeEndpointsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribeEndpoints
func (c *Client) DescribeEndpointsRequest(input *DescribeEndpointsInput) DescribeEndpointsRequest {
	op := &aws.Operation{
		Name:       opDescribeEndpoints,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeEndpointsInput{}
	}

	req := c.newRequest(op, input, &DescribeEndpointsOutput{})

	return DescribeEndpointsRequest{Request: req, Input: input, Copy: c.DescribeEndpointsRequest}
}

// DescribeEndpointsRequest is the request type for the
// DescribeEndpoints API operation.
type DescribeEndpointsRequest struct {
	*aws.Request
	Input *DescribeEndpointsInput
	Copy  func(*DescribeEndpointsInput) DescribeEndpointsRequest
}

// Send marshals and sends the DescribeEndpoints API request.
func (r DescribeEndpointsRequest) Send(ctx context.Context) (*DescribeEndpointsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEndpointsResponse{
		DescribeEndpointsOutput: r.Request.Data.(*DescribeEndpointsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeEndpointsRequestPaginator returns a paginator for DescribeEndpoints.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeEndpointsRequest(input)
//   p := databasemigrationservice.NewDescribeEndpointsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeEndpointsPaginator(req DescribeEndpointsRequest) DescribeEndpointsPaginator {
	return DescribeEndpointsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeEndpointsInput
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

// DescribeEndpointsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeEndpointsPaginator struct {
	aws.Pager
}

func (p *DescribeEndpointsPaginator) CurrentPage() *DescribeEndpointsOutput {
	return p.Pager.CurrentPage().(*DescribeEndpointsOutput)
}

// DescribeEndpointsResponse is the response type for the
// DescribeEndpoints API operation.
type DescribeEndpointsResponse struct {
	*DescribeEndpointsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEndpoints request.
func (r *DescribeEndpointsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
