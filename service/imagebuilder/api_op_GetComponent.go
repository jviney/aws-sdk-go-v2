// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetComponentInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the component that you want to retrieve.
	// Regex requires "/\d+$" suffix.
	//
	// ComponentBuildVersionArn is a required field
	ComponentBuildVersionArn *string `location:"querystring" locationName:"componentBuildVersionArn" type:"string" required:"true"`
}

// String returns the string representation
func (s GetComponentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetComponentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetComponentInput"}

	if s.ComponentBuildVersionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ComponentBuildVersionArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetComponentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ComponentBuildVersionArn != nil {
		v := *s.ComponentBuildVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "componentBuildVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetComponentOutput struct {
	_ struct{} `type:"structure"`

	// The component object associated with the specified ARN.
	Component *Component `locationName:"component" type:"structure"`

	// The request ID that uniquely identifies this request.
	RequestId *string `locationName:"requestId" min:"1" type:"string"`
}

// String returns the string representation
func (s GetComponentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetComponentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Component != nil {
		v := s.Component

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "component", v, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetComponent = "GetComponent"

// GetComponentRequest returns a request value for making API operation for
// EC2 Image Builder.
//
// Gets a component object.
//
//    // Example sending a request using GetComponentRequest.
//    req := client.GetComponentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/imagebuilder-2019-12-02/GetComponent
func (c *Client) GetComponentRequest(input *GetComponentInput) GetComponentRequest {
	op := &aws.Operation{
		Name:       opGetComponent,
		HTTPMethod: "GET",
		HTTPPath:   "/GetComponent",
	}

	if input == nil {
		input = &GetComponentInput{}
	}

	req := c.newRequest(op, input, &GetComponentOutput{})

	return GetComponentRequest{Request: req, Input: input, Copy: c.GetComponentRequest}
}

// GetComponentRequest is the request type for the
// GetComponent API operation.
type GetComponentRequest struct {
	*aws.Request
	Input *GetComponentInput
	Copy  func(*GetComponentInput) GetComponentRequest
}

// Send marshals and sends the GetComponent API request.
func (r GetComponentRequest) Send(ctx context.Context) (*GetComponentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetComponentResponse{
		GetComponentOutput: r.Request.Data.(*GetComponentOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetComponentResponse is the response type for the
// GetComponent API operation.
type GetComponentResponse struct {
	*GetComponentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetComponent request.
func (r *GetComponentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
