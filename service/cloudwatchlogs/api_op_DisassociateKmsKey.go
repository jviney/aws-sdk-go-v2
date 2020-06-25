// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DisassociateKmsKeyInput struct {
	_ struct{} `type:"structure"`

	// The name of the log group.
	//
	// LogGroupName is a required field
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateKmsKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateKmsKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateKmsKeyInput"}

	if s.LogGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogGroupName"))
	}
	if s.LogGroupName != nil && len(*s.LogGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisassociateKmsKeyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateKmsKeyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateKmsKey = "DisassociateKmsKey"

// DisassociateKmsKeyRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Disassociates the associated AWS Key Management Service (AWS KMS) customer
// master key (CMK) from the specified log group.
//
// After the AWS KMS CMK is disassociated from the log group, AWS CloudWatch
// Logs stops encrypting newly ingested data for the log group. All previously
// ingested data remains encrypted, and AWS CloudWatch Logs requires permissions
// for the CMK whenever the encrypted data is requested.
//
// Note that it can take up to 5 minutes for this operation to take effect.
//
//    // Example sending a request using DisassociateKmsKeyRequest.
//    req := client.DisassociateKmsKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DisassociateKmsKey
func (c *Client) DisassociateKmsKeyRequest(input *DisassociateKmsKeyInput) DisassociateKmsKeyRequest {
	op := &aws.Operation{
		Name:       opDisassociateKmsKey,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateKmsKeyInput{}
	}

	req := c.newRequest(op, input, &DisassociateKmsKeyOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DisassociateKmsKeyRequest{Request: req, Input: input, Copy: c.DisassociateKmsKeyRequest}
}

// DisassociateKmsKeyRequest is the request type for the
// DisassociateKmsKey API operation.
type DisassociateKmsKeyRequest struct {
	*aws.Request
	Input *DisassociateKmsKeyInput
	Copy  func(*DisassociateKmsKeyInput) DisassociateKmsKeyRequest
}

// Send marshals and sends the DisassociateKmsKey API request.
func (r DisassociateKmsKeyRequest) Send(ctx context.Context) (*DisassociateKmsKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateKmsKeyResponse{
		DisassociateKmsKeyOutput: r.Request.Data.(*DisassociateKmsKeyOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateKmsKeyResponse is the response type for the
// DisassociateKmsKey API operation.
type DisassociateKmsKeyResponse struct {
	*DisassociateKmsKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateKmsKey request.
func (r *DisassociateKmsKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
