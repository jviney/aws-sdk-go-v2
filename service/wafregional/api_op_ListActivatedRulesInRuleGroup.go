// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListActivatedRulesInRuleGroupInput struct {
	_ struct{} `type:"structure"`

	// Specifies the number of ActivatedRules that you want AWS WAF to return for
	// this request. If you have more ActivatedRules than the number that you specify
	// for Limit, the response includes a NextMarker value that you can use to get
	// another batch of ActivatedRules.
	Limit *int64 `type:"integer"`

	// If you specify a value for Limit and you have more ActivatedRules than the
	// value of Limit, AWS WAF returns a NextMarker value in the response that allows
	// you to list another group of ActivatedRules. For the second and subsequent
	// ListActivatedRulesInRuleGroup requests, specify the value of NextMarker from
	// the previous response to get information about another batch of ActivatedRules.
	NextMarker *string `min:"1" type:"string"`

	// The RuleGroupId of the RuleGroup for which you want to get a list of ActivatedRule
	// objects.
	RuleGroupId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListActivatedRulesInRuleGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListActivatedRulesInRuleGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListActivatedRulesInRuleGroupInput"}
	if s.NextMarker != nil && len(*s.NextMarker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextMarker", 1))
	}
	if s.RuleGroupId != nil && len(*s.RuleGroupId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleGroupId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListActivatedRulesInRuleGroupOutput struct {
	_ struct{} `type:"structure"`

	// An array of ActivatedRules objects.
	ActivatedRules []ActivatedRule `type:"list"`

	// If you have more ActivatedRules than the number that you specified for Limit
	// in the request, the response includes a NextMarker value. To list more ActivatedRules,
	// submit another ListActivatedRulesInRuleGroup request, and specify the NextMarker
	// value from the response in the NextMarker value in the next request.
	NextMarker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListActivatedRulesInRuleGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opListActivatedRulesInRuleGroup = "ListActivatedRulesInRuleGroup"

// ListActivatedRulesInRuleGroupRequest returns a request value for making API operation for
// AWS WAF Regional.
//
//
// This is AWS WAF Classic documentation. For more information, see AWS WAF
// Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the AWS
// WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// With the latest version, AWS WAF has a single set of endpoints for regional
// and global use.
//
// Returns an array of ActivatedRule objects.
//
//    // Example sending a request using ListActivatedRulesInRuleGroupRequest.
//    req := client.ListActivatedRulesInRuleGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/ListActivatedRulesInRuleGroup
func (c *Client) ListActivatedRulesInRuleGroupRequest(input *ListActivatedRulesInRuleGroupInput) ListActivatedRulesInRuleGroupRequest {
	op := &aws.Operation{
		Name:       opListActivatedRulesInRuleGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListActivatedRulesInRuleGroupInput{}
	}

	req := c.newRequest(op, input, &ListActivatedRulesInRuleGroupOutput{})

	return ListActivatedRulesInRuleGroupRequest{Request: req, Input: input, Copy: c.ListActivatedRulesInRuleGroupRequest}
}

// ListActivatedRulesInRuleGroupRequest is the request type for the
// ListActivatedRulesInRuleGroup API operation.
type ListActivatedRulesInRuleGroupRequest struct {
	*aws.Request
	Input *ListActivatedRulesInRuleGroupInput
	Copy  func(*ListActivatedRulesInRuleGroupInput) ListActivatedRulesInRuleGroupRequest
}

// Send marshals and sends the ListActivatedRulesInRuleGroup API request.
func (r ListActivatedRulesInRuleGroupRequest) Send(ctx context.Context) (*ListActivatedRulesInRuleGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListActivatedRulesInRuleGroupResponse{
		ListActivatedRulesInRuleGroupOutput: r.Request.Data.(*ListActivatedRulesInRuleGroupOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListActivatedRulesInRuleGroupResponse is the response type for the
// ListActivatedRulesInRuleGroup API operation.
type ListActivatedRulesInRuleGroupResponse struct {
	*ListActivatedRulesInRuleGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListActivatedRulesInRuleGroup request.
func (r *ListActivatedRulesInRuleGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
