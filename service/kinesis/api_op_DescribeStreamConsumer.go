// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeStreamConsumerInput struct {
	_ struct{} `type:"structure"`

	// The ARN returned by Kinesis Data Streams when you registered the consumer.
	ConsumerARN *string `min:"1" type:"string"`

	// The name that you gave to the consumer.
	ConsumerName *string `min:"1" type:"string"`

	// The ARN of the Kinesis data stream that the consumer is registered with.
	// For more information, see Amazon Resource Names (ARNs) and AWS Service Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kinesis-streams).
	StreamARN *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeStreamConsumerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStreamConsumerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStreamConsumerInput"}
	if s.ConsumerARN != nil && len(*s.ConsumerARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConsumerARN", 1))
	}
	if s.ConsumerName != nil && len(*s.ConsumerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConsumerName", 1))
	}
	if s.StreamARN != nil && len(*s.StreamARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamARN", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeStreamConsumerOutput struct {
	_ struct{} `type:"structure"`

	// An object that represents the details of the consumer.
	//
	// ConsumerDescription is a required field
	ConsumerDescription *ConsumerDescription `type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeStreamConsumerOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeStreamConsumer = "DescribeStreamConsumer"

// DescribeStreamConsumerRequest returns a request value for making API operation for
// Amazon Kinesis.
//
// To get the description of a registered consumer, provide the ARN of the consumer.
// Alternatively, you can provide the ARN of the data stream and the name you
// gave the consumer when you registered it. You may also provide all three
// parameters, as long as they don't conflict with each other. If you don't
// know the name or ARN of the consumer that you want to describe, you can use
// the ListStreamConsumers operation to get a list of the descriptions of all
// the consumers that are currently registered with a given data stream.
//
// This operation has a limit of 20 transactions per second per account.
//
//    // Example sending a request using DescribeStreamConsumerRequest.
//    req := client.DescribeStreamConsumerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/DescribeStreamConsumer
func (c *Client) DescribeStreamConsumerRequest(input *DescribeStreamConsumerInput) DescribeStreamConsumerRequest {
	op := &aws.Operation{
		Name:       opDescribeStreamConsumer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeStreamConsumerInput{}
	}

	req := c.newRequest(op, input, &DescribeStreamConsumerOutput{})

	return DescribeStreamConsumerRequest{Request: req, Input: input, Copy: c.DescribeStreamConsumerRequest}
}

// DescribeStreamConsumerRequest is the request type for the
// DescribeStreamConsumer API operation.
type DescribeStreamConsumerRequest struct {
	*aws.Request
	Input *DescribeStreamConsumerInput
	Copy  func(*DescribeStreamConsumerInput) DescribeStreamConsumerRequest
}

// Send marshals and sends the DescribeStreamConsumer API request.
func (r DescribeStreamConsumerRequest) Send(ctx context.Context) (*DescribeStreamConsumerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStreamConsumerResponse{
		DescribeStreamConsumerOutput: r.Request.Data.(*DescribeStreamConsumerOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeStreamConsumerResponse is the response type for the
// DescribeStreamConsumer API operation.
type DescribeStreamConsumerResponse struct {
	*DescribeStreamConsumerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStreamConsumer request.
func (r *DescribeStreamConsumerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
