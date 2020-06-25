// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// The input argument to the TerminationProtection operation.
type SetTerminationProtectionInput struct {
	_ struct{} `type:"structure"`

	// A list of strings that uniquely identify the clusters to protect. This identifier
	// is returned by RunJobFlow and can also be obtained from DescribeJobFlows .
	//
	// JobFlowIds is a required field
	JobFlowIds []string `type:"list" required:"true"`

	// A Boolean that indicates whether to protect the cluster and prevent the Amazon
	// EC2 instances in the cluster from shutting down due to API calls, user intervention,
	// or job-flow error.
	//
	// TerminationProtected is a required field
	TerminationProtected *bool `type:"boolean" required:"true"`
}

// String returns the string representation
func (s SetTerminationProtectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetTerminationProtectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetTerminationProtectionInput"}

	if s.JobFlowIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobFlowIds"))
	}

	if s.TerminationProtected == nil {
		invalidParams.Add(aws.NewErrParamRequired("TerminationProtected"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SetTerminationProtectionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetTerminationProtectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetTerminationProtection = "SetTerminationProtection"

// SetTerminationProtectionRequest returns a request value for making API operation for
// Amazon Elastic MapReduce.
//
// SetTerminationProtection locks a cluster (job flow) so the EC2 instances
// in the cluster cannot be terminated by user intervention, an API call, or
// in the event of a job-flow error. The cluster still terminates upon successful
// completion of the job flow. Calling SetTerminationProtection on a cluster
// is similar to calling the Amazon EC2 DisableAPITermination API on all EC2
// instances in a cluster.
//
// SetTerminationProtection is used to prevent accidental termination of a cluster
// and to ensure that in the event of an error, the instances persist so that
// you can recover any data stored in their ephemeral instance storage.
//
// To terminate a cluster that has been locked by setting SetTerminationProtection
// to true, you must first unlock the job flow by a subsequent call to SetTerminationProtection
// in which you set the value to false.
//
// For more information, seeManaging Cluster Termination (https://docs.aws.amazon.com/emr/latest/ManagementGuide/UsingEMR_TerminationProtection.html)
// in the Amazon EMR Management Guide.
//
//    // Example sending a request using SetTerminationProtectionRequest.
//    req := client.SetTerminationProtectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/SetTerminationProtection
func (c *Client) SetTerminationProtectionRequest(input *SetTerminationProtectionInput) SetTerminationProtectionRequest {
	op := &aws.Operation{
		Name:       opSetTerminationProtection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetTerminationProtectionInput{}
	}

	req := c.newRequest(op, input, &SetTerminationProtectionOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return SetTerminationProtectionRequest{Request: req, Input: input, Copy: c.SetTerminationProtectionRequest}
}

// SetTerminationProtectionRequest is the request type for the
// SetTerminationProtection API operation.
type SetTerminationProtectionRequest struct {
	*aws.Request
	Input *SetTerminationProtectionInput
	Copy  func(*SetTerminationProtectionInput) SetTerminationProtectionRequest
}

// Send marshals and sends the SetTerminationProtection API request.
func (r SetTerminationProtectionRequest) Send(ctx context.Context) (*SetTerminationProtectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetTerminationProtectionResponse{
		SetTerminationProtectionOutput: r.Request.Data.(*SetTerminationProtectionOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetTerminationProtectionResponse is the response type for the
// SetTerminationProtection API operation.
type SetTerminationProtectionResponse struct {
	*SetTerminationProtectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetTerminationProtection request.
func (r *SetTerminationProtectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
