// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/checksum"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/service/s3/internal/arn"
)

type PutObjectRetentionInput struct {
	_ struct{} `type:"structure" payload:"Retention"`

	// The bucket name that contains the object you want to apply this Object Retention
	// configuration to.
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

	// Indicates whether this operation should bypass Governance-mode restrictions.
	BypassGovernanceRetention *bool `location:"header" locationName:"x-amz-bypass-governance-retention" type:"boolean"`

	// The key name for the object that you want to apply this Object Retention
	// configuration to.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// The container element for the Object Retention configuration.
	Retention *ObjectLockRetention `locationName:"Retention" type:"structure" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`

	// The version ID for the object that you want to apply this Object Retention
	// configuration to.
	VersionId *string `location:"querystring" locationName:"versionId" type:"string"`
}

// String returns the string representation
func (s PutObjectRetentionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutObjectRetentionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutObjectRetentionInput"}

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

func (s *PutObjectRetentionInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutObjectRetentionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.BypassGovernanceRetention != nil {
		v := *s.BypassGovernanceRetention

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-bypass-governance-retention", protocol.BoolValue(v), metadata)
	}
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
	if s.Retention != nil {
		v := s.Retention

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "Retention", v, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "versionId", protocol.StringValue(v), metadata)
	}
	return nil
}

func (s *PutObjectRetentionInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *PutObjectRetentionInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type PutObjectRetentionOutput struct {
	_ struct{} `type:"structure"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`
}

// String returns the string representation
func (s PutObjectRetentionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutObjectRetentionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	return nil
}

const opPutObjectRetention = "PutObjectRetention"

// PutObjectRetentionRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Places an Object Retention configuration on an object.
//
// Related Resources
//
//    * Locking Objects (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html)
//
//    // Example sending a request using PutObjectRetentionRequest.
//    req := client.PutObjectRetentionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutObjectRetention
func (c *Client) PutObjectRetentionRequest(input *PutObjectRetentionInput) PutObjectRetentionRequest {
	op := &aws.Operation{
		Name:       opPutObjectRetention,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}?retention",
	}

	if input == nil {
		input = &PutObjectRetentionInput{}
	}

	req := c.newRequest(op, input, &PutObjectRetentionOutput{})

	req.Handlers.Build.PushBackNamed(aws.NamedHandler{
		Name: "contentMd5Handler",
		Fn:   checksum.AddBodyContentMD5Handler,
	})

	return PutObjectRetentionRequest{Request: req, Input: input, Copy: c.PutObjectRetentionRequest}
}

// PutObjectRetentionRequest is the request type for the
// PutObjectRetention API operation.
type PutObjectRetentionRequest struct {
	*aws.Request
	Input *PutObjectRetentionInput
	Copy  func(*PutObjectRetentionInput) PutObjectRetentionRequest
}

// Send marshals and sends the PutObjectRetention API request.
func (r PutObjectRetentionRequest) Send(ctx context.Context) (*PutObjectRetentionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutObjectRetentionResponse{
		PutObjectRetentionOutput: r.Request.Data.(*PutObjectRetentionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutObjectRetentionResponse is the response type for the
// PutObjectRetention API operation.
type PutObjectRetentionResponse struct {
	*PutObjectRetentionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutObjectRetention request.
func (r *PutObjectRetentionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
