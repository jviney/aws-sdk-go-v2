// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// The input for the DescribeDeliveryChannels action.
type DescribeDeliveryChannelsInput struct {
	_ struct{} `type:"structure"`

	// A list of delivery channel names.
	DeliveryChannelNames []string `type:"list"`
}

// String returns the string representation
func (s DescribeDeliveryChannelsInput) String() string {
	return awsutil.Prettify(s)
}

// The output for the DescribeDeliveryChannels action.
type DescribeDeliveryChannelsOutput struct {
	_ struct{} `type:"structure"`

	// A list that contains the descriptions of the specified delivery channel.
	DeliveryChannels []DeliveryChannel `type:"list"`
}

// String returns the string representation
func (s DescribeDeliveryChannelsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDeliveryChannels = "DescribeDeliveryChannels"

// DescribeDeliveryChannelsRequest returns a request value for making API operation for
// AWS Config.
//
// Returns details about the specified delivery channel. If a delivery channel
// is not specified, this action returns the details of all delivery channels
// associated with the account.
//
// Currently, you can specify only one delivery channel per region in your account.
//
//    // Example sending a request using DescribeDeliveryChannelsRequest.
//    req := client.DescribeDeliveryChannelsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DescribeDeliveryChannels
func (c *Client) DescribeDeliveryChannelsRequest(input *DescribeDeliveryChannelsInput) DescribeDeliveryChannelsRequest {
	op := &aws.Operation{
		Name:       opDescribeDeliveryChannels,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDeliveryChannelsInput{}
	}

	req := c.newRequest(op, input, &DescribeDeliveryChannelsOutput{})

	return DescribeDeliveryChannelsRequest{Request: req, Input: input, Copy: c.DescribeDeliveryChannelsRequest}
}

// DescribeDeliveryChannelsRequest is the request type for the
// DescribeDeliveryChannels API operation.
type DescribeDeliveryChannelsRequest struct {
	*aws.Request
	Input *DescribeDeliveryChannelsInput
	Copy  func(*DescribeDeliveryChannelsInput) DescribeDeliveryChannelsRequest
}

// Send marshals and sends the DescribeDeliveryChannels API request.
func (r DescribeDeliveryChannelsRequest) Send(ctx context.Context) (*DescribeDeliveryChannelsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDeliveryChannelsResponse{
		DescribeDeliveryChannelsOutput: r.Request.Data.(*DescribeDeliveryChannelsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDeliveryChannelsResponse is the response type for the
// DescribeDeliveryChannels API operation.
type DescribeDeliveryChannelsResponse struct {
	*DescribeDeliveryChannelsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDeliveryChannels request.
func (r *DescribeDeliveryChannelsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
