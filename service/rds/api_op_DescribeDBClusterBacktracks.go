// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeDBClusterBacktracksInput struct {
	_ struct{} `type:"structure"`

	// If specified, this value is the backtrack identifier of the backtrack to
	// be described.
	//
	// Constraints:
	//
	//    * Must contain a valid universally unique identifier (UUID). For more
	//    information about UUIDs, see A Universally Unique Identifier (UUID) URN
	//    Namespace (http://www.ietf.org/rfc/rfc4122.txt).
	//
	// Example: 123e4567-e89b-12d3-a456-426655440000
	BacktrackIdentifier *string `type:"string"`

	// The DB cluster identifier of the DB cluster to be described. This parameter
	// is stored as a lowercase string.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 alphanumeric characters or hyphens.
	//
	//    * First character must be a letter.
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example: my-cluster1
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// A filter that specifies one or more DB clusters to describe. Supported filters
	// include the following:
	//
	//    * db-cluster-backtrack-id - Accepts backtrack identifiers. The results
	//    list includes information about only the backtracks identified by these
	//    identifiers.
	//
	//    * db-cluster-backtrack-status - Accepts any of the following backtrack
	//    status values: applying completed failed pending The results list includes
	//    information about only the backtracks identified by these values.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous DescribeDBClusterBacktracks
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeDBClusterBacktracksInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBClusterBacktracksInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDBClusterBacktracksInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
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

// Contains the result of a successful invocation of the DescribeDBClusterBacktracks
// action.
type DescribeDBClusterBacktracksOutput struct {
	_ struct{} `type:"structure"`

	// Contains a list of backtracks for the user.
	DBClusterBacktracks []DBClusterBacktrack `locationNameList:"DBClusterBacktrack" type:"list"`

	// A pagination token that can be used in a later DescribeDBClusterBacktracks
	// request.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBClusterBacktracksOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDBClusterBacktracks = "DescribeDBClusterBacktracks"

// DescribeDBClusterBacktracksRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Returns information about backtracks for a DB cluster.
//
// For more information on Amazon Aurora, see What Is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using DescribeDBClusterBacktracksRequest.
//    req := client.DescribeDBClusterBacktracksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBClusterBacktracks
func (c *Client) DescribeDBClusterBacktracksRequest(input *DescribeDBClusterBacktracksInput) DescribeDBClusterBacktracksRequest {
	op := &aws.Operation{
		Name:       opDescribeDBClusterBacktracks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDBClusterBacktracksInput{}
	}

	req := c.newRequest(op, input, &DescribeDBClusterBacktracksOutput{})

	return DescribeDBClusterBacktracksRequest{Request: req, Input: input, Copy: c.DescribeDBClusterBacktracksRequest}
}

// DescribeDBClusterBacktracksRequest is the request type for the
// DescribeDBClusterBacktracks API operation.
type DescribeDBClusterBacktracksRequest struct {
	*aws.Request
	Input *DescribeDBClusterBacktracksInput
	Copy  func(*DescribeDBClusterBacktracksInput) DescribeDBClusterBacktracksRequest
}

// Send marshals and sends the DescribeDBClusterBacktracks API request.
func (r DescribeDBClusterBacktracksRequest) Send(ctx context.Context) (*DescribeDBClusterBacktracksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBClusterBacktracksResponse{
		DescribeDBClusterBacktracksOutput: r.Request.Data.(*DescribeDBClusterBacktracksOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDBClusterBacktracksResponse is the response type for the
// DescribeDBClusterBacktracks API operation.
type DescribeDBClusterBacktracksResponse struct {
	*DescribeDBClusterBacktracksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBClusterBacktracks request.
func (r *DescribeDBClusterBacktracksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
