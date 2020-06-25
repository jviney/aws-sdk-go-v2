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

type DeleteObjectTaggingInput struct {
	_ struct{} `type:"structure"`

	// The bucket name containing the objects from which to remove the tags.
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

	// Name of the tag.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// The versionId of the object that the tag-set will be removed from.
	VersionId *string `location:"querystring" locationName:"versionId" type:"string"`
}

// String returns the string representation
func (s DeleteObjectTaggingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteObjectTaggingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteObjectTaggingInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *DeleteObjectTaggingInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteObjectTaggingInput) MarshalFields(e protocol.FieldEncoder) error {

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
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "versionId", protocol.StringValue(v), metadata)
	}
	return nil
}

func (s *DeleteObjectTaggingInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *DeleteObjectTaggingInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type DeleteObjectTaggingOutput struct {
	_ struct{} `type:"structure"`

	// The versionId of the object the tag-set was removed from.
	VersionId *string `location:"header" locationName:"x-amz-version-id" type:"string"`
}

// String returns the string representation
func (s DeleteObjectTaggingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteObjectTaggingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-version-id", protocol.StringValue(v), metadata)
	}
	return nil
}

const opDeleteObjectTagging = "DeleteObjectTagging"

// DeleteObjectTaggingRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Removes the entire tag set from the specified object. For more information
// about managing object tags, see Object Tagging (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-tagging.html).
//
// To use this operation, you must have permission to perform the s3:DeleteObjectTagging
// action.
//
// To delete tags of a specific object version, add the versionId query parameter
// in the request. You will need permission for the s3:DeleteObjectVersionTagging
// action.
//
// The following operations are related to DeleteBucketMetricsConfiguration:
//
//    * PutObjectTagging
//
//    * GetObjectTagging
//
//    // Example sending a request using DeleteObjectTaggingRequest.
//    req := client.DeleteObjectTaggingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteObjectTagging
func (c *Client) DeleteObjectTaggingRequest(input *DeleteObjectTaggingInput) DeleteObjectTaggingRequest {
	op := &aws.Operation{
		Name:       opDeleteObjectTagging,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}/{Key+}?tagging",
	}

	if input == nil {
		input = &DeleteObjectTaggingInput{}
	}

	req := c.newRequest(op, input, &DeleteObjectTaggingOutput{})

	return DeleteObjectTaggingRequest{Request: req, Input: input, Copy: c.DeleteObjectTaggingRequest}
}

// DeleteObjectTaggingRequest is the request type for the
// DeleteObjectTagging API operation.
type DeleteObjectTaggingRequest struct {
	*aws.Request
	Input *DeleteObjectTaggingInput
	Copy  func(*DeleteObjectTaggingInput) DeleteObjectTaggingRequest
}

// Send marshals and sends the DeleteObjectTagging API request.
func (r DeleteObjectTaggingRequest) Send(ctx context.Context) (*DeleteObjectTaggingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteObjectTaggingResponse{
		DeleteObjectTaggingOutput: r.Request.Data.(*DeleteObjectTaggingOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteObjectTaggingResponse is the response type for the
// DeleteObjectTagging API operation.
type DeleteObjectTaggingResponse struct {
	*DeleteObjectTaggingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteObjectTagging request.
func (r *DeleteObjectTaggingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
