// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Input for the OptInPhoneNumber action.
type OptInPhoneNumberInput struct {
	_ struct{} `type:"structure"`

	// The phone number to opt in.
	//
	// PhoneNumber is a required field
	PhoneNumber *string `locationName:"phoneNumber" type:"string" required:"true"`
}

// String returns the string representation
func (s OptInPhoneNumberInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *OptInPhoneNumberInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "OptInPhoneNumberInput"}

	if s.PhoneNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("PhoneNumber"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The response for the OptInPhoneNumber action.
type OptInPhoneNumberOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s OptInPhoneNumberOutput) String() string {
	return awsutil.Prettify(s)
}

const opOptInPhoneNumber = "OptInPhoneNumber"

// OptInPhoneNumberRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Use this request to opt in a phone number that is opted out, which enables
// you to resume sending SMS messages to the number.
//
// You can opt in a phone number only once every 30 days.
//
//    // Example sending a request using OptInPhoneNumberRequest.
//    req := client.OptInPhoneNumberRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/OptInPhoneNumber
func (c *Client) OptInPhoneNumberRequest(input *OptInPhoneNumberInput) OptInPhoneNumberRequest {
	op := &aws.Operation{
		Name:       opOptInPhoneNumber,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &OptInPhoneNumberInput{}
	}

	req := c.newRequest(op, input, &OptInPhoneNumberOutput{})

	return OptInPhoneNumberRequest{Request: req, Input: input, Copy: c.OptInPhoneNumberRequest}
}

// OptInPhoneNumberRequest is the request type for the
// OptInPhoneNumber API operation.
type OptInPhoneNumberRequest struct {
	*aws.Request
	Input *OptInPhoneNumberInput
	Copy  func(*OptInPhoneNumberInput) OptInPhoneNumberRequest
}

// Send marshals and sends the OptInPhoneNumber API request.
func (r OptInPhoneNumberRequest) Send(ctx context.Context) (*OptInPhoneNumberResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &OptInPhoneNumberResponse{
		OptInPhoneNumberOutput: r.Request.Data.(*OptInPhoneNumberOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// OptInPhoneNumberResponse is the response type for the
// OptInPhoneNumber API operation.
type OptInPhoneNumberResponse struct {
	*OptInPhoneNumberOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// OptInPhoneNumber request.
func (r *OptInPhoneNumberResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
