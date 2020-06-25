// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type CreateElasticsearchDomainInput struct {
	_ struct{} `type:"structure"`

	// IAM access policy as a JSON-formatted string.
	AccessPolicies *string `type:"string"`

	// Option to allow references to indices in an HTTP request body. Must be false
	// when configuring access to individual sub-resources. By default, the value
	// is true. See Configuration Advanced Options (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-advanced-options)
	// for more information.
	AdvancedOptions map[string]string `type:"map"`

	// Specifies advanced security options.
	AdvancedSecurityOptions *AdvancedSecurityOptionsInput `type:"structure"`

	// Options to specify the Cognito user and identity pools for Kibana authentication.
	// For more information, see Amazon Cognito Authentication for Kibana (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-cognito-auth.html).
	CognitoOptions *CognitoOptions `type:"structure"`

	// Options to specify configuration that will be applied to the domain endpoint.
	DomainEndpointOptions *DomainEndpointOptions `type:"structure"`

	// The name of the Elasticsearch domain that you are creating. Domain names
	// are unique across the domains owned by an account within an AWS region. Domain
	// names must start with a lowercase letter and can contain the following characters:
	// a-z (lowercase), 0-9, and - (hyphen).
	//
	// DomainName is a required field
	DomainName *string `min:"3" type:"string" required:"true"`

	// Options to enable, disable and specify the type and size of EBS storage volumes.
	EBSOptions *EBSOptions `type:"structure"`

	// Configuration options for an Elasticsearch domain. Specifies the instance
	// type and number of instances in the domain cluster.
	ElasticsearchClusterConfig *ElasticsearchClusterConfig `type:"structure"`

	// String of format X.Y to specify version for the Elasticsearch domain eg.
	// "1.5" or "2.3". For more information, see Creating Elasticsearch Domains
	// (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomains)
	// in the Amazon Elasticsearch Service Developer Guide.
	ElasticsearchVersion *string `type:"string"`

	// Specifies the Encryption At Rest Options.
	EncryptionAtRestOptions *EncryptionAtRestOptions `type:"structure"`

	// Map of LogType and LogPublishingOption, each containing options to publish
	// a given type of Elasticsearch log.
	LogPublishingOptions map[string]LogPublishingOption `type:"map"`

	// Specifies the NodeToNodeEncryptionOptions.
	NodeToNodeEncryptionOptions *NodeToNodeEncryptionOptions `type:"structure"`

	// Option to set time, in UTC format, of the daily automated snapshot. Default
	// value is 0 hours.
	SnapshotOptions *SnapshotOptions `type:"structure"`

	// Options to specify the subnets and security groups for VPC endpoint. For
	// more information, see Creating a VPC (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-creating-vpc)
	// in VPC Endpoints for Amazon Elasticsearch Service Domains
	VPCOptions *VPCOptions `type:"structure"`
}

// String returns the string representation
func (s CreateElasticsearchDomainInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateElasticsearchDomainInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateElasticsearchDomainInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 3))
	}
	if s.AdvancedSecurityOptions != nil {
		if err := s.AdvancedSecurityOptions.Validate(); err != nil {
			invalidParams.AddNested("AdvancedSecurityOptions", err.(aws.ErrInvalidParams))
		}
	}
	if s.CognitoOptions != nil {
		if err := s.CognitoOptions.Validate(); err != nil {
			invalidParams.AddNested("CognitoOptions", err.(aws.ErrInvalidParams))
		}
	}
	if s.EncryptionAtRestOptions != nil {
		if err := s.EncryptionAtRestOptions.Validate(); err != nil {
			invalidParams.AddNested("EncryptionAtRestOptions", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateElasticsearchDomainInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccessPolicies != nil {
		v := *s.AccessPolicies

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AccessPolicies", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AdvancedOptions != nil {
		v := s.AdvancedOptions

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "AdvancedOptions", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.AdvancedSecurityOptions != nil {
		v := s.AdvancedSecurityOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "AdvancedSecurityOptions", v, metadata)
	}
	if s.CognitoOptions != nil {
		v := s.CognitoOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "CognitoOptions", v, metadata)
	}
	if s.DomainEndpointOptions != nil {
		v := s.DomainEndpointOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DomainEndpointOptions", v, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DomainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EBSOptions != nil {
		v := s.EBSOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "EBSOptions", v, metadata)
	}
	if s.ElasticsearchClusterConfig != nil {
		v := s.ElasticsearchClusterConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ElasticsearchClusterConfig", v, metadata)
	}
	if s.ElasticsearchVersion != nil {
		v := *s.ElasticsearchVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ElasticsearchVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EncryptionAtRestOptions != nil {
		v := s.EncryptionAtRestOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "EncryptionAtRestOptions", v, metadata)
	}
	if s.LogPublishingOptions != nil {
		v := s.LogPublishingOptions

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "LogPublishingOptions", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.NodeToNodeEncryptionOptions != nil {
		v := s.NodeToNodeEncryptionOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "NodeToNodeEncryptionOptions", v, metadata)
	}
	if s.SnapshotOptions != nil {
		v := s.SnapshotOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SnapshotOptions", v, metadata)
	}
	if s.VPCOptions != nil {
		v := s.VPCOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "VPCOptions", v, metadata)
	}
	return nil
}

// The result of a CreateElasticsearchDomain operation. Contains the status
// of the newly created Elasticsearch domain.
type CreateElasticsearchDomainOutput struct {
	_ struct{} `type:"structure"`

	// The status of the newly created Elasticsearch domain.
	DomainStatus *ElasticsearchDomainStatus `type:"structure"`
}

// String returns the string representation
func (s CreateElasticsearchDomainOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateElasticsearchDomainOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DomainStatus != nil {
		v := s.DomainStatus

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DomainStatus", v, metadata)
	}
	return nil
}

const opCreateElasticsearchDomain = "CreateElasticsearchDomain"

// CreateElasticsearchDomainRequest returns a request value for making API operation for
// Amazon Elasticsearch Service.
//
// Creates a new Elasticsearch domain. For more information, see Creating Elasticsearch
// Domains (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomains)
// in the Amazon Elasticsearch Service Developer Guide.
//
//    // Example sending a request using CreateElasticsearchDomainRequest.
//    req := client.CreateElasticsearchDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateElasticsearchDomainRequest(input *CreateElasticsearchDomainInput) CreateElasticsearchDomainRequest {
	op := &aws.Operation{
		Name:       opCreateElasticsearchDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-01-01/es/domain",
	}

	if input == nil {
		input = &CreateElasticsearchDomainInput{}
	}

	req := c.newRequest(op, input, &CreateElasticsearchDomainOutput{})

	return CreateElasticsearchDomainRequest{Request: req, Input: input, Copy: c.CreateElasticsearchDomainRequest}
}

// CreateElasticsearchDomainRequest is the request type for the
// CreateElasticsearchDomain API operation.
type CreateElasticsearchDomainRequest struct {
	*aws.Request
	Input *CreateElasticsearchDomainInput
	Copy  func(*CreateElasticsearchDomainInput) CreateElasticsearchDomainRequest
}

// Send marshals and sends the CreateElasticsearchDomain API request.
func (r CreateElasticsearchDomainRequest) Send(ctx context.Context) (*CreateElasticsearchDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateElasticsearchDomainResponse{
		CreateElasticsearchDomainOutput: r.Request.Data.(*CreateElasticsearchDomainOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateElasticsearchDomainResponse is the response type for the
// CreateElasticsearchDomain API operation.
type CreateElasticsearchDomainResponse struct {
	*CreateElasticsearchDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateElasticsearchDomain request.
func (r *CreateElasticsearchDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
