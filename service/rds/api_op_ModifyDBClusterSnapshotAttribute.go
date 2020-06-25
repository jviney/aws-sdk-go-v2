// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ModifyDBClusterSnapshotAttributeInput struct {
	_ struct{} `type:"structure"`

	// The name of the DB cluster snapshot attribute to modify.
	//
	// To manage authorization for other AWS accounts to copy or restore a manual
	// DB cluster snapshot, set this value to restore.
	//
	// AttributeName is a required field
	AttributeName *string `type:"string" required:"true"`

	// The identifier for the DB cluster snapshot to modify the attributes for.
	//
	// DBClusterSnapshotIdentifier is a required field
	DBClusterSnapshotIdentifier *string `type:"string" required:"true"`

	// A list of DB cluster snapshot attributes to add to the attribute specified
	// by AttributeName.
	//
	// To authorize other AWS accounts to copy or restore a manual DB cluster snapshot,
	// set this list to include one or more AWS account IDs, or all to make the
	// manual DB cluster snapshot restorable by any AWS account. Do not add the
	// all value for any manual DB cluster snapshots that contain private information
	// that you don't want available to all AWS accounts.
	ValuesToAdd []string `locationNameList:"AttributeValue" type:"list"`

	// A list of DB cluster snapshot attributes to remove from the attribute specified
	// by AttributeName.
	//
	// To remove authorization for other AWS accounts to copy or restore a manual
	// DB cluster snapshot, set this list to include one or more AWS account identifiers,
	// or all to remove authorization for any AWS account to copy or restore the
	// DB cluster snapshot. If you specify all, an AWS account whose account ID
	// is explicitly added to the restore attribute can still copy or restore a
	// manual DB cluster snapshot.
	ValuesToRemove []string `locationNameList:"AttributeValue" type:"list"`
}

// String returns the string representation
func (s ModifyDBClusterSnapshotAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBClusterSnapshotAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyDBClusterSnapshotAttributeInput"}

	if s.AttributeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AttributeName"))
	}

	if s.DBClusterSnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterSnapshotIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyDBClusterSnapshotAttributeOutput struct {
	_ struct{} `type:"structure"`

	// Contains the results of a successful call to the DescribeDBClusterSnapshotAttributes
	// API action.
	//
	// Manual DB cluster snapshot attributes are used to authorize other AWS accounts
	// to copy or restore a manual DB cluster snapshot. For more information, see
	// the ModifyDBClusterSnapshotAttribute API action.
	DBClusterSnapshotAttributesResult *DBClusterSnapshotAttributesResult `type:"structure"`
}

// String returns the string representation
func (s ModifyDBClusterSnapshotAttributeOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyDBClusterSnapshotAttribute = "ModifyDBClusterSnapshotAttribute"

// ModifyDBClusterSnapshotAttributeRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Adds an attribute and values to, or removes an attribute and values from,
// a manual DB cluster snapshot.
//
// To share a manual DB cluster snapshot with other AWS accounts, specify restore
// as the AttributeName and use the ValuesToAdd parameter to add a list of IDs
// of the AWS accounts that are authorized to restore the manual DB cluster
// snapshot. Use the value all to make the manual DB cluster snapshot public,
// which means that it can be copied or restored by all AWS accounts. Do not
// add the all value for any manual DB cluster snapshots that contain private
// information that you don't want available to all AWS accounts. If a manual
// DB cluster snapshot is encrypted, it can be shared, but only by specifying
// a list of authorized AWS account IDs for the ValuesToAdd parameter. You can't
// use all as a value for that parameter in this case.
//
// To view which AWS accounts have access to copy or restore a manual DB cluster
// snapshot, or whether a manual DB cluster snapshot public or private, use
// the DescribeDBClusterSnapshotAttributes API action.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using ModifyDBClusterSnapshotAttributeRequest.
//    req := client.ModifyDBClusterSnapshotAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ModifyDBClusterSnapshotAttribute
func (c *Client) ModifyDBClusterSnapshotAttributeRequest(input *ModifyDBClusterSnapshotAttributeInput) ModifyDBClusterSnapshotAttributeRequest {
	op := &aws.Operation{
		Name:       opModifyDBClusterSnapshotAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBClusterSnapshotAttributeInput{}
	}

	req := c.newRequest(op, input, &ModifyDBClusterSnapshotAttributeOutput{})

	return ModifyDBClusterSnapshotAttributeRequest{Request: req, Input: input, Copy: c.ModifyDBClusterSnapshotAttributeRequest}
}

// ModifyDBClusterSnapshotAttributeRequest is the request type for the
// ModifyDBClusterSnapshotAttribute API operation.
type ModifyDBClusterSnapshotAttributeRequest struct {
	*aws.Request
	Input *ModifyDBClusterSnapshotAttributeInput
	Copy  func(*ModifyDBClusterSnapshotAttributeInput) ModifyDBClusterSnapshotAttributeRequest
}

// Send marshals and sends the ModifyDBClusterSnapshotAttribute API request.
func (r ModifyDBClusterSnapshotAttributeRequest) Send(ctx context.Context) (*ModifyDBClusterSnapshotAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyDBClusterSnapshotAttributeResponse{
		ModifyDBClusterSnapshotAttributeOutput: r.Request.Data.(*ModifyDBClusterSnapshotAttributeOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyDBClusterSnapshotAttributeResponse is the response type for the
// ModifyDBClusterSnapshotAttribute API operation.
type ModifyDBClusterSnapshotAttributeResponse struct {
	*ModifyDBClusterSnapshotAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyDBClusterSnapshotAttribute request.
func (r *ModifyDBClusterSnapshotAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
