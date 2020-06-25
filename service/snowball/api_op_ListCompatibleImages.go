// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package snowball

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListCompatibleImagesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results for the list of compatible images. Currently,
	// a Snowball Edge device can store 10 AMIs.
	MaxResults *int64 `type:"integer"`

	// HTTP requests are stateless. To identify what object comes "next" in the
	// list of compatible images, you can specify a value for NextToken as the starting
	// point for your list of returned images.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListCompatibleImagesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListCompatibleImagesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListCompatibleImagesInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListCompatibleImagesOutput struct {
	_ struct{} `type:"structure"`

	// A JSON-formatted object that describes a compatible AMI, including the ID
	// and name for a Snowball Edge AMI.
	CompatibleImages []CompatibleImage `type:"list"`

	// Because HTTP requests are stateless, this is the starting point for your
	// next list of returned images.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListCompatibleImagesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListCompatibleImages = "ListCompatibleImages"

// ListCompatibleImagesRequest returns a request value for making API operation for
// Amazon Import/Export Snowball.
//
// This action returns a list of the different Amazon EC2 Amazon Machine Images
// (AMIs) that are owned by your AWS account that would be supported for use
// on a Snowball Edge device. Currently, supported AMIs are based on the CentOS
// 7 (x86_64) - with Updates HVM, Ubuntu Server 14.04 LTS (HVM), and Ubuntu
// 16.04 LTS - Xenial (HVM) images, available on the AWS Marketplace.
//
//    // Example sending a request using ListCompatibleImagesRequest.
//    req := client.ListCompatibleImagesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/ListCompatibleImages
func (c *Client) ListCompatibleImagesRequest(input *ListCompatibleImagesInput) ListCompatibleImagesRequest {
	op := &aws.Operation{
		Name:       opListCompatibleImages,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListCompatibleImagesInput{}
	}

	req := c.newRequest(op, input, &ListCompatibleImagesOutput{})

	return ListCompatibleImagesRequest{Request: req, Input: input, Copy: c.ListCompatibleImagesRequest}
}

// ListCompatibleImagesRequest is the request type for the
// ListCompatibleImages API operation.
type ListCompatibleImagesRequest struct {
	*aws.Request
	Input *ListCompatibleImagesInput
	Copy  func(*ListCompatibleImagesInput) ListCompatibleImagesRequest
}

// Send marshals and sends the ListCompatibleImages API request.
func (r ListCompatibleImagesRequest) Send(ctx context.Context) (*ListCompatibleImagesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListCompatibleImagesResponse{
		ListCompatibleImagesOutput: r.Request.Data.(*ListCompatibleImagesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListCompatibleImagesResponse is the response type for the
// ListCompatibleImages API operation.
type ListCompatibleImagesResponse struct {
	*ListCompatibleImagesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListCompatibleImages request.
func (r *ListCompatibleImagesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
