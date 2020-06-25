// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type BatchPutAssetPropertyValueInput struct {
	_ struct{} `type:"structure"`

	// The list of asset property value entries for the batch put request. You can
	// specify up to 10 entries per request.
	//
	// Entries is a required field
	Entries []PutAssetPropertyValueEntry `locationName:"entries" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchPutAssetPropertyValueInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchPutAssetPropertyValueInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchPutAssetPropertyValueInput"}

	if s.Entries == nil {
		invalidParams.Add(aws.NewErrParamRequired("Entries"))
	}
	if s.Entries != nil {
		for i, v := range s.Entries {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Entries", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchPutAssetPropertyValueInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Entries != nil {
		v := s.Entries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "entries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

type BatchPutAssetPropertyValueOutput struct {
	_ struct{} `type:"structure"`

	// A list of the errors (if any) associated with the batch put request. Each
	// error entry contains the entryId of the entry that failed.
	//
	// ErrorEntries is a required field
	ErrorEntries []BatchPutAssetPropertyErrorEntry `locationName:"errorEntries" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchPutAssetPropertyValueOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchPutAssetPropertyValueOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ErrorEntries != nil {
		v := s.ErrorEntries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "errorEntries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opBatchPutAssetPropertyValue = "BatchPutAssetPropertyValue"

// BatchPutAssetPropertyValueRequest returns a request value for making API operation for
// AWS IoT SiteWise.
//
// Sends a list of asset property values to AWS IoT SiteWise. Each value is
// a timestamp-quality-value (TQV) data point. For more information, see Ingesting
// Data Using the API (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/ingest-api.html)
// in the AWS IoT SiteWise User Guide.
//
// To identify an asset property, you must specify one of the following:
//
//    * The assetId and propertyId of an asset property.
//
//    * A propertyAlias, which is a data stream alias (for example, /company/windfarm/3/turbine/7/temperature).
//    To define an asset property's alias, see UpdateAssetProperty (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetProperty.html).
//
// With respect to Unix epoch time, AWS IoT SiteWise accepts only TQVs that
// have a timestamp of no more than 15 minutes in the past and no more than
// 5 minutes in the future. AWS IoT SiteWise rejects timestamps outside of the
// inclusive range of [-15, +5] minutes and returns a TimestampOutOfRangeException
// error.
//
// For each asset property, AWS IoT SiteWise overwrites TQVs with duplicate
// timestamps unless the newer TQV has a different quality. For example, if
// you store a TQV {T1, GOOD, V1}, then storing {T1, GOOD, V2} replaces the
// existing TQV.
//
//    // Example sending a request using BatchPutAssetPropertyValueRequest.
//    req := client.BatchPutAssetPropertyValueRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotsitewise-2019-12-02/BatchPutAssetPropertyValue
func (c *Client) BatchPutAssetPropertyValueRequest(input *BatchPutAssetPropertyValueInput) BatchPutAssetPropertyValueRequest {
	op := &aws.Operation{
		Name:       opBatchPutAssetPropertyValue,
		HTTPMethod: "POST",
		HTTPPath:   "/properties",
	}

	if input == nil {
		input = &BatchPutAssetPropertyValueInput{}
	}

	req := c.newRequest(op, input, &BatchPutAssetPropertyValueOutput{})
	req.Handlers.Build.PushBackNamed(protocol.NewHostPrefixHandler("data.", nil))
	req.Handlers.Build.PushBackNamed(protocol.ValidateEndpointHostHandler)

	return BatchPutAssetPropertyValueRequest{Request: req, Input: input, Copy: c.BatchPutAssetPropertyValueRequest}
}

// BatchPutAssetPropertyValueRequest is the request type for the
// BatchPutAssetPropertyValue API operation.
type BatchPutAssetPropertyValueRequest struct {
	*aws.Request
	Input *BatchPutAssetPropertyValueInput
	Copy  func(*BatchPutAssetPropertyValueInput) BatchPutAssetPropertyValueRequest
}

// Send marshals and sends the BatchPutAssetPropertyValue API request.
func (r BatchPutAssetPropertyValueRequest) Send(ctx context.Context) (*BatchPutAssetPropertyValueResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchPutAssetPropertyValueResponse{
		BatchPutAssetPropertyValueOutput: r.Request.Data.(*BatchPutAssetPropertyValueOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchPutAssetPropertyValueResponse is the response type for the
// BatchPutAssetPropertyValue API operation.
type BatchPutAssetPropertyValueResponse struct {
	*BatchPutAssetPropertyValueOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchPutAssetPropertyValue request.
func (r *BatchPutAssetPropertyValueResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
