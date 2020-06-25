// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eventbridge

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateEventBusInput struct {
	_ struct{} `type:"structure"`

	// If you are creating a partner event bus, this specifies the partner event
	// source that the new event bus will be matched with.
	EventSourceName *string `min:"1" type:"string"`

	// The name of the new event bus.
	//
	// Event bus names cannot contain the / character. You can't use the name default
	// for a custom event bus, as this name is already used for your account's default
	// event bus.
	//
	// If this is a partner event bus, the name must exactly match the name of the
	// partner event source that this event bus is matched to.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// Tags to associate with the event bus.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateEventBusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateEventBusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateEventBusInput"}
	if s.EventSourceName != nil && len(*s.EventSourceName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventSourceName", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateEventBusOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the new event bus.
	EventBusArn *string `type:"string"`
}

// String returns the string representation
func (s CreateEventBusOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateEventBus = "CreateEventBus"

// CreateEventBusRequest returns a request value for making API operation for
// Amazon EventBridge.
//
// Creates a new event bus within your account. This can be a custom event bus
// which you can use to receive events from your custom applications and services,
// or it can be a partner event bus which can be matched to a partner event
// source.
//
//    // Example sending a request using CreateEventBusRequest.
//    req := client.CreateEventBusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/CreateEventBus
func (c *Client) CreateEventBusRequest(input *CreateEventBusInput) CreateEventBusRequest {
	op := &aws.Operation{
		Name:       opCreateEventBus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateEventBusInput{}
	}

	req := c.newRequest(op, input, &CreateEventBusOutput{})

	return CreateEventBusRequest{Request: req, Input: input, Copy: c.CreateEventBusRequest}
}

// CreateEventBusRequest is the request type for the
// CreateEventBus API operation.
type CreateEventBusRequest struct {
	*aws.Request
	Input *CreateEventBusInput
	Copy  func(*CreateEventBusInput) CreateEventBusRequest
}

// Send marshals and sends the CreateEventBus API request.
func (r CreateEventBusRequest) Send(ctx context.Context) (*CreateEventBusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateEventBusResponse{
		CreateEventBusOutput: r.Request.Data.(*CreateEventBusOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateEventBusResponse is the response type for the
// CreateEventBus API operation.
type CreateEventBusResponse struct {
	*CreateEventBusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateEventBus request.
func (r *CreateEventBusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
