// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// GetInfrastructureConfiguration request object.
type GetInfrastructureConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the infrastructure configuration that you
	// want to retrieve.
	//
	// InfrastructureConfigurationArn is a required field
	InfrastructureConfigurationArn *string `location:"querystring" locationName:"infrastructureConfigurationArn" type:"string" required:"true"`
}

// String returns the string representation
func (s GetInfrastructureConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetInfrastructureConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetInfrastructureConfigurationInput"}

	if s.InfrastructureConfigurationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("InfrastructureConfigurationArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetInfrastructureConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.InfrastructureConfigurationArn != nil {
		v := *s.InfrastructureConfigurationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "infrastructureConfigurationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// GetInfrastructureConfiguration response object.
type GetInfrastructureConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The infrastructure configuration object.
	InfrastructureConfiguration *InfrastructureConfiguration `locationName:"infrastructureConfiguration" type:"structure"`

	// The request ID that uniquely identifies this request.
	RequestId *string `locationName:"requestId" min:"1" type:"string"`
}

// String returns the string representation
func (s GetInfrastructureConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetInfrastructureConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.InfrastructureConfiguration != nil {
		v := s.InfrastructureConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "infrastructureConfiguration", v, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetInfrastructureConfiguration = "GetInfrastructureConfiguration"

// GetInfrastructureConfigurationRequest returns a request value for making API operation for
// EC2 Image Builder.
//
// Gets an infrastructure configuration.
//
//    // Example sending a request using GetInfrastructureConfigurationRequest.
//    req := client.GetInfrastructureConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/imagebuilder-2019-12-02/GetInfrastructureConfiguration
func (c *Client) GetInfrastructureConfigurationRequest(input *GetInfrastructureConfigurationInput) GetInfrastructureConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetInfrastructureConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/GetInfrastructureConfiguration",
	}

	if input == nil {
		input = &GetInfrastructureConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetInfrastructureConfigurationOutput{})

	return GetInfrastructureConfigurationRequest{Request: req, Input: input, Copy: c.GetInfrastructureConfigurationRequest}
}

// GetInfrastructureConfigurationRequest is the request type for the
// GetInfrastructureConfiguration API operation.
type GetInfrastructureConfigurationRequest struct {
	*aws.Request
	Input *GetInfrastructureConfigurationInput
	Copy  func(*GetInfrastructureConfigurationInput) GetInfrastructureConfigurationRequest
}

// Send marshals and sends the GetInfrastructureConfiguration API request.
func (r GetInfrastructureConfigurationRequest) Send(ctx context.Context) (*GetInfrastructureConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetInfrastructureConfigurationResponse{
		GetInfrastructureConfigurationOutput: r.Request.Data.(*GetInfrastructureConfigurationOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetInfrastructureConfigurationResponse is the response type for the
// GetInfrastructureConfiguration API operation.
type GetInfrastructureConfigurationResponse struct {
	*GetInfrastructureConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetInfrastructureConfiguration request.
func (r *GetInfrastructureConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
