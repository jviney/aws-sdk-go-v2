// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetPartitionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog where the partition in question resides. If none
	// is provided, the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The name of the catalog database where the partition resides.
	//
	// DatabaseName is a required field
	DatabaseName *string `min:"1" type:"string" required:"true"`

	// The values that define the partition.
	//
	// PartitionValues is a required field
	PartitionValues []string `type:"list" required:"true"`

	// The name of the partition's table.
	//
	// TableName is a required field
	TableName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPartitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPartitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPartitionInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatabaseName"))
	}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}

	if s.PartitionValues == nil {
		invalidParams.Add(aws.NewErrParamRequired("PartitionValues"))
	}

	if s.TableName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableName"))
	}
	if s.TableName != nil && len(*s.TableName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TableName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetPartitionOutput struct {
	_ struct{} `type:"structure"`

	// The requested information, in the form of a Partition object.
	Partition *Partition `type:"structure"`
}

// String returns the string representation
func (s GetPartitionOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetPartition = "GetPartition"

// GetPartitionRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves information about a specified partition.
//
//    // Example sending a request using GetPartitionRequest.
//    req := client.GetPartitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetPartition
func (c *Client) GetPartitionRequest(input *GetPartitionInput) GetPartitionRequest {
	op := &aws.Operation{
		Name:       opGetPartition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetPartitionInput{}
	}

	req := c.newRequest(op, input, &GetPartitionOutput{})

	return GetPartitionRequest{Request: req, Input: input, Copy: c.GetPartitionRequest}
}

// GetPartitionRequest is the request type for the
// GetPartition API operation.
type GetPartitionRequest struct {
	*aws.Request
	Input *GetPartitionInput
	Copy  func(*GetPartitionInput) GetPartitionRequest
}

// Send marshals and sends the GetPartition API request.
func (r GetPartitionRequest) Send(ctx context.Context) (*GetPartitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPartitionResponse{
		GetPartitionOutput: r.Request.Data.(*GetPartitionOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPartitionResponse is the response type for the
// GetPartition API operation.
type GetPartitionResponse struct {
	*GetPartitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPartition request.
func (r *GetPartitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
