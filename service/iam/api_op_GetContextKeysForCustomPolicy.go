// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetContextKeysForCustomPolicyInput struct {
	_ struct{} `type:"structure"`

	// A list of policies for which you want the list of context keys referenced
	// in those policies. Each document is specified as a string containing the
	// complete, valid JSON text of an IAM policy.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) used to validate this
	// parameter is a string of characters consisting of the following:
	//
	//    * Any printable ASCII character ranging from the space character (\u0020)
	//    through the end of the ASCII character range
	//
	//    * The printable characters in the Basic Latin and Latin-1 Supplement character
	//    set (through \u00FF)
	//
	//    * The special characters tab (\u0009), line feed (\u000A), and carriage
	//    return (\u000D)
	//
	// PolicyInputList is a required field
	PolicyInputList []string `type:"list" required:"true"`
}

// String returns the string representation
func (s GetContextKeysForCustomPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetContextKeysForCustomPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetContextKeysForCustomPolicyInput"}

	if s.PolicyInputList == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyInputList"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful GetContextKeysForPrincipalPolicy or
// GetContextKeysForCustomPolicy request.
type GetContextKeysForCustomPolicyOutput struct {
	_ struct{} `type:"structure"`

	// The list of context keys that are referenced in the input policies.
	ContextKeyNames []string `type:"list"`
}

// String returns the string representation
func (s GetContextKeysForCustomPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetContextKeysForCustomPolicy = "GetContextKeysForCustomPolicy"

// GetContextKeysForCustomPolicyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Gets a list of all of the context keys referenced in the input policies.
// The policies are supplied as a list of one or more strings. To get the context
// keys from policies associated with an IAM user, group, or role, use GetContextKeysForPrincipalPolicy.
//
// Context keys are variables maintained by AWS and its services that provide
// details about the context of an API query request. Context keys can be evaluated
// by testing against a value specified in an IAM policy. Use GetContextKeysForCustomPolicy
// to understand what key names and values you must supply when you call SimulateCustomPolicy.
// Note that all parameters are shown in unencoded form here for clarity but
// must be URL encoded to be included as a part of a real HTML request.
//
//    // Example sending a request using GetContextKeysForCustomPolicyRequest.
//    req := client.GetContextKeysForCustomPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetContextKeysForCustomPolicy
func (c *Client) GetContextKeysForCustomPolicyRequest(input *GetContextKeysForCustomPolicyInput) GetContextKeysForCustomPolicyRequest {
	op := &aws.Operation{
		Name:       opGetContextKeysForCustomPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetContextKeysForCustomPolicyInput{}
	}

	req := c.newRequest(op, input, &GetContextKeysForCustomPolicyOutput{})

	return GetContextKeysForCustomPolicyRequest{Request: req, Input: input, Copy: c.GetContextKeysForCustomPolicyRequest}
}

// GetContextKeysForCustomPolicyRequest is the request type for the
// GetContextKeysForCustomPolicy API operation.
type GetContextKeysForCustomPolicyRequest struct {
	*aws.Request
	Input *GetContextKeysForCustomPolicyInput
	Copy  func(*GetContextKeysForCustomPolicyInput) GetContextKeysForCustomPolicyRequest
}

// Send marshals and sends the GetContextKeysForCustomPolicy API request.
func (r GetContextKeysForCustomPolicyRequest) Send(ctx context.Context) (*GetContextKeysForCustomPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetContextKeysForCustomPolicyResponse{
		GetContextKeysForCustomPolicyOutput: r.Request.Data.(*GetContextKeysForCustomPolicyOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetContextKeysForCustomPolicyResponse is the response type for the
// GetContextKeysForCustomPolicy API operation.
type GetContextKeysForCustomPolicyResponse struct {
	*GetContextKeysForCustomPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetContextKeysForCustomPolicy request.
func (r *GetContextKeysForCustomPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
