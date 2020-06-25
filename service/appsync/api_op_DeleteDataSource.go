// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DeleteDataSourceInput struct {
	_ struct{} `type:"structure"`

	// The API ID.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// The name of the data source.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDataSourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDataSourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDataSourceInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
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

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDataSourceInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDataSourceOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteDataSource = "DeleteDataSource"

// DeleteDataSourceRequest returns a request value for making API operation for
// AWS AppSync.
//
// Deletes a DataSource object.
//
//    // Example sending a request using DeleteDataSourceRequest.
//    req := client.DeleteDataSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/DeleteDataSource
func (c *Client) DeleteDataSourceRequest(input *DeleteDataSourceInput) DeleteDataSourceRequest {
	op := &aws.Operation{
		Name:       opDeleteDataSource,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/apis/{apiId}/datasources/{name}",
	}

	if input == nil {
		input = &DeleteDataSourceInput{}
	}

	req := c.newRequest(op, input, &DeleteDataSourceOutput{})

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
