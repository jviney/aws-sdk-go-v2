// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DescribeMitigationActionInput struct {
	_ struct{} `type:"structure"`

	// The friendly name that uniquely identifies the mitigation action.
	//
	// ActionName is a required field
	ActionName *string `location:"uri" locationName:"actionName" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeMitigationActionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeMitigationActionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeMitigationActionInput"}

	if s.ActionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActionName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeMitigationActionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ActionName != nil {
		v := *s.ActionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "actionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeMitigationActionOutput struct {
	_ struct{} `type:"structure"`

	// The ARN that identifies this migration action.
	ActionArn *string `locationName:"actionArn" type:"string"`

	// A unique identifier for this action.
	ActionId *string `locationName:"actionId" type:"string"`

	// The friendly name that uniquely identifies the mitigation action.
	ActionName *string `locationName:"actionName" type:"string"`

	// Parameters that control how the mitigation action is applied, specific to
	// the type of mitigation action.
	ActionParams *MitigationActionParams `locationName:"actionParams" type:"structure"`

	// The type of mitigation action.
	ActionType MitigationActionType `locationName:"actionType" type:"string" enum:"true"`

	// The date and time when the mitigation action was added to your AWS account.
	CreationDate *time.Time `locationName:"creationDate" type:"timestamp"`

	// The date and time when the mitigation action was last changed.
	LastModifiedDate *time.Time `locationName:"lastModifiedDate" type:"timestamp"`

	// The ARN of the IAM role used to apply this action.
	RoleArn *string `locationName:"roleArn" min:"20" type:"string"`
}

// String returns the string representation
func (s DescribeMitigationActionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeMitigationActionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ActionArn != nil {
		v := *s.ActionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "actionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ActionId != nil {
		v := *s.ActionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "actionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ActionName != nil {
		v := *s.ActionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "actionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ActionParams != nil {
		v := s.ActionParams

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "actionParams", v, metadata)
	}
	if len(s.ActionType) > 0 {
		v := s.ActionType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "actionType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.CreationDate != nil {
		v := *s.CreationDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "creationDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.LastModifiedDate != nil {
		v := *s.LastModifiedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastModifiedDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.RoleArn != nil {
		v := *s.RoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "roleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeMitigationAction = "DescribeMitigationAction"

// DescribeMitigationActionRequest returns a request value for making API operation for
// AWS IoT.
//
// Gets information about a mitigation action.
//
//    // Example sending a request using DescribeMitigationActionRequest.
//    req := client.DescribeMitigationActionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeMitigationActionRequest(input *DescribeMitigationActionInput) DescribeMitigationActionRequest {
	op := &aws.Operation{
		Name:       opDescribeMitigationAction,
		HTTPMethod: "GET",
		HTTPPath:   "/mitigationactions/actions/{actionName}",
	}

	if input == nil {
		input = &DescribeMitigationActionInput{}
	}

	req := c.newRequest(op, input, &DescribeMitigationActionOutput{})

	return DescribeMitigationActionRequest{Request: req, Input: input, Copy: c.DescribeMitigationActionRequest}
}

// DescribeMitigationActionRequest is the request type for the
// DescribeMitigationAction API operation.
type DescribeMitigationActionRequest struct {
	*aws.Request
	Input *DescribeMitigationActionInput
	Copy  func(*DescribeMitigationActionInput) DescribeMitigationActionRequest
}

// Send marshals and sends the DescribeMitigationAction API request.
func (r DescribeMitigationActionRequest) Send(ctx context.Context) (*DescribeMitigationActionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeMitigationActionResponse{
		DescribeMitigationActionOutput: r.Request.Data.(*DescribeMitigationActionOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeMitigationActionResponse is the response type for the
// DescribeMitigationAction API operation.
type DescribeMitigationActionResponse struct {
	*DescribeMitigationActionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeMitigationAction request.
func (r *DescribeMitigationActionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
