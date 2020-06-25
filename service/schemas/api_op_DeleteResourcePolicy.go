// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteResourcePolicyInput struct {
	_ struct{} `type:"structure"`

	RegistryName *string `location:"querystring" locationName:"registryName" type:"string"`
}

// String returns the string representation
func (s DeleteResourcePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteResourcePolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RegistryName != nil {
		v := *s.RegistryName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "registryName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteResourcePolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteResourcePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteResourcePolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteResourcePolicy = "DeleteResourcePolicy"

// DeleteResourcePolicyRequest returns a request value for making API operation for
// Schemas.
//
// Delete the resource-based policy attached to the specified registry.
//
//    // Example sending a request using DeleteResourcePolicyRequest.
//    req := client.DeleteResourcePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/schemas-2019-12-02/DeleteResourcePolicy
func (c *Client) DeleteResourcePolicyRequest(input *DeleteResourcePolicyInput) DeleteResourcePolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteResourcePolicy,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/policy",
	}

	if input == nil {
		input = &DeleteResourcePolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteResourcePolicyOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteResourcePolicyRequest{Request: req, Input: input, Copy: c.DeleteResourcePolicyRequest}
}

// DeleteResourcePolicyRequest is the request type for the
// DeleteResourcePolicy API operation.
type DeleteResourcePolicyRequest struct {
	*aws.Request
	Input *DeleteResourcePolicyInput
	Copy  func(*DeleteResourcePolicyInput) DeleteResourcePolicyRequest
}

// Send marshals and sends the DeleteResourcePolicy API request.
func (r DeleteResourcePolicyRequest) Send(ctx context.Context) (*DeleteResourcePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteResourcePolicyResponse{
		DeleteResourcePolicyOutput: r.Request.Data.(*DeleteResourcePolicyOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteResourcePolicyResponse is the response type for the
// DeleteResourcePolicy API operation.
type DeleteResourcePolicyResponse struct {
	*DeleteResourcePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteResourcePolicy request.
func (r *DeleteResourcePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
