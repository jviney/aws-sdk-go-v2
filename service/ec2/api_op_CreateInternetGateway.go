// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateInternetGatewayInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`
}

// String returns the string representation
func (s CreateInternetGatewayInput) String() string {
	return awsutil.Prettify(s)
}

type CreateInternetGatewayOutput struct {
	_ struct{} `type:"structure"`

	// Information about the internet gateway.
	InternetGateway *InternetGateway `locationName:"internetGateway" type:"structure"`
}

// String returns the string representation
func (s CreateInternetGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateInternetGateway = "CreateInternetGateway"

// CreateInternetGatewayRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates an internet gateway for use with a VPC. After creating the internet
// gateway, you attach it to a VPC using AttachInternetGateway.
//
// For more information about your VPC and internet gateway, see the Amazon
// Virtual Private Cloud User Guide (https://docs.aws.amazon.com/vpc/latest/userguide/).
//
//    // Example sending a request using CreateInternetGatewayRequest.
//    req := client.CreateInternetGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateInternetGateway
func (c *Client) CreateInternetGatewayRequest(input *CreateInternetGatewayInput) CreateInternetGatewayRequest {
	op := &aws.Operation{
		Name:       opCreateInternetGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateInternetGatewayInput{}
	}

	req := c.newRequest(op, input, &CreateInternetGatewayOutput{})

	return CreateInternetGatewayRequest{Request: req, Input: input, Copy: c.CreateInternetGatewayRequest}
}

// CreateInternetGatewayRequest is the request type for the
// CreateInternetGateway API operation.
type CreateInternetGatewayRequest struct {
	*aws.Request
	Input *CreateInternetGatewayInput
	Copy  func(*CreateInternetGatewayInput) CreateInternetGatewayRequest
}

// Send marshals and sends the CreateInternetGateway API request.
func (r CreateInternetGatewayRequest) Send(ctx context.Context) (*CreateInternetGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateInternetGatewayResponse{
		CreateInternetGatewayOutput: r.Request.Data.(*CreateInternetGatewayOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateInternetGatewayResponse is the response type for the
// CreateInternetGateway API operation.
type CreateInternetGatewayResponse struct {
	*CreateInternetGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateInternetGateway request.
func (r *CreateInternetGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
