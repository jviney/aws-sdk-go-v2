// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateAcceleratorInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether an accelerator is enabled. The value is true or false.
	// The default value is true.
	//
	// If the value is set to true, an accelerator cannot be deleted. If set to
	// false, the accelerator can be deleted.
	Enabled *bool `type:"boolean"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency—that
	// is, the uniqueness—of an accelerator.
	//
	// IdempotencyToken is a required field
	IdempotencyToken *string `type:"string" required:"true" idempotencyToken:"true"`

	// The value for the address type must be IPv4.
	IpAddressType IpAddressType `type:"string" enum:"true"`

	// Optionally, if you've added your own IP address pool to Global Accelerator,
	// you can choose IP addresses from your own pool to use for the accelerator's
	// static IP addresses. You can specify one or two addresses, separated by a
	// comma. Do not include the /32 suffix.
	//
	// If you specify only one IP address from your IP address range, Global Accelerator
	// assigns a second static IP address for the accelerator from the AWS IP address
	// pool.
	//
	// For more information, see Bring Your Own IP Addresses (BYOIP) (https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html)
	// in the AWS Global Accelerator Developer Guide.
	IpAddresses []string `type:"list"`

	// The name of an accelerator. The name can have a maximum of 32 characters,
	// must contain only alphanumeric characters or hyphens (-), and must not begin
	// or end with a hyphen.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// Create tags for an accelerator.
	//
	// For more information, see Tagging in AWS Global Accelerator (https://docs.aws.amazon.com/global-accelerator/latest/dg/tagging-in-global-accelerator.html)
	// in the AWS Global Accelerator Developer Guide.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateAcceleratorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAcceleratorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAcceleratorInput"}

	if s.IdempotencyToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdempotencyToken"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateAcceleratorOutput struct {
	_ struct{} `type:"structure"`

	// The accelerator that is created by specifying a listener and the supported
	// IP address types.
	Accelerator *Accelerator `type:"structure"`
}

// String returns the string representation
func (s CreateAcceleratorOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateAccelerator = "CreateAccelerator"

// CreateAcceleratorRequest returns a request value for making API operation for
// AWS Global Accelerator.
//
// Create an accelerator. An accelerator includes one or more listeners that
// process inbound connections and direct traffic to one or more endpoint groups,
// each of which includes endpoints, such as Network Load Balancers. To see
// an AWS CLI example of creating an accelerator, scroll down to Example.
//
// If you bring your own IP address ranges to AWS Global Accelerator (BYOIP),
// you can assign IP addresses from your own pool to your accelerator as the
// static IP address entry points. Only one IP address from each of your IP
// address ranges can be used for each accelerator.
//
// You must specify the US West (Oregon) Region to create or update accelerators.
//
//    // Example sending a request using CreateAcceleratorRequest.
//    req := client.CreateAcceleratorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/CreateAccelerator
func (c *Client) CreateAcceleratorRequest(input *CreateAcceleratorInput) CreateAcceleratorRequest {
	op := &aws.Operation{
		Name:       opCreateAccelerator,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateAcceleratorInput{}
	}

	req := c.newRequest(op, input, &CreateAcceleratorOutput{})

	return CreateAcceleratorRequest{Request: req, Input: input, Copy: c.CreateAcceleratorRequest}
}

// CreateAcceleratorRequest is the request type for the
// CreateAccelerator API operation.
type CreateAcceleratorRequest struct {
	*aws.Request
	Input *CreateAcceleratorInput
	Copy  func(*CreateAcceleratorInput) CreateAcceleratorRequest
}

// Send marshals and sends the CreateAccelerator API request.
func (r CreateAcceleratorRequest) Send(ctx context.Context) (*CreateAcceleratorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAcceleratorResponse{
		CreateAcceleratorOutput: r.Request.Data.(*CreateAcceleratorOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAcceleratorResponse is the response type for the
// CreateAccelerator API operation.
type CreateAcceleratorResponse struct {
	*CreateAcceleratorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateAccelerator request.
func (r *CreateAcceleratorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
