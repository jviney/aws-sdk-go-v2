// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteWebhookInput struct {
	_ struct{} `type:"structure"`

	// The name of the webhook you want to delete.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteWebhookInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteWebhookInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteWebhookInput"}

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

type DeleteWebhookOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteWebhookOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteWebhook = "DeleteWebhook"

// DeleteWebhookRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Deletes a previously created webhook by name. Deleting the webhook stops
// AWS CodePipeline from starting a pipeline every time an external event occurs.
// The API returns successfully when trying to delete a webhook that is already
// deleted. If a deleted webhook is re-created by calling PutWebhook with the
// same name, it will have a different URL.
//
//    // Example sending a request using DeleteWebhookRequest.
//    req := client.DeleteWebhookRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/DeleteWebhook
func (c *Client) DeleteWebhookRequest(input *DeleteWebhookInput) DeleteWebhookRequest {
	op := &aws.Operation{
		Name:       opDeleteWebhook,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteWebhookInput{}
	}

	req := c.newRequest(op, input, &DeleteWebhookOutput{})

	return DeleteWebhookRequest{Request: req, Input: input, Copy: c.DeleteWebhookRequest}
}

// DeleteWebhookRequest is the request type for the
// DeleteWebhook API operation.
type DeleteWebhookRequest struct {
	*aws.Request
	Input *DeleteWebhookInput
	Copy  func(*DeleteWebhookInput) DeleteWebhookRequest
}

// Send marshals and sends the DeleteWebhook API request.
func (r DeleteWebhookRequest) Send(ctx context.Context) (*DeleteWebhookResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteWebhookResponse{
		DeleteWebhookOutput: r.Request.Data.(*DeleteWebhookOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteWebhookResponse is the response type for the
// DeleteWebhook API operation.
type DeleteWebhookResponse struct {
	*DeleteWebhookOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteWebhook request.
func (r *DeleteWebhookResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
