// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type CreateGatewayInput struct {
	_ struct{} `type:"structure"`

	// A unique, friendly name for the gateway.
	//
	// GatewayName is a required field
	GatewayName *string `locationName:"gatewayName" min:"1" type:"string" required:"true"`

	// The gateway's platform. You can only specify one platform in a gateway.
	//
	// GatewayPlatform is a required field
	GatewayPlatform *GatewayPlatform `locationName:"gatewayPlatform" type:"structure" required:"true"`

	// A list of key-value pairs that contain metadata for the gateway. For more
	// information, see Tagging your AWS IoT SiteWise resources (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html)
	// in the AWS IoT SiteWise User Guide.
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`
}

// String returns the string representation
func (s CreateGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateGatewayInput"}

	if s.GatewayName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayName"))
	}
	if s.GatewayName != nil && len(*s.GatewayName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayName", 1))
	}

	if s.GatewayPlatform == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayPlatform"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.GatewayPlatform != nil {
		if err := s.GatewayPlatform.Validate(); err != nil {
			invalidParams.AddNested("GatewayPlatform", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateGatewayInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GatewayName != nil {
		v := *s.GatewayName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "gatewayName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GatewayPlatform != nil {
		v := s.GatewayPlatform

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "gatewayPlatform", v, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

type CreateGatewayOutput struct {
	_ struct{} `type:"structure"`

	// The ARN (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the gateway, which has the following format.
	//
	// arn:${Partition}:iotsitewise:${Region}:${Account}:gateway/${GatewayId}
	//
	// GatewayArn is a required field
	GatewayArn *string `locationName:"gatewayArn" min:"1" type:"string" required:"true"`

	// The ID of the gateway device. You can use this ID when you call other AWS
	// IoT SiteWise APIs.
	//
	// GatewayId is a required field
	GatewayId *string `locationName:"gatewayId" min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateGatewayOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.GatewayArn != nil {
		v := *s.GatewayArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "gatewayArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GatewayId != nil {
		v := *s.GatewayId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "gatewayId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateGateway = "CreateGateway"

// CreateGatewayRequest returns a request value for making API operation for
// AWS IoT SiteWise.
//
// Creates a gateway, which is a virtual or edge device that delivers industrial
// data streams from local servers to AWS IoT SiteWise. For more information,
// see Ingesting data using a gateway (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/gateway-connector.html)
// in the AWS IoT SiteWise User Guide.
//
//    // Example sending a request using CreateGatewayRequest.
//    req := client.CreateGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotsitewise-2019-12-02/CreateGateway
func (c *Client) CreateGatewayRequest(input *CreateGatewayInput) CreateGatewayRequest {
	op := &aws.Operation{
		Name:       opCreateGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/20200301/gateways",
	}

	if input == nil {
		input = &CreateGatewayInput{}
	}

	req := c.newRequest(op, input, &CreateGatewayOutput{})
	req.Handlers.Build.PushBackNamed(protocol.NewHostPrefixHandler("edge.", nil))
	req.Handlers.Build.PushBackNamed(protocol.ValidateEndpointHostHandler)

	return CreateGatewayRequest{Request: req, Input: input, Copy: c.CreateGatewayRequest}
}

// CreateGatewayRequest is the request type for the
// CreateGateway API operation.
type CreateGatewayRequest struct {
	*aws.Request
	Input *CreateGatewayInput
	Copy  func(*CreateGatewayInput) CreateGatewayRequest
}

// Send marshals and sends the CreateGateway API request.
func (r CreateGatewayRequest) Send(ctx context.Context) (*CreateGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateGatewayResponse{
		CreateGatewayOutput: r.Request.Data.(*CreateGatewayOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateGatewayResponse is the response type for the
// CreateGateway API operation.
type CreateGatewayResponse struct {
	*CreateGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateGateway request.
func (r *CreateGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
