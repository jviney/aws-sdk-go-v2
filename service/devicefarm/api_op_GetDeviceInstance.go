// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetDeviceInstanceInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the instance you're requesting information
	// about.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDeviceInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDeviceInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDeviceInstanceInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetDeviceInstanceOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains information about your device instance.
	DeviceInstance *DeviceInstance `locationName:"deviceInstance" type:"structure"`
}

// String returns the string representation
func (s GetDeviceInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDeviceInstance = "GetDeviceInstance"

// GetDeviceInstanceRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Returns information about a device instance that belongs to a private device
// fleet.
//
//    // Example sending a request using GetDeviceInstanceRequest.
//    req := client.GetDeviceInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/GetDeviceInstance
func (c *Client) GetDeviceInstanceRequest(input *GetDeviceInstanceInput) GetDeviceInstanceRequest {
	op := &aws.Operation{
		Name:       opGetDeviceInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDeviceInstanceInput{}
	}

	req := c.newRequest(op, input, &GetDeviceInstanceOutput{})

	return GetDeviceInstanceRequest{Request: req, Input: input, Copy: c.GetDeviceInstanceRequest}
}

// GetDeviceInstanceRequest is the request type for the
// GetDeviceInstance API operation.
type GetDeviceInstanceRequest struct {
	*aws.Request
	Input *GetDeviceInstanceInput
	Copy  func(*GetDeviceInstanceInput) GetDeviceInstanceRequest
}

// Send marshals and sends the GetDeviceInstance API request.
func (r GetDeviceInstanceRequest) Send(ctx context.Context) (*GetDeviceInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDeviceInstanceResponse{
		GetDeviceInstanceOutput: r.Request.Data.(*GetDeviceInstanceOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDeviceInstanceResponse is the response type for the
// GetDeviceInstance API operation.
type GetDeviceInstanceResponse struct {
	*GetDeviceInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDeviceInstance request.
func (r *GetDeviceInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
