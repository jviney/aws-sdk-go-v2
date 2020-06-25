// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreatePartitionInput struct {
	_ struct{} `type:"structure"`

	// The AWS account ID of the catalog in which the partition is to be created.
	CatalogId *string `min:"1" type:"string"`

	// The name of the metadata database in which the partition is to be created.
	//
	// DatabaseName is a required field
	DatabaseName *string `min:"1" type:"string" required:"true"`

	// A PartitionInput structure defining the partition to be created.
	//
	// PartitionInput is a required field
	PartitionInput *PartitionInput `type:"structure" required:"true"`

	// The name of the metadata table in which the partition is to be created.
	//
	// TableName is a required field
	TableName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreatePartitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePartitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePartitionInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatabaseName"))
	}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}

	if s.PartitionInput == nil {
		invalidParams.Add(aws.NewErrParamRequired("PartitionInput"))
	}

	if s.TableName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableName"))
	}
	if s.TableName != nil && len(*s.TableName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TableName", 1))
	}
	if s.PartitionInput != nil {
		if err := s.PartitionInput.Validate(); err != nil {
			invalidParams.AddNested("PartitionInput", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreatePartitionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreatePartitionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreatePartition = "CreatePartition"

// CreatePartitionRequest returns a request value for making API operation for
// AWS Glue.
//
// Creates a new partition.
//
//    // Example sending a request using CreatePartitionRequest.
//    req := client.CreatePartitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/CreatePartition
func (c *Client) CreatePartitionRequest(input *CreatePartitionInput) CreatePartitionRequest {
	op := &aws.Operation{
		Name:       opCreatePartition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreatePartitionInput{}
	}

	req := c.newRequest(op, input, &CreatePartitionOutput{})

	return CreatePartitionRequest{Request: req, Input: input, Copy: c.CreatePartitionRequest}
}

// CreatePartitionRequest is the request type for the
// CreatePartition API operation.
type CreatePartitionRequest struct {
	*aws.Request
	Input *CreatePartitionInput
	Copy  func(*CreatePartitionInput) CreatePartitionRequest
}

// Send marshals and sends the CreatePartition API request.
func (r CreatePartitionRequest) Send(ctx context.Context) (*CreatePartitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePartitionResponse{
		CreatePartitionOutput: r.Request.Data.(*CreatePartitionOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePartitionResponse is the response type for the
// CreatePartition API operation.
type CreatePartitionResponse struct {
	*CreatePartitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePartition request.
func (r *CreatePartitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
