// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeTableStatisticsInput struct {
	_ struct{} `type:"structure"`

	// Filters applied to the describe table statistics action.
	//
	// Valid filter names: schema-name | table-name | table-state
	//
	// A combination of filters creates an AND condition where each record matches
	// all specified filters.
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
	// Constraints: Minimum 20, maximum 500.
	MaxRecords *int64 `type:"integer"`

	// The Amazon Resource Name (ARN) of the replication task.
	//
	// ReplicationTaskArn is a required field
	ReplicationTaskArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeTableStatisticsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTableStatisticsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTableStatisticsInput"}

	if s.ReplicationTaskArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationTaskArn"))
	}
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

type DescribeTableStatisticsOutput struct {
	_ struct{} `type:"structure"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The Amazon Resource Name (ARN) of the replication task.
	ReplicationTaskArn *string `type:"string"`

	// The table statistics.
	TableStatistics []TableStatistics `type:"list"`
}

// String returns the string representation
func (s DescribeTableStatisticsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTableStatistics = "DescribeTableStatistics"

// DescribeTableStatisticsRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Returns table statistics on the database migration task, including table
// name, rows inserted, rows updated, and rows deleted.
//
// Note that the "last updated" column the DMS console only indicates the time
// that AWS DMS last updated the table statistics record for a table. It does
// not indicate the time of the last update to the table.
//
//    // Example sending a request using DescribeTableStatisticsRequest.
//    req := client.DescribeTableStatisticsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribeTableStatistics
func (c *Client) DescribeTableStatisticsRequest(input *DescribeTableStatisticsInput) DescribeTableStatisticsRequest {
	op := &aws.Operation{
		Name:       opDescribeTableStatistics,
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
		input = &DescribeTableStatisticsInput{}
	}

	req := c.newRequest(op, input, &DescribeTableStatisticsOutput{})

	return DescribeTableStatisticsRequest{Request: req, Input: input, Copy: c.DescribeTableStatisticsRequest}
}

// DescribeTableStatisticsRequest is the request type for the
// DescribeTableStatistics API operation.
type DescribeTableStatisticsRequest struct {
	*aws.Request
	Input *DescribeTableStatisticsInput
	Copy  func(*DescribeTableStatisticsInput) DescribeTableStatisticsRequest
}

// Send marshals and sends the DescribeTableStatistics API request.
func (r DescribeTableStatisticsRequest) Send(ctx context.Context) (*DescribeTableStatisticsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTableStatisticsResponse{
		DescribeTableStatisticsOutput: r.Request.Data.(*DescribeTableStatisticsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeTableStatisticsRequestPaginator returns a paginator for DescribeTableStatistics.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeTableStatisticsRequest(input)
//   p := databasemigrationservice.NewDescribeTableStatisticsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeTableStatisticsPaginator(req DescribeTableStatisticsRequest) DescribeTableStatisticsPaginator {
	return DescribeTableStatisticsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeTableStatisticsInput
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

// DescribeTableStatisticsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeTableStatisticsPaginator struct {
	aws.Pager
}

func (p *DescribeTableStatisticsPaginator) CurrentPage() *DescribeTableStatisticsOutput {
	return p.Pager.CurrentPage().(*DescribeTableStatisticsOutput)
}

// DescribeTableStatisticsResponse is the response type for the
// DescribeTableStatistics API operation.
type DescribeTableStatisticsResponse struct {
	*DescribeTableStatisticsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTableStatistics request.
func (r *DescribeTableStatisticsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
