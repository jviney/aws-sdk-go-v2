// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appconfig

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type CreateApplicationInput struct {
	_ struct{} `type:"structure"`

	// A description of the application.
	Description *string `type:"string"`

	// A name for the application.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// Metadata to assign to the application. Tags help organize and categorize
	// your AppConfig resources. Each tag consists of a key and an optional value,
	// both of which you define.
	Tags map[string]string `type:"map"`
}

// String returns the string representation
func (s CreateApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateApplicationInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateApplicationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

type CreateApplicationOutput struct {
	_ struct{} `type:"structure"`

	// The description of the application.
	Description *string `type:"string"`

	// The application ID.
	Id *string `type:"string"`

	// The application name.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateApplicationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateApplication = "CreateApplication"

// CreateApplicationRequest returns a request value for making API operation for
// Amazon AppConfig.
//
// An application in AppConfig is a logical unit of code that provides capabilities
// for your customers. For example, an application can be a microservice that
// runs on Amazon EC2 instances, a mobile application installed by your users,
// a serverless application using Amazon API Gateway and AWS Lambda, or any
// system you run on behalf of others.
//
//    // Example sending a request using CreateApplicationRequest.
//    req := client.CreateApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appconfig-2019-10-09/CreateApplication
func (c *Client) CreateApplicationRequest(input *CreateApplicationInput) CreateApplicationRequest {
	op := &aws.Operation{
		Name:       opCreateApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/applications",
	}

	if input == nil {
		input = &CreateApplicationInput{}
	}

	req := c.newRequest(op, input, &CreateApplicationOutput{})

	return CreateApplicationRequest{Request: req, Input: input, Copy: c.CreateApplicationRequest}
}

// CreateApplicationRequest is the request type for the
// CreateApplication API operation.
type CreateApplicationRequest struct {
	*aws.Request
	Input *CreateApplicationInput
	Copy  func(*CreateApplicationInput) CreateApplicationRequest
}

// Send marshals and sends the CreateApplication API request.
func (r CreateApplicationRequest) Send(ctx context.Context) (*CreateApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateApplicationResponse{
		CreateApplicationOutput: r.Request.Data.(*CreateApplicationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateApplicationResponse is the response type for the
// CreateApplication API operation.
type CreateApplicationResponse struct {
	*CreateApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateApplication request.
func (r *CreateApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
