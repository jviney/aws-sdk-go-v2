// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// A request to change information about the BasePathMapping resource.
type UpdateBasePathMappingInput struct {
	_ struct{} `type:"structure"`

	// [Required] The base path of the BasePathMapping resource to change.
	//
	// To specify an empty base path, set this parameter to '(none)'.
	//
	// BasePath is a required field
	BasePath *string `location:"uri" locationName:"base_path" type:"string" required:"true"`

	// [Required] The domain name of the BasePathMapping resource to change.
	//
	// DomainName is a required field
	DomainName *string `location:"uri" locationName:"domain_name" type:"string" required:"true"`

	// A list of update operations to be applied to the specified resource and in
	// the order specified in this list.
	PatchOperations []PatchOperation `locationName:"patchOperations" type:"list"`
}

// String returns the string representation
func (s UpdateBasePathMappingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateBasePathMappingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateBasePathMappingInput"}

	if s.BasePath == nil {
		invalidParams.Add(aws.NewErrParamRequired("BasePath"))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateBasePathMappingInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.PatchOperations != nil {
		v := s.PatchOperations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "patchOperations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.BasePath != nil {
		v := *s.BasePath

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "base_path", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "domain_name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents the base path that callers of the API must provide as part of
// the URL after the domain name.
//
// A custom domain name plus a BasePathMapping specification identifies a deployed
// RestApi in a given stage of the owner Account.
//
// Use Custom Domain Names (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html)
type UpdateBasePathMappingOutput struct {
	_ struct{} `type:"structure"`

	// The base path name that callers of the API must provide as part of the URL
	// after the domain name.
	BasePath *string `locationName:"basePath" type:"string"`

	// The string identifier of the associated RestApi.
	RestApiId *string `locationName:"restApiId" type:"string"`

	// The name of the associated stage.
	Stage *string `locationName:"stage" type:"string"`
}

// String returns the string representation
func (s UpdateBasePathMappingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateBasePathMappingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BasePath != nil {
		v := *s.BasePath

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "basePath", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "restApiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Stage != nil {
		v := *s.Stage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateBasePathMapping = "UpdateBasePathMapping"

// UpdateBasePathMappingRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Changes information about the BasePathMapping resource.
//
//    // Example sending a request using UpdateBasePathMappingRequest.
//    req := client.UpdateBasePathMappingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateBasePathMappingRequest(input *UpdateBasePathMappingInput) UpdateBasePathMappingRequest {
	op := &aws.Operation{
		Name:       opUpdateBasePathMapping,
		HTTPMethod: "PATCH",
		HTTPPath:   "/domainnames/{domain_name}/basepathmappings/{base_path}",
	}

	if input == nil {
		input = &UpdateBasePathMappingInput{}
	}

	req := c.newRequest(op, input, &UpdateBasePathMappingOutput{})

	return UpdateBasePathMappingRequest{Request: req, Input: input, Copy: c.UpdateBasePathMappingRequest}
}

// UpdateBasePathMappingRequest is the request type for the
// UpdateBasePathMapping API operation.
type UpdateBasePathMappingRequest struct {
	*aws.Request
	Input *UpdateBasePathMappingInput
	Copy  func(*UpdateBasePathMappingInput) UpdateBasePathMappingRequest
}

// Send marshals and sends the UpdateBasePathMapping API request.
func (r UpdateBasePathMappingRequest) Send(ctx context.Context) (*UpdateBasePathMappingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateBasePathMappingResponse{
		UpdateBasePathMappingOutput: r.Request.Data.(*UpdateBasePathMappingOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateBasePathMappingResponse is the response type for the
// UpdateBasePathMapping API operation.
type UpdateBasePathMappingResponse struct {
	*UpdateBasePathMappingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateBasePathMapping request.
func (r *UpdateBasePathMappingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
