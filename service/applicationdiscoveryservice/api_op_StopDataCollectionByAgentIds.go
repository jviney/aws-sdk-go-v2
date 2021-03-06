// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type StopDataCollectionByAgentIdsInput struct {
	_ struct{} `type:"structure"`

	// The IDs of the agents or connectors from which to stop collecting data.
	//
	// AgentIds is a required field
	AgentIds []string `locationName:"agentIds" type:"list" required:"true"`
}

// String returns the string representation
func (s StopDataCollectionByAgentIdsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopDataCollectionByAgentIdsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopDataCollectionByAgentIdsInput"}

	if s.AgentIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("AgentIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopDataCollectionByAgentIdsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the agents or connector that were instructed to stop collecting
	// data. Information includes the agent/connector ID, a description of the operation
	// performed, and whether the agent/connector configuration was updated.
	AgentsConfigurationStatus []AgentConfigurationStatus `locationName:"agentsConfigurationStatus" type:"list"`
}

// String returns the string representation
func (s StopDataCollectionByAgentIdsOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopDataCollectionByAgentIds = "StopDataCollectionByAgentIds"

// StopDataCollectionByAgentIdsRequest returns a request value for making API operation for
// AWS Application Discovery Service.
//
// Instructs the specified agents or connectors to stop collecting data.
//
//    // Example sending a request using StopDataCollectionByAgentIdsRequest.
//    req := client.StopDataCollectionByAgentIdsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/StopDataCollectionByAgentIds
func (c *Client) StopDataCollectionByAgentIdsRequest(input *StopDataCollectionByAgentIdsInput) StopDataCollectionByAgentIdsRequest {
	op := &aws.Operation{
		Name:       opStopDataCollectionByAgentIds,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopDataCollectionByAgentIdsInput{}
	}

	req := c.newRequest(op, input, &StopDataCollectionByAgentIdsOutput{})

	return StopDataCollectionByAgentIdsRequest{Request: req, Input: input, Copy: c.StopDataCollectionByAgentIdsRequest}
}

// StopDataCollectionByAgentIdsRequest is the request type for the
// StopDataCollectionByAgentIds API operation.
type StopDataCollectionByAgentIdsRequest struct {
	*aws.Request
	Input *StopDataCollectionByAgentIdsInput
	Copy  func(*StopDataCollectionByAgentIdsInput) StopDataCollectionByAgentIdsRequest
}

// Send marshals and sends the StopDataCollectionByAgentIds API request.
func (r StopDataCollectionByAgentIdsRequest) Send(ctx context.Context) (*StopDataCollectionByAgentIdsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopDataCollectionByAgentIdsResponse{
		StopDataCollectionByAgentIdsOutput: r.Request.Data.(*StopDataCollectionByAgentIdsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopDataCollectionByAgentIdsResponse is the response type for the
// StopDataCollectionByAgentIds API operation.
type StopDataCollectionByAgentIdsResponse struct {
	*StopDataCollectionByAgentIdsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopDataCollectionByAgentIds request.
func (r *StopDataCollectionByAgentIdsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
