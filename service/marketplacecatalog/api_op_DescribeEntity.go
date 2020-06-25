// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplacecatalog

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DescribeEntityInput struct {
	_ struct{} `type:"structure"`

	// Required. The catalog related to the request. Fixed value: AWSMarketplace
	//
	// Catalog is a required field
	Catalog *string `location:"querystring" locationName:"catalog" min:"1" type:"string" required:"true"`

	// Required. The unique ID of the entity to describe.
	//
	// EntityId is a required field
	EntityId *string `location:"querystring" locationName:"entityId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeEntityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEntityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEntityInput"}

	if s.Catalog == nil {
		invalidParams.Add(aws.NewErrParamRequired("Catalog"))
	}
	if s.Catalog != nil && len(*s.Catalog) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Catalog", 1))
	}

	if s.EntityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntityId"))
	}
	if s.EntityId != nil && len(*s.EntityId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeEntityInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Catalog != nil {
		v := *s.Catalog

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "catalog", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EntityId != nil {
		v := *s.EntityId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "entityId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeEntityOutput struct {
	_ struct{} `type:"structure"`

	// This stringified JSON object includes the details of the entity.
	Details *string `min:"2" type:"string"`

	// The ARN associated to the unique identifier for the change set referenced
	// in this request.
	EntityArn *string `min:"1" type:"string"`

	// The identifier of the entity, in the format of EntityId@RevisionId.
	EntityIdentifier *string `min:"1" type:"string"`

	// The named type of the entity, in the format of EntityType@Version.
	EntityType *string `min:"1" type:"string"`

	// The last modified date of the entity, in ISO 8601 format (2018-02-27T13:45:22Z).
	LastModifiedDate *string `type:"string"`
}

// String returns the string representation
func (s DescribeEntityOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeEntityOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Details != nil {
		v := *s.Details

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Details", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EntityArn != nil {
		v := *s.EntityArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EntityArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EntityIdentifier != nil {
		v := *s.EntityIdentifier

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EntityIdentifier", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EntityType != nil {
		v := *s.EntityType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EntityType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastModifiedDate != nil {
		v := *s.LastModifiedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastModifiedDate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeEntity = "DescribeEntity"

// DescribeEntityRequest returns a request value for making API operation for
// AWS Marketplace Catalog Service.
//
// Returns the metadata and content of the entity.
//
//    // Example sending a request using DescribeEntityRequest.
//    req := client.DescribeEntityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/marketplace-catalog-2018-09-17/DescribeEntity
func (c *Client) DescribeEntityRequest(input *DescribeEntityInput) DescribeEntityRequest {
	op := &aws.Operation{
		Name:       opDescribeEntity,
		HTTPMethod: "GET",
		HTTPPath:   "/DescribeEntity",
	}

	if input == nil {
		input = &DescribeEntityInput{}
	}

	req := c.newRequest(op, input, &DescribeEntityOutput{})

	return DescribeEntityRequest{Request: req, Input: input, Copy: c.DescribeEntityRequest}
}

// DescribeEntityRequest is the request type for the
// DescribeEntity API operation.
type DescribeEntityRequest struct {
	*aws.Request
	Input *DescribeEntityInput
	Copy  func(*DescribeEntityInput) DescribeEntityRequest
}

// Send marshals and sends the DescribeEntity API request.
func (r DescribeEntityRequest) Send(ctx context.Context) (*DescribeEntityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEntityResponse{
		DescribeEntityOutput: r.Request.Data.(*DescribeEntityOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEntityResponse is the response type for the
// DescribeEntity API operation.
type DescribeEntityResponse struct {
	*DescribeEntityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEntity request.
func (r *DescribeEntityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
