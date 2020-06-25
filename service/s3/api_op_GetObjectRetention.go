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

type GetObjectRetentionInput struct {
	_ struct{} `type:"structure"`

	// The bucket name containing the object whose retention settings you want to
	// retrieve.
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

	// The key name for the object whose retention settings you want to retrieve.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// The version ID for the object whose retention settings you want to retrieve.
	VersionId *string `location:"querystring" locationName:"versionId" type:"string"`
}

// String returns the string representation
func (s GetObjectRetentionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetObjectRetentionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetObjectRetentionInput"}

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

func (s *GetObjectRetentionInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectRetentionInput) MarshalFields(e protocol.FieldEncoder) error {

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
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "versionId", protocol.StringValue(v), metadata)
	}
	return nil
}

func (s *GetObjectRetentionInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *GetObjectRetentionInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type GetObjectRetentionOutput struct {
	_ struct{} `type:"structure" payload:"Retention"`

	// The container element for an object's retention settings.
	Retention *ObjectLockRetention `type:"structure"`
}

// String returns the string representation
func (s GetObjectRetentionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectRetentionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Retention != nil {
		v := s.Retention

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "Retention", v, metadata)
	}
	return nil
}

const opGetObjectRetention = "GetObjectRetention"

// GetObjectRetentionRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Retrieves an object's retention settings. For more information, see Locking
// Objects (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html).
//
//    // Example sending a request using GetObjectRetentionRequest.
//    req := client.GetObjectRetentionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetObjectRetention
func (c *Client) GetObjectRetentionRequest(input *GetObjectRetentionInput) GetObjectRetentionRequest {
	op := &aws.Operation{
		Name:       opGetObjectRetention,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}/{Key+}?retention",
	}

	if input == nil {
		input = &GetObjectRetentionInput{}
	}

	req := c.newRequest(op, input, &GetObjectRetentionOutput{})

	return GetObjectRetentionRequest{Request: req, Input: input, Copy: c.GetObjectRetentionRequest}
}

// GetObjectRetentionRequest is the request type for the
// GetObjectRetention API operation.
type GetObjectRetentionRequest struct {
	*aws.Request
	Input *GetObjectRetentionInput
	Copy  func(*GetObjectRetentionInput) GetObjectRetentionRequest
}

// Send marshals and sends the GetObjectRetention API request.
func (r GetObjectRetentionRequest) Send(ctx context.Context) (*GetObjectRetentionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetObjectRetentionResponse{
		GetObjectRetentionOutput: r.Request.Data.(*GetObjectRetentionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetObjectRetentionResponse is the response type for the
// GetObjectRetention API operation.
type GetObjectRetentionResponse struct {
	*GetObjectRetentionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetObjectRetention request.
func (r *GetObjectRetentionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
