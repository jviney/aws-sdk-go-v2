// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sfn

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type UntagResourceInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) for the Step Functions state machine or activity.
	//
	// ResourceArn is a required field
	ResourceArn *string `locationName:"resourceArn" min:"1" type:"string" required:"true"`

	// The list of tags to remove from the resource.
	//
	// TagKeys is a required field
	TagKeys []string `locationName:"tagKeys" type:"list" required:"true"`
}

// String returns the string representation
func (s UntagResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UntagResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UntagResourceInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}
	if s.ResourceArn != nil && len(*s.ResourceArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 1))
	}

	if s.TagKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagKeys"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UntagResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UntagResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opUntagResource = "UntagResource"

// UntagResourceRequest returns a request value for making API operation for
// AWS Step Functions.
//
// Remove a tag from a Step Functions resource
//
//    // Example sending a request using UntagResourceRequest.
//    req := client.UntagResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23/UntagResource
func (c *Client) UntagResourceRequest(input *UntagResourceInput) UntagResourceRequest {
	op := &aws.Operation{
		Name:       opUntagResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UntagResourceInput{}
	}

	req := c.newRequest(op, input, &UntagResourceOutput{})

	return UntagResourceRequest{Request: req, Input: input, Copy: c.UntagResourceRequest}
}

// UntagResourceRequest is the request type for the
// UntagResource API operation.
type UntagResourceRequest struct {
	*aws.Request
	Input *UntagResourceInput
	Copy  func(*UntagResourceInput) UntagResourceRequest
}

// Send marshals and sends the UntagResource API request.
func (r UntagResourceRequest) Send(ctx context.Context) (*UntagResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UntagResourceResponse{
		UntagResourceOutput: r.Request.Data.(*UntagResourceOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UntagResourceResponse is the response type for the
// UntagResource API operation.
type UntagResourceResponse struct {
	*UntagResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UntagResource request.
func (r *UntagResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
