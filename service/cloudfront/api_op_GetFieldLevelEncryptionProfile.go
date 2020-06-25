// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetFieldLevelEncryptionProfileInput struct {
	_ struct{} `type:"structure"`

	// Get the ID for the field-level encryption profile information.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetFieldLevelEncryptionProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFieldLevelEncryptionProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetFieldLevelEncryptionProfileInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFieldLevelEncryptionProfileInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	return nil
}

type GetFieldLevelEncryptionProfileOutput struct {
	_ struct{} `type:"structure" payload:"FieldLevelEncryptionProfile"`

	// The current version of the field level encryption profile. For example: E2QWRUHAPOMQZL.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// Return the field-level encryption profile information.
	FieldLevelEncryptionProfile *FieldLevelEncryptionProfile `type:"structure"`
}

// String returns the string representation
func (s GetFieldLevelEncryptionProfileOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFieldLevelEncryptionProfileOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.FieldLevelEncryptionProfile != nil {
		v := s.FieldLevelEncryptionProfile

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "FieldLevelEncryptionProfile", v, metadata)
	}
	return nil
}

const opGetFieldLevelEncryptionProfile = "GetFieldLevelEncryptionProfile2019_03_26"

// GetFieldLevelEncryptionProfileRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Get the field-level encryption profile information.
//
//    // Example sending a request using GetFieldLevelEncryptionProfileRequest.
//    req := client.GetFieldLevelEncryptionProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/GetFieldLevelEncryptionProfile
func (c *Client) GetFieldLevelEncryptionProfileRequest(input *GetFieldLevelEncryptionProfileInput) GetFieldLevelEncryptionProfileRequest {
	op := &aws.Operation{
		Name:       opGetFieldLevelEncryptionProfile,
		HTTPMethod: "GET",
		HTTPPath:   "/2019-03-26/field-level-encryption-profile/{Id}",
	}

	if input == nil {
		input = &GetFieldLevelEncryptionProfileInput{}
	}

	req := c.newRequest(op, input, &GetFieldLevelEncryptionProfileOutput{})

	return GetFieldLevelEncryptionProfileRequest{Request: req, Input: input, Copy: c.GetFieldLevelEncryptionProfileRequest}
}

// GetFieldLevelEncryptionProfileRequest is the request type for the
// GetFieldLevelEncryptionProfile API operation.
type GetFieldLevelEncryptionProfileRequest struct {
	*aws.Request
	Input *GetFieldLevelEncryptionProfileInput
	Copy  func(*GetFieldLevelEncryptionProfileInput) GetFieldLevelEncryptionProfileRequest
}

// Send marshals and sends the GetFieldLevelEncryptionProfile API request.
func (r GetFieldLevelEncryptionProfileRequest) Send(ctx context.Context) (*GetFieldLevelEncryptionProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFieldLevelEncryptionProfileResponse{
		GetFieldLevelEncryptionProfileOutput: r.Request.Data.(*GetFieldLevelEncryptionProfileOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetFieldLevelEncryptionProfileResponse is the response type for the
// GetFieldLevelEncryptionProfile API operation.
type GetFieldLevelEncryptionProfileResponse struct {
	*GetFieldLevelEncryptionProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFieldLevelEncryptionProfile request.
func (r *GetFieldLevelEncryptionProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
