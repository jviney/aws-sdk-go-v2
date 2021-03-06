// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetSendStatisticsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetSendStatisticsInput) String() string {
	return awsutil.Prettify(s)
}

// Represents a list of data points. This list contains aggregated data from
// the previous two weeks of your sending activity with Amazon SES.
type GetSendStatisticsOutput struct {
	_ struct{} `type:"structure"`

	// A list of data points, each of which represents 15 minutes of activity.
	SendDataPoints []SendDataPoint `type:"list"`
}

// String returns the string representation
func (s GetSendStatisticsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetSendStatistics = "GetSendStatistics"

// GetSendStatisticsRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Provides sending statistics for the current AWS Region. The result is a list
// of data points, representing the last two weeks of sending activity. Each
// data point in the list contains statistics for a 15-minute period of time.
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using GetSendStatisticsRequest.
//    req := client.GetSendStatisticsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/GetSendStatistics
func (c *Client) GetSendStatisticsRequest(input *GetSendStatisticsInput) GetSendStatisticsRequest {
	op := &aws.Operation{
		Name:       opGetSendStatistics,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSendStatisticsInput{}
	}

	req := c.newRequest(op, input, &GetSendStatisticsOutput{})

	return GetSendStatisticsRequest{Request: req, Input: input, Copy: c.GetSendStatisticsRequest}
}

// GetSendStatisticsRequest is the request type for the
// GetSendStatistics API operation.
type GetSendStatisticsRequest struct {
	*aws.Request
	Input *GetSendStatisticsInput
	Copy  func(*GetSendStatisticsInput) GetSendStatisticsRequest
}

// Send marshals and sends the GetSendStatistics API request.
func (r GetSendStatisticsRequest) Send(ctx context.Context) (*GetSendStatisticsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSendStatisticsResponse{
		GetSendStatisticsOutput: r.Request.Data.(*GetSendStatisticsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSendStatisticsResponse is the response type for the
// GetSendStatistics API operation.
type GetSendStatisticsResponse struct {
	*GetSendStatisticsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSendStatistics request.
func (r *GetSendStatisticsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
