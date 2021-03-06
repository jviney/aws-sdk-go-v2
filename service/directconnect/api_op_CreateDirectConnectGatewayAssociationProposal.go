// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateDirectConnectGatewayAssociationProposalInput struct {
	_ struct{} `type:"structure"`

	// The Amazon VPC prefixes to advertise to the Direct Connect gateway.
	AddAllowedPrefixesToDirectConnectGateway []RouteFilterPrefix `locationName:"addAllowedPrefixesToDirectConnectGateway" type:"list"`

	// The ID of the Direct Connect gateway.
	//
	// DirectConnectGatewayId is a required field
	DirectConnectGatewayId *string `locationName:"directConnectGatewayId" type:"string" required:"true"`

	// The ID of the AWS account that owns the Direct Connect gateway.
	//
	// DirectConnectGatewayOwnerAccount is a required field
	DirectConnectGatewayOwnerAccount *string `locationName:"directConnectGatewayOwnerAccount" type:"string" required:"true"`

	// The ID of the virtual private gateway or transit gateway.
	//
	// GatewayId is a required field
	GatewayId *string `locationName:"gatewayId" type:"string" required:"true"`

	// The Amazon VPC prefixes to no longer advertise to the Direct Connect gateway.
	RemoveAllowedPrefixesToDirectConnectGateway []RouteFilterPrefix `locationName:"removeAllowedPrefixesToDirectConnectGateway" type:"list"`
}

// String returns the string representation
func (s CreateDirectConnectGatewayAssociationProposalInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDirectConnectGatewayAssociationProposalInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDirectConnectGatewayAssociationProposalInput"}

	if s.DirectConnectGatewayId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectConnectGatewayId"))
	}

	if s.DirectConnectGatewayOwnerAccount == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectConnectGatewayOwnerAccount"))
	}

	if s.GatewayId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDirectConnectGatewayAssociationProposalOutput struct {
	_ struct{} `type:"structure"`

	// Information about the Direct Connect gateway proposal.
	DirectConnectGatewayAssociationProposal *DirectConnectGatewayAssociationProposal `locationName:"directConnectGatewayAssociationProposal" type:"structure"`
}

// String returns the string representation
func (s CreateDirectConnectGatewayAssociationProposalOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDirectConnectGatewayAssociationProposal = "CreateDirectConnectGatewayAssociationProposal"

// CreateDirectConnectGatewayAssociationProposalRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Creates a proposal to associate the specified virtual private gateway or
// transit gateway with the specified Direct Connect gateway.
//
// You can only associate a Direct Connect gateway and virtual private gateway
// or transit gateway when the account that owns the Direct Connect gateway
// and the account that owns the virtual private gateway or transit gateway
// have the same AWS Payer ID.
//
//    // Example sending a request using CreateDirectConnectGatewayAssociationProposalRequest.
//    req := client.CreateDirectConnectGatewayAssociationProposalRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/CreateDirectConnectGatewayAssociationProposal
func (c *Client) CreateDirectConnectGatewayAssociationProposalRequest(input *CreateDirectConnectGatewayAssociationProposalInput) CreateDirectConnectGatewayAssociationProposalRequest {
	op := &aws.Operation{
		Name:       opCreateDirectConnectGatewayAssociationProposal,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDirectConnectGatewayAssociationProposalInput{}
	}

	req := c.newRequest(op, input, &CreateDirectConnectGatewayAssociationProposalOutput{})

	return CreateDirectConnectGatewayAssociationProposalRequest{Request: req, Input: input, Copy: c.CreateDirectConnectGatewayAssociationProposalRequest}
}

// CreateDirectConnectGatewayAssociationProposalRequest is the request type for the
// CreateDirectConnectGatewayAssociationProposal API operation.
type CreateDirectConnectGatewayAssociationProposalRequest struct {
	*aws.Request
	Input *CreateDirectConnectGatewayAssociationProposalInput
	Copy  func(*CreateDirectConnectGatewayAssociationProposalInput) CreateDirectConnectGatewayAssociationProposalRequest
}

// Send marshals and sends the CreateDirectConnectGatewayAssociationProposal API request.
func (r CreateDirectConnectGatewayAssociationProposalRequest) Send(ctx context.Context) (*CreateDirectConnectGatewayAssociationProposalResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDirectConnectGatewayAssociationProposalResponse{
		CreateDirectConnectGatewayAssociationProposalOutput: r.Request.Data.(*CreateDirectConnectGatewayAssociationProposalOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDirectConnectGatewayAssociationProposalResponse is the response type for the
// CreateDirectConnectGatewayAssociationProposal API operation.
type CreateDirectConnectGatewayAssociationProposalResponse struct {
	*CreateDirectConnectGatewayAssociationProposalOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDirectConnectGatewayAssociationProposal request.
func (r *CreateDirectConnectGatewayAssociationProposalResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
