// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ResetFpgaImageAttributeInput struct {
	_ struct{} `type:"structure"`

	// The attribute.
	Attribute ResetFpgaImageAttributeName `type:"string" enum:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the AFI.
	//
	// FpgaImageId is a required field
	FpgaImageId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ResetFpgaImageAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResetFpgaImageAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResetFpgaImageAttributeInput"}

	if s.FpgaImageId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FpgaImageId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ResetFpgaImageAttributeOutput struct {
	_ struct{} `type:"structure"`

	// Is true if the request succeeds, and an error otherwise.
	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s ResetFpgaImageAttributeOutput) String() string {
	return awsutil.Prettify(s)
}

const opResetFpgaImageAttribute = "ResetFpgaImageAttribute"

// ResetFpgaImageAttributeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Resets the specified attribute of the specified Amazon FPGA Image (AFI) to
// its default value. You can only reset the load permission attribute.
//
//    // Example sending a request using ResetFpgaImageAttributeRequest.
//    req := client.ResetFpgaImageAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ResetFpgaImageAttribute
func (c *Client) ResetFpgaImageAttributeRequest(input *ResetFpgaImageAttributeInput) ResetFpgaImageAttributeRequest {
	op := &aws.Operation{
		Name:       opResetFpgaImageAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResetFpgaImageAttributeInput{}
	}

	req := c.newRequest(op, input, &ResetFpgaImageAttributeOutput{})

	return ResetFpgaImageAttributeRequest{Request: req, Input: input, Copy: c.ResetFpgaImageAttributeRequest}
}

// ResetFpgaImageAttributeRequest is the request type for the
// ResetFpgaImageAttribute API operation.
type ResetFpgaImageAttributeRequest struct {
	*aws.Request
	Input *ResetFpgaImageAttributeInput
	Copy  func(*ResetFpgaImageAttributeInput) ResetFpgaImageAttributeRequest
}

// Send marshals and sends the ResetFpgaImageAttribute API request.
func (r ResetFpgaImageAttributeRequest) Send(ctx context.Context) (*ResetFpgaImageAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResetFpgaImageAttributeResponse{
		ResetFpgaImageAttributeOutput: r.Request.Data.(*ResetFpgaImageAttributeOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResetFpgaImageAttributeResponse is the response type for the
// ResetFpgaImageAttribute API operation.
type ResetFpgaImageAttributeResponse struct {
	*ResetFpgaImageAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResetFpgaImageAttribute request.
func (r *ResetFpgaImageAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
