// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteTrialComponentInput struct {
	_ struct{} `type:"structure"`

	// The name of the component to delete.
	//
	// TrialComponentName is a required field
	TrialComponentName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTrialComponentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTrialComponentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTrialComponentInput"}

	if s.TrialComponentName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrialComponentName"))
	}
	if s.TrialComponentName != nil && len(*s.TrialComponentName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TrialComponentName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteTrialComponentOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the component is being deleted.
	TrialComponentArn *string `type:"string"`
}

// String returns the string representation
func (s DeleteTrialComponentOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTrialComponent = "DeleteTrialComponent"

// DeleteTrialComponentRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Deletes the specified trial component. A trial component must be disassociated
// from all trials before the trial component can be deleted. To disassociate
// a trial component from a trial, call the DisassociateTrialComponent API.
//
//    // Example sending a request using DeleteTrialComponentRequest.
//    req := client.DeleteTrialComponentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DeleteTrialComponent
func (c *Client) DeleteTrialComponentRequest(input *DeleteTrialComponentInput) DeleteTrialComponentRequest {
	op := &aws.Operation{
		Name:       opDeleteTrialComponent,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTrialComponentInput{}
	}

	req := c.newRequest(op, input, &DeleteTrialComponentOutput{})

	return DeleteTrialComponentRequest{Request: req, Input: input, Copy: c.DeleteTrialComponentRequest}
}

// DeleteTrialComponentRequest is the request type for the
// DeleteTrialComponent API operation.
type DeleteTrialComponentRequest struct {
	*aws.Request
	Input *DeleteTrialComponentInput
	Copy  func(*DeleteTrialComponentInput) DeleteTrialComponentRequest
}

// Send marshals and sends the DeleteTrialComponent API request.
func (r DeleteTrialComponentRequest) Send(ctx context.Context) (*DeleteTrialComponentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTrialComponentResponse{
		DeleteTrialComponentOutput: r.Request.Data.(*DeleteTrialComponentOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTrialComponentResponse is the response type for the
// DeleteTrialComponent API operation.
type DeleteTrialComponentResponse struct {
	*DeleteTrialComponentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTrialComponent request.
func (r *DeleteTrialComponentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
