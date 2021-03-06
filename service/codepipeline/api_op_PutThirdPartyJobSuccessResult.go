// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Represents the input of a PutThirdPartyJobSuccessResult action.
type PutThirdPartyJobSuccessResultInput struct {
	_ struct{} `type:"structure"`

	// The clientToken portion of the clientId and clientToken pair used to verify
	// that the calling entity is allowed access to the job and its details.
	//
	// ClientToken is a required field
	ClientToken *string `locationName:"clientToken" min:"1" type:"string" required:"true"`

	// A token generated by a job worker, such as an AWS CodeDeploy deployment ID,
	// that a successful job provides to identify a partner action in progress.
	// Future jobs use this token to identify the running instance of the action.
	// It can be reused to return more information about the progress of the partner
	// action. When the action is complete, no continuation token should be supplied.
	ContinuationToken *string `locationName:"continuationToken" min:"1" type:"string"`

	// Represents information about a current revision.
	CurrentRevision *CurrentRevision `locationName:"currentRevision" type:"structure"`

	// The details of the actions taken and results produced on an artifact as it
	// passes through stages in the pipeline.
	ExecutionDetails *ExecutionDetails `locationName:"executionDetails" type:"structure"`

	// The ID of the job that successfully completed. This is the same ID returned
	// from PollForThirdPartyJobs.
	//
	// JobId is a required field
	JobId *string `locationName:"jobId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutThirdPartyJobSuccessResultInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutThirdPartyJobSuccessResultInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutThirdPartyJobSuccessResultInput"}

	if s.ClientToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientToken"))
	}
	if s.ClientToken != nil && len(*s.ClientToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 1))
	}
	if s.ContinuationToken != nil && len(*s.ContinuationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContinuationToken", 1))
	}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}
	if s.CurrentRevision != nil {
		if err := s.CurrentRevision.Validate(); err != nil {
			invalidParams.AddNested("CurrentRevision", err.(aws.ErrInvalidParams))
		}
	}
	if s.ExecutionDetails != nil {
		if err := s.ExecutionDetails.Validate(); err != nil {
			invalidParams.AddNested("ExecutionDetails", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutThirdPartyJobSuccessResultOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutThirdPartyJobSuccessResultOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutThirdPartyJobSuccessResult = "PutThirdPartyJobSuccessResult"

// PutThirdPartyJobSuccessResultRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Represents the success of a third party job as returned to the pipeline by
// a job worker. Used for partner actions only.
//
//    // Example sending a request using PutThirdPartyJobSuccessResultRequest.
//    req := client.PutThirdPartyJobSuccessResultRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/PutThirdPartyJobSuccessResult
func (c *Client) PutThirdPartyJobSuccessResultRequest(input *PutThirdPartyJobSuccessResultInput) PutThirdPartyJobSuccessResultRequest {
	op := &aws.Operation{
		Name:       opPutThirdPartyJobSuccessResult,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutThirdPartyJobSuccessResultInput{}
	}

	req := c.newRequest(op, input, &PutThirdPartyJobSuccessResultOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return PutThirdPartyJobSuccessResultRequest{Request: req, Input: input, Copy: c.PutThirdPartyJobSuccessResultRequest}
}

// PutThirdPartyJobSuccessResultRequest is the request type for the
// PutThirdPartyJobSuccessResult API operation.
type PutThirdPartyJobSuccessResultRequest struct {
	*aws.Request
	Input *PutThirdPartyJobSuccessResultInput
	Copy  func(*PutThirdPartyJobSuccessResultInput) PutThirdPartyJobSuccessResultRequest
}

// Send marshals and sends the PutThirdPartyJobSuccessResult API request.
func (r PutThirdPartyJobSuccessResultRequest) Send(ctx context.Context) (*PutThirdPartyJobSuccessResultResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutThirdPartyJobSuccessResultResponse{
		PutThirdPartyJobSuccessResultOutput: r.Request.Data.(*PutThirdPartyJobSuccessResultOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutThirdPartyJobSuccessResultResponse is the response type for the
// PutThirdPartyJobSuccessResult API operation.
type PutThirdPartyJobSuccessResultResponse struct {
	*PutThirdPartyJobSuccessResultOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutThirdPartyJobSuccessResult request.
func (r *PutThirdPartyJobSuccessResultResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
