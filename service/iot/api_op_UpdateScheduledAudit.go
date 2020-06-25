// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type UpdateScheduledAuditInput struct {
	_ struct{} `type:"structure"`

	// The day of the month on which the scheduled audit takes place. Can be "1"
	// through "31" or "LAST". This field is required if the "frequency" parameter
	// is set to "MONTHLY". If days 29-31 are specified, and the month does not
	// have that many days, the audit takes place on the "LAST" day of the month.
	DayOfMonth *string `locationName:"dayOfMonth" type:"string"`

	// The day of the week on which the scheduled audit takes place. Can be one
	// of "SUN", "MON", "TUE", "WED", "THU", "FRI", or "SAT". This field is required
	// if the "frequency" parameter is set to "WEEKLY" or "BIWEEKLY".
	DayOfWeek DayOfWeek `locationName:"dayOfWeek" type:"string" enum:"true"`

	// How often the scheduled audit takes place. Can be one of "DAILY", "WEEKLY",
	// "BIWEEKLY", or "MONTHLY". The start time of each audit is determined by the
	// system.
	Frequency AuditFrequency `locationName:"frequency" type:"string" enum:"true"`

	// The name of the scheduled audit. (Max. 128 chars)
	//
	// ScheduledAuditName is a required field
	ScheduledAuditName *string `location:"uri" locationName:"scheduledAuditName" min:"1" type:"string" required:"true"`

	// Which checks are performed during the scheduled audit. Checks must be enabled
	// for your account. (Use DescribeAccountAuditConfiguration to see the list
	// of all checks, including those that are enabled or use UpdateAccountAuditConfiguration
	// to select which checks are enabled.)
	TargetCheckNames []string `locationName:"targetCheckNames" type:"list"`
}

// String returns the string representation
func (s UpdateScheduledAuditInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateScheduledAuditInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateScheduledAuditInput"}

	if s.ScheduledAuditName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScheduledAuditName"))
	}
	if s.ScheduledAuditName != nil && len(*s.ScheduledAuditName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ScheduledAuditName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateScheduledAuditInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DayOfMonth != nil {
		v := *s.DayOfMonth

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dayOfMonth", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.DayOfWeek) > 0 {
		v := s.DayOfWeek

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dayOfWeek", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.Frequency) > 0 {
		v := s.Frequency

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "frequency", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.TargetCheckNames != nil {
		v := s.TargetCheckNames

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "targetCheckNames", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.ScheduledAuditName != nil {
		v := *s.ScheduledAuditName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "scheduledAuditName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateScheduledAuditOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the scheduled audit.
	ScheduledAuditArn *string `locationName:"scheduledAuditArn" type:"string"`
}

// String returns the string representation
func (s UpdateScheduledAuditOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateScheduledAuditOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ScheduledAuditArn != nil {
		v := *s.ScheduledAuditArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "scheduledAuditArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateScheduledAudit = "UpdateScheduledAudit"

// UpdateScheduledAuditRequest returns a request value for making API operation for
// AWS IoT.
//
// Updates a scheduled audit, including which checks are performed and how often
// the audit takes place.
//
//    // Example sending a request using UpdateScheduledAuditRequest.
//    req := client.UpdateScheduledAuditRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateScheduledAuditRequest(input *UpdateScheduledAuditInput) UpdateScheduledAuditRequest {
	op := &aws.Operation{
		Name:       opUpdateScheduledAudit,
		HTTPMethod: "PATCH",
		HTTPPath:   "/audit/scheduledaudits/{scheduledAuditName}",
	}

	if input == nil {
		input = &UpdateScheduledAuditInput{}
	}

	req := c.newRequest(op, input, &UpdateScheduledAuditOutput{})

	return UpdateScheduledAuditRequest{Request: req, Input: input, Copy: c.UpdateScheduledAuditRequest}
}

// UpdateScheduledAuditRequest is the request type for the
// UpdateScheduledAudit API operation.
type UpdateScheduledAuditRequest struct {
	*aws.Request
	Input *UpdateScheduledAuditInput
	Copy  func(*UpdateScheduledAuditInput) UpdateScheduledAuditRequest
}

// Send marshals and sends the UpdateScheduledAudit API request.
func (r UpdateScheduledAuditRequest) Send(ctx context.Context) (*UpdateScheduledAuditResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateScheduledAuditResponse{
		UpdateScheduledAuditOutput: r.Request.Data.(*UpdateScheduledAuditOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateScheduledAuditResponse is the response type for the
// UpdateScheduledAudit API operation.
type UpdateScheduledAuditResponse struct {
	*UpdateScheduledAuditOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateScheduledAudit request.
func (r *UpdateScheduledAuditResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
