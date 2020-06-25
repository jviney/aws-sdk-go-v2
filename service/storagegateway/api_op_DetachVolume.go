// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// AttachVolumeInput
type DetachVolumeInput struct {
	_ struct{} `type:"structure"`

	// Set to true to forcibly remove the iSCSI connection of the target volume
	// and detach the volume. The default is false. If this value is set to false,
	// you must manually disconnect the iSCSI connection from the target volume.
	ForceDetach *bool `type:"boolean"`

	// The Amazon Resource Name (ARN) of the volume to detach from the gateway.
	//
	// VolumeARN is a required field
	VolumeARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s DetachVolumeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachVolumeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetachVolumeInput"}

	if s.VolumeARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeARN"))
	}
	if s.VolumeARN != nil && len(*s.VolumeARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("VolumeARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// AttachVolumeOutput
type DetachVolumeOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the volume that was detached.
	VolumeARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s DetachVolumeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDetachVolume = "DetachVolume"

// DetachVolumeRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Disconnects a volume from an iSCSI connection and then detaches the volume
// from the specified gateway. Detaching and attaching a volume enables you
// to recover your data from one gateway to a different gateway without creating
// a snapshot. It also makes it easier to move your volumes from an on-premises
// gateway to a gateway hosted on an Amazon EC2 instance. This operation is
// only supported in the volume gateway type.
//
//    // Example sending a request using DetachVolumeRequest.
//    req := client.DetachVolumeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DetachVolume
func (c *Client) DetachVolumeRequest(input *DetachVolumeInput) DetachVolumeRequest {
	op := &aws.Operation{
		Name:       opDetachVolume,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachVolumeInput{}
	}

	req := c.newRequest(op, input, &DetachVolumeOutput{})

	return DetachVolumeRequest{Request: req, Input: input, Copy: c.DetachVolumeRequest}
}

// DetachVolumeRequest is the request type for the
// DetachVolume API operation.
type DetachVolumeRequest struct {
	*aws.Request
	Input *DetachVolumeInput
	Copy  func(*DetachVolumeInput) DetachVolumeRequest
}

// Send marshals and sends the DetachVolume API request.
func (r DetachVolumeRequest) Send(ctx context.Context) (*DetachVolumeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachVolumeResponse{
		DetachVolumeOutput: r.Request.Data.(*DetachVolumeOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachVolumeResponse is the response type for the
// DetachVolume API operation.
type DetachVolumeResponse struct {
	*DetachVolumeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachVolume request.
func (r *DetachVolumeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
