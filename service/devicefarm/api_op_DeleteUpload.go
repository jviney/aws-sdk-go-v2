// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to the delete upload operation.
type DeleteUploadInput struct {
	_ struct{} `type:"structure"`

	// Represents the Amazon Resource Name (ARN) of the Device Farm upload to delete.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteUploadInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteUploadInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteUploadInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the result of a delete upload request.
type DeleteUploadOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteUploadOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteUpload = "DeleteUpload"

// DeleteUploadRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Deletes an upload given the upload ARN.
//
//    // Example sending a request using DeleteUploadRequest.
//    req := client.DeleteUploadRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/DeleteUpload
func (c *Client) DeleteUploadRequest(input *DeleteUploadInput) DeleteUploadRequest {
	op := &aws.Operation{
		Name:       opDeleteUpload,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteUploadInput{}
	}

	req := c.newRequest(op, input, &DeleteUploadOutput{})

	return DeleteUploadRequest{Request: req, Input: input, Copy: c.DeleteUploadRequest}
}

// DeleteUploadRequest is the request type for the
// DeleteUpload API operation.
type DeleteUploadRequest struct {
	*aws.Request
	Input *DeleteUploadInput
	Copy  func(*DeleteUploadInput) DeleteUploadRequest
}

// Send marshals and sends the DeleteUpload API request.
func (r DeleteUploadRequest) Send(ctx context.Context) (*DeleteUploadResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteUploadResponse{
		DeleteUploadOutput: r.Request.Data.(*DeleteUploadOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteUploadResponse is the response type for the
// DeleteUpload API operation.
type DeleteUploadResponse struct {
	*DeleteUploadOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteUpload request.
func (r *DeleteUploadResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
