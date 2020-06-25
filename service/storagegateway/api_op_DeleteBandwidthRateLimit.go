// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing the following fields:
//
//    * DeleteBandwidthRateLimitInput$BandwidthType
type DeleteBandwidthRateLimitInput struct {
	_ struct{} `type:"structure"`

	// One of the BandwidthType values that indicates the gateway bandwidth rate
	// limit to delete.
	//
	// Valid Values: Upload, Download, All.
	//
	// BandwidthType is a required field
	BandwidthType *string `min:"3" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBandwidthRateLimitInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBandwidthRateLimitInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBandwidthRateLimitInput"}

	if s.BandwidthType == nil {
		invalidParams.Add(aws.NewErrParamRequired("BandwidthType"))
	}
	if s.BandwidthType != nil && len(*s.BandwidthType) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("BandwidthType", 3))
	}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway whose
// bandwidth rate information was deleted.
type DeleteBandwidthRateLimitOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s DeleteBandwidthRateLimitOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteBandwidthRateLimit = "DeleteBandwidthRateLimit"

// DeleteBandwidthRateLimitRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Deletes the bandwidth rate limits of a gateway. You can delete either the
// upload and download bandwidth rate limit, or you can delete both. If you
// delete only one of the limits, the other limit remains unchanged. To specify
// which gateway to work with, use the Amazon Resource Name (ARN) of the gateway
// in your request. This operation is supported for the stored volume, cached
// volume and tape gateway types.
//
//    // Example sending a request using DeleteBandwidthRateLimitRequest.
//    req := client.DeleteBandwidthRateLimitRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DeleteBandwidthRateLimit
func (c *Client) DeleteBandwidthRateLimitRequest(input *DeleteBandwidthRateLimitInput) DeleteBandwidthRateLimitRequest {
	op := &aws.Operation{
		Name:       opDeleteBandwidthRateLimit,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteBandwidthRateLimitInput{}
	}

	req := c.newRequest(op, input, &DeleteBandwidthRateLimitOutput{})

	return DeleteBandwidthRateLimitRequest{Request: req, Input: input, Copy: c.DeleteBandwidthRateLimitRequest}
}

// DeleteBandwidthRateLimitRequest is the request type for the
// DeleteBandwidthRateLimit API operation.
type DeleteBandwidthRateLimitRequest struct {
	*aws.Request
	Input *DeleteBandwidthRateLimitInput
	Copy  func(*DeleteBandwidthRateLimitInput) DeleteBandwidthRateLimitRequest
}

// Send marshals and sends the DeleteBandwidthRateLimit API request.
func (r DeleteBandwidthRateLimitRequest) Send(ctx context.Context) (*DeleteBandwidthRateLimitResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBandwidthRateLimitResponse{
		DeleteBandwidthRateLimitOutput: r.Request.Data.(*DeleteBandwidthRateLimitOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBandwidthRateLimitResponse is the response type for the
// DeleteBandwidthRateLimit API operation.
type DeleteBandwidthRateLimitResponse struct {
	*DeleteBandwidthRateLimitOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBandwidthRateLimit request.
func (r *DeleteBandwidthRateLimitResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
