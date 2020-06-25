// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/query"
)

type DeleteHsmConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the Amazon Redshift HSM configuration to be deleted.
	//
	// HsmConfigurationIdentifier is a required field
	HsmConfigurationIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteHsmConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteHsmConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteHsmConfigurationInput"}

	if s.HsmConfigurationIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("HsmConfigurationIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteHsmConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteHsmConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteHsmConfiguration = "DeleteHsmConfiguration"

// DeleteHsmConfigurationRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Deletes the specified Amazon Redshift HSM configuration.
//
//    // Example sending a request using DeleteHsmConfigurationRequest.
//    req := client.DeleteHsmConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DeleteHsmConfiguration
func (c *Client) DeleteHsmConfigurationRequest(input *DeleteHsmConfigurationInput) DeleteHsmConfigurationRequest {
	op := &aws.Operation{
		Name:       opDeleteHsmConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteHsmConfigurationInput{}
	}

	req := c.newRequest(op, input, &DeleteHsmConfigurationOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteHsmConfigurationRequest{Request: req, Input: input, Copy: c.DeleteHsmConfigurationRequest}
}

// DeleteHsmConfigurationRequest is the request type for the
// DeleteHsmConfiguration API operation.
type DeleteHsmConfigurationRequest struct {
	*aws.Request
	Input *DeleteHsmConfigurationInput
	Copy  func(*DeleteHsmConfigurationInput) DeleteHsmConfigurationRequest
}

// Send marshals and sends the DeleteHsmConfiguration API request.
func (r DeleteHsmConfigurationRequest) Send(ctx context.Context) (*DeleteHsmConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteHsmConfigurationResponse{
		DeleteHsmConfigurationOutput: r.Request.Data.(*DeleteHsmConfigurationOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteHsmConfigurationResponse is the response type for the
// DeleteHsmConfiguration API operation.
type DeleteHsmConfigurationResponse struct {
	*DeleteHsmConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteHsmConfiguration request.
func (r *DeleteHsmConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
