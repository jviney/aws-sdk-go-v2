// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/query"
)

// Represents a request to enable or disable the email sending capabilities
// for your entire Amazon SES account.
type UpdateAccountSendingEnabledInput struct {
	_ struct{} `type:"structure"`

	// Describes whether email sending is enabled or disabled for your Amazon SES
	// account in the current AWS Region.
	Enabled *bool `type:"boolean"`
}

// String returns the string representation
func (s UpdateAccountSendingEnabledInput) String() string {
	return awsutil.Prettify(s)
}

type UpdateAccountSendingEnabledOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateAccountSendingEnabledOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateAccountSendingEnabled = "UpdateAccountSendingEnabled"

// UpdateAccountSendingEnabledRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Enables or disables email sending across your entire Amazon SES account in
// the current AWS Region. You can use this operation in conjunction with Amazon
// CloudWatch alarms to temporarily pause email sending across your Amazon SES
// account in a given AWS Region when reputation metrics (such as your bounce
// or complaint rates) reach certain thresholds.
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using UpdateAccountSendingEnabledRequest.
//    req := client.UpdateAccountSendingEnabledRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/UpdateAccountSendingEnabled
func (c *Client) UpdateAccountSendingEnabledRequest(input *UpdateAccountSendingEnabledInput) UpdateAccountSendingEnabledRequest {
	op := &aws.Operation{
		Name:       opUpdateAccountSendingEnabled,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateAccountSendingEnabledInput{}
	}

	req := c.newRequest(op, input, &UpdateAccountSendingEnabledOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return UpdateAccountSendingEnabledRequest{Request: req, Input: input, Copy: c.UpdateAccountSendingEnabledRequest}
}

// UpdateAccountSendingEnabledRequest is the request type for the
// UpdateAccountSendingEnabled API operation.
type UpdateAccountSendingEnabledRequest struct {
	*aws.Request
	Input *UpdateAccountSendingEnabledInput
	Copy  func(*UpdateAccountSendingEnabledInput) UpdateAccountSendingEnabledRequest
}

// Send marshals and sends the UpdateAccountSendingEnabled API request.
func (r UpdateAccountSendingEnabledRequest) Send(ctx context.Context) (*UpdateAccountSendingEnabledResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAccountSendingEnabledResponse{
		UpdateAccountSendingEnabledOutput: r.Request.Data.(*UpdateAccountSendingEnabledOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAccountSendingEnabledResponse is the response type for the
// UpdateAccountSendingEnabled API operation.
type UpdateAccountSendingEnabledResponse struct {
	*UpdateAccountSendingEnabledOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAccountSendingEnabled request.
func (r *UpdateAccountSendingEnabledResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
