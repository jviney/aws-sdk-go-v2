// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package efs

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteFileSystemPolicyInput struct {
	_ struct{} `type:"structure"`

	// Specifies the EFS file system for which to delete the FileSystemPolicy.
	//
	// FileSystemId is a required field
	FileSystemId *string `location:"uri" locationName:"FileSystemId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteFileSystemPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFileSystemPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFileSystemPolicyInput"}

	if s.FileSystemId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FileSystemId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteFileSystemPolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FileSystemId != nil {
		v := *s.FileSystemId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FileSystemId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteFileSystemPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteFileSystemPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteFileSystemPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteFileSystemPolicy = "DeleteFileSystemPolicy"

// DeleteFileSystemPolicyRequest returns a request value for making API operation for
// Amazon Elastic File System.
//
// Deletes the FileSystemPolicy for the specified file system. The default FileSystemPolicy
// goes into effect once the existing policy is deleted. For more information
// about the default file system policy, see Using Resource-based Policies with
// EFS (https://docs.aws.amazon.com/efs/latest/ug/res-based-policies-efs.html).
//
// This operation requires permissions for the elasticfilesystem:DeleteFileSystemPolicy
// action.
//
//    // Example sending a request using DeleteFileSystemPolicyRequest.
//    req := client.DeleteFileSystemPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/DeleteFileSystemPolicy
func (c *Client) DeleteFileSystemPolicyRequest(input *DeleteFileSystemPolicyInput) DeleteFileSystemPolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteFileSystemPolicy,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2015-02-01/file-systems/{FileSystemId}/policy",
	}

	if input == nil {
		input = &DeleteFileSystemPolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteFileSystemPolicyOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteFileSystemPolicyRequest{Request: req, Input: input, Copy: c.DeleteFileSystemPolicyRequest}
}

// DeleteFileSystemPolicyRequest is the request type for the
// DeleteFileSystemPolicy API operation.
type DeleteFileSystemPolicyRequest struct {
	*aws.Request
	Input *DeleteFileSystemPolicyInput
	Copy  func(*DeleteFileSystemPolicyInput) DeleteFileSystemPolicyRequest
}

// Send marshals and sends the DeleteFileSystemPolicy API request.
func (r DeleteFileSystemPolicyRequest) Send(ctx context.Context) (*DeleteFileSystemPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteFileSystemPolicyResponse{
		DeleteFileSystemPolicyOutput: r.Request.Data.(*DeleteFileSystemPolicyOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteFileSystemPolicyResponse is the response type for the
// DeleteFileSystemPolicy API operation.
type DeleteFileSystemPolicyResponse struct {
	*DeleteFileSystemPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteFileSystemPolicy request.
func (r *DeleteFileSystemPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
