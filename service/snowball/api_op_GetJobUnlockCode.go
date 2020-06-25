// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package snowball

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetJobUnlockCodeInput struct {
	_ struct{} `type:"structure"`

	// The ID for the job that you want to get the UnlockCode value for, for example
	// JID123e4567-e89b-12d3-a456-426655440000.
	//
	// JobId is a required field
	JobId *string `min:"39" type:"string" required:"true"`
}

// String returns the string representation
func (s GetJobUnlockCodeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetJobUnlockCodeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetJobUnlockCodeInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 39 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 39))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetJobUnlockCodeOutput struct {
	_ struct{} `type:"structure"`

	// The UnlockCode value for the specified job. The UnlockCode value can be accessed
	// for up to 90 days after the job has been created.
	UnlockCode *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetJobUnlockCodeOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetJobUnlockCode = "GetJobUnlockCode"

// GetJobUnlockCodeRequest returns a request value for making API operation for
// Amazon Import/Export Snowball.
//
// Returns the UnlockCode code value for the specified job. A particular UnlockCode
// value can be accessed for up to 90 days after the associated job has been
// created.
//
// The UnlockCode value is a 29-character code with 25 alphanumeric characters
// and 4 hyphens. This code is used to decrypt the manifest file when it is
// passed along with the manifest to the Snowball through the Snowball client
// when the client is started for the first time.
//
// As a best practice, we recommend that you don't save a copy of the UnlockCode
// in the same location as the manifest file for that job. Saving these separately
// helps prevent unauthorized parties from gaining access to the Snowball associated
// with that job.
//
//    // Example sending a request using GetJobUnlockCodeRequest.
//    req := client.GetJobUnlockCodeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/GetJobUnlockCode
func (c *Client) GetJobUnlockCodeRequest(input *GetJobUnlockCodeInput) GetJobUnlockCodeRequest {
	op := &aws.Operation{
		Name:       opGetJobUnlockCode,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetJobUnlockCodeInput{}
	}

	req := c.newRequest(op, input, &GetJobUnlockCodeOutput{})

	return GetJobUnlockCodeRequest{Request: req, Input: input, Copy: c.GetJobUnlockCodeRequest}
}

// GetJobUnlockCodeRequest is the request type for the
// GetJobUnlockCode API operation.
type GetJobUnlockCodeRequest struct {
	*aws.Request
	Input *GetJobUnlockCodeInput
	Copy  func(*GetJobUnlockCodeInput) GetJobUnlockCodeRequest
}

// Send marshals and sends the GetJobUnlockCode API request.
func (r GetJobUnlockCodeRequest) Send(ctx context.Context) (*GetJobUnlockCodeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetJobUnlockCodeResponse{
		GetJobUnlockCodeOutput: r.Request.Data.(*GetJobUnlockCodeOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetJobUnlockCodeResponse is the response type for the
// GetJobUnlockCode API operation.
type GetJobUnlockCodeResponse struct {
	*GetJobUnlockCodeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetJobUnlockCode request.
func (r *GetJobUnlockCodeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
