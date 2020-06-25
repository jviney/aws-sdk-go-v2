// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// DescribeSMBFileSharesInput
type DescribeSMBFileSharesInput struct {
	_ struct{} `type:"structure"`

	// An array containing the Amazon Resource Name (ARN) of each file share to
	// be described.
	//
	// FileShareARNList is a required field
	FileShareARNList []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeSMBFileSharesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSMBFileSharesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSMBFileSharesInput"}

	if s.FileShareARNList == nil {
		invalidParams.Add(aws.NewErrParamRequired("FileShareARNList"))
	}
	if s.FileShareARNList != nil && len(s.FileShareARNList) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FileShareARNList", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// DescribeSMBFileSharesOutput
type DescribeSMBFileSharesOutput struct {
	_ struct{} `type:"structure"`

	// An array containing a description for each requested file share.
	SMBFileShareInfoList []SMBFileShareInfo `type:"list"`
}

// String returns the string representation
func (s DescribeSMBFileSharesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSMBFileShares = "DescribeSMBFileShares"

// DescribeSMBFileSharesRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Gets a description for one or more Server Message Block (SMB) file shares
// from a file gateway. This operation is only supported for file gateways.
//
//    // Example sending a request using DescribeSMBFileSharesRequest.
//    req := client.DescribeSMBFileSharesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DescribeSMBFileShares
func (c *Client) DescribeSMBFileSharesRequest(input *DescribeSMBFileSharesInput) DescribeSMBFileSharesRequest {
	op := &aws.Operation{
		Name:       opDescribeSMBFileShares,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSMBFileSharesInput{}
	}

	req := c.newRequest(op, input, &DescribeSMBFileSharesOutput{})

	return DescribeSMBFileSharesRequest{Request: req, Input: input, Copy: c.DescribeSMBFileSharesRequest}
}

// DescribeSMBFileSharesRequest is the request type for the
// DescribeSMBFileShares API operation.
type DescribeSMBFileSharesRequest struct {
	*aws.Request
	Input *DescribeSMBFileSharesInput
	Copy  func(*DescribeSMBFileSharesInput) DescribeSMBFileSharesRequest
}

// Send marshals and sends the DescribeSMBFileShares API request.
func (r DescribeSMBFileSharesRequest) Send(ctx context.Context) (*DescribeSMBFileSharesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSMBFileSharesResponse{
		DescribeSMBFileSharesOutput: r.Request.Data.(*DescribeSMBFileSharesOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSMBFileSharesResponse is the response type for the
// DescribeSMBFileShares API operation.
type DescribeSMBFileSharesResponse struct {
	*DescribeSMBFileSharesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSMBFileShares request.
func (r *DescribeSMBFileSharesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
