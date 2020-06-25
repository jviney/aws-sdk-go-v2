// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type AcceptPortfolioShareInput struct {
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

	// The type of shared portfolios to accept. The default is to accept imported
	// portfolios.
	//
	//    * AWS_ORGANIZATIONS - Accept portfolios shared by the master account of
	//    your organization.
	//
	//    * IMPORTED - Accept imported portfolios.
	//
	//    * AWS_SERVICECATALOG - Not supported. (Throws ResourceNotFoundException.)
	//
	// For example, aws servicecatalog accept-portfolio-share --portfolio-id "port-2qwzkwxt3y5fk"
	// --portfolio-share-type AWS_ORGANIZATIONS
	PortfolioShareType PortfolioShareType `type:"string" enum:"true"`
}

// String returns the string representation
func (s AcceptPortfolioShareInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AcceptPortfolioShareInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AcceptPortfolioShareInput"}

	if s.PortfolioId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PortfolioId"))
	}
	if s.PortfolioId != nil && len(*s.PortfolioId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PortfolioId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AcceptPortfolioShareOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AcceptPortfolioShareOutput) String() string {
	return awsutil.Prettify(s)
}

const opAcceptPortfolioShare = "AcceptPortfolioShare"

// AcceptPortfolioShareRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Accepts an offer to share the specified portfolio.
//
//    // Example sending a request using AcceptPortfolioShareRequest.
//    req := client.AcceptPortfolioShareRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/AcceptPortfolioShare
func (c *Client) AcceptPortfolioShareRequest(input *AcceptPortfolioShareInput) AcceptPortfolioShareRequest {
	op := &aws.Operation{
		Name:       opAcceptPortfolioShare,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AcceptPortfolioShareInput{}
	}

	req := c.newRequest(op, input, &AcceptPortfolioShareOutput{})

	return AcceptPortfolioShareRequest{Request: req, Input: input, Copy: c.AcceptPortfolioShareRequest}
}

// AcceptPortfolioShareRequest is the request type for the
// AcceptPortfolioShare API operation.
type AcceptPortfolioShareRequest struct {
	*aws.Request
	Input *AcceptPortfolioShareInput
	Copy  func(*AcceptPortfolioShareInput) AcceptPortfolioShareRequest
}

// Send marshals and sends the AcceptPortfolioShare API request.
func (r AcceptPortfolioShareRequest) Send(ctx context.Context) (*AcceptPortfolioShareResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcceptPortfolioShareResponse{
		AcceptPortfolioShareOutput: r.Request.Data.(*AcceptPortfolioShareOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcceptPortfolioShareResponse is the response type for the
// AcceptPortfolioShare API operation.
type AcceptPortfolioShareResponse struct {
	*AcceptPortfolioShareOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcceptPortfolioShare request.
func (r *AcceptPortfolioShareResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
