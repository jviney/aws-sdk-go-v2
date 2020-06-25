// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DecreaseNodeGroupsInGlobalReplicationGroupInput struct {
	_ struct{} `type:"structure"`

	// Indicates that the shard reconfiguration process begins immediately. At present,
	// the only permitted value for this parameter is true.
	//
	// ApplyImmediately is a required field
	ApplyImmediately *bool `type:"boolean" required:"true"`

	// If the value of NodeGroupCount is less than the current number of node groups
	// (shards), then either NodeGroupsToRemove or NodeGroupsToRetain is required.
	// NodeGroupsToRemove is a list of NodeGroupIds to remove from the cluster.
	// ElastiCache for Redis will attempt to remove all node groups listed by NodeGroupsToRemove
	// from the cluster.
	GlobalNodeGroupsToRemove []string `locationNameList:"GlobalNodeGroupId" type:"list"`

	// If the value of NodeGroupCount is less than the current number of node groups
	// (shards), then either NodeGroupsToRemove or NodeGroupsToRetain is required.
	// NodeGroupsToRemove is a list of NodeGroupIds to remove from the cluster.
	// ElastiCache for Redis will attempt to remove all node groups listed by NodeGroupsToRemove
	// from the cluster.
	GlobalNodeGroupsToRetain []string `locationNameList:"GlobalNodeGroupId" type:"list"`

	// The name of the Global Datastore
	//
	// GlobalReplicationGroupId is a required field
	GlobalReplicationGroupId *string `type:"string" required:"true"`

	// The number of node groups (shards) that results from the modification of
	// the shard configuration
	//
	// NodeGroupCount is a required field
	NodeGroupCount *int64 `type:"integer" required:"true"`
}

// String returns the string representation
func (s DecreaseNodeGroupsInGlobalReplicationGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DecreaseNodeGroupsInGlobalReplicationGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DecreaseNodeGroupsInGlobalReplicationGroupInput"}

	if s.ApplyImmediately == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplyImmediately"))
	}

	if s.GlobalReplicationGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GlobalReplicationGroupId"))
	}

	if s.NodeGroupCount == nil {
		invalidParams.Add(aws.NewErrParamRequired("NodeGroupCount"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DecreaseNodeGroupsInGlobalReplicationGroupOutput struct {
	_ struct{} `type:"structure"`

	// Consists of a primary cluster that accepts writes and an associated secondary
	// cluster that resides in a different AWS region. The secondary cluster accepts
	// only reads. The primary cluster automatically replicates updates to the secondary
	// cluster.
	//
	//    * The GlobalReplicationGroupIdSuffix represents the name of the Global
	//    Datastore, which is what you use to associate a secondary cluster.
	GlobalReplicationGroup *GlobalReplicationGroup `type:"structure"`
}

// String returns the string representation
func (s DecreaseNodeGroupsInGlobalReplicationGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDecreaseNodeGroupsInGlobalReplicationGroup = "DecreaseNodeGroupsInGlobalReplicationGroup"

// DecreaseNodeGroupsInGlobalReplicationGroupRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Decreases the number of node groups in a Global Datastore
//
//    // Example sending a request using DecreaseNodeGroupsInGlobalReplicationGroupRequest.
//    req := client.DecreaseNodeGroupsInGlobalReplicationGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DecreaseNodeGroupsInGlobalReplicationGroup
func (c *Client) DecreaseNodeGroupsInGlobalReplicationGroupRequest(input *DecreaseNodeGroupsInGlobalReplicationGroupInput) DecreaseNodeGroupsInGlobalReplicationGroupRequest {
	op := &aws.Operation{
		Name:       opDecreaseNodeGroupsInGlobalReplicationGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DecreaseNodeGroupsInGlobalReplicationGroupInput{}
	}

	req := c.newRequest(op, input, &DecreaseNodeGroupsInGlobalReplicationGroupOutput{})

	return DecreaseNodeGroupsInGlobalReplicationGroupRequest{Request: req, Input: input, Copy: c.DecreaseNodeGroupsInGlobalReplicationGroupRequest}
}

// DecreaseNodeGroupsInGlobalReplicationGroupRequest is the request type for the
// DecreaseNodeGroupsInGlobalReplicationGroup API operation.
type DecreaseNodeGroupsInGlobalReplicationGroupRequest struct {
	*aws.Request
	Input *DecreaseNodeGroupsInGlobalReplicationGroupInput
	Copy  func(*DecreaseNodeGroupsInGlobalReplicationGroupInput) DecreaseNodeGroupsInGlobalReplicationGroupRequest
}

// Send marshals and sends the DecreaseNodeGroupsInGlobalReplicationGroup API request.
func (r DecreaseNodeGroupsInGlobalReplicationGroupRequest) Send(ctx context.Context) (*DecreaseNodeGroupsInGlobalReplicationGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DecreaseNodeGroupsInGlobalReplicationGroupResponse{
		DecreaseNodeGroupsInGlobalReplicationGroupOutput: r.Request.Data.(*DecreaseNodeGroupsInGlobalReplicationGroupOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DecreaseNodeGroupsInGlobalReplicationGroupResponse is the response type for the
// DecreaseNodeGroupsInGlobalReplicationGroup API operation.
type DecreaseNodeGroupsInGlobalReplicationGroupResponse struct {
	*DecreaseNodeGroupsInGlobalReplicationGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DecreaseNodeGroupsInGlobalReplicationGroup request.
func (r *DecreaseNodeGroupsInGlobalReplicationGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
