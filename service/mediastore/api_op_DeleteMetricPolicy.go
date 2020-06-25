// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastore

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteMetricPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name of the container that is associated with the metric policy that
	// you want to delete.
	//
	// ContainerName is a required field
	ContainerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMetricPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMetricPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteMetricPolicyInput"}

	if s.ContainerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContainerName"))
	}
	if s.ContainerName != nil && len(*s.ContainerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContainerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteMetricPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteMetricPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteMetricPolicy = "DeleteMetricPolicy"

// DeleteMetricPolicyRequest returns a request value for making API operation for
// AWS Elemental MediaStore.
//
// Deletes the metric policy that is associated with the specified container.
// If there is no metric policy associated with the container, MediaStore doesn't
// send metrics to CloudWatch.
//
//    // Example sending a request using DeleteMetricPolicyRequest.
//    req := client.DeleteMetricPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/DeleteMetricPolicy
func (c *Client) DeleteMetricPolicyRequest(input *DeleteMetricPolicyInput) DeleteMetricPolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteMetricPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteMetricPolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteMetricPolicyOutput{})

	return DeleteMetricPolicyRequest{Request: req, Input: input, Copy: c.DeleteMetricPolicyRequest}
}

// DeleteMetricPolicyRequest is the request type for the
// DeleteMetricPolicy API operation.
type DeleteMetricPolicyRequest struct {
	*aws.Request
	Input *DeleteMetricPolicyInput
	Copy  func(*DeleteMetricPolicyInput) DeleteMetricPolicyRequest
}

// Send marshals and sends the DeleteMetricPolicy API request.
func (r DeleteMetricPolicyRequest) Send(ctx context.Context) (*DeleteMetricPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteMetricPolicyResponse{
		DeleteMetricPolicyOutput: r.Request.Data.(*DeleteMetricPolicyOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteMetricPolicyResponse is the response type for the
// DeleteMetricPolicy API operation.
type DeleteMetricPolicyResponse struct {
	*DeleteMetricPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteMetricPolicy request.
func (r *DeleteMetricPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
