// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acm

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type ResendValidationEmailInput struct {
	_ struct{} `type:"structure"`

	// String that contains the ARN of the requested certificate. The certificate
	// ARN is generated and returned by the RequestCertificate action as soon as
	// the request is made. By default, using this parameter causes email to be
	// sent to all top-level domains you specified in the certificate request. The
	// ARN must be of the form:
	//
	// arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012
	//
	// CertificateArn is a required field
	CertificateArn *string `min:"20" type:"string" required:"true"`

	// The fully qualified domain name (FQDN) of the certificate that needs to be
	// validated.
	//
	// Domain is a required field
	Domain *string `min:"1" type:"string" required:"true"`

	// The base validation domain that will act as the suffix of the email addresses
	// that are used to send the emails. This must be the same as the Domain value
	// or a superdomain of the Domain value. For example, if you requested a certificate
	// for site.subdomain.example.com and specify a ValidationDomain of subdomain.example.com,
	// ACM sends email to the domain registrant, technical contact, and administrative
	// contact in WHOIS and the following five addresses:
	//
	//    * admin@subdomain.example.com
	//
	//    * administrator@subdomain.example.com
	//
	//    * hostmaster@subdomain.example.com
	//
	//    * postmaster@subdomain.example.com
	//
	//    * webmaster@subdomain.example.com
	//
	// ValidationDomain is a required field
	ValidationDomain *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ResendValidationEmailInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResendValidationEmailInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResendValidationEmailInput"}

	if s.CertificateArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateArn"))
	}
	if s.CertificateArn != nil && len(*s.CertificateArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateArn", 20))
	}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 1))
	}

	if s.ValidationDomain == nil {
		invalidParams.Add(aws.NewErrParamRequired("ValidationDomain"))
	}
	if s.ValidationDomain != nil && len(*s.ValidationDomain) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ValidationDomain", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ResendValidationEmailOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ResendValidationEmailOutput) String() string {
	return awsutil.Prettify(s)
}

const opResendValidationEmail = "ResendValidationEmail"

// ResendValidationEmailRequest returns a request value for making API operation for
// AWS Certificate Manager.
//
// Resends the email that requests domain ownership validation. The domain owner
// or an authorized representative must approve the ACM certificate before it
// can be issued. The certificate can be approved by clicking a link in the
// mail to navigate to the Amazon certificate approval website and then clicking
// I Approve. However, the validation email can be blocked by spam filters.
// Therefore, if you do not receive the original mail, you can request that
// the mail be resent within 72 hours of requesting the ACM certificate. If
// more than 72 hours have elapsed since your original request or since your
// last attempt to resend validation mail, you must request a new certificate.
// For more information about setting up your contact email addresses, see Configure
// Email for your Domain (https://docs.aws.amazon.com/acm/latest/userguide/setup-email.html).
//
//    // Example sending a request using ResendValidationEmailRequest.
//    req := client.ResendValidationEmailRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-2015-12-08/ResendValidationEmail
func (c *Client) ResendValidationEmailRequest(input *ResendValidationEmailInput) ResendValidationEmailRequest {
	op := &aws.Operation{
		Name:       opResendValidationEmail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResendValidationEmailInput{}
	}

	req := c.newRequest(op, input, &ResendValidationEmailOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return ResendValidationEmailRequest{Request: req, Input: input, Copy: c.ResendValidationEmailRequest}
}

// ResendValidationEmailRequest is the request type for the
// ResendValidationEmail API operation.
type ResendValidationEmailRequest struct {
	*aws.Request
	Input *ResendValidationEmailInput
	Copy  func(*ResendValidationEmailInput) ResendValidationEmailRequest
}

// Send marshals and sends the ResendValidationEmail API request.
func (r ResendValidationEmailRequest) Send(ctx context.Context) (*ResendValidationEmailResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResendValidationEmailResponse{
		ResendValidationEmailOutput: r.Request.Data.(*ResendValidationEmailOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResendValidationEmailResponse is the response type for the
// ResendValidationEmail API operation.
type ResendValidationEmailResponse struct {
	*ResendValidationEmailOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResendValidationEmail request.
func (r *ResendValidationEmailResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
