// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// DeleteAgentRequest
type DeleteAgentInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the agent to delete. Use the ListAgents
	// operation to return a list of agents for your account and AWS Region.
	//
	// AgentArn is a required field
	AgentArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAgentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAgentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAgentInput"}

	if s.AgentArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AgentArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteAgentOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAgentOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteAgent = "DeleteAgent"

// DeleteAgentRequest returns a request value for making API operation for
// AWS DataSync.
//
// Deletes an agent. To specify which agent to delete, use the Amazon Resource
// Name (ARN) of the agent in your request. The operation disassociates the
// agent from your AWS account. However, it doesn't delete the agent virtual
// machine (VM) from your on-premises environment.
//
//    // Example sending a request using DeleteAgentRequest.
//    req := client.DeleteAgentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/DeleteAgent
func (c *Client) DeleteAgentRequest(input *DeleteAgentInput) DeleteAgentRequest {
	op := &aws.Operation{
		Name:       opDeleteAgent,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteAgentInput{}
	}

	req := c.newRequest(op, input, &DeleteAgentOutput{})

	return DeleteAgentRequest{Request: req, Input: input, Copy: c.DeleteAgentRequest}
}

// DeleteAgentRequest is the request type for the
// DeleteAgent API operation.
type DeleteAgentRequest struct {
	*aws.Request
	Input *DeleteAgentInput
	Copy  func(*DeleteAgentInput) DeleteAgentRequest
}

// Send marshals and sends the DeleteAgent API request.
func (r DeleteAgentRequest) Send(ctx context.Context) (*DeleteAgentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAgentResponse{
		DeleteAgentOutput: r.Request.Data.(*DeleteAgentOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAgentResponse is the response type for the
// DeleteAgent API operation.
type DeleteAgentResponse struct {
	*DeleteAgentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAgent request.
func (r *DeleteAgentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
