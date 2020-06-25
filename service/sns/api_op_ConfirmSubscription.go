// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Input for ConfirmSubscription action.
type ConfirmSubscriptionInput struct {
	_ struct{} `type:"structure"`

	// Disallows unauthenticated unsubscribes of the subscription. If the value
	// of this parameter is true and the request has an AWS signature, then only
	// the topic owner and the subscription owner can unsubscribe the endpoint.
	// The unsubscribe action requires AWS authentication.
	AuthenticateOnUnsubscribe *string `type:"string"`

	// Short-lived token sent to an endpoint during the Subscribe action.
	//
	// Token is a required field
	Token *string `type:"string" required:"true"`

	// The ARN of the topic for which you wish to confirm a subscription.
	//
	// TopicArn is a required field
	TopicArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ConfirmSubscriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ConfirmSubscriptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ConfirmSubscriptionInput"}

	if s.Token == nil {
		invalidParams.Add(aws.NewErrParamRequired("Token"))
	}

	if s.TopicArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("TopicArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response for ConfirmSubscriptions action.
type ConfirmSubscriptionOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the created subscription.
	SubscriptionArn *string `type:"string"`
}

// String returns the string representation
func (s ConfirmSubscriptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opConfirmSubscription = "ConfirmSubscription"

// ConfirmSubscriptionRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Verifies an endpoint owner's intent to receive messages by validating the
// token sent to the endpoint by an earlier Subscribe action. If the token is
// valid, the action creates a new subscription and returns its Amazon Resource
// Name (ARN). This call requires an AWS signature only when the AuthenticateOnUnsubscribe
// flag is set to "true".
//
//    // Example sending a request using ConfirmSubscriptionRequest.
//    req := client.ConfirmSubscriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/ConfirmSubscription
func (c *Client) ConfirmSubscriptionRequest(input *ConfirmSubscriptionInput) ConfirmSubscriptionRequest {
	op := &aws.Operation{
		Name:       opConfirmSubscription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ConfirmSubscriptionInput{}
	}

	req := c.newRequest(op, input, &ConfirmSubscriptionOutput{})

	return ConfirmSubscriptionRequest{Request: req, Input: input, Copy: c.ConfirmSubscriptionRequest}
}

// ConfirmSubscriptionRequest is the request type for the
// ConfirmSubscription API operation.
type ConfirmSubscriptionRequest struct {
	*aws.Request
	Input *ConfirmSubscriptionInput
	Copy  func(*ConfirmSubscriptionInput) ConfirmSubscriptionRequest
}

// Send marshals and sends the ConfirmSubscription API request.
func (r ConfirmSubscriptionRequest) Send(ctx context.Context) (*ConfirmSubscriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ConfirmSubscriptionResponse{
		ConfirmSubscriptionOutput: r.Request.Data.(*ConfirmSubscriptionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ConfirmSubscriptionResponse is the response type for the
// ConfirmSubscription API operation.
type ConfirmSubscriptionResponse struct {
	*ConfirmSubscriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ConfirmSubscription request.
func (r *ConfirmSubscriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
