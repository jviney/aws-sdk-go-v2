// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package budgets

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Request of CreateSubscriber
type CreateSubscriberInput struct {
	_ struct{} `type:"structure"`

	// The accountId that is associated with the budget that you want to create
	// a subscriber for.
	//
	// AccountId is a required field
	AccountId *string `min:"12" type:"string" required:"true"`

	// The name of the budget that you want to subscribe to. Budget names must be
	// unique within an account.
	//
	// BudgetName is a required field
	BudgetName *string `min:"1" type:"string" required:"true"`

	// The notification that you want to create a subscriber for.
	//
	// Notification is a required field
	Notification *Notification `type:"structure" required:"true"`

	// The subscriber that you want to associate with a budget notification.
	//
	// Subscriber is a required field
	Subscriber *Subscriber `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateSubscriberInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSubscriberInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSubscriberInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}
	if s.AccountId != nil && len(*s.AccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountId", 12))
	}

	if s.BudgetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BudgetName"))
	}
	if s.BudgetName != nil && len(*s.BudgetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BudgetName", 1))
	}

	if s.Notification == nil {
		invalidParams.Add(aws.NewErrParamRequired("Notification"))
	}

	if s.Subscriber == nil {
		invalidParams.Add(aws.NewErrParamRequired("Subscriber"))
	}
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			invalidParams.AddNested("Notification", err.(aws.ErrInvalidParams))
		}
	}
	if s.Subscriber != nil {
		if err := s.Subscriber.Validate(); err != nil {
			invalidParams.AddNested("Subscriber", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response of CreateSubscriber
type CreateSubscriberOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateSubscriberOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateSubscriber = "CreateSubscriber"

// CreateSubscriberRequest returns a request value for making API operation for
// AWS Budgets.
//
// Creates a subscriber. You must create the associated budget and notification
// before you create the subscriber.
//
//    // Example sending a request using CreateSubscriberRequest.
//    req := client.CreateSubscriberRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateSubscriberRequest(input *CreateSubscriberInput) CreateSubscriberRequest {
	op := &aws.Operation{
		Name:       opCreateSubscriber,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSubscriberInput{}
	}

	req := c.newRequest(op, input, &CreateSubscriberOutput{})

	return CreateSubscriberRequest{Request: req, Input: input, Copy: c.CreateSubscriberRequest}
}

// CreateSubscriberRequest is the request type for the
// CreateSubscriber API operation.
type CreateSubscriberRequest struct {
	*aws.Request
	Input *CreateSubscriberInput
	Copy  func(*CreateSubscriberInput) CreateSubscriberRequest
}

// Send marshals and sends the CreateSubscriber API request.
func (r CreateSubscriberRequest) Send(ctx context.Context) (*CreateSubscriberResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSubscriberResponse{
		CreateSubscriberOutput: r.Request.Data.(*CreateSubscriberOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSubscriberResponse is the response type for the
// CreateSubscriber API operation.
type CreateSubscriberResponse struct {
	*CreateSubscriberOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSubscriber request.
func (r *CreateSubscriberResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
