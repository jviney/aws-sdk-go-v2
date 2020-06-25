// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DescribeCustomerGateways.
type DescribeCustomerGatewaysInput struct {
	_ struct{} `type:"structure"`

	// One or more customer gateway IDs.
	//
	// Default: Describes all your customer gateways.
	CustomerGatewayIds []string `locationName:"CustomerGatewayId" locationNameList:"CustomerGatewayId" type:"list"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * bgp-asn - The customer gateway's Border Gateway Protocol (BGP) Autonomous
	//    System Number (ASN).
	//
	//    * customer-gateway-id - The ID of the customer gateway.
	//
	//    * ip-address - The IP address of the customer gateway's Internet-routable
	//    external interface.
	//
	//    * state - The state of the customer gateway (pending | available | deleting
	//    | deleted).
	//
	//    * type - The type of customer gateway. Currently, the only supported type
	//    is ipsec.1.
	//
	//    * tag:<key> - The key/value combination of a tag assigned to the resource.
	//    Use the tag key in the filter name and the tag value as the filter value.
	//    For example, to find all resources that have a tag with the key Owner
	//    and the value TeamA, specify tag:Owner for the filter name and TeamA for
	//    the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. Use this filter
	//    to find all resources assigned a tag with a specific key, regardless of
	//    the tag value.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`
}

// String returns the string representation
func (s DescribeCustomerGatewaysInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the output of DescribeCustomerGateways.
type DescribeCustomerGatewaysOutput struct {
	_ struct{} `type:"structure"`

	// Information about one or more customer gateways.
	CustomerGateways []CustomerGateway `locationName:"customerGatewaySet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeCustomerGatewaysOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCustomerGateways = "DescribeCustomerGateways"

// DescribeCustomerGatewaysRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more of your VPN customer gateways.
//
// For more information, see AWS Site-to-Site VPN (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html)
// in the AWS Site-to-Site VPN User Guide.
//
//    // Example sending a request using DescribeCustomerGatewaysRequest.
//    req := client.DescribeCustomerGatewaysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeCustomerGateways
func (c *Client) DescribeCustomerGatewaysRequest(input *DescribeCustomerGatewaysInput) DescribeCustomerGatewaysRequest {
	op := &aws.Operation{
		Name:       opDescribeCustomerGateways,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCustomerGatewaysInput{}
	}

	req := c.newRequest(op, input, &DescribeCustomerGatewaysOutput{})

	return DescribeCustomerGatewaysRequest{Request: req, Input: input, Copy: c.DescribeCustomerGatewaysRequest}
}

// DescribeCustomerGatewaysRequest is the request type for the
// DescribeCustomerGateways API operation.
type DescribeCustomerGatewaysRequest struct {
	*aws.Request
	Input *DescribeCustomerGatewaysInput
	Copy  func(*DescribeCustomerGatewaysInput) DescribeCustomerGatewaysRequest
}

// Send marshals and sends the DescribeCustomerGateways API request.
func (r DescribeCustomerGatewaysRequest) Send(ctx context.Context) (*DescribeCustomerGatewaysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCustomerGatewaysResponse{
		DescribeCustomerGatewaysOutput: r.Request.Data.(*DescribeCustomerGatewaysOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCustomerGatewaysResponse is the response type for the
// DescribeCustomerGateways API operation.
type DescribeCustomerGatewaysResponse struct {
	*DescribeCustomerGatewaysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCustomerGateways request.
func (r *DescribeCustomerGatewaysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
