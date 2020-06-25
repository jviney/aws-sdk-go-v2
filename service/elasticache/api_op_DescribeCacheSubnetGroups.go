// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DescribeCacheSubnetGroups operation.
type DescribeCacheSubnetGroupsInput struct {
	_ struct{} `type:"structure"`

	// The name of the cache subnet group to return details for.
	CacheSubnetGroupName *string `type:"string"`

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a marker is included in the response
	// so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: minimum 20; maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeCacheSubnetGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Represents the output of a DescribeCacheSubnetGroups operation.
type DescribeCacheSubnetGroupsOutput struct {
	_ struct{} `type:"structure"`

	// A list of cache subnet groups. Each element in the list contains detailed
	// information about one group.
	CacheSubnetGroups []CacheSubnetGroup `locationNameList:"CacheSubnetGroup" type:"list"`

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeCacheSubnetGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCacheSubnetGroups = "DescribeCacheSubnetGroups"

// DescribeCacheSubnetGroupsRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Returns a list of cache subnet group descriptions. If a subnet group name
// is specified, the list contains only the description of that group. This
// is applicable only when you have ElastiCache in VPC setup. All ElastiCache
// clusters now launch in VPC by default.
//
//    // Example sending a request using DescribeCacheSubnetGroupsRequest.
//    req := client.DescribeCacheSubnetGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DescribeCacheSubnetGroups
func (c *Client) DescribeCacheSubnetGroupsRequest(input *DescribeCacheSubnetGroupsInput) DescribeCacheSubnetGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribeCacheSubnetGroups,
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
		input = &DescribeCacheSubnetGroupsInput{}
	}

	req := c.newRequest(op, input, &DescribeCacheSubnetGroupsOutput{})

	return DescribeCacheSubnetGroupsRequest{Request: req, Input: input, Copy: c.DescribeCacheSubnetGroupsRequest}
}

// DescribeCacheSubnetGroupsRequest is the request type for the
// DescribeCacheSubnetGroups API operation.
type DescribeCacheSubnetGroupsRequest struct {
	*aws.Request
	Input *DescribeCacheSubnetGroupsInput
	Copy  func(*DescribeCacheSubnetGroupsInput) DescribeCacheSubnetGroupsRequest
}

// Send marshals and sends the DescribeCacheSubnetGroups API request.
func (r DescribeCacheSubnetGroupsRequest) Send(ctx context.Context) (*DescribeCacheSubnetGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCacheSubnetGroupsResponse{
		DescribeCacheSubnetGroupsOutput: r.Request.Data.(*DescribeCacheSubnetGroupsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeCacheSubnetGroupsRequestPaginator returns a paginator for DescribeCacheSubnetGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeCacheSubnetGroupsRequest(input)
//   p := elasticache.NewDescribeCacheSubnetGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeCacheSubnetGroupsPaginator(req DescribeCacheSubnetGroupsRequest) DescribeCacheSubnetGroupsPaginator {
	return DescribeCacheSubnetGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeCacheSubnetGroupsInput
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

// DescribeCacheSubnetGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeCacheSubnetGroupsPaginator struct {
	aws.Pager
}

func (p *DescribeCacheSubnetGroupsPaginator) CurrentPage() *DescribeCacheSubnetGroupsOutput {
	return p.Pager.CurrentPage().(*DescribeCacheSubnetGroupsOutput)
}

// DescribeCacheSubnetGroupsResponse is the response type for the
// DescribeCacheSubnetGroups API operation.
type DescribeCacheSubnetGroupsResponse struct {
	*DescribeCacheSubnetGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCacheSubnetGroups request.
func (r *DescribeCacheSubnetGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
