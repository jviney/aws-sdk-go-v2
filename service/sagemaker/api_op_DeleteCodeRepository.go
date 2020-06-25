// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteCodeRepositoryInput struct {
	_ struct{} `type:"structure"`

	// The name of the Git repository to delete.
	//
	// CodeRepositoryName is a required field
	CodeRepositoryName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCodeRepositoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCodeRepositoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCodeRepositoryInput"}

	if s.CodeRepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CodeRepositoryName"))
	}
	if s.CodeRepositoryName != nil && len(*s.CodeRepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CodeRepositoryName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteCodeRepositoryOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteCodeRepositoryOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteCodeRepository = "DeleteCodeRepository"

// DeleteCodeRepositoryRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Deletes the specified Git repository from your account.
//
//    // Example sending a request using DeleteCodeRepositoryRequest.
//    req := client.DeleteCodeRepositoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DeleteCodeRepository
func (c *Client) DeleteCodeRepositoryRequest(input *DeleteCodeRepositoryInput) DeleteCodeRepositoryRequest {
	op := &aws.Operation{
		Name:       opDeleteCodeRepository,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCodeRepositoryInput{}
	}

	req := c.newRequest(op, input, &DeleteCodeRepositoryOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteCodeRepositoryRequest{Request: req, Input: input, Copy: c.DeleteCodeRepositoryRequest}
}

// DeleteCodeRepositoryRequest is the request type for the
// DeleteCodeRepository API operation.
type DeleteCodeRepositoryRequest struct {
	*aws.Request
	Input *DeleteCodeRepositoryInput
	Copy  func(*DeleteCodeRepositoryInput) DeleteCodeRepositoryRequest
}

// Send marshals and sends the DeleteCodeRepository API request.
func (r DeleteCodeRepositoryRequest) Send(ctx context.Context) (*DeleteCodeRepositoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCodeRepositoryResponse{
		DeleteCodeRepositoryOutput: r.Request.Data.(*DeleteCodeRepositoryOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCodeRepositoryResponse is the response type for the
// DeleteCodeRepository API operation.
type DeleteCodeRepositoryResponse struct {
	*DeleteCodeRepositoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCodeRepository request.
func (r *DeleteCodeRepositoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
