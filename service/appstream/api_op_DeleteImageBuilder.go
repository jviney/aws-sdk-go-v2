// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteImageBuilderInput struct {
	_ struct{} `type:"structure"`

	// The name of the image builder.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteImageBuilderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteImageBuilderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteImageBuilderInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteImageBuilderOutput struct {
	_ struct{} `type:"structure"`

	// Information about the image builder.
	ImageBuilder *ImageBuilder `type:"structure"`
}

// String returns the string representation
func (s DeleteImageBuilderOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteImageBuilder = "DeleteImageBuilder"

// DeleteImageBuilderRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Deletes the specified image builder and releases the capacity.
//
//    // Example sending a request using DeleteImageBuilderRequest.
//    req := client.DeleteImageBuilderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DeleteImageBuilder
func (c *Client) DeleteImageBuilderRequest(input *DeleteImageBuilderInput) DeleteImageBuilderRequest {
	op := &aws.Operation{
		Name:       opDeleteImageBuilder,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteImageBuilderInput{}
	}

	req := c.newRequest(op, input, &DeleteImageBuilderOutput{})

	return DeleteImageBuilderRequest{Request: req, Input: input, Copy: c.DeleteImageBuilderRequest}
}

// DeleteImageBuilderRequest is the request type for the
// DeleteImageBuilder API operation.
type DeleteImageBuilderRequest struct {
	*aws.Request
	Input *DeleteImageBuilderInput
	Copy  func(*DeleteImageBuilderInput) DeleteImageBuilderRequest
}

// Send marshals and sends the DeleteImageBuilder API request.
func (r DeleteImageBuilderRequest) Send(ctx context.Context) (*DeleteImageBuilderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteImageBuilderResponse{
		DeleteImageBuilderOutput: r.Request.Data.(*DeleteImageBuilderOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteImageBuilderResponse is the response type for the
// DeleteImageBuilder API operation.
type DeleteImageBuilderResponse struct {
	*DeleteImageBuilderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteImageBuilder request.
func (r *DeleteImageBuilderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
