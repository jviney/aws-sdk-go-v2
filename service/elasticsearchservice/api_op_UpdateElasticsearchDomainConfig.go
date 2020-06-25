// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// Container for the parameters to the UpdateElasticsearchDomain operation.
// Specifies the type and number of instances in the domain cluster.
type UpdateElasticsearchDomainConfigInput struct {
	_ struct{} `type:"structure"`

	// IAM access policy as a JSON-formatted string.
	AccessPolicies *string `type:"string"`

	// Modifies the advanced option to allow references to indices in an HTTP request
	// body. Must be false when configuring access to individual sub-resources.
	// By default, the value is true. See Configuration Advanced Options (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-advanced-options)
	// for more information.
	AdvancedOptions map[string]string `type:"map"`

	// Specifies advanced security options.
	AdvancedSecurityOptions *AdvancedSecurityOptionsInput `type:"structure"`

	// Options to specify the Cognito user and identity pools for Kibana authentication.
	// For more information, see Amazon Cognito Authentication for Kibana (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-cognito-auth.html).
	CognitoOptions *CognitoOptions `type:"structure"`

	// Options to specify configuration that will be applied to the domain endpoint.
	DomainEndpointOptions *DomainEndpointOptions `type:"structure"`

	// The name of the Elasticsearch domain that you are updating.
	//
	// DomainName is a required field
	DomainName *string `location:"uri" locationName:"DomainName" min:"3" type:"string" required:"true"`

	// Specify the type and size of the EBS volume that you want to use.
	EBSOptions *EBSOptions `type:"structure"`

	// The type and number of instances to instantiate for the domain cluster.
	ElasticsearchClusterConfig *ElasticsearchClusterConfig `type:"structure"`

	// Map of LogType and LogPublishingOption, each containing options to publish
	// a given type of Elasticsearch log.
	LogPublishingOptions map[string]LogPublishingOption `type:"map"`

	// Option to set the time, in UTC format, for the daily automated snapshot.
	// Default value is 0 hours.
	SnapshotOptions *SnapshotOptions `type:"structure"`

	// Options to specify the subnets and security groups for VPC endpoint. For
	// more information, see Creating a VPC (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-creating-vpc)
	// in VPC Endpoints for Amazon Elasticsearch Service Domains
	VPCOptions *VPCOptions `type:"structure"`
}

// String returns the string representation
func (s UpdateElasticsearchDomainConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateElasticsearchDomainConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateElasticsearchDomainConfigInput"}

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

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateElasticsearchDomainConfigInput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DomainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The result of an UpdateElasticsearchDomain request. Contains the status of
// the Elasticsearch domain being updated.
type UpdateElasticsearchDomainConfigOutput struct {
	_ struct{} `type:"structure"`

	// The status of the updated Elasticsearch domain.
	//
	// DomainConfig is a required field
	DomainConfig *ElasticsearchDomainConfig `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateElasticsearchDomainConfigOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateElasticsearchDomainConfigOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DomainConfig != nil {
		v := s.DomainConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DomainConfig", v, metadata)
	}
	return nil
}

const opUpdateElasticsearchDomainConfig = "UpdateElasticsearchDomainConfig"

// UpdateElasticsearchDomainConfigRequest returns a request value for making API operation for
// Amazon Elasticsearch Service.
//
// Modifies the cluster configuration of the specified Elasticsearch domain,
// setting as setting the instance type and the number of instances.
//
//    // Example sending a request using UpdateElasticsearchDomainConfigRequest.
//    req := client.UpdateElasticsearchDomainConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateElasticsearchDomainConfigRequest(input *UpdateElasticsearchDomainConfigInput) UpdateElasticsearchDomainConfigRequest {
	op := &aws.Operation{
		Name:       opUpdateElasticsearchDomainConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-01-01/es/domain/{DomainName}/config",
	}

	if input == nil {
		input = &UpdateElasticsearchDomainConfigInput{}
	}

	req := c.newRequest(op, input, &UpdateElasticsearchDomainConfigOutput{})

	return UpdateElasticsearchDomainConfigRequest{Request: req, Input: input, Copy: c.UpdateElasticsearchDomainConfigRequest}
}

// UpdateElasticsearchDomainConfigRequest is the request type for the
// UpdateElasticsearchDomainConfig API operation.
type UpdateElasticsearchDomainConfigRequest struct {
	*aws.Request
	Input *UpdateElasticsearchDomainConfigInput
	Copy  func(*UpdateElasticsearchDomainConfigInput) UpdateElasticsearchDomainConfigRequest
}

// Send marshals and sends the UpdateElasticsearchDomainConfig API request.
func (r UpdateElasticsearchDomainConfigRequest) Send(ctx context.Context) (*UpdateElasticsearchDomainConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateElasticsearchDomainConfigResponse{
		UpdateElasticsearchDomainConfigOutput: r.Request.Data.(*UpdateElasticsearchDomainConfigOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateElasticsearchDomainConfigResponse is the response type for the
// UpdateElasticsearchDomainConfig API operation.
type UpdateElasticsearchDomainConfigResponse struct {
	*UpdateElasticsearchDomainConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateElasticsearchDomainConfig request.
func (r *UpdateElasticsearchDomainConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
