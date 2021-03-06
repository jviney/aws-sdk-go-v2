// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/jviney/aws-sdk-go-v2/service/s3/internal/arn"
)

type DeleteBucketWebsiteInput struct {
	_ struct{} `type:"structure"`

	// The bucket name for which you want to remove the website configuration.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBucketWebsiteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBucketWebsiteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBucketWebsiteInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *DeleteBucketWebsiteInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBucketWebsiteInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

func (s *DeleteBucketWebsiteInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *DeleteBucketWebsiteInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type DeleteBucketWebsiteOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteBucketWebsiteOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBucketWebsiteOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteBucketWebsite = "DeleteBucketWebsite"

// DeleteBucketWebsiteRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// This operation removes the website configuration for a bucket. Amazon S3
// returns a 200 OK response upon successfully deleting a website configuration
// on the specified bucket. You will get a 200 OK response if the website configuration
// you are trying to delete does not exist on the bucket. Amazon S3 returns
// a 404 response if the bucket specified in the request does not exist.
//
// This DELETE operation requires the S3:DeleteBucketWebsite permission. By
// default, only the bucket owner can delete the website configuration attached
// to a bucket. However, bucket owners can grant other users permission to delete
// the website configuration by writing a bucket policy granting them the S3:DeleteBucketWebsite
// permission.
//
// For more information about hosting websites, see Hosting Websites on Amazon
// S3 (https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html).
//
// The following operations are related to DeleteBucketWebsite:
//
//    * GetBucketWebsite
//
//    * PutBucketWebsite
//
//    // Example sending a request using DeleteBucketWebsiteRequest.
//    req := client.DeleteBucketWebsiteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteBucketWebsite
func (c *Client) DeleteBucketWebsiteRequest(input *DeleteBucketWebsiteInput) DeleteBucketWebsiteRequest {
	op := &aws.Operation{
		Name:       opDeleteBucketWebsite,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?website",
	}

	if input == nil {
		input = &DeleteBucketWebsiteInput{}
	}

	req := c.newRequest(op, input, &DeleteBucketWebsiteOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteBucketWebsiteRequest{Request: req, Input: input, Copy: c.DeleteBucketWebsiteRequest}
}

// DeleteBucketWebsiteRequest is the request type for the
// DeleteBucketWebsite API operation.
type DeleteBucketWebsiteRequest struct {
	*aws.Request
	Input *DeleteBucketWebsiteInput
	Copy  func(*DeleteBucketWebsiteInput) DeleteBucketWebsiteRequest
}

// Send marshals and sends the DeleteBucketWebsite API request.
func (r DeleteBucketWebsiteRequest) Send(ctx context.Context) (*DeleteBucketWebsiteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBucketWebsiteResponse{
		DeleteBucketWebsiteOutput: r.Request.Data.(*DeleteBucketWebsiteOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBucketWebsiteResponse is the response type for the
// DeleteBucketWebsite API operation.
type DeleteBucketWebsiteResponse struct {
	*DeleteBucketWebsiteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBucketWebsite request.
func (r *DeleteBucketWebsiteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
