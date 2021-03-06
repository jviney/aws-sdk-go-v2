// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// A request to change the settings for an event destination for a configuration
// set.
type UpdateConfigurationSetEventDestinationInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set that contains the event destination that
	// you want to modify.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `location:"uri" locationName:"ConfigurationSetName" type:"string" required:"true"`

	// An object that defines the event destination.
	//
	// EventDestination is a required field
	EventDestination *EventDestinationDefinition `type:"structure" required:"true"`

	// The name of the event destination that you want to modify.
	//
	// EventDestinationName is a required field
	EventDestinationName *string `location:"uri" locationName:"EventDestinationName" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateConfigurationSetEventDestinationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateConfigurationSetEventDestinationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateConfigurationSetEventDestinationInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if s.EventDestination == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventDestination"))
	}

	if s.EventDestinationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventDestinationName"))
	}
	if s.EventDestination != nil {
		if err := s.EventDestination.Validate(); err != nil {
			invalidParams.AddNested("EventDestination", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateConfigurationSetEventDestinationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.EventDestination != nil {
		v := s.EventDestination

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "EventDestination", v, metadata)
	}
	if s.ConfigurationSetName != nil {
		v := *s.ConfigurationSetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ConfigurationSetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EventDestinationName != nil {
		v := *s.EventDestinationName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "EventDestinationName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
type UpdateConfigurationSetEventDestinationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateConfigurationSetEventDestinationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateConfigurationSetEventDestinationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateConfigurationSetEventDestination = "UpdateConfigurationSetEventDestination"

// UpdateConfigurationSetEventDestinationRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Update the configuration of an event destination for a configuration set.
//
// In Amazon Pinpoint, events include message sends, deliveries, opens, clicks,
// bounces, and complaints. Event destinations are places that you can send
// information about these events to. For example, you can send event data to
// Amazon SNS to receive notifications when you receive bounces or complaints,
// or you can use Amazon Kinesis Data Firehose to stream data to Amazon S3 for
// long-term storage.
//
//    // Example sending a request using UpdateConfigurationSetEventDestinationRequest.
//    req := client.UpdateConfigurationSetEventDestinationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/UpdateConfigurationSetEventDestination
func (c *Client) UpdateConfigurationSetEventDestinationRequest(input *UpdateConfigurationSetEventDestinationInput) UpdateConfigurationSetEventDestinationRequest {
	op := &aws.Operation{
		Name:       opUpdateConfigurationSetEventDestination,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/email/configuration-sets/{ConfigurationSetName}/event-destinations/{EventDestinationName}",
	}

	if input == nil {
		input = &UpdateConfigurationSetEventDestinationInput{}
	}

	req := c.newRequest(op, input, &UpdateConfigurationSetEventDestinationOutput{})

	return UpdateConfigurationSetEventDestinationRequest{Request: req, Input: input, Copy: c.UpdateConfigurationSetEventDestinationRequest}
}

// UpdateConfigurationSetEventDestinationRequest is the request type for the
// UpdateConfigurationSetEventDestination API operation.
type UpdateConfigurationSetEventDestinationRequest struct {
	*aws.Request
	Input *UpdateConfigurationSetEventDestinationInput
	Copy  func(*UpdateConfigurationSetEventDestinationInput) UpdateConfigurationSetEventDestinationRequest
}

// Send marshals and sends the UpdateConfigurationSetEventDestination API request.
func (r UpdateConfigurationSetEventDestinationRequest) Send(ctx context.Context) (*UpdateConfigurationSetEventDestinationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateConfigurationSetEventDestinationResponse{
		UpdateConfigurationSetEventDestinationOutput: r.Request.Data.(*UpdateConfigurationSetEventDestinationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateConfigurationSetEventDestinationResponse is the response type for the
// UpdateConfigurationSetEventDestination API operation.
type UpdateConfigurationSetEventDestinationResponse struct {
	*UpdateConfigurationSetEventDestinationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateConfigurationSetEventDestination request.
func (r *UpdateConfigurationSetEventDestinationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
