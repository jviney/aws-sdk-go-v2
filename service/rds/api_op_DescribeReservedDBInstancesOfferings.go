// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeReservedDBInstancesOfferingsInput struct {
	_ struct{} `type:"structure"`

	// The DB instance class filter value. Specify this parameter to show only the
	// available offerings matching the specified DB instance class.
	DBInstanceClass *string `type:"string"`

	// Duration filter value, specified in years or seconds. Specify this parameter
	// to show only reservations for this duration.
	//
	// Valid Values: 1 | 3 | 31536000 | 94608000
	Duration *string `type:"string"`

	// This parameter isn't currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more than the
	// MaxRecords value is available, a pagination token called a marker is included
	// in the response so you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// A value that indicates whether to show only those reservations that support
	// Multi-AZ.
	MultiAZ *bool `type:"boolean"`

	// The offering type filter value. Specify this parameter to show only the available
	// offerings matching the specified offering type.
	//
	// Valid Values: "Partial Upfront" | "All Upfront" | "No Upfront"
	OfferingType *string `type:"string"`

	// Product description filter value. Specify this parameter to show only the
	// available offerings that contain the specified product description.
	//
	// The results show offerings that partially match the filter value.
	ProductDescription *string `type:"string"`

	// The offering identifier filter value. Specify this parameter to show only
	// the available offering that matches the specified reservation identifier.
	//
	// Example: 438012d3-4052-4cc7-b2e3-8d3372e0e706
	ReservedDBInstancesOfferingId *string `type:"string"`
}

// String returns the string representation
func (s DescribeReservedDBInstancesOfferingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeReservedDBInstancesOfferingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeReservedDBInstancesOfferingsInput"}
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

// Contains the result of a successful invocation of the DescribeReservedDBInstancesOfferings
// action.
type DescribeReservedDBInstancesOfferingsOutput struct {
	_ struct{} `type:"structure"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// A list of reserved DB instance offerings.
	ReservedDBInstancesOfferings []ReservedDBInstancesOffering `locationNameList:"ReservedDBInstancesOffering" type:"list"`
}

// String returns the string representation
func (s DescribeReservedDBInstancesOfferingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeReservedDBInstancesOfferings = "DescribeReservedDBInstancesOfferings"

// DescribeReservedDBInstancesOfferingsRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Lists available reserved DB instance offerings.
//
//    // Example sending a request using DescribeReservedDBInstancesOfferingsRequest.
//    req := client.DescribeReservedDBInstancesOfferingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeReservedDBInstancesOfferings
func (c *Client) DescribeReservedDBInstancesOfferingsRequest(input *DescribeReservedDBInstancesOfferingsInput) DescribeReservedDBInstancesOfferingsRequest {
	op := &aws.Operation{
		Name:       opDescribeReservedDBInstancesOfferings,
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
		input = &DescribeReservedDBInstancesOfferingsInput{}
	}

	req := c.newRequest(op, input, &DescribeReservedDBInstancesOfferingsOutput{})

	return DescribeReservedDBInstancesOfferingsRequest{Request: req, Input: input, Copy: c.DescribeReservedDBInstancesOfferingsRequest}
}

// DescribeReservedDBInstancesOfferingsRequest is the request type for the
// DescribeReservedDBInstancesOfferings API operation.
type DescribeReservedDBInstancesOfferingsRequest struct {
	*aws.Request
	Input *DescribeReservedDBInstancesOfferingsInput
	Copy  func(*DescribeReservedDBInstancesOfferingsInput) DescribeReservedDBInstancesOfferingsRequest
}

// Send marshals and sends the DescribeReservedDBInstancesOfferings API request.
func (r DescribeReservedDBInstancesOfferingsRequest) Send(ctx context.Context) (*DescribeReservedDBInstancesOfferingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeReservedDBInstancesOfferingsResponse{
		DescribeReservedDBInstancesOfferingsOutput: r.Request.Data.(*DescribeReservedDBInstancesOfferingsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeReservedDBInstancesOfferingsRequestPaginator returns a paginator for DescribeReservedDBInstancesOfferings.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeReservedDBInstancesOfferingsRequest(input)
//   p := rds.NewDescribeReservedDBInstancesOfferingsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeReservedDBInstancesOfferingsPaginator(req DescribeReservedDBInstancesOfferingsRequest) DescribeReservedDBInstancesOfferingsPaginator {
	return DescribeReservedDBInstancesOfferingsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeReservedDBInstancesOfferingsInput
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

// DescribeReservedDBInstancesOfferingsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeReservedDBInstancesOfferingsPaginator struct {
	aws.Pager
}

func (p *DescribeReservedDBInstancesOfferingsPaginator) CurrentPage() *DescribeReservedDBInstancesOfferingsOutput {
	return p.Pager.CurrentPage().(*DescribeReservedDBInstancesOfferingsOutput)
}

// DescribeReservedDBInstancesOfferingsResponse is the response type for the
// DescribeReservedDBInstancesOfferings API operation.
type DescribeReservedDBInstancesOfferingsResponse struct {
	*DescribeReservedDBInstancesOfferingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeReservedDBInstancesOfferings request.
func (r *DescribeReservedDBInstancesOfferingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
