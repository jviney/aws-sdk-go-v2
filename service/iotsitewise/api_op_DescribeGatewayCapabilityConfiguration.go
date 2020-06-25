// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DescribeGatewayCapabilityConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The namespace of the capability configuration. For example, if you configure
	// OPC-UA sources from the AWS IoT SiteWise console, your OPC-UA capability
	// configuration has the namespace iotsitewise:opcuacollector:version, where
	// version is a number such as 1.
	//
	// CapabilityNamespace is a required field
	CapabilityNamespace *string `location:"uri" locationName:"capabilityNamespace" min:"1" type:"string" required:"true"`

	// The ID of the gateway that defines the capability configuration.
	//
	// GatewayId is a required field
	GatewayId *string `location:"uri" locationName:"gatewayId" min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeGatewayCapabilityConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeGatewayCapabilityConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeGatewayCapabilityConfigurationInput"}

	if s.CapabilityNamespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("CapabilityNamespace"))
	}
	if s.CapabilityNamespace != nil && len(*s.CapabilityNamespace) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CapabilityNamespace", 1))
	}

	if s.GatewayId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayId"))
	}
	if s.GatewayId != nil && len(*s.GatewayId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeGatewayCapabilityConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CapabilityNamespace != nil {
		v := *s.CapabilityNamespace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "capabilityNamespace", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GatewayId != nil {
		v := *s.GatewayId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "gatewayId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeGatewayCapabilityConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The JSON document that defines the gateway capability's configuration. For
	// more information, see Configuring data sources (CLI) (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/configure-sources.html#configure-source-cli)
	// in the AWS IoT SiteWise User Guide.
	//
	// CapabilityConfiguration is a required field
	CapabilityConfiguration *string `locationName:"capabilityConfiguration" min:"1" type:"string" required:"true"`

	// The namespace of the gateway capability.
	//
	// CapabilityNamespace is a required field
	CapabilityNamespace *string `locationName:"capabilityNamespace" min:"1" type:"string" required:"true"`

	// The synchronization status of the capability configuration. The sync status
	// can be one of the following:
	//
	//    * IN_SYNC – The gateway is running the capability configuration.
	//
	//    * OUT_OF_SYNC – The gateway hasn't received the capability configuration.
	//
	//    * SYNC_FAILED – The gateway rejected the capability configuration.
	//
	// CapabilitySyncStatus is a required field
	CapabilitySyncStatus CapabilitySyncStatus `locationName:"capabilitySyncStatus" type:"string" required:"true" enum:"true"`

	// The ID of the gateway that defines the capability configuration.
	//
	// GatewayId is a required field
	GatewayId *string `locationName:"gatewayId" min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeGatewayCapabilityConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeGatewayCapabilityConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CapabilityConfiguration != nil {
		v := *s.CapabilityConfiguration

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "capabilityConfiguration", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CapabilityNamespace != nil {
		v := *s.CapabilityNamespace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "capabilityNamespace", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.CapabilitySyncStatus) > 0 {
		v := s.CapabilitySyncStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "capabilitySyncStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.GatewayId != nil {
		v := *s.GatewayId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "gatewayId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeGatewayCapabilityConfiguration = "DescribeGatewayCapabilityConfiguration"

// DescribeGatewayCapabilityConfigurationRequest returns a request value for making API operation for
// AWS IoT SiteWise.
//
// Retrieves information about a gateway capability configuration. Each gateway
// capability defines data sources for a gateway. A capability configuration
// can contain multiple data source configurations. If you define OPC-UA sources
// for a gateway in the AWS IoT SiteWise console, all of your OPC-UA sources
// are stored in one capability configuration. To list all capability configurations
// for a gateway, use DescribeGateway (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeGateway.html).
//
//    // Example sending a request using DescribeGatewayCapabilityConfigurationRequest.
//    req := client.DescribeGatewayCapabilityConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotsitewise-2019-12-02/DescribeGatewayCapabilityConfiguration
func (c *Client) DescribeGatewayCapabilityConfigurationRequest(input *DescribeGatewayCapabilityConfigurationInput) DescribeGatewayCapabilityConfigurationRequest {
	op := &aws.Operation{
		Name:       opDescribeGatewayCapabilityConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/20200301/gateways/{gatewayId}/capability/{capabilityNamespace}",
	}

	if input == nil {
		input = &DescribeGatewayCapabilityConfigurationInput{}
	}

	req := c.newRequest(op, input, &DescribeGatewayCapabilityConfigurationOutput{})
	req.Handlers.Build.PushBackNamed(protocol.NewHostPrefixHandler("edge.", nil))
	req.Handlers.Build.PushBackNamed(protocol.ValidateEndpointHostHandler)

	return DescribeGatewayCapabilityConfigurationRequest{Request: req, Input: input, Copy: c.DescribeGatewayCapabilityConfigurationRequest}
}

// DescribeGatewayCapabilityConfigurationRequest is the request type for the
// DescribeGatewayCapabilityConfiguration API operation.
type DescribeGatewayCapabilityConfigurationRequest struct {
	*aws.Request
	Input *DescribeGatewayCapabilityConfigurationInput
	Copy  func(*DescribeGatewayCapabilityConfigurationInput) DescribeGatewayCapabilityConfigurationRequest
}

// Send marshals and sends the DescribeGatewayCapabilityConfiguration API request.
func (r DescribeGatewayCapabilityConfigurationRequest) Send(ctx context.Context) (*DescribeGatewayCapabilityConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeGatewayCapabilityConfigurationResponse{
		DescribeGatewayCapabilityConfigurationOutput: r.Request.Data.(*DescribeGatewayCapabilityConfigurationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeGatewayCapabilityConfigurationResponse is the response type for the
// DescribeGatewayCapabilityConfiguration API operation.
type DescribeGatewayCapabilityConfigurationResponse struct {
	*DescribeGatewayCapabilityConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeGatewayCapabilityConfiguration request.
func (r *DescribeGatewayCapabilityConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
