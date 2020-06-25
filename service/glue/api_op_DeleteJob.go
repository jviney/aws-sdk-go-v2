// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteJobInput struct {
	_ struct{} `type:"structure"`

	// The name of the job definition to delete.
	//
	// JobName is a required field
	JobName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteJobInput"}

	if s.JobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobName"))
	}
	if s.JobName != nil && len(*s.JobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteJobOutput struct {
	_ struct{} `type:"structure"`

	// The name of the job definition that was deleted.
	JobName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteJob = "DeleteJob"

// DeleteJobRequest returns a request value for making API operation for
// AWS Glue.
//
// Deletes a specified job definition. If the job definition is not found, no
// exception is thrown.
//
//    // Example sending a request using DeleteJobRequest.
//    req := client.DeleteJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/DeleteJob
func (c *Client) DeleteJobRequest(input *DeleteJobInput) DeleteJobRequest {
	op := &aws.Operation{
		Name:       opDeleteJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteJobInput{}
	}

	req := c.newRequest(op, input, &DeleteJobOutput{})

	return DeleteJobRequest{Request: req, Input: input, Copy: c.DeleteJobRequest}
}

// DeleteJobRequest is the request type for the
// DeleteJob API operation.
type DeleteJobRequest struct {
	*aws.Request
	Input *DeleteJobInput
	Copy  func(*DeleteJobInput) DeleteJobRequest
}

// Send marshals and sends the DeleteJob API request.
func (r DeleteJobRequest) Send(ctx context.Context) (*DeleteJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteJobResponse{
		DeleteJobOutput: r.Request.Data.(*DeleteJobOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteJobResponse is the response type for the
// DeleteJob API operation.
type DeleteJobResponse struct {
	*DeleteJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteJob request.
func (r *DeleteJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
