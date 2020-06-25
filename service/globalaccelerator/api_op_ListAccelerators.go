// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListAcceleratorsInput struct {
	_ struct{} `type:"structure"`

	// The number of Global Accelerator objects that you want to return with this
	// call. The default value is 10.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListAcceleratorsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAcceleratorsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAcceleratorsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListAcceleratorsOutput struct {
	_ struct{} `type:"structure"`

	// The list of accelerators for a customer account.
	Accelerators []Accelerator `type:"list"`

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListAcceleratorsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListAccelerators = "ListAccelerators"

// ListAcceleratorsRequest returns a request value for making API operation for
// AWS Global Accelerator.
//
// List the accelerators for an AWS account. To see an AWS CLI example of listing
// the accelerators for an AWS account, scroll down to Example.
//
//    // Example sending a request using ListAcceleratorsRequest.
//    req := client.ListAcceleratorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/ListAccelerators
func (c *Client) ListAcceleratorsRequest(input *ListAcceleratorsInput) ListAcceleratorsRequest {
	op := &aws.Operation{
		Name:       opListAccelerators,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAcceleratorsInput{}
	}

	req := c.newRequest(op, input, &ListAcceleratorsOutput{})

	return ListAcceleratorsRequest{Request: req, Input: input, Copy: c.ListAcceleratorsRequest}
}

// ListAcceleratorsRequest is the request type for the
// ListAccelerators API operation.
type ListAcceleratorsRequest struct {
	*aws.Request
	Input *ListAcceleratorsInput
	Copy  func(*ListAcceleratorsInput) ListAcceleratorsRequest
}

// Send marshals and sends the ListAccelerators API request.
func (r ListAcceleratorsRequest) Send(ctx context.Context) (*ListAcceleratorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAcceleratorsResponse{
		ListAcceleratorsOutput: r.Request.Data.(*ListAcceleratorsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListAcceleratorsResponse is the response type for the
// ListAccelerators API operation.
type ListAcceleratorsResponse struct {
	*ListAcceleratorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAccelerators request.
func (r *ListAcceleratorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
