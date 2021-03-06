// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type UpdateHITTypeOfHITInput struct {
	_ struct{} `type:"structure"`

	// The HIT to update.
	//
	// HITId is a required field
	HITId *string `min:"1" type:"string" required:"true"`

	// The ID of the new HIT type.
	//
	// HITTypeId is a required field
	HITTypeId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateHITTypeOfHITInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateHITTypeOfHITInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateHITTypeOfHITInput"}

	if s.HITId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HITId"))
	}
	if s.HITId != nil && len(*s.HITId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HITId", 1))
	}

	if s.HITTypeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HITTypeId"))
	}
	if s.HITTypeId != nil && len(*s.HITTypeId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HITTypeId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateHITTypeOfHITOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateHITTypeOfHITOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateHITTypeOfHIT = "UpdateHITTypeOfHIT"

// UpdateHITTypeOfHITRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The UpdateHITTypeOfHIT operation allows you to change the HITType properties
// of a HIT. This operation disassociates the HIT from its old HITType properties
// and associates it with the new HITType properties. The HIT takes on the properties
// of the new HITType in place of the old ones.
//
//    // Example sending a request using UpdateHITTypeOfHITRequest.
//    req := client.UpdateHITTypeOfHITRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/UpdateHITTypeOfHIT
func (c *Client) UpdateHITTypeOfHITRequest(input *UpdateHITTypeOfHITInput) UpdateHITTypeOfHITRequest {
	op := &aws.Operation{
		Name:       opUpdateHITTypeOfHIT,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateHITTypeOfHITInput{}
	}

	req := c.newRequest(op, input, &UpdateHITTypeOfHITOutput{})

	return UpdateHITTypeOfHITRequest{Request: req, Input: input, Copy: c.UpdateHITTypeOfHITRequest}
}

// UpdateHITTypeOfHITRequest is the request type for the
// UpdateHITTypeOfHIT API operation.
type UpdateHITTypeOfHITRequest struct {
	*aws.Request
	Input *UpdateHITTypeOfHITInput
	Copy  func(*UpdateHITTypeOfHITInput) UpdateHITTypeOfHITRequest
}

// Send marshals and sends the UpdateHITTypeOfHIT API request.
func (r UpdateHITTypeOfHITRequest) Send(ctx context.Context) (*UpdateHITTypeOfHITResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateHITTypeOfHITResponse{
		UpdateHITTypeOfHITOutput: r.Request.Data.(*UpdateHITTypeOfHITOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateHITTypeOfHITResponse is the response type for the
// UpdateHITTypeOfHIT API operation.
type UpdateHITTypeOfHITResponse struct {
	*UpdateHITTypeOfHITOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateHITTypeOfHIT request.
func (r *UpdateHITTypeOfHITResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
