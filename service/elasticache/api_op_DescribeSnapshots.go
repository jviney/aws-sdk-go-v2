// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DescribeSnapshotsMessage operation.
type DescribeSnapshotsInput struct {
	_ struct{} `type:"structure"`

	// A user-supplied cluster identifier. If this parameter is specified, only
	// snapshots associated with that specific cluster are described.
	CacheClusterId *string `type:"string"`

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a marker is included in the response
	// so that the remaining results can be retrieved.
	//
	// Default: 50
	//
	// Constraints: minimum 20; maximum 50.
	MaxRecords *int64 `type:"integer"`

	// A user-supplied replication group identifier. If this parameter is specified,
	// only snapshots associated with that specific replication group are described.
	ReplicationGroupId *string `type:"string"`

	// A Boolean value which if true, the node group (shard) configuration is included
	// in the snapshot description.
	ShowNodeGroupConfig *bool `type:"boolean"`

	// A user-supplied name of the snapshot. If this parameter is specified, only
	// this snapshot are described.
	SnapshotName *string `type:"string"`

	// If set to system, the output shows snapshots that were automatically created
	// by ElastiCache. If set to user the output shows snapshots that were manually
	// created. If omitted, the output shows both automatically and manually created
	// snapshots.
	SnapshotSource *string `type:"string"`
}

// String returns the string representation
func (s DescribeSnapshotsInput) String() string {
	return awsutil.Prettify(s)
}

// Represents the output of a DescribeSnapshots operation.
type DescribeSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// A list of snapshots. Each item in the list contains detailed information
	// about one snapshot.
	Snapshots []Snapshot `locationNameList:"Snapshot" type:"list"`
}

// String returns the string representation
func (s DescribeSnapshotsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSnapshots = "DescribeSnapshots"

// DescribeSnapshotsRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Returns information about cluster or replication group snapshots. By default,
// DescribeSnapshots lists all of your snapshots; it can optionally describe
// a single snapshot, or just the snapshots associated with a particular cache
// cluster.
//
// This operation is valid for Redis only.
//
//    // Example sending a request using DescribeSnapshotsRequest.
//    req := client.DescribeSnapshotsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DescribeSnapshots
func (c *Client) DescribeSnapshotsRequest(input *DescribeSnapshotsInput) DescribeSnapshotsRequest {
	op := &aws.Operation{
		Name:       opDescribeSnapshots,
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
		input = &DescribeSnapshotsInput{}
	}

	req := c.newRequest(op, input, &DescribeSnapshotsOutput{})

	return DescribeSnapshotsRequest{Request: req, Input: input, Copy: c.DescribeSnapshotsRequest}
}

// DescribeSnapshotsRequest is the request type for the
// DescribeSnapshots API operation.
type DescribeSnapshotsRequest struct {
	*aws.Request
	Input *DescribeSnapshotsInput
	Copy  func(*DescribeSnapshotsInput) DescribeSnapshotsRequest
}

// Send marshals and sends the DescribeSnapshots API request.
func (r DescribeSnapshotsRequest) Send(ctx context.Context) (*DescribeSnapshotsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSnapshotsResponse{
		DescribeSnapshotsOutput: r.Request.Data.(*DescribeSnapshotsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeSnapshotsRequestPaginator returns a paginator for DescribeSnapshots.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeSnapshotsRequest(input)
//   p := elasticache.NewDescribeSnapshotsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeSnapshotsPaginator(req DescribeSnapshotsRequest) DescribeSnapshotsPaginator {
	return DescribeSnapshotsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeSnapshotsInput
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

// DescribeSnapshotsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeSnapshotsPaginator struct {
	aws.Pager
}

func (p *DescribeSnapshotsPaginator) CurrentPage() *DescribeSnapshotsOutput {
	return p.Pager.CurrentPage().(*DescribeSnapshotsOutput)
}

// DescribeSnapshotsResponse is the response type for the
// DescribeSnapshots API operation.
type DescribeSnapshotsResponse struct {
	*DescribeSnapshotsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSnapshots request.
func (r *DescribeSnapshotsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
