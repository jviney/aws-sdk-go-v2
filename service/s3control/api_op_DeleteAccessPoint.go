// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3control

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/restxml"
)

type DeleteAccessPointInput struct {
	_ struct{} `type:"structure"`

	// The account ID for the account that owns the specified access point.
	//
	// AccountId is a required field
	AccountId *string `location:"header" locationName:"x-amz-account-id" type:"string" required:"true"`

	// The name of the access point you want to delete.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAccessPointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAccessPointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAccessPointInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAccessPointInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-account-id", protocol.StringValue(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.StringValue(v), metadata)
	}
	return nil
}

type DeleteAccessPointOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAccessPointOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAccessPointOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteAccessPoint = "DeleteAccessPoint"

// DeleteAccessPointRequest returns a request value for making API operation for
// AWS S3 Control.
//
// Deletes the specified access point.
//
//    // Example sending a request using DeleteAccessPointRequest.
//    req := client.DeleteAccessPointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3control-2018-08-20/DeleteAccessPoint
func (c *Client) DeleteAccessPointRequest(input *DeleteAccessPointInput) DeleteAccessPointRequest {
	op := &aws.Operation{
		Name:       opDeleteAccessPoint,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v20180820/accesspoint/{name}",
	}

	if input == nil {
		input = &DeleteAccessPointInput{}
	}

	req := c.newRequest(op, input, &DeleteAccessPointOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	req.Handlers.Build.PushBackNamed(buildPrefixHostHandler("AccountID", aws.StringValue(input.AccountId)))
	req.Handlers.Build.PushBackNamed(buildRemoveHeaderHandler("X-Amz-Account-Id"))

	return DeleteAccessPointRequest{Request: req, Input: input, Copy: c.DeleteAccessPointRequest}
}

// DeleteAccessPointRequest is the request type for the
// DeleteAccessPoint API operation.
type DeleteAccessPointRequest struct {
	*aws.Request
	Input *DeleteAccessPointInput
	Copy  func(*DeleteAccessPointInput) DeleteAccessPointRequest
}

// Send marshals and sends the DeleteAccessPoint API request.
func (r DeleteAccessPointRequest) Send(ctx context.Context) (*DeleteAccessPointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAccessPointResponse{
		DeleteAccessPointOutput: r.Request.Data.(*DeleteAccessPointOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAccessPointResponse is the response type for the
// DeleteAccessPoint API operation.
type DeleteAccessPointResponse struct {
	*DeleteAccessPointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAccessPoint request.
func (r *DeleteAccessPointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
