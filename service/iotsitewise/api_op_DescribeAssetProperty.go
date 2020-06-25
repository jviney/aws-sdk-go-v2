// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DescribeAssetPropertyInput struct {
	_ struct{} `type:"structure"`

	// The ID of the asset.
	//
	// AssetId is a required field
	AssetId *string `location:"uri" locationName:"assetId" min:"36" type:"string" required:"true"`

	// The ID of the asset property.
	//
	// PropertyId is a required field
	PropertyId *string `location:"uri" locationName:"propertyId" min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeAssetPropertyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAssetPropertyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAssetPropertyInput"}

	if s.AssetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssetId"))
	}
	if s.AssetId != nil && len(*s.AssetId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("AssetId", 36))
	}

	if s.PropertyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PropertyId"))
	}
	if s.PropertyId != nil && len(*s.PropertyId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("PropertyId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeAssetPropertyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AssetId != nil {
		v := *s.AssetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "assetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PropertyId != nil {
		v := *s.PropertyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "propertyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeAssetPropertyOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the asset.
	//
	// AssetId is a required field
	AssetId *string `locationName:"assetId" min:"36" type:"string" required:"true"`

	// The ID of the asset model.
	//
	// AssetModelId is a required field
	AssetModelId *string `locationName:"assetModelId" min:"36" type:"string" required:"true"`

	// The name of the asset.
	//
	// AssetName is a required field
	AssetName *string `locationName:"assetName" min:"1" type:"string" required:"true"`

	// The asset property's definition, alias, and notification state.
	//
	// AssetProperty is a required field
	AssetProperty *Property `locationName:"assetProperty" type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeAssetPropertyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeAssetPropertyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AssetId != nil {
		v := *s.AssetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "assetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AssetModelId != nil {
		v := *s.AssetModelId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "assetModelId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AssetName != nil {
		v := *s.AssetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "assetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AssetProperty != nil {
		v := s.AssetProperty

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "assetProperty", v, metadata)
	}
	return nil
}

const opDescribeAssetProperty = "DescribeAssetProperty"

// DescribeAssetPropertyRequest returns a request value for making API operation for
// AWS IoT SiteWise.
//
// Retrieves information about an asset's property.
//
//    // Example sending a request using DescribeAssetPropertyRequest.
//    req := client.DescribeAssetPropertyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotsitewise-2019-12-02/DescribeAssetProperty
func (c *Client) DescribeAssetPropertyRequest(input *DescribeAssetPropertyInput) DescribeAssetPropertyRequest {
	op := &aws.Operation{
		Name:       opDescribeAssetProperty,
		HTTPMethod: "GET",
		HTTPPath:   "/assets/{assetId}/properties/{propertyId}",
	}

	if input == nil {
		input = &DescribeAssetPropertyInput{}
	}

	req := c.newRequest(op, input, &DescribeAssetPropertyOutput{})
	req.Handlers.Build.PushBackNamed(protocol.NewHostPrefixHandler("model.", nil))
	req.Handlers.Build.PushBackNamed(protocol.ValidateEndpointHostHandler)

	return DescribeAssetPropertyRequest{Request: req, Input: input, Copy: c.DescribeAssetPropertyRequest}
}

// DescribeAssetPropertyRequest is the request type for the
// DescribeAssetProperty API operation.
type DescribeAssetPropertyRequest struct {
	*aws.Request
	Input *DescribeAssetPropertyInput
	Copy  func(*DescribeAssetPropertyInput) DescribeAssetPropertyRequest
}

// Send marshals and sends the DescribeAssetProperty API request.
func (r DescribeAssetPropertyRequest) Send(ctx context.Context) (*DescribeAssetPropertyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAssetPropertyResponse{
		DescribeAssetPropertyOutput: r.Request.Data.(*DescribeAssetPropertyOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAssetPropertyResponse is the response type for the
// DescribeAssetProperty API operation.
type DescribeAssetPropertyResponse struct {
	*DescribeAssetPropertyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAssetProperty request.
func (r *DescribeAssetPropertyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
