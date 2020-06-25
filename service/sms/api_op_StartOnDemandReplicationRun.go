// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type StartOnDemandReplicationRunInput struct {
	_ struct{} `type:"structure"`

	// The description of the replication run.
	Description *string `locationName:"description" type:"string"`

	// The identifier of the replication job.
	//
	// ReplicationJobId is a required field
	ReplicationJobId *string `locationName:"replicationJobId" type:"string" required:"true"`
}

// String returns the string representation
func (s StartOnDemandReplicationRunInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartOnDemandReplicationRunInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartOnDemandReplicationRunInput"}

	if s.ReplicationJobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationJobId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartOnDemandReplicationRunOutput struct {
	_ struct{} `type:"structure"`

	// The identifier of the replication run.
	ReplicationRunId *string `locationName:"replicationRunId" type:"string"`
}

// String returns the string representation
func (s StartOnDemandReplicationRunOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartOnDemandReplicationRun = "StartOnDemandReplicationRun"

// StartOnDemandReplicationRunRequest returns a request value for making API operation for
// AWS Server Migration Service.
//
// Starts an on-demand replication run for the specified replication job. This
// replication run starts immediately. This replication run is in addition to
// the ones already scheduled.
//
// There is a limit on the number of on-demand replications runs you can request
// in a 24-hour period.
//
//    // Example sending a request using StartOnDemandReplicationRunRequest.
//    req := client.StartOnDemandReplicationRunRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/StartOnDemandReplicationRun
func (c *Client) StartOnDemandReplicationRunRequest(input *StartOnDemandReplicationRunInput) StartOnDemandReplicationRunRequest {
	op := &aws.Operation{
		Name:       opStartOnDemandReplicationRun,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartOnDemandReplicationRunInput{}
	}

	req := c.newRequest(op, input, &StartOnDemandReplicationRunOutput{})

	return StartOnDemandReplicationRunRequest{Request: req, Input: input, Copy: c.StartOnDemandReplicationRunRequest}
}

// StartOnDemandReplicationRunRequest is the request type for the
// StartOnDemandReplicationRun API operation.
type StartOnDemandReplicationRunRequest struct {
	*aws.Request
	Input *StartOnDemandReplicationRunInput
	Copy  func(*StartOnDemandReplicationRunInput) StartOnDemandReplicationRunRequest
}

// Send marshals and sends the StartOnDemandReplicationRun API request.
func (r StartOnDemandReplicationRunRequest) Send(ctx context.Context) (*StartOnDemandReplicationRunResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartOnDemandReplicationRunResponse{
		StartOnDemandReplicationRunOutput: r.Request.Data.(*StartOnDemandReplicationRunOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartOnDemandReplicationRunResponse is the response type for the
// StartOnDemandReplicationRun API operation.
type StartOnDemandReplicationRunResponse struct {
	*StartOnDemandReplicationRunOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartOnDemandReplicationRun request.
func (r *StartOnDemandReplicationRunResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
