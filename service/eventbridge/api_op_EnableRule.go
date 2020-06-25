// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eventbridge

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type EnableRuleInput struct {
	_ struct{} `type:"structure"`

	// The event bus associated with the rule. If you omit this, the default event
	// bus is used.
	EventBusName *string `min:"1" type:"string"`

	// The name of the rule.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s EnableRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EnableRuleInput"}
	if s.EventBusName != nil && len(*s.EventBusName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventBusName", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type EnableRuleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EnableRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opEnableRule = "EnableRule"

// EnableRuleRequest returns a request value for making API operation for
// Amazon EventBridge.
//
// Enables the specified rule. If the rule does not exist, the operation fails.
//
// When you enable a rule, incoming events might not immediately start matching
// to a newly enabled rule. Allow a short period of time for changes to take
// effect.
//
//    // Example sending a request using EnableRuleRequest.
//    req := client.EnableRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/EnableRule
func (c *Client) EnableRuleRequest(input *EnableRuleInput) EnableRuleRequest {
	op := &aws.Operation{
		Name:       opEnableRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableRuleInput{}
	}

	req := c.newRequest(op, input, &EnableRuleOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return EnableRuleRequest{Request: req, Input: input, Copy: c.EnableRuleRequest}
}

// EnableRuleRequest is the request type for the
// EnableRule API operation.
type EnableRuleRequest struct {
	*aws.Request
	Input *EnableRuleInput
	Copy  func(*EnableRuleInput) EnableRuleRequest
}

// Send marshals and sends the EnableRule API request.
func (r EnableRuleRequest) Send(ctx context.Context) (*EnableRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableRuleResponse{
		EnableRuleOutput: r.Request.Data.(*EnableRuleOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableRuleResponse is the response type for the
// EnableRule API operation.
type EnableRuleResponse struct {
	*EnableRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableRule request.
func (r *EnableRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
