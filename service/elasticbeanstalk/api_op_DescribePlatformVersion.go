// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribePlatformVersionInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the platform version.
	PlatformArn *string `type:"string"`
}

// String returns the string representation
func (s DescribePlatformVersionInput) String() string {
	return awsutil.Prettify(s)
}

type DescribePlatformVersionOutput struct {
	_ struct{} `type:"structure"`

	// Detailed information about the platform version.
	PlatformDescription *PlatformDescription `type:"structure"`
}

// String returns the string representation
func (s DescribePlatformVersionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribePlatformVersion = "DescribePlatformVersion"

// DescribePlatformVersionRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Describes a platform version. Provides full details. Compare to ListPlatformVersions,
// which provides summary information about a list of platform versions.
//
// For definitions of platform version and other platform-related terms, see
// AWS Elastic Beanstalk Platforms Glossary (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/platforms-glossary.html).
//
//    // Example sending a request using DescribePlatformVersionRequest.
//    req := client.DescribePlatformVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/DescribePlatformVersion
func (c *Client) DescribePlatformVersionRequest(input *DescribePlatformVersionInput) DescribePlatformVersionRequest {
	op := &aws.Operation{
		Name:       opDescribePlatformVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribePlatformVersionInput{}
	}

	req := c.newRequest(op, input, &DescribePlatformVersionOutput{})

	return DescribePlatformVersionRequest{Request: req, Input: input, Copy: c.DescribePlatformVersionRequest}
}

// DescribePlatformVersionRequest is the request type for the
// DescribePlatformVersion API operation.
type DescribePlatformVersionRequest struct {
	*aws.Request
	Input *DescribePlatformVersionInput
	Copy  func(*DescribePlatformVersionInput) DescribePlatformVersionRequest
}

// Send marshals and sends the DescribePlatformVersion API request.
func (r DescribePlatformVersionRequest) Send(ctx context.Context) (*DescribePlatformVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePlatformVersionResponse{
		DescribePlatformVersionOutput: r.Request.Data.(*DescribePlatformVersionOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribePlatformVersionResponse is the response type for the
// DescribePlatformVersion API operation.
type DescribePlatformVersionResponse struct {
	*DescribePlatformVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePlatformVersion request.
func (r *DescribePlatformVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
