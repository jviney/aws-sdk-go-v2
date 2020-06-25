// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ModifyReplicationSubnetGroupInput struct {
	_ struct{} `type:"structure"`

	// A description for the replication instance subnet group.
	ReplicationSubnetGroupDescription *string `type:"string"`

	// The name of the replication instance subnet group.
	//
	// ReplicationSubnetGroupIdentifier is a required field
	ReplicationSubnetGroupIdentifier *string `type:"string" required:"true"`

	// A list of subnet IDs.
	//
	// SubnetIds is a required field
	SubnetIds []string `type:"list" required:"true"`
}

// String returns the string representation
func (s ModifyReplicationSubnetGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyReplicationSubnetGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyReplicationSubnetGroupInput"}

	if s.ReplicationSubnetGroupIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationSubnetGroupIdentifier"))
	}

	if s.SubnetIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyReplicationSubnetGroupOutput struct {
	_ struct{} `type:"structure"`

	// The modified replication subnet group.
	ReplicationSubnetGroup *ReplicationSubnetGroup `type:"structure"`
}

// String returns the string representation
func (s ModifyReplicationSubnetGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyReplicationSubnetGroup = "ModifyReplicationSubnetGroup"

// ModifyReplicationSubnetGroupRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Modifies the settings for the specified replication subnet group.
//
//    // Example sending a request using ModifyReplicationSubnetGroupRequest.
//    req := client.ModifyReplicationSubnetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/ModifyReplicationSubnetGroup
func (c *Client) ModifyReplicationSubnetGroupRequest(input *ModifyReplicationSubnetGroupInput) ModifyReplicationSubnetGroupRequest {
	op := &aws.Operation{
		Name:       opModifyReplicationSubnetGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyReplicationSubnetGroupInput{}
	}

	req := c.newRequest(op, input, &ModifyReplicationSubnetGroupOutput{})

	return ModifyReplicationSubnetGroupRequest{Request: req, Input: input, Copy: c.ModifyReplicationSubnetGroupRequest}
}

// ModifyReplicationSubnetGroupRequest is the request type for the
// ModifyReplicationSubnetGroup API operation.
type ModifyReplicationSubnetGroupRequest struct {
	*aws.Request
	Input *ModifyReplicationSubnetGroupInput
	Copy  func(*ModifyReplicationSubnetGroupInput) ModifyReplicationSubnetGroupRequest
}

// Send marshals and sends the ModifyReplicationSubnetGroup API request.
func (r ModifyReplicationSubnetGroupRequest) Send(ctx context.Context) (*ModifyReplicationSubnetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyReplicationSubnetGroupResponse{
		ModifyReplicationSubnetGroupOutput: r.Request.Data.(*ModifyReplicationSubnetGroupOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyReplicationSubnetGroupResponse is the response type for the
// ModifyReplicationSubnetGroup API operation.
type ModifyReplicationSubnetGroupResponse struct {
	*ModifyReplicationSubnetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyReplicationSubnetGroup request.
func (r *ModifyReplicationSubnetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
