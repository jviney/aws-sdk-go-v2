// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointsmsvoice

import (
	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// An object that defines a message that contains text formatted using Amazon
// Pinpoint Voice Instructions markup.
type CallInstructionsMessageType struct {
	_ struct{} `type:"structure"`

	// The language to use when delivering the message. For a complete list of supported
	// languages, see the Amazon Polly Developer Guide.
	Text *string `type:"string"`
}

// String returns the string representation
func (s CallInstructionsMessageType) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CallInstructionsMessageType) MarshalFields(e protocol.FieldEncoder) error {
	if s.Text != nil {
		v := *s.Text

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Text", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An object that contains information about an event destination that sends
// data to Amazon CloudWatch Logs.
type CloudWatchLogsDestination struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of an Amazon Identity and Access Management
	// (IAM) role that is able to write event data to an Amazon CloudWatch destination.
	IamRoleArn *string `type:"string"`

	// The name of the Amazon CloudWatch Log Group that you want to record events
	// in.
	LogGroupArn *string `type:"string"`
}

// String returns the string representation
func (s CloudWatchLogsDestination) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CloudWatchLogsDestination) MarshalFields(e protocol.FieldEncoder) error {
	if s.IamRoleArn != nil {
		v := *s.IamRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IamRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LogGroupArn != nil {
		v := *s.LogGroupArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LogGroupArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An object that defines an event destination.
type EventDestination struct {
	_ struct{} `type:"structure"`

	// An object that contains information about an event destination that sends
	// data to Amazon CloudWatch Logs.
	CloudWatchLogsDestination *CloudWatchLogsDestination `type:"structure"`

	// Indicates whether or not the event destination is enabled. If the event destination
	// is enabled, then Amazon Pinpoint sends response data to the specified event
	// destination.
	Enabled *bool `type:"boolean"`

	// An object that contains information about an event destination that sends
	// data to Amazon Kinesis Data Firehose.
	KinesisFirehoseDestination *KinesisFirehoseDestination `type:"structure"`

	// An array of EventDestination objects. Each EventDestination object includes
	// ARNs and other information that define an event destination.
	MatchingEventTypes []EventType `type:"list"`

	// A name that identifies the event destination configuration.
	Name *string `type:"string"`

	// An object that contains information about an event destination that sends
	// data to Amazon SNS.
	SnsDestination *SnsDestination `type:"structure"`
}

// String returns the string representation
func (s EventDestination) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s EventDestination) MarshalFields(e protocol.FieldEncoder) error {
	if s.CloudWatchLogsDestination != nil {
		v := s.CloudWatchLogsDestination

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "CloudWatchLogsDestination", v, metadata)
	}
	if s.Enabled != nil {
		v := *s.Enabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Enabled", protocol.BoolValue(v), metadata)
	}
	if s.KinesisFirehoseDestination != nil {
		v := s.KinesisFirehoseDestination

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "KinesisFirehoseDestination", v, metadata)
	}
	if s.MatchingEventTypes != nil {
		v := s.MatchingEventTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "MatchingEventTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SnsDestination != nil {
		v := s.SnsDestination

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SnsDestination", v, metadata)
	}
	return nil
}

// An object that defines a single event destination.
type EventDestinationDefinition struct {
	_ struct{} `type:"structure"`

	// An object that contains information about an event destination that sends
	// data to Amazon CloudWatch Logs.
	CloudWatchLogsDestination *CloudWatchLogsDestination `type:"structure"`

	// Indicates whether or not the event destination is enabled. If the event destination
	// is enabled, then Amazon Pinpoint sends response data to the specified event
	// destination.
	Enabled *bool `type:"boolean"`

	// An object that contains information about an event destination that sends
	// data to Amazon Kinesis Data Firehose.
	KinesisFirehoseDestination *KinesisFirehoseDestination `type:"structure"`

	// An array of EventDestination objects. Each EventDestination object includes
	// ARNs and other information that define an event destination.
	MatchingEventTypes []EventType `type:"list"`

	// An object that contains information about an event destination that sends
	// data to Amazon SNS.
	SnsDestination *SnsDestination `type:"structure"`
}

