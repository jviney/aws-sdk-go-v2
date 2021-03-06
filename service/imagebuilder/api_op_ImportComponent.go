// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ImportComponentInput struct {
	_ struct{} `type:"structure"`

	// The change description of the component. Describes what change has been made
	// in this version, or what makes this version different from other versions
	// of this component.
	ChangeDescription *string `locationName:"changeDescription" min:"1" type:"string"`

	// The idempotency token of the component.
	//
	// ClientToken is a required field
	ClientToken *string `locationName:"clientToken" min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The data of the component. Used to specify the data inline. Either data or
	// uri can be used to specify the data within the component.
	Data *string `locationName:"data" min:"1" type:"string"`

	// The description of the component. Describes the contents of the component.
	Description *string `locationName:"description" min:"1" type:"string"`

	// The format of the resource that you want to import as a component.
	//
	// Format is a required field
	Format ComponentFormat `locationName:"format" type:"string" required:"true" enum:"true"`

	// The ID of the KMS key that should be used to encrypt this component.
	KmsKeyId *string `locationName:"kmsKeyId" min:"1" type:"string"`

	// The name of the component.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// The platform of the component.
	//
	// Platform is a required field
	Platform Platform `locationName:"platform" type:"string" required:"true" enum:"true"`

	// The semantic version of the component. This version follows the semantic
	// version syntax. For example, major.minor.patch. This could be versioned like
	// software (2.0.1) or like a date (2019.12.01).
	//
	// SemanticVersion is a required field
	SemanticVersion *string `locationName:"semanticVersion" type:"string" required:"true"`

	// The tags of the component.
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`

	// The type of the component denotes whether the component is used to build
	// the image or only to test it.
	//
	// Type is a required field
	Type ComponentType `locationName:"type" type:"string" required:"true" enum:"true"`

	// The uri of the component. Must be an S3 URL and the requester must have permission
	// to access the S3 bucket. If you use S3, you can specify component content
	// up to your service quota. Either data or uri can be used to specify the data
	// within the component.
	Uri *string `locationName:"uri" type:"string"`
}

// String returns the string representation
func (s ImportComponentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ImportComponentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ImportComponentInput"}
	if s.ChangeDescription != nil && len(*s.ChangeDescription) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeDescription", 1))
	}

	if s.ClientToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientToken"))
	}
	if s.ClientToken != nil && len(*s.ClientToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 1))
	}
	if s.Data != nil && len(*s.Data) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Data", 1))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if len(s.Format) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Format"))
	}
	if s.KmsKeyId != nil && len(*s.KmsKeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KmsKeyId", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if len(s.Platform) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Platform"))
	}

	if s.SemanticVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("SemanticVersion"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ImportComponentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ChangeDescription != nil {
		v := *s.ChangeDescription

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "changeDescription", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Data != nil {
		v := *s.Data

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "data", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Format) > 0 {
		v := s.Format

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "format", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.KmsKeyId != nil {
		v := *s.KmsKeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "kmsKeyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Platform) > 0 {
		v := s.Platform

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "platform", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.SemanticVersion != nil {
		v := *s.SemanticVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "semanticVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Uri != nil {
		v := *s.Uri

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "uri", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ImportComponentOutput struct {
	_ struct{} `type:"structure"`

	// The idempotency token used to make this request idempotent.
	ClientToken *string `locationName:"clientToken" min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the imported component.
	ComponentBuildVersionArn *string `locationName:"componentBuildVersionArn" type:"string"`

	// The request ID that uniquely identifies this request.
	RequestId *string `locationName:"requestId" min:"1" type:"string"`
}

// String returns the string representation
func (s ImportComponentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ImportComponentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ComponentBuildVersionArn != nil {
		v := *s.ComponentBuildVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "componentBuildVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opImportComponent = "ImportComponent"

// ImportComponentRequest returns a request value for making API operation for
// EC2 Image Builder.
//
// Imports a component and transforms its data into a component document.
//
//    // Example sending a request using ImportComponentRequest.
//    req := client.ImportComponentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/imagebuilder-2019-12-02/ImportComponent
func (c *Client) ImportComponentRequest(input *ImportComponentInput) ImportComponentRequest {
	op := &aws.Operation{
		Name:       opImportComponent,
		HTTPMethod: "PUT",
		HTTPPath:   "/ImportComponent",
	}

	if input == nil {
		input = &ImportComponentInput{}
	}

	req := c.newRequest(op, input, &ImportComponentOutput{})

	return ImportComponentRequest{Request: req, Input: input, Copy: c.ImportComponentRequest}
}

// ImportComponentRequest is the request type for the
// ImportComponent API operation.
type ImportComponentRequest struct {
	*aws.Request
	Input *ImportComponentInput
	Copy  func(*ImportComponentInput) ImportComponentRequest
}

// Send marshals and sends the ImportComponent API request.
func (r ImportComponentRequest) Send(ctx context.Context) (*ImportComponentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportComponentResponse{
		ImportComponentOutput: r.Request.Data.(*ImportComponentOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportComponentResponse is the response type for the
// ImportComponent API operation.
type ImportComponentResponse struct {
	*ImportComponentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportComponent request.
func (r *ImportComponentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
