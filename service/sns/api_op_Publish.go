// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Input for Publish action.
type PublishInput struct {
	_ struct{} `type:"structure"`

	// The message you want to send.
	//
	// If you are publishing to a topic and you want to send the same message to
	// all transport protocols, include the text of the message as a String value.
	// If you want to send different messages for each transport protocol, set the
	// value of the MessageStructure parameter to json and use a JSON object for
	// the Message parameter.
	//
	// Constraints:
	//
	//    * With the exception of SMS, messages must be UTF-8 encoded strings and
	//    at most 256 KB in size (262,144 bytes, not 262,144 characters).
	//
	//    * For SMS, each message can contain up to 140 characters. This character
	//    limit depends on the encoding schema. For example, an SMS message can
	//    contain 160 GSM characters, 140 ASCII characters, or 70 UCS-2 characters.
	//    If you publish a message that exceeds this size limit, Amazon SNS sends
	//    the message as multiple messages, each fitting within the size limit.
	//    Messages aren't truncated mid-word but are cut off at whole-word boundaries.
	//    The total size limit for a single SMS Publish action is 1,600 characters.
	//
	// JSON-specific constraints:
	//
	//    * Keys in the JSON object that correspond to supported transport protocols
	//    must have simple JSON string values.
	//
	//    * The values will be parsed (unescaped) before they are used in outgoing
	//    messages.
	//
	//    * Outbound notifications are JSON encoded (meaning that the characters
	//    will be reescaped for sending).
	//
	//    * Values have a minimum length of 0 (the empty string, "", is allowed).
	//
	//    * Values have a maximum length bounded by the overall message size (so,
	//    including multiple protocols may limit message sizes).
	//
	//    * Non-string values will cause the key to be ignored.
	//
	//    * Keys that do not correspond to supported transport protocols are ignored.
	//
	//    * Duplicate keys are not allowed.
	//
	//    * Failure to parse or validate any key or value in the message will cause
	//    the Publish call to return an error (no partial delivery).
	//
	// Message is a required field
	Message *string `type:"string" required:"true"`

	// Message attributes for Publish action.
	MessageAttributes map[string]MessageAttributeValue `locationNameKey:"Name" locationNameValue:"Value" type:"map"`

	// Set MessageStructure to json if you want to send a different message for
	// each protocol. For example, using one publish action, you can send a short
	// message to your SMS subscribers and a longer message to your email subscribers.
	// If you set MessageStructure to json, the value of the Message parameter must:
	//
	//    * be a syntactically valid JSON object; and
	//
	//    * contain at least a top-level JSON key of "default" with a value that
	//    is a string.
	//
	// You can define other top-level keys that define the message you want to send
	// to a specific transport protocol (e.g., "http").
	//
	// Valid value: json
	MessageStructure *string `type:"string"`

	// The phone number to which you want to deliver an SMS message. Use E.164 format.
	//
	// If you don't specify a value for the PhoneNumber parameter, you must specify
	// a value for the TargetArn or TopicArn parameters.
	PhoneNumber *string `type:"string"`

	// Optional parameter to be used as the "Subject" line when the message is delivered
	// to email endpoints. This field will also be included, if present, in the
	// standard JSON messages delivered to other endpoints.
	//
	// Constraints: Subjects must be ASCII text that begins with a letter, number,
	// or punctuation mark; must not include line breaks or control characters;
	// and must be less than 100 characters long.
	Subject *string `type:"string"`

	// If you don't specify a value for the TargetArn parameter, you must specify
	// a value for the PhoneNumber or TopicArn parameters.
	TargetArn *string `type:"string"`

	// The topic you want to publish to.
	//
	// If you don't specify a value for the TopicArn parameter, you must specify
	// a value for the PhoneNumber or TargetArn parameters.
	TopicArn *string `type:"string"`
}

// String returns the string representation
func (s PublishInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PublishInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PublishInput"}

	if s.Message == nil {
		invalidParams.Add(aws.NewErrParamRequired("Message"))
	}
	if s.MessageAttributes != nil {
		for i, v := range s.MessageAttributes {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "MessageAttributes", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response for Publish action.
type PublishOutput struct {
	_ struct{} `type:"structure"`

	// Unique identifier assigned to the published message.
	//
	// Length Constraint: Maximum 100 characters
	MessageId *string `type:"string"`
}

// String returns the string representation
func (s PublishOutput) String() string {
	return awsutil.Prettify(s)
}

const opPublish = "Publish"

// PublishRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Sends a message to an Amazon SNS topic or sends a text message (SMS message)
// directly to a phone number.
//
// If you send a message to a topic, Amazon SNS delivers the message to each
// endpoint that is subscribed to the topic. The format of the message depends
// on the notification protocol for each subscribed endpoint.
//
// When a messageId is returned, the message has been saved and Amazon SNS will
// attempt to deliver it shortly.
//
// To use the Publish action for sending a message to a mobile endpoint, such
// as an app on a Kindle device or mobile phone, you must specify the EndpointArn
// for the TargetArn parameter. The EndpointArn is returned when making a call
// with the CreatePlatformEndpoint action.
//
// For more information about formatting messages, see Send Custom Platform-Specific
// Payloads in Messages to Mobile Devices (https://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-custommessage.html).
//
//    // Example sending a request using PublishRequest.
//    req := client.PublishRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/Publish
func (c *Client) PublishRequest(input *PublishInput) PublishRequest {
	op := &aws.Operation{
		Name:       opPublish,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PublishInput{}
	}

	req := c.newRequest(op, input, &PublishOutput{})

	return PublishRequest{Request: req, Input: input, Copy: c.PublishRequest}
}

// PublishRequest is the request type for the
// Publish API operation.
type PublishRequest struct {
	*aws.Request
	Input *PublishInput
	Copy  func(*PublishInput) PublishRequest
}

// Send marshals and sends the Publish API request.
func (r PublishRequest) Send(ctx context.Context) (*PublishResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PublishResponse{
		PublishOutput: r.Request.Data.(*PublishOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PublishResponse is the response type for the
// Publish API operation.
type PublishResponse struct {
	*PublishOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// Publish request.
func (r *PublishResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
