// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DescribeNodegroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the Amazon EKS cluster associated with the node group.
	//
	// ClusterName is a required field
	ClusterName *string `location:"uri" locationName:"name" type:"string" required:"true"`

	// The name of the node group to describe.
	//
	// NodegroupName is a required field
	NodegroupName *string `location:"uri" locationName:"nodegroupName" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeNodegroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeNodegroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeNodegroupInput"}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}

	if s.NodegroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("NodegroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeNodegroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ClusterName != nil {
		v := *s.ClusterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NodegroupName != nil {
		v := *s.NodegroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "nodegroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeNodegroupOutput struct {
	_ struct{} `type:"structure"`

	// The full description of your node group.
	Nodegroup *Nodegroup `locationName:"nodegroup" type:"structure"`
}

// String returns the string representation
func (s DescribeNodegroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeNodegroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Nodegroup != nil {
		v := s.Nodegroup

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "nodegroup", v, metadata)
	}
	return nil
}

const opDescribeNodegroup = "DescribeNodegroup"

// DescribeNodegroupRequest returns a request value for making API operation for
// Amazon Elastic Kubernetes Service.
//
// Returns descriptive information about an Amazon EKS node group.
//
//    // Example sending a request using DescribeNodegroupRequest.
//    req := client.DescribeNodegroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/DescribeNodegroup
func (c *Client) DescribeNodegroupRequest(input *DescribeNodegroupInput) DescribeNodegroupRequest {
	op := &aws.Operation{
		Name:       opDescribeNodegroup,
		HTTPMethod: "GET",
		HTTPPath:   "/clusters/{name}/node-groups/{nodegroupName}",
	}

	if input == nil {
		input = &DescribeNodegroupInput{}
	}

	req := c.newRequest(op, input, &DescribeNodegroupOutput{})

	return DescribeNodegroupRequest{Request: req, Input: input, Copy: c.DescribeNodegroupRequest}
}

// DescribeNodegroupRequest is the request type for the
// DescribeNodegroup API operation.
type DescribeNodegroupRequest struct {
	*aws.Request
	Input *DescribeNodegroupInput
	Copy  func(*DescribeNodegroupInput) DescribeNodegroupRequest
}

// Send marshals and sends the DescribeNodegroup API request.
func (r DescribeNodegroupRequest) Send(ctx context.Context) (*DescribeNodegroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeNodegroupResponse{
		DescribeNodegroupOutput: r.Request.Data.(*DescribeNodegroupOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeNodegroupResponse is the response type for the
// DescribeNodegroup API operation.
type DescribeNodegroupResponse struct {
	*DescribeNodegroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeNodegroup request.
func (r *DescribeNodegroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
