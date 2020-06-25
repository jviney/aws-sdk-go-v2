// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CopyDBParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// The identifier or ARN for the source DB parameter group. For information
	// about creating an ARN, see Constructing an ARN for Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.ARN.html#USER_Tagging.ARN.Constructing)
	// in the Amazon RDS User Guide.
	//
	// Constraints:
	//
	//    * Must specify a valid DB parameter group.
	//
	//    * Must specify a valid DB parameter group identifier, for example my-db-param-group,
	//    or a valid ARN.
	//
	// SourceDBParameterGroupIdentifier is a required field
	SourceDBParameterGroupIdentifier *string `type:"string" required:"true"`

	// A list of tags. For more information, see Tagging Amazon RDS Resources (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html)
	// in the Amazon RDS User Guide.
	Tags []Tag `locationNameList:"Tag" type:"list"`

	// A description for the copied DB parameter group.
	//
	// TargetDBParameterGroupDescription is a required field
	TargetDBParameterGroupDescription *string `type:"string" required:"true"`

	// The identifier for the copied DB parameter group.
	//
	// Constraints:
	//
	//    * Can't be null, empty, or blank
	//
	//    * Must contain from 1 to 255 letters, numbers, or hyphens
	//
	//    * First character must be a letter
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens
	//
	// Example: my-db-parameter-group
	//
	// TargetDBParameterGroupIdentifier is a required field
	TargetDBParameterGroupIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CopyDBParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CopyDBParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CopyDBParameterGroupInput"}

	if s.SourceDBParameterGroupIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceDBParameterGroupIdentifier"))
	}

	if s.TargetDBParameterGroupDescription == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetDBParameterGroupDescription"))
	}

	if s.TargetDBParameterGroupIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetDBParameterGroupIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CopyDBParameterGroupOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon RDS DB parameter group.
	//
	// This data type is used as a response element in the DescribeDBParameterGroups
	// action.
	DBParameterGroup *DBParameterGroup `type:"structure"`
}

// String returns the string representation
func (s CopyDBParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCopyDBParameterGroup = "CopyDBParameterGroup"

// CopyDBParameterGroupRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Copies the specified DB parameter group.
//
//    // Example sending a request using CopyDBParameterGroupRequest.
//    req := client.CopyDBParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CopyDBParameterGroup
func (c *Client) CopyDBParameterGroupRequest(input *CopyDBParameterGroupInput) CopyDBParameterGroupRequest {
	op := &aws.Operation{
		Name:       opCopyDBParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CopyDBParameterGroupInput{}
	}

	req := c.newRequest(op, input, &CopyDBParameterGroupOutput{})

	return CopyDBParameterGroupRequest{Request: req, Input: input, Copy: c.CopyDBParameterGroupRequest}
}

// CopyDBParameterGroupRequest is the request type for the
// CopyDBParameterGroup API operation.
type CopyDBParameterGroupRequest struct {
	*aws.Request
	Input *CopyDBParameterGroupInput
	Copy  func(*CopyDBParameterGroupInput) CopyDBParameterGroupRequest
}

// Send marshals and sends the CopyDBParameterGroup API request.
func (r CopyDBParameterGroupRequest) Send(ctx context.Context) (*CopyDBParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CopyDBParameterGroupResponse{
		CopyDBParameterGroupOutput: r.Request.Data.(*CopyDBParameterGroupOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CopyDBParameterGroupResponse is the response type for the
// CopyDBParameterGroup API operation.
type CopyDBParameterGroupResponse struct {
	*CopyDBParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CopyDBParameterGroup request.
func (r *CopyDBParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
