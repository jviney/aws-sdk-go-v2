// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type InitiateLayerUploadInput struct {
	_ struct{} `type:"structure"`

	// The AWS account ID associated with the registry to which you intend to upload
	// layers. If you do not specify a registry, the default registry is assumed.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The name of the repository to which you intend to upload layers.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s InitiateLayerUploadInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InitiateLayerUploadInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "InitiateLayerUploadInput"}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type InitiateLayerUploadOutput struct {
	_ struct{} `type:"structure"`

	// The size, in bytes, that Amazon ECR expects future layer part uploads to
	// be.
	PartSize *int64 `locationName:"partSize" type:"long"`

	// The upload ID for the layer upload. This parameter is passed to further UploadLayerPart
	// and CompleteLayerUpload operations.
	UploadId *string `locationName:"uploadId" type:"string"`
}

// String returns the string representation
func (s InitiateLayerUploadOutput) String() string {
	return awsutil.Prettify(s)
}

const opInitiateLayerUpload = "InitiateLayerUpload"

// InitiateLayerUploadRequest returns a request value for making API operation for
// Amazon EC2 Container Registry.
//
// Notifies Amazon ECR that you intend to upload an image layer.
//
// When an image is pushed, the InitiateLayerUpload API is called once per image
// layer that has not already been uploaded. Whether or not an image layer has
// been uploaded is determined by the BatchCheckLayerAvailability API action.
//
// This operation is used by the Amazon ECR proxy and is not generally used
// by customers for pulling and pushing images. In most cases, you should use
// the docker CLI to pull, tag, and push images.
//
//    // Example sending a request using InitiateLayerUploadRequest.
//    req := client.InitiateLayerUploadRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/InitiateLayerUpload
func (c *Client) InitiateLayerUploadRequest(input *InitiateLayerUploadInput) InitiateLayerUploadRequest {
	op := &aws.Operation{
		Name:       opInitiateLayerUpload,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &InitiateLayerUploadInput{}
	}

	req := c.newRequest(op, input, &InitiateLayerUploadOutput{})

	return InitiateLayerUploadRequest{Request: req, Input: input, Copy: c.InitiateLayerUploadRequest}
}

// InitiateLayerUploadRequest is the request type for the
// InitiateLayerUpload API operation.
type InitiateLayerUploadRequest struct {
	*aws.Request
	Input *InitiateLayerUploadInput
	Copy  func(*InitiateLayerUploadInput) InitiateLayerUploadRequest
}

// Send marshals and sends the InitiateLayerUpload API request.
func (r InitiateLayerUploadRequest) Send(ctx context.Context) (*InitiateLayerUploadResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &InitiateLayerUploadResponse{
		InitiateLayerUploadOutput: r.Request.Data.(*InitiateLayerUploadOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// InitiateLayerUploadResponse is the response type for the
// InitiateLayerUpload API operation.
type InitiateLayerUploadResponse struct {
	*InitiateLayerUploadOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// InitiateLayerUpload request.
func (r *InitiateLayerUploadResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
