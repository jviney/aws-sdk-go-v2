// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteRegexMatchSetInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// The RegexMatchSetId of the RegexMatchSet that you want to delete. RegexMatchSetId
	// is returned by CreateRegexMatchSet and by ListRegexMatchSets.
	//
	// RegexMatchSetId is a required field
	RegexMatchSetId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRegexMatchSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRegexMatchSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRegexMatchSetInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.RegexMatchSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RegexMatchSetId"))
	}
	if s.RegexMatchSetId != nil && len(*s.RegexMatchSetId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RegexMatchSetId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteRegexMatchSetOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the DeleteRegexMatchSet request.
	// You can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteRegexMatchSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteRegexMatchSet = "DeleteRegexMatchSet"

// DeleteRegexMatchSetRequest returns a request value for making API operation for
// AWS WAF.
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
// Permanently deletes a RegexMatchSet. You can't delete a RegexMatchSet if
// it's still used in any Rules or if it still includes any RegexMatchTuples
// objects (any filters).
//
// If you just want to remove a RegexMatchSet from a Rule, use UpdateRule.
//
// To permanently delete a RegexMatchSet, perform the following steps:
//
// Update the RegexMatchSet to remove filters, if any. For more information,
// see UpdateRegexMatchSet.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a DeleteRegexMatchSet request.
//
// Submit a DeleteRegexMatchSet request.
//
//    // Example sending a request using DeleteRegexMatchSetRequest.
//    req := client.DeleteRegexMatchSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/DeleteRegexMatchSet
func (c *Client) DeleteRegexMatchSetRequest(input *DeleteRegexMatchSetInput) DeleteRegexMatchSetRequest {
	op := &aws.Operation{
		Name:       opDeleteRegexMatchSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRegexMatchSetInput{}
	}

	req := c.newRequest(op, input, &DeleteRegexMatchSetOutput{})

	return DeleteRegexMatchSetRequest{Request: req, Input: input, Copy: c.DeleteRegexMatchSetRequest}
}

// DeleteRegexMatchSetRequest is the request type for the
// DeleteRegexMatchSet API operation.
type DeleteRegexMatchSetRequest struct {
	*aws.Request
	Input *DeleteRegexMatchSetInput
	Copy  func(*DeleteRegexMatchSetInput) DeleteRegexMatchSetRequest
}

// Send marshals and sends the DeleteRegexMatchSet API request.
func (r DeleteRegexMatchSetRequest) Send(ctx context.Context) (*DeleteRegexMatchSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRegexMatchSetResponse{
		DeleteRegexMatchSetOutput: r.Request.Data.(*DeleteRegexMatchSetOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRegexMatchSetResponse is the response type for the
// DeleteRegexMatchSet API operation.
type DeleteRegexMatchSetResponse struct {
	*DeleteRegexMatchSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRegexMatchSet request.
func (r *DeleteRegexMatchSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
