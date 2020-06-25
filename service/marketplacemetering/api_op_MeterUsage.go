// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplacemetering

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type MeterUsageInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the permissions required for the action, but does
	// not make the request. If you have the permissions, the request returns DryRunOperation;
	// otherwise, it returns UnauthorizedException. Defaults to false if not specified.
	DryRun *bool `type:"boolean"`

	// Product code is used to uniquely identify a product in AWS Marketplace. The
	// product code should be the same as the one used during the publishing of
	// a new product.
	//
	// ProductCode is a required field
	ProductCode *string `min:"1" type:"string" required:"true"`

	// Timestamp, in UTC, for which the usage is being reported. Your application
	// can meter usage for up to one hour in the past. Make sure the timestamp value
	// is not before the start of the software usage.
	//
	// Timestamp is a required field
	Timestamp *time.Time `type:"timestamp" required:"true"`

	// It will be one of the fcp dimension name provided during the publishing of
	// the product.
	//
	// UsageDimension is a required field
	UsageDimension *string `min:"1" type:"string" required:"true"`

	// Consumption value for the hour. Defaults to 0 if not specified.
	UsageQuantity *int64 `type:"integer"`
}

// String returns the string representation
func (s MeterUsageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MeterUsageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MeterUsageInput"}

	if s.ProductCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductCode"))
	}
	if s.ProductCode != nil && len(*s.ProductCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProductCode", 1))
	}

	if s.Timestamp == nil {
		invalidParams.Add(aws.NewErrParamRequired("Timestamp"))
	}

	if s.UsageDimension == nil {
		invalidParams.Add(aws.NewErrParamRequired("UsageDimension"))
	}
	if s.UsageDimension != nil && len(*s.UsageDimension) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UsageDimension", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type MeterUsageOutput struct {
	_ struct{} `type:"structure"`

	// Metering record id.
	MeteringRecordId *string `type:"string"`
}

// String returns the string representation
func (s MeterUsageOutput) String() string {
	return awsutil.Prettify(s)
}

const opMeterUsage = "MeterUsage"

// MeterUsageRequest returns a request value for making API operation for
// AWSMarketplace Metering.
//
// API to emit metering records. For identical requests, the API is idempotent.
// It simply returns the metering record ID.
//
// MeterUsage is authenticated on the buyer's AWS account using credentials
// from the EC2 instance, ECS task, or EKS pod.
//
//    // Example sending a request using MeterUsageRequest.
//    req := client.MeterUsageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/MeterUsage
func (c *Client) MeterUsageRequest(input *MeterUsageInput) MeterUsageRequest {
	op := &aws.Operation{
		Name:       opMeterUsage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &MeterUsageInput{}
	}

	req := c.newRequest(op, input, &MeterUsageOutput{})

	return MeterUsageRequest{Request: req, Input: input, Copy: c.MeterUsageRequest}
}

// MeterUsageRequest is the request type for the
// MeterUsage API operation.
type MeterUsageRequest struct {
	*aws.Request
	Input *MeterUsageInput
	Copy  func(*MeterUsageInput) MeterUsageRequest
}

// Send marshals and sends the MeterUsage API request.
func (r MeterUsageRequest) Send(ctx context.Context) (*MeterUsageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &MeterUsageResponse{
		MeterUsageOutput: r.Request.Data.(*MeterUsageOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// MeterUsageResponse is the response type for the
// MeterUsage API operation.
type MeterUsageResponse struct {
	*MeterUsageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// MeterUsage request.
func (r *MeterUsageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
