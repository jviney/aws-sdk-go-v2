// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteActivationInput struct {
	_ struct{} `type:"structure"`

	// The ID of the activation that you want to delete.
	//
	// ActivationId is a required field
	ActivationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteActivationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteActivationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteActivationInput"}

	if s.ActivationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActivationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteActivationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteActivationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteActivation = "DeleteActivation"

// DeleteActivationRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Deletes an activation. You are not required to delete an activation. If you
// delete an activation, you can no longer use it to register additional managed
// instances. Deleting an activation does not de-register managed instances.
// You must manually de-register managed instances.
//
//    // Example sending a request using DeleteActivationRequest.
//    req := client.DeleteActivationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeleteActivation
func (c *Client) DeleteActivationRequest(input *DeleteActivationInput) DeleteActivationRequest {
	op := &aws.Operation{
		Name:       opDeleteActivation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteActivationInput{}
	}

	req := c.newRequest(op, input, &DeleteActivationOutput{})

	return DeleteActivationRequest{Request: req, Input: input, Copy: c.DeleteActivationRequest}
}

// DeleteActivationRequest is the request type for the
// DeleteActivation API operation.
type DeleteActivationRequest struct {
	*aws.Request
	Input *DeleteActivationInput
	Copy  func(*DeleteActivationInput) DeleteActivationRequest
}

// Send marshals and sends the DeleteActivation API request.
func (r DeleteActivationRequest) Send(ctx context.Context) (*DeleteActivationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteActivationResponse{
		DeleteActivationOutput: r.Request.Data.(*DeleteActivationOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteActivationResponse is the response type for the
// DeleteActivation API operation.
type DeleteActivationResponse struct {
	*DeleteActivationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteActivation request.
func (r *DeleteActivationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
