// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type UpdateRegexPatternSetInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// The RegexPatternSetId of the RegexPatternSet that you want to update. RegexPatternSetId
	// is returned by CreateRegexPatternSet and by ListRegexPatternSets.
	//
	// RegexPatternSetId is a required field
	RegexPatternSetId *string `min:"1" type:"string" required:"true"`

	// An array of RegexPatternSetUpdate objects that you want to insert into or
	// delete from a RegexPatternSet.
	//
	// Updates is a required field
	Updates []RegexPatternSetUpdate `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateRegexPatternSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateRegexPatternSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateRegexPatternSetInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.RegexPatternSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RegexPatternSetId"))
	}
	if s.RegexPatternSetId != nil && len(*s.RegexPatternSetId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RegexPatternSetId", 1))
	}

	if s.Updates == nil {
		invalidParams.Add(aws.NewErrParamRequired("Updates"))
	}
	if s.Updates != nil && len(s.Updates) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Updates", 1))
	}
	if s.Updates != nil {
		for i, v := range s.Updates {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Updates", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateRegexPatternSetOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the UpdateRegexPatternSet request.
	// You can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateRegexPatternSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateRegexPatternSet = "UpdateRegexPatternSet"

// UpdateRegexPatternSetRequest returns a request value for making API operation for
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
// Inserts or deletes RegexPatternString objects in a RegexPatternSet. For each
// RegexPatternString object, you specify the following values:
//
//    * Whether to insert or delete the RegexPatternString.
//
//    * The regular expression pattern that you want to insert or delete. For
//    more information, see RegexPatternSet.
//
// For example, you can create a RegexPatternString such as B[a@]dB[o0]t. AWS
// WAF will match this RegexPatternString to:
//
//    * BadBot
//
//    * BadB0t
//
//    * B@dBot
//
//    * B@dB0t
//
// To create and configure a RegexPatternSet, perform the following steps:
//
// Create a RegexPatternSet. For more information, see CreateRegexPatternSet.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateRegexPatternSet request.
//
// Submit an UpdateRegexPatternSet request to specify the regular expression
// pattern that you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using UpdateRegexPatternSetRequest.
//    req := client.UpdateRegexPatternSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/UpdateRegexPatternSet
func (c *Client) UpdateRegexPatternSetRequest(input *UpdateRegexPatternSetInput) UpdateRegexPatternSetRequest {
	op := &aws.Operation{
		Name:       opUpdateRegexPatternSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateRegexPatternSetInput{}
	}

	req := c.newRequest(op, input, &UpdateRegexPatternSetOutput{})

	return UpdateRegexPatternSetRequest{Request: req, Input: input, Copy: c.UpdateRegexPatternSetRequest}
}

// UpdateRegexPatternSetRequest is the request type for the
// UpdateRegexPatternSet API operation.
type UpdateRegexPatternSetRequest struct {
	*aws.Request
	Input *UpdateRegexPatternSetInput
	Copy  func(*UpdateRegexPatternSetInput) UpdateRegexPatternSetRequest
}

// Send marshals and sends the UpdateRegexPatternSet API request.
func (r UpdateRegexPatternSetRequest) Send(ctx context.Context) (*UpdateRegexPatternSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateRegexPatternSetResponse{
		UpdateRegexPatternSetOutput: r.Request.Data.(*UpdateRegexPatternSetOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateRegexPatternSetResponse is the response type for the
// UpdateRegexPatternSet API operation.
type UpdateRegexPatternSetResponse struct {
	*UpdateRegexPatternSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateRegexPatternSet request.
func (r *UpdateRegexPatternSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
