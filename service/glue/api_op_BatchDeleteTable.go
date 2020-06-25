// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type BatchDeleteTableInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog where the table resides. If none is provided,
	// the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The name of the catalog database in which the tables to delete reside. For
	// Hive compatibility, this name is entirely lowercase.
	//
	// DatabaseName is a required field
	DatabaseName *string `min:"1" type:"string" required:"true"`

	// A list of the table to delete.
	//
	// TablesToDelete is a required field
	TablesToDelete []string `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDeleteTableInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchDeleteTableInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchDeleteTableInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatabaseName"))
	}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}

	if s.TablesToDelete == nil {
		invalidParams.Add(aws.NewErrParamRequired("TablesToDelete"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchDeleteTableOutput struct {
	_ struct{} `type:"structure"`

	// A list of errors encountered in attempting to delete the specified tables.
	Errors []TableError `type:"list"`
}

// String returns the string representation
func (s BatchDeleteTableOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchDeleteTable = "BatchDeleteTable"

// BatchDeleteTableRequest returns a request value for making API operation for
// AWS Glue.
//
// Deletes multiple tables at once.
//
// After completing this operation, you no longer have access to the table versions
// and partitions that belong to the deleted table. AWS Glue deletes these "orphaned"
// resources asynchronously in a timely manner, at the discretion of the service.
//
// To ensure the immediate deletion of all related resources, before calling
// BatchDeleteTable, use DeleteTableVersion or BatchDeleteTableVersion, and
// DeletePartition or BatchDeletePartition, to delete any resources that belong
// to the table.
//
//    // Example sending a request using BatchDeleteTableRequest.
//    req := client.BatchDeleteTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/BatchDeleteTable
func (c *Client) BatchDeleteTableRequest(input *BatchDeleteTableInput) BatchDeleteTableRequest {
	op := &aws.Operation{
		Name:       opBatchDeleteTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchDeleteTableInput{}
	}

	req := c.newRequest(op, input, &BatchDeleteTableOutput{})

	return BatchDeleteTableRequest{Request: req, Input: input, Copy: c.BatchDeleteTableRequest}
}

// BatchDeleteTableRequest is the request type for the
// BatchDeleteTable API operation.
type BatchDeleteTableRequest struct {
	*aws.Request
	Input *BatchDeleteTableInput
	Copy  func(*BatchDeleteTableInput) BatchDeleteTableRequest
}

// Send marshals and sends the BatchDeleteTable API request.
func (r BatchDeleteTableRequest) Send(ctx context.Context) (*BatchDeleteTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDeleteTableResponse{
		BatchDeleteTableOutput: r.Request.Data.(*BatchDeleteTableOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDeleteTableResponse is the response type for the
// BatchDeleteTable API operation.
type BatchDeleteTableResponse struct {
	*BatchDeleteTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDeleteTable request.
func (r *BatchDeleteTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
