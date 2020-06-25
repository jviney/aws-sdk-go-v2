// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sesv2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// A request to enable or disable the ability of Amazon SES to send emails that
// use a specific configuration set.
type PutConfigurationSetSendingOptionsInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set that you want to enable or disable email
	// sending for.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `location:"uri" locationName:"ConfigurationSetName" type:"string" required:"true"`

	// If true, email sending is enabled for the configuration set. If false, email
	// sending is disabled for the configuration set.
	SendingEnabled *bool `type:"boolean"`
}

// String returns the string representation
func (s PutConfigurationSetSendingOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutConfigurationSetSendingOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutConfigurationSetSendingOptionsInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutConfigurationSetSendingOptionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.SendingEnabled != nil {
		v := *s.SendingEnabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SendingEnabled", protocol.BoolValue(v), metadata)
	}
	if s.ConfigurationSetName != nil {
		v := *s.ConfigurationSetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ConfigurationSetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
type PutConfigurationSetSendingOptionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutConfigurationSetSendingOptionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutConfigurationSetSendingOptionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutConfigurationSetSendingOptions = "PutConfigurationSetSendingOptions"

// PutConfigurationSetSendingOptionsRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Enable or disable email sending for messages that use a particular configuration
// set in a specific AWS Region.
//
//    // Example sending a request using PutConfigurationSetSendingOptionsRequest.
//    req := client.PutConfigurationSetSendingOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sesv2-2019-09-27/PutConfigurationSetSendingOptions
func (c *Client) PutConfigurationSetSendingOptionsRequest(input *PutConfigurationSetSendingOptionsInput) PutConfigurationSetSendingOptionsRequest {
	op := &aws.Operation{
		Name:       opPutConfigurationSetSendingOptions,
		HTTPMethod: "PUT",
		HTTPPath:   "/v2/email/configuration-sets/{ConfigurationSetName}/sending",
	}

	if input == nil {
		input = &PutConfigurationSetSendingOptionsInput{}
	}

	req := c.newRequest(op, input, &PutConfigurationSetSendingOptionsOutput{})

	return PutConfigurationSetSendingOptionsRequest{Request: req, Input: input, Copy: c.PutConfigurationSetSendingOptionsRequest}
}

// PutConfigurationSetSendingOptionsRequest is the request type for the
// PutConfigurationSetSendingOptions API operation.
type PutConfigurationSetSendingOptionsRequest struct {
	*aws.Request
	Input *PutConfigurationSetSendingOptionsInput
	Copy  func(*PutConfigurationSetSendingOptionsInput) PutConfigurationSetSendingOptionsRequest
}

// Send marshals and sends the PutConfigurationSetSendingOptions API request.
func (r PutConfigurationSetSendingOptionsRequest) Send(ctx context.Context) (*PutConfigurationSetSendingOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutConfigurationSetSendingOptionsResponse{
		PutConfigurationSetSendingOptionsOutput: r.Request.Data.(*PutConfigurationSetSendingOptionsOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutConfigurationSetSendingOptionsResponse is the response type for the
// PutConfigurationSetSendingOptions API operation.
type PutConfigurationSetSendingOptionsResponse struct {
	*PutConfigurationSetSendingOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutConfigurationSetSendingOptions request.
func (r *PutConfigurationSetSendingOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
