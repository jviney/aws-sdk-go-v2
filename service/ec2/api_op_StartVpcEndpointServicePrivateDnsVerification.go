// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type StartVpcEndpointServicePrivateDnsVerificationInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the endpoint service.
	//
	// ServiceId is a required field
	ServiceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StartVpcEndpointServicePrivateDnsVerificationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartVpcEndpointServicePrivateDnsVerificationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartVpcEndpointServicePrivateDnsVerificationInput"}

	if s.ServiceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartVpcEndpointServicePrivateDnsVerificationOutput struct {
	_ struct{} `type:"structure"`

	// Returns true if the request succeeds; otherwise, it returns an error.
	ReturnValue *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s StartVpcEndpointServicePrivateDnsVerificationOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartVpcEndpointServicePrivateDnsVerification = "StartVpcEndpointServicePrivateDnsVerification"

// StartVpcEndpointServicePrivateDnsVerificationRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Initiates the verification process to prove that the service provider owns
// the private DNS name domain for the endpoint service.
//
// The service provider must successfully perform the verification before the
// consumer can use the name to access the service.
//
// Before the service provider runs this command, they must add a record to
// the DNS server. For more information, see Adding a TXT Record to Your Domain's
// DNS Server (https://docs.aws.amazon.com/vpc/latest/userguide/ndpoint-services-dns-validation.html#add-dns-txt-record)
// in the Amazon VPC User Guide.
//
//    // Example sending a request using StartVpcEndpointServicePrivateDnsVerificationRequest.
//    req := client.StartVpcEndpointServicePrivateDnsVerificationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/StartVpcEndpointServicePrivateDnsVerification
func (c *Client) StartVpcEndpointServicePrivateDnsVerificationRequest(input *StartVpcEndpointServicePrivateDnsVerificationInput) StartVpcEndpointServicePrivateDnsVerificationRequest {
	op := &aws.Operation{
		Name:       opStartVpcEndpointServicePrivateDnsVerification,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartVpcEndpointServicePrivateDnsVerificationInput{}
	}

	req := c.newRequest(op, input, &StartVpcEndpointServicePrivateDnsVerificationOutput{})

	return StartVpcEndpointServicePrivateDnsVerificationRequest{Request: req, Input: input, Copy: c.StartVpcEndpointServicePrivateDnsVerificationRequest}
}

// StartVpcEndpointServicePrivateDnsVerificationRequest is the request type for the
// StartVpcEndpointServicePrivateDnsVerification API operation.
type StartVpcEndpointServicePrivateDnsVerificationRequest struct {
	*aws.Request
	Input *StartVpcEndpointServicePrivateDnsVerificationInput
	Copy  func(*StartVpcEndpointServicePrivateDnsVerificationInput) StartVpcEndpointServicePrivateDnsVerificationRequest
}

// Send marshals and sends the StartVpcEndpointServicePrivateDnsVerification API request.
func (r StartVpcEndpointServicePrivateDnsVerificationRequest) Send(ctx context.Context) (*StartVpcEndpointServicePrivateDnsVerificationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartVpcEndpointServicePrivateDnsVerificationResponse{
		StartVpcEndpointServicePrivateDnsVerificationOutput: r.Request.Data.(*StartVpcEndpointServicePrivateDnsVerificationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartVpcEndpointServicePrivateDnsVerificationResponse is the response type for the
// StartVpcEndpointServicePrivateDnsVerification API operation.
type StartVpcEndpointServicePrivateDnsVerificationResponse struct {
	*StartVpcEndpointServicePrivateDnsVerificationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartVpcEndpointServicePrivateDnsVerification request.
func (r *StartVpcEndpointServicePrivateDnsVerificationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
