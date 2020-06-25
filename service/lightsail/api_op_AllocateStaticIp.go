// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type AllocateStaticIpInput struct {
	_ struct{} `type:"structure"`

	// The name of the static IP address.
	//
	// StaticIpName is a required field
	StaticIpName *string `locationName:"staticIpName" type:"string" required:"true"`
}

// String returns the string representation
func (s AllocateStaticIpInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AllocateStaticIpInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AllocateStaticIpInput"}

	if s.StaticIpName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StaticIpName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AllocateStaticIpOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected
	// by the request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s AllocateStaticIpOutput) String() string {
	return awsutil.Prettify(s)
}

const opAllocateStaticIp = "AllocateStaticIp"

// AllocateStaticIpRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Allocates a static IP address.
//
//    // Example sending a request using AllocateStaticIpRequest.
//    req := client.AllocateStaticIpRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/AllocateStaticIp
func (c *Client) AllocateStaticIpRequest(input *AllocateStaticIpInput) AllocateStaticIpRequest {
	op := &aws.Operation{
		Name:       opAllocateStaticIp,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AllocateStaticIpInput{}
	}

	req := c.newRequest(op, input, &AllocateStaticIpOutput{})

	return AllocateStaticIpRequest{Request: req, Input: input, Copy: c.AllocateStaticIpRequest}
}

// AllocateStaticIpRequest is the request type for the
// AllocateStaticIp API operation.
type AllocateStaticIpRequest struct {
	*aws.Request
	Input *AllocateStaticIpInput
	Copy  func(*AllocateStaticIpInput) AllocateStaticIpRequest
}

// Send marshals and sends the AllocateStaticIp API request.
func (r AllocateStaticIpRequest) Send(ctx context.Context) (*AllocateStaticIpResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AllocateStaticIpResponse{
		AllocateStaticIpOutput: r.Request.Data.(*AllocateStaticIpOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AllocateStaticIpResponse is the response type for the
// AllocateStaticIp API operation.
type AllocateStaticIpResponse struct {
	*AllocateStaticIpOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AllocateStaticIp request.
func (r *AllocateStaticIpResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
