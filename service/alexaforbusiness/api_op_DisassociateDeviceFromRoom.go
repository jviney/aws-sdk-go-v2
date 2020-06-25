// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DisassociateDeviceFromRoomInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the device to disassociate from a room. Required.
	DeviceArn *string `type:"string"`
}

// String returns the string representation
func (s DisassociateDeviceFromRoomInput) String() string {
	return awsutil.Prettify(s)
}

type DisassociateDeviceFromRoomOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateDeviceFromRoomOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateDeviceFromRoom = "DisassociateDeviceFromRoom"

// DisassociateDeviceFromRoomRequest returns a request value for making API operation for
// Alexa For Business.
//
// Disassociates a device from its current room. The device continues to be
// connected to the Wi-Fi network and is still registered to the account. The
// device settings and skills are removed from the room.
//
//    // Example sending a request using DisassociateDeviceFromRoomRequest.
//    req := client.DisassociateDeviceFromRoomRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DisassociateDeviceFromRoom
func (c *Client) DisassociateDeviceFromRoomRequest(input *DisassociateDeviceFromRoomInput) DisassociateDeviceFromRoomRequest {
	op := &aws.Operation{
		Name:       opDisassociateDeviceFromRoom,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateDeviceFromRoomInput{}
	}

	req := c.newRequest(op, input, &DisassociateDeviceFromRoomOutput{})

	return DisassociateDeviceFromRoomRequest{Request: req, Input: input, Copy: c.DisassociateDeviceFromRoomRequest}
}

// DisassociateDeviceFromRoomRequest is the request type for the
// DisassociateDeviceFromRoom API operation.
type DisassociateDeviceFromRoomRequest struct {
	*aws.Request
	Input *DisassociateDeviceFromRoomInput
	Copy  func(*DisassociateDeviceFromRoomInput) DisassociateDeviceFromRoomRequest
}

// Send marshals and sends the DisassociateDeviceFromRoom API request.
func (r DisassociateDeviceFromRoomRequest) Send(ctx context.Context) (*DisassociateDeviceFromRoomResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateDeviceFromRoomResponse{
		DisassociateDeviceFromRoomOutput: r.Request.Data.(*DisassociateDeviceFromRoomOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateDeviceFromRoomResponse is the response type for the
// DisassociateDeviceFromRoom API operation.
type DisassociateDeviceFromRoomResponse struct {
	*DisassociateDeviceFromRoomOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateDeviceFromRoom request.
func (r *DisassociateDeviceFromRoomResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
