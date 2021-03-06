// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteDatabaseInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog in which the database resides. If none is provided,
	// the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The name of the database to delete. For Hive compatibility, this must be
	// all lowercase.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDatabaseInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDatabaseInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDatabaseInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDatabaseOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDatabaseOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDatabase = "DeleteDatabase"

// DeleteDatabaseRequest returns a request value for making API operation for
// AWS Glue.
//
// Removes a specified database from a Data Catalog.
//
// After completing this operation, you no longer have access to the tables
// (and all table versions and partitions that might belong to the tables) and
// the user-defined functions in the deleted database. AWS Glue deletes these
// "orphaned" resources asynchronously in a timely manner, at the discretion
// of the service.
//
// To ensure the immediate deletion of all related resources, before calling
// DeleteDatabase, use DeleteTableVersion or BatchDeleteTableVersion, DeletePartition
// or BatchDeletePartition, DeleteUserDefinedFunction, and DeleteTable or BatchDeleteTable,
// to delete any resources that belong to the database.
//
//    // Example sending a request using DeleteDatabaseRequest.
//    req := client.DeleteDatabaseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/DeleteDatabase
func (c *Client) DeleteDatabaseRequest(input *DeleteDatabaseInput) DeleteDatabaseRequest {
	op := &aws.Operation{
		Name:       opDeleteDatabase,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDatabaseInput{}
	}

	req := c.newRequest(op, input, &DeleteDatabaseOutput{})

	return DeleteDatabaseRequest{Request: req, Input: input, Copy: c.DeleteDatabaseRequest}
}

// DeleteDatabaseRequest is the request type for the
// DeleteDatabase API operation.
type DeleteDatabaseRequest struct {
	*aws.Request
	Input *DeleteDatabaseInput
	Copy  func(*DeleteDatabaseInput) DeleteDatabaseRequest
}

// Send marshals and sends the DeleteDatabase API request.
func (r DeleteDatabaseRequest) Send(ctx context.Context) (*DeleteDatabaseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDatabaseResponse{
		DeleteDatabaseOutput: r.Request.Data.(*DeleteDatabaseOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDatabaseResponse is the response type for the
// DeleteDatabase API operation.
type DeleteDatabaseResponse struct {
	*DeleteDatabaseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDatabase request.
func (r *DeleteDatabaseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
