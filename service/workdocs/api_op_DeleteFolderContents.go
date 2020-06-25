// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteFolderContentsInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string" sensitive:"true"`

	// The ID of the folder.
	//
	// FolderId is a required field
	FolderId *string `location:"uri" locationName:"FolderId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteFolderContentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFolderContentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFolderContentsInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}

	if s.FolderId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FolderId"))
	}
	if s.FolderId != nil && len(*s.FolderId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FolderId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteFolderContentsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AuthenticationToken != nil {
		v := *s.AuthenticationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Authentication", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FolderId != nil {
		v := *s.FolderId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FolderId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteFolderContentsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteFolderContentsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteFolderContentsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteFolderContents = "DeleteFolderContents"

// DeleteFolderContentsRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Deletes the contents of the specified folder.
//
//    // Example sending a request using DeleteFolderContentsRequest.
//    req := client.DeleteFolderContentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DeleteFolderContents
func (c *Client) DeleteFolderContentsRequest(input *DeleteFolderContentsInput) DeleteFolderContentsRequest {
	op := &aws.Operation{
		Name:       opDeleteFolderContents,
		HTTPMethod: "DELETE",
		HTTPPath:   "/api/v1/folders/{FolderId}/contents",
	}

	if input == nil {
		input = &DeleteFolderContentsInput{}
	}

	req := c.newRequest(op, input, &DeleteFolderContentsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteFolderContentsRequest{Request: req, Input: input, Copy: c.DeleteFolderContentsRequest}
}

// DeleteFolderContentsRequest is the request type for the
// DeleteFolderContents API operation.
type DeleteFolderContentsRequest struct {
	*aws.Request
	Input *DeleteFolderContentsInput
	Copy  func(*DeleteFolderContentsInput) DeleteFolderContentsRequest
}

// Send marshals and sends the DeleteFolderContents API request.
func (r DeleteFolderContentsRequest) Send(ctx context.Context) (*DeleteFolderContentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteFolderContentsResponse{
		DeleteFolderContentsOutput: r.Request.Data.(*DeleteFolderContentsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteFolderContentsResponse is the response type for the
// DeleteFolderContents API operation.
type DeleteFolderContentsResponse struct {
	*DeleteFolderContentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteFolderContents request.
func (r *DeleteFolderContentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
