// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeTypeRegistrationInput struct {
	_ struct{} `type:"structure"`

	// The identifier for this registration request.
	//
	// This registration token is generated by CloudFormation when you initiate
	// a registration request using RegisterType .
	//
	// RegistrationToken is a required field
	RegistrationToken *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeTypeRegistrationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTypeRegistrationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTypeRegistrationInput"}

	if s.RegistrationToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("RegistrationToken"))
	}
	if s.RegistrationToken != nil && len(*s.RegistrationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RegistrationToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeTypeRegistrationOutput struct {
	_ struct{} `type:"structure"`

	// The description of the type registration request.
	Description *string `min:"1" type:"string"`

	// The current status of the type registration request.
	ProgressStatus RegistrationStatus `type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) of the type being registered.
	//
	// For registration requests with a ProgressStatus of other than COMPLETE, this
	// will be null.
	TypeArn *string `type:"string"`

	// The Amazon Resource Name (ARN) of this specific version of the type being
	// registered.
	//
	// For registration requests with a ProgressStatus of other than COMPLETE, this
	// will be null.
	TypeVersionArn *string `type:"string"`
}

// String returns the string representation
func (s DescribeTypeRegistrationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTypeRegistration = "DescribeTypeRegistration"

// DescribeTypeRegistrationRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns information about a type's registration, including its current status
// and type and version identifiers.
//
// When you initiate a registration request using RegisterType , you can then
// use DescribeTypeRegistration to monitor the progress of that registration
// request.
//
// Once the registration request has completed, use DescribeType to return detailed
// informaiton about a type.
//
//    // Example sending a request using DescribeTypeRegistrationRequest.
//    req := client.DescribeTypeRegistrationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeTypeRegistration
func (c *Client) DescribeTypeRegistrationRequest(input *DescribeTypeRegistrationInput) DescribeTypeRegistrationRequest {
	op := &aws.Operation{
		Name:       opDescribeTypeRegistration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTypeRegistrationInput{}
	}

	req := c.newRequest(op, input, &DescribeTypeRegistrationOutput{})

	return DescribeTypeRegistrationRequest{Request: req, Input: input, Copy: c.DescribeTypeRegistrationRequest}
}

// DescribeTypeRegistrationRequest is the request type for the
// DescribeTypeRegistration API operation.
type DescribeTypeRegistrationRequest struct {
	*aws.Request
	Input *DescribeTypeRegistrationInput
	Copy  func(*DescribeTypeRegistrationInput) DescribeTypeRegistrationRequest
}

// Send marshals and sends the DescribeTypeRegistration API request.
func (r DescribeTypeRegistrationRequest) Send(ctx context.Context) (*DescribeTypeRegistrationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTypeRegistrationResponse{
		DescribeTypeRegistrationOutput: r.Request.Data.(*DescribeTypeRegistrationOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTypeRegistrationResponse is the response type for the
// DescribeTypeRegistration API operation.
type DescribeTypeRegistrationResponse struct {
	*DescribeTypeRegistrationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTypeRegistration request.
func (r *DescribeTypeRegistrationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
