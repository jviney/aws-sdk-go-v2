// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteRelationalDatabaseInput struct {
	_ struct{} `type:"structure"`

	// The name of the database snapshot created if skip final snapshot is false,
	// which is the default value for that parameter.
	//
	// Specifying this parameter and also specifying the skip final snapshot parameter
	// to true results in an error.
	//
	// Constraints:
	//
	//    * Must contain from 2 to 255 alphanumeric characters, or hyphens.
	//
	//    * The first and last character must be a letter or number.
	FinalRelationalDatabaseSnapshotName *string `locationName:"finalRelationalDatabaseSnapshotName" type:"string"`

	// The name of the database that you are deleting.
	//
	// RelationalDatabaseName is a required field
	RelationalDatabaseName *string `locationName:"relationalDatabaseName" type:"string" required:"true"`

	// Determines whether a final database snapshot is created before your database
	// is deleted. If true is specified, no database snapshot is created. If false
	// is specified, a database snapshot is created before your database is deleted.
	//
	// You must specify the final relational database snapshot name parameter if
	// the skip final snapshot parameter is false.
	//
	// Default: false
	SkipFinalSnapshot *bool `locationName:"skipFinalSnapshot" type:"boolean"`
}

// String returns the string representation
func (s DeleteRelationalDatabaseInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRelationalDatabaseInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRelationalDatabaseInput"}

	if s.RelationalDatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RelationalDatabaseName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteRelationalDatabaseOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected
	// by the request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s DeleteRelationalDatabaseOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteRelationalDatabase = "DeleteRelationalDatabase"

// DeleteRelationalDatabaseRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Deletes a database in Amazon Lightsail.
//
// The delete relational database operation supports tag-based access control
// via resource tags applied to the resource identified by relationalDatabaseName.
// For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using DeleteRelationalDatabaseRequest.
//    req := client.DeleteRelationalDatabaseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DeleteRelationalDatabase
func (c *Client) DeleteRelationalDatabaseRequest(input *DeleteRelationalDatabaseInput) DeleteRelationalDatabaseRequest {
	op := &aws.Operation{
		Name:       opDeleteRelationalDatabase,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRelationalDatabaseInput{}
	}

	req := c.newRequest(op, input, &DeleteRelationalDatabaseOutput{})

	return DeleteRelationalDatabaseRequest{Request: req, Input: input, Copy: c.DeleteRelationalDatabaseRequest}
}

// DeleteRelationalDatabaseRequest is the request type for the
// DeleteRelationalDatabase API operation.
type DeleteRelationalDatabaseRequest struct {
	*aws.Request
	Input *DeleteRelationalDatabaseInput
	Copy  func(*DeleteRelationalDatabaseInput) DeleteRelationalDatabaseRequest
}

// Send marshals and sends the DeleteRelationalDatabase API request.
func (r DeleteRelationalDatabaseRequest) Send(ctx context.Context) (*DeleteRelationalDatabaseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRelationalDatabaseResponse{
		DeleteRelationalDatabaseOutput: r.Request.Data.(*DeleteRelationalDatabaseOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRelationalDatabaseResponse is the response type for the
// DeleteRelationalDatabase API operation.
type DeleteRelationalDatabaseResponse struct {
	*DeleteRelationalDatabaseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRelationalDatabase request.
func (r *DeleteRelationalDatabaseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
