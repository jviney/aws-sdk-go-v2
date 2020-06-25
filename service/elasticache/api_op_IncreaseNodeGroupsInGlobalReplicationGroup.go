// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type IncreaseNodeGroupsInGlobalReplicationGroupInput struct {
	_ struct{} `type:"structure"`

	// Indicates that the process begins immediately. At present, the only permitted
	// value for this parameter is true.
	//
	// ApplyImmediately is a required field
	ApplyImmediately *bool `type:"boolean" required:"true"`

	// The name of the Global Datastore
	//
	// GlobalReplicationGroupId is a required field
	GlobalReplicationGroupId *string `type:"string" required:"true"`

	// The number of node groups you wish to add
	//
	// NodeGroupCount is a required field
	NodeGroupCount *int64 `type:"integer" required:"true"`

	// Describes the replication group IDs, the AWS regions where they are stored
	// and the shard configuration for each that comprise the Global Datastore
	RegionalConfigurations []RegionalConfiguration `locationNameList:"RegionalConfiguration" type:"list"`
}

// String returns the string representation
func (s IncreaseNodeGroupsInGlobalReplicationGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *IncreaseNodeGroupsInGlobalReplicationGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "IncreaseNodeGroupsInGlobalReplicationGroupInput"}

	if s.ApplyImmediately == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplyImmediately"))
	}

	if s.GlobalReplicationGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GlobalReplicationGroupId"))
	}

	if s.NodeGroupCount == nil {
		invalidParams.Add(aws.NewErrParamRequired("NodeGroupCount"))
	}
	if s.RegionalConfigurations != nil {
		for i, v := range s.RegionalConfigurations {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "RegionalConfigurations", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type IncreaseNodeGroupsInGlobalReplicationGroupOutput struct {
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
func (s IncreaseNodeGroupsInGlobalReplicationGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opIncreaseNodeGroupsInGlobalReplicationGroup = "IncreaseNodeGroupsInGlobalReplicationGroup"

// IncreaseNodeGroupsInGlobalReplicationGroupRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Increase the number of node groups in the Global Datastore
//
//    // Example sending a request using IncreaseNodeGroupsInGlobalReplicationGroupRequest.
//    req := client.IncreaseNodeGroupsInGlobalReplicationGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/IncreaseNodeGroupsInGlobalReplicationGroup
func (c *Client) IncreaseNodeGroupsInGlobalReplicationGroupRequest(input *IncreaseNodeGroupsInGlobalReplicationGroupInput) IncreaseNodeGroupsInGlobalReplicationGroupRequest {
	op := &aws.Operation{
		Name:       opIncreaseNodeGroupsInGlobalReplicationGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &IncreaseNodeGroupsInGlobalReplicationGroupInput{}
	}

	req := c.newRequest(op, input, &IncreaseNodeGroupsInGlobalReplicationGroupOutput{})

	return IncreaseNodeGroupsInGlobalReplicationGroupRequest{Request: req, Input: input, Copy: c.IncreaseNodeGroupsInGlobalReplicationGroupRequest}
}

// IncreaseNodeGroupsInGlobalReplicationGroupRequest is the request type for the
// IncreaseNodeGroupsInGlobalReplicationGroup API operation.
type IncreaseNodeGroupsInGlobalReplicationGroupRequest struct {
	*aws.Request
	Input *IncreaseNodeGroupsInGlobalReplicationGroupInput
	Copy  func(*IncreaseNodeGroupsInGlobalReplicationGroupInput) IncreaseNodeGroupsInGlobalReplicationGroupRequest
}

// Send marshals and sends the IncreaseNodeGroupsInGlobalReplicationGroup API request.
func (r IncreaseNodeGroupsInGlobalReplicationGroupRequest) Send(ctx context.Context) (*IncreaseNodeGroupsInGlobalReplicationGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &IncreaseNodeGroupsInGlobalReplicationGroupResponse{
		IncreaseNodeGroupsInGlobalReplicationGroupOutput: r.Request.Data.(*IncreaseNodeGroupsInGlobalReplicationGroupOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// IncreaseNodeGroupsInGlobalReplicationGroupResponse is the response type for the
// IncreaseNodeGroupsInGlobalReplicationGroup API operation.
type IncreaseNodeGroupsInGlobalReplicationGroupResponse struct {
	*IncreaseNodeGroupsInGlobalReplicationGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// IncreaseNodeGroupsInGlobalReplicationGroup request.
func (r *IncreaseNodeGroupsInGlobalReplicationGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
