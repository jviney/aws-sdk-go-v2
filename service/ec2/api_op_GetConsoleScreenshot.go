// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetConsoleScreenshotInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the instance.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// When set to true, acts as keystroke input and wakes up an instance that's
	// in standby or "sleep" mode.
	WakeUp *bool `type:"boolean"`
}

// String returns the string representation
func (s GetConsoleScreenshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetConsoleScreenshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetConsoleScreenshotInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetConsoleScreenshotOutput struct {
	_ struct{} `type:"structure"`

	// The data that comprises the image.
	ImageData *string `locationName:"imageData" type:"string"`

	// The ID of the instance.
	InstanceId *string `locationName:"instanceId" type:"string"`
}

// String returns the string representation
func (s GetConsoleScreenshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetConsoleScreenshot = "GetConsoleScreenshot"

// GetConsoleScreenshotRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Retrieve a JPG-format screenshot of a running instance to help with troubleshooting.
//
// The returned content is Base64-encoded.
//
//    // Example sending a request using GetConsoleScreenshotRequest.
//    req := client.GetConsoleScreenshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetConsoleScreenshot
func (c *Client) GetConsoleScreenshotRequest(input *GetConsoleScreenshotInput) GetConsoleScreenshotRequest {
	op := &aws.Operation{
		Name:       opGetConsoleScreenshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetConsoleScreenshotInput{}
	}

	req := c.newRequest(op, input, &GetConsoleScreenshotOutput{})

	return GetConsoleScreenshotRequest{Request: req, Input: input, Copy: c.GetConsoleScreenshotRequest}
}

// GetConsoleScreenshotRequest is the request type for the
// GetConsoleScreenshot API operation.
type GetConsoleScreenshotRequest struct {
	*aws.Request
	Input *GetConsoleScreenshotInput
	Copy  func(*GetConsoleScreenshotInput) GetConsoleScreenshotRequest
}

// Send marshals and sends the GetConsoleScreenshot API request.
func (r GetConsoleScreenshotRequest) Send(ctx context.Context) (*GetConsoleScreenshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetConsoleScreenshotResponse{
		GetConsoleScreenshotOutput: r.Request.Data.(*GetConsoleScreenshotOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetConsoleScreenshotResponse is the response type for the
// GetConsoleScreenshot API operation.
type GetConsoleScreenshotResponse struct {
	*GetConsoleScreenshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetConsoleScreenshot request.
func (r *GetConsoleScreenshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
