// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type UpdateResolverRuleInput struct {
	_ struct{} `type:"structure"`

	// The new settings for the resolver rule.
	//
	// Config is a required field
	Config *ResolverRuleConfig `type:"structure" required:"true"`

	// The ID of the resolver rule that you want to update.
	//
	// ResolverRuleId is a required field
	ResolverRuleId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateResolverRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateResolverRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateResolverRuleInput"}

	if s.Config == nil {
		invalidParams.Add(aws.NewErrParamRequired("Config"))
	}

	if s.ResolverRuleId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResolverRuleId"))
	}
	if s.ResolverRuleId != nil && len(*s.ResolverRuleId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResolverRuleId", 1))
	}
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			invalidParams.AddNested("Config", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateResolverRuleOutput struct {
	_ struct{} `type:"structure"`

	// The response to an UpdateResolverRule request.
	ResolverRule *ResolverRule `type:"structure"`
}

// String returns the string representation
func (s UpdateResolverRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateResolverRule = "UpdateResolverRule"

// UpdateResolverRuleRequest returns a request value for making API operation for
// Amazon Route 53 Resolver.
//
// Updates settings for a specified resolver rule. ResolverRuleId is required,
// and all other parameters are optional. If you don't specify a parameter,
// it retains its current value.
//
//    // Example sending a request using UpdateResolverRuleRequest.
//    req := client.UpdateResolverRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/UpdateResolverRule
func (c *Client) UpdateResolverRuleRequest(input *UpdateResolverRuleInput) UpdateResolverRuleRequest {
	op := &aws.Operation{
		Name:       opUpdateResolverRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateResolverRuleInput{}
	}

	req := c.newRequest(op, input, &UpdateResolverRuleOutput{})

	return UpdateResolverRuleRequest{Request: req, Input: input, Copy: c.UpdateResolverRuleRequest}
}

// UpdateResolverRuleRequest is the request type for the
// UpdateResolverRule API operation.
type UpdateResolverRuleRequest struct {
	*aws.Request
	Input *UpdateResolverRuleInput
	Copy  func(*UpdateResolverRuleInput) UpdateResolverRuleRequest
}

// Send marshals and sends the UpdateResolverRule API request.
func (r UpdateResolverRuleRequest) Send(ctx context.Context) (*UpdateResolverRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateResolverRuleResponse{
		UpdateResolverRuleOutput: r.Request.Data.(*UpdateResolverRuleOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateResolverRuleResponse is the response type for the
// UpdateResolverRule API operation.
type UpdateResolverRuleResponse struct {
	*UpdateResolverRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateResolverRule request.
func (r *UpdateResolverRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
