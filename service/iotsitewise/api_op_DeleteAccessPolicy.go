// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DeleteAccessPolicyInput struct {
	_ struct{} `type:"structure"`

	// The ID of the access policy to be deleted.
	//
	// AccessPolicyId is a required field
	AccessPolicyId *string `location:"uri" locationName:"accessPolicyId" min:"36" type:"string" required:"true"`

	// A unique case-sensitive identifier that you can provide to ensure the idempotency
	// of the request. Don't reuse this client token if a new idempotent request
	// is required.
	ClientToken *string `location:"querystring" locationName:"clientToken" min:"36" type:"string" idempotencyToken:"true"`
}

// String returns the string representation
func (s DeleteAccessPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAccessPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAccessPolicyInput"}

	if s.AccessPolicyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessPolicyId"))
	}
	if s.AccessPolicyId != nil && len(*s.AccessPolicyId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("AccessPolicyId", 36))
	}
	if s.ClientToken != nil && len(*s.ClientToken) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAccessPolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccessPolicyId != nil {
		v := *s.AccessPolicyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accessPolicyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
		e.SetValue(protocol.QueryTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteAccessPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAccessPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAccessPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteAccessPolicy = "DeleteAccessPolicy"

// DeleteAccessPolicyRequest returns a request value for making API operation for
// AWS IoT SiteWise.
//
// Deletes an access policy that grants the specified AWS Single Sign-On identity
// access to the specified AWS IoT SiteWise Monitor resource. You can use this
// operation to revoke access to an AWS IoT SiteWise Monitor resource.
//
//    // Example sending a request using DeleteAccessPolicyRequest.
//    req := client.DeleteAccessPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotsitewise-2019-12-02/DeleteAccessPolicy
func (c *Client) DeleteAccessPolicyRequest(input *DeleteAccessPolicyInput) DeleteAccessPolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteAccessPolicy,
		HTTPMethod: "DELETE",
		HTTPPath:   "/access-policies/{accessPolicyId}",
	}

	if input == nil {
		input = &DeleteAccessPolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteAccessPolicyOutput{})
	req.Handlers.Build.PushBackNamed(protocol.NewHostPrefixHandler("monitor.", nil))
	req.Handlers.Build.PushBackNamed(protocol.ValidateEndpointHostHandler)

	return DeleteAccessPolicyRequest{Request: req, Input: input, Copy: c.DeleteAccessPolicyRequest}
}

// DeleteAccessPolicyRequest is the request type for the
// DeleteAccessPolicy API operation.
type DeleteAccessPolicyRequest struct {
	*aws.Request
	Input *DeleteAccessPolicyInput
	Copy  func(*DeleteAccessPolicyInput) DeleteAccessPolicyRequest
}

// Send marshals and sends the DeleteAccessPolicy API request.
func (r DeleteAccessPolicyRequest) Send(ctx context.Context) (*DeleteAccessPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAccessPolicyResponse{
		DeleteAccessPolicyOutput: r.Request.Data.(*DeleteAccessPolicyOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAccessPolicyResponse is the response type for the
// DeleteAccessPolicy API operation.
type DeleteAccessPolicyResponse struct {
	*DeleteAccessPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAccessPolicy request.
func (r *DeleteAccessPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
