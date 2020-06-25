// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeAutoScalingNotificationTypesInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeAutoScalingNotificationTypesInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeAutoScalingNotificationTypesOutput struct {
	_ struct{} `type:"structure"`

	// The notification types.
	AutoScalingNotificationTypes []string `type:"list"`
}

// String returns the string representation
func (s DescribeAutoScalingNotificationTypesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAutoScalingNotificationTypes = "DescribeAutoScalingNotificationTypes"

// DescribeAutoScalingNotificationTypesRequest returns a request value for making API operation for
// Auto Scaling.
//
// Describes the notification types that are supported by Amazon EC2 Auto Scaling.
//
//    // Example sending a request using DescribeAutoScalingNotificationTypesRequest.
//    req := client.DescribeAutoScalingNotificationTypesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/DescribeAutoScalingNotificationTypes
func (c *Client) DescribeAutoScalingNotificationTypesRequest(input *DescribeAutoScalingNotificationTypesInput) DescribeAutoScalingNotificationTypesRequest {
	op := &aws.Operation{
		Name:       opDescribeAutoScalingNotificationTypes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAutoScalingNotificationTypesInput{}
	}

	req := c.newRequest(op, input, &DescribeAutoScalingNotificationTypesOutput{})

	return DescribeAutoScalingNotificationTypesRequest{Request: req, Input: input, Copy: c.DescribeAutoScalingNotificationTypesRequest}
}

// DescribeAutoScalingNotificationTypesRequest is the request type for the
// DescribeAutoScalingNotificationTypes API operation.
type DescribeAutoScalingNotificationTypesRequest struct {
	*aws.Request
	Input *DescribeAutoScalingNotificationTypesInput
	Copy  func(*DescribeAutoScalingNotificationTypesInput) DescribeAutoScalingNotificationTypesRequest
}

// Send marshals and sends the DescribeAutoScalingNotificationTypes API request.
func (r DescribeAutoScalingNotificationTypesRequest) Send(ctx context.Context) (*DescribeAutoScalingNotificationTypesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAutoScalingNotificationTypesResponse{
		DescribeAutoScalingNotificationTypesOutput: r.Request.Data.(*DescribeAutoScalingNotificationTypesOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAutoScalingNotificationTypesResponse is the response type for the
// DescribeAutoScalingNotificationTypes API operation.
type DescribeAutoScalingNotificationTypesResponse struct {
	*DescribeAutoScalingNotificationTypesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAutoScalingNotificationTypes request.
func (r *DescribeAutoScalingNotificationTypesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
