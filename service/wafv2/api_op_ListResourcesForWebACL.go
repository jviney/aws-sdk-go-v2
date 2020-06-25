// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafv2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListResourcesForWebACLInput struct {
	_ struct{} `type:"structure"`

	// Used for web ACLs that are scoped for regional applications. A regional application
	// can be an Application Load Balancer (ALB) or an API Gateway stage.
	ResourceType ResourceType `type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) of the Web ACL.
	//
	// WebACLArn is a required field
	WebACLArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s ListResourcesForWebACLInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListResourcesForWebACLInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListResourcesForWebACLInput"}

	if s.WebACLArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("WebACLArn"))
	}
	if s.WebACLArn != nil && len(*s.WebACLArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("WebACLArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListResourcesForWebACLOutput struct {
	_ struct{} `type:"structure"`

	// The array of Amazon Resource Names (ARNs) of the associated resources.
	ResourceArns []string `type:"list"`
}

// String returns the string representation
func (s ListResourcesForWebACLOutput) String() string {
	return awsutil.Prettify(s)
}

const opListResourcesForWebACL = "ListResourcesForWebACL"

// ListResourcesForWebACLRequest returns a request value for making API operation for
// AWS WAFV2.
//
//
// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from
// the prior release, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
//
// Retrieves an array of the Amazon Resource Names (ARNs) for the regional resources
// that are associated with the specified web ACL. If you want the list of AWS
// CloudFront resources, use the AWS CloudFront call ListDistributionsByWebACLId.
//
//    // Example sending a request using ListResourcesForWebACLRequest.
//    req := client.ListResourcesForWebACLRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/wafv2-2019-07-29/ListResourcesForWebACL
func (c *Client) ListResourcesForWebACLRequest(input *ListResourcesForWebACLInput) ListResourcesForWebACLRequest {
	op := &aws.Operation{
		Name:       opListResourcesForWebACL,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListResourcesForWebACLInput{}
	}

	req := c.newRequest(op, input, &ListResourcesForWebACLOutput{})

	return ListResourcesForWebACLRequest{Request: req, Input: input, Copy: c.ListResourcesForWebACLRequest}
}

// ListResourcesForWebACLRequest is the request type for the
// ListResourcesForWebACL API operation.
type ListResourcesForWebACLRequest struct {
	*aws.Request
	Input *ListResourcesForWebACLInput
	Copy  func(*ListResourcesForWebACLInput) ListResourcesForWebACLRequest
}

// Send marshals and sends the ListResourcesForWebACL API request.
func (r ListResourcesForWebACLRequest) Send(ctx context.Context) (*ListResourcesForWebACLResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListResourcesForWebACLResponse{
		ListResourcesForWebACLOutput: r.Request.Data.(*ListResourcesForWebACLOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListResourcesForWebACLResponse is the response type for the
// ListResourcesForWebACL API operation.
type ListResourcesForWebACLResponse struct {
	*ListResourcesForWebACLOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListResourcesForWebACL request.
func (r *ListResourcesForWebACLResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
