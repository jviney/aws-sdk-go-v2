// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type PutRetentionSettingsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The retention settings.
	//
	// RetentionSettings is a required field
	RetentionSettings *RetentionSettings `type:"structure" required:"true"`
}

// String returns the string representation
func (s PutRetentionSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutRetentionSettingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutRetentionSettingsInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.RetentionSettings == nil {
		invalidParams.Add(aws.NewErrParamRequired("RetentionSettings"))
	}
	if s.RetentionSettings != nil {
		if err := s.RetentionSettings.Validate(); err != nil {
			invalidParams.AddNested("RetentionSettings", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutRetentionSettingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RetentionSettings != nil {
		v := s.RetentionSettings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "RetentionSettings", v, metadata)
	}
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type PutRetentionSettingsOutput struct {
	_ struct{} `type:"structure"`

	// The timestamp representing the time at which the specified items are permanently
	// deleted, in ISO 8601 format.
	InitiateDeletionTimestamp *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The retention settings.
	RetentionSettings *RetentionSettings `type:"structure"`
}

// String returns the string representation
func (s PutRetentionSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutRetentionSettingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.InitiateDeletionTimestamp != nil {
		v := *s.InitiateDeletionTimestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "InitiateDeletionTimestamp",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.RetentionSettings != nil {
		v := s.RetentionSettings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "RetentionSettings", v, metadata)
	}
	return nil
}

const opPutRetentionSettings = "PutRetentionSettings"

// PutRetentionSettingsRequest returns a request value for making API operation for
// Amazon Chime.
//
// Puts retention settings for the specified Amazon Chime Enterprise account.
// We recommend using AWS CloudTrail to monitor usage of this API for your account.
// For more information, see Logging Amazon Chime API Calls with AWS CloudTrail
// (https://docs.aws.amazon.com/chime/latest/ag/cloudtrail.html) in the Amazon
// Chime Administration Guide.
//
// To turn off existing retention settings, remove the number of days from the
// corresponding RetentionDays field in the RetentionSettings object. For more
// information about retention settings, see Managing Chat Retention Policies
// (https://docs.aws.amazon.com/chime/latest/ag/chat-retention.html) in the
// Amazon Chime Administration Guide.
//
//    // Example sending a request using PutRetentionSettingsRequest.
//    req := client.PutRetentionSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/PutRetentionSettings
func (c *Client) PutRetentionSettingsRequest(input *PutRetentionSettingsInput) PutRetentionSettingsRequest {
	op := &aws.Operation{
		Name:       opPutRetentionSettings,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{accountId}/retention-settings",
	}

	if input == nil {
		input = &PutRetentionSettingsInput{}
	}

	req := c.newRequest(op, input, &PutRetentionSettingsOutput{})

	return PutRetentionSettingsRequest{Request: req, Input: input, Copy: c.PutRetentionSettingsRequest}
}

// PutRetentionSettingsRequest is the request type for the
// PutRetentionSettings API operation.
type PutRetentionSettingsRequest struct {
	*aws.Request
	Input *PutRetentionSettingsInput
	Copy  func(*PutRetentionSettingsInput) PutRetentionSettingsRequest
}

// Send marshals and sends the PutRetentionSettings API request.
func (r PutRetentionSettingsRequest) Send(ctx context.Context) (*PutRetentionSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutRetentionSettingsResponse{
		PutRetentionSettingsOutput: r.Request.Data.(*PutRetentionSettingsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutRetentionSettingsResponse is the response type for the
// PutRetentionSettings API operation.
type PutRetentionSettingsResponse struct {
	*PutRetentionSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutRetentionSettings request.
func (r *PutRetentionSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
