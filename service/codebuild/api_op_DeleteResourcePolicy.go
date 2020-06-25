// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteResourcePolicyInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the resource that is associated with the resource policy.
	//
	// ResourceArn is a required field
	ResourceArn *string `locationName:"resourceArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteResourcePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteResourcePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteResourcePolicyInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}
	if s.ResourceArn != nil && len(*s.ResourceArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteResourcePolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteResourcePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteResourcePolicy = "DeleteResourcePolicy"

// DeleteResourcePolicyRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// Deletes a resource policy that is identified by its resource ARN.
//
//    // Example sending a request using DeleteResourcePolicyRequest.
//    req := client.DeleteResourcePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/DeleteResourcePolicy
func (c *Client) DeleteResourcePolicyRequest(input *DeleteResourcePolicyInput) DeleteResourcePolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteResourcePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteResourcePolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteResourcePolicyOutput{})

	return DeleteResourcePolicyRequest{Request: req, Input: input, Copy: c.DeleteResourcePolicyRequest}
}

// DeleteResourcePolicyRequest is the request type for the
// DeleteResourcePolicy API operation.
type DeleteResourcePolicyRequest struct {
	*aws.Request
	Input *DeleteResourcePolicyInput
	Copy  func(*DeleteResourcePolicyInput) DeleteResourcePolicyRequest
}

// Send marshals and sends the DeleteResourcePolicy API request.
func (r DeleteResourcePolicyRequest) Send(ctx context.Context) (*DeleteResourcePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteResourcePolicyResponse{
		DeleteResourcePolicyOutput: r.Request.Data.(*DeleteResourcePolicyOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteResourcePolicyResponse is the response type for the
// DeleteResourcePolicy API operation.
type DeleteResourcePolicyResponse struct {
	*DeleteResourcePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteResourcePolicy request.
func (r *DeleteResourcePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
