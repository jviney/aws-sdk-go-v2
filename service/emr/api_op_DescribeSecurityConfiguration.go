// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeSecurityConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The name of the security configuration.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeSecurityConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSecurityConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSecurityConfigurationInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeSecurityConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The date and time the security configuration was created
	CreationDateTime *time.Time `type:"timestamp"`

	// The name of the security configuration.
	Name *string `type:"string"`

	// The security configuration details in JSON format.
	SecurityConfiguration *string `type:"string"`
}

// String returns the string representation
func (s DescribeSecurityConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSecurityConfiguration = "DescribeSecurityConfiguration"

// DescribeSecurityConfigurationRequest returns a request value for making API operation for
// Amazon Elastic MapReduce.
//
// Provides the details of a security configuration by returning the configuration
// JSON.
//
//    // Example sending a request using DescribeSecurityConfigurationRequest.
//    req := client.DescribeSecurityConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/DescribeSecurityConfiguration
func (c *Client) DescribeSecurityConfigurationRequest(input *DescribeSecurityConfigurationInput) DescribeSecurityConfigurationRequest {
	op := &aws.Operation{
		Name:       opDescribeSecurityConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSecurityConfigurationInput{}
	}

	req := c.newRequest(op, input, &DescribeSecurityConfigurationOutput{})

	return DescribeSecurityConfigurationRequest{Request: req, Input: input, Copy: c.DescribeSecurityConfigurationRequest}
}

// DescribeSecurityConfigurationRequest is the request type for the
// DescribeSecurityConfiguration API operation.
type DescribeSecurityConfigurationRequest struct {
	*aws.Request
	Input *DescribeSecurityConfigurationInput
	Copy  func(*DescribeSecurityConfigurationInput) DescribeSecurityConfigurationRequest
}

// Send marshals and sends the DescribeSecurityConfiguration API request.
func (r DescribeSecurityConfigurationRequest) Send(ctx context.Context) (*DescribeSecurityConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSecurityConfigurationResponse{
		DescribeSecurityConfigurationOutput: r.Request.Data.(*DescribeSecurityConfigurationOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSecurityConfigurationResponse is the response type for the
// DescribeSecurityConfiguration API operation.
type DescribeSecurityConfigurationResponse struct {
	*DescribeSecurityConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSecurityConfiguration request.
func (r *DescribeSecurityConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
