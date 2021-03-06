// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DescribeReservedInstancesOfferings.
type DescribeReservedInstancesOfferingsInput struct {
	_ struct{} `type:"structure"`

	// The Availability Zone in which the Reserved Instance can be used.
	AvailabilityZone *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * availability-zone - The Availability Zone where the Reserved Instance
	//    can be used.
	//
	//    * duration - The duration of the Reserved Instance (for example, one year
	//    or three years), in seconds (31536000 | 94608000).
	//
	//    * fixed-price - The purchase price of the Reserved Instance (for example,
	//    9800.0).
	//
	//    * instance-type - The instance type that is covered by the reservation.
	//
	//    * marketplace - Set to true to show only Reserved Instance Marketplace
	//    offerings. When this filter is not used, which is the default behavior,
	//    all offerings from both AWS and the Reserved Instance Marketplace are
	//    listed.
	//
	//    * product-description - The Reserved Instance product platform description.
	//    Instances that include (Amazon VPC) in the product platform description
	//    will only be displayed to EC2-Classic account holders and are for use
	//    with Amazon VPC. (Linux/UNIX | Linux/UNIX (Amazon VPC) | SUSE Linux |
	//    SUSE Linux (Amazon VPC) | Red Hat Enterprise Linux | Red Hat Enterprise
	//    Linux (Amazon VPC) | Windows | Windows (Amazon VPC) | Windows with SQL
	//    Server Standard | Windows with SQL Server Standard (Amazon VPC) | Windows
	//    with SQL Server Web | Windows with SQL Server Web (Amazon VPC) | Windows
	//    with SQL Server Enterprise | Windows with SQL Server Enterprise (Amazon
	//    VPC))
	//
	//    * reserved-instances-offering-id - The Reserved Instances offering ID.
	//
	//    * scope - The scope of the Reserved Instance (Availability Zone or Region).
	//
	//    * usage-price - The usage price of the Reserved Instance, per hour (for
	//    example, 0.84).
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// Include Reserved Instance Marketplace offerings in the response.
	IncludeMarketplace *bool `type:"boolean"`

	// The tenancy of the instances covered by the reservation. A Reserved Instance
	// with a tenancy of dedicated is applied to instances that run in a VPC on
	// single-tenant hardware (i.e., Dedicated Instances).
	//
	// Important: The host value cannot be used with this parameter. Use the default
	// or dedicated values only.
	//
	// Default: default
	InstanceTenancy Tenancy `locationName:"instanceTenancy" type:"string" enum:"true"`

	// The instance type that the reservation will cover (for example, m1.small).
	// For more information, see Instance Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	InstanceType InstanceType `type:"string" enum:"true"`

	// The maximum duration (in seconds) to filter when searching for offerings.
	//
	// Default: 94608000 (3 years)
	MaxDuration *int64 `type:"long"`

	// The maximum number of instances to filter when searching for offerings.
	//
	// Default: 20
	MaxInstanceCount *int64 `type:"integer"`

	// The maximum number of results to return for the request in a single page.
	// The remaining results of the initial request can be seen by sending another
	// request with the returned NextToken value. The maximum is 100.
	//
	// Default: 100
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The minimum duration (in seconds) to filter when searching for offerings.
	//
	// Default: 2592000 (1 month)
	MinDuration *int64 `type:"long"`

	// The token to retrieve the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The offering class of the Reserved Instance. Can be standard or convertible.
	OfferingClass OfferingClassType `type:"string" enum:"true"`

	// The Reserved Instance offering type. If you are using tools that predate
	// the 2011-11-01 API version, you only have access to the Medium Utilization
	// Reserved Instance offering type.
	OfferingType OfferingTypeValues `locationName:"offeringType" type:"string" enum:"true"`

	// The Reserved Instance product platform description. Instances that include
	// (Amazon VPC) in the description are for use with Amazon VPC.
	ProductDescription RIProductDescription `type:"string" enum:"true"`

	// One or more Reserved Instances offering IDs.
	ReservedInstancesOfferingIds []string `locationName:"ReservedInstancesOfferingId" type:"list"`
}

// String returns the string representation
func (s DescribeReservedInstancesOfferingsInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the output of DescribeReservedInstancesOfferings.
type DescribeReservedInstancesOfferingsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// A list of Reserved Instances offerings.
	ReservedInstancesOfferings []ReservedInstancesOffering `locationName:"reservedInstancesOfferingsSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeReservedInstancesOfferingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeReservedInstancesOfferings = "DescribeReservedInstancesOfferings"

// DescribeReservedInstancesOfferingsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes Reserved Instance offerings that are available for purchase. With
// Reserved Instances, you purchase the right to launch instances for a period
// of time. During that time period, you do not receive insufficient capacity
// errors, and you pay a lower usage rate than the rate charged for On-Demand
// instances for the actual time used.
//
// If you have listed your own Reserved Instances for sale in the Reserved Instance
// Marketplace, they will be excluded from these results. This is to ensure
// that you do not purchase your own Reserved Instances.
//
// For more information, see Reserved Instance Marketplace (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-market-general.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DescribeReservedInstancesOfferingsRequest.
//    req := client.DescribeReservedInstancesOfferingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeReservedInstancesOfferings
func (c *Client) DescribeReservedInstancesOfferingsRequest(input *DescribeReservedInstancesOfferingsInput) DescribeReservedInstancesOfferingsRequest {
	op := &aws.Operation{
		Name:       opDescribeReservedInstancesOfferings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeReservedInstancesOfferingsInput{}
	}

	req := c.newRequest(op, input, &DescribeReservedInstancesOfferingsOutput{})

	return DescribeReservedInstancesOfferingsRequest{Request: req, Input: input, Copy: c.DescribeReservedInstancesOfferingsRequest}
}

// DescribeReservedInstancesOfferingsRequest is the request type for the
// DescribeReservedInstancesOfferings API operation.
type DescribeReservedInstancesOfferingsRequest struct {
	*aws.Request
	Input *DescribeReservedInstancesOfferingsInput
	Copy  func(*DescribeReservedInstancesOfferingsInput) DescribeReservedInstancesOfferingsRequest
}

// Send marshals and sends the DescribeReservedInstancesOfferings API request.
func (r DescribeReservedInstancesOfferingsRequest) Send(ctx context.Context) (*DescribeReservedInstancesOfferingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeReservedInstancesOfferingsResponse{
		DescribeReservedInstancesOfferingsOutput: r.Request.Data.(*DescribeReservedInstancesOfferingsOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeReservedInstancesOfferingsRequestPaginator returns a paginator for DescribeReservedInstancesOfferings.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeReservedInstancesOfferingsRequest(input)
//   p := ec2.NewDescribeReservedInstancesOfferingsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeReservedInstancesOfferingsPaginator(req DescribeReservedInstancesOfferingsRequest) DescribeReservedInstancesOfferingsPaginator {
	return DescribeReservedInstancesOfferingsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeReservedInstancesOfferingsInput
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

// DescribeReservedInstancesOfferingsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeReservedInstancesOfferingsPaginator struct {
	aws.Pager
}

func (p *DescribeReservedInstancesOfferingsPaginator) CurrentPage() *DescribeReservedInstancesOfferingsOutput {
	return p.Pager.CurrentPage().(*DescribeReservedInstancesOfferingsOutput)
}

// DescribeReservedInstancesOfferingsResponse is the response type for the
// DescribeReservedInstancesOfferings API operation.
type DescribeReservedInstancesOfferingsResponse struct {
	*DescribeReservedInstancesOfferingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeReservedInstancesOfferings request.
func (r *DescribeReservedInstancesOfferingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
