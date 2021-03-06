// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package frauddetector

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateModelVersionInput struct {
	_ struct{} `type:"structure"`

	// The model version description.
	Description *string `locationName:"description" min:"1" type:"string"`

	// The model ID.
	//
	// ModelId is a required field
	ModelId *string `locationName:"modelId" min:"1" type:"string" required:"true"`

	// The model type.
	//
	// ModelType is a required field
	ModelType ModelTypeEnum `locationName:"modelType" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreateModelVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateModelVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateModelVersionInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}

	if s.ModelId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ModelId"))
	}
	if s.ModelId != nil && len(*s.ModelId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ModelId", 1))
	}
	if len(s.ModelType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ModelType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateModelVersionOutput struct {
	_ struct{} `type:"structure"`

	// The model ID.
	ModelId *string `locationName:"modelId" min:"1" type:"string"`

	// The model type.
	ModelType ModelTypeEnum `locationName:"modelType" type:"string" enum:"true"`

	// The version of the model.
	ModelVersionNumber *string `locationName:"modelVersionNumber" min:"1" type:"string"`

	// The model version status.
	Status *string `locationName:"status" type:"string"`
}

// String returns the string representation
func (s CreateModelVersionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateModelVersion = "CreateModelVersion"

// CreateModelVersionRequest returns a request value for making API operation for
// Amazon Fraud Detector.
//
// Creates a version of the model using the specified model type.
//
//    // Example sending a request using CreateModelVersionRequest.
//    req := client.CreateModelVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/frauddetector-2019-11-15/CreateModelVersion
func (c *Client) CreateModelVersionRequest(input *CreateModelVersionInput) CreateModelVersionRequest {
	op := &aws.Operation{
		Name:       opCreateModelVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateModelVersionInput{}
	}

	req := c.newRequest(op, input, &CreateModelVersionOutput{})

	return CreateModelVersionRequest{Request: req, Input: input, Copy: c.CreateModelVersionRequest}
}

// CreateModelVersionRequest is the request type for the
// CreateModelVersion API operation.
type CreateModelVersionRequest struct {
	*aws.Request
	Input *CreateModelVersionInput
	Copy  func(*CreateModelVersionInput) CreateModelVersionRequest
}

// Send marshals and sends the CreateModelVersion API request.
func (r CreateModelVersionRequest) Send(ctx context.Context) (*CreateModelVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateModelVersionResponse{
		CreateModelVersionOutput: r.Request.Data.(*CreateModelVersionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateModelVersionResponse is the response type for the
// CreateModelVersion API operation.
type CreateModelVersionResponse struct {
	*CreateModelVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateModelVersion request.
func (r *CreateModelVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
