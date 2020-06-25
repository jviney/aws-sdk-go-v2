// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type UpdateAccessPolicyInput struct {
	_ struct{} `type:"structure"`

	// The ID of the access policy.
	//
	// AccessPolicyId is a required field
	AccessPolicyId *string `location:"uri" locationName:"accessPolicyId" min:"36" type:"string" required:"true"`

	// The identity for this access policy. Choose either a user or a group but
	// not both.
	//
	// AccessPolicyIdentity is a required field
	AccessPolicyIdentity *Identity `locationName:"accessPolicyIdentity" type:"structure" required:"true"`

	// The permission level for this access policy. Note that a project ADMINISTRATOR
	// is also known as a project owner.
	//
	// AccessPolicyPermission is a required field
	AccessPolicyPermission Permission `locationName:"accessPolicyPermission" type:"string" required:"true" enum:"true"`

	// The AWS IoT SiteWise Monitor resource for this access policy. Choose either
	// portal or project but not both.
	//
	// AccessPolicyResource is a required field
	AccessPolicyResource *Resource `locationName:"accessPolicyResource" type:"structure" required:"true"`

	// A unique case-sensitive identifier that you can provide to ensure the idempotency
	// of the request. Don't reuse this client token if a new idempotent request
	// is required.
	ClientToken *string `locationName:"clientToken" min:"36" type:"string" idempotencyToken:"true"`
}

// String returns the string representation
func (s UpdateAccessPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAccessPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAccessPolicyInput"}

	if s.AccessPolicyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessPolicyId"))
	}
	if s.AccessPolicyId != nil && len(*s.AccessPolicyId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("AccessPolicyId", 36))
	}

	if s.AccessPolicyIdentity == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessPolicyIdentity"))
	}
	if len(s.AccessPolicyPermission) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("AccessPolicyPermission"))
	}

	if s.AccessPolicyResource == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessPolicyResource"))
	}
	if s.ClientToken != nil && len(*s.ClientToken) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 36))
	}
	if s.AccessPolicyIdentity != nil {
		if err := s.AccessPolicyIdentity.Validate(); err != nil {
			invalidParams.AddNested("AccessPolicyIdentity", err.(aws.ErrInvalidParams))
		}
	}
	if s.AccessPolicyResource != nil {
		if err := s.AccessPolicyResource.Validate(); err != nil {
			invalidParams.AddNested("AccessPolicyResource", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateAccessPolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccessPolicyIdentity != nil {
		v := s.AccessPolicyIdentity

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "accessPolicyIdentity", v, metadata)
	}
	if len(s.AccessPolicyPermission) > 0 {
		v := s.AccessPolicyPermission

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "accessPolicyPermission", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.AccessPolicyResource != nil {
		v := s.AccessPolicyResource

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "accessPolicyResource", v, metadata)
	}
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
	if s.AccessPolicyId != nil {
		v := *s.AccessPolicyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accessPolicyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateAccessPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateAccessPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateAccessPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateAccessPolicy = "UpdateAccessPolicy"

// UpdateAccessPolicyRequest returns a request value for making API operation for
// AWS IoT SiteWise.
//
// Updates an existing access policy that specifies an AWS SSO user or group's
// access to an AWS IoT SiteWise Monitor portal or project resource.
//
//    // Example sending a request using UpdateAccessPolicyRequest.
//    req := client.UpdateAccessPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotsitewise-2019-12-02/UpdateAccessPolicy
func (c *Client) UpdateAccessPolicyRequest(input *UpdateAccessPolicyInput) UpdateAccessPolicyRequest {
	op := &aws.Operation{
		Name:       opUpdateAccessPolicy,
		HTTPMethod: "PUT",
		HTTPPath:   "/access-policies/{accessPolicyId}",
	}

	if input == nil {
		input = &UpdateAccessPolicyInput{}
	}

	req := c.newRequest(op, input, &UpdateAccessPolicyOutput{})
	req.Handlers.Build.PushBackNamed(protocol.NewHostPrefixHandler("monitor.", nil))
	req.Handlers.Build.PushBackNamed(protocol.ValidateEndpointHostHandler)

	return UpdateAccessPolicyRequest{Request: req, Input: input, Copy: c.UpdateAccessPolicyRequest}
}

// UpdateAccessPolicyRequest is the request type for the
// UpdateAccessPolicy API operation.
type UpdateAccessPolicyRequest struct {
	*aws.Request
	Input *UpdateAccessPolicyInput
	Copy  func(*UpdateAccessPolicyInput) UpdateAccessPolicyRequest
}

// Send marshals and sends the UpdateAccessPolicy API request.
func (r UpdateAccessPolicyRequest) Send(ctx context.Context) (*UpdateAccessPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAccessPolicyResponse{
		UpdateAccessPolicyOutput: r.Request.Data.(*UpdateAccessPolicyOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAccessPolicyResponse is the response type for the
// UpdateAccessPolicy API operation.
type UpdateAccessPolicyResponse struct {
	*UpdateAccessPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAccessPolicy request.
func (r *UpdateAccessPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
