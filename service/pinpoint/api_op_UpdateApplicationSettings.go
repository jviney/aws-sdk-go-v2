// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type UpdateApplicationSettingsInput struct {
	_ struct{} `type:"structure" payload:"WriteApplicationSettingsRequest"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// Specifies the default settings for an application.
	//
	// WriteApplicationSettingsRequest is a required field
	WriteApplicationSettingsRequest *WriteApplicationSettingsRequest `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateApplicationSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateApplicationSettingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateApplicationSettingsInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.WriteApplicationSettingsRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("WriteApplicationSettingsRequest"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateApplicationSettingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.WriteApplicationSettingsRequest != nil {
		v := s.WriteApplicationSettingsRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "WriteApplicationSettingsRequest", v, metadata)
	}
	return nil
}

type UpdateApplicationSettingsOutput struct {
	_ struct{} `type:"structure" payload:"ApplicationSettingsResource"`

	// Provides information about an application, including the default settings
	// for an application.
	//
	// ApplicationSettingsResource is a required field
	ApplicationSettingsResource *ApplicationSettingsResource `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateApplicationSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateApplicationSettingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApplicationSettingsResource != nil {
		v := s.ApplicationSettingsResource

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "ApplicationSettingsResource", v, metadata)
	}
	return nil
}

const opUpdateApplicationSettings = "UpdateApplicationSettings"

// UpdateApplicationSettingsRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Updates the settings for an application.
//
//    // Example sending a request using UpdateApplicationSettingsRequest.
//    req := client.UpdateApplicationSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateApplicationSettings
func (c *Client) UpdateApplicationSettingsRequest(input *UpdateApplicationSettingsInput) UpdateApplicationSettingsRequest {
	op := &aws.Operation{
		Name:       opUpdateApplicationSettings,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/apps/{application-id}/settings",
	}

	if input == nil {
		input = &UpdateApplicationSettingsInput{}
	}

	req := c.newRequest(op, input, &UpdateApplicationSettingsOutput{})

	return UpdateApplicationSettingsRequest{Request: req, Input: input, Copy: c.UpdateApplicationSettingsRequest}
}

// UpdateApplicationSettingsRequest is the request type for the
// UpdateApplicationSettings API operation.
type UpdateApplicationSettingsRequest struct {
	*aws.Request
	Input *UpdateApplicationSettingsInput
	Copy  func(*UpdateApplicationSettingsInput) UpdateApplicationSettingsRequest
}

// Send marshals and sends the UpdateApplicationSettings API request.
func (r UpdateApplicationSettingsRequest) Send(ctx context.Context) (*UpdateApplicationSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateApplicationSettingsResponse{
		UpdateApplicationSettingsOutput: r.Request.Data.(*UpdateApplicationSettingsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateApplicationSettingsResponse is the response type for the
// UpdateApplicationSettings API operation.
type UpdateApplicationSettingsResponse struct {
	*UpdateApplicationSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApplicationSettings request.
func (r *UpdateApplicationSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
