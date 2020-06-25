// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplacecommerceanalytics

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Container for the parameters to the GenerateDataSet operation.
type GenerateDataSetInput struct {
	_ struct{} `type:"structure"`

	// (Optional) Key-value pairs which will be returned, unmodified, in the Amazon
	// SNS notification message and the data set metadata file. These key-value
	// pairs can be used to correlated responses with tracking information from
	// other systems.
	CustomerDefinedValues map[string]string `locationName:"customerDefinedValues" min:"1" type:"map"`

	// The date a data set was published. For daily data sets, provide a date with
	// day-level granularity for the desired day. For monthly data sets except those
	// with prefix disbursed_amount, provide a date with month-level granularity
	// for the desired month (the day value will be ignored). For data sets with
	// prefix disbursed_amount, provide a date with day-level granularity for the
	// desired day. For these data sets we will look backwards in time over the
	// range of 31 days until the first data set is found (the latest one).
	//
	// DataSetPublicationDate is a required field
	DataSetPublicationDate *time.Time `locationName:"dataSetPublicationDate" type:"timestamp" required:"true"`

	// The desired data set type.
	//
	//    * customer_subscriber_hourly_monthly_subscriptions From 2017-09-15 to
	//    present: Available daily by 24:00 UTC.
	//
	//    * customer_subscriber_annual_subscriptions From 2017-09-15 to present:
	//    Available daily by 24:00 UTC.
	//
	//    * daily_business_usage_by_instance_type From 2017-09-15 to present: Available
	//    daily by 24:00 UTC.
	//
	//    * daily_business_fees From 2017-09-15 to present: Available daily by 24:00
	//    UTC.
	//
	//    * daily_business_free_trial_conversions From 2017-09-15 to present: Available
	//    daily by 24:00 UTC.
	//
	//    * daily_business_new_instances From 2017-09-15 to present: Available daily
	//    by 24:00 UTC.
	//
	//    * daily_business_new_product_subscribers From 2017-09-15 to present: Available
	//    daily by 24:00 UTC.
	//
	//    * daily_business_canceled_product_subscribers From 2017-09-15 to present:
	//    Available daily by 24:00 UTC.
	//
	//    * monthly_revenue_billing_and_revenue_data From 2017-09-15 to present:
	//    Available monthly on the 15th day of the month by 24:00 UTC. Data includes
	//    metered transactions (e.g. hourly) from one month prior.
	//
	//    * monthly_revenue_annual_subscriptions From 2017-09-15 to present: Available
	//    monthly on the 15th day of the month by 24:00 UTC. Data includes up-front
	//    software charges (e.g. annual) from one month prior.
	//
	//    * monthly_revenue_field_demonstration_usage From 2018-03-15 to present:
	//    Available monthly on the 15th day of the month by 24:00 UTC.
	//
	//    * monthly_revenue_flexible_payment_schedule From 2018-11-15 to present:
	//    Available monthly on the 15th day of the month by 24:00 UTC.
	//
	//    * disbursed_amount_by_product From 2017-09-15 to present: Available every
	//    30 days by 24:00 UTC.
	//
	//    * disbursed_amount_by_instance_hours From 2017-09-15 to present: Available
	//    every 30 days by 24:00 UTC.
	//
	//    * disbursed_amount_by_customer_geo From 2017-09-15 to present: Available
	//    every 30 days by 24:00 UTC.
	//
	//    * disbursed_amount_by_age_of_uncollected_funds From 2017-09-15 to present:
	//    Available every 30 days by 24:00 UTC.
	//
	//    * disbursed_amount_by_age_of_disbursed_funds From 2017-09-15 to present:
	//    Available every 30 days by 24:00 UTC.
	//
	//    * disbursed_amount_by_age_of_past_due_funds From 2018-04-07 to present:
	//    Available every 30 days by 24:00 UTC.
	//
	//    * disbursed_amount_by_uncollected_funds_breakdown From 2019-10-04 to present:
	//    Available every 30 days by 24:00 UTC.
	//
	//    * sales_compensation_billed_revenue From 2017-09-15 to present: Available
	//    monthly on the 15th day of the month by 24:00 UTC. Data includes metered
	//    transactions (e.g. hourly) from one month prior, and up-front software
	//    charges (e.g. annual) from one month prior.
	//
	//    * us_sales_and_use_tax_records From 2017-09-15 to present: Available monthly
	//    on the 15th day of the month by 24:00 UTC.
	//
	// DataSetType is a required field
	DataSetType DataSetType `locationName:"dataSetType" min:"1" type:"string" required:"true" enum:"true"`

	// The name (friendly name, not ARN) of the destination S3 bucket.
	//
	// DestinationS3BucketName is a required field
	DestinationS3BucketName *string `locationName:"destinationS3BucketName" min:"1" type:"string" required:"true"`

	// (Optional) The desired S3 prefix for the published data set, similar to a
	// directory path in standard file systems. For example, if given the bucket
	// name "mybucket" and the prefix "myprefix/mydatasets", the output file "outputfile"
	// would be published to "s3://mybucket/myprefix/mydatasets/outputfile". If
	// the prefix directory structure does not exist, it will be created. If no
	// prefix is provided, the data set will be published to the S3 bucket root.
	DestinationS3Prefix *string `locationName:"destinationS3Prefix" type:"string"`

	// The Amazon Resource Name (ARN) of the Role with an attached permissions policy
	// to interact with the provided AWS services.
	//
	// RoleNameArn is a required field
	RoleNameArn *string `locationName:"roleNameArn" min:"1" type:"string" required:"true"`

	// Amazon Resource Name (ARN) for the SNS Topic that will be notified when the
	// data set has been published or if an error has occurred.
	//
	// SnsTopicArn is a required field
	SnsTopicArn *string `locationName:"snsTopicArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GenerateDataSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GenerateDataSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GenerateDataSetInput"}
	if s.CustomerDefinedValues != nil && len(s.CustomerDefinedValues) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CustomerDefinedValues", 1))
	}

	if s.DataSetPublicationDate == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSetPublicationDate"))
	}
	if len(s.DataSetType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("DataSetType"))
	}

	if s.DestinationS3BucketName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationS3BucketName"))
	}
	if s.DestinationS3BucketName != nil && len(*s.DestinationS3BucketName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DestinationS3BucketName", 1))
	}

	if s.RoleNameArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleNameArn"))
	}
	if s.RoleNameArn != nil && len(*s.RoleNameArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleNameArn", 1))
	}

	if s.SnsTopicArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnsTopicArn"))
	}
	if s.SnsTopicArn != nil && len(*s.SnsTopicArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SnsTopicArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Container for the result of the GenerateDataSet operation.
type GenerateDataSetOutput struct {
	_ struct{} `type:"structure"`

	// A unique identifier representing a specific request to the GenerateDataSet
	// operation. This identifier can be used to correlate a request with notifications
	// from the SNS topic.
	DataSetRequestId *string `locationName:"dataSetRequestId" type:"string"`
}

// String returns the string representation
func (s GenerateDataSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opGenerateDataSet = "GenerateDataSet"

// GenerateDataSetRequest returns a request value for making API operation for
// AWS Marketplace Commerce Analytics.
//
// Given a data set type and data set publication date, asynchronously publishes
// the requested data set to the specified S3 bucket and notifies the specified
// SNS topic once the data is available. Returns a unique request identifier
// that can be used to correlate requests with notifications from the SNS topic.
// Data sets will be published in comma-separated values (CSV) format with the
// file name {data_set_type}_YYYY-MM-DD.csv. If a file with the same name already
// exists (e.g. if the same data set is requested twice), the original file
// will be overwritten by the new file. Requires a Role with an attached permissions
// policy providing Allow permissions for the following actions: s3:PutObject,
// s3:GetBucketLocation, sns:GetTopicAttributes, sns:Publish, iam:GetRolePolicy.
//
//    // Example sending a request using GenerateDataSetRequest.
//    req := client.GenerateDataSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/marketplacecommerceanalytics-2015-07-01/GenerateDataSet
func (c *Client) GenerateDataSetRequest(input *GenerateDataSetInput) GenerateDataSetRequest {
	op := &aws.Operation{
		Name:       opGenerateDataSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GenerateDataSetInput{}
	}

	req := c.newRequest(op, input, &GenerateDataSetOutput{})

	return GenerateDataSetRequest{Request: req, Input: input, Copy: c.GenerateDataSetRequest}
}

// GenerateDataSetRequest is the request type for the
// GenerateDataSet API operation.
type GenerateDataSetRequest struct {
	*aws.Request
	Input *GenerateDataSetInput
	Copy  func(*GenerateDataSetInput) GenerateDataSetRequest
}

// Send marshals and sends the GenerateDataSet API request.
func (r GenerateDataSetRequest) Send(ctx context.Context) (*GenerateDataSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GenerateDataSetResponse{
		GenerateDataSetOutput: r.Request.Data.(*GenerateDataSetOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GenerateDataSetResponse is the response type for the
// GenerateDataSet API operation.
type GenerateDataSetResponse struct {
	*GenerateDataSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GenerateDataSet request.
func (r *GenerateDataSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
