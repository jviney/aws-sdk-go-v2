// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dax

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type RebootNodeInput struct {
	_ struct{} `type:"structure"`

	// The name of the DAX cluster containing the node to be rebooted.
	//
	// ClusterName is a required field
	ClusterName *string `type:"string" required:"true"`

	// The system-assigned ID of the node to be rebooted.
	//
	// NodeId is a required field
	NodeId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RebootNodeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RebootNodeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RebootNodeInput"}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}

	if s.NodeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NodeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RebootNodeOutput struct {
	_ struct{} `type:"structure"`

	// A description of the DAX cluster after a node has been rebooted.
	Cluster *Cluster `type:"structure"`
}

// String returns the string representation
func (s RebootNodeOutput) String() string {
	return awsutil.Prettify(s)
}

const opRebootNode = "RebootNode"

// RebootNodeRequest returns a request value for making API operation for
// Amazon DynamoDB Accelerator (DAX).
//
// Reboots a single node of a DAX cluster. The reboot action takes place as
// soon as possible. During the reboot, the node status is set to REBOOTING.
//
// RebootNode restarts the DAX engine process and does not remove the contents
// of the cache.
//
//    // Example sending a request using RebootNodeRequest.
//    req := client.RebootNodeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/RebootNode
func (c *Client) RebootNodeRequest(input *RebootNodeInput) RebootNodeRequest {
	op := &aws.Operation{
		Name:       opRebootNode,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RebootNodeInput{}
	}

	req := c.newRequest(op, input, &RebootNodeOutput{})

	return RebootNodeRequest{Request: req, Input: input, Copy: c.RebootNodeRequest}
}

// RebootNodeRequest is the request type for the
// RebootNode API operation.
type RebootNodeRequest struct {
	*aws.Request
	Input *RebootNodeInput
	Copy  func(*RebootNodeInput) RebootNodeRequest
}

// Send marshals and sends the RebootNode API request.
func (r RebootNodeRequest) Send(ctx context.Context) (*RebootNodeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RebootNodeResponse{
		RebootNodeOutput: r.Request.Data.(*RebootNodeOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RebootNodeResponse is the response type for the
// RebootNode API operation.
type RebootNodeResponse struct {
	*RebootNodeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RebootNode request.
func (r *RebootNodeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
