// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteCorsConfigurationInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCorsConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCorsConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCorsConfigurationInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteCorsConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteCorsConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteCorsConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteCorsConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteCorsConfiguration = "DeleteCorsConfiguration"

// DeleteCorsConfigurationRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Deletes a CORS configuration.
//
//    // Example sending a request using DeleteCorsConfigurationRequest.
//    req := client.DeleteCorsConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/DeleteCorsConfiguration
func (c *Client) DeleteCorsConfigurationRequest(input *DeleteCorsConfigurationInput) DeleteCorsConfigurationRequest {
	op := &aws.Operation{
		Name:       opDeleteCorsConfiguration,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v2/apis/{apiId}/cors",
	}

	if input == nil {
		input = &DeleteCorsConfigurationInput{}
	}

	req := c.newRequest(op, input, &DeleteCorsConfigurationOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteCorsConfigurationRequest{Request: req, Input: input, Copy: c.DeleteCorsConfigurationRequest}
}

// DeleteCorsConfigurationRequest is the request type for the
// DeleteCorsConfiguration API operation.
type DeleteCorsConfigurationRequest struct {
	*aws.Request
	Input *DeleteCorsConfigurationInput
	Copy  func(*DeleteCorsConfigurationInput) DeleteCorsConfigurationRequest
}

// Send marshals and sends the DeleteCorsConfiguration API request.
func (r DeleteCorsConfigurationRequest) Send(ctx context.Context) (*DeleteCorsConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCorsConfigurationResponse{
		DeleteCorsConfigurationOutput: r.Request.Data.(*DeleteCorsConfigurationOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCorsConfigurationResponse is the response type for the
// DeleteCorsConfiguration API operation.
type DeleteCorsConfigurationResponse struct {
	*DeleteCorsConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCorsConfiguration request.
func (r *DeleteCorsConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
