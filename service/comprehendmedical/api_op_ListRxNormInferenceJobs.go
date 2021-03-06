// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehendmedical

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListRxNormInferenceJobsInput struct {
	_ struct{} `type:"structure"`

	// Filters the jobs that are returned. You can filter jobs based on their names,
	// status, or the date and time that they were submitted. You can only set one
	// filter at a time.
	Filter *ComprehendMedicalAsyncJobFilter `type:"structure"`

	// Identifies the next page of results to return.
	MaxResults *int64 `min:"1" type:"integer"`

	// Identifies the next page of results to return.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListRxNormInferenceJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRxNormInferenceJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRxNormInferenceJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			invalidParams.AddNested("Filter", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListRxNormInferenceJobsOutput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in each page. The default is 100.
	ComprehendMedicalAsyncJobPropertiesList []ComprehendMedicalAsyncJobProperties `type:"list"`

	// Identifies the next page of results to return.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListRxNormInferenceJobsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListRxNormInferenceJobs = "ListRxNormInferenceJobs"

// ListRxNormInferenceJobsRequest returns a request value for making API operation for
// AWS Comprehend Medical.
//
// Gets a list of InferRxNorm jobs that you have submitted.
//
//    // Example sending a request using ListRxNormInferenceJobsRequest.
//    req := client.ListRxNormInferenceJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehendmedical-2018-10-30/ListRxNormInferenceJobs
func (c *Client) ListRxNormInferenceJobsRequest(input *ListRxNormInferenceJobsInput) ListRxNormInferenceJobsRequest {
	op := &aws.Operation{
		Name:       opListRxNormInferenceJobs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListRxNormInferenceJobsInput{}
	}

	req := c.newRequest(op, input, &ListRxNormInferenceJobsOutput{})

	return ListRxNormInferenceJobsRequest{Request: req, Input: input, Copy: c.ListRxNormInferenceJobsRequest}
}

// ListRxNormInferenceJobsRequest is the request type for the
// ListRxNormInferenceJobs API operation.
type ListRxNormInferenceJobsRequest struct {
	*aws.Request
	Input *ListRxNormInferenceJobsInput
	Copy  func(*ListRxNormInferenceJobsInput) ListRxNormInferenceJobsRequest
}

// Send marshals and sends the ListRxNormInferenceJobs API request.
func (r ListRxNormInferenceJobsRequest) Send(ctx context.Context) (*ListRxNormInferenceJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRxNormInferenceJobsResponse{
		ListRxNormInferenceJobsOutput: r.Request.Data.(*ListRxNormInferenceJobsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListRxNormInferenceJobsResponse is the response type for the
// ListRxNormInferenceJobs API operation.
type ListRxNormInferenceJobsResponse struct {
	*ListRxNormInferenceJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRxNormInferenceJobs request.
func (r *ListRxNormInferenceJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
