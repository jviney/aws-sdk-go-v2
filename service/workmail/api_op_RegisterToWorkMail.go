// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type RegisterToWorkMailInput struct {
	_ struct{} `type:"structure"`

	// The email for the user, group, or resource to be updated.
	//
	// Email is a required field
	Email *string `min:"1" type:"string" required:"true"`

	// The identifier for the user, group, or resource to be updated.
	//
	// EntityId is a required field
	EntityId *string `min:"12" type:"string" required:"true"`

	// The identifier for the organization under which the user, group, or resource
	// exists.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RegisterToWorkMailInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterToWorkMailInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterToWorkMailInput"}

	if s.Email == nil {
		invalidParams.Add(aws.NewErrParamRequired("Email"))
	}
	if s.Email != nil && len(*s.Email) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Email", 1))
	}

	if s.EntityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntityId"))
	}
	if s.EntityId != nil && len(*s.EntityId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityId", 12))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RegisterToWorkMailOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RegisterToWorkMailOutput) String() string {
	return awsutil.Prettify(s)
}

const opRegisterToWorkMail = "RegisterToWorkMail"

// RegisterToWorkMailRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Registers an existing and disabled user, group, or resource for Amazon WorkMail
// use by associating a mailbox and calendaring capabilities. It performs no
// change if the user, group, or resource is enabled and fails if the user,
// group, or resource is deleted. This operation results in the accumulation
// of costs. For more information, see Pricing (https://aws.amazon.com/workmail/pricing).
// The equivalent console functionality for this operation is Enable.
//
// Users can either be created by calling the CreateUser API operation or they
// can be synchronized from your directory. For more information, see DeregisterFromWorkMail.
//
//    // Example sending a request using RegisterToWorkMailRequest.
//    req := client.RegisterToWorkMailRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/RegisterToWorkMail
func (c *Client) RegisterToWorkMailRequest(input *RegisterToWorkMailInput) RegisterToWorkMailRequest {
	op := &aws.Operation{
		Name:       opRegisterToWorkMail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterToWorkMailInput{}
	}

	req := c.newRequest(op, input, &RegisterToWorkMailOutput{})

	return RegisterToWorkMailRequest{Request: req, Input: input, Copy: c.RegisterToWorkMailRequest}
}

// RegisterToWorkMailRequest is the request type for the
// RegisterToWorkMail API operation.
type RegisterToWorkMailRequest struct {
	*aws.Request
	Input *RegisterToWorkMailInput
	Copy  func(*RegisterToWorkMailInput) RegisterToWorkMailRequest
}

// Send marshals and sends the RegisterToWorkMail API request.
func (r RegisterToWorkMailRequest) Send(ctx context.Context) (*RegisterToWorkMailResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterToWorkMailResponse{
		RegisterToWorkMailOutput: r.Request.Data.(*RegisterToWorkMailOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterToWorkMailResponse is the response type for the
// RegisterToWorkMail API operation.
type RegisterToWorkMailResponse struct {
	*RegisterToWorkMailOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterToWorkMail request.
func (r *RegisterToWorkMailResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
