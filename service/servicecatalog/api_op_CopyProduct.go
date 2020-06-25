// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CopyProductInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The copy options. If the value is CopyTags, the tags from the source product
	// are copied to the target product.
	CopyOptions []CopyOption `type:"list"`

	// A unique identifier that you provide to ensure idempotency. If multiple requests
	// differ only by the idempotency token, the same response is returned for each
	// repeated request.
	//
	// IdempotencyToken is a required field
	IdempotencyToken *string `min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The Amazon Resource Name (ARN) of the source product.
	//
	// SourceProductArn is a required field
	SourceProductArn *string `min:"1" type:"string" required:"true"`

	// The identifiers of the provisioning artifacts (also known as versions) of
	// the product to copy. By default, all provisioning artifacts are copied.
	SourceProvisioningArtifactIdentifiers []map[string]string `type:"list"`

	// The identifier of the target product. By default, a new product is created.
	TargetProductId *string `min:"1" type:"string"`

	// A name for the target product. The default is the name of the source product.
	TargetProductName *string `type:"string"`
}

// String returns the string representation
func (s CopyProductInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CopyProductInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CopyProductInput"}

	if s.IdempotencyToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdempotencyToken"))
	}
	if s.IdempotencyToken != nil && len(*s.IdempotencyToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdempotencyToken", 1))
	}

	if s.SourceProductArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceProductArn"))
	}
	if s.SourceProductArn != nil && len(*s.SourceProductArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SourceProductArn", 1))
	}
	if s.TargetProductId != nil && len(*s.TargetProductId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetProductId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CopyProductOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to track the progress of the operation.
	CopyProductToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CopyProductOutput) String() string {
	return awsutil.Prettify(s)
}

const opCopyProduct = "CopyProduct"

// CopyProductRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Copies the specified source product to the specified target product or a
// new product.
//
// You can copy a product to the same account or another account. You can copy
// a product to the same region or another region.
//
// This operation is performed asynchronously. To track the progress of the
// operation, use DescribeCopyProductStatus.
//
//    // Example sending a request using CopyProductRequest.
//    req := client.CopyProductRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/CopyProduct
func (c *Client) CopyProductRequest(input *CopyProductInput) CopyProductRequest {
	op := &aws.Operation{
		Name:       opCopyProduct,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CopyProductInput{}
	}

	req := c.newRequest(op, input, &CopyProductOutput{})

	return CopyProductRequest{Request: req, Input: input, Copy: c.CopyProductRequest}
}

// CopyProductRequest is the request type for the
// CopyProduct API operation.
type CopyProductRequest struct {
	*aws.Request
	Input *CopyProductInput
	Copy  func(*CopyProductInput) CopyProductRequest
}

// Send marshals and sends the CopyProduct API request.
func (r CopyProductRequest) Send(ctx context.Context) (*CopyProductResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CopyProductResponse{
		CopyProductOutput: r.Request.Data.(*CopyProductOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CopyProductResponse is the response type for the
// CopyProduct API operation.
type CopyProductResponse struct {
	*CopyProductOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CopyProduct request.
func (r *CopyProductResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