// String returns the string representation
func (s EventDestinationDefinition) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s EventDestinationDefinition) MarshalFields(e protocol.FieldEncoder) error {
	if s.CloudWatchLogsDestination != nil {
		v := s.CloudWatchLogsDestination

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "CloudWatchLogsDestination", v, metadata)
	}
	if s.Enabled != nil {
		v := *s.Enabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Enabled", protocol.BoolValue(v), metadata)
	}
	if s.KinesisFirehoseDestination != nil {
		v := s.KinesisFirehoseDestination

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "KinesisFirehoseDestination", v, metadata)
	}
	if s.MatchingEventTypes != nil {
		v := s.MatchingEventTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "MatchingEventTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.SnsDestination != nil {
		v := s.SnsDestination

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SnsDestination", v, metadata)
	}
	return nil
}

// An object that contains information about an event destination that sends
// data to Amazon Kinesis Data Firehose.
type KinesisFirehoseDestination struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of an IAM role that can write data to an Amazon
	// Kinesis Data Firehose stream.
	DeliveryStreamArn *string `type:"string"`

	// The Amazon Resource Name (ARN) of the Amazon Kinesis Data Firehose destination
	// that you want to use in the event destination.
	IamRoleArn *string `type:"string"`
}

// String returns the string representation
func (s KinesisFirehoseDestination) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s KinesisFirehoseDestination) MarshalFields(e protocol.FieldEncoder) error {
	if s.DeliveryStreamArn != nil {
		v := *s.DeliveryStreamArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeliveryStreamArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IamRoleArn != nil {
		v := *s.IamRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IamRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An object that defines a message that contains unformatted text.
type PlainTextMessageType struct {
	_ struct{} `type:"structure"`

	// The language to use when delivering the message. For a complete list of supported
	// languages, see the Amazon Polly Developer Guide.
	LanguageCode *string `type:"string"`

	// The plain (not SSML-formatted) text to deliver to the recipient.
	Text *string `type:"string"`

	// The name of the voice that you want to use to deliver the message. For a
	// complete list of supported voices, see the Amazon Polly Developer Guide.
	VoiceId *string `type:"string"`
}

// String returns the string representation
func (s PlainTextMessageType) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PlainTextMessageType) MarshalFields(e protocol.FieldEncoder) error {
	if s.LanguageCode != nil {
		v := *s.LanguageCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LanguageCode", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Text != nil {
		v := *s.Text

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Text", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VoiceId != nil {
		v := *s.VoiceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VoiceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An object that defines a message that contains SSML-formatted text.
type SSMLMessageType struct {
	_ struct{} `type:"structure"`

	// The language to use when delivering the message. For a complete list of supported
	// languages, see the Amazon Polly Developer Guide.
	LanguageCode *string `type:"string"`

	// The SSML-formatted text to deliver to the recipient.
	Text *string `type:"string"`

	// The name of the voice that you want to use to deliver the message. For a
	// complete list of supported voices, see the Amazon Polly Developer Guide.
	VoiceId *string `type:"string"`
}

// String returns the string representation
func (s SSMLMessageType) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SSMLMessageType) MarshalFields(e protocol.FieldEncoder) error {
	if s.LanguageCode != nil {
		v := *s.LanguageCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LanguageCode", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Text != nil {
		v := *s.Text

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Text", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VoiceId != nil {
		v := *s.VoiceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VoiceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An object that contains information about an event destination that sends
// data to Amazon SNS.
type SnsDestination struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the Amazon SNS topic that you want to publish
	// events to.
	TopicArn *string `type:"string"`
}

// String returns the string representation
func (s SnsDestination) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SnsDestination) MarshalFields(e protocol.FieldEncoder) error {
	if s.TopicArn != nil {
		v := *s.TopicArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TopicArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An object that contains a voice message and information about the recipient
// that you want to send it to.
type VoiceMessageContent struct {
	_ struct{} `type:"structure"`

	// An object that defines a message that contains text formatted using Amazon
	// Pinpoint Voice Instructions markup.
	CallInstructionsMessage *CallInstructionsMessageType `type:"structure"`

	// An object that defines a message that contains unformatted text.
	PlainTextMessage *PlainTextMessageType `type:"structure"`

	// An object that defines a message that contains SSML-formatted text.
	SSMLMessage *SSMLMessageType `type:"structure"`
}

// String returns the string representation
func (s VoiceMessageContent) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s VoiceMessageContent) MarshalFields(e protocol.FieldEncoder) error {
	if s.CallInstructionsMessage != nil {
		v := s.CallInstructionsMessage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "CallInstructionsMessage", v, metadata)
	}
	if s.PlainTextMessage != nil {
		v := s.PlainTextMessage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "PlainTextMessage", v, metadata)
	}
	if s.SSMLMessage != nil {
		v := s.SSMLMessage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SSMLMessage", v, metadata)
	}
	return nil
}
