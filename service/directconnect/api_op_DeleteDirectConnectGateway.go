// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteDirectConnectGatewayInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Direct Connect gateway.
	//
	// DirectConnectGatewayId is a required field
	DirectConnectGatewayId *string `locationName:"directConnectGatewayId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDirectConnectGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDirectConnectGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDirectConnectGatewayInput"}

	if s.DirectConnectGatewayId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectConnectGatewayId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDirectConnectGatewayOutput struct {
	_ struct{} `type:"structure"`

	// The Direct Connect gateway.
	DirectConnectGateway *DirectConnectGateway `locationName:"directConnectGateway" type:"structure"`
}

// String returns the string representation
func (s DeleteDirectConnectGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDirectConnectGateway = "DeleteDirectConnectGateway"

// DeleteDirectConnectGatewayRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Deletes the specified Direct Connect gateway. You must first delete all virtual
// interfaces that are attached to the Direct Connect gateway and disassociate
// all virtual private gateways associated with the Direct Connect gateway.
//
//    // Example sending a request using DeleteDirectConnectGatewayRequest.
//    req := client.DeleteDirectConnectGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/DeleteDirectConnectGateway
func (c *Client) DeleteDirectConnectGatewayRequest(input *DeleteDirectConnectGatewayInput) DeleteDirectConnectGatewayRequest {
	op := &aws.Operation{
		Name:       opDeleteDirectConnectGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDirectConnectGatewayInput{}
	}

	req := c.newRequest(op, input, &DeleteDirectConnectGatewayOutput{})

	return DeleteDirectConnectGatewayRequest{Request: req, Input: input, Copy: c.DeleteDirectConnectGatewayRequest}
}

// DeleteDirectConnectGatewayRequest is the request type for the
// DeleteDirectConnectGateway API operation.
type DeleteDirectConnectGatewayRequest struct {
	*aws.Request
	Input *DeleteDirectConnectGatewayInput
	Copy  func(*DeleteDirectConnectGatewayInput) DeleteDirectConnectGatewayRequest
}

// Send marshals and sends the DeleteDirectConnectGateway API request.
func (r DeleteDirectConnectGatewayRequest) Send(ctx context.Context) (*DeleteDirectConnectGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDirectConnectGatewayResponse{
		DeleteDirectConnectGatewayOutput: r.Request.Data.(*DeleteDirectConnectGatewayOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDirectConnectGatewayResponse is the response type for the
// DeleteDirectConnectGateway API operation.
type DeleteDirectConnectGatewayResponse struct {
	*DeleteDirectConnectGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDirectConnectGateway request.
func (r *DeleteDirectConnectGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
