// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type SetSubnetsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the load balancer.
	//
	// LoadBalancerArn is a required field
	LoadBalancerArn *string `type:"string" required:"true"`

	// The IDs of the public subnets. You can specify only one subnet per Availability
	// Zone. You must specify either subnets or subnet mappings.
	//
	// [Application Load Balancers] You must specify subnets from at least two Availability
	// Zones. You cannot specify Elastic IP addresses for your subnets.
	//
	// [Network Load Balancers] You can specify subnets from one or more Availability
	// Zones. If you need static IP addresses for your internet-facing load balancer,
	// you can specify one Elastic IP address per subnet. For internal load balancers,
	// you can specify one private IP address per subnet from the IPv4 range of
	// the subnet.
	SubnetMappings []SubnetMapping `type:"list"`

	// The IDs of the public subnets. You must specify subnets from at least two
	// Availability Zones. You can specify only one subnet per Availability Zone.
	// You must specify either subnets or subnet mappings.
	Subnets []string `type:"list"`
}

// String returns the string representation
func (s SetSubnetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetSubnetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetSubnetsInput"}

	if s.LoadBalancerArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SetSubnetsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the subnet and Availability Zone.
	AvailabilityZones []AvailabilityZone `type:"list"`
}

// String returns the string representation
func (s SetSubnetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetSubnets = "SetSubnets"

// SetSubnetsRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Enables the Availability Zones for the specified public subnets for the specified
// load balancer. The specified subnets replace the previously enabled subnets.
//
// When you specify subnets for a Network Load Balancer, you must include all
// subnets that were enabled previously, with their existing configurations,
// plus any additional subnets.
//
//    // Example sending a request using SetSubnetsRequest.
//    req := client.SetSubnetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/SetSubnets
func (c *Client) SetSubnetsRequest(input *SetSubnetsInput) SetSubnetsRequest {
	op := &aws.Operation{
		Name:       opSetSubnets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetSubnetsInput{}
	}

	req := c.newRequest(op, input, &SetSubnetsOutput{})

	return SetSubnetsRequest{Request: req, Input: input, Copy: c.SetSubnetsRequest}
}

// SetSubnetsRequest is the request type for the
// SetSubnets API operation.
type SetSubnetsRequest struct {
	*aws.Request
	Input *SetSubnetsInput
	Copy  func(*SetSubnetsInput) SetSubnetsRequest
}

// Send marshals and sends the SetSubnets API request.
func (r SetSubnetsRequest) Send(ctx context.Context) (*SetSubnetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetSubnetsResponse{
		SetSubnetsOutput: r.Request.Data.(*SetSubnetsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetSubnetsResponse is the response type for the
// SetSubnets API operation.
type SetSubnetsResponse struct {
	*SetSubnetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetSubnets request.
func (r *SetSubnetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
