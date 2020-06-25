// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type UpdateEndpointInput struct {
	_ struct{} `type:"structure"`

	// The desired number of inference units to be used by the model using this
	// endpoint. Each inference unit represents of a throughput of 100 characters
	// per second.
	//
	// DesiredInferenceUnits is a required field
	DesiredInferenceUnits *int64 `min:"1" type:"integer" required:"true"`

	// The Amazon Resource Number (ARN) of the endpoint being updated.
	//
	// EndpointArn is a required field
	EndpointArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateEndpointInput"}

	if s.DesiredInferenceUnits == nil {
		invalidParams.Add(aws.NewErrParamRequired("DesiredInferenceUnits"))
	}
	if s.DesiredInferenceUnits != nil && *s.DesiredInferenceUnits < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("DesiredInferenceUnits", 1))
	}

	if s.EndpointArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateEndpointOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateEndpoint = "UpdateEndpoint"

// UpdateEndpointRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Updates information about the specified endpoint.
//
//    // Example sending a request using UpdateEndpointRequest.
//    req := client.UpdateEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/UpdateEndpoint
func (c *Client) UpdateEndpointRequest(input *UpdateEndpointInput) UpdateEndpointRequest {
	op := &aws.Operation{
		Name:       opUpdateEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateEndpointInput{}
	}

	req := c.newRequest(op, input, &UpdateEndpointOutput{})

	return UpdateEndpointRequest{Request: req, Input: input, Copy: c.UpdateEndpointRequest}
}

// UpdateEndpointRequest is the request type for the
// UpdateEndpoint API operation.
type UpdateEndpointRequest struct {
	*aws.Request
	Input *UpdateEndpointInput
	Copy  func(*UpdateEndpointInput) UpdateEndpointRequest
}

// Send marshals and sends the UpdateEndpoint API request.
func (r UpdateEndpointRequest) Send(ctx context.Context) (*UpdateEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateEndpointResponse{
		UpdateEndpointOutput: r.Request.Data.(*UpdateEndpointOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateEndpointResponse is the response type for the
// UpdateEndpoint API operation.
type UpdateEndpointResponse struct {
	*UpdateEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateEndpoint request.
func (r *UpdateEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
