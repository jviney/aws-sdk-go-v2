// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafv2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteWebACLInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the Web ACL. This ID is returned in the responses
	// to create and list commands. You provide it to operations like update and
	// delete.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// A token used for optimistic locking. AWS WAF returns a token to your get
	// and list requests, to mark the state of the entity at the time of the request.
	// To make changes to the entity associated with the token, you provide the
	// token to operations like update and delete. AWS WAF uses the token to ensure
	// that no changes have been made to the entity since you last retrieved it.
	// If a change has been made, the update fails with a WAFOptimisticLockException.
	// If this happens, perform another get, and use the new token returned by that
	// operation.
	//
	// LockToken is a required field
	LockToken *string `min:"1" type:"string" required:"true"`

	// The name of the Web ACL. You cannot change the name of a Web ACL after you
	// create it.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// Specifies whether this is for an AWS CloudFront distribution or for a regional
	// application. A regional application can be an Application Load Balancer (ALB)
	// or an API Gateway stage.
	//
	// To work with CloudFront, you must also specify the Region US East (N. Virginia)
	// as follows:
	//
	//    * CLI - Specify the Region when you use the CloudFront scope: --scope=CLOUDFRONT
	//    --region=us-east-1.
	//
	//    * API and SDKs - For all calls, use the Region endpoint us-east-1.
	//
	// Scope is a required field
	Scope Scope `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s DeleteWebACLInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteWebACLInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteWebACLInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if s.LockToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("LockToken"))
	}
	if s.LockToken != nil && len(*s.LockToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LockToken", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if len(s.Scope) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Scope"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteWebACLOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteWebACLOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteWebACL = "DeleteWebACL"

// DeleteWebACLRequest returns a request value for making API operation for
// AWS WAFV2.
//
//
// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from
// the prior release, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
//
// Deletes the specified WebACL.
//
// You can only use this if ManagedByFirewallManager is false in the specified
// WebACL.
//
//    // Example sending a request using DeleteWebACLRequest.
//    req := client.DeleteWebACLRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/wafv2-2019-07-29/DeleteWebACL
func (c *Client) DeleteWebACLRequest(input *DeleteWebACLInput) DeleteWebACLRequest {
	op := &aws.Operation{
		Name:       opDeleteWebACL,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteWebACLInput{}
	}

	req := c.newRequest(op, input, &DeleteWebACLOutput{})

	return DeleteWebACLRequest{Request: req, Input: input, Copy: c.DeleteWebACLRequest}
}

// DeleteWebACLRequest is the request type for the
// DeleteWebACL API operation.
type DeleteWebACLRequest struct {
	*aws.Request
	Input *DeleteWebACLInput
	Copy  func(*DeleteWebACLInput) DeleteWebACLRequest
}

// Send marshals and sends the DeleteWebACL API request.
func (r DeleteWebACLRequest) Send(ctx context.Context) (*DeleteWebACLResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteWebACLResponse{
		DeleteWebACLOutput: r.Request.Data.(*DeleteWebACLOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteWebACLResponse is the response type for the
// DeleteWebACL API operation.
type DeleteWebACLResponse struct {
	*DeleteWebACLOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteWebACL request.
func (r *DeleteWebACLResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
