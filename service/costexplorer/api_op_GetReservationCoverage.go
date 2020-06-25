// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costexplorer

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// You can use the following request parameters to query for how much of your
// instance usage a reservation covered.
type GetReservationCoverageInput struct {
	_ struct{} `type:"structure"`

	// Filters utilization data by dimensions. You can filter by the following dimensions:
	//
	//    * AZ
	//
	//    * CACHE_ENGINE
	//
	//    * DATABASE_ENGINE
	//
	//    * DEPLOYMENT_OPTION
	//
	//    * INSTANCE_TYPE
	//
	//    * LINKED_ACCOUNT
	//
	//    * OPERATING_SYSTEM
	//
	//    * PLATFORM
	//
	//    * REGION
	//
	//    * SERVICE
	//
	//    * TAG
	//
	//    * TENANCY
	//
	// GetReservationCoverage uses the same Expression (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html)
	// object as the other operations, but only AND is supported among each dimension.
	// You can nest only one level deep. If there are multiple values for a dimension,
	// they are OR'd together.
	//
	// If you don't provide a SERVICE filter, Cost Explorer defaults to EC2.
	//
	// Cost category is also supported.
	Filter *Expression `type:"structure"`

	// The granularity of the AWS cost data for the reservation. Valid values are
	// MONTHLY and DAILY.
	//
	// If GroupBy is set, Granularity can't be set. If Granularity isn't set, the
	// response object doesn't include Granularity, either MONTHLY or DAILY.
	//
	// The GetReservationCoverage operation supports only DAILY and MONTHLY granularities.
	Granularity Granularity `type:"string" enum:"true"`

	// You can group the data by the following attributes:
	//
	//    * AZ
	//
	//    * CACHE_ENGINE
	//
	//    * DATABASE_ENGINE
	//
	//    * DEPLOYMENT_OPTION
	//
	//    * INSTANCE_TYPE
	//
	//    * LINKED_ACCOUNT
	//
	//    * OPERATING_SYSTEM
	//
	//    * PLATFORM
	//
	//    * REGION
	//
	//    * TENANCY
	GroupBy []GroupDefinition `type:"list"`

	// The measurement that you want your reservation coverage reported in.
	//
	// Valid values are Hour, Unit, and Cost. You can use multiple values in a request.
	Metrics []string `type:"list"`

	// The token to retrieve the next set of results. AWS provides the token when
	// the response from a previous call has more results than the maximum page
	// size.
	NextPageToken *string `type:"string"`

	// The start and end dates of the period that you want to retrieve data about
	// reservation coverage for. You can retrieve data for a maximum of 13 months:
	// the last 12 months and the current month. The start date is inclusive, but
	// the end date is exclusive. For example, if start is 2017-01-01 and end is
	// 2017-05-01, then the cost and usage data is retrieved from 2017-01-01 up
	// to and including 2017-04-30 but not including 2017-05-01.
	//
	// TimePeriod is a required field
	TimePeriod *DateInterval `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetReservationCoverageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetReservationCoverageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetReservationCoverageInput"}

	if s.TimePeriod == nil {
		invalidParams.Add(aws.NewErrParamRequired("TimePeriod"))
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			invalidParams.AddNested("Filter", err.(aws.ErrInvalidParams))
		}
	}
	if s.TimePeriod != nil {
		if err := s.TimePeriod.Validate(); err != nil {
			invalidParams.AddNested("TimePeriod", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetReservationCoverageOutput struct {
	_ struct{} `type:"structure"`

	// The amount of time that your reservations covered.
	//
	// CoveragesByTime is a required field
	CoveragesByTime []CoverageByTime `type:"list" required:"true"`

	// The token for the next set of retrievable results. AWS provides the token
	// when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string `type:"string"`

	// The total amount of instance usage that a reservation covered.
	Total *Coverage `type:"structure"`
}

// String returns the string representation
func (s GetReservationCoverageOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetReservationCoverage = "GetReservationCoverage"

// GetReservationCoverageRequest returns a request value for making API operation for
// AWS Cost Explorer Service.
//
// Retrieves the reservation coverage for your account. This enables you to
// see how much of your Amazon Elastic Compute Cloud, Amazon ElastiCache, Amazon
// Relational Database Service, or Amazon Redshift usage is covered by a reservation.
// An organization's master account can see the coverage of the associated member
// accounts. This supports dimensions, Cost Categories, and nested expressions.
// For any time period, you can filter data about reservation usage by the following
// dimensions:
//
//    * AZ
//
//    * CACHE_ENGINE
//
//    * DATABASE_ENGINE
//
//    * DEPLOYMENT_OPTION
//
//    * INSTANCE_TYPE
//
//    * LINKED_ACCOUNT
//
//    * OPERATING_SYSTEM
//
//    * PLATFORM
//
//    * REGION
//
//    * SERVICE
//
//    * TAG
//
//    * TENANCY
//
// To determine valid values for a dimension, use the GetDimensionValues operation.
//
//    // Example sending a request using GetReservationCoverageRequest.
//    req := client.GetReservationCoverageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetReservationCoverage
func (c *Client) GetReservationCoverageRequest(input *GetReservationCoverageInput) GetReservationCoverageRequest {
	op := &aws.Operation{
		Name:       opGetReservationCoverage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetReservationCoverageInput{}
	}

	req := c.newRequest(op, input, &GetReservationCoverageOutput{})

	return GetReservationCoverageRequest{Request: req, Input: input, Copy: c.GetReservationCoverageRequest}
}

// GetReservationCoverageRequest is the request type for the
// GetReservationCoverage API operation.
type GetReservationCoverageRequest struct {
	*aws.Request
	Input *GetReservationCoverageInput
	Copy  func(*GetReservationCoverageInput) GetReservationCoverageRequest
}

// Send marshals and sends the GetReservationCoverage API request.
func (r GetReservationCoverageRequest) Send(ctx context.Context) (*GetReservationCoverageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetReservationCoverageResponse{
		GetReservationCoverageOutput: r.Request.Data.(*GetReservationCoverageOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetReservationCoverageResponse is the response type for the
// GetReservationCoverage API operation.
type GetReservationCoverageResponse struct {
	*GetReservationCoverageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetReservationCoverage request.
func (r *GetReservationCoverageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
