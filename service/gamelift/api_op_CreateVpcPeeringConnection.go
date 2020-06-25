// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type CreateVpcPeeringConnectionInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for a fleet. You can use either the fleet ID or ARN value.
	// This tells Amazon GameLift which GameLift VPC to peer with.
	//
	// FleetId is a required field
	FleetId *string `type:"string" required:"true"`

	// A unique identifier for the AWS account with the VPC that you want to peer
	// your Amazon GameLift fleet with. You can find your Account ID in the AWS
	// Management Console under account settings.
	//
	// PeerVpcAwsAccountId is a required field
	PeerVpcAwsAccountId *string `min:"1" type:"string" required:"true"`

	// A unique identifier for a VPC with resources to be accessed by your Amazon
	// GameLift fleet. The VPC must be in the same Region where your fleet is deployed.
	// Look up a VPC ID using the VPC Dashboard (https://console.aws.amazon.com/vpc/)
	// in the AWS Management Console. Learn more about VPC peering in VPC Peering
	// with Amazon GameLift Fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html).
	//
	// PeerVpcId is a required field
	PeerVpcId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateVpcPeeringConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVpcPeeringConnectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVpcPeeringConnectionInput"}

	if s.FleetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetId"))
	}

	if s.PeerVpcAwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PeerVpcAwsAccountId"))
	}
	if s.PeerVpcAwsAccountId != nil && len(*s.PeerVpcAwsAccountId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PeerVpcAwsAccountId", 1))
	}

	if s.PeerVpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PeerVpcId"))
	}
	if s.PeerVpcId != nil && len(*s.PeerVpcId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PeerVpcId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateVpcPeeringConnectionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateVpcPeeringConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateVpcPeeringConnection = "CreateVpcPeeringConnection"

// CreateVpcPeeringConnectionRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Establishes a VPC peering connection between a virtual private cloud (VPC)
// in an AWS account with the VPC for your Amazon GameLift fleet. VPC peering
// enables the game servers on your fleet to communicate directly with other
// AWS resources. You can peer with VPCs in any AWS account that you have access
// to, including the account that you use to manage your Amazon GameLift fleets.
// You cannot peer with VPCs that are in different Regions. For more information,
// see VPC Peering with Amazon GameLift Fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html).
//
// Before calling this operation to establish the peering connection, you first
// need to call CreateVpcPeeringAuthorization and identify the VPC you want
// to peer with. Once the authorization for the specified VPC is issued, you
// have 24 hours to establish the connection. These two operations handle all
// tasks necessary to peer the two VPCs, including acceptance, updating routing
// tables, etc.
//
// To establish the connection, call this operation from the AWS account that
// is used to manage the Amazon GameLift fleets. Identify the following values:
// (1) The ID of the fleet you want to be enable a VPC peering connection for;
// (2) The AWS account with the VPC that you want to peer with; and (3) The
// ID of the VPC you want to peer with. This operation is asynchronous. If successful,
// a VpcPeeringConnection request is created. You can use continuous polling
// to track the request's status using DescribeVpcPeeringConnections, or by
// monitoring fleet events for success or failure using DescribeFleetEvents.
//
//    * CreateVpcPeeringAuthorization
//
//    * DescribeVpcPeeringAuthorizations
//
//    * DeleteVpcPeeringAuthorization
//
//    * CreateVpcPeeringConnection
//
//    * DescribeVpcPeeringConnections
//
//    * DeleteVpcPeeringConnection
//
//    // Example sending a request using CreateVpcPeeringConnectionRequest.
//    req := client.CreateVpcPeeringConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/CreateVpcPeeringConnection
func (c *Client) CreateVpcPeeringConnectionRequest(input *CreateVpcPeeringConnectionInput) CreateVpcPeeringConnectionRequest {
	op := &aws.Operation{
		Name:       opCreateVpcPeeringConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVpcPeeringConnectionInput{}
	}

	req := c.newRequest(op, input, &CreateVpcPeeringConnectionOutput{})

	return CreateVpcPeeringConnectionRequest{Request: req, Input: input, Copy: c.CreateVpcPeeringConnectionRequest}
}

// CreateVpcPeeringConnectionRequest is the request type for the
// CreateVpcPeeringConnection API operation.
type CreateVpcPeeringConnectionRequest struct {
	*aws.Request
	Input *CreateVpcPeeringConnectionInput
	Copy  func(*CreateVpcPeeringConnectionInput) CreateVpcPeeringConnectionRequest
}

// Send marshals and sends the CreateVpcPeeringConnection API request.
func (r CreateVpcPeeringConnectionRequest) Send(ctx context.Context) (*CreateVpcPeeringConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVpcPeeringConnectionResponse{
		CreateVpcPeeringConnectionOutput: r.Request.Data.(*CreateVpcPeeringConnectionOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVpcPeeringConnectionResponse is the response type for the
// CreateVpcPeeringConnection API operation.
type CreateVpcPeeringConnectionResponse struct {
	*CreateVpcPeeringConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVpcPeeringConnection request.
func (r *CreateVpcPeeringConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
