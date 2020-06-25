// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DeleteEventSourceMappingInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the event source mapping.
	//
	// UUID is a required field
	UUID *string `location:"uri" locationName:"UUID" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteEventSourceMappingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteEventSourceMappingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteEventSourceMappingInput"}

	if s.UUID == nil {
		invalidParams.Add(aws.NewErrParamRequired("UUID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteEventSourceMappingInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.UUID != nil {
		v := *s.UUID

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "UUID", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// A mapping between an AWS resource and an AWS Lambda function. See CreateEventSourceMapping
// for details.
type DeleteEventSourceMappingOutput struct {
	_ struct{} `type:"structure"`

	// The maximum number of items to retrieve in a single batch.
	BatchSize *int64 `min:"1" type:"integer"`

	// (Streams) If the function returns an error, split the batch in two and retry.
	BisectBatchOnFunctionError *bool `type:"boolean"`

	// (Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded
	// records.
	DestinationConfig *DestinationConfig `type:"structure"`

	// The Amazon Resource Name (ARN) of the event source.
	EventSourceArn *string `type:"string"`

	// The ARN of the Lambda function.
	FunctionArn *string `type:"string"`

	// The date that the event source mapping was last updated, or its state changed.
	LastModified *time.Time `type:"timestamp"`

	// The result of the last AWS Lambda invocation of your Lambda function.
	LastProcessingResult *string `type:"string"`

	// (Streams) The maximum amount of time to gather records before invoking the
	// function, in seconds.
	MaximumBatchingWindowInSeconds *int64 `type:"integer"`

	// (Streams) The maximum age of a record that Lambda sends to a function for
	// processing.
	MaximumRecordAgeInSeconds *int64 `min:"60" type:"integer"`

	// (Streams) The maximum number of times to retry when the function returns
	// an error.
	MaximumRetryAttempts *int64 `type:"integer"`

	// (Streams) The number of batches to process from each shard concurrently.
	ParallelizationFactor *int64 `min:"1" type:"integer"`

	// The state of the event source mapping. It can be one of the following: Creating,
	// Enabling, Enabled, Disabling, Disabled, Updating, or Deleting.
	State *string `type:"string"`

	// Indicates whether the last change to the event source mapping was made by
	// a user, or by the Lambda service.
	StateTransitionReason *string `type:"string"`

	// The identifier of the event source mapping.
	UUID *string `type:"string"`
}

// String returns the string representation
func (s DeleteEventSourceMappingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteEventSourceMappingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BatchSize != nil {
		v := *s.BatchSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BatchSize", protocol.Int64Value(v), metadata)
	}
	if s.BisectBatchOnFunctionError != nil {
		v := *s.BisectBatchOnFunctionError

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BisectBatchOnFunctionError", protocol.BoolValue(v), metadata)
	}
	if s.DestinationConfig != nil {
		v := s.DestinationConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DestinationConfig", v, metadata)
	}
	if s.EventSourceArn != nil {
		v := *s.EventSourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EventSourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FunctionArn != nil {
		v := *s.FunctionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FunctionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastModified != nil {
		v := *s.LastModified

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastModified",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.LastProcessingResult != nil {
		v := *s.LastProcessingResult

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastProcessingResult", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaximumBatchingWindowInSeconds != nil {
		v := *s.MaximumBatchingWindowInSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaximumBatchingWindowInSeconds", protocol.Int64Value(v), metadata)
	}
	if s.MaximumRecordAgeInSeconds != nil {
		v := *s.MaximumRecordAgeInSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaximumRecordAgeInSeconds", protocol.Int64Value(v), metadata)
	}
	if s.MaximumRetryAttempts != nil {
		v := *s.MaximumRetryAttempts

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaximumRetryAttempts", protocol.Int64Value(v), metadata)
	}
	if s.ParallelizationFactor != nil {
		v := *s.ParallelizationFactor

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ParallelizationFactor", protocol.Int64Value(v), metadata)
	}
	if s.State != nil {
		v := *s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "State", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StateTransitionReason != nil {
		v := *s.StateTransitionReason

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StateTransitionReason", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UUID != nil {
		v := *s.UUID

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "UUID", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDeleteEventSourceMapping = "DeleteEventSourceMapping"

// DeleteEventSourceMappingRequest returns a request value for making API operation for
// AWS Lambda.
//
// Deletes an event source mapping (https://docs.aws.amazon.com/lambda/latest/dg/intro-invocation-modes.html).
// You can get the identifier of a mapping from the output of ListEventSourceMappings.
//
// When you delete an event source mapping, it enters a Deleting state and might
// not be completely deleted for several seconds.
//
//    // Example sending a request using DeleteEventSourceMappingRequest.
//    req := client.DeleteEventSourceMappingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/DeleteEventSourceMapping
func (c *Client) DeleteEventSourceMappingRequest(input *DeleteEventSourceMappingInput) DeleteEventSourceMappingRequest {
	op := &aws.Operation{
		Name:       opDeleteEventSourceMapping,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2015-03-31/event-source-mappings/{UUID}",
	}

	if input == nil {
		input = &DeleteEventSourceMappingInput{}
	}

	req := c.newRequest(op, input, &DeleteEventSourceMappingOutput{})

	return DeleteEventSourceMappingRequest{Request: req, Input: input, Copy: c.DeleteEventSourceMappingRequest}
}

// DeleteEventSourceMappingRequest is the request type for the
// DeleteEventSourceMapping API operation.
type DeleteEventSourceMappingRequest struct {
	*aws.Request
	Input *DeleteEventSourceMappingInput
	Copy  func(*DeleteEventSourceMappingInput) DeleteEventSourceMappingRequest
}

// Send marshals and sends the DeleteEventSourceMapping API request.
func (r DeleteEventSourceMappingRequest) Send(ctx context.Context) (*DeleteEventSourceMappingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteEventSourceMappingResponse{
		DeleteEventSourceMappingOutput: r.Request.Data.(*DeleteEventSourceMappingOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteEventSourceMappingResponse is the response type for the
// DeleteEventSourceMapping API operation.
type DeleteEventSourceMappingResponse struct {
	*DeleteEventSourceMappingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteEventSourceMapping request.
func (r *DeleteEventSourceMappingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
