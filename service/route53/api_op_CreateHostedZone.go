// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// A complex type that contains information about the request to create a public
// or private hosted zone.
type CreateHostedZoneInput struct {
	_ struct{} `locationName:"CreateHostedZoneRequest" type:"structure" xmlURI:"https://route53.amazonaws.com/doc/2013-04-01/"`

	// A unique string that identifies the request and that allows failed CreateHostedZone
	// requests to be retried without the risk of executing the operation twice.
	// You must use a unique CallerReference string every time you submit a CreateHostedZone
	// request. CallerReference can be any unique string, for example, a date/time
	// stamp.
	//
	// CallerReference is a required field
	CallerReference *string `min:"1" type:"string" required:"true"`

	// If you want to associate a reusable delegation set with this hosted zone,
	// the ID that Amazon Route 53 assigned to the reusable delegation set when
	// you created it. For more information about reusable delegation sets, see
	// CreateReusableDelegationSet (https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateReusableDelegationSet.html).
	DelegationSetId *string `type:"string"`

	// (Optional) A complex type that contains the following optional values:
	//
	//    * For public and private hosted zones, an optional comment
	//
	//    * For private hosted zones, an optional PrivateZone element
	//
	// If you don't specify a comment or the PrivateZone element, omit HostedZoneConfig
	// and the other elements.
	HostedZoneConfig *HostedZoneConfig `type:"structure"`

	// The name of the domain. Specify a fully qualified domain name, for example,
	// www.example.com. The trailing dot is optional; Amazon Route 53 assumes that
	// the domain name is fully qualified. This means that Route 53 treats www.example.com
	// (without a trailing dot) and www.example.com. (with a trailing dot) as identical.
	//
	// If you're creating a public hosted zone, this is the name you have registered
	// with your DNS registrar. If your domain name is registered with a registrar
	// other than Route 53, change the name servers for your domain to the set of
	// NameServers that CreateHostedZone returns in DelegationSet.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// (Private hosted zones only) A complex type that contains information about
	// the Amazon VPC that you're associating with this hosted zone.
	//
	// You can specify only one Amazon VPC when you create a private hosted zone.
	// To associate additional Amazon VPCs with the hosted zone, use AssociateVPCWithHostedZone
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_AssociateVPCWithHostedZone.html)
	// after you create a hosted zone.
	VPC *VPC `type:"structure"`
}

// String returns the string representation
func (s CreateHostedZoneInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateHostedZoneInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateHostedZoneInput"}

	if s.CallerReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("CallerReference"))
	}
	if s.CallerReference != nil && len(*s.CallerReference) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CallerReference", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.VPC != nil {
		if err := s.VPC.Validate(); err != nil {
			invalidParams.AddNested("VPC", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateHostedZoneInput) MarshalFields(e protocol.FieldEncoder) error {

	e.SetFields(protocol.BodyTarget, "CreateHostedZoneRequest", protocol.FieldMarshalerFunc(func(e protocol.FieldEncoder) error {
		if s.CallerReference != nil {
			v := *s.CallerReference

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "CallerReference", protocol.StringValue(v), metadata)
		}
		if s.DelegationSetId != nil {
			v := *s.DelegationSetId

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "DelegationSetId", protocol.StringValue(v), metadata)
		}
		if s.HostedZoneConfig != nil {
			v := s.HostedZoneConfig

			metadata := protocol.Metadata{}
			e.SetFields(protocol.BodyTarget, "HostedZoneConfig", v, metadata)
		}
		if s.Name != nil {
			v := *s.Name

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "Name", protocol.StringValue(v), metadata)
		}
		if s.VPC != nil {
			v := s.VPC

			metadata := protocol.Metadata{}
			e.SetFields(protocol.BodyTarget, "VPC", v, metadata)
		}
		return nil
	}), protocol.Metadata{XMLNamespaceURI: "https://route53.amazonaws.com/doc/2013-04-01/"})
	return nil
}

// A complex type containing the response information for the hosted zone.
type CreateHostedZoneOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains information about the CreateHostedZone request.
	//
	// ChangeInfo is a required field
	ChangeInfo *ChangeInfo `type:"structure" required:"true"`

	// A complex type that describes the name servers for this hosted zone.
	//
	// DelegationSet is a required field
	DelegationSet *DelegationSet `type:"structure" required:"true"`

	// A complex type that contains general information about the hosted zone.
	//
	// HostedZone is a required field
	HostedZone *HostedZone `type:"structure" required:"true"`

	// The unique URL representing the new hosted zone.
	//
	// Location is a required field
	Location *string `location:"header" locationName:"Location" type:"string" required:"true"`

	// A complex type that contains information about an Amazon VPC that you associated
	// with this hosted zone.
	VPC *VPC `type:"structure"`
}

