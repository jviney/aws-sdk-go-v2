// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteServiceActionInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The self-service action identifier. For example, act-fs7abcd89wxyz.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteServiceActionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteServiceActionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteServiceActionInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteServiceActionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteServiceActionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteServiceAction = "DeleteServiceAction"

// DeleteServiceActionRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Deletes a self-service action.
//
//    // Example sending a request using DeleteServiceActionRequest.
//    req := client.DeleteServiceActionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DeleteServiceAction
func (c *Client) DeleteServiceActionRequest(input *DeleteServiceActionInput) DeleteServiceActionRequest {
	op := &aws.Operation{
		Name:       opDeleteServiceAction,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteServiceActionInput{}
	}

	req := c.newRequest(op, input, &DeleteServiceActionOutput{})

	return DeleteServiceActionRequest{Request: req, Input: input, Copy: c.DeleteServiceActionRequest}
}

// DeleteServiceActionRequest is the request type for the
// DeleteServiceAction API operation.
type DeleteServiceActionRequest struct {
	*aws.Request
	Input *DeleteServiceActionInput
	Copy  func(*DeleteServiceActionInput) DeleteServiceActionRequest
}

// Send marshals and sends the DeleteServiceAction API request.
func (r DeleteServiceActionRequest) Send(ctx context.Context) (*DeleteServiceActionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteServiceActionResponse{
		DeleteServiceActionOutput: r.Request.Data.(*DeleteServiceActionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteServiceActionResponse is the response type for the
// DeleteServiceAction API operation.
type DeleteServiceActionResponse struct {
	*DeleteServiceActionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteServiceAction request.
func (r *DeleteServiceActionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
