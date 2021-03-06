// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetRegexMatchSetInput struct {
	_ struct{} `type:"structure"`

	// The RegexMatchSetId of the RegexMatchSet that you want to get. RegexMatchSetId
	// is returned by CreateRegexMatchSet and by ListRegexMatchSets.
	//
	// RegexMatchSetId is a required field
	RegexMatchSetId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRegexMatchSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRegexMatchSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRegexMatchSetInput"}

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

type GetRegexMatchSetOutput struct {
	_ struct{} `type:"structure"`

	// Information about the RegexMatchSet that you specified in the GetRegexMatchSet
	// request. For more information, see RegexMatchTuple.
	RegexMatchSet *RegexMatchSet `type:"structure"`
}

// String returns the string representation
func (s GetRegexMatchSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRegexMatchSet = "GetRegexMatchSet"

// GetRegexMatchSetRequest returns a request value for making API operation for
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
// Returns the RegexMatchSet specified by RegexMatchSetId.
//
//    // Example sending a request using GetRegexMatchSetRequest.
//    req := client.GetRegexMatchSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/GetRegexMatchSet
func (c *Client) GetRegexMatchSetRequest(input *GetRegexMatchSetInput) GetRegexMatchSetRequest {
	op := &aws.Operation{
		Name:       opGetRegexMatchSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRegexMatchSetInput{}
	}

	req := c.newRequest(op, input, &GetRegexMatchSetOutput{})

	return GetRegexMatchSetRequest{Request: req, Input: input, Copy: c.GetRegexMatchSetRequest}
}

// GetRegexMatchSetRequest is the request type for the
// GetRegexMatchSet API operation.
type GetRegexMatchSetRequest struct {
	*aws.Request
	Input *GetRegexMatchSetInput
	Copy  func(*GetRegexMatchSetInput) GetRegexMatchSetRequest
}

// Send marshals and sends the GetRegexMatchSet API request.
func (r GetRegexMatchSetRequest) Send(ctx context.Context) (*GetRegexMatchSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRegexMatchSetResponse{
		GetRegexMatchSetOutput: r.Request.Data.(*GetRegexMatchSetOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRegexMatchSetResponse is the response type for the
// GetRegexMatchSet API operation.
type GetRegexMatchSetResponse struct {
	*GetRegexMatchSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRegexMatchSet request.
func (r *GetRegexMatchSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