// String returns the string representation
func (s CreateHostedZoneOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateHostedZoneOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ChangeInfo != nil {
		v := s.ChangeInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ChangeInfo", v, metadata)
	}
	if s.DelegationSet != nil {
		v := s.DelegationSet

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DelegationSet", v, metadata)
	}
	if s.HostedZone != nil {
		v := s.HostedZone

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "HostedZone", v, metadata)
	}
	if s.VPC != nil {
		v := s.VPC

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "VPC", v, metadata)
	}
	if s.Location != nil {
		v := *s.Location

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Location", protocol.StringValue(v), metadata)
	}
	return nil
}

const opCreateHostedZone = "CreateHostedZone"

// CreateHostedZoneRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Creates a new public or private hosted zone. You create records in a public
// hosted zone to define how you want to route traffic on the internet for a
// domain, such as example.com, and its subdomains (apex.example.com, acme.example.com).
// You create records in a private hosted zone to define how you want to route
// traffic for a domain and its subdomains within one or more Amazon Virtual
// Private Clouds (Amazon VPCs).
//
// You can't convert a public hosted zone to a private hosted zone or vice versa.
// Instead, you must create a new hosted zone with the same name and create
// new resource record sets.
//
// For more information about charges for hosted zones, see Amazon Route 53
// Pricing (http://aws.amazon.com/route53/pricing/).
//
// Note the following:
//
//    * You can't create a hosted zone for a top-level domain (TLD) such as
//    .com.
//
//    * For public hosted zones, Route 53 automatically creates a default SOA
//    record and four NS records for the zone. For more information about SOA
//    and NS records, see NS and SOA Records that Route 53 Creates for a Hosted
//    Zone (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/SOA-NSrecords.html)
//    in the Amazon Route 53 Developer Guide. If you want to use the same name
//    servers for multiple public hosted zones, you can optionally associate
//    a reusable delegation set with the hosted zone. See the DelegationSetId
//    element.
//
//    * If your domain is registered with a registrar other than Route 53, you
//    must update the name servers with your registrar to make Route 53 the
//    DNS service for the domain. For more information, see Migrating DNS Service
//    for an Existing Domain to Amazon Route 53 (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/MigratingDNS.html)
//    in the Amazon Route 53 Developer Guide.
//
// When you submit a CreateHostedZone request, the initial status of the hosted
// zone is PENDING. For public hosted zones, this means that the NS and SOA
// records are not yet available on all Route 53 DNS servers. When the NS and
// SOA records are available, the status of the zone changes to INSYNC.
//
//    // Example sending a request using CreateHostedZoneRequest.
//    req := client.CreateHostedZoneRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/CreateHostedZone
func (c *Client) CreateHostedZoneRequest(input *CreateHostedZoneInput) CreateHostedZoneRequest {
	op := &aws.Operation{
		Name:       opCreateHostedZone,
		HTTPMethod: "POST",
		HTTPPath:   "/2013-04-01/hostedzone",
	}

	if input == nil {
		input = &CreateHostedZoneInput{}
	}

	req := c.newRequest(op, input, &CreateHostedZoneOutput{})

	return CreateHostedZoneRequest{Request: req, Input: input, Copy: c.CreateHostedZoneRequest}
}

// CreateHostedZoneRequest is the request type for the
// CreateHostedZone API operation.
type CreateHostedZoneRequest struct {
	*aws.Request
	Input *CreateHostedZoneInput
	Copy  func(*CreateHostedZoneInput) CreateHostedZoneRequest
}

// Send marshals and sends the CreateHostedZone API request.
func (r CreateHostedZoneRequest) Send(ctx context.Context) (*CreateHostedZoneResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateHostedZoneResponse{
		CreateHostedZoneOutput: r.Request.Data.(*CreateHostedZoneOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateHostedZoneResponse is the response type for the
// CreateHostedZone API operation.
type CreateHostedZoneResponse struct {
	*CreateHostedZoneOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateHostedZone request.
func (r *CreateHostedZoneResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
