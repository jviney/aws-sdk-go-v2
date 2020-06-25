// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteRouteInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// RouteId is a required field
	RouteId *string `location:"uri" locationName:"routeId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRouteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRouteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRouteInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.RouteId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRouteInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

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

type DeleteRouteOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRouteOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRouteOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteRoute = "DeleteRoute"

// DeleteRouteRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Deletes a Route.
//
//    // Example sending a request using DeleteRouteRequest.
//    req := client.DeleteRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/DeleteRoute
func (c *Client) DeleteRouteRequest(input *DeleteRouteInput) DeleteRouteRequest {
	op := &aws.Operation{
		Name:       opDeleteRoute,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v2/apis/{apiId}/routes/{routeId}",
	}

	if input == nil {
		input = &DeleteRouteInput{}
	}

	req := c.newRequest(op, input, &DeleteRouteOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteRouteRequest{Request: req, Input: input, Copy: c.DeleteRouteRequest}
}

// DeleteRouteRequest is the request type for the
// DeleteRoute API operation.
type DeleteRouteRequest struct {
	*aws.Request
	Input *DeleteRouteInput
	Copy  func(*DeleteRouteInput) DeleteRouteRequest
}

// Send marshals and sends the DeleteRoute API request.
func (r DeleteRouteRequest) Send(ctx context.Context) (*DeleteRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRouteResponse{
		DeleteRouteOutput: r.Request.Data.(*DeleteRouteOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRouteResponse is the response type for the
// DeleteRoute API operation.
type DeleteRouteResponse struct {
	*DeleteRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRoute request.
func (r *DeleteRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
