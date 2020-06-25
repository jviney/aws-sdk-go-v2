// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type AllocatePublicVirtualInterfaceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the connection on which the public virtual interface is provisioned.
	//
	// ConnectionId is a required field
	ConnectionId *string `locationName:"connectionId" type:"string" required:"true"`

	// Information about the public virtual interface.
	//
	// NewPublicVirtualInterfaceAllocation is a required field
	NewPublicVirtualInterfaceAllocation *NewPublicVirtualInterfaceAllocation `locationName:"newPublicVirtualInterfaceAllocation" type:"structure" required:"true"`

	// The ID of the AWS account that owns the public virtual interface.
	//
	// OwnerAccount is a required field
	OwnerAccount *string `locationName:"ownerAccount" type:"string" required:"true"`
}

// String returns the string representation
func (s AllocatePublicVirtualInterfaceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AllocatePublicVirtualInterfaceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AllocatePublicVirtualInterfaceInput"}

	if s.ConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionId"))
	}

	if s.NewPublicVirtualInterfaceAllocation == nil {
		invalidParams.Add(aws.NewErrParamRequired("NewPublicVirtualInterfaceAllocation"))
	}

	if s.OwnerAccount == nil {
		invalidParams.Add(aws.NewErrParamRequired("OwnerAccount"))
	}
	if s.NewPublicVirtualInterfaceAllocation != nil {
		if err := s.NewPublicVirtualInterfaceAllocation.Validate(); err != nil {
			invalidParams.AddNested("NewPublicVirtualInterfaceAllocation", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about a virtual interface.
type AllocatePublicVirtualInterfaceOutput struct {
	_ struct{} `type:"structure"`

	// The address family for the BGP peer.
	AddressFamily AddressFamily `locationName:"addressFamily" type:"string" enum:"true"`

	// The IP address assigned to the Amazon interface.
	AmazonAddress *string `locationName:"amazonAddress" type:"string"`

	// The autonomous system number (ASN) for the Amazon side of the connection.
	AmazonSideAsn *int64 `locationName:"amazonSideAsn" type:"long"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	//
	// The valid values are 1-2147483647.
	Asn *int64 `locationName:"asn" type:"integer"`

	// The authentication key for BGP configuration. This string has a minimum length
	// of 6 characters and and a maximun lenth of 80 characters.
	AuthKey *string `locationName:"authKey" type:"string"`

	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDeviceV2 *string `locationName:"awsDeviceV2" type:"string"`

	// The BGP peers configured on this virtual interface.
	BgpPeers []BGPPeer `locationName:"bgpPeers" type:"list"`

	// The ID of the connection.
	ConnectionId *string `locationName:"connectionId" type:"string"`

	// The IP address assigned to the customer interface.
	CustomerAddress *string `locationName:"customerAddress" type:"string"`

	// The customer router configuration.
	CustomerRouterConfig *string `locationName:"customerRouterConfig" type:"string"`

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string `locationName:"directConnectGatewayId" type:"string"`

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool `locationName:"jumboFrameCapable" type:"boolean"`

	// The location of the connection.
	Location *string `locationName:"location" type:"string"`

	// The maximum transmission unit (MTU), in bytes. The supported values are 1500
	// and 9001. The default value is 1500.
	Mtu *int64 `locationName:"mtu" type:"integer"`

	// The ID of the AWS account that owns the virtual interface.
	OwnerAccount *string `locationName:"ownerAccount" type:"string"`

	// The AWS Region where the virtual interface is located.
	Region *string `locationName:"region" type:"string"`

	// The routes to be advertised to the AWS network in this Region. Applies to
	// public virtual interfaces.
	RouteFilterPrefixes []RouteFilterPrefix `locationName:"routeFilterPrefixes" type:"list"`

	// The tags associated with the virtual interface.
	Tags []Tag `locationName:"tags" min:"1" type:"list"`

	// The ID of the virtual private gateway. Applies only to private virtual interfaces.
	VirtualGatewayId *string `locationName:"virtualGatewayId" type:"string"`

	// The ID of the virtual interface.
	VirtualInterfaceId *string `locationName:"virtualInterfaceId" type:"string"`

	// The name of the virtual interface assigned by the customer network.
	VirtualInterfaceName *string `locationName:"virtualInterfaceName" type:"string"`

	// The state of the virtual interface. The following are the possible values:
	//
	//    * confirming: The creation of the virtual interface is pending confirmation
	//    from the virtual interface owner. If the owner of the virtual interface
	//    is different from the owner of the connection on which it is provisioned,
	//    then the virtual interface will remain in this state until it is confirmed
	//    by the virtual interface owner.
	//
	//    * verifying: This state only applies to public virtual interfaces. Each
	//    public virtual interface needs validation before the virtual interface
	//    can be created.
	//
	//    * pending: A virtual interface is in this state from the time that it
	//    is created until the virtual interface is ready to forward traffic.
	//
	//    * available: A virtual interface that is able to forward traffic.
	//
	//    * down: A virtual interface that is BGP down.
	//
	//    * deleting: A virtual interface is in this state immediately after calling
	//    DeleteVirtualInterface until it can no longer forward traffic.
	//
	//    * deleted: A virtual interface that cannot forward traffic.
	//
	//    * rejected: The virtual interface owner has declined creation of the virtual
	//    interface. If a virtual interface in the Confirming state is deleted by
	//    the virtual interface owner, the virtual interface enters the Rejected
	//    state.
	//
	//    * unknown: The state of the virtual interface is not available.
	VirtualInterfaceState VirtualInterfaceState `locationName:"virtualInterfaceState" type:"string" enum:"true"`

	// The type of virtual interface. The possible values are private and public.
	VirtualInterfaceType *string `locationName:"virtualInterfaceType" type:"string"`

	// The ID of the VLAN.
	Vlan *int64 `locationName:"vlan" type:"integer"`
}

// String returns the string representation
func (s AllocatePublicVirtualInterfaceOutput) String() string {
	return awsutil.Prettify(s)
}

const opAllocatePublicVirtualInterface = "AllocatePublicVirtualInterface"

// AllocatePublicVirtualInterfaceRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Provisions a public virtual interface to be owned by the specified AWS account.
//
// The owner of a connection calls this function to provision a public virtual
// interface to be owned by the specified AWS account.
//
// Virtual interfaces created using this function must be confirmed by the owner
// using ConfirmPublicVirtualInterface. Until this step has been completed,
// the virtual interface is in the confirming state and is not available to
// handle traffic.
//
// When creating an IPv6 public virtual interface, omit the Amazon address and
// customer address. IPv6 addresses are automatically assigned from the Amazon
// pool of IPv6 addresses; you cannot specify custom IPv6 addresses.
//
//    // Example sending a request using AllocatePublicVirtualInterfaceRequest.
//    req := client.AllocatePublicVirtualInterfaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/AllocatePublicVirtualInterface
func (c *Client) AllocatePublicVirtualInterfaceRequest(input *AllocatePublicVirtualInterfaceInput) AllocatePublicVirtualInterfaceRequest {
	op := &aws.Operation{
		Name:       opAllocatePublicVirtualInterface,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AllocatePublicVirtualInterfaceInput{}
	}

	req := c.newRequest(op, input, &AllocatePublicVirtualInterfaceOutput{})

	return AllocatePublicVirtualInterfaceRequest{Request: req, Input: input, Copy: c.AllocatePublicVirtualInterfaceRequest}
}

// AllocatePublicVirtualInterfaceRequest is the request type for the
// AllocatePublicVirtualInterface API operation.
type AllocatePublicVirtualInterfaceRequest struct {
	*aws.Request
	Input *AllocatePublicVirtualInterfaceInput
	Copy  func(*AllocatePublicVirtualInterfaceInput) AllocatePublicVirtualInterfaceRequest
}

// Send marshals and sends the AllocatePublicVirtualInterface API request.
func (r AllocatePublicVirtualInterfaceRequest) Send(ctx context.Context) (*AllocatePublicVirtualInterfaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AllocatePublicVirtualInterfaceResponse{
		AllocatePublicVirtualInterfaceOutput: r.Request.Data.(*AllocatePublicVirtualInterfaceOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AllocatePublicVirtualInterfaceResponse is the response type for the
// AllocatePublicVirtualInterface API operation.
type AllocatePublicVirtualInterfaceResponse struct {
	*AllocatePublicVirtualInterfaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AllocatePublicVirtualInterface request.
func (r *AllocatePublicVirtualInterfaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
