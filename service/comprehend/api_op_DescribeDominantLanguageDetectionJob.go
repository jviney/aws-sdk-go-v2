// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeDominantLanguageDetectionJobInput struct {
	_ struct{} `type:"structure"`

	// The identifier that Amazon Comprehend generated for the job. The operation
	// returns this identifier in its response.
	//
	// JobId is a required field
	JobId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDominantLanguageDetectionJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDominantLanguageDetectionJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDominantLanguageDetectionJobInput"}

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

type DescribeDominantLanguageDetectionJobOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains the properties associated with a dominant language
	// detection job.
	DominantLanguageDetectionJobProperties *DominantLanguageDetectionJobProperties `type:"structure"`
}

// String returns the string representation
func (s DescribeDominantLanguageDetectionJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDominantLanguageDetectionJob = "DescribeDominantLanguageDetectionJob"

// DescribeDominantLanguageDetectionJobRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Gets the properties associated with a dominant language detection job. Use
// this operation to get the status of a detection job.
//
//    // Example sending a request using DescribeDominantLanguageDetectionJobRequest.
//    req := client.DescribeDominantLanguageDetectionJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/DescribeDominantLanguageDetectionJob
func (c *Client) DescribeDominantLanguageDetectionJobRequest(input *DescribeDominantLanguageDetectionJobInput) DescribeDominantLanguageDetectionJobRequest {
	op := &aws.Operation{
		Name:       opDescribeDominantLanguageDetectionJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDominantLanguageDetectionJobInput{}
	}

	req := c.newRequest(op, input, &DescribeDominantLanguageDetectionJobOutput{})

	return DescribeDominantLanguageDetectionJobRequest{Request: req, Input: input, Copy: c.DescribeDominantLanguageDetectionJobRequest}
}

// DescribeDominantLanguageDetectionJobRequest is the request type for the
// DescribeDominantLanguageDetectionJob API operation.
type DescribeDominantLanguageDetectionJobRequest struct {
	*aws.Request
	Input *DescribeDominantLanguageDetectionJobInput
	Copy  func(*DescribeDominantLanguageDetectionJobInput) DescribeDominantLanguageDetectionJobRequest
}

// Send marshals and sends the DescribeDominantLanguageDetectionJob API request.
func (r DescribeDominantLanguageDetectionJobRequest) Send(ctx context.Context) (*DescribeDominantLanguageDetectionJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDominantLanguageDetectionJobResponse{
		DescribeDominantLanguageDetectionJobOutput: r.Request.Data.(*DescribeDominantLanguageDetectionJobOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDominantLanguageDetectionJobResponse is the response type for the
// DescribeDominantLanguageDetectionJob API operation.
type DescribeDominantLanguageDetectionJobResponse struct {
	*DescribeDominantLanguageDetectionJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDominantLanguageDetectionJob request.
func (r *DescribeDominantLanguageDetectionJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
