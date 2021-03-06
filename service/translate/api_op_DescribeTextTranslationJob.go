// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package translate

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeTextTranslationJobInput struct {
	_ struct{} `type:"structure"`

	// The identifier that Amazon Translate generated for the job. The StartTextTranslationJob
	// operation returns this identifier in its response.
	//
	// JobId is a required field
	JobId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeTextTranslationJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTextTranslationJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTextTranslationJobInput"}

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

type DescribeTextTranslationJobOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains the properties associated with an asynchronous batch
	// translation job.
	TextTranslationJobProperties *TextTranslationJobProperties `type:"structure"`
}

// String returns the string representation
func (s DescribeTextTranslationJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTextTranslationJob = "DescribeTextTranslationJob"

// DescribeTextTranslationJobRequest returns a request value for making API operation for
// Amazon Translate.
//
// Gets the properties associated with an asycnhronous batch translation job
// including name, ID, status, source and target languages, input/output S3
// buckets, and so on.
//
//    // Example sending a request using DescribeTextTranslationJobRequest.
//    req := client.DescribeTextTranslationJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/translate-2017-07-01/DescribeTextTranslationJob
func (c *Client) DescribeTextTranslationJobRequest(input *DescribeTextTranslationJobInput) DescribeTextTranslationJobRequest {
	op := &aws.Operation{
		Name:       opDescribeTextTranslationJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTextTranslationJobInput{}
	}

	req := c.newRequest(op, input, &DescribeTextTranslationJobOutput{})

	return DescribeTextTranslationJobRequest{Request: req, Input: input, Copy: c.DescribeTextTranslationJobRequest}
}

// DescribeTextTranslationJobRequest is the request type for the
// DescribeTextTranslationJob API operation.
type DescribeTextTranslationJobRequest struct {
	*aws.Request
	Input *DescribeTextTranslationJobInput
	Copy  func(*DescribeTextTranslationJobInput) DescribeTextTranslationJobRequest
}

// Send marshals and sends the DescribeTextTranslationJob API request.
func (r DescribeTextTranslationJobRequest) Send(ctx context.Context) (*DescribeTextTranslationJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTextTranslationJobResponse{
		DescribeTextTranslationJobOutput: r.Request.Data.(*DescribeTextTranslationJobOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTextTranslationJobResponse is the response type for the
// DescribeTextTranslationJob API operation.
type DescribeTextTranslationJobResponse struct {
	*DescribeTextTranslationJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTextTranslationJob request.
func (r *DescribeTextTranslationJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
