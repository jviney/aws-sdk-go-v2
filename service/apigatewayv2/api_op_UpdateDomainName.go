// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type UpdateDomainNameInput struct {
	_ struct{} `type:"structure"`

	// DomainName is a required field
	DomainName *string `location:"uri" locationName:"domainName" type:"string" required:"true"`

	// The domain name configurations.
	DomainNameConfigurations []DomainNameConfiguration `locationName:"domainNameConfigurations" type:"list"`
}

// String returns the string representation
func (s UpdateDomainNameInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDomainNameInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDomainNameInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDomainNameInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DomainNameConfigurations != nil {
		v := s.DomainNameConfigurations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "domainNameConfigurations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateDomainNameOutput struct {
	_ struct{} `type:"structure"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	ApiMappingSelectionExpression *string `locationName:"apiMappingSelectionExpression" type:"string"`

	// A string with a length between [1-512].
	DomainName *string `locationName:"domainName" type:"string"`

	// The domain name configurations.
	DomainNameConfigurations []DomainNameConfiguration `locationName:"domainNameConfigurations" type:"list"`

	// Represents a collection of tags associated with the resource.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s UpdateDomainNameOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDomainNameOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApiMappingSelectionExpression != nil {
		v := *s.ApiMappingSelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiMappingSelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainNameConfigurations != nil {
		v := s.DomainNameConfigurations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "domainNameConfigurations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

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
	return nil
}

const opUpdateDomainName = "UpdateDomainName"

// UpdateDomainNameRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Updates a domain name.
//
//    // Example sending a request using UpdateDomainNameRequest.
//    req := client.UpdateDomainNameRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/UpdateDomainName
func (c *Client) UpdateDomainNameRequest(input *UpdateDomainNameInput) UpdateDomainNameRequest {
	op := &aws.Operation{
		Name:       opUpdateDomainName,
		HTTPMethod: "PATCH",
		HTTPPath:   "/v2/domainnames/{domainName}",
	}

	if input == nil {
		input = &UpdateDomainNameInput{}
	}

	req := c.newRequest(op, input, &UpdateDomainNameOutput{})

	return UpdateDomainNameRequest{Request: req, Input: input, Copy: c.UpdateDomainNameRequest}
}

// UpdateDomainNameRequest is the request type for the
// UpdateDomainName API operation.
type UpdateDomainNameRequest struct {
	*aws.Request
	Input *UpdateDomainNameInput
	Copy  func(*UpdateDomainNameInput) UpdateDomainNameRequest
}

// Send marshals and sends the UpdateDomainName API request.
func (r UpdateDomainNameRequest) Send(ctx context.Context) (*UpdateDomainNameResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDomainNameResponse{
		UpdateDomainNameOutput: r.Request.Data.(*UpdateDomainNameOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDomainNameResponse is the response type for the
// UpdateDomainName API operation.
type UpdateDomainNameResponse struct {
	*UpdateDomainNameOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDomainName request.
func (r *UpdateDomainNameResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
