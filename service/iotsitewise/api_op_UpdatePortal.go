// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type UpdatePortalInput struct {
	_ struct{} `type:"structure"`

	// A unique case-sensitive identifier that you can provide to ensure the idempotency
	// of the request. Don't reuse this client token if a new idempotent request
	// is required.
	ClientToken *string `locationName:"clientToken" min:"36" type:"string" idempotencyToken:"true"`

	// The AWS administrator's contact email address.
	//
	// PortalContactEmail is a required field
	PortalContactEmail *string `locationName:"portalContactEmail" min:"1" type:"string" required:"true"`

	// A new description for the portal.
	PortalDescription *string `locationName:"portalDescription" min:"1" type:"string"`

	// The ID of the portal to update.
	//
	// PortalId is a required field
	PortalId *string `location:"uri" locationName:"portalId" min:"36" type:"string" required:"true"`

	// Contains an image that is one of the following:
	//
	//    * An image file. Choose this option to upload a new image.
	//
	//    * The ID of an existing image. Choose this option to keep an existing
	//    image.
	PortalLogoImage *Image `locationName:"portalLogoImage" type:"structure"`

	// A new friendly name for the portal.
	//
	// PortalName is a required field
	PortalName *string `locationName:"portalName" min:"1" type:"string" required:"true"`

	// The ARN (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of a service role that allows the portal's users to access your AWS IoT SiteWise
	// resources on your behalf. For more information, see Using service roles for
	// AWS IoT SiteWise Monitor (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/monitor-service-role.html)
	// in the AWS IoT SiteWise User Guide.
	//
	// RoleArn is a required field
	RoleArn *string `locationName:"roleArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdatePortalInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdatePortalInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdatePortalInput"}
	if s.ClientToken != nil && len(*s.ClientToken) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 36))
	}

	if s.PortalContactEmail == nil {
		invalidParams.Add(aws.NewErrParamRequired("PortalContactEmail"))
	}
	if s.PortalContactEmail != nil && len(*s.PortalContactEmail) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PortalContactEmail", 1))
	}
	if s.PortalDescription != nil && len(*s.PortalDescription) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PortalDescription", 1))
	}

	if s.PortalId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PortalId"))
	}
	if s.PortalId != nil && len(*s.PortalId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("PortalId", 36))
	}

	if s.PortalName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PortalName"))
	}
	if s.PortalName != nil && len(*s.PortalName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PortalName", 1))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 1))
	}
	if s.PortalLogoImage != nil {
		if err := s.PortalLogoImage.Validate(); err != nil {
			invalidParams.AddNested("PortalLogoImage", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdatePortalInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	var ClientToken string
	if s.ClientToken != nil {
		ClientToken = *s.ClientToken
	} else {
		ClientToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PortalContactEmail != nil {
		v := *s.PortalContactEmail

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "portalContactEmail", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PortalDescription != nil {
		v := *s.PortalDescription

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "portalDescription", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PortalLogoImage != nil {
		v := s.PortalLogoImage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "portalLogoImage", v, metadata)
	}
	if s.PortalName != nil {
		v := *s.PortalName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "portalName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RoleArn != nil {
		v := *s.RoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "roleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PortalId != nil {
		v := *s.PortalId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "portalId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdatePortalOutput struct {
	_ struct{} `type:"structure"`

	// The status of the portal, which contains a state (UPDATING after successfully
	// calling this operation) and any error message.
	//
	// PortalStatus is a required field
	PortalStatus *PortalStatus `locationName:"portalStatus" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdatePortalOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdatePortalOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.PortalStatus != nil {
		v := s.PortalStatus

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "portalStatus", v, metadata)
	}
	return nil
}

const opUpdatePortal = "UpdatePortal"

// UpdatePortalRequest returns a request value for making API operation for
// AWS IoT SiteWise.
//
// Updates an AWS IoT SiteWise Monitor portal.
//
//    // Example sending a request using UpdatePortalRequest.
//    req := client.UpdatePortalRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotsitewise-2019-12-02/UpdatePortal
func (c *Client) UpdatePortalRequest(input *UpdatePortalInput) UpdatePortalRequest {
	op := &aws.Operation{
		Name:       opUpdatePortal,
		HTTPMethod: "PUT",
		HTTPPath:   "/portals/{portalId}",
	}

	if input == nil {
		input = &UpdatePortalInput{}
	}

	req := c.newRequest(op, input, &UpdatePortalOutput{})
	req.Handlers.Build.PushBackNamed(protocol.NewHostPrefixHandler("monitor.", nil))
	req.Handlers.Build.PushBackNamed(protocol.ValidateEndpointHostHandler)

	return UpdatePortalRequest{Request: req, Input: input, Copy: c.UpdatePortalRequest}
}

// UpdatePortalRequest is the request type for the
// UpdatePortal API operation.
type UpdatePortalRequest struct {
	*aws.Request
	Input *UpdatePortalInput
	Copy  func(*UpdatePortalInput) UpdatePortalRequest
}

// Send marshals and sends the UpdatePortal API request.
func (r UpdatePortalRequest) Send(ctx context.Context) (*UpdatePortalResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdatePortalResponse{
		UpdatePortalOutput: r.Request.Data.(*UpdatePortalOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdatePortalResponse is the response type for the
// UpdatePortal API operation.
type UpdatePortalResponse struct {
	*UpdatePortalOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdatePortal request.
func (r *UpdatePortalResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
