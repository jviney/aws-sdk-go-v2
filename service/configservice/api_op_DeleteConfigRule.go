// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteConfigRuleInput struct {
	_ struct{} `type:"structure"`

	// The name of the AWS Config rule that you want to delete.
	//
	// ConfigRuleName is a required field
	ConfigRuleName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteConfigRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteConfigRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteConfigRuleInput"}

	if s.ConfigRuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigRuleName"))
	}
	if s.ConfigRuleName != nil && len(*s.ConfigRuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConfigRuleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteConfigRuleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteConfigRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteConfigRule = "DeleteConfigRule"

// DeleteConfigRuleRequest returns a request value for making API operation for
// AWS Config.
//
// Deletes the specified AWS Config rule and all of its evaluation results.
//
// AWS Config sets the state of a rule to DELETING until the deletion is complete.
// You cannot update a rule while it is in this state. If you make a PutConfigRule
// or DeleteConfigRule request for the rule, you will receive a ResourceInUseException.
//
// You can check the state of a rule by using the DescribeConfigRules request.
//
//    // Example sending a request using DeleteConfigRuleRequest.
//    req := client.DeleteConfigRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DeleteConfigRule
func (c *Client) DeleteConfigRuleRequest(input *DeleteConfigRuleInput) DeleteConfigRuleRequest {
	op := &aws.Operation{
		Name:       opDeleteConfigRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteConfigRuleInput{}
	}

	req := c.newRequest(op, input, &DeleteConfigRuleOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteConfigRuleRequest{Request: req, Input: input, Copy: c.DeleteConfigRuleRequest}
}

// DeleteConfigRuleRequest is the request type for the
// DeleteConfigRule API operation.
type DeleteConfigRuleRequest struct {
	*aws.Request
	Input *DeleteConfigRuleInput
	Copy  func(*DeleteConfigRuleInput) DeleteConfigRuleRequest
}

// Send marshals and sends the DeleteConfigRule API request.
func (r DeleteConfigRuleRequest) Send(ctx context.Context) (*DeleteConfigRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteConfigRuleResponse{
		DeleteConfigRuleOutput: r.Request.Data.(*DeleteConfigRuleOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteConfigRuleResponse is the response type for the
// DeleteConfigRule API operation.
type DeleteConfigRuleResponse struct {
	*DeleteConfigRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteConfigRule request.
func (r *DeleteConfigRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
