// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/query"
)

type RebuildEnvironmentInput struct {
	_ struct{} `type:"structure"`

	// The ID of the environment to rebuild.
	//
	// Condition: You must specify either this or an EnvironmentName, or both. If
	// you do not specify either, AWS Elastic Beanstalk returns MissingRequiredParameter
	// error.
	EnvironmentId *string `type:"string"`

	// The name of the environment to rebuild.
	//
	// Condition: You must specify either this or an EnvironmentId, or both. If
	// you do not specify either, AWS Elastic Beanstalk returns MissingRequiredParameter
	// error.
	EnvironmentName *string `min:"4" type:"string"`
}

// String returns the string representation
func (s RebuildEnvironmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RebuildEnvironmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RebuildEnvironmentInput"}
	if s.EnvironmentName != nil && len(*s.EnvironmentName) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("EnvironmentName", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RebuildEnvironmentOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RebuildEnvironmentOutput) String() string {
	return awsutil.Prettify(s)
}

const opRebuildEnvironment = "RebuildEnvironment"

// RebuildEnvironmentRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Deletes and recreates all of the AWS resources (for example: the Auto Scaling
// group, load balancer, etc.) for a specified environment and forces a restart.
//
//    // Example sending a request using RebuildEnvironmentRequest.
//    req := client.RebuildEnvironmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/RebuildEnvironment
func (c *Client) RebuildEnvironmentRequest(input *RebuildEnvironmentInput) RebuildEnvironmentRequest {
	op := &aws.Operation{
		Name:       opRebuildEnvironment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RebuildEnvironmentInput{}
	}

	req := c.newRequest(op, input, &RebuildEnvironmentOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return RebuildEnvironmentRequest{Request: req, Input: input, Copy: c.RebuildEnvironmentRequest}
}

// RebuildEnvironmentRequest is the request type for the
// RebuildEnvironment API operation.
type RebuildEnvironmentRequest struct {
	*aws.Request
	Input *RebuildEnvironmentInput
	Copy  func(*RebuildEnvironmentInput) RebuildEnvironmentRequest
}

// Send marshals and sends the RebuildEnvironment API request.
func (r RebuildEnvironmentRequest) Send(ctx context.Context) (*RebuildEnvironmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RebuildEnvironmentResponse{
		RebuildEnvironmentOutput: r.Request.Data.(*RebuildEnvironmentOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RebuildEnvironmentResponse is the response type for the
// RebuildEnvironment API operation.
type RebuildEnvironmentResponse struct {
	*RebuildEnvironmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RebuildEnvironment request.
func (r *RebuildEnvironmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
