// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafv2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListWebACLsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of objects that you want AWS WAF to return for this request.
	// If more objects are available, in the response, AWS WAF provides a NextMarker
	// value that you can use in a subsequent call to get the next batch of objects.
	Limit *int64 `min:"1" type:"integer"`

	// When you request a list of objects with a Limit setting, if the number of
	// objects that are still available for retrieval exceeds the limit, AWS WAF
	// returns a NextMarker value in the response. To retrieve the next batch of
	// objects, provide the marker from the prior call in your next request.
	NextMarker *string `min:"1" type:"string"`

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
func (s ListWebACLsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListWebACLsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListWebACLsInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.NextMarker != nil && len(*s.NextMarker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextMarker", 1))
	}
	if len(s.Scope) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Scope"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListWebACLsOutput struct {
	_ struct{} `type:"structure"`

	// When you request a list of objects with a Limit setting, if the number of
	// objects that are still available for retrieval exceeds the limit, AWS WAF
	// returns a NextMarker value in the response. To retrieve the next batch of
	// objects, provide the marker from the prior call in your next request.
	NextMarker *string `min:"1" type:"string"`

	WebACLs []WebACLSummary `type:"list"`
}

// String returns the string representation
func (s ListWebACLsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListWebACLs = "ListWebACLs"

// ListWebACLsRequest returns a request value for making API operation for
// AWS WAFV2.
//
//
// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from
// the prior release, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
//
// Retrieves an array of WebACLSummary objects for the web ACLs that you manage.
//
//    // Example sending a request using ListWebACLsRequest.
//    req := client.ListWebACLsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/wafv2-2019-07-29/ListWebACLs
func (c *Client) ListWebACLsRequest(input *ListWebACLsInput) ListWebACLsRequest {
	op := &aws.Operation{
		Name:       opListWebACLs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListWebACLsInput{}
	}

	req := c.newRequest(op, input, &ListWebACLsOutput{})

	return ListWebACLsRequest{Request: req, Input: input, Copy: c.ListWebACLsRequest}
}

// ListWebACLsRequest is the request type for the
// ListWebACLs API operation.
type ListWebACLsRequest struct {
	*aws.Request
	Input *ListWebACLsInput
	Copy  func(*ListWebACLsInput) ListWebACLsRequest
}

// Send marshals and sends the ListWebACLs API request.
func (r ListWebACLsRequest) Send(ctx context.Context) (*ListWebACLsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListWebACLsResponse{
		ListWebACLsOutput: r.Request.Data.(*ListWebACLsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListWebACLsResponse is the response type for the
// ListWebACLs API operation.
type ListWebACLsResponse struct {
	*ListWebACLsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListWebACLs request.
func (r *ListWebACLsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
