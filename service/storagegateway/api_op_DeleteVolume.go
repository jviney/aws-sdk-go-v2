// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing the DeleteVolumeInput$VolumeARN to delete.
type DeleteVolumeInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the volume. Use the ListVolumes operation
	// to return a list of gateway volumes.
	//
	// VolumeARN is a required field
	VolumeARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVolumeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVolumeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVolumeInput"}

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

// A JSON object containing the Amazon Resource Name (ARN) of the storage volume
// that was deleted
type DeleteVolumeOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the storage volume that was deleted. It
	// is the same ARN you provided in the request.
	VolumeARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s DeleteVolumeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteVolume = "DeleteVolume"

// DeleteVolumeRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Deletes the specified storage volume that you previously created using the
// CreateCachediSCSIVolume or CreateStorediSCSIVolume API. This operation is
// only supported in the cached volume and stored volume types. For stored volume
// gateways, the local disk that was configured as the storage volume is not
// deleted. You can reuse the local disk to create another storage volume.
//
// Before you delete a volume, make sure there are no iSCSI connections to the
// volume you are deleting. You should also make sure there is no snapshot in
// progress. You can use the Amazon Elastic Compute Cloud (Amazon EC2) API to
// query snapshots on the volume you are deleting and check the snapshot status.
// For more information, go to DescribeSnapshots (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSnapshots.html)
// in the Amazon Elastic Compute Cloud API Reference.
//
// In the request, you must provide the Amazon Resource Name (ARN) of the storage
// volume you want to delete.
//
//    // Example sending a request using DeleteVolumeRequest.
//    req := client.DeleteVolumeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DeleteVolume
func (c *Client) DeleteVolumeRequest(input *DeleteVolumeInput) DeleteVolumeRequest {
	op := &aws.Operation{
		Name:       opDeleteVolume,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteVolumeInput{}
	}

	req := c.newRequest(op, input, &DeleteVolumeOutput{})

	return DeleteVolumeRequest{Request: req, Input: input, Copy: c.DeleteVolumeRequest}
}

// DeleteVolumeRequest is the request type for the
// DeleteVolume API operation.
type DeleteVolumeRequest struct {
	*aws.Request
	Input *DeleteVolumeInput
	Copy  func(*DeleteVolumeInput) DeleteVolumeRequest
}

// Send marshals and sends the DeleteVolume API request.
func (r DeleteVolumeRequest) Send(ctx context.Context) (*DeleteVolumeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVolumeResponse{
		DeleteVolumeOutput: r.Request.Data.(*DeleteVolumeOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVolumeResponse is the response type for the
// DeleteVolume API operation.
type DeleteVolumeResponse struct {
	*DeleteVolumeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVolume request.
func (r *DeleteVolumeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
