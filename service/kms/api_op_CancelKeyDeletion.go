// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CancelKeyDeletionInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the customer master key (CMK) for which to cancel
	// deletion.
	//
	// Specify the key ID or the Amazon Resource Name (ARN) of the CMK.
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// KeyId is a required field
	KeyId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CancelKeyDeletionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelKeyDeletionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelKeyDeletionInput"}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CancelKeyDeletionOutput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the master key for which deletion is canceled.
	KeyId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CancelKeyDeletionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCancelKeyDeletion = "CancelKeyDeletion"

// CancelKeyDeletionRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Cancels the deletion of a customer master key (CMK). When this operation
// succeeds, the key state of the CMK is Disabled. To enable the CMK, use EnableKey.
// You cannot perform this operation on a CMK in a different AWS account.
//
// For more information about scheduling and canceling deletion of a CMK, see
// Deleting Customer Master Keys (https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html)
// in the AWS Key Management Service Developer Guide.
//
// The CMK that you use for this operation must be in a compatible key state.
// For details, see How Key State Affects Use of a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the AWS Key Management Service Developer Guide.
//
//    // Example sending a request using CancelKeyDeletionRequest.
//    req := client.CancelKeyDeletionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/CancelKeyDeletion
func (c *Client) CancelKeyDeletionRequest(input *CancelKeyDeletionInput) CancelKeyDeletionRequest {
	op := &aws.Operation{
		Name:       opCancelKeyDeletion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CancelKeyDeletionInput{}
	}

	req := c.newRequest(op, input, &CancelKeyDeletionOutput{})

	return CancelKeyDeletionRequest{Request: req, Input: input, Copy: c.CancelKeyDeletionRequest}
}

// CancelKeyDeletionRequest is the request type for the
// CancelKeyDeletion API operation.
type CancelKeyDeletionRequest struct {
	*aws.Request
	Input *CancelKeyDeletionInput
	Copy  func(*CancelKeyDeletionInput) CancelKeyDeletionRequest
}

// Send marshals and sends the CancelKeyDeletion API request.
func (r CancelKeyDeletionRequest) Send(ctx context.Context) (*CancelKeyDeletionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelKeyDeletionResponse{
		CancelKeyDeletionOutput: r.Request.Data.(*CancelKeyDeletionOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelKeyDeletionResponse is the response type for the
// CancelKeyDeletion API operation.
type CancelKeyDeletionResponse struct {
	*CancelKeyDeletionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelKeyDeletion request.
func (r *CancelKeyDeletionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
