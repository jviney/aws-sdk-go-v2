// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteEndpointGroupInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the endpoint group to delete.
	//
	// EndpointGroupArn is a required field
	EndpointGroupArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteEndpointGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteEndpointGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteEndpointGroupInput"}

	if s.EndpointGroupArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointGroupArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteEndpointGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteEndpointGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteEndpointGroup = "DeleteEndpointGroup"

// DeleteEndpointGroupRequest returns a request value for making API operation for
// AWS Global Accelerator.
//
// Delete an endpoint group from a listener.
//
//    // Example sending a request using DeleteEndpointGroupRequest.
//    req := client.DeleteEndpointGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/DeleteEndpointGroup
func (c *Client) DeleteEndpointGroupRequest(input *DeleteEndpointGroupInput) DeleteEndpointGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteEndpointGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteEndpointGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteEndpointGroupOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteEndpointGroupRequest{Request: req, Input: input, Copy: c.DeleteEndpointGroupRequest}
}

// DeleteEndpointGroupRequest is the request type for the
// DeleteEndpointGroup API operation.
type DeleteEndpointGroupRequest struct {
	*aws.Request
	Input *DeleteEndpointGroupInput
	Copy  func(*DeleteEndpointGroupInput) DeleteEndpointGroupRequest
}

// Send marshals and sends the DeleteEndpointGroup API request.
func (r DeleteEndpointGroupRequest) Send(ctx context.Context) (*DeleteEndpointGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteEndpointGroupResponse{
		DeleteEndpointGroupOutput: r.Request.Data.(*DeleteEndpointGroupOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteEndpointGroupResponse is the response type for the
// DeleteEndpointGroup API operation.
type DeleteEndpointGroupResponse struct {
	*DeleteEndpointGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteEndpointGroup request.
func (r *DeleteEndpointGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
