// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteGatewayGroupInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the gateway group to delete.
	//
	// GatewayGroupArn is a required field
	GatewayGroupArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteGatewayGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteGatewayGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteGatewayGroupInput"}

	if s.GatewayGroupArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayGroupArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteGatewayGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteGatewayGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteGatewayGroup = "DeleteGatewayGroup"

// DeleteGatewayGroupRequest returns a request value for making API operation for
// Alexa For Business.
//
// Deletes a gateway group.
//
//    // Example sending a request using DeleteGatewayGroupRequest.
//    req := client.DeleteGatewayGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DeleteGatewayGroup
func (c *Client) DeleteGatewayGroupRequest(input *DeleteGatewayGroupInput) DeleteGatewayGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteGatewayGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteGatewayGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteGatewayGroupOutput{})

	return DeleteGatewayGroupRequest{Request: req, Input: input, Copy: c.DeleteGatewayGroupRequest}
}

// DeleteGatewayGroupRequest is the request type for the
// DeleteGatewayGroup API operation.
type DeleteGatewayGroupRequest struct {
	*aws.Request
	Input *DeleteGatewayGroupInput
	Copy  func(*DeleteGatewayGroupInput) DeleteGatewayGroupRequest
}

// Send marshals and sends the DeleteGatewayGroup API request.
func (r DeleteGatewayGroupRequest) Send(ctx context.Context) (*DeleteGatewayGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteGatewayGroupResponse{
		DeleteGatewayGroupOutput: r.Request.Data.(*DeleteGatewayGroupOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteGatewayGroupResponse is the response type for the
// DeleteGatewayGroup API operation.
type DeleteGatewayGroupResponse struct {
	*DeleteGatewayGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteGatewayGroup request.
func (r *DeleteGatewayGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
