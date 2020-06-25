// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to begin Amazon SES domain verification and to generate
// the TXT records that you must publish to the DNS server of your domain to
// complete the verification. For information about domain verification, see
// the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-domains.html).
type VerifyDomainIdentityInput struct {
	_ struct{} `type:"structure"`

	// The domain to be verified.
	//
	// Domain is a required field
	Domain *string `type:"string" required:"true"`
}

// String returns the string representation
func (s VerifyDomainIdentityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *VerifyDomainIdentityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "VerifyDomainIdentityInput"}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Returns a TXT record that you must publish to the DNS server of your domain
// to complete domain verification with Amazon SES.
type VerifyDomainIdentityOutput struct {
	_ struct{} `type:"structure"`

	// A TXT record that you must place in the DNS settings of the domain to complete
	// domain verification with Amazon SES.
	//
	// As Amazon SES searches for the TXT record, the domain's verification status
	// is "Pending". When Amazon SES detects the record, the domain's verification
	// status changes to "Success". If Amazon SES is unable to detect the record
	// within 72 hours, the domain's verification status changes to "Failed." In
	// that case, if you still want to verify the domain, you must restart the verification
	// process from the beginning.
	//
	// VerificationToken is a required field
	VerificationToken *string `type:"string" required:"true"`
}

// String returns the string representation
func (s VerifyDomainIdentityOutput) String() string {
	return awsutil.Prettify(s)
}

const opVerifyDomainIdentity = "VerifyDomainIdentity"

// VerifyDomainIdentityRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Adds a domain to the list of identities for your Amazon SES account in the
// current AWS Region and attempts to verify it. For more information about
// verifying domains, see Verifying Email Addresses and Domains (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-addresses-and-domains.html)
// in the Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using VerifyDomainIdentityRequest.
//    req := client.VerifyDomainIdentityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/VerifyDomainIdentity
func (c *Client) VerifyDomainIdentityRequest(input *VerifyDomainIdentityInput) VerifyDomainIdentityRequest {
	op := &aws.Operation{
		Name:       opVerifyDomainIdentity,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &VerifyDomainIdentityInput{}
	}

	req := c.newRequest(op, input, &VerifyDomainIdentityOutput{})

	return VerifyDomainIdentityRequest{Request: req, Input: input, Copy: c.VerifyDomainIdentityRequest}
}

// VerifyDomainIdentityRequest is the request type for the
// VerifyDomainIdentity API operation.
type VerifyDomainIdentityRequest struct {
	*aws.Request
	Input *VerifyDomainIdentityInput
	Copy  func(*VerifyDomainIdentityInput) VerifyDomainIdentityRequest
}

// Send marshals and sends the VerifyDomainIdentity API request.
func (r VerifyDomainIdentityRequest) Send(ctx context.Context) (*VerifyDomainIdentityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &VerifyDomainIdentityResponse{
		VerifyDomainIdentityOutput: r.Request.Data.(*VerifyDomainIdentityOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// VerifyDomainIdentityResponse is the response type for the
// VerifyDomainIdentity API operation.
type VerifyDomainIdentityResponse struct {
	*VerifyDomainIdentityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// VerifyDomainIdentity request.
func (r *VerifyDomainIdentityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
