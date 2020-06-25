// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetRuleGroupInput struct {
	_ struct{} `type:"structure"`

	// The RuleGroupId of the RuleGroup that you want to get. RuleGroupId is returned
	// by CreateRuleGroup and by ListRuleGroups.
	//
	// RuleGroupId is a required field
	RuleGroupId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRuleGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRuleGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRuleGroupInput"}

	if s.RuleGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleGroupId"))
	}
	if s.RuleGroupId != nil && len(*s.RuleGroupId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleGroupId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetRuleGroupOutput struct {
	_ struct{} `type:"structure"`

	// Information about the RuleGroup that you specified in the GetRuleGroup request.
	RuleGroup *RuleGroup `type:"structure"`
}

// String returns the string representation
func (s GetRuleGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRuleGroup = "GetRuleGroup"

// GetRuleGroupRequest returns a request value for making API operation for
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
// Returns the RuleGroup that is specified by the RuleGroupId that you included
// in the GetRuleGroup request.
//
// To view the rules in a rule group, use ListActivatedRulesInRuleGroup.
//
//    // Example sending a request using GetRuleGroupRequest.
//    req := client.GetRuleGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetRuleGroup
func (c *Client) GetRuleGroupRequest(input *GetRuleGroupInput) GetRuleGroupRequest {
	op := &aws.Operation{
		Name:       opGetRuleGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRuleGroupInput{}
	}

	req := c.newRequest(op, input, &GetRuleGroupOutput{})

	return GetRuleGroupRequest{Request: req, Input: input, Copy: c.GetRuleGroupRequest}
}

// GetRuleGroupRequest is the request type for the
// GetRuleGroup API operation.
type GetRuleGroupRequest struct {
	*aws.Request
	Input *GetRuleGroupInput
	Copy  func(*GetRuleGroupInput) GetRuleGroupRequest
}

// Send marshals and sends the GetRuleGroup API request.
func (r GetRuleGroupRequest) Send(ctx context.Context) (*GetRuleGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRuleGroupResponse{
		GetRuleGroupOutput: r.Request.Data.(*GetRuleGroupOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRuleGroupResponse is the response type for the
// GetRuleGroup API operation.
type GetRuleGroupResponse struct {
	*GetRuleGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRuleGroup request.
func (r *GetRuleGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
