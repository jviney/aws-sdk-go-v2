// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing the Amazon Resource Name (ARN) of the gateway to
// start.
type StartGatewayInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s StartGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartGatewayInput"}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway that
// was restarted.
type StartGatewayOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s StartGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartGateway = "StartGateway"

// StartGatewayRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Starts a gateway that you previously shut down (see ShutdownGateway). After
// the gateway starts, you can then make other API calls, your applications
// can read from or write to the gateway's storage volumes and you will be able
// to take snapshot backups.
//
// When you make a request, you will get a 200 OK success response immediately.
// However, it might take some time for the gateway to be ready. You should
// call DescribeGatewayInformation and check the status before making any additional
// API calls. For more information, see ActivateGateway.
//
// To specify which gateway to start, use the Amazon Resource Name (ARN) of
// the gateway in your request.
//
//    // Example sending a request using StartGatewayRequest.
//    req := client.StartGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/StartGateway
func (c *Client) StartGatewayRequest(input *StartGatewayInput) StartGatewayRequest {
	op := &aws.Operation{
		Name:       opStartGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartGatewayInput{}
	}

	req := c.newRequest(op, input, &StartGatewayOutput{})

	return StartGatewayRequest{Request: req, Input: input, Copy: c.StartGatewayRequest}
}

// StartGatewayRequest is the request type for the
// StartGateway API operation.
type StartGatewayRequest struct {
	*aws.Request
	Input *StartGatewayInput
	Copy  func(*StartGatewayInput) StartGatewayRequest
}

// Send marshals and sends the StartGateway API request.
func (r StartGatewayRequest) Send(ctx context.Context) (*StartGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartGatewayResponse{
		StartGatewayOutput: r.Request.Data.(*StartGatewayOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartGatewayResponse is the response type for the
// StartGateway API operation.
type StartGatewayResponse struct {
	*StartGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartGateway request.
func (r *StartGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
