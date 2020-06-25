// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package fms

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type PutNotificationChannelInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the IAM role that allows Amazon SNS to
	// record AWS Firewall Manager activity.
	//
	// SnsRoleName is a required field
	SnsRoleName *string `min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the SNS topic that collects notifications
	// from AWS Firewall Manager.
	//
	// SnsTopicArn is a required field
	SnsTopicArn *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutNotificationChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutNotificationChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutNotificationChannelInput"}

	if s.SnsRoleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnsRoleName"))
	}
	if s.SnsRoleName != nil && len(*s.SnsRoleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SnsRoleName", 1))
	}

	if s.SnsTopicArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnsTopicArn"))
	}
	if s.SnsTopicArn != nil && len(*s.SnsTopicArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SnsTopicArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutNotificationChannelOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutNotificationChannelOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutNotificationChannel = "PutNotificationChannel"

// PutNotificationChannelRequest returns a request value for making API operation for
// Firewall Management Service.
//
// Designates the IAM role and Amazon Simple Notification Service (SNS) topic
// that AWS Firewall Manager uses to record SNS logs.
//
//    // Example sending a request using PutNotificationChannelRequest.
//    req := client.PutNotificationChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/fms-2018-01-01/PutNotificationChannel
func (c *Client) PutNotificationChannelRequest(input *PutNotificationChannelInput) PutNotificationChannelRequest {
	op := &aws.Operation{
		Name:       opPutNotificationChannel,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutNotificationChannelInput{}
	}

	req := c.newRequest(op, input, &PutNotificationChannelOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return PutNotificationChannelRequest{Request: req, Input: input, Copy: c.PutNotificationChannelRequest}
}

// PutNotificationChannelRequest is the request type for the
// PutNotificationChannel API operation.
type PutNotificationChannelRequest struct {
	*aws.Request
	Input *PutNotificationChannelInput
	Copy  func(*PutNotificationChannelInput) PutNotificationChannelRequest
}

// Send marshals and sends the PutNotificationChannel API request.
func (r PutNotificationChannelRequest) Send(ctx context.Context) (*PutNotificationChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutNotificationChannelResponse{
		PutNotificationChannelOutput: r.Request.Data.(*PutNotificationChannelOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutNotificationChannelResponse is the response type for the
// PutNotificationChannel API operation.
type PutNotificationChannelResponse struct {
	*PutNotificationChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutNotificationChannel request.
func (r *PutNotificationChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
