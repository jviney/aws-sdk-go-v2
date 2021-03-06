// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// A request to create a SqlInjectionMatchSet.
type CreateSqlInjectionMatchSetInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// A friendly name or description for the SqlInjectionMatchSet that you're creating.
	// You can't change Name after you create the SqlInjectionMatchSet.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateSqlInjectionMatchSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSqlInjectionMatchSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSqlInjectionMatchSetInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The response to a CreateSqlInjectionMatchSet request.
type CreateSqlInjectionMatchSetOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the CreateSqlInjectionMatchSet request.
	// You can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`

	// A SqlInjectionMatchSet.
	SqlInjectionMatchSet *SqlInjectionMatchSet `type:"structure"`
}

// String returns the string representation
func (s CreateSqlInjectionMatchSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateSqlInjectionMatchSet = "CreateSqlInjectionMatchSet"

// CreateSqlInjectionMatchSetRequest returns a request value for making API operation for
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
// Creates a SqlInjectionMatchSet, which you use to allow, block, or count requests
// that contain snippets of SQL code in a specified part of web requests. AWS
// WAF searches for character sequences that are likely to be malicious strings.
//
// To create and configure a SqlInjectionMatchSet, perform the following steps:
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a CreateSqlInjectionMatchSet request.
//
// Submit a CreateSqlInjectionMatchSet request.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateSqlInjectionMatchSet request.
//
// Submit an UpdateSqlInjectionMatchSet request to specify the parts of web
// requests in which you want to allow, block, or count malicious SQL code.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using CreateSqlInjectionMatchSetRequest.
//    req := client.CreateSqlInjectionMatchSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/CreateSqlInjectionMatchSet
func (c *Client) CreateSqlInjectionMatchSetRequest(input *CreateSqlInjectionMatchSetInput) CreateSqlInjectionMatchSetRequest {
	op := &aws.Operation{
		Name:       opCreateSqlInjectionMatchSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSqlInjectionMatchSetInput{}
	}

	req := c.newRequest(op, input, &CreateSqlInjectionMatchSetOutput{})

	return CreateSqlInjectionMatchSetRequest{Request: req, Input: input, Copy: c.CreateSqlInjectionMatchSetRequest}
}

// CreateSqlInjectionMatchSetRequest is the request type for the
// CreateSqlInjectionMatchSet API operation.
type CreateSqlInjectionMatchSetRequest struct {
	*aws.Request
	Input *CreateSqlInjectionMatchSetInput
	Copy  func(*CreateSqlInjectionMatchSetInput) CreateSqlInjectionMatchSetRequest
}

// Send marshals and sends the CreateSqlInjectionMatchSet API request.
func (r CreateSqlInjectionMatchSetRequest) Send(ctx context.Context) (*CreateSqlInjectionMatchSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSqlInjectionMatchSetResponse{
		CreateSqlInjectionMatchSetOutput: r.Request.Data.(*CreateSqlInjectionMatchSetOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSqlInjectionMatchSetResponse is the response type for the
// CreateSqlInjectionMatchSet API operation.
type CreateSqlInjectionMatchSetResponse struct {
	*CreateSqlInjectionMatchSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSqlInjectionMatchSet request.
func (r *CreateSqlInjectionMatchSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
