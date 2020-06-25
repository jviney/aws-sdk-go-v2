// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type AddApplicationVpcConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The name of an existing application.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// The version of the application to which you want to add the input processing
	// configuration. You can use the DescribeApplication operation to get the current
	// application version. If the version specified is not the current version,
	// the ConcurrentModificationException is returned.
	//
	// CurrentApplicationVersionId is a required field
	CurrentApplicationVersionId *int64 `min:"1" type:"long" required:"true"`

	// Description of the VPC to add to the application.
	//
	// VpcConfiguration is a required field
	VpcConfiguration *VpcConfiguration `type:"structure" required:"true"`
}

// String returns the string representation
func (s AddApplicationVpcConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddApplicationVpcConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddApplicationVpcConfigurationInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if s.CurrentApplicationVersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CurrentApplicationVersionId"))
	}
	if s.CurrentApplicationVersionId != nil && *s.CurrentApplicationVersionId < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("CurrentApplicationVersionId", 1))
	}

	if s.VpcConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcConfiguration"))
	}
	if s.VpcConfiguration != nil {
		if err := s.VpcConfiguration.Validate(); err != nil {
			invalidParams.AddNested("VpcConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AddApplicationVpcConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the application.
	ApplicationARN *string `min:"1" type:"string"`

	// Provides the current application version. Kinesis Data Analytics updates
	// the ApplicationVersionId each time you update the application.
	ApplicationVersionId *int64 `min:"1" type:"long"`

	// The parameters of the new VPC configuration.
	VpcConfigurationDescription *VpcConfigurationDescription `type:"structure"`
}

// String returns the string representation
func (s AddApplicationVpcConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddApplicationVpcConfiguration = "AddApplicationVpcConfiguration"

// AddApplicationVpcConfigurationRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
// Adds a Virtual Private Cloud (VPC) configuration to the application. Applications
// can use VPCs to store and access resources securely.
//
// Note the following about VPC configurations for Kinesis Data Analytics applications:
//
//    * VPC configurations are not supported for SQL applications.
//
//    * When a VPC is added to a Kinesis Data Analytics application, the application
//    can no longer be accessed from the Internet directly. To enable Internet
//    access to the application, add an Internet gateway to your VPC.
//
//    // Example sending a request using AddApplicationVpcConfigurationRequest.
//    req := client.AddApplicationVpcConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/AddApplicationVpcConfiguration
func (c *Client) AddApplicationVpcConfigurationRequest(input *AddApplicationVpcConfigurationInput) AddApplicationVpcConfigurationRequest {
	op := &aws.Operation{
		Name:       opAddApplicationVpcConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddApplicationVpcConfigurationInput{}
	}

	req := c.newRequest(op, input, &AddApplicationVpcConfigurationOutput{})

	return AddApplicationVpcConfigurationRequest{Request: req, Input: input, Copy: c.AddApplicationVpcConfigurationRequest}
}

// AddApplicationVpcConfigurationRequest is the request type for the
// AddApplicationVpcConfiguration API operation.
type AddApplicationVpcConfigurationRequest struct {
	*aws.Request
	Input *AddApplicationVpcConfigurationInput
	Copy  func(*AddApplicationVpcConfigurationInput) AddApplicationVpcConfigurationRequest
}

// Send marshals and sends the AddApplicationVpcConfiguration API request.
func (r AddApplicationVpcConfigurationRequest) Send(ctx context.Context) (*AddApplicationVpcConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddApplicationVpcConfigurationResponse{
		AddApplicationVpcConfigurationOutput: r.Request.Data.(*AddApplicationVpcConfigurationOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddApplicationVpcConfigurationResponse is the response type for the
// AddApplicationVpcConfiguration API operation.
type AddApplicationVpcConfigurationResponse struct {
	*AddApplicationVpcConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddApplicationVpcConfiguration request.
func (r *AddApplicationVpcConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
