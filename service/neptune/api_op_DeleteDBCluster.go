// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteDBClusterInput struct {
	_ struct{} `type:"structure"`

	// The DB cluster identifier for the DB cluster to be deleted. This parameter
	// isn't case-sensitive.
	//
	// Constraints:
	//
	//    * Must match an existing DBClusterIdentifier.
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// The DB cluster snapshot identifier of the new DB cluster snapshot created
	// when SkipFinalSnapshot is set to false.
	//
	// Specifying this parameter and also setting the SkipFinalShapshot parameter
	// to true results in an error.
	//
	// Constraints:
	//
	//    * Must be 1 to 255 letters, numbers, or hyphens.
	//
	//    * First character must be a letter
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens
	FinalDBSnapshotIdentifier *string `type:"string"`

	// Determines whether a final DB cluster snapshot is created before the DB cluster
	// is deleted. If true is specified, no DB cluster snapshot is created. If false
	// is specified, a DB cluster snapshot is created before the DB cluster is deleted.
	//
	// You must specify a FinalDBSnapshotIdentifier parameter if SkipFinalSnapshot
	// is false.
	//
	// Default: false
	SkipFinalSnapshot *bool `type:"boolean"`
}

// String returns the string representation
func (s DeleteDBClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDBClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDBClusterInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDBClusterOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon Neptune DB cluster.
	//
	// This data type is used as a response element in the DescribeDBClusters action.
	DBCluster *DBCluster `type:"structure"`
}

// String returns the string representation
func (s DeleteDBClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDBCluster = "DeleteDBCluster"

// DeleteDBClusterRequest returns a request value for making API operation for
// Amazon Neptune.
//
// The DeleteDBCluster action deletes a previously provisioned DB cluster. When
// you delete a DB cluster, all automated backups for that DB cluster are deleted
// and can't be recovered. Manual DB cluster snapshots of the specified DB cluster
// are not deleted.
//
// Note that the DB Cluster cannot be deleted if deletion protection is enabled.
// To delete it, you must first set its DeletionProtection field to False.
//
//    // Example sending a request using DeleteDBClusterRequest.
//    req := client.DeleteDBClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/DeleteDBCluster
func (c *Client) DeleteDBClusterRequest(input *DeleteDBClusterInput) DeleteDBClusterRequest {
	op := &aws.Operation{
		Name:       opDeleteDBCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDBClusterInput{}
	}

	req := c.newRequest(op, input, &DeleteDBClusterOutput{})

	return DeleteDBClusterRequest{Request: req, Input: input, Copy: c.DeleteDBClusterRequest}
}

// DeleteDBClusterRequest is the request type for the
// DeleteDBCluster API operation.
type DeleteDBClusterRequest struct {
	*aws.Request
	Input *DeleteDBClusterInput
	Copy  func(*DeleteDBClusterInput) DeleteDBClusterRequest
}

// Send marshals and sends the DeleteDBCluster API request.
func (r DeleteDBClusterRequest) Send(ctx context.Context) (*DeleteDBClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDBClusterResponse{
		DeleteDBClusterOutput: r.Request.Data.(*DeleteDBClusterOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDBClusterResponse is the response type for the
// DeleteDBCluster API operation.
type DeleteDBClusterResponse struct {
	*DeleteDBClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDBCluster request.
func (r *DeleteDBClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
