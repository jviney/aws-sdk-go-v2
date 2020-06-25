// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/query"
)

type AbortEnvironmentUpdateInput struct {
	_ struct{} `type:"structure"`

	// This specifies the ID of the environment with the in-progress update that
	// you want to cancel.
	EnvironmentId *string `type:"string"`

	// This specifies the name of the environment with the in-progress update that
	// you want to cancel.
	EnvironmentName *string `min:"4" type:"string"`
}

// String returns the string representation
func (s AbortEnvironmentUpdateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AbortEnvironmentUpdateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AbortEnvironmentUpdateInput"}
	if s.EnvironmentName != nil && len(*s.EnvironmentName) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("EnvironmentName", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AbortEnvironmentUpdateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AbortEnvironmentUpdateOutput) String() string {
	return awsutil.Prettify(s)
}

const opAbortEnvironmentUpdate = "AbortEnvironmentUpdate"

// AbortEnvironmentUpdateRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Cancels in-progress environment configuration update or application version
// deployment.
//
//    // Example sending a request using AbortEnvironmentUpdateRequest.
//    req := client.AbortEnvironmentUpdateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/AbortEnvironmentUpdate
func (c *Client) AbortEnvironmentUpdateRequest(input *AbortEnvironmentUpdateInput) AbortEnvironmentUpdateRequest {
	op := &aws.Operation{
		Name:       opAbortEnvironmentUpdate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AbortEnvironmentUpdateInput{}
	}

	req := c.newRequest(op, input, &AbortEnvironmentUpdateOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return AbortEnvironmentUpdateRequest{Request: req, Input: input, Copy: c.AbortEnvironmentUpdateRequest}
}

// AbortEnvironmentUpdateRequest is the request type for the
// AbortEnvironmentUpdate API operation.
type AbortEnvironmentUpdateRequest struct {
	*aws.Request
	Input *AbortEnvironmentUpdateInput
	Copy  func(*AbortEnvironmentUpdateInput) AbortEnvironmentUpdateRequest
}

// Send marshals and sends the AbortEnvironmentUpdate API request.
func (r AbortEnvironmentUpdateRequest) Send(ctx context.Context) (*AbortEnvironmentUpdateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AbortEnvironmentUpdateResponse{
		AbortEnvironmentUpdateOutput: r.Request.Data.(*AbortEnvironmentUpdateOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AbortEnvironmentUpdateResponse is the response type for the
// AbortEnvironmentUpdate API operation.
type AbortEnvironmentUpdateResponse struct {
	*AbortEnvironmentUpdateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AbortEnvironmentUpdate request.
func (r *AbortEnvironmentUpdateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
