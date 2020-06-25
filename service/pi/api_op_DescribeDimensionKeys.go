// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pi

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeDimensionKeysInput struct {
	_ struct{} `type:"structure"`

	// The date and time specifying the end of the requested time series data. The
	// value specified is exclusive - data points less than (but not equal to) EndTime
	// will be returned.
	//
	// The value for EndTime must be later than the value for StartTime.
	//
	// EndTime is a required field
	EndTime *time.Time `type:"timestamp" required:"true"`

	// One or more filters to apply in the request. Restrictions:
	//
	//    * Any number of filters by the same dimension, as specified in the GroupBy
	//    or Partition parameters.
	//
	//    * A single filter for any other dimension in this dimension group.
	Filter map[string]string `type:"map"`

	// A specification for how to aggregate the data points from a query result.
	// You must specify a valid dimension group. Performance Insights will return
	// all of the dimensions within that group, unless you provide the names of
	// specific dimensions within that group. You can also request that Performance
	// Insights return a limited number of values for a dimension.
	//
	// GroupBy is a required field
	GroupBy *DimensionGroup `type:"structure" required:"true"`

	// An immutable, AWS Region-unique identifier for a data source. Performance
	// Insights gathers metrics from this data source.
	//
	// To use an Amazon RDS instance as a data source, you specify its DbiResourceId
	// value - for example: db-FAIHNTYBKTGAUSUZQYPDS2GW4A
	//
	// Identifier is a required field
	Identifier *string `type:"string" required:"true"`

	// The maximum number of items to return in the response. If more items exist
	// than the specified MaxRecords value, a pagination token is included in the
	// response so that the remaining results can be retrieved.
	MaxResults *int64 `type:"integer"`

	// The name of a Performance Insights metric to be measured.
	//
	// Valid values for Metric are:
	//
	//    * db.load.avg - a scaled representation of the number of active sessions
	//    for the database engine.
	//
	//    * db.sampledload.avg - the raw number of active sessions for the database
	//    engine.
	//
	// Metric is a required field
	Metric *string `type:"string" required:"true"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the token, up to
	// the value specified by MaxRecords.
	NextToken *string `type:"string"`

	// For each dimension specified in GroupBy, specify a secondary dimension to
	// further subdivide the partition keys in the response.
	PartitionBy *DimensionGroup `type:"structure"`

	// The granularity, in seconds, of the data points returned from Performance
	// Insights. A period can be as short as one second, or as long as one day (86400
	// seconds). Valid values are:
	//
	//    * 1 (one second)
	//
	//    * 60 (one minute)
	//
	//    * 300 (five minutes)
	//
	//    * 3600 (one hour)
	//
	//    * 86400 (twenty-four hours)
	//
	// If you don't specify PeriodInSeconds, then Performance Insights will choose
	// a value for you, with a goal of returning roughly 100-200 data points in
	// the response.
	PeriodInSeconds *int64 `type:"integer"`

	// The AWS service for which Performance Insights will return metrics. The only
	// valid value for ServiceType is: RDS
	//
	// ServiceType is a required field
	ServiceType ServiceType `type:"string" required:"true" enum:"true"`

	// The date and time specifying the beginning of the requested time series data.
	// You can't specify a StartTime that's earlier than 7 days ago. The value specified
	// is inclusive - data points equal to or greater than StartTime will be returned.
	//
	// The value for StartTime must be earlier than the value for EndTime.
	//
	// StartTime is a required field
	StartTime *time.Time `type:"timestamp" required:"true"`
}

// String returns the string representation
func (s DescribeDimensionKeysInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDimensionKeysInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDimensionKeysInput"}

	if s.EndTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndTime"))
	}

	if s.GroupBy == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupBy"))
	}

	if s.Identifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("Identifier"))
	}

	if s.Metric == nil {
		invalidParams.Add(aws.NewErrParamRequired("Metric"))
	}
	if len(s.ServiceType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ServiceType"))
	}

	if s.StartTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartTime"))
	}
	if s.GroupBy != nil {
		if err := s.GroupBy.Validate(); err != nil {
			invalidParams.AddNested("GroupBy", err.(aws.ErrInvalidParams))
		}
	}
	if s.PartitionBy != nil {
		if err := s.PartitionBy.Validate(); err != nil {
			invalidParams.AddNested("PartitionBy", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeDimensionKeysOutput struct {
	_ struct{} `type:"structure"`

	// The end time for the returned dimension keys, after alignment to a granular
	// boundary (as specified by PeriodInSeconds). AlignedEndTime will be greater
	// than or equal to the value of the user-specified Endtime.
	AlignedEndTime *time.Time `type:"timestamp"`

	// The start time for the returned dimension keys, after alignment to a granular
	// boundary (as specified by PeriodInSeconds). AlignedStartTime will be less
	// than or equal to the value of the user-specified StartTime.
	AlignedStartTime *time.Time `type:"timestamp"`

	// The dimension keys that were requested.
	Keys []DimensionKeyDescription `type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the token, up to
	// the value specified by MaxRecords.
	NextToken *string `type:"string"`

	// If PartitionBy was present in the request, PartitionKeys contains the breakdown
	// of dimension keys by the specified partitions.
	PartitionKeys []ResponsePartitionKey `type:"list"`
}

// String returns the string representation
func (s DescribeDimensionKeysOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDimensionKeys = "DescribeDimensionKeys"

// DescribeDimensionKeysRequest returns a request value for making API operation for
// AWS Performance Insights.
//
// For a specific time period, retrieve the top N dimension keys for a metric.
//
//    // Example sending a request using DescribeDimensionKeysRequest.
//    req := client.DescribeDimensionKeysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/DescribeDimensionKeys
func (c *Client) DescribeDimensionKeysRequest(input *DescribeDimensionKeysInput) DescribeDimensionKeysRequest {
	op := &aws.Operation{
		Name:       opDescribeDimensionKeys,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDimensionKeysInput{}
	}

	req := c.newRequest(op, input, &DescribeDimensionKeysOutput{})

	return DescribeDimensionKeysRequest{Request: req, Input: input, Copy: c.DescribeDimensionKeysRequest}
}

// DescribeDimensionKeysRequest is the request type for the
// DescribeDimensionKeys API operation.
type DescribeDimensionKeysRequest struct {
	*aws.Request
	Input *DescribeDimensionKeysInput
	Copy  func(*DescribeDimensionKeysInput) DescribeDimensionKeysRequest
}

// Send marshals and sends the DescribeDimensionKeys API request.
func (r DescribeDimensionKeysRequest) Send(ctx context.Context) (*DescribeDimensionKeysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDimensionKeysResponse{
		DescribeDimensionKeysOutput: r.Request.Data.(*DescribeDimensionKeysOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDimensionKeysResponse is the response type for the
// DescribeDimensionKeys API operation.
type DescribeDimensionKeysResponse struct {
	*DescribeDimensionKeysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDimensionKeys request.
func (r *DescribeDimensionKeysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
