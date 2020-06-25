// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetWorkflowRunPropertiesInput struct {
	_ struct{} `type:"structure"`

	// Name of the workflow which was run.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The ID of the workflow run whose run properties should be returned.
	//
	// RunId is a required field
	RunId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetWorkflowRunPropertiesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetWorkflowRunPropertiesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetWorkflowRunPropertiesInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.RunId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RunId"))
	}
	if s.RunId != nil && len(*s.RunId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RunId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetWorkflowRunPropertiesOutput struct {
	_ struct{} `type:"structure"`

	// The workflow run properties which were set during the specified run.
	RunProperties map[string]string `type:"map"`
}

// String returns the string representation
func (s GetWorkflowRunPropertiesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetWorkflowRunProperties = "GetWorkflowRunProperties"

// GetWorkflowRunPropertiesRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves the workflow run properties which were set during the run.
//
//    // Example sending a request using GetWorkflowRunPropertiesRequest.
//    req := client.GetWorkflowRunPropertiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetWorkflowRunProperties
func (c *Client) GetWorkflowRunPropertiesRequest(input *GetWorkflowRunPropertiesInput) GetWorkflowRunPropertiesRequest {
	op := &aws.Operation{
		Name:       opGetWorkflowRunProperties,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetWorkflowRunPropertiesInput{}
	}

	req := c.newRequest(op, input, &GetWorkflowRunPropertiesOutput{})

	return GetWorkflowRunPropertiesRequest{Request: req, Input: input, Copy: c.GetWorkflowRunPropertiesRequest}
}

// GetWorkflowRunPropertiesRequest is the request type for the
// GetWorkflowRunProperties API operation.
type GetWorkflowRunPropertiesRequest struct {
	*aws.Request
	Input *GetWorkflowRunPropertiesInput
	Copy  func(*GetWorkflowRunPropertiesInput) GetWorkflowRunPropertiesRequest
}

// Send marshals and sends the GetWorkflowRunProperties API request.
func (r GetWorkflowRunPropertiesRequest) Send(ctx context.Context) (*GetWorkflowRunPropertiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetWorkflowRunPropertiesResponse{
		GetWorkflowRunPropertiesOutput: r.Request.Data.(*GetWorkflowRunPropertiesOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetWorkflowRunPropertiesResponse is the response type for the
// GetWorkflowRunProperties API operation.
type GetWorkflowRunPropertiesResponse struct {
	*GetWorkflowRunPropertiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetWorkflowRunProperties request.
func (r *GetWorkflowRunPropertiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
