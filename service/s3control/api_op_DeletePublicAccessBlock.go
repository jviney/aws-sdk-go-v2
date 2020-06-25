// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3control

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/restxml"
)

type DeletePublicAccessBlockInput struct {
	_ struct{} `type:"structure"`

	// The account ID for the Amazon Web Services account whose PublicAccessBlock
	// configuration you want to remove.
	//
	// AccountId is a required field
	AccountId *string `location:"header" locationName:"x-amz-account-id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePublicAccessBlockInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePublicAccessBlockInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePublicAccessBlockInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeletePublicAccessBlockInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-account-id", protocol.StringValue(v), metadata)
	}
	return nil
}

type DeletePublicAccessBlockOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeletePublicAccessBlockOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeletePublicAccessBlockOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeletePublicAccessBlock = "DeletePublicAccessBlock"

// DeletePublicAccessBlockRequest returns a request value for making API operation for
// AWS S3 Control.
//
// Removes the PublicAccessBlock configuration for an Amazon Web Services account.
//
//    // Example sending a request using DeletePublicAccessBlockRequest.
//    req := client.DeletePublicAccessBlockRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3control-2018-08-20/DeletePublicAccessBlock
func (c *Client) DeletePublicAccessBlockRequest(input *DeletePublicAccessBlockInput) DeletePublicAccessBlockRequest {
	op := &aws.Operation{
		Name:       opDeletePublicAccessBlock,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v20180820/configuration/publicAccessBlock",
	}

	if input == nil {
		input = &DeletePublicAccessBlockInput{}
	}

	req := c.newRequest(op, input, &DeletePublicAccessBlockOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	req.Handlers.Build.PushBackNamed(buildPrefixHostHandler("AccountID", aws.StringValue(input.AccountId)))
	req.Handlers.Build.PushBackNamed(buildRemoveHeaderHandler("X-Amz-Account-Id"))

	return DeletePublicAccessBlockRequest{Request: req, Input: input, Copy: c.DeletePublicAccessBlockRequest}
}

// DeletePublicAccessBlockRequest is the request type for the
// DeletePublicAccessBlock API operation.
type DeletePublicAccessBlockRequest struct {
	*aws.Request
	Input *DeletePublicAccessBlockInput
	Copy  func(*DeletePublicAccessBlockInput) DeletePublicAccessBlockRequest
}

// Send marshals and sends the DeletePublicAccessBlock API request.
func (r DeletePublicAccessBlockRequest) Send(ctx context.Context) (*DeletePublicAccessBlockResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePublicAccessBlockResponse{
		DeletePublicAccessBlockOutput: r.Request.Data.(*DeletePublicAccessBlockOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePublicAccessBlockResponse is the response type for the
// DeletePublicAccessBlock API operation.
type DeletePublicAccessBlockResponse struct {
	*DeletePublicAccessBlockOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePublicAccessBlock request.
func (r *DeletePublicAccessBlockResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
