// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dataexchange

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// The request to update an asset.
type UpdateAssetInput struct {
	_ struct{} `type:"structure"`

	// AssetId is a required field
	AssetId *string `location:"uri" locationName:"AssetId" type:"string" required:"true"`

	// DataSetId is a required field
	DataSetId *string `location:"uri" locationName:"DataSetId" type:"string" required:"true"`

	// The name of the asset. When importing from Amazon S3, the S3 object key is
	// used as the asset name. When exporting to Amazon S3, the asset name is used
	// as default target S3 object key.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// RevisionId is a required field
	RevisionId *string `location:"uri" locationName:"RevisionId" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateAssetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAssetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAssetInput"}

	if s.AssetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssetId"))
	}

	if s.DataSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSetId"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.RevisionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RevisionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateAssetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AssetId != nil {
		v := *s.AssetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AssetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DataSetId != nil {
		v := *s.DataSetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DataSetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "RevisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateAssetOutput struct {
	_ struct{} `type:"structure"`

	// An Amazon Resource Name (ARN) that uniquely identifies an AWS resource.
	Arn *string `type:"string"`

	AssetDetails *AssetDetails `type:"structure"`

	// The type of file your data is stored in. Currently, the supported asset type
	// is S3_SNAPSHOT.
	AssetType AssetType `type:"string" enum:"true"`

	// Dates and times in AWS Data Exchange are recorded in ISO 8601 format.
	CreatedAt *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// A unique identifier.
	DataSetId *string `type:"string"`

	// A unique identifier.
	Id *string `type:"string"`

	// The name of the asset. When importing from Amazon S3, the S3 object key is
	// used as the asset name. When exporting to Amazon S3, the asset name is used
	// as default target S3 object key.
	Name *string `type:"string"`

	// A unique identifier.
	RevisionId *string `type:"string"`

	// A unique identifier.
	SourceId *string `type:"string"`

	// Dates and times in AWS Data Exchange are recorded in ISO 8601 format.
	UpdatedAt *time.Time `type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s UpdateAssetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateAssetOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AssetDetails != nil {
		v := s.AssetDetails

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "AssetDetails", v, metadata)
	}
	if len(s.AssetType) > 0 {
		v := s.AssetType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AssetType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.CreatedAt != nil {
		v := *s.CreatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreatedAt",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.DataSetId != nil {
		v := *s.DataSetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DataSetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RevisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceId != nil {
		v := *s.SourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SourceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UpdatedAt != nil {
		v := *s.UpdatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "UpdatedAt",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	return nil
}

const opUpdateAsset = "UpdateAsset"

// UpdateAssetRequest returns a request value for making API operation for
// AWS Data Exchange.
//
// This operation updates an asset.
//
//    // Example sending a request using UpdateAssetRequest.
//    req := client.UpdateAssetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dataexchange-2017-07-25/UpdateAsset
func (c *Client) UpdateAssetRequest(input *UpdateAssetInput) UpdateAssetRequest {
	op := &aws.Operation{
		Name:       opUpdateAsset,
		HTTPMethod: "PATCH",
		HTTPPath:   "/v1/data-sets/{DataSetId}/revisions/{RevisionId}/assets/{AssetId}",
	}

	if input == nil {
		input = &UpdateAssetInput{}
	}

	req := c.newRequest(op, input, &UpdateAssetOutput{})

	return UpdateAssetRequest{Request: req, Input: input, Copy: c.UpdateAssetRequest}
}

// UpdateAssetRequest is the request type for the
// UpdateAsset API operation.
type UpdateAssetRequest struct {
	*aws.Request
	Input *UpdateAssetInput
	Copy  func(*UpdateAssetInput) UpdateAssetRequest
}

// Send marshals and sends the UpdateAsset API request.
func (r UpdateAssetRequest) Send(ctx context.Context) (*UpdateAssetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAssetResponse{
		UpdateAssetOutput: r.Request.Data.(*UpdateAssetOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAssetResponse is the response type for the
// UpdateAsset API operation.
type UpdateAssetResponse struct {
	*UpdateAssetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAsset request.
func (r *UpdateAssetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
