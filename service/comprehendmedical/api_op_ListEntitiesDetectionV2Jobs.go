// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehendmedical

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListEntitiesDetectionV2JobsInput struct {
	_ struct{} `type:"structure"`

	// Filters the jobs that are returned. You can filter jobs based on their names,
	// status, or the date and time that they were submitted. You can only set one
	// filter at a time.
	Filter *ComprehendMedicalAsyncJobFilter `type:"structure"`

	// The maximum number of results to return in each page. The default is 100.
	MaxResults *int64 `min:"1" type:"integer"`

	// Identifies the next page of results to return.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListEntitiesDetectionV2JobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListEntitiesDetectionV2JobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListEntitiesDetectionV2JobsInput"}
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

type ListEntitiesDetectionV2JobsOutput struct {
	_ struct{} `type:"structure"`

	// A list containing the properties of each job returned.
	ComprehendMedicalAsyncJobPropertiesList []ComprehendMedicalAsyncJobProperties `type:"list"`

	// Identifies the next page of results to return.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListEntitiesDetectionV2JobsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListEntitiesDetectionV2Jobs = "ListEntitiesDetectionV2Jobs"

// ListEntitiesDetectionV2JobsRequest returns a request value for making API operation for
// AWS Comprehend Medical.
//
// Gets a list of medical entity detection jobs that you have submitted.
//
//    // Example sending a request using ListEntitiesDetectionV2JobsRequest.
//    req := client.ListEntitiesDetectionV2JobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehendmedical-2018-10-30/ListEntitiesDetectionV2Jobs
func (c *Client) ListEntitiesDetectionV2JobsRequest(input *ListEntitiesDetectionV2JobsInput) ListEntitiesDetectionV2JobsRequest {
	op := &aws.Operation{
		Name:       opListEntitiesDetectionV2Jobs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListEntitiesDetectionV2JobsInput{}
	}

	req := c.newRequest(op, input, &ListEntitiesDetectionV2JobsOutput{})

	return ListEntitiesDetectionV2JobsRequest{Request: req, Input: input, Copy: c.ListEntitiesDetectionV2JobsRequest}
}

// ListEntitiesDetectionV2JobsRequest is the request type for the
// ListEntitiesDetectionV2Jobs API operation.
type ListEntitiesDetectionV2JobsRequest struct {
	*aws.Request
	Input *ListEntitiesDetectionV2JobsInput
	Copy  func(*ListEntitiesDetectionV2JobsInput) ListEntitiesDetectionV2JobsRequest
}

// Send marshals and sends the ListEntitiesDetectionV2Jobs API request.
func (r ListEntitiesDetectionV2JobsRequest) Send(ctx context.Context) (*ListEntitiesDetectionV2JobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListEntitiesDetectionV2JobsResponse{
		ListEntitiesDetectionV2JobsOutput: r.Request.Data.(*ListEntitiesDetectionV2JobsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListEntitiesDetectionV2JobsResponse is the response type for the
// ListEntitiesDetectionV2Jobs API operation.
type ListEntitiesDetectionV2JobsResponse struct {
	*ListEntitiesDetectionV2JobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListEntitiesDetectionV2Jobs request.
func (r *ListEntitiesDetectionV2JobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
