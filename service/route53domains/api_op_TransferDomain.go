// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// The TransferDomain request includes the following elements.
type TransferDomainInput struct {
	_ struct{} `type:"structure"`

	// Provides detailed contact information.
	//
	// AdminContact is a required field
	AdminContact *ContactDetail `type:"structure" required:"true" sensitive:"true"`

	// The authorization code for the domain. You get this value from the current
	// registrar.
	AuthCode *string `type:"string" sensitive:"true"`

	// Indicates whether the domain will be automatically renewed (true) or not
	// (false). Autorenewal only takes effect after the account is charged.
	//
	// Default: true
	AutoRenew *bool `type:"boolean"`

	// The name of the domain that you want to transfer to Route 53. The top-level
	// domain (TLD), such as .com, must be a TLD that Route 53 supports. For a list
	// of supported TLDs, see Domains that You Can Register with Amazon Route 53
	// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/registrar-tld-list.html)
	// in the Amazon Route 53 Developer Guide.
	//
	// The domain name can contain only the following characters:
	//
	//    * Letters a through z. Domain names are not case sensitive.
	//
	//    * Numbers 0 through 9.
	//
	//    * Hyphen (-). You can't specify a hyphen at the beginning or end of a
	//    label.
	//
	//    * Period (.) to separate the labels in the name, such as the . in example.com.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`

	// The number of years that you want to register the domain for. Domains are
	// registered for a minimum of one year. The maximum period depends on the top-level
	// domain.
	//
	// Default: 1
	//
	// DurationInYears is a required field
	DurationInYears *int64 `min:"1" type:"integer" required:"true"`

	// Reserved for future use.
	IdnLangCode *string `type:"string"`

	// Contains details for the host and glue IP addresses.
	Nameservers []Nameserver `type:"list"`

	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true, WHOIS ("who is") queries return contact information either
	// for Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false, WHOIS queries
	// return the information that you entered for the admin contact.
	//
	// Default: true
	PrivacyProtectAdminContact *bool `type:"boolean"`

	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true, WHOIS ("who is") queries return contact information either
	// for Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false, WHOIS queries
	// return the information that you entered for the registrant contact (domain
	// owner).
	//
	// Default: true
	PrivacyProtectRegistrantContact *bool `type:"boolean"`

	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true, WHOIS ("who is") queries return contact information either
	// for Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false, WHOIS queries
	// return the information that you entered for the technical contact.
	//
	// Default: true
	PrivacyProtectTechContact *bool `type:"boolean"`

	// Provides detailed contact information.
	//
	// RegistrantContact is a required field
	RegistrantContact *ContactDetail `type:"structure" required:"true" sensitive:"true"`

	// Provides detailed contact information.
	//
	// TechContact is a required field
	TechContact *ContactDetail `type:"structure" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s TransferDomainInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TransferDomainInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TransferDomainInput"}

	if s.AdminContact == nil {
		invalidParams.Add(aws.NewErrParamRequired("AdminContact"))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if s.DurationInYears == nil {
		invalidParams.Add(aws.NewErrParamRequired("DurationInYears"))
	}
	if s.DurationInYears != nil && *s.DurationInYears < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("DurationInYears", 1))
	}

	if s.RegistrantContact == nil {
		invalidParams.Add(aws.NewErrParamRequired("RegistrantContact"))
	}

	if s.TechContact == nil {
		invalidParams.Add(aws.NewErrParamRequired("TechContact"))
	}
	if s.AdminContact != nil {
		if err := s.AdminContact.Validate(); err != nil {
			invalidParams.AddNested("AdminContact", err.(aws.ErrInvalidParams))
		}
	}
	if s.Nameservers != nil {
		for i, v := range s.Nameservers {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Nameservers", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.RegistrantContact != nil {
		if err := s.RegistrantContact.Validate(); err != nil {
			invalidParams.AddNested("RegistrantContact", err.(aws.ErrInvalidParams))
		}
	}
	if s.TechContact != nil {
		if err := s.TechContact.Validate(); err != nil {
			invalidParams.AddNested("TechContact", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The TransferDomain response includes the following element.
type TransferDomainOutput struct {
	_ struct{} `type:"structure"`

	// Identifier for tracking the progress of the request. To query the operation
	// status, use GetOperationDetail (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html).
	//
	// OperationId is a required field
	OperationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s TransferDomainOutput) String() string {
	return awsutil.Prettify(s)
}

const opTransferDomain = "TransferDomain"

// TransferDomainRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// Transfers a domain from another registrar to Amazon Route 53. When the transfer
// is complete, the domain is registered either with Amazon Registrar (for .com,
// .net, and .org domains) or with our registrar associate, Gandi (for all other
// TLDs).
//
// For more information about transferring domains, see the following topics:
//
//    * For transfer requirements, a detailed procedure, and information about
//    viewing the status of a domain that you're transferring to Route 53, see
//    Transferring Registration for a Domain to Amazon Route 53 (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-transfer-to-route-53.html)
//    in the Amazon Route 53 Developer Guide.
//
//    * For information about how to transfer a domain from one AWS account
//    to another, see TransferDomainToAnotherAwsAccount (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html).
//
//    * For information about how to transfer a domain to another domain registrar,
//    see Transferring a Domain from Amazon Route 53 to Another Registrar (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-transfer-from-route-53.html)
//    in the Amazon Route 53 Developer Guide.
//
// If the registrar for your domain is also the DNS service provider for the
// domain, we highly recommend that you transfer your DNS service to Route 53
// or to another DNS service provider before you transfer your registration.
// Some registrars provide free DNS service when you purchase a domain registration.
// When you transfer the registration, the previous registrar will not renew
// your domain registration and could end your DNS service at any time.
//
// If the registrar for your domain is also the DNS service provider for the
// domain and you don't transfer DNS service to another provider, your website,
// email, and the web applications associated with the domain might become unavailable.
//
// If the transfer is successful, this method returns an operation ID that you
// can use to track the progress and completion of the action. If the transfer
// doesn't complete successfully, the domain registrant will be notified by
// email.
//
//    // Example sending a request using TransferDomainRequest.
//    req := client.TransferDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/TransferDomain
func (c *Client) TransferDomainRequest(input *TransferDomainInput) TransferDomainRequest {
	op := &aws.Operation{
		Name:       opTransferDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TransferDomainInput{}
	}

	req := c.newRequest(op, input, &TransferDomainOutput{})

	return TransferDomainRequest{Request: req, Input: input, Copy: c.TransferDomainRequest}
}

// TransferDomainRequest is the request type for the
// TransferDomain API operation.
type TransferDomainRequest struct {
	*aws.Request
	Input *TransferDomainInput
	Copy  func(*TransferDomainInput) TransferDomainRequest
}

// Send marshals and sends the TransferDomain API request.
func (r TransferDomainRequest) Send(ctx context.Context) (*TransferDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TransferDomainResponse{
		TransferDomainOutput: r.Request.Data.(*TransferDomainOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TransferDomainResponse is the response type for the
// TransferDomain API operation.
type TransferDomainResponse struct {
	*TransferDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TransferDomain request.
func (r *TransferDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
