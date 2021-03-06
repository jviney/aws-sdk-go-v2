// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeClustersInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of a cluster whose properties you are requesting. This
	// parameter is case sensitive.
	//
	// The default is that all clusters defined for an account are returned.
	ClusterIdentifier *string `type:"string"`

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeClusters request exceed the
	// value specified in MaxRecords, AWS returns a value in the Marker field of
	// the response. You can retrieve the next set of response records by providing
	// the returned marker value in the Marker parameter and retrying the request.
	//
	// Constraints: You can specify either the ClusterIdentifier parameter or the
	// Marker parameter, but not both.
	Marker *string `type:"string"`

	// The maximum number of response records to return in each call. If the number
	// of remaining response records exceeds the specified MaxRecords value, a value
	// is returned in a marker field of the response. You can retrieve the next
	// set of records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// A tag key or keys for which you want to return all matching clusters that
	// are associated with the specified key or keys. For example, suppose that
	// you have clusters that are tagged with keys called owner and environment.
	// If you specify both of these tag keys in the request, Amazon Redshift returns
	// a response with the clusters that have either or both of these tag keys associated
	// with them.
	TagKeys []string `locationNameList:"TagKey" type:"list"`

	// A tag value or values for which you want to return all matching clusters
	// that are associated with the specified tag value or values. For example,
	// suppose that you have clusters that are tagged with values called admin and
	// test. If you specify both of these tag values in the request, Amazon Redshift
	// returns a response with the clusters that have either or both of these tag
	// values associated with them.
	TagValues []string `locationNameList:"TagValue" type:"list"`
}

// String returns the string representation
func (s DescribeClustersInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the output from the DescribeClusters action.
type DescribeClustersOutput struct {
	_ struct{} `type:"structure"`

	// A list of Cluster objects, where each object describes one cluster.
	Clusters []Cluster `locationNameList:"Cluster" type:"list"`

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeClustersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeClusters = "DescribeClusters"

// DescribeClustersRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Returns properties of provisioned clusters including general cluster properties,
// cluster database properties, maintenance and backup properties, and security
// and access properties. This operation supports pagination. For more information
// about managing clusters, go to Amazon Redshift Clusters (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html)
// in the Amazon Redshift Cluster Management Guide.
//
// If you specify both tag keys and tag values in the same request, Amazon Redshift
// returns all clusters that match any combination of the specified keys and
// values. For example, if you have owner and environment for tag keys, and
// admin and test for tag values, all clusters that have any combination of
// those values are returned.
//
// If both tag keys and values are omitted from the request, clusters are returned
// regardless of whether they have tag keys or values associated with them.
//
//    // Example sending a request using DescribeClustersRequest.
//    req := client.DescribeClustersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeClusters
func (c *Client) DescribeClustersRequest(input *DescribeClustersInput) DescribeClustersRequest {
	op := &aws.Operation{
		Name:       opDescribeClusters,
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
		input = &DescribeClustersInput{}
	}

	req := c.newRequest(op, input, &DescribeClustersOutput{})

	return DescribeClustersRequest{Request: req, Input: input, Copy: c.DescribeClustersRequest}
}

// DescribeClustersRequest is the request type for the
// DescribeClusters API operation.
type DescribeClustersRequest struct {
	*aws.Request
	Input *DescribeClustersInput
	Copy  func(*DescribeClustersInput) DescribeClustersRequest
}

// Send marshals and sends the DescribeClusters API request.
func (r DescribeClustersRequest) Send(ctx context.Context) (*DescribeClustersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeClustersResponse{
		DescribeClustersOutput: r.Request.Data.(*DescribeClustersOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeClustersRequestPaginator returns a paginator for DescribeClusters.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeClustersRequest(input)
//   p := redshift.NewDescribeClustersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeClustersPaginator(req DescribeClustersRequest) DescribeClustersPaginator {
	return DescribeClustersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeClustersInput
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

// DescribeClustersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeClustersPaginator struct {
	aws.Pager
}

func (p *DescribeClustersPaginator) CurrentPage() *DescribeClustersOutput {
	return p.Pager.CurrentPage().(*DescribeClustersOutput)
}

// DescribeClustersResponse is the response type for the
// DescribeClusters API operation.
type DescribeClustersResponse struct {
	*DescribeClustersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeClusters request.
func (r *DescribeClustersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
