// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type CreateVirtualNodeInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `locationName:"clientToken" type:"string" idempotencyToken:"true"`

	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	MeshOwner *string `location:"querystring" locationName:"meshOwner" min:"12" type:"string"`

	// An object that represents the specification of a virtual node.
	//
	// Spec is a required field
	Spec *VirtualNodeSpec `locationName:"spec" type:"structure" required:"true"`

	Tags []TagRef `locationName:"tags" type:"list"`

	// VirtualNodeName is a required field
	VirtualNodeName *string `locationName:"virtualNodeName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateVirtualNodeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVirtualNodeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVirtualNodeInput"}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}
	if s.MeshOwner != nil && len(*s.MeshOwner) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshOwner", 12))
	}

	if s.Spec == nil {
		invalidParams.Add(aws.NewErrParamRequired("Spec"))
	}

	if s.VirtualNodeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualNodeName"))
	}
	if s.VirtualNodeName != nil && len(*s.VirtualNodeName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VirtualNodeName", 1))
	}
	if s.Spec != nil {
		if err := s.Spec.Validate(); err != nil {
			invalidParams.AddNested("Spec", err.(aws.ErrInvalidParams))
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

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateVirtualNodeInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	var ClientToken string
	if s.ClientToken != nil {
		ClientToken = *s.ClientToken
	} else {
		ClientToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Spec != nil {
		v := s.Spec

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "spec", v, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.VirtualNodeName != nil {
		v := *s.VirtualNodeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "virtualNodeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MeshName != nil {
		v := *s.MeshName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meshName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MeshOwner != nil {
		v := *s.MeshOwner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "meshOwner", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateVirtualNodeOutput struct {
	_ struct{} `type:"structure" payload:"VirtualNode"`

	// An object that represents a virtual node returned by a describe operation.
	//
	// VirtualNode is a required field
	VirtualNode *VirtualNodeData `locationName:"virtualNode" type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateVirtualNodeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateVirtualNodeOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.VirtualNode != nil {
		v := s.VirtualNode

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "virtualNode", v, metadata)
	}
	return nil
}

const opCreateVirtualNode = "CreateVirtualNode"

// CreateVirtualNodeRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Creates a virtual node within a service mesh.
//
// A virtual node acts as a logical pointer to a particular task group, such
// as an Amazon ECS service or a Kubernetes deployment. When you create a virtual
// node, you can specify the service discovery information for your task group,
// and whether the proxy running in a task group will communicate with other
// proxies using Transport Layer Security (TLS).
//
// You define a listener for any inbound traffic that your virtual node expects.
// Any virtual service that your virtual node expects to communicate to is specified
// as a backend.
//
// The response metadata for your new virtual node contains the arn that is
// associated with the virtual node. Set this value (either the full ARN or
// the truncated resource name: for example, mesh/default/virtualNode/simpleapp)
// as the APPMESH_VIRTUAL_NODE_NAME environment variable for your task group's
// Envoy proxy container in your task definition or pod spec. This is then mapped
// to the node.id and node.cluster Envoy parameters.
//
// If you require your Envoy stats or tracing to use a different name, you can
// override the node.cluster value that is set by APPMESH_VIRTUAL_NODE_NAME
// with the APPMESH_VIRTUAL_NODE_CLUSTER environment variable.
//
// For more information about virtual nodes, see Virtual nodes (https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_nodes.html).
//
//    // Example sending a request using CreateVirtualNodeRequest.
//    req := client.CreateVirtualNodeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/CreateVirtualNode
func (c *Client) CreateVirtualNodeRequest(input *CreateVirtualNodeInput) CreateVirtualNodeRequest {
	op := &aws.Operation{
		Name:       opCreateVirtualNode,
		HTTPMethod: "PUT",
		HTTPPath:   "/v20190125/meshes/{meshName}/virtualNodes",
	}

	if input == nil {
		input = &CreateVirtualNodeInput{}
	}

	req := c.newRequest(op, input, &CreateVirtualNodeOutput{})

	return CreateVirtualNodeRequest{Request: req, Input: input, Copy: c.CreateVirtualNodeRequest}
}

// CreateVirtualNodeRequest is the request type for the
// CreateVirtualNode API operation.
type CreateVirtualNodeRequest struct {
	*aws.Request
	Input *CreateVirtualNodeInput
	Copy  func(*CreateVirtualNodeInput) CreateVirtualNodeRequest
}

// Send marshals and sends the CreateVirtualNode API request.
func (r CreateVirtualNodeRequest) Send(ctx context.Context) (*CreateVirtualNodeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVirtualNodeResponse{
		CreateVirtualNodeOutput: r.Request.Data.(*CreateVirtualNodeOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVirtualNodeResponse is the response type for the
// CreateVirtualNode API operation.
type CreateVirtualNodeResponse struct {
	*CreateVirtualNodeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVirtualNode request.
func (r *CreateVirtualNodeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
