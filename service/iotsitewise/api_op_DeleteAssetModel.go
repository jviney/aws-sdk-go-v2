// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DeleteAssetModelInput struct {
	_ struct{} `type:"structure"`

	// The ID of the asset model to delete.
	//
	// AssetModelId is a required field
	AssetModelId *string `location:"uri" locationName:"assetModelId" min:"36" type:"string" required:"true"`

	// A unique case-sensitive identifier that you can provide to ensure the idempotency
	// of the request. Don't reuse this client token if a new idempotent request
	// is required.
	ClientToken *string `location:"querystring" locationName:"clientToken" min:"36" type:"string" idempotencyToken:"true"`
}

// String returns the string representation
func (s DeleteAssetModelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAssetModelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAssetModelInput"}

	if s.AssetModelId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssetModelId"))
	}
	if s.AssetModelId != nil && len(*s.AssetModelId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("AssetModelId", 36))
	}
	if s.ClientToken != nil && len(*s.ClientToken) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAssetModelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AssetModelId != nil {
		v := *s.AssetModelId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "assetModelId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	var ClientToken string
	if s.ClientToken != nil {
		ClientToken = *s.ClientToken
	} else {
		ClientToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteAssetModelOutput struct {
	_ struct{} `type:"structure"`

	// The status of the asset model, which contains a state (DELETING after successfully
	// calling this operation) and any error message.
	//
	// AssetModelStatus is a required field
	AssetModelStatus *AssetModelStatus `locationName:"assetModelStatus" type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteAssetModelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAssetModelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AssetModelStatus != nil {
		v := s.AssetModelStatus

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "assetModelStatus", v, metadata)
	}
	return nil
}

const opDeleteAssetModel = "DeleteAssetModel"

// DeleteAssetModelRequest returns a request value for making API operation for
// AWS IoT SiteWise.
//
// Deletes an asset model. This action can't be undone. You must delete all
// assets created from an asset model before you can delete the model. Also,
// you can't delete an asset model if a parent asset model exists that contains
// a property formula expression that depends on the asset model that you want
// to delete. For more information, see Deleting Assets and Models (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/delete-assets-and-models.html)
// in the AWS IoT SiteWise User Guide.
//
//    // Example sending a request using DeleteAssetModelRequest.
//    req := client.DeleteAssetModelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotsitewise-2019-12-02/DeleteAssetModel
func (c *Client) DeleteAssetModelRequest(input *DeleteAssetModelInput) DeleteAssetModelRequest {
	op := &aws.Operation{
		Name:       opDeleteAssetModel,
		HTTPMethod: "DELETE",
		HTTPPath:   "/asset-models/{assetModelId}",
	}

	if input == nil {
		input = &DeleteAssetModelInput{}
	}

	req := c.newRequest(op, input, &DeleteAssetModelOutput{})
	req.Handlers.Build.PushBackNamed(protocol.NewHostPrefixHandler("model.", nil))
	req.Handlers.Build.PushBackNamed(protocol.ValidateEndpointHostHandler)

	return DeleteAssetModelRequest{Request: req, Input: input, Copy: c.DeleteAssetModelRequest}
}

// DeleteAssetModelRequest is the request type for the
// DeleteAssetModel API operation.
type DeleteAssetModelRequest struct {
	*aws.Request
	Input *DeleteAssetModelInput
	Copy  func(*DeleteAssetModelInput) DeleteAssetModelRequest
}

// Send marshals and sends the DeleteAssetModel API request.
func (r DeleteAssetModelRequest) Send(ctx context.Context) (*DeleteAssetModelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAssetModelResponse{
		DeleteAssetModelOutput: r.Request.Data.(*DeleteAssetModelOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAssetModelResponse is the response type for the
// DeleteAssetModel API operation.
type DeleteAssetModelResponse struct {
	*DeleteAssetModelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAssetModel request.
func (r *DeleteAssetModelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
