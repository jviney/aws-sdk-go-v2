// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to the get job operation.
type GetJobInput struct {
	_ struct{} `type:"structure"`

	// The job's ARN.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s GetJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetJobInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the result of a get job request.
type GetJobOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains information about the requested job.
	Job *Job `locationName:"job" type:"structure"`
}

// String returns the string representation
func (s GetJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetJob = "GetJob"

// GetJobRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Gets information about a job.
//
//    // Example sending a request using GetJobRequest.
//    req := client.GetJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/GetJob
func (c *Client) GetJobRequest(input *GetJobInput) GetJobRequest {
	op := &aws.Operation{
		Name:       opGetJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetJobInput{}
	}

	req := c.newRequest(op, input, &GetJobOutput{})

	return GetJobRequest{Request: req, Input: input, Copy: c.GetJobRequest}
}

// GetJobRequest is the request type for the
// GetJob API operation.
type GetJobRequest struct {
	*aws.Request
	Input *GetJobInput
	Copy  func(*GetJobInput) GetJobRequest
}

// Send marshals and sends the GetJob API request.
func (r GetJobRequest) Send(ctx context.Context) (*GetJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetJobResponse{
		GetJobOutput: r.Request.Data.(*GetJobOutput),
		response:     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetJobResponse is the response type for the
// GetJob API operation.
type GetJobResponse struct {
	*GetJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetJob request.
func (r *GetJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
