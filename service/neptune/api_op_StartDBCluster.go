// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type StartDBClusterInput struct {
	_ struct{} `type:"structure"`

	// The DB cluster identifier of the Neptune DB cluster to be started. This parameter
	// is stored as a lowercase string.
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StartDBClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartDBClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartDBClusterInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartDBClusterOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon Neptune DB cluster.
	//
	// This data type is used as a response element in the DescribeDBClusters action.
	DBCluster *DBCluster `type:"structure"`
}

// String returns the string representation
func (s StartDBClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartDBCluster = "StartDBCluster"

// StartDBClusterRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Starts an Amazon Neptune DB cluster that was stopped using the AWS console,
// the AWS CLI stop-db-cluster command, or the StopDBCluster API.
//
//    // Example sending a request using StartDBClusterRequest.
//    req := client.StartDBClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/StartDBCluster
func (c *Client) StartDBClusterRequest(input *StartDBClusterInput) StartDBClusterRequest {
	op := &aws.Operation{
		Name:       opStartDBCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartDBClusterInput{}
	}

	req := c.newRequest(op, input, &StartDBClusterOutput{})

	return StartDBClusterRequest{Request: req, Input: input, Copy: c.StartDBClusterRequest}
}

// StartDBClusterRequest is the request type for the
// StartDBCluster API operation.
type StartDBClusterRequest struct {
	*aws.Request
	Input *StartDBClusterInput
	Copy  func(*StartDBClusterInput) StartDBClusterRequest
}

// Send marshals and sends the StartDBCluster API request.
func (r StartDBClusterRequest) Send(ctx context.Context) (*StartDBClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartDBClusterResponse{
		StartDBClusterOutput: r.Request.Data.(*StartDBClusterOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartDBClusterResponse is the response type for the
// StartDBCluster API operation.
type StartDBClusterResponse struct {
	*StartDBClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartDBCluster request.
func (r *StartDBClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
