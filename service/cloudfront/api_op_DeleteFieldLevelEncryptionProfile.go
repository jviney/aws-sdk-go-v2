// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/restxml"
)

type DeleteFieldLevelEncryptionProfileInput struct {
	_ struct{} `type:"structure"`

	// Request the ID of the profile you want to delete from CloudFront.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`

	// The value of the ETag header that you received when retrieving the profile
	// to delete. For example: E2QWRUHAPOMQZL.
	IfMatch *string `location:"header" locationName:"If-Match" type:"string"`
}

// String returns the string representation
func (s DeleteFieldLevelEncryptionProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFieldLevelEncryptionProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFieldLevelEncryptionProfileInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteFieldLevelEncryptionProfileInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.IfMatch != nil {
		v := *s.IfMatch

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "If-Match", protocol.StringValue(v), metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	return nil
}

type DeleteFieldLevelEncryptionProfileOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteFieldLevelEncryptionProfileOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteFieldLevelEncryptionProfileOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteFieldLevelEncryptionProfile = "DeleteFieldLevelEncryptionProfile2019_03_26"

// DeleteFieldLevelEncryptionProfileRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Remove a field-level encryption profile.
//
//    // Example sending a request using DeleteFieldLevelEncryptionProfileRequest.
//    req := client.DeleteFieldLevelEncryptionProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/DeleteFieldLevelEncryptionProfile
func (c *Client) DeleteFieldLevelEncryptionProfileRequest(input *DeleteFieldLevelEncryptionProfileInput) DeleteFieldLevelEncryptionProfileRequest {
	op := &aws.Operation{
		Name:       opDeleteFieldLevelEncryptionProfile,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2019-03-26/field-level-encryption-profile/{Id}",
	}

	if input == nil {
		input = &DeleteFieldLevelEncryptionProfileInput{}
	}

	req := c.newRequest(op, input, &DeleteFieldLevelEncryptionProfileOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteFieldLevelEncryptionProfileRequest{Request: req, Input: input, Copy: c.DeleteFieldLevelEncryptionProfileRequest}
}

// DeleteFieldLevelEncryptionProfileRequest is the request type for the
// DeleteFieldLevelEncryptionProfile API operation.
type DeleteFieldLevelEncryptionProfileRequest struct {
	*aws.Request
	Input *DeleteFieldLevelEncryptionProfileInput
	Copy  func(*DeleteFieldLevelEncryptionProfileInput) DeleteFieldLevelEncryptionProfileRequest
}

// Send marshals and sends the DeleteFieldLevelEncryptionProfile API request.
func (r DeleteFieldLevelEncryptionProfileRequest) Send(ctx context.Context) (*DeleteFieldLevelEncryptionProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteFieldLevelEncryptionProfileResponse{
		DeleteFieldLevelEncryptionProfileOutput: r.Request.Data.(*DeleteFieldLevelEncryptionProfileOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteFieldLevelEncryptionProfileResponse is the response type for the
// DeleteFieldLevelEncryptionProfile API operation.
type DeleteFieldLevelEncryptionProfileResponse struct {
	*DeleteFieldLevelEncryptionProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteFieldLevelEncryptionProfile request.
func (r *DeleteFieldLevelEncryptionProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
