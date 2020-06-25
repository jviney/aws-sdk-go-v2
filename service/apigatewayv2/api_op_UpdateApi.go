// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type UpdateApiInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	ApiKeySelectionExpression *string `locationName:"apiKeySelectionExpression" type:"string"`

	// Represents a CORS configuration. Supported only for HTTP APIs. See Configuring
	// CORS (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html)
	// for more information.
	CorsConfiguration *Cors `locationName:"corsConfiguration" type:"structure"`

	// Represents an Amazon Resource Name (ARN).
	CredentialsArn *string `locationName:"credentialsArn" type:"string"`

	// A string with a length between [0-1024].
	Description *string `locationName:"description" type:"string"`

	DisableSchemaValidation *bool `locationName:"disableSchemaValidation" type:"boolean"`

	// A string with a length between [1-128].
	Name *string `locationName:"name" type:"string"`

	// After evaluating a selection expression, the result is compared against one
	// or more selection keys to find a matching key. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for a list of expressions and each expression's associated selection key
	// type.
	RouteKey *string `locationName:"routeKey" type:"string"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	RouteSelectionExpression *string `locationName:"routeSelectionExpression" type:"string"`

	// A string representation of a URI with a length between [1-2048].
	Target *string `locationName:"target" type:"string"`

	// A string with a length between [1-64].
	Version *string `locationName:"version" type:"string"`
}

// String returns the string representation
func (s UpdateApiInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateApiInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateApiInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}
	if s.CorsConfiguration != nil {
		if err := s.CorsConfiguration.Validate(); err != nil {
			invalidParams.AddNested("CorsConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateApiInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiKeySelectionExpression != nil {
		v := *s.ApiKeySelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiKeySelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CorsConfiguration != nil {
		v := s.CorsConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "corsConfiguration", v, metadata)
	}
	if s.CredentialsArn != nil {
		v := *s.CredentialsArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "credentialsArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DisableSchemaValidation != nil {
		v := *s.DisableSchemaValidation

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "disableSchemaValidation", protocol.BoolValue(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RouteKey != nil {
		v := *s.RouteKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "routeKey", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RouteSelectionExpression != nil {
		v := *s.RouteSelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "routeSelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Target != nil {
		v := *s.Target

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "target", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateApiOutput struct {
	_ struct{} `type:"structure"`

	ApiEndpoint *string `locationName:"apiEndpoint" type:"string"`

	// The identifier.
	ApiId *string `locationName:"apiId" type:"string"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	ApiKeySelectionExpression *string `locationName:"apiKeySelectionExpression" type:"string"`

	// Represents a CORS configuration. Supported only for HTTP APIs. See Configuring
	// CORS (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html)
	// for more information.
	CorsConfiguration *Cors `locationName:"corsConfiguration" type:"structure"`

	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp" timestampFormat:"iso8601"`

	// A string with a length between [0-1024].
	Description *string `locationName:"description" type:"string"`

	DisableSchemaValidation *bool `locationName:"disableSchemaValidation" type:"boolean"`

	ImportInfo []string `locationName:"importInfo" type:"list"`

	// A string with a length between [1-128].
	Name *string `locationName:"name" type:"string"`

	// Represents a protocol type.
	ProtocolType ProtocolType `locationName:"protocolType" type:"string" enum:"true"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	RouteSelectionExpression *string `locationName:"routeSelectionExpression" type:"string"`

	// Represents a collection of tags associated with the resource.
	Tags map[string]string `locationName:"tags" type:"map"`

	// A string with a length between [1-64].
	Version *string `locationName:"version" type:"string"`

	Warnings []string `locationName:"warnings" type:"list"`
}

// String returns the string representation
func (s UpdateApiOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateApiOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApiEndpoint != nil {
		v := *s.ApiEndpoint

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiEndpoint", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiKeySelectionExpression != nil {
		v := *s.ApiKeySelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiKeySelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CorsConfiguration != nil {
		v := s.CorsConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "corsConfiguration", v, metadata)
	}
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DisableSchemaValidation != nil {
		v := *s.DisableSchemaValidation

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "disableSchemaValidation", protocol.BoolValue(v), metadata)
	}
	if s.ImportInfo != nil {
		v := s.ImportInfo

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "importInfo", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ProtocolType) > 0 {
		v := s.ProtocolType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "protocolType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.RouteSelectionExpression != nil {
		v := *s.RouteSelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "routeSelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Warnings != nil {
		v := s.Warnings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "warnings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opUpdateApi = "UpdateApi"

// UpdateApiRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Updates an Api resource.
//
//    // Example sending a request using UpdateApiRequest.
//    req := client.UpdateApiRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/UpdateApi
func (c *Client) UpdateApiRequest(input *UpdateApiInput) UpdateApiRequest {
	op := &aws.Operation{
		Name:       opUpdateApi,
		HTTPMethod: "PATCH",
		HTTPPath:   "/v2/apis/{apiId}",
	}

	if input == nil {
		input = &UpdateApiInput{}
	}

	req := c.newRequest(op, input, &UpdateApiOutput{})

	return UpdateApiRequest{Request: req, Input: input, Copy: c.UpdateApiRequest}
}

// UpdateApiRequest is the request type for the
// UpdateApi API operation.
type UpdateApiRequest struct {
	*aws.Request
	Input *UpdateApiInput
	Copy  func(*UpdateApiInput) UpdateApiRequest
}

// Send marshals and sends the UpdateApi API request.
func (r UpdateApiRequest) Send(ctx context.Context) (*UpdateApiResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateApiResponse{
		UpdateApiOutput: r.Request.Data.(*UpdateApiOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateApiResponse is the response type for the
// UpdateApi API operation.
type UpdateApiResponse struct {
	*UpdateApiOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApi request.
func (r *UpdateApiResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
