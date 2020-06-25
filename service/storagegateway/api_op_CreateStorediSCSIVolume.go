// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing one or more of the following fields:
//
//    * CreateStorediSCSIVolumeInput$DiskId
//
//    * CreateStorediSCSIVolumeInput$NetworkInterfaceId
//
//    * CreateStorediSCSIVolumeInput$PreserveExistingData
//
//    * CreateStorediSCSIVolumeInput$SnapshotId
//
//    * CreateStorediSCSIVolumeInput$TargetName
type CreateStorediSCSIVolumeInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the gateway local disk that is configured as a
	// stored volume. Use ListLocalDisks (https://docs.aws.amazon.com/storagegateway/latest/userguide/API_ListLocalDisks.html)
	// to list disk IDs for a gateway.
	//
	// DiskId is a required field
	DiskId *string `min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`

	// True to use Amazon S3 server-side encryption with your own AWS KMS key, or
	// false to use a key managed by Amazon S3. Optional.
	KMSEncrypted *bool `type:"boolean"`

	// The Amazon Resource Name (ARN) of the KMS key used for Amazon S3 server-side
	// encryption. This value can only be set when KMSEncrypted is true. Optional.
	KMSKey *string `min:"7" type:"string"`

	// The network interface of the gateway on which to expose the iSCSI target.
	// Only IPv4 addresses are accepted. Use DescribeGatewayInformation to get a
	// list of the network interfaces available on a gateway.
	//
	// Valid Values: A valid IP address.
	//
	// NetworkInterfaceId is a required field
	NetworkInterfaceId *string `type:"string" required:"true"`

	// Specify this field as true if you want to preserve the data on the local
	// disk. Otherwise, specifying this field as false creates an empty volume.
	//
	// Valid Values: true, false
	//
	// PreserveExistingData is a required field
	PreserveExistingData *bool `type:"boolean" required:"true"`

	// The snapshot ID (e.g. "snap-1122aabb") of the snapshot to restore as the
	// new stored volume. Specify this field if you want to create the iSCSI storage
	// volume from a snapshot otherwise do not include this field. To list snapshots
	// for your account use DescribeSnapshots (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSnapshots.html)
	// in the Amazon Elastic Compute Cloud API Reference.
	SnapshotId *string `type:"string"`

	// A list of up to 50 tags that can be assigned to a stored volume. Each tag
	// is a key-value pair.
	//
	// Valid characters for key and value are letters, spaces, and numbers representable
	// in UTF-8 format, and the following special characters: + - = . _ : / @. The
	// maximum length of a tag's key is 128 characters, and the maximum length for
	// a tag's value is 256.
	Tags []Tag `type:"list"`

	// The name of the iSCSI target used by an initiator to connect to a volume
	// and used as a suffix for the target ARN. For example, specifying TargetName
	// as myvolume results in the target ARN of arn:aws:storagegateway:us-east-2:111122223333:gateway/sgw-12A3456B/target/iqn.1997-05.com.amazon:myvolume.
	// The target name must be unique across all volumes on a gateway.
	//
	// If you don't specify a value, Storage Gateway uses the value that was previously
	// used for this volume as the new target name.
	//
	// TargetName is a required field
	TargetName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateStorediSCSIVolumeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateStorediSCSIVolumeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateStorediSCSIVolumeInput"}

	if s.DiskId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DiskId"))
	}
	if s.DiskId != nil && len(*s.DiskId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DiskId", 1))
	}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}
	if s.KMSKey != nil && len(*s.KMSKey) < 7 {
		invalidParams.Add(aws.NewErrParamMinLen("KMSKey", 7))
	}

	if s.NetworkInterfaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkInterfaceId"))
	}

	if s.PreserveExistingData == nil {
		invalidParams.Add(aws.NewErrParamRequired("PreserveExistingData"))
	}

	if s.TargetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetName"))
	}
	if s.TargetName != nil && len(*s.TargetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetName", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A JSON object containing the following fields:
type CreateStorediSCSIVolumeOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the volume target, which includes the iSCSI
	// name that initiators can use to connect to the target.
	TargetARN *string `min:"50" type:"string"`

	// The Amazon Resource Name (ARN) of the configured volume.
	VolumeARN *string `min:"50" type:"string"`

	// The size of the volume in bytes.
	VolumeSizeInBytes *int64 `type:"long"`
}

// String returns the string representation
func (s CreateStorediSCSIVolumeOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateStorediSCSIVolume = "CreateStorediSCSIVolume"

// CreateStorediSCSIVolumeRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Creates a volume on a specified gateway. This operation is only supported
// in the stored volume gateway type.
//
// The size of the volume to create is inferred from the disk size. You can
// choose to preserve existing data on the disk, create volume from an existing
// snapshot, or create an empty volume. If you choose to create an empty gateway
// volume, then any existing data on the disk is erased.
//
// In the request you must specify the gateway and the disk information on which
// you are creating the volume. In response, the gateway creates the volume
// and returns volume information such as the volume Amazon Resource Name (ARN),
// its size, and the iSCSI target ARN that initiators can use to connect to
// the volume target.
//
//    // Example sending a request using CreateStorediSCSIVolumeRequest.
//    req := client.CreateStorediSCSIVolumeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/CreateStorediSCSIVolume
func (c *Client) CreateStorediSCSIVolumeRequest(input *CreateStorediSCSIVolumeInput) CreateStorediSCSIVolumeRequest {
	op := &aws.Operation{
		Name:       opCreateStorediSCSIVolume,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateStorediSCSIVolumeInput{}
	}

	req := c.newRequest(op, input, &CreateStorediSCSIVolumeOutput{})

	return CreateStorediSCSIVolumeRequest{Request: req, Input: input, Copy: c.CreateStorediSCSIVolumeRequest}
}

// CreateStorediSCSIVolumeRequest is the request type for the
// CreateStorediSCSIVolume API operation.
type CreateStorediSCSIVolumeRequest struct {
	*aws.Request
	Input *CreateStorediSCSIVolumeInput
	Copy  func(*CreateStorediSCSIVolumeInput) CreateStorediSCSIVolumeRequest
}

// Send marshals and sends the CreateStorediSCSIVolume API request.
func (r CreateStorediSCSIVolumeRequest) Send(ctx context.Context) (*CreateStorediSCSIVolumeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateStorediSCSIVolumeResponse{
		CreateStorediSCSIVolumeOutput: r.Request.Data.(*CreateStorediSCSIVolumeOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateStorediSCSIVolumeResponse is the response type for the
// CreateStorediSCSIVolume API operation.
type CreateStorediSCSIVolumeResponse struct {
	*CreateStorediSCSIVolumeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateStorediSCSIVolume request.
func (r *CreateStorediSCSIVolumeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
