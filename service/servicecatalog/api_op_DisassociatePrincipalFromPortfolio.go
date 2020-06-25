// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DisassociatePrincipalFromPortfolioInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The portfolio identifier.
	//
	// PortfolioId is a required field
	PortfolioId *string `min:"1" type:"string" required:"true"`

	// The ARN of the principal (IAM user, role, or group).
	//
	// PrincipalARN is a required field
	PrincipalARN *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociatePrincipalFromPortfolioInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociatePrincipalFromPortfolioInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociatePrincipalFromPortfolioInput"}

	if s.PortfolioId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PortfolioId"))
	}
	if s.PortfolioId != nil && len(*s.PortfolioId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PortfolioId", 1))
	}

	if s.PrincipalARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("PrincipalARN"))
	}
	if s.PrincipalARN != nil && len(*s.PrincipalARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PrincipalARN", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisassociatePrincipalFromPortfolioOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociatePrincipalFromPortfolioOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociatePrincipalFromPortfolio = "DisassociatePrincipalFromPortfolio"

// DisassociatePrincipalFromPortfolioRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Disassociates a previously associated principal ARN from a specified portfolio.
//
//    // Example sending a request using DisassociatePrincipalFromPortfolioRequest.
//    req := client.DisassociatePrincipalFromPortfolioRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DisassociatePrincipalFromPortfolio
func (c *Client) DisassociatePrincipalFromPortfolioRequest(input *DisassociatePrincipalFromPortfolioInput) DisassociatePrincipalFromPortfolioRequest {
	op := &aws.Operation{
		Name:       opDisassociatePrincipalFromPortfolio,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociatePrincipalFromPortfolioInput{}
	}

	req := c.newRequest(op, input, &DisassociatePrincipalFromPortfolioOutput{})

	return DisassociatePrincipalFromPortfolioRequest{Request: req, Input: input, Copy: c.DisassociatePrincipalFromPortfolioRequest}
}

// DisassociatePrincipalFromPortfolioRequest is the request type for the
// DisassociatePrincipalFromPortfolio API operation.
type DisassociatePrincipalFromPortfolioRequest struct {
	*aws.Request
	Input *DisassociatePrincipalFromPortfolioInput
	Copy  func(*DisassociatePrincipalFromPortfolioInput) DisassociatePrincipalFromPortfolioRequest
}

// Send marshals and sends the DisassociatePrincipalFromPortfolio API request.
func (r DisassociatePrincipalFromPortfolioRequest) Send(ctx context.Context) (*DisassociatePrincipalFromPortfolioResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociatePrincipalFromPortfolioResponse{
		DisassociatePrincipalFromPortfolioOutput: r.Request.Data.(*DisassociatePrincipalFromPortfolioOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociatePrincipalFromPortfolioResponse is the response type for the
// DisassociatePrincipalFromPortfolio API operation.
type DisassociatePrincipalFromPortfolioResponse struct {
	*DisassociatePrincipalFromPortfolioOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociatePrincipalFromPortfolio request.
func (r *DisassociatePrincipalFromPortfolioResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
