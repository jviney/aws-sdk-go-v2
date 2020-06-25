// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DetectCustomLabelsInput struct {
	_ struct{} `type:"structure"`

	// Provides the input image either as bytes or an S3 object.
	//
	// You pass image bytes to an Amazon Rekognition API operation by using the
	// Bytes property. For example, you would use the Bytes property to pass an
	// image loaded from a local file system. Image bytes passed by using the Bytes
	// property must be base64-encoded. Your code may not need to encode image bytes
	// if you are using an AWS SDK to call Amazon Rekognition API operations.
	//
	// For more information, see Analyzing an Image Loaded from a Local File System
	// in the Amazon Rekognition Developer Guide.
	//
	// You pass images stored in an S3 bucket to an Amazon Rekognition API operation
	// by using the S3Object property. Images stored in an S3 bucket do not need
	// to be base64-encoded.
	//
	// The region for the S3 bucket containing the S3 object must match the region
	// you use for Amazon Rekognition operations.
	//
	// If you use the AWS CLI to call Amazon Rekognition operations, passing image
	// bytes using the Bytes property is not supported. You must first upload the
	// image to an Amazon S3 bucket and then call the operation using the S3Object
	// property.
	//
	// For Amazon Rekognition to process an S3 object, the user must have permission
	// to access the S3 object. For more information, see Resource Based Policies
	// in the Amazon Rekognition Developer Guide.
	//
	// Image is a required field
	Image *Image `type:"structure" required:"true"`

	// Maximum number of results you want the service to return in the response.
	// The service returns the specified number of highest confidence labels ranked
	// from highest confidence to lowest.
	MaxResults *int64 `type:"integer"`

	// Specifies the minimum confidence level for the labels to return. Amazon Rekognition
	// doesn't return any labels with a confidence lower than this specified value.
	// If you specify a value of 0, all labels are return, regardless of the default
	// thresholds that the model version applies.
	MinConfidence *float64 `type:"float"`

	// The ARN of the model version that you want to use.
	//
	// ProjectVersionArn is a required field
	ProjectVersionArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DetectCustomLabelsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetectCustomLabelsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetectCustomLabelsInput"}

	if s.Image == nil {
		invalidParams.Add(aws.NewErrParamRequired("Image"))
	}

	if s.ProjectVersionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectVersionArn"))
	}
	if s.ProjectVersionArn != nil && len(*s.ProjectVersionArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectVersionArn", 20))
	}
	if s.Image != nil {
		if err := s.Image.Validate(); err != nil {
			invalidParams.AddNested("Image", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DetectCustomLabelsOutput struct {
	_ struct{} `type:"structure"`

	// An array of custom labels detected in the input image.
	CustomLabels []CustomLabel `type:"list"`
}

// String returns the string representation
func (s DetectCustomLabelsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDetectCustomLabels = "DetectCustomLabels"

// DetectCustomLabelsRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Detects custom labels in a supplied image by using an Amazon Rekognition
// Custom Labels model.
//
// You specify which version of a model version to use by using the ProjectVersionArn
// input parameter.
//
// You pass the input image as base64-encoded image bytes or as a reference
// to an image in an Amazon S3 bucket. If you use the AWS CLI to call Amazon
// Rekognition operations, passing image bytes is not supported. The image must
// be either a PNG or JPEG formatted file.
//
// For each object that the model version detects on an image, the API returns
// a (CustomLabel) object in an array (CustomLabels). Each CustomLabel object
// provides the label name (Name), the level of confidence that the image contains
// the object (Confidence), and object location information, if it exists, for
// the label on the image (Geometry).
//
// During training model calculates a threshold value that determines if a prediction
// for a label is true. By default, DetectCustomLabels doesn't return labels
// whose confidence value is below the model's calculated threshold value. To
// filter labels that are returned, specify a value for MinConfidence that is
// higher than the model's calculated threshold. You can get the model's calculated
// threshold from the model's training results shown in the Amazon Rekognition
// Custom Labels console. To get all labels, regardless of confidence, specify
// a MinConfidence value of 0.
//
// You can also add the MaxResults parameter to limit the number of labels returned.
//
// This is a stateless API operation. That is, the operation does not persist
// any data.
//
// This operation requires permissions to perform the rekognition:DetectCustomLabels
// action.
//
//    // Example sending a request using DetectCustomLabelsRequest.
//    req := client.DetectCustomLabelsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DetectCustomLabelsRequest(input *DetectCustomLabelsInput) DetectCustomLabelsRequest {
	op := &aws.Operation{
		Name:       opDetectCustomLabels,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetectCustomLabelsInput{}
	}

	req := c.newRequest(op, input, &DetectCustomLabelsOutput{})

	return DetectCustomLabelsRequest{Request: req, Input: input, Copy: c.DetectCustomLabelsRequest}
}

// DetectCustomLabelsRequest is the request type for the
// DetectCustomLabels API operation.
type DetectCustomLabelsRequest struct {
	*aws.Request
	Input *DetectCustomLabelsInput
	Copy  func(*DetectCustomLabelsInput) DetectCustomLabelsRequest
}

// Send marshals and sends the DetectCustomLabels API request.
func (r DetectCustomLabelsRequest) Send(ctx context.Context) (*DetectCustomLabelsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetectCustomLabelsResponse{
		DetectCustomLabelsOutput: r.Request.Data.(*DetectCustomLabelsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetectCustomLabelsResponse is the response type for the
// DetectCustomLabels API operation.
type DetectCustomLabelsResponse struct {
	*DetectCustomLabelsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetectCustomLabels request.
func (r *DetectCustomLabelsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
