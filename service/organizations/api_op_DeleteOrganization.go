// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteOrganizationInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteOrganizationInput) String() string {
	return awsutil.Prettify(s)
}

type DeleteOrganizationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteOrganizationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteOrganization = "DeleteOrganization"

// DeleteOrganizationRequest returns a request value for making API operation for
// AWS Organizations.
//
// Deletes the organization. You can delete an organization only by using credentials
// from the master account. The organization must be empty of member accounts.
//
//    // Example sending a request using DeleteOrganizationRequest.
//    req := client.DeleteOrganizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/DeleteOrganization
func (c *Client) DeleteOrganizationRequest(input *DeleteOrganizationInput) DeleteOrganizationRequest {
	op := &aws.Operation{
		Name:       opDeleteOrganization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteOrganizationInput{}
	}

	req := c.newRequest(op, input, &DeleteOrganizationOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteOrganizationRequest{Request: req, Input: input, Copy: c.DeleteOrganizationRequest}
}

// DeleteOrganizationRequest is the request type for the
// DeleteOrganization API operation.
type DeleteOrganizationRequest struct {
	*aws.Request
	Input *DeleteOrganizationInput
	Copy  func(*DeleteOrganizationInput) DeleteOrganizationRequest
}

// Send marshals and sends the DeleteOrganization API request.
func (r DeleteOrganizationRequest) Send(ctx context.Context) (*DeleteOrganizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteOrganizationResponse{
		DeleteOrganizationOutput: r.Request.Data.(*DeleteOrganizationOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteOrganizationResponse is the response type for the
// DeleteOrganization API operation.
type DeleteOrganizationResponse struct {
	*DeleteOrganizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteOrganization request.
func (r *DeleteOrganizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
