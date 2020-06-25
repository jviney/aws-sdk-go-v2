// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetTagsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the resource for which to retrieve tags.
	//
	// ResourceArn is a required field
	ResourceArn *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTagsInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}
	if s.ResourceArn != nil && len(*s.ResourceArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetTagsOutput struct {
	_ struct{} `type:"structure"`

	// The requested tags.
	Tags map[string]string `type:"map"`
}

// String returns the string representation
func (s GetTagsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetTags = "GetTags"

// GetTagsRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves a list of tags associated with a resource.
//
//    // Example sending a request using GetTagsRequest.
//    req := client.GetTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetTags
func (c *Client) GetTagsRequest(input *GetTagsInput) GetTagsRequest {
	op := &aws.Operation{
		Name:       opGetTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetTagsInput{}
	}

	req := c.newRequest(op, input, &GetTagsOutput{})

	return GetTagsRequest{Request: req, Input: input, Copy: c.GetTagsRequest}
}

// GetTagsRequest is the request type for the
// GetTags API operation.
type GetTagsRequest struct {
	*aws.Request
	Input *GetTagsInput
	Copy  func(*GetTagsInput) GetTagsRequest
}

// Send marshals and sends the GetTags API request.
func (r GetTagsRequest) Send(ctx context.Context) (*GetTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTagsResponse{
		GetTagsOutput: r.Request.Data.(*GetTagsOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetTagsResponse is the response type for the
// GetTags API operation.
type GetTagsResponse struct {
	*GetTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTags request.
func (r *GetTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
