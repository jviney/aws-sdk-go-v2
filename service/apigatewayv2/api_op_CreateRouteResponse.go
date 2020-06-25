// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type CreateRouteResponseInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	ModelSelectionExpression *string `locationName:"modelSelectionExpression" type:"string"`

	// The route models.
	ResponseModels map[string]string `locationName:"responseModels" type:"map"`

	// The route parameters.
	ResponseParameters map[string]ParameterConstraints `locationName:"responseParameters" type:"map"`

	// RouteId is a required field
	RouteId *string `location:"uri" locationName:"routeId" type:"string" required:"true"`

	// After evaluating a selection expression, the result is compared against one
	// or more selection keys to find a matching key. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for a list of expressions and each expression's associated selection key
	// type.
	//
	// RouteResponseKey is a required field
	RouteResponseKey *string `locationName:"routeResponseKey" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateRouteResponseInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRouteResponseInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRouteResponseInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.RouteId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteId"))
	}

	if s.RouteResponseKey == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteResponseKey"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateRouteResponseInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ModelSelectionExpression != nil {
		v := *s.ModelSelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "modelSelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResponseModels != nil {
		v := s.ResponseModels

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "responseModels", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.ResponseParameters != nil {
		v := s.ResponseParameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "responseParameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.RouteResponseKey != nil {
		v := *s.RouteResponseKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "routeResponseKey", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RouteId != nil {
		v := *s.RouteId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "routeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateRouteResponseOutput struct {
	_ struct{} `type:"structure"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	ModelSelectionExpression *string `locationName:"modelSelectionExpression" type:"string"`

	// The route models.
	ResponseModels map[string]string `locationName:"responseModels" type:"map"`

	// The route parameters.
	ResponseParameters map[string]ParameterConstraints `locationName:"responseParameters" type:"map"`

	// The identifier.
	RouteResponseId *string `locationName:"routeResponseId" type:"string"`

	// After evaluating a selection expression, the result is compared against one
	// or more selection keys to find a matching key. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for a list of expressions and each expression's associated selection key
	// type.
	RouteResponseKey *string `locationName:"routeResponseKey" type:"string"`
}

// String returns the string representation
func (s CreateRouteResponseOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateRouteResponseOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ModelSelectionExpression != nil {
		v := *s.ModelSelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "modelSelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResponseModels != nil {
		v := s.ResponseModels

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "responseModels", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.ResponseParameters != nil {
		v := s.ResponseParameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "responseParameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.RouteResponseId != nil {
		v := *s.RouteResponseId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "routeResponseId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RouteResponseKey != nil {
		v := *s.RouteResponseKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "routeResponseKey", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateRouteResponse = "CreateRouteResponse"

// CreateRouteResponseRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Creates a RouteResponse for a Route.
//
//    // Example sending a request using CreateRouteResponseRequest.
//    req := client.CreateRouteResponseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/CreateRouteResponse
func (c *Client) CreateRouteResponseRequest(input *CreateRouteResponseInput) CreateRouteResponseRequest {
	op := &aws.Operation{
		Name:       opCreateRouteResponse,
		HTTPMethod: "POST",
		HTTPPath:   "/v2/apis/{apiId}/routes/{routeId}/routeresponses",
	}

	if input == nil {
		input = &CreateRouteResponseInput{}
	}

	req := c.newRequest(op, input, &CreateRouteResponseOutput{})

	return CreateRouteResponseRequest{Request: req, Input: input, Copy: c.CreateRouteResponseRequest}
}

// CreateRouteResponseRequest is the request type for the
// CreateRouteResponse API operation.
type CreateRouteResponseRequest struct {
	*aws.Request
	Input *CreateRouteResponseInput
	Copy  func(*CreateRouteResponseInput) CreateRouteResponseRequest
}

// Send marshals and sends the CreateRouteResponse API request.
func (r CreateRouteResponseRequest) Send(ctx context.Context) (*CreateRouteResponseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRouteResponseResponse{
		CreateRouteResponseOutput: r.Request.Data.(*CreateRouteResponseOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRouteResponseResponse is the response type for the
// CreateRouteResponse API operation.
type CreateRouteResponseResponse struct {
	*CreateRouteResponseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRouteResponse request.
func (r *CreateRouteResponseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
