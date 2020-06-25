// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// The GetDomainDetail request includes the following element.
type GetDomainDetailInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain that you want to get detailed information about.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetDomainDetailInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDomainDetailInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDomainDetailInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The GetDomainDetail response includes the following elements.
type GetDomainDetailOutput struct {
	_ struct{} `type:"structure"`

	// Email address to contact to report incorrect contact information for a domain,
	// to report that the domain is being used to send spam, to report that someone
	// is cybersquatting on a domain name, or report some other type of abuse.
	AbuseContactEmail *string `type:"string"`

	// Phone number for reporting abuse.
	AbuseContactPhone *string `type:"string"`

	// Provides details about the domain administrative contact.
	//
	// AdminContact is a required field
	AdminContact *ContactDetail `type:"structure" required:"true" sensitive:"true"`

	// Specifies whether contact information is concealed from WHOIS queries. If
	// the value is true, WHOIS ("who is") queries return contact information either
	// for Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If the value is false, WHOIS queries
	// return the information that you entered for the admin contact.
	AdminPrivacy *bool `type:"boolean"`

	// Specifies whether the domain registration is set to renew automatically.
	AutoRenew *bool `type:"boolean"`

	// The date when the domain was created as found in the response to a WHOIS
	// query. The date and time is in Unix time format and Coordinated Universal
	// time (UTC).
	CreationDate *time.Time `type:"timestamp"`

	// Reserved for future use.
	DnsSec *string `type:"string"`

	// The name of a domain.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`

	// The date when the registration for the domain is set to expire. The date
	// and time is in Unix time format and Coordinated Universal time (UTC).
	ExpirationDate *time.Time `type:"timestamp"`

	// The name of the domain.
	//
	// Nameservers is a required field
	Nameservers []Nameserver `type:"list" required:"true"`

	// Provides details about the domain registrant.
	//
	// RegistrantContact is a required field
	RegistrantContact *ContactDetail `type:"structure" required:"true" sensitive:"true"`

	// Specifies whether contact information is concealed from WHOIS queries. If
	// the value is true, WHOIS ("who is") queries return contact information either
	// for Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If the value is false, WHOIS queries
	// return the information that you entered for the registrant contact (domain
	// owner).
	RegistrantPrivacy *bool `type:"boolean"`

	// Name of the registrar of the domain as identified in the registry. Domains
	// with a .com, .net, or .org TLD are registered by Amazon Registrar. All other
	// domains are registered by our registrar associate, Gandi. The value for domains
	// that are registered by Gandi is "GANDI SAS".
	RegistrarName *string `type:"string"`

	// Web address of the registrar.
	RegistrarUrl *string `type:"string"`

	// Reserved for future use.
	RegistryDomainId *string `type:"string"`

	// Reseller of the domain. Domains registered or transferred using Route 53
	// domains will have "Amazon" as the reseller.
	Reseller *string `type:"string"`

	// An array of domain name status codes, also known as Extensible Provisioning
	// Protocol (EPP) status codes.
	//
	// ICANN, the organization that maintains a central database of domain names,
	// has developed a set of domain name status codes that tell you the status
	// of a variety of operations on a domain name, for example, registering a domain
	// name, transferring a domain name to another registrar, renewing the registration
	// for a domain name, and so on. All registrars use this same set of status
	// codes.
	//
	// For a current list of domain name status codes and an explanation of what
	// each code means, go to the ICANN website (https://www.icann.org/) and search
	// for epp status codes. (Search on the ICANN website; web searches sometimes
	// return an old version of the document.)
	StatusList []string `type:"list"`

	// Provides details about the domain technical contact.
	//
	// TechContact is a required field
	TechContact *ContactDetail `type:"structure" required:"true" sensitive:"true"`

	// Specifies whether contact information is concealed from WHOIS queries. If
	// the value is true, WHOIS ("who is") queries return contact information either
	// for Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If the value is false, WHOIS queries
	// return the information that you entered for the technical contact.
	TechPrivacy *bool `type:"boolean"`

	// The last updated date of the domain as found in the response to a WHOIS query.
	// The date and time is in Unix time format and Coordinated Universal time (UTC).
	UpdatedDate *time.Time `type:"timestamp"`

	// The fully qualified name of the WHOIS server that can answer the WHOIS query
	// for the domain.
	WhoIsServer *string `type:"string"`
}

// String returns the string representation
func (s GetDomainDetailOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDomainDetail = "GetDomainDetail"

// GetDomainDetailRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// This operation returns detailed information about a specified domain that
// is associated with the current AWS account. Contact information for the domain
// is also returned as part of the output.
//
//    // Example sending a request using GetDomainDetailRequest.
//    req := client.GetDomainDetailRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/GetDomainDetail
func (c *Client) GetDomainDetailRequest(input *GetDomainDetailInput) GetDomainDetailRequest {
	op := &aws.Operation{
		Name:       opGetDomainDetail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDomainDetailInput{}
	}

	req := c.newRequest(op, input, &GetDomainDetailOutput{})

	return GetDomainDetailRequest{Request: req, Input: input, Copy: c.GetDomainDetailRequest}
}

// GetDomainDetailRequest is the request type for the
// GetDomainDetail API operation.
type GetDomainDetailRequest struct {
	*aws.Request
	Input *GetDomainDetailInput
	Copy  func(*GetDomainDetailInput) GetDomainDetailRequest
}

// Send marshals and sends the GetDomainDetail API request.
func (r GetDomainDetailRequest) Send(ctx context.Context) (*GetDomainDetailResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDomainDetailResponse{
		GetDomainDetailOutput: r.Request.Data.(*GetDomainDetailOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDomainDetailResponse is the response type for the
// GetDomainDetail API operation.
type GetDomainDetailResponse struct {
	*GetDomainDetailOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDomainDetail request.
func (r *GetDomainDetailResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
