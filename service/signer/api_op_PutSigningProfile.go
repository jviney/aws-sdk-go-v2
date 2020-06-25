// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package signer

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type PutSigningProfileInput struct {
	_ struct{} `type:"structure"`

	// A subfield of platform. This specifies any different configuration options
	// that you want to apply to the chosen platform (such as a different hash-algorithm
	// or signing-algorithm).
	Overrides *SigningPlatformOverrides `locationName:"overrides" type:"structure"`

	// The ID of the signing platform to be created.
	//
	// PlatformId is a required field
	PlatformId *string `locationName:"platformId" type:"string" required:"true"`

	// The name of the signing profile to be created.
	//
	// ProfileName is a required field
	ProfileName *string `location:"uri" locationName:"profileName" min:"2" type:"string" required:"true"`

	// The AWS Certificate Manager certificate that will be used to sign code with
	// the new signing profile.
	//
	// SigningMaterial is a required field
	SigningMaterial *SigningMaterial `locationName:"signingMaterial" type:"structure" required:"true"`

	// Map of key-value pairs for signing. These can include any information that
	// you want to use during signing.
	SigningParameters map[string]string `locationName:"signingParameters" type:"map"`

	// Tags to be associated with the signing profile that is being created.
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`
}

// String returns the string representation
func (s PutSigningProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutSigningProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutSigningProfileInput"}

	if s.PlatformId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlatformId"))
	}

	if s.ProfileName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProfileName"))
	}
	if s.ProfileName != nil && len(*s.ProfileName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("ProfileName", 2))
	}

	if s.SigningMaterial == nil {
		invalidParams.Add(aws.NewErrParamRequired("SigningMaterial"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.SigningMaterial != nil {
		if err := s.SigningMaterial.Validate(); err != nil {
			invalidParams.AddNested("SigningMaterial", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutSigningProfileInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Overrides != nil {
		v := s.Overrides

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "overrides", v, metadata)
	}
	if s.PlatformId != nil {
		v := *s.PlatformId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "platformId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SigningMaterial != nil {
		v := s.SigningMaterial

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "signingMaterial", v, metadata)
	}
	if s.SigningParameters != nil {
		v := s.SigningParameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "signingParameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

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
	if s.ProfileName != nil {
		v := *s.ProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "profileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type PutSigningProfileOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the signing profile created.
	Arn *string `locationName:"arn" type:"string"`
}

// String returns the string representation
func (s PutSigningProfileOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutSigningProfileOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opPutSigningProfile = "PutSigningProfile"

// PutSigningProfileRequest returns a request value for making API operation for
// AWS Signer.
//
// Creates a signing profile. A signing profile is a code signing template that
// can be used to carry out a pre-defined signing job. For more information,
// see http://docs.aws.amazon.com/signer/latest/developerguide/gs-profile.html
// (http://docs.aws.amazon.com/signer/latest/developerguide/gs-profile.html)
//
//    // Example sending a request using PutSigningProfileRequest.
//    req := client.PutSigningProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/signer-2017-08-25/PutSigningProfile
func (c *Client) PutSigningProfileRequest(input *PutSigningProfileInput) PutSigningProfileRequest {
	op := &aws.Operation{
		Name:       opPutSigningProfile,
		HTTPMethod: "PUT",
		HTTPPath:   "/signing-profiles/{profileName}",
	}

	if input == nil {
		input = &PutSigningProfileInput{}
	}

	req := c.newRequest(op, input, &PutSigningProfileOutput{})

	return PutSigningProfileRequest{Request: req, Input: input, Copy: c.PutSigningProfileRequest}
}

// PutSigningProfileRequest is the request type for the
// PutSigningProfile API operation.
type PutSigningProfileRequest struct {
	*aws.Request
	Input *PutSigningProfileInput
	Copy  func(*PutSigningProfileInput) PutSigningProfileRequest
}

// Send marshals and sends the PutSigningProfile API request.
func (r PutSigningProfileRequest) Send(ctx context.Context) (*PutSigningProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutSigningProfileResponse{
		PutSigningProfileOutput: r.Request.Data.(*PutSigningProfileOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutSigningProfileResponse is the response type for the
// PutSigningProfile API operation.
type PutSigningProfileResponse struct {
	*PutSigningProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutSigningProfile request.
func (r *PutSigningProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
