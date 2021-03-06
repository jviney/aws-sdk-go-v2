// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package docdb

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input to FailoverDBCluster.
type FailoverDBClusterInput struct {
	_ struct{} `type:"structure"`

	// A cluster identifier to force a failover for. This parameter is not case
	// sensitive.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing DBCluster.
	DBClusterIdentifier *string `type:"string"`

	// The name of the instance to promote to the primary instance.
	//
	// You must specify the instance identifier for an Amazon DocumentDB replica
	// in the cluster. For example, mydbcluster-replica1.
	TargetDBInstanceIdentifier *string `type:"string"`
}

// String returns the string representation
func (s FailoverDBClusterInput) String() string {
	return awsutil.Prettify(s)
}

type FailoverDBClusterOutput struct {
	_ struct{} `type:"structure"`

	// Detailed information about a cluster.
	DBCluster *DBCluster `type:"structure"`
}

// String returns the string representation
func (s FailoverDBClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opFailoverDBCluster = "FailoverDBCluster"

// FailoverDBClusterRequest returns a request value for making API operation for
// Amazon DocumentDB with MongoDB compatibility.
//
// Forces a failover for a cluster.
//
// A failover for a cluster promotes one of the Amazon DocumentDB replicas (read-only
// instances) in the cluster to be the primary instance (the cluster writer).
//
// If the primary instance fails, Amazon DocumentDB automatically fails over
// to an Amazon DocumentDB replica, if one exists. You can force a failover
// when you want to simulate a failure of a primary instance for testing.
//
//    // Example sending a request using FailoverDBClusterRequest.
//    req := client.FailoverDBClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/FailoverDBCluster
func (c *Client) FailoverDBClusterRequest(input *FailoverDBClusterInput) FailoverDBClusterRequest {
	op := &aws.Operation{
		Name:       opFailoverDBCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &FailoverDBClusterInput{}
	}

	req := c.newRequest(op, input, &FailoverDBClusterOutput{})

	return FailoverDBClusterRequest{Request: req, Input: input, Copy: c.FailoverDBClusterRequest}
}

// FailoverDBClusterRequest is the request type for the
// FailoverDBCluster API operation.
type FailoverDBClusterRequest struct {
	*aws.Request
	Input *FailoverDBClusterInput
	Copy  func(*FailoverDBClusterInput) FailoverDBClusterRequest
}

// Send marshals and sends the FailoverDBCluster API request.
func (r FailoverDBClusterRequest) Send(ctx context.Context) (*FailoverDBClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &FailoverDBClusterResponse{
		FailoverDBClusterOutput: r.Request.Data.(*FailoverDBClusterOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// FailoverDBClusterResponse is the response type for the
// FailoverDBCluster API operation.
type FailoverDBClusterResponse struct {
	*FailoverDBClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// FailoverDBCluster request.
func (r *FailoverDBClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
