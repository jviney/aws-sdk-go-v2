// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetApiMappingInput struct {
	_ struct{} `type:"structure"`

	// ApiMappingId is a required field
	ApiMappingId *string `location:"uri" locationName:"apiMappingId" type:"string" required:"true"`

	// DomainName is a required field
	DomainName *string `location:"uri" locationName:"domainName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetApiMappingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetApiMappingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetApiMappingInput"}

	if s.ApiMappingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiMappingId"))
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
func (s GetApiMappingInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiMappingId != nil {
		v := *s.ApiMappingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiMappingId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetApiMappingOutput struct {
	_ struct{} `type:"structure"`

	// The identifier.
	ApiId *string `locationName:"apiId" type:"string"`

	// The identifier.
	ApiMappingId *string `locationName:"apiMappingId" type:"string"`

	// After evaluating a selection expression, the result is compared against one
	// or more selection keys to find a matching key. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for a list of expressions and each expression's associated selection key
	// type.
	ApiMappingKey *string `locationName:"apiMappingKey" type:"string"`

	// A string with a length between [1-128].
	Stage *string `locationName:"stage" type:"string"`
}

// String returns the string representation
func (s GetApiMappingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetApiMappingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiMappingId != nil {
		v := *s.ApiMappingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiMappingId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiMappingKey != nil {
		v := *s.ApiMappingKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiMappingKey", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Stage != nil {
		v := *s.Stage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetApiMapping = "GetApiMapping"

// GetApiMappingRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Gets an API mapping.
//
//    // Example sending a request using GetApiMappingRequest.
//    req := client.GetApiMappingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetApiMapping
func (c *Client) GetApiMappingRequest(input *GetApiMappingInput) GetApiMappingRequest {
	op := &aws.Operation{
		Name:       opGetApiMapping,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/domainnames/{domainName}/apimappings/{apiMappingId}",
	}

	if input == nil {
		input = &GetApiMappingInput{}
	}

	req := c.newRequest(op, input, &GetApiMappingOutput{})

	return GetApiMappingRequest{Request: req, Input: input, Copy: c.GetApiMappingRequest}
}

// GetApiMappingRequest is the request type for the
// GetApiMapping API operation.
type GetApiMappingRequest struct {
	*aws.Request
	Input *GetApiMappingInput
	Copy  func(*GetApiMappingInput) GetApiMappingRequest
}

// Send marshals and sends the GetApiMapping API request.
func (r GetApiMappingRequest) Send(ctx context.Context) (*GetApiMappingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetApiMappingResponse{
		GetApiMappingOutput: r.Request.Data.(*GetApiMappingOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetApiMappingResponse is the response type for the
// GetApiMapping API operation.
type GetApiMappingResponse struct {
	*GetApiMappingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetApiMapping request.
func (r *GetApiMappingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
