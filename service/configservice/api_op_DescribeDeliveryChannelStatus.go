// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// The input for the DeliveryChannelStatus action.
type DescribeDeliveryChannelStatusInput struct {
	_ struct{} `type:"structure"`

	// A list of delivery channel names.
	DeliveryChannelNames []string `type:"list"`
}

// String returns the string representation
func (s DescribeDeliveryChannelStatusInput) String() string {
	return awsutil.Prettify(s)
}

// The output for the DescribeDeliveryChannelStatus action.
type DescribeDeliveryChannelStatusOutput struct {
	_ struct{} `type:"structure"`

	// A list that contains the status of a specified delivery channel.
	DeliveryChannelsStatus []DeliveryChannelStatus `type:"list"`
}

// String returns the string representation
func (s DescribeDeliveryChannelStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDeliveryChannelStatus = "DescribeDeliveryChannelStatus"

// DescribeDeliveryChannelStatusRequest returns a request value for making API operation for
// AWS Config.
//
// Returns the current status of the specified delivery channel. If a delivery
// channel is not specified, this action returns the current status of all delivery
// channels associated with the account.
//
// Currently, you can specify only one delivery channel per region in your account.
//
//    // Example sending a request using DescribeDeliveryChannelStatusRequest.
//    req := client.DescribeDeliveryChannelStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DescribeDeliveryChannelStatus
func (c *Client) DescribeDeliveryChannelStatusRequest(input *DescribeDeliveryChannelStatusInput) DescribeDeliveryChannelStatusRequest {
	op := &aws.Operation{
		Name:       opDescribeDeliveryChannelStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDeliveryChannelStatusInput{}
	}

	req := c.newRequest(op, input, &DescribeDeliveryChannelStatusOutput{})

	return DescribeDeliveryChannelStatusRequest{Request: req, Input: input, Copy: c.DescribeDeliveryChannelStatusRequest}
}

// DescribeDeliveryChannelStatusRequest is the request type for the
// DescribeDeliveryChannelStatus API operation.
type DescribeDeliveryChannelStatusRequest struct {
	*aws.Request
	Input *DescribeDeliveryChannelStatusInput
	Copy  func(*DescribeDeliveryChannelStatusInput) DescribeDeliveryChannelStatusRequest
}

// Send marshals and sends the DescribeDeliveryChannelStatus API request.
func (r DescribeDeliveryChannelStatusRequest) Send(ctx context.Context) (*DescribeDeliveryChannelStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDeliveryChannelStatusResponse{
		DescribeDeliveryChannelStatusOutput: r.Request.Data.(*DescribeDeliveryChannelStatusOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDeliveryChannelStatusResponse is the response type for the
// DescribeDeliveryChannelStatus API operation.
type DescribeDeliveryChannelStatusResponse struct {
	*DescribeDeliveryChannelStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDeliveryChannelStatus request.
func (r *DescribeDeliveryChannelStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
