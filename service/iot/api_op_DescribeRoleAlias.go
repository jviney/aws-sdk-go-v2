// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DescribeRoleAliasInput struct {
	_ struct{} `type:"structure"`

	// The role alias to describe.
	//
	// RoleAlias is a required field
	RoleAlias *string `location:"uri" locationName:"roleAlias" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeRoleAliasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRoleAliasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeRoleAliasInput"}

	if s.RoleAlias == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleAlias"))
	}
	if s.RoleAlias != nil && len(*s.RoleAlias) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleAlias", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeRoleAliasInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RoleAlias != nil {
		v := *s.RoleAlias

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "roleAlias", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeRoleAliasOutput struct {
	_ struct{} `type:"structure"`

	// The role alias description.
	RoleAliasDescription *RoleAliasDescription `locationName:"roleAliasDescription" type:"structure"`
}

// String returns the string representation
func (s DescribeRoleAliasOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeRoleAliasOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.RoleAliasDescription != nil {
		v := s.RoleAliasDescription

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "roleAliasDescription", v, metadata)
	}
	return nil
}

const opDescribeRoleAlias = "DescribeRoleAlias"

// DescribeRoleAliasRequest returns a request value for making API operation for
// AWS IoT.
//
// Describes a role alias.
//
//    // Example sending a request using DescribeRoleAliasRequest.
//    req := client.DescribeRoleAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeRoleAliasRequest(input *DescribeRoleAliasInput) DescribeRoleAliasRequest {
	op := &aws.Operation{
		Name:       opDescribeRoleAlias,
		HTTPMethod: "GET",
		HTTPPath:   "/role-aliases/{roleAlias}",
	}

	if input == nil {
		input = &DescribeRoleAliasInput{}
	}

	req := c.newRequest(op, input, &DescribeRoleAliasOutput{})

	return DescribeRoleAliasRequest{Request: req, Input: input, Copy: c.DescribeRoleAliasRequest}
}

// DescribeRoleAliasRequest is the request type for the
// DescribeRoleAlias API operation.
type DescribeRoleAliasRequest struct {
	*aws.Request
	Input *DescribeRoleAliasInput
	Copy  func(*DescribeRoleAliasInput) DescribeRoleAliasRequest
}

// Send marshals and sends the DescribeRoleAlias API request.
func (r DescribeRoleAliasRequest) Send(ctx context.Context) (*DescribeRoleAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRoleAliasResponse{
		DescribeRoleAliasOutput: r.Request.Data.(*DescribeRoleAliasOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeRoleAliasResponse is the response type for the
// DescribeRoleAlias API operation.
type DescribeRoleAliasResponse struct {
	*DescribeRoleAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRoleAlias request.
func (r *DescribeRoleAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
