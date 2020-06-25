// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DeleteScheduledAuditInput struct {
	_ struct{} `type:"structure"`

	// The name of the scheduled audit you want to delete.
	//
	// ScheduledAuditName is a required field
	ScheduledAuditName *string `location:"uri" locationName:"scheduledAuditName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteScheduledAuditInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteScheduledAuditInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteScheduledAuditInput"}

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
func (s DeleteScheduledAuditInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ScheduledAuditName != nil {
		v := *s.ScheduledAuditName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "scheduledAuditName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteScheduledAuditOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteScheduledAuditOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteScheduledAuditOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteScheduledAudit = "DeleteScheduledAudit"

// DeleteScheduledAuditRequest returns a request value for making API operation for
// AWS IoT.
//
// Deletes a scheduled audit.
//
//    // Example sending a request using DeleteScheduledAuditRequest.
//    req := client.DeleteScheduledAuditRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteScheduledAuditRequest(input *DeleteScheduledAuditInput) DeleteScheduledAuditRequest {
	op := &aws.Operation{
		Name:       opDeleteScheduledAudit,
		HTTPMethod: "DELETE",
		HTTPPath:   "/audit/scheduledaudits/{scheduledAuditName}",
	}

	if input == nil {
		input = &DeleteScheduledAuditInput{}
	}

	req := c.newRequest(op, input, &DeleteScheduledAuditOutput{})

	return DeleteScheduledAuditRequest{Request: req, Input: input, Copy: c.DeleteScheduledAuditRequest}
}

// DeleteScheduledAuditRequest is the request type for the
// DeleteScheduledAudit API operation.
type DeleteScheduledAuditRequest struct {
	*aws.Request
	Input *DeleteScheduledAuditInput
	Copy  func(*DeleteScheduledAuditInput) DeleteScheduledAuditRequest
}

// Send marshals and sends the DeleteScheduledAudit API request.
func (r DeleteScheduledAuditRequest) Send(ctx context.Context) (*DeleteScheduledAuditResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteScheduledAuditResponse{
		DeleteScheduledAuditOutput: r.Request.Data.(*DeleteScheduledAuditOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteScheduledAuditResponse is the response type for the
// DeleteScheduledAudit API operation.
type DeleteScheduledAuditResponse struct {
	*DeleteScheduledAuditOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteScheduledAudit request.
func (r *DeleteScheduledAuditResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
