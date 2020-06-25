// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to delete a configuration set event destination. Configuration
// set event destinations are associated with configuration sets, which enable
// you to publish email sending events. For information about using configuration
// sets, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
type DeleteConfigurationSetEventDestinationInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set from which to delete the event destination.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `type:"string" required:"true"`

	// The name of the event destination to delete.
	//
	// EventDestinationName is a required field
	EventDestinationName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteConfigurationSetEventDestinationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteConfigurationSetEventDestinationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteConfigurationSetEventDestinationInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if s.EventDestinationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventDestinationName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An empty element returned on a successful request.
type DeleteConfigurationSetEventDestinationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteConfigurationSetEventDestinationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteConfigurationSetEventDestination = "DeleteConfigurationSetEventDestination"

// DeleteConfigurationSetEventDestinationRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Deletes a configuration set event destination. Configuration set event destinations
// are associated with configuration sets, which enable you to publish email
// sending events. For information about using configuration sets, see the Amazon
// SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using DeleteConfigurationSetEventDestinationRequest.
//    req := client.DeleteConfigurationSetEventDestinationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DeleteConfigurationSetEventDestination
func (c *Client) DeleteConfigurationSetEventDestinationRequest(input *DeleteConfigurationSetEventDestinationInput) DeleteConfigurationSetEventDestinationRequest {
	op := &aws.Operation{
		Name:       opDeleteConfigurationSetEventDestination,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteConfigurationSetEventDestinationInput{}
	}

	req := c.newRequest(op, input, &DeleteConfigurationSetEventDestinationOutput{})

	return DeleteConfigurationSetEventDestinationRequest{Request: req, Input: input, Copy: c.DeleteConfigurationSetEventDestinationRequest}
}

// DeleteConfigurationSetEventDestinationRequest is the request type for the
// DeleteConfigurationSetEventDestination API operation.
type DeleteConfigurationSetEventDestinationRequest struct {
	*aws.Request
	Input *DeleteConfigurationSetEventDestinationInput
	Copy  func(*DeleteConfigurationSetEventDestinationInput) DeleteConfigurationSetEventDestinationRequest
}

// Send marshals and sends the DeleteConfigurationSetEventDestination API request.
func (r DeleteConfigurationSetEventDestinationRequest) Send(ctx context.Context) (*DeleteConfigurationSetEventDestinationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteConfigurationSetEventDestinationResponse{
		DeleteConfigurationSetEventDestinationOutput: r.Request.Data.(*DeleteConfigurationSetEventDestinationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteConfigurationSetEventDestinationResponse is the response type for the
// DeleteConfigurationSetEventDestination API operation.
type DeleteConfigurationSetEventDestinationResponse struct {
	*DeleteConfigurationSetEventDestinationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteConfigurationSetEventDestination request.
func (r *DeleteConfigurationSetEventDestinationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
