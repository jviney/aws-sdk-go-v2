// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateLagInput struct {
	_ struct{} `type:"structure"`

	// The tags to associate with the automtically created LAGs.
	ChildConnectionTags []Tag `locationName:"childConnectionTags" min:"1" type:"list"`

	// The ID of an existing connection to migrate to the LAG.
	ConnectionId *string `locationName:"connectionId" type:"string"`

	// The bandwidth of the individual physical connections bundled by the LAG.
	// The possible values are 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps,
	// 1Gbps, 2Gbps, 5Gbps, and 10Gbps.
	//
	// ConnectionsBandwidth is a required field
	ConnectionsBandwidth *string `locationName:"connectionsBandwidth" type:"string" required:"true"`

	// The name of the LAG.
	//
	// LagName is a required field
	LagName *string `locationName:"lagName" type:"string" required:"true"`

	// The location for the LAG.
	//
	// Location is a required field
	Location *string `locationName:"location" type:"string" required:"true"`

	// The number of physical connections initially provisioned and bundled by the
	// LAG.
	//
	// NumberOfConnections is a required field
	NumberOfConnections *int64 `locationName:"numberOfConnections" type:"integer" required:"true"`

	// The name of the service provider associated with the LAG.
	ProviderName *string `locationName:"providerName" type:"string"`

	// The tags to associate with the LAG.
	Tags []Tag `locationName:"tags" min:"1" type:"list"`
}

// String returns the string representation
func (s CreateLagInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLagInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLagInput"}
	if s.ChildConnectionTags != nil && len(s.ChildConnectionTags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChildConnectionTags", 1))
	}

	if s.ConnectionsBandwidth == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionsBandwidth"))
	}

	if s.LagName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LagName"))
	}

	if s.Location == nil {
		invalidParams.Add(aws.NewErrParamRequired("Location"))
	}

	if s.NumberOfConnections == nil {
		invalidParams.Add(aws.NewErrParamRequired("NumberOfConnections"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.ChildConnectionTags != nil {
		for i, v := range s.ChildConnectionTags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ChildConnectionTags", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about a link aggregation group (LAG).
type CreateLagOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether the LAG can host other connections.
	AllowsHostedConnections *bool `locationName:"allowsHostedConnections" type:"boolean"`

	// The AWS Direct Connect endpoint that hosts the LAG.
	AwsDevice *string `locationName:"awsDevice" deprecated:"true" type:"string"`

	// The AWS Direct Connect endpoint that hosts the LAG.
	AwsDeviceV2 *string `locationName:"awsDeviceV2" type:"string"`

	// The connections bundled by the LAG.
	Connections []Connection `locationName:"connections" type:"list"`

	// The individual bandwidth of the physical connections bundled by the LAG.
	// The possible values are 1Gbps and 10Gbps.
	ConnectionsBandwidth *string `locationName:"connectionsBandwidth" type:"string"`

	// Indicates whether the LAG supports a secondary BGP peer in the same address
	// family (IPv4/IPv6).
	HasLogicalRedundancy HasLogicalRedundancy `locationName:"hasLogicalRedundancy" type:"string" enum:"true"`

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool `locationName:"jumboFrameCapable" type:"boolean"`

	// The ID of the LAG.
	LagId *string `locationName:"lagId" type:"string"`

	// The name of the LAG.
	LagName *string `locationName:"lagName" type:"string"`

	// The state of the LAG. The following are the possible values:
	//
	//    * requested: The initial state of a LAG. The LAG stays in the requested
	//    state until the Letter of Authorization (LOA) is available.
	//
	//    * pending: The LAG has been approved and is being initialized.
	//
	//    * available: The network link is established and the LAG is ready for
	//    use.
	//
	//    * down: The network link is down.
	//
	//    * deleting: The LAG is being deleted.
	//
	//    * deleted: The LAG is deleted.
	//
	//    * unknown: The state of the LAG is not available.
	LagState LagState `locationName:"lagState" type:"string" enum:"true"`

	// The location of the LAG.
	Location *string `locationName:"location" type:"string"`

	// The minimum number of physical connections that must be operational for the
	// LAG itself to be operational.
	MinimumLinks *int64 `locationName:"minimumLinks" type:"integer"`

	// The number of physical connections bundled by the LAG, up to a maximum of
	// 10.
	NumberOfConnections *int64 `locationName:"numberOfConnections" type:"integer"`

	// The ID of the AWS account that owns the LAG.
	OwnerAccount *string `locationName:"ownerAccount" type:"string"`

	// The name of the service provider associated with the LAG.
	ProviderName *string `locationName:"providerName" type:"string"`

	// The AWS Region where the connection is located.
	Region *string `locationName:"region" type:"string"`

	// The tags associated with the LAG.
	Tags []Tag `locationName:"tags" min:"1" type:"list"`
}

// String returns the string representation
func (s CreateLagOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateLag = "CreateLag"

// CreateLagRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Creates a link aggregation group (LAG) with the specified number of bundled
// physical connections between the customer network and a specific AWS Direct
// Connect location. A LAG is a logical interface that uses the Link Aggregation
// Control Protocol (LACP) to aggregate multiple interfaces, enabling you to
// treat them as a single interface.
//
// All connections in a LAG must use the same bandwidth and must terminate at
// the same AWS Direct Connect endpoint.
//
// You can have up to 10 connections per LAG. Regardless of this limit, if you
// request more connections for the LAG than AWS Direct Connect can allocate
// on a single endpoint, no LAG is created.
//
// You can specify an existing physical connection or interconnect to include
// in the LAG (which counts towards the total number of connections). Doing
// so interrupts the current physical connection or hosted connections, and
// re-establishes them as a member of the LAG. The LAG will be created on the
// same AWS Direct Connect endpoint to which the connection terminates. Any
// virtual interfaces associated with the connection are automatically disassociated
// and re-associated with the LAG. The connection ID does not change.
//
// If the AWS account used to create a LAG is a registered AWS Direct Connect
// Partner, the LAG is automatically enabled to host sub-connections. For a
// LAG owned by a partner, any associated virtual interfaces cannot be directly
// configured.
//
//    // Example sending a request using CreateLagRequest.
//    req := client.CreateLagRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/CreateLag
func (c *Client) CreateLagRequest(input *CreateLagInput) CreateLagRequest {
	op := &aws.Operation{
		Name:       opCreateLag,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateLagInput{}
	}

	req := c.newRequest(op, input, &CreateLagOutput{})

	return CreateLagRequest{Request: req, Input: input, Copy: c.CreateLagRequest}
}

// CreateLagRequest is the request type for the
// CreateLag API operation.
type CreateLagRequest struct {
	*aws.Request
	Input *CreateLagInput
	Copy  func(*CreateLagInput) CreateLagRequest
}

// Send marshals and sends the CreateLag API request.
func (r CreateLagRequest) Send(ctx context.Context) (*CreateLagResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLagResponse{
		CreateLagOutput: r.Request.Data.(*CreateLagOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLagResponse is the response type for the
// CreateLag API operation.
type CreateLagResponse struct {
	*CreateLagOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLag request.
func (r *CreateLagResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
