// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kendra

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteDataSourceInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the data source to delete.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// The unique identifier of the index associated with the data source.
	//
	// IndexId is a required field
	IndexId *string `min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDataSourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDataSourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDataSourceInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if s.IndexId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IndexId"))
	}
	if s.IndexId != nil && len(*s.IndexId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("IndexId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDataSourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDataSourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDataSource = "DeleteDataSource"

// DeleteDataSourceRequest returns a request value for making API operation for
// AWSKendraFrontendService.
//
// Deletes an Amazon Kendra data source. An exception is not thrown if the data
// source is already being deleted. While the data source is being deleted,
// the Status field returned by a call to the operation is set to DELETING.
// For more information, see Deleting Data Sources (https://docs.aws.amazon.com/kendra/latest/dg/delete-data-source.html).
//
//    // Example sending a request using DeleteDataSourceRequest.
//    req := client.DeleteDataSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kendra-2019-02-03/DeleteDataSource
func (c *Client) DeleteDataSourceRequest(input *DeleteDataSourceInput) DeleteDataSourceRequest {
	op := &aws.Operation{
		Name:       opDeleteDataSource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDataSourceInput{}
	}

	req := c.newRequest(op, input, &DeleteDataSourceOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteDataSourceRequest{Request: req, Input: input, Copy: c.DeleteDataSourceRequest}
}

// DeleteDataSourceRequest is the request type for the
// DeleteDataSource API operation.
type DeleteDataSourceRequest struct {
	*aws.Request
	Input *DeleteDataSourceInput
	Copy  func(*DeleteDataSourceInput) DeleteDataSourceRequest
}

// Send marshals and sends the DeleteDataSource API request.
func (r DeleteDataSourceRequest) Send(ctx context.Context) (*DeleteDataSourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDataSourceResponse{
		DeleteDataSourceOutput: r.Request.Data.(*DeleteDataSourceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDataSourceResponse is the response type for the
// DeleteDataSource API operation.
type DeleteDataSourceResponse struct {
	*DeleteDataSourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDataSource request.
func (r *DeleteDataSourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
