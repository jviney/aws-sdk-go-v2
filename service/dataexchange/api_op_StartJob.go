// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dataexchange

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type StartJobInput struct {
	_ struct{} `type:"structure"`

	// JobId is a required field
	JobId *string `location:"uri" locationName:"JobId" type:"string" required:"true"`
}

// String returns the string representation
func (s StartJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "JobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type StartJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opStartJob = "StartJob"

// StartJobRequest returns a request value for making API operation for
// AWS Data Exchange.
//
// This operation starts a job.
//
//    // Example sending a request using StartJobRequest.
//    req := client.StartJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dataexchange-2017-07-25/StartJob
func (c *Client) StartJobRequest(input *StartJobInput) StartJobRequest {
	op := &aws.Operation{
		Name:       opStartJob,
		HTTPMethod: "PATCH",
		HTTPPath:   "/v1/jobs/{JobId}",
	}

	if input == nil {
		input = &StartJobInput{}
	}

	req := c.newRequest(op, input, &StartJobOutput{})

	return StartJobRequest{Request: req, Input: input, Copy: c.StartJobRequest}
}

// StartJobRequest is the request type for the
// StartJob API operation.
type StartJobRequest struct {
	*aws.Request
	Input *StartJobInput
	Copy  func(*StartJobInput) StartJobRequest
}

// Send marshals and sends the StartJob API request.
func (r StartJobRequest) Send(ctx context.Context) (*StartJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartJobResponse{
		StartJobOutput: r.Request.Data.(*StartJobOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartJobResponse is the response type for the
// StartJob API operation.
type StartJobResponse struct {
	*StartJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartJob request.
func (r *StartJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
