// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type UpgradeAppliedSchemaInput struct {
	_ struct{} `type:"structure"`

	// The ARN for the directory to which the upgraded schema will be applied.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `type:"string" required:"true"`

	// Used for testing whether the major version schemas are backward compatible
	// or not. If schema compatibility fails, an exception would be thrown else
	// the call would succeed but no changes will be saved. This parameter is optional.
	DryRun *bool `type:"boolean"`

	// The revision of the published schema to upgrade the directory to.
	//
	// PublishedSchemaArn is a required field
	PublishedSchemaArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpgradeAppliedSchemaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpgradeAppliedSchemaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpgradeAppliedSchemaInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if s.PublishedSchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PublishedSchemaArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpgradeAppliedSchemaInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DirectoryArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DryRun != nil {
		v := *s.DryRun

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DryRun", protocol.BoolValue(v), metadata)
	}
	if s.PublishedSchemaArn != nil {
		v := *s.PublishedSchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PublishedSchemaArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpgradeAppliedSchemaOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the directory that is returned as part of the response.
	DirectoryArn *string `type:"string"`

	// The ARN of the upgraded schema that is returned as part of the response.
	UpgradedSchemaArn *string `type:"string"`
}

// String returns the string representation
func (s UpgradeAppliedSchemaOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpgradeAppliedSchemaOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DirectoryArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UpgradedSchemaArn != nil {
		v := *s.UpgradedSchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "UpgradedSchemaArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpgradeAppliedSchema = "UpgradeAppliedSchema"

// UpgradeAppliedSchemaRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Upgrades a single directory in-place using the PublishedSchemaArn with schema
// updates found in MinorVersion. Backwards-compatible minor version upgrades
// are instantaneously available for readers on all objects in the directory.
// Note: This is a synchronous API call and upgrades only one schema on a given
// directory per call. To upgrade multiple directories from one schema, you
// would need to call this API on each directory.
//
//    // Example sending a request using UpgradeAppliedSchemaRequest.
//    req := client.UpgradeAppliedSchemaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/UpgradeAppliedSchema
func (c *Client) UpgradeAppliedSchemaRequest(input *UpgradeAppliedSchemaInput) UpgradeAppliedSchemaRequest {
	op := &aws.Operation{
		Name:       opUpgradeAppliedSchema,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/schema/upgradeapplied",
	}

	if input == nil {
		input = &UpgradeAppliedSchemaInput{}
	}

	req := c.newRequest(op, input, &UpgradeAppliedSchemaOutput{})

	return UpgradeAppliedSchemaRequest{Request: req, Input: input, Copy: c.UpgradeAppliedSchemaRequest}
}

// UpgradeAppliedSchemaRequest is the request type for the
// UpgradeAppliedSchema API operation.
type UpgradeAppliedSchemaRequest struct {
	*aws.Request
	Input *UpgradeAppliedSchemaInput
	Copy  func(*UpgradeAppliedSchemaInput) UpgradeAppliedSchemaRequest
}

// Send marshals and sends the UpgradeAppliedSchema API request.
func (r UpgradeAppliedSchemaRequest) Send(ctx context.Context) (*UpgradeAppliedSchemaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpgradeAppliedSchemaResponse{
		UpgradeAppliedSchemaOutput: r.Request.Data.(*UpgradeAppliedSchemaOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpgradeAppliedSchemaResponse is the response type for the
// UpgradeAppliedSchema API operation.
type UpgradeAppliedSchemaResponse struct {
	*UpgradeAppliedSchemaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpgradeAppliedSchema request.
func (r *UpgradeAppliedSchemaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
