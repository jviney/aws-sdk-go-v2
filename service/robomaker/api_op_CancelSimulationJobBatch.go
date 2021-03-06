// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type CancelSimulationJobBatchInput struct {
	_ struct{} `type:"structure"`

	// The id of the batch to cancel.
	//
	// Batch is a required field
	Batch *string `locationName:"batch" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CancelSimulationJobBatchInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelSimulationJobBatchInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelSimulationJobBatchInput"}

	if s.Batch == nil {
		invalidParams.Add(aws.NewErrParamRequired("Batch"))
	}
	if s.Batch != nil && len(*s.Batch) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Batch", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelSimulationJobBatchInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Batch != nil {
		v := *s.Batch

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "batch", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CancelSimulationJobBatchOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CancelSimulationJobBatchOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelSimulationJobBatchOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCancelSimulationJobBatch = "CancelSimulationJobBatch"

// CancelSimulationJobBatchRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Cancels a simulation job batch. When you cancel a simulation job batch, you
// are also cancelling all of the active simulation jobs created as part of
// the batch.
//
//    // Example sending a request using CancelSimulationJobBatchRequest.
//    req := client.CancelSimulationJobBatchRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CancelSimulationJobBatch
func (c *Client) CancelSimulationJobBatchRequest(input *CancelSimulationJobBatchInput) CancelSimulationJobBatchRequest {
	op := &aws.Operation{
		Name:       opCancelSimulationJobBatch,
		HTTPMethod: "POST",
		HTTPPath:   "/cancelSimulationJobBatch",
	}

	if input == nil {
		input = &CancelSimulationJobBatchInput{}
	}

	req := c.newRequest(op, input, &CancelSimulationJobBatchOutput{})

	return CancelSimulationJobBatchRequest{Request: req, Input: input, Copy: c.CancelSimulationJobBatchRequest}
}

// CancelSimulationJobBatchRequest is the request type for the
// CancelSimulationJobBatch API operation.
type CancelSimulationJobBatchRequest struct {
	*aws.Request
	Input *CancelSimulationJobBatchInput
	Copy  func(*CancelSimulationJobBatchInput) CancelSimulationJobBatchRequest
}

// Send marshals and sends the CancelSimulationJobBatch API request.
func (r CancelSimulationJobBatchRequest) Send(ctx context.Context) (*CancelSimulationJobBatchResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelSimulationJobBatchResponse{
		CancelSimulationJobBatchOutput: r.Request.Data.(*CancelSimulationJobBatchOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelSimulationJobBatchResponse is the response type for the
// CancelSimulationJobBatch API operation.
type CancelSimulationJobBatchResponse struct {
	*CancelSimulationJobBatchOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelSimulationJobBatch request.
func (r *CancelSimulationJobBatchResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
