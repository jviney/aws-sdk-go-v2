// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// The CancelDomainTransferToAnotherAwsAccount request includes the following
// element.
type CancelDomainTransferToAnotherAwsAccountInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain for which you want to cancel the transfer to another
	// AWS account.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CancelDomainTransferToAnotherAwsAccountInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelDomainTransferToAnotherAwsAccountInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelDomainTransferToAnotherAwsAccountInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The CancelDomainTransferToAnotherAwsAccount response includes the following
// element.
type CancelDomainTransferToAnotherAwsAccountOutput struct {
	_ struct{} `type:"structure"`

	// The identifier that TransferDomainToAnotherAwsAccount returned to track the
	// progress of the request. Because the transfer request was canceled, the value
	// is no longer valid, and you can't use GetOperationDetail to query the operation
	// status.
	OperationId *string `type:"string"`
}

// String returns the string representation
func (s CancelDomainTransferToAnotherAwsAccountOutput) String() string {
	return awsutil.Prettify(s)
}

const opCancelDomainTransferToAnotherAwsAccount = "CancelDomainTransferToAnotherAwsAccount"

// CancelDomainTransferToAnotherAwsAccountRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// Cancels the transfer of a domain from the current AWS account to another
// AWS account. You initiate a transfer between AWS accounts using TransferDomainToAnotherAwsAccount
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html).
//
// You must cancel the transfer before the other AWS account accepts the transfer
// using AcceptDomainTransferFromAnotherAwsAccount (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_AcceptDomainTransferFromAnotherAwsAccount.html).
//
// Use either ListOperations (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ListOperations.html)
// or GetOperationDetail (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html)
// to determine whether the operation succeeded. GetOperationDetail (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html)
// provides additional information, for example, Domain Transfer from Aws Account
// 111122223333 has been cancelled.
//
//    // Example sending a request using CancelDomainTransferToAnotherAwsAccountRequest.
//    req := client.CancelDomainTransferToAnotherAwsAccountRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/CancelDomainTransferToAnotherAwsAccount
func (c *Client) CancelDomainTransferToAnotherAwsAccountRequest(input *CancelDomainTransferToAnotherAwsAccountInput) CancelDomainTransferToAnotherAwsAccountRequest {
	op := &aws.Operation{
		Name:       opCancelDomainTransferToAnotherAwsAccount,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CancelDomainTransferToAnotherAwsAccountInput{}
	}

	req := c.newRequest(op, input, &CancelDomainTransferToAnotherAwsAccountOutput{})

	return CancelDomainTransferToAnotherAwsAccountRequest{Request: req, Input: input, Copy: c.CancelDomainTransferToAnotherAwsAccountRequest}
}

// CancelDomainTransferToAnotherAwsAccountRequest is the request type for the
// CancelDomainTransferToAnotherAwsAccount API operation.
type CancelDomainTransferToAnotherAwsAccountRequest struct {
	*aws.Request
	Input *CancelDomainTransferToAnotherAwsAccountInput
	Copy  func(*CancelDomainTransferToAnotherAwsAccountInput) CancelDomainTransferToAnotherAwsAccountRequest
}

// Send marshals and sends the CancelDomainTransferToAnotherAwsAccount API request.
func (r CancelDomainTransferToAnotherAwsAccountRequest) Send(ctx context.Context) (*CancelDomainTransferToAnotherAwsAccountResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelDomainTransferToAnotherAwsAccountResponse{
		CancelDomainTransferToAnotherAwsAccountOutput: r.Request.Data.(*CancelDomainTransferToAnotherAwsAccountOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelDomainTransferToAnotherAwsAccountResponse is the response type for the
// CancelDomainTransferToAnotherAwsAccount API operation.
type CancelDomainTransferToAnotherAwsAccountResponse struct {
	*CancelDomainTransferToAnotherAwsAccountOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelDomainTransferToAnotherAwsAccount request.
func (r *CancelDomainTransferToAnotherAwsAccountResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
