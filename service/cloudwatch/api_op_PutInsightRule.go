// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type PutInsightRuleInput struct {
	_ struct{} `type:"structure"`

	// The definition of the rule, as a JSON object. For details on the valid syntax,
	// see Contributor Insights Rule Syntax (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/ContributorInsights-RuleSyntax.html).
	//
	// RuleDefinition is a required field
	RuleDefinition *string `min:"1" type:"string" required:"true"`

	// A unique name for the rule.
	//
	// RuleName is a required field
	RuleName *string `min:"1" type:"string" required:"true"`

	// The state of the rule. Valid values are ENABLED and DISABLED.
	RuleState *string `min:"1" type:"string"`

	// A list of key-value pairs to associate with the Contributor Insights rule.
	// You can associate as many as 50 tags with a rule.
	//
	// Tags can help you organize and categorize your resources. You can also use
	// them to scope user permissions, by granting a user permission to access or
	// change only the resources that have certain tag values.
	//
	// To be able to associate tags with a rule, you must have the cloudwatch:TagResource
	// permission in addition to the cloudwatch:PutInsightRule permission.
	//
	// If you are using this operation to update an existing Contributor Insights
	// rule, any tags you specify in this parameter are ignored. To change the tags
	// of an existing rule, use TagResource (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_TagResource.html).
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s PutInsightRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutInsightRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutInsightRuleInput"}

	if s.RuleDefinition == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleDefinition"))
	}
	if s.RuleDefinition != nil && len(*s.RuleDefinition) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleDefinition", 1))
	}

	if s.RuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleName"))
	}
	if s.RuleName != nil && len(*s.RuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleName", 1))
	}
	if s.RuleState != nil && len(*s.RuleState) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleState", 1))
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

type PutInsightRuleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutInsightRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutInsightRule = "PutInsightRule"

// PutInsightRuleRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// Creates a Contributor Insights rule. Rules evaluate log events in a CloudWatch
// Logs log group, enabling you to find contributor data for the log events
// in that log group. For more information, see Using Contributor Insights to
// Analyze High-Cardinality Data (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/ContributorInsights.html).
//
// If you create a rule, delete it, and then re-create it with the same name,
// historical data from the first time the rule was created may or may not be
// available.
//
//    // Example sending a request using PutInsightRuleRequest.
//    req := client.PutInsightRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/PutInsightRule
func (c *Client) PutInsightRuleRequest(input *PutInsightRuleInput) PutInsightRuleRequest {
	op := &aws.Operation{
		Name:       opPutInsightRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutInsightRuleInput{}
	}

	req := c.newRequest(op, input, &PutInsightRuleOutput{})

	return PutInsightRuleRequest{Request: req, Input: input, Copy: c.PutInsightRuleRequest}
}

// PutInsightRuleRequest is the request type for the
// PutInsightRule API operation.
type PutInsightRuleRequest struct {
	*aws.Request
	Input *PutInsightRuleInput
	Copy  func(*PutInsightRuleInput) PutInsightRuleRequest
}

// Send marshals and sends the PutInsightRule API request.
func (r PutInsightRuleRequest) Send(ctx context.Context) (*PutInsightRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutInsightRuleResponse{
		PutInsightRuleOutput: r.Request.Data.(*PutInsightRuleOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutInsightRuleResponse is the response type for the
// PutInsightRule API operation.
type PutInsightRuleResponse struct {
	*PutInsightRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutInsightRule request.
func (r *PutInsightRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
