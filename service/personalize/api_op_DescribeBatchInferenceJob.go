// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeBatchInferenceJobInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the batch inference job to describe.
	//
	// BatchInferenceJobArn is a required field
	BatchInferenceJobArn *string `locationName:"batchInferenceJobArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeBatchInferenceJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBatchInferenceJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeBatchInferenceJobInput"}

	if s.BatchInferenceJobArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("BatchInferenceJobArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeBatchInferenceJobOutput struct {
	_ struct{} `type:"structure"`

	// Information on the specified batch inference job.
	BatchInferenceJob *BatchInferenceJob `locationName:"batchInferenceJob" type:"structure"`
}

// String returns the string representation
func (s DescribeBatchInferenceJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeBatchInferenceJob = "DescribeBatchInferenceJob"

// DescribeBatchInferenceJobRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Gets the properties of a batch inference job including name, Amazon Resource
// Name (ARN), status, input and output configurations, and the ARN of the solution
// version used to generate the recommendations.
//
//    // Example sending a request using DescribeBatchInferenceJobRequest.
//    req := client.DescribeBatchInferenceJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/DescribeBatchInferenceJob
func (c *Client) DescribeBatchInferenceJobRequest(input *DescribeBatchInferenceJobInput) DescribeBatchInferenceJobRequest {
	op := &aws.Operation{
		Name:       opDescribeBatchInferenceJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeBatchInferenceJobInput{}
	}

	req := c.newRequest(op, input, &DescribeBatchInferenceJobOutput{})

	return DescribeBatchInferenceJobRequest{Request: req, Input: input, Copy: c.DescribeBatchInferenceJobRequest}
}

// DescribeBatchInferenceJobRequest is the request type for the
// DescribeBatchInferenceJob API operation.
type DescribeBatchInferenceJobRequest struct {
	*aws.Request
	Input *DescribeBatchInferenceJobInput
	Copy  func(*DescribeBatchInferenceJobInput) DescribeBatchInferenceJobRequest
}

// Send marshals and sends the DescribeBatchInferenceJob API request.
func (r DescribeBatchInferenceJobRequest) Send(ctx context.Context) (*DescribeBatchInferenceJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeBatchInferenceJobResponse{
		DescribeBatchInferenceJobOutput: r.Request.Data.(*DescribeBatchInferenceJobOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeBatchInferenceJobResponse is the response type for the
// DescribeBatchInferenceJob API operation.
type DescribeBatchInferenceJobResponse struct {
	*DescribeBatchInferenceJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeBatchInferenceJob request.
func (r *DescribeBatchInferenceJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
