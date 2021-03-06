// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/service/s3/internal/arn"
)

type AbortMultipartUploadInput struct {
	_ struct{} `type:"structure"`

	// The bucket name to which the upload was taking place.
	//
	// When using this API with an access point, you must direct requests to the
	// access point hostname. The access point hostname takes the form AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com.
	// When using this operation using an access point through the AWS SDKs, you
	// provide the access point ARN in place of the bucket name. For more information
	// about access point ARNs, see Using Access Points (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html)
	// in the Amazon Simple Storage Service Developer Guide.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Key of the object for which the multipart upload was initiated.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// Upload ID that identifies the multipart upload.
	//
	// UploadId is a required field
	UploadId *string `location:"querystring" locationName:"uploadId" type:"string" required:"true"`
}

// String returns the string representation
func (s AbortMultipartUploadInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AbortMultipartUploadInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AbortMultipartUploadInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.UploadId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UploadId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *AbortMultipartUploadInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AbortMultipartUploadInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.RequestPayer) > 0 {
		v := s.RequestPayer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-payer", v, metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Key", protocol.StringValue(v), metadata)
	}
	if s.UploadId != nil {
		v := *s.UploadId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "uploadId", protocol.StringValue(v), metadata)
	}
	return nil
}

func (s *AbortMultipartUploadInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *AbortMultipartUploadInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type AbortMultipartUploadOutput struct {
	_ struct{} `type:"structure"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`
}

// String returns the string representation
func (s AbortMultipartUploadOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AbortMultipartUploadOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	return nil
}

const opAbortMultipartUpload = "AbortMultipartUpload"

// AbortMultipartUploadRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// This operation aborts a multipart upload. After a multipart upload is aborted,
// no additional parts can be uploaded using that upload ID. The storage consumed
// by any previously uploaded parts will be freed. However, if any part uploads
// are currently in progress, those part uploads might or might not succeed.
// As a result, it might be necessary to abort a given multipart upload multiple
// times in order to completely free all storage consumed by all parts.
//
// To verify that all parts have been removed, so you don't get charged for
// the part storage, you should call the ListParts operation and ensure that
// the parts list is empty.
//
// For information about permissions required to use the multipart upload API,
// see Multipart Upload API and Permissions (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html).
//
// The following operations are related to AbortMultipartUpload:
//
//    * CreateMultipartUpload
//
//    * UploadPart
//
//    * CompleteMultipartUpload
//
//    * ListParts
//
//    * ListMultipartUploads
//
//    // Example sending a request using AbortMultipartUploadRequest.
//    req := client.AbortMultipartUploadRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/AbortMultipartUpload
func (c *Client) AbortMultipartUploadRequest(input *AbortMultipartUploadInput) AbortMultipartUploadRequest {
	op := &aws.Operation{
		Name:       opAbortMultipartUpload,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &AbortMultipartUploadInput{}
	}

	req := c.newRequest(op, input, &AbortMultipartUploadOutput{})

	return AbortMultipartUploadRequest{Request: req, Input: input, Copy: c.AbortMultipartUploadRequest}
}

// AbortMultipartUploadRequest is the request type for the
// AbortMultipartUpload API operation.
type AbortMultipartUploadRequest struct {
	*aws.Request
	Input *AbortMultipartUploadInput
	Copy  func(*AbortMultipartUploadInput) AbortMultipartUploadRequest
}

// Send marshals and sends the AbortMultipartUpload API request.
func (r AbortMultipartUploadRequest) Send(ctx context.Context) (*AbortMultipartUploadResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AbortMultipartUploadResponse{
		AbortMultipartUploadOutput: r.Request.Data.(*AbortMultipartUploadOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AbortMultipartUploadResponse is the response type for the
// AbortMultipartUpload API operation.
type AbortMultipartUploadResponse struct {
	*AbortMultipartUploadOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AbortMultipartUpload request.
func (r *AbortMultipartUploadResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
