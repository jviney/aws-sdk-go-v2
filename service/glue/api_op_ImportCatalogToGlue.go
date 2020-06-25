// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ImportCatalogToGlueInput struct {
	_ struct{} `type:"structure"`

	// The ID of the catalog to import. Currently, this should be the AWS account
	// ID.
	CatalogId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ImportCatalogToGlueInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ImportCatalogToGlueInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ImportCatalogToGlueInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ImportCatalogToGlueOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ImportCatalogToGlueOutput) String() string {
	return awsutil.Prettify(s)
}

const opImportCatalogToGlue = "ImportCatalogToGlue"

// ImportCatalogToGlueRequest returns a request value for making API operation for
// AWS Glue.
//
// Imports an existing Amazon Athena Data Catalog to AWS Glue
//
//    // Example sending a request using ImportCatalogToGlueRequest.
//    req := client.ImportCatalogToGlueRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/ImportCatalogToGlue
func (c *Client) ImportCatalogToGlueRequest(input *ImportCatalogToGlueInput) ImportCatalogToGlueRequest {
	op := &aws.Operation{
		Name:       opImportCatalogToGlue,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ImportCatalogToGlueInput{}
	}

	req := c.newRequest(op, input, &ImportCatalogToGlueOutput{})

	return ImportCatalogToGlueRequest{Request: req, Input: input, Copy: c.ImportCatalogToGlueRequest}
}

// ImportCatalogToGlueRequest is the request type for the
// ImportCatalogToGlue API operation.
type ImportCatalogToGlueRequest struct {
	*aws.Request
	Input *ImportCatalogToGlueInput
	Copy  func(*ImportCatalogToGlueInput) ImportCatalogToGlueRequest
}

// Send marshals and sends the ImportCatalogToGlue API request.
func (r ImportCatalogToGlueRequest) Send(ctx context.Context) (*ImportCatalogToGlueResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportCatalogToGlueResponse{
		ImportCatalogToGlueOutput: r.Request.Data.(*ImportCatalogToGlueOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportCatalogToGlueResponse is the response type for the
// ImportCatalogToGlue API operation.
type ImportCatalogToGlueResponse struct {
	*ImportCatalogToGlueOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportCatalogToGlue request.
func (r *ImportCatalogToGlueResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
