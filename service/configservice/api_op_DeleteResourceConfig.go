// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteResourceConfigInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier of the resource.
	//
	// ResourceId is a required field
	ResourceId *string `min:"1" type:"string" required:"true"`

	// The type of the resource.
	//
	// ResourceType is a required field
	ResourceType *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteResourceConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteResourceConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteResourceConfigInput"}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if s.ResourceId != nil && len(*s.ResourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceId", 1))
	}

	if s.ResourceType == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceType"))
	}
	if s.ResourceType != nil && len(*s.ResourceType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceType", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteResourceConfigOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteResourceConfigOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteResourceConfig = "DeleteResourceConfig"

// DeleteResourceConfigRequest returns a request value for making API operation for
// AWS Config.
//
// Records the configuration state for a custom resource that has been deleted.
// This API records a new ConfigurationItem with a ResourceDeleted status. You
// can retrieve the ConfigurationItems recorded for this resource in your AWS
// Config History.
//
//    // Example sending a request using DeleteResourceConfigRequest.
//    req := client.DeleteResourceConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DeleteResourceConfig
func (c *Client) DeleteResourceConfigRequest(input *DeleteResourceConfigInput) DeleteResourceConfigRequest {
	op := &aws.Operation{
		Name:       opDeleteResourceConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteResourceConfigInput{}
	}

	req := c.newRequest(op, input, &DeleteResourceConfigOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteResourceConfigRequest{Request: req, Input: input, Copy: c.DeleteResourceConfigRequest}
}

// DeleteResourceConfigRequest is the request type for the
// DeleteResourceConfig API operation.
type DeleteResourceConfigRequest struct {
	*aws.Request
	Input *DeleteResourceConfigInput
	Copy  func(*DeleteResourceConfigInput) DeleteResourceConfigRequest
}

// Send marshals and sends the DeleteResourceConfig API request.
func (r DeleteResourceConfigRequest) Send(ctx context.Context) (*DeleteResourceConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteResourceConfigResponse{
		DeleteResourceConfigOutput: r.Request.Data.(*DeleteResourceConfigOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteResourceConfigResponse is the response type for the
// DeleteResourceConfig API operation.
type DeleteResourceConfigResponse struct {
	*DeleteResourceConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteResourceConfig request.
func (r *DeleteResourceConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
