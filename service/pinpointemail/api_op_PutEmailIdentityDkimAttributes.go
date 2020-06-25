// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// A request to enable or disable DKIM signing of email that you send from an
// email identity.
type PutEmailIdentityDkimAttributesInput struct {
	_ struct{} `type:"structure"`

	// The email identity that you want to change the DKIM settings for.
	//
	// EmailIdentity is a required field
	EmailIdentity *string `location:"uri" locationName:"EmailIdentity" type:"string" required:"true"`

	// Sets the DKIM signing configuration for the identity.
	//
	// When you set this value true, then the messages that Amazon Pinpoint sends
	// from the identity are DKIM-signed. When you set this value to false, then
	// the messages that Amazon Pinpoint sends from the identity aren't DKIM-signed.
	SigningEnabled *bool `type:"boolean"`
}

// String returns the string representation
func (s PutEmailIdentityDkimAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutEmailIdentityDkimAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutEmailIdentityDkimAttributesInput"}

	if s.EmailIdentity == nil {
		invalidParams.Add(aws.NewErrParamRequired("EmailIdentity"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutEmailIdentityDkimAttributesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.SigningEnabled != nil {
		v := *s.SigningEnabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SigningEnabled", protocol.BoolValue(v), metadata)
	}
	if s.EmailIdentity != nil {
		v := *s.EmailIdentity

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "EmailIdentity", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
type PutEmailIdentityDkimAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutEmailIdentityDkimAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutEmailIdentityDkimAttributesOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutEmailIdentityDkimAttributes = "PutEmailIdentityDkimAttributes"

// PutEmailIdentityDkimAttributesRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Used to enable or disable DKIM authentication for an email identity.
//
//    // Example sending a request using PutEmailIdentityDkimAttributesRequest.
//    req := client.PutEmailIdentityDkimAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/PutEmailIdentityDkimAttributes
func (c *Client) PutEmailIdentityDkimAttributesRequest(input *PutEmailIdentityDkimAttributesInput) PutEmailIdentityDkimAttributesRequest {
	op := &aws.Operation{
		Name:       opPutEmailIdentityDkimAttributes,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/email/identities/{EmailIdentity}/dkim",
	}

	if input == nil {
		input = &PutEmailIdentityDkimAttributesInput{}
	}

	req := c.newRequest(op, input, &PutEmailIdentityDkimAttributesOutput{})

	return PutEmailIdentityDkimAttributesRequest{Request: req, Input: input, Copy: c.PutEmailIdentityDkimAttributesRequest}
}

// PutEmailIdentityDkimAttributesRequest is the request type for the
// PutEmailIdentityDkimAttributes API operation.
type PutEmailIdentityDkimAttributesRequest struct {
	*aws.Request
	Input *PutEmailIdentityDkimAttributesInput
	Copy  func(*PutEmailIdentityDkimAttributesInput) PutEmailIdentityDkimAttributesRequest
}

// Send marshals and sends the PutEmailIdentityDkimAttributes API request.
func (r PutEmailIdentityDkimAttributesRequest) Send(ctx context.Context) (*PutEmailIdentityDkimAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutEmailIdentityDkimAttributesResponse{
		PutEmailIdentityDkimAttributesOutput: r.Request.Data.(*PutEmailIdentityDkimAttributesOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutEmailIdentityDkimAttributesResponse is the response type for the
// PutEmailIdentityDkimAttributes API operation.
type PutEmailIdentityDkimAttributesResponse struct {
	*PutEmailIdentityDkimAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutEmailIdentityDkimAttributes request.
func (r *PutEmailIdentityDkimAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
