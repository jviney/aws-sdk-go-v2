// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dataexchange

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// A request to create a data set that contains one or more revisions.
type CreateDataSetInput struct {
	_ struct{} `type:"structure"`

	// The type of file your data is stored in. Currently, the supported asset type
	// is S3_SNAPSHOT.
	//
	// AssetType is a required field
	AssetType AssetType `type:"string" required:"true" enum:"true"`

	// A description for the data set. This value can be up to 16,348 characters
	// long.
	//
	// Description is a required field
	Description *string `type:"string" required:"true"`

	// The name of the data set.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// A data set tag is an optional label that you can assign to a data set when
	// you create it. Each tag consists of a key and an optional value, both of
	// which you define. When you use tagging, you can also use tag-based access
	// control in IAM policies to control access to these data sets and revisions.
	Tags map[string]string `type:"map"`
}

// String returns the string representation
func (s CreateDataSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDataSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDataSetInput"}
	if len(s.AssetType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("AssetType"))
	}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDataSetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if len(s.AssetType) > 0 {
		v := s.AssetType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AssetType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

type CreateDataSetOutput struct {
	_ struct{} `type:"structure"`

	// An Amazon Resource Name (ARN) that uniquely identifies an AWS resource.
	Arn *string `type:"string"`

	// The type of file your data is stored in. Currently, the supported asset type
	// is S3_SNAPSHOT.
	AssetType AssetType `type:"string" enum:"true"`

	// Dates and times in AWS Data Exchange are recorded in ISO 8601 format.
	CreatedAt *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// A description of a resource.
	Description *string `type:"string"`

	// A unique identifier.
	Id *string `type:"string"`

	// The name of the model.
	Name *string `type:"string"`

	// A property that defines the data set as OWNED by the account (for providers)
	// or ENTITLED to the account (for subscribers). When an owned data set is published
	// in a product, AWS Data Exchange creates a copy of the data set. Subscribers
	// can access that copy of the data set as an entitled data set.
	Origin Origin `type:"string" enum:"true"`

	OriginDetails *OriginDetails `type:"structure"`

	// A unique identifier.
	SourceId *string `type:"string"`

	Tags map[string]string `type:"map"`

	// Dates and times in AWS Data Exchange are recorded in ISO 8601 format.
	UpdatedAt *time.Time `type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s CreateDataSetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDataSetOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if len(s.Origin) > 0 {
		v := s.Origin

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Origin", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.OriginDetails != nil {
		v := s.OriginDetails

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "OriginDetails", v, metadata)
	}
	if s.SourceId != nil {
		v := *s.SourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SourceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.UpdatedAt != nil {
		v := *s.UpdatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "UpdatedAt",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	return nil
}

const opCreateDataSet = "CreateDataSet"

// CreateDataSetRequest returns a request value for making API operation for
// AWS Data Exchange.
//
// This operation creates a data set.
//
//    // Example sending a request using CreateDataSetRequest.
//    req := client.CreateDataSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dataexchange-2017-07-25/CreateDataSet
func (c *Client) CreateDataSetRequest(input *CreateDataSetInput) CreateDataSetRequest {
	op := &aws.Operation{
		Name:       opCreateDataSet,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/data-sets",
	}

	if input == nil {
		input = &CreateDataSetInput{}
	}

	req := c.newRequest(op, input, &CreateDataSetOutput{})

	return CreateDataSetRequest{Request: req, Input: input, Copy: c.CreateDataSetRequest}
}

// CreateDataSetRequest is the request type for the
// CreateDataSet API operation.
type CreateDataSetRequest struct {
	*aws.Request
	Input *CreateDataSetInput
	Copy  func(*CreateDataSetInput) CreateDataSetRequest
}

// Send marshals and sends the CreateDataSet API request.
func (r CreateDataSetRequest) Send(ctx context.Context) (*CreateDataSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDataSetResponse{
		CreateDataSetOutput: r.Request.Data.(*CreateDataSetOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDataSetResponse is the response type for the
// CreateDataSet API operation.
type CreateDataSetResponse struct {
	*CreateDataSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDataSet request.
func (r *CreateDataSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
