// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteConnectionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog in which the connection resides. If none is provided,
	// the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The name of the connection to delete.
	//
	// ConnectionName is a required field
	ConnectionName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteConnectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteConnectionInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.ConnectionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionName"))
	}
	if s.ConnectionName != nil && len(*s.ConnectionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConnectionName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteConnectionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteConnection = "DeleteConnection"

// DeleteConnectionRequest returns a request value for making API operation for
// AWS Glue.
//
// Deletes a connection from the Data Catalog.
//
//    // Example sending a request using DeleteConnectionRequest.
//    req := client.DeleteConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/DeleteConnection
func (c *Client) DeleteConnectionRequest(input *DeleteConnectionInput) DeleteConnectionRequest {
	op := &aws.Operation{
		Name:       opDeleteConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteConnectionInput{}
	}

	req := c.newRequest(op, input, &DeleteConnectionOutput{})

	return DeleteConnectionRequest{Request: req, Input: input, Copy: c.DeleteConnectionRequest}
}

// DeleteConnectionRequest is the request type for the
// DeleteConnection API operation.
type DeleteConnectionRequest struct {
	*aws.Request
	Input *DeleteConnectionInput
	Copy  func(*DeleteConnectionInput) DeleteConnectionRequest
}

// Send marshals and sends the DeleteConnection API request.
func (r DeleteConnectionRequest) Send(ctx context.Context) (*DeleteConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteConnectionResponse{
		DeleteConnectionOutput: r.Request.Data.(*DeleteConnectionOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteConnectionResponse is the response type for the
// DeleteConnection API operation.
type DeleteConnectionResponse struct {
	*DeleteConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteConnection request.
func (r *DeleteConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
