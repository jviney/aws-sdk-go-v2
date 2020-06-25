// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/restjson"
)

type AttachPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name of the policy to attach.
	//
	// PolicyName is a required field
	PolicyName *string `location:"uri" locationName:"policyName" min:"1" type:"string" required:"true"`

	// The identity (https://docs.aws.amazon.com/iot/latest/developerguide/iot-security-identity.html)
	// to which the policy is attached.
	//
	// Target is a required field
	Target *string `locationName:"target" type:"string" required:"true"`
}

// String returns the string representation
func (s AttachPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AttachPolicyInput"}

	if s.PolicyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyName"))
	}
	if s.PolicyName != nil && len(*s.PolicyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyName", 1))
	}

	if s.Target == nil {
		invalidParams.Add(aws.NewErrParamRequired("Target"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AttachPolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Target != nil {
		v := *s.Target

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "target", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PolicyName != nil {
		v := *s.PolicyName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "policyName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type AttachPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AttachPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AttachPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opAttachPolicy = "AttachPolicy"

// AttachPolicyRequest returns a request value for making API operation for
// AWS IoT.
//
// Attaches a policy to the specified target.
//
//    // Example sending a request using AttachPolicyRequest.
//    req := client.AttachPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) AttachPolicyRequest(input *AttachPolicyInput) AttachPolicyRequest {
	op := &aws.Operation{
		Name:       opAttachPolicy,
		HTTPMethod: "PUT",
		HTTPPath:   "/target-policies/{policyName}",
	}

	if input == nil {
		input = &AttachPolicyInput{}
	}

	req := c.newRequest(op, input, &AttachPolicyOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return AttachPolicyRequest{Request: req, Input: input, Copy: c.AttachPolicyRequest}
}

// AttachPolicyRequest is the request type for the
// AttachPolicy API operation.
type AttachPolicyRequest struct {
	*aws.Request
	Input *AttachPolicyInput
	Copy  func(*AttachPolicyInput) AttachPolicyRequest
}

// Send marshals and sends the AttachPolicy API request.
func (r AttachPolicyRequest) Send(ctx context.Context) (*AttachPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachPolicyResponse{
		AttachPolicyOutput: r.Request.Data.(*AttachPolicyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachPolicyResponse is the response type for the
// AttachPolicy API operation.
type AttachPolicyResponse struct {
	*AttachPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachPolicy request.
func (r *AttachPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
