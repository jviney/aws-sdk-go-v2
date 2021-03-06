// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeregisterTransitGatewayMulticastGroupMembersInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The IP address assigned to the transit gateway multicast group.
	GroupIpAddress *string `type:"string"`

	// The IDs of the group members' network interfaces.
	NetworkInterfaceIds []string `locationNameList:"item" type:"list"`

	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId *string `type:"string"`
}

// String returns the string representation
func (s DeregisterTransitGatewayMulticastGroupMembersInput) String() string {
	return awsutil.Prettify(s)
}

type DeregisterTransitGatewayMulticastGroupMembersOutput struct {
	_ struct{} `type:"structure"`

	// Information about the deregistered members.
	DeregisteredMulticastGroupMembers *TransitGatewayMulticastDeregisteredGroupMembers `locationName:"deregisteredMulticastGroupMembers" type:"structure"`
}

// String returns the string representation
func (s DeregisterTransitGatewayMulticastGroupMembersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeregisterTransitGatewayMulticastGroupMembers = "DeregisterTransitGatewayMulticastGroupMembers"

// DeregisterTransitGatewayMulticastGroupMembersRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deregisters the specified members (network interfaces) from the transit gateway
// multicast group.
//
//    // Example sending a request using DeregisterTransitGatewayMulticastGroupMembersRequest.
//    req := client.DeregisterTransitGatewayMulticastGroupMembersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeregisterTransitGatewayMulticastGroupMembers
func (c *Client) DeregisterTransitGatewayMulticastGroupMembersRequest(input *DeregisterTransitGatewayMulticastGroupMembersInput) DeregisterTransitGatewayMulticastGroupMembersRequest {
	op := &aws.Operation{
		Name:       opDeregisterTransitGatewayMulticastGroupMembers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeregisterTransitGatewayMulticastGroupMembersInput{}
	}

	req := c.newRequest(op, input, &DeregisterTransitGatewayMulticastGroupMembersOutput{})

	return DeregisterTransitGatewayMulticastGroupMembersRequest{Request: req, Input: input, Copy: c.DeregisterTransitGatewayMulticastGroupMembersRequest}
}

// DeregisterTransitGatewayMulticastGroupMembersRequest is the request type for the
// DeregisterTransitGatewayMulticastGroupMembers API operation.
type DeregisterTransitGatewayMulticastGroupMembersRequest struct {
	*aws.Request
	Input *DeregisterTransitGatewayMulticastGroupMembersInput
	Copy  func(*DeregisterTransitGatewayMulticastGroupMembersInput) DeregisterTransitGatewayMulticastGroupMembersRequest
}

// Send marshals and sends the DeregisterTransitGatewayMulticastGroupMembers API request.
func (r DeregisterTransitGatewayMulticastGroupMembersRequest) Send(ctx context.Context) (*DeregisterTransitGatewayMulticastGroupMembersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterTransitGatewayMulticastGroupMembersResponse{
		DeregisterTransitGatewayMulticastGroupMembersOutput: r.Request.Data.(*DeregisterTransitGatewayMulticastGroupMembersOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterTransitGatewayMulticastGroupMembersResponse is the response type for the
// DeregisterTransitGatewayMulticastGroupMembers API operation.
type DeregisterTransitGatewayMulticastGroupMembersResponse struct {
	*DeregisterTransitGatewayMulticastGroupMembersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterTransitGatewayMulticastGroupMembers request.
func (r *DeregisterTransitGatewayMulticastGroupMembersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
