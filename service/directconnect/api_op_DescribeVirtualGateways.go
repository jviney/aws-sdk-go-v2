// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeVirtualGatewaysInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeVirtualGatewaysInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeVirtualGatewaysOutput struct {
	_ struct{} `type:"structure"`

	// The virtual private gateways.
	VirtualGateways []VirtualGateway `locationName:"virtualGateways" type:"list"`
}

// String returns the string representation
func (s DescribeVirtualGatewaysOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeVirtualGateways = "DescribeVirtualGateways"

// DescribeVirtualGatewaysRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Lists the virtual private gateways owned by the AWS account.
//
// You can create one or more AWS Direct Connect private virtual interfaces
// linked to a virtual private gateway.
//
//    // Example sending a request using DescribeVirtualGatewaysRequest.
//    req := client.DescribeVirtualGatewaysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/DescribeVirtualGateways
func (c *Client) DescribeVirtualGatewaysRequest(input *DescribeVirtualGatewaysInput) DescribeVirtualGatewaysRequest {
	op := &aws.Operation{
		Name:       opDescribeVirtualGateways,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeVirtualGatewaysInput{}
	}

	req := c.newRequest(op, input, &DescribeVirtualGatewaysOutput{})

	return DescribeVirtualGatewaysRequest{Request: req, Input: input, Copy: c.DescribeVirtualGatewaysRequest}
}

// DescribeVirtualGatewaysRequest is the request type for the
// DescribeVirtualGateways API operation.
type DescribeVirtualGatewaysRequest struct {
	*aws.Request
	Input *DescribeVirtualGatewaysInput
	Copy  func(*DescribeVirtualGatewaysInput) DescribeVirtualGatewaysRequest
}

// Send marshals and sends the DescribeVirtualGateways API request.
func (r DescribeVirtualGatewaysRequest) Send(ctx context.Context) (*DescribeVirtualGatewaysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVirtualGatewaysResponse{
		DescribeVirtualGatewaysOutput: r.Request.Data.(*DescribeVirtualGatewaysOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeVirtualGatewaysResponse is the response type for the
// DescribeVirtualGateways API operation.
type DescribeVirtualGatewaysResponse struct {
	*DescribeVirtualGatewaysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVirtualGateways request.
func (r *DescribeVirtualGatewaysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
