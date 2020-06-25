// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DeleteMitigationActionInput struct {
	_ struct{} `type:"structure"`

	// The name of the mitigation action that you want to delete.
	//
	// ActionName is a required field
	ActionName *string `location:"uri" locationName:"actionName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMitigationActionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMitigationActionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteMitigationActionInput"}

	if s.ActionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActionName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteMitigationActionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ActionName != nil {
		v := *s.ActionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "actionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteMitigationActionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteMitigationActionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteMitigationActionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteMitigationAction = "DeleteMitigationAction"

// DeleteMitigationActionRequest returns a request value for making API operation for
// AWS IoT.
//
// Deletes a defined mitigation action from your AWS account.
//
//    // Example sending a request using DeleteMitigationActionRequest.
//    req := client.DeleteMitigationActionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteMitigationActionRequest(input *DeleteMitigationActionInput) DeleteMitigationActionRequest {
	op := &aws.Operation{
		Name:       opDeleteMitigationAction,
		HTTPMethod: "DELETE",
		HTTPPath:   "/mitigationactions/actions/{actionName}",
	}

	if input == nil {
		input = &DeleteMitigationActionInput{}
	}

	req := c.newRequest(op, input, &DeleteMitigationActionOutput{})

	return DeleteMitigationActionRequest{Request: req, Input: input, Copy: c.DeleteMitigationActionRequest}
}

// DeleteMitigationActionRequest is the request type for the
// DeleteMitigationAction API operation.
type DeleteMitigationActionRequest struct {
	*aws.Request
	Input *DeleteMitigationActionInput
	Copy  func(*DeleteMitigationActionInput) DeleteMitigationActionRequest
}

// Send marshals and sends the DeleteMitigationAction API request.
func (r DeleteMitigationActionRequest) Send(ctx context.Context) (*DeleteMitigationActionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteMitigationActionResponse{
		DeleteMitigationActionOutput: r.Request.Data.(*DeleteMitigationActionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteMitigationActionResponse is the response type for the
// DeleteMitigationAction API operation.
type DeleteMitigationActionResponse struct {
	*DeleteMitigationActionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteMitigationAction request.
func (r *DeleteMitigationActionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
