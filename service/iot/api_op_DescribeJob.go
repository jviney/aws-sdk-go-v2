// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DescribeJobInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier you assigned to this job when it was created.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeJobOutput struct {
	_ struct{} `type:"structure"`

	// An S3 link to the job document.
	DocumentSource *string `locationName:"documentSource" min:"1" type:"string"`

	// Information about the job.
	Job *Job `locationName:"job" type:"structure"`
}

// String returns the string representation
func (s DescribeJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DocumentSource != nil {
		v := *s.DocumentSource

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "documentSource", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Job != nil {
		v := s.Job

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "job", v, metadata)
	}
	return nil
}

const opDescribeJob = "DescribeJob"

// DescribeJobRequest returns a request value for making API operation for
// AWS IoT.
//
// Describes a job.
//
//    // Example sending a request using DescribeJobRequest.
//    req := client.DescribeJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeJobRequest(input *DescribeJobInput) DescribeJobRequest {
	op := &aws.Operation{
		Name:       opDescribeJob,
		HTTPMethod: "GET",
		HTTPPath:   "/jobs/{jobId}",
	}

	if input == nil {
		input = &DescribeJobInput{}
	}

	req := c.newRequest(op, input, &DescribeJobOutput{})

	return DescribeJobRequest{Request: req, Input: input, Copy: c.DescribeJobRequest}
}

// DescribeJobRequest is the request type for the
// DescribeJob API operation.
type DescribeJobRequest struct {
	*aws.Request
	Input *DescribeJobInput
	Copy  func(*DescribeJobInput) DescribeJobRequest
}

// Send marshals and sends the DescribeJob API request.
func (r DescribeJobRequest) Send(ctx context.Context) (*DescribeJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeJobResponse{
		DescribeJobOutput: r.Request.Data.(*DescribeJobOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeJobResponse is the response type for the
// DescribeJob API operation.
type DescribeJobResponse struct {
	*DescribeJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeJob request.
func (r *DescribeJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
