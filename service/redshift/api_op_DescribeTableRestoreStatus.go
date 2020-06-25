// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeTableRestoreStatusInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Redshift cluster that the table is being restored to.
	ClusterIdentifier *string `type:"string"`

	// An optional pagination token provided by a previous DescribeTableRestoreStatus
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by the MaxRecords parameter.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	MaxRecords *int64 `type:"integer"`

	// The identifier of the table restore request to return status for. If you
	// don't specify a TableRestoreRequestId value, then DescribeTableRestoreStatus
	// returns the status of all in-progress table restore requests.
	TableRestoreRequestId *string `type:"string"`
}

// String returns the string representation
func (s DescribeTableRestoreStatusInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeTableRestoreStatusOutput struct {
	_ struct{} `type:"structure"`

	// A pagination token that can be used in a subsequent DescribeTableRestoreStatus
	// request.
	Marker *string `type:"string"`

	// A list of status details for one or more table restore requests.
	TableRestoreStatusDetails []TableRestoreStatus `locationNameList:"TableRestoreStatus" type:"list"`
}

// String returns the string representation
func (s DescribeTableRestoreStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTableRestoreStatus = "DescribeTableRestoreStatus"

// DescribeTableRestoreStatusRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Lists the status of one or more table restore requests made using the RestoreTableFromClusterSnapshot
// API action. If you don't specify a value for the TableRestoreRequestId parameter,
// then DescribeTableRestoreStatus returns the status of all table restore requests
// ordered by the date and time of the request in ascending order. Otherwise
// DescribeTableRestoreStatus returns the status of the table specified by TableRestoreRequestId.
//
//    // Example sending a request using DescribeTableRestoreStatusRequest.
//    req := client.DescribeTableRestoreStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeTableRestoreStatus
func (c *Client) DescribeTableRestoreStatusRequest(input *DescribeTableRestoreStatusInput) DescribeTableRestoreStatusRequest {
	op := &aws.Operation{
		Name:       opDescribeTableRestoreStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTableRestoreStatusInput{}
	}

	req := c.newRequest(op, input, &DescribeTableRestoreStatusOutput{})

	return DescribeTableRestoreStatusRequest{Request: req, Input: input, Copy: c.DescribeTableRestoreStatusRequest}
}

// DescribeTableRestoreStatusRequest is the request type for the
// DescribeTableRestoreStatus API operation.
type DescribeTableRestoreStatusRequest struct {
	*aws.Request
	Input *DescribeTableRestoreStatusInput
	Copy  func(*DescribeTableRestoreStatusInput) DescribeTableRestoreStatusRequest
}

// Send marshals and sends the DescribeTableRestoreStatus API request.
func (r DescribeTableRestoreStatusRequest) Send(ctx context.Context) (*DescribeTableRestoreStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTableRestoreStatusResponse{
		DescribeTableRestoreStatusOutput: r.Request.Data.(*DescribeTableRestoreStatusOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTableRestoreStatusResponse is the response type for the
// DescribeTableRestoreStatus API operation.
type DescribeTableRestoreStatusResponse struct {
	*DescribeTableRestoreStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTableRestoreStatus request.
func (r *DescribeTableRestoreStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
