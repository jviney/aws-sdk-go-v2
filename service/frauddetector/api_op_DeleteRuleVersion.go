// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package frauddetector

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteRuleVersionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the detector that includes the rule version to delete.
	//
	// DetectorId is a required field
	DetectorId *string `locationName:"detectorId" min:"1" type:"string" required:"true"`

	// The rule ID of the rule version to delete.
	//
	// RuleId is a required field
	RuleId *string `locationName:"ruleId" min:"1" type:"string" required:"true"`

	// The rule version to delete.
	//
	// RuleVersion is a required field
	RuleVersion *string `locationName:"ruleVersion" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRuleVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRuleVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRuleVersionInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if s.RuleId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleId"))
	}
	if s.RuleId != nil && len(*s.RuleId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleId", 1))
	}

	if s.RuleVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleVersion"))
	}
	if s.RuleVersion != nil && len(*s.RuleVersion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleVersion", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteRuleVersionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRuleVersionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteRuleVersion = "DeleteRuleVersion"

// DeleteRuleVersionRequest returns a request value for making API operation for
// Amazon Fraud Detector.
//
// Deletes the rule version. You cannot delete a rule version if it is used
// by an ACTIVE or INACTIVE detector version.
//
//    // Example sending a request using DeleteRuleVersionRequest.
//    req := client.DeleteRuleVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/frauddetector-2019-11-15/DeleteRuleVersion
func (c *Client) DeleteRuleVersionRequest(input *DeleteRuleVersionInput) DeleteRuleVersionRequest {
	op := &aws.Operation{
		Name:       opDeleteRuleVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRuleVersionInput{}
	}

	req := c.newRequest(op, input, &DeleteRuleVersionOutput{})

	return DeleteRuleVersionRequest{Request: req, Input: input, Copy: c.DeleteRuleVersionRequest}
}

// DeleteRuleVersionRequest is the request type for the
// DeleteRuleVersion API operation.
type DeleteRuleVersionRequest struct {
	*aws.Request
	Input *DeleteRuleVersionInput
	Copy  func(*DeleteRuleVersionInput) DeleteRuleVersionRequest
}

// Send marshals and sends the DeleteRuleVersion API request.
func (r DeleteRuleVersionRequest) Send(ctx context.Context) (*DeleteRuleVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRuleVersionResponse{
		DeleteRuleVersionOutput: r.Request.Data.(*DeleteRuleVersionOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRuleVersionResponse is the response type for the
// DeleteRuleVersion API operation.
type DeleteRuleVersionResponse struct {
	*DeleteRuleVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRuleVersion request.
func (r *DeleteRuleVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
