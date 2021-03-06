// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/ec2query"
)

type AttachInternetGatewayInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the internet gateway.
	//
	// InternetGatewayId is a required field
	InternetGatewayId *string `locationName:"internetGatewayId" type:"string" required:"true"`

	// The ID of the VPC.
	//
	// VpcId is a required field
	VpcId *string `locationName:"vpcId" type:"string" required:"true"`
}

// String returns the string representation
func (s AttachInternetGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachInternetGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AttachInternetGatewayInput"}

	if s.InternetGatewayId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InternetGatewayId"))
	}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AttachInternetGatewayOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AttachInternetGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

const opAttachInternetGateway = "AttachInternetGateway"

// AttachInternetGatewayRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Attaches an internet gateway or a virtual private gateway to a VPC, enabling
// connectivity between the internet and the VPC. For more information about
// your VPC and internet gateway, see the Amazon Virtual Private Cloud User
// Guide (https://docs.aws.amazon.com/vpc/latest/userguide/).
//
//    // Example sending a request using AttachInternetGatewayRequest.
//    req := client.AttachInternetGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AttachInternetGateway
func (c *Client) AttachInternetGatewayRequest(input *AttachInternetGatewayInput) AttachInternetGatewayRequest {
	op := &aws.Operation{
		Name:       opAttachInternetGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachInternetGatewayInput{}
	}

	req := c.newRequest(op, input, &AttachInternetGatewayOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return AttachInternetGatewayRequest{Request: req, Input: input, Copy: c.AttachInternetGatewayRequest}
}

// AttachInternetGatewayRequest is the request type for the
// AttachInternetGateway API operation.
type AttachInternetGatewayRequest struct {
	*aws.Request
	Input *AttachInternetGatewayInput
	Copy  func(*AttachInternetGatewayInput) AttachInternetGatewayRequest
}

// Send marshals and sends the AttachInternetGateway API request.
func (r AttachInternetGatewayRequest) Send(ctx context.Context) (*AttachInternetGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachInternetGatewayResponse{
		AttachInternetGatewayOutput: r.Request.Data.(*AttachInternetGatewayOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachInternetGatewayResponse is the response type for the
// AttachInternetGateway API operation.
type AttachInternetGatewayResponse struct {
	*AttachInternetGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachInternetGateway request.
func (r *AttachInternetGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
