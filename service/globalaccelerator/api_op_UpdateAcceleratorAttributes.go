// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type UpdateAcceleratorAttributesInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the accelerator that you want to update.
	//
	// AcceleratorArn is a required field
	AcceleratorArn *string `type:"string" required:"true"`

	// Update whether flow logs are enabled. The default value is false. If the
	// value is true, FlowLogsS3Bucket and FlowLogsS3Prefix must be specified.
	//
	// For more information, see Flow Logs (https://docs.aws.amazon.com/global-accelerator/latest/dg/monitoring-global-accelerator.flow-logs.html)
	// in the AWS Global Accelerator Developer Guide.
	FlowLogsEnabled *bool `type:"boolean"`

	// The name of the Amazon S3 bucket for the flow logs. Attribute is required
	// if FlowLogsEnabled is true. The bucket must exist and have a bucket policy
	// that grants AWS Global Accelerator permission to write to the bucket.
	FlowLogsS3Bucket *string `type:"string"`

	// Update the prefix for the location in the Amazon S3 bucket for the flow logs.
	// Attribute is required if FlowLogsEnabled is true.
	//
	// If you don’t specify a prefix, the flow logs are stored in the root of
	// the bucket. If you specify slash (/) for the S3 bucket prefix, the log file
	// bucket folder structure will include a double slash (//), like the following:
	//
	// s3-bucket_name//AWSLogs/aws_account_id
	FlowLogsS3Prefix *string `type:"string"`
}

// String returns the string representation
func (s UpdateAcceleratorAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAcceleratorAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAcceleratorAttributesInput"}

	if s.AcceleratorArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AcceleratorArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateAcceleratorAttributesOutput struct {
	_ struct{} `type:"structure"`

	// Updated attributes for the accelerator.
	AcceleratorAttributes *AcceleratorAttributes `type:"structure"`
}

// String returns the string representation
func (s UpdateAcceleratorAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateAcceleratorAttributes = "UpdateAcceleratorAttributes"

// UpdateAcceleratorAttributesRequest returns a request value for making API operation for
// AWS Global Accelerator.
//
// Update the attributes for an accelerator. To see an AWS CLI example of updating
// an accelerator to enable flow logs, scroll down to Example.
//
//    // Example sending a request using UpdateAcceleratorAttributesRequest.
//    req := client.UpdateAcceleratorAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/UpdateAcceleratorAttributes
func (c *Client) UpdateAcceleratorAttributesRequest(input *UpdateAcceleratorAttributesInput) UpdateAcceleratorAttributesRequest {
	op := &aws.Operation{
		Name:       opUpdateAcceleratorAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateAcceleratorAttributesInput{}
	}

	req := c.newRequest(op, input, &UpdateAcceleratorAttributesOutput{})

	return UpdateAcceleratorAttributesRequest{Request: req, Input: input, Copy: c.UpdateAcceleratorAttributesRequest}
}

// UpdateAcceleratorAttributesRequest is the request type for the
// UpdateAcceleratorAttributes API operation.
type UpdateAcceleratorAttributesRequest struct {
	*aws.Request
	Input *UpdateAcceleratorAttributesInput
	Copy  func(*UpdateAcceleratorAttributesInput) UpdateAcceleratorAttributesRequest
}

// Send marshals and sends the UpdateAcceleratorAttributes API request.
func (r UpdateAcceleratorAttributesRequest) Send(ctx context.Context) (*UpdateAcceleratorAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAcceleratorAttributesResponse{
		UpdateAcceleratorAttributesOutput: r.Request.Data.(*UpdateAcceleratorAttributesOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAcceleratorAttributesResponse is the response type for the
// UpdateAcceleratorAttributes API operation.
type UpdateAcceleratorAttributesResponse struct {
	*UpdateAcceleratorAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAcceleratorAttributes request.
func (r *UpdateAcceleratorAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
