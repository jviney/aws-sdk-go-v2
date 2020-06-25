// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ImportServerCatalogInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ImportServerCatalogInput) String() string {
	return awsutil.Prettify(s)
}

type ImportServerCatalogOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ImportServerCatalogOutput) String() string {
	return awsutil.Prettify(s)
}

const opImportServerCatalog = "ImportServerCatalog"

// ImportServerCatalogRequest returns a request value for making API operation for
// AWS Server Migration Service.
//
// Gathers a complete list of on-premises servers. Connectors must be installed
// and monitoring all servers that you want to import.
//
// This call returns immediately, but might take additional time to retrieve
// all the servers.
//
//    // Example sending a request using ImportServerCatalogRequest.
//    req := client.ImportServerCatalogRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/ImportServerCatalog
func (c *Client) ImportServerCatalogRequest(input *ImportServerCatalogInput) ImportServerCatalogRequest {
	op := &aws.Operation{
		Name:       opImportServerCatalog,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ImportServerCatalogInput{}
	}

	req := c.newRequest(op, input, &ImportServerCatalogOutput{})

	return ImportServerCatalogRequest{Request: req, Input: input, Copy: c.ImportServerCatalogRequest}
}

// ImportServerCatalogRequest is the request type for the
// ImportServerCatalog API operation.
type ImportServerCatalogRequest struct {
	*aws.Request
	Input *ImportServerCatalogInput
	Copy  func(*ImportServerCatalogInput) ImportServerCatalogRequest
}

// Send marshals and sends the ImportServerCatalog API request.
func (r ImportServerCatalogRequest) Send(ctx context.Context) (*ImportServerCatalogResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportServerCatalogResponse{
		ImportServerCatalogOutput: r.Request.Data.(*ImportServerCatalogOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportServerCatalogResponse is the response type for the
// ImportServerCatalog API operation.
type ImportServerCatalogResponse struct {
	*ImportServerCatalogOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportServerCatalog request.
func (r *ImportServerCatalogResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
