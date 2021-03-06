// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateBGPPeerInput struct {
	_ struct{} `type:"structure"`

	// Information about the BGP peer.
	NewBGPPeer *NewBGPPeer `locationName:"newBGPPeer" type:"structure"`

	// The ID of the virtual interface.
	VirtualInterfaceId *string `locationName:"virtualInterfaceId" type:"string"`
}

// String returns the string representation
func (s CreateBGPPeerInput) String() string {
	return awsutil.Prettify(s)
}

type CreateBGPPeerOutput struct {
	_ struct{} `type:"structure"`

	// The virtual interface.
	VirtualInterface *VirtualInterface `locationName:"virtualInterface" type:"structure"`
}

// String returns the string representation
func (s CreateBGPPeerOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateBGPPeer = "CreateBGPPeer"

// CreateBGPPeerRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Creates a BGP peer on the specified virtual interface.
//
// You must create a BGP peer for the corresponding address family (IPv4/IPv6)
// in order to access AWS resources that also use that address family.
//
// If logical redundancy is not supported by the connection, interconnect, or
// LAG, the BGP peer cannot be in the same address family as an existing BGP
// peer on the virtual interface.
//
// When creating a IPv6 BGP peer, omit the Amazon address and customer address.
// IPv6 addresses are automatically assigned from the Amazon pool of IPv6 addresses;
// you cannot specify custom IPv6 addresses.
//
// For a public virtual interface, the Autonomous System Number (ASN) must be
// private or already whitelisted for the virtual interface.
//
//    // Example sending a request using CreateBGPPeerRequest.
//    req := client.CreateBGPPeerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/CreateBGPPeer
func (c *Client) CreateBGPPeerRequest(input *CreateBGPPeerInput) CreateBGPPeerRequest {
	op := &aws.Operation{
		Name:       opCreateBGPPeer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateBGPPeerInput{}
	}

	req := c.newRequest(op, input, &CreateBGPPeerOutput{})

	return CreateBGPPeerRequest{Request: req, Input: input, Copy: c.CreateBGPPeerRequest}
}

// CreateBGPPeerRequest is the request type for the
// CreateBGPPeer API operation.
type CreateBGPPeerRequest struct {
	*aws.Request
	Input *CreateBGPPeerInput
	Copy  func(*CreateBGPPeerInput) CreateBGPPeerRequest
}

// Send marshals and sends the CreateBGPPeer API request.
func (r CreateBGPPeerRequest) Send(ctx context.Context) (*CreateBGPPeerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateBGPPeerResponse{
		CreateBGPPeerOutput: r.Request.Data.(*CreateBGPPeerOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateBGPPeerResponse is the response type for the
// CreateBGPPeer API operation.
type CreateBGPPeerResponse struct {
	*CreateBGPPeerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateBGPPeer request.
func (r *CreateBGPPeerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
