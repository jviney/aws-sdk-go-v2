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

type DeletePublicAccessBlockInput struct {
	_ struct{} `type:"structure"`

	// The Amazon S3 bucket whose PublicAccessBlock configuration you want to delete.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePublicAccessBlockInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePublicAccessBlockInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePublicAccessBlockInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *DeletePublicAccessBlockInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeletePublicAccessBlockInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

func (s *DeletePublicAccessBlockInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *DeletePublicAccessBlockInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type DeletePublicAccessBlockOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeletePublicAccessBlockOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeletePublicAccessBlockOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeletePublicAccessBlock = "DeletePublicAccessBlock"

// DeletePublicAccessBlockRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Removes the PublicAccessBlock configuration for an Amazon S3 bucket. To use
// this operation, you must have the s3:PutBucketPublicAccessBlock permission.
// For more information about permissions, see Permissions Related to Bucket
// Subresource Operations (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html).
//
// The following operations are related to DeletePublicAccessBlock:
//
//    * Using Amazon S3 Block Public Access (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html)
//
//    * GetPublicAccessBlock
//
//    * PutPublicAccessBlock
//
//    * GetBucketPolicyStatus
//
//    // Example sending a request using DeletePublicAccessBlockRequest.
//    req := client.DeletePublicAccessBlockRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeletePublicAccessBlock
func (c *Client) DeletePublicAccessBlockRequest(input *DeletePublicAccessBlockInput) DeletePublicAccessBlockRequest {
	op := &aws.Operation{
		Name:       opDeletePublicAccessBlock,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?publicAccessBlock",
	}

	if input == nil {
		input = &DeletePublicAccessBlockInput{}
	}

	req := c.newRequest(op, input, &DeletePublicAccessBlockOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeletePublicAccessBlockRequest{Request: req, Input: input, Copy: c.DeletePublicAccessBlockRequest}
}

// DeletePublicAccessBlockRequest is the request type for the
// DeletePublicAccessBlock API operation.
type DeletePublicAccessBlockRequest struct {
	*aws.Request
	Input *DeletePublicAccessBlockInput
	Copy  func(*DeletePublicAccessBlockInput) DeletePublicAccessBlockRequest
}

// Send marshals and sends the DeletePublicAccessBlock API request.
func (r DeletePublicAccessBlockRequest) Send(ctx context.Context) (*DeletePublicAccessBlockResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePublicAccessBlockResponse{
		DeletePublicAccessBlockOutput: r.Request.Data.(*DeletePublicAccessBlockOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePublicAccessBlockResponse is the response type for the
// DeletePublicAccessBlock API operation.
type DeletePublicAccessBlockResponse struct {
	*DeletePublicAccessBlockOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePublicAccessBlock request.
func (r *DeletePublicAccessBlockResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
