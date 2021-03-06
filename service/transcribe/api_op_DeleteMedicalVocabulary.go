// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transcribe

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteMedicalVocabularyInput struct {
	_ struct{} `type:"structure"`

	// The name of the vocabulary you are choosing to delete.
	//
	// VocabularyName is a required field
	VocabularyName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMedicalVocabularyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMedicalVocabularyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteMedicalVocabularyInput"}

	if s.VocabularyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VocabularyName"))
	}
	if s.VocabularyName != nil && len(*s.VocabularyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VocabularyName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteMedicalVocabularyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteMedicalVocabularyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteMedicalVocabulary = "DeleteMedicalVocabulary"

// DeleteMedicalVocabularyRequest returns a request value for making API operation for
// Amazon Transcribe Service.
//
// Deletes a vocabulary from Amazon Transcribe Medical.
//
//    // Example sending a request using DeleteMedicalVocabularyRequest.
//    req := client.DeleteMedicalVocabularyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/DeleteMedicalVocabulary
func (c *Client) DeleteMedicalVocabularyRequest(input *DeleteMedicalVocabularyInput) DeleteMedicalVocabularyRequest {
	op := &aws.Operation{
		Name:       opDeleteMedicalVocabulary,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteMedicalVocabularyInput{}
	}

	req := c.newRequest(op, input, &DeleteMedicalVocabularyOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteMedicalVocabularyRequest{Request: req, Input: input, Copy: c.DeleteMedicalVocabularyRequest}
}

// DeleteMedicalVocabularyRequest is the request type for the
// DeleteMedicalVocabulary API operation.
type DeleteMedicalVocabularyRequest struct {
	*aws.Request
	Input *DeleteMedicalVocabularyInput
	Copy  func(*DeleteMedicalVocabularyInput) DeleteMedicalVocabularyRequest
}

// Send marshals and sends the DeleteMedicalVocabulary API request.
func (r DeleteMedicalVocabularyRequest) Send(ctx context.Context) (*DeleteMedicalVocabularyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteMedicalVocabularyResponse{
		DeleteMedicalVocabularyOutput: r.Request.Data.(*DeleteMedicalVocabularyOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteMedicalVocabularyResponse is the response type for the
// DeleteMedicalVocabulary API operation.
type DeleteMedicalVocabularyResponse struct {
	*DeleteMedicalVocabularyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteMedicalVocabulary request.
func (r *DeleteMedicalVocabularyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
