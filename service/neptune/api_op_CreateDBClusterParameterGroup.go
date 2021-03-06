// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateDBClusterParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the DB cluster parameter group.
	//
	// Constraints:
	//
	//    * Must match the name of an existing DBClusterParameterGroup.
	//
	// This value is stored as a lowercase string.
	//
	// DBClusterParameterGroupName is a required field
	DBClusterParameterGroupName *string `type:"string" required:"true"`

	// The DB cluster parameter group family name. A DB cluster parameter group
	// can be associated with one and only one DB cluster parameter group family,
	// and can be applied only to a DB cluster running a database engine and engine
	// version compatible with that DB cluster parameter group family.
	//
	// DBParameterGroupFamily is a required field
	DBParameterGroupFamily *string `type:"string" required:"true"`

	// The description for the DB cluster parameter group.
	//
	// Description is a required field
	Description *string `type:"string" required:"true"`

	// The tags to be assigned to the new DB cluster parameter group.
	Tags []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s CreateDBClusterParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDBClusterParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDBClusterParameterGroupInput"}

	if s.DBClusterParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterParameterGroupName"))
	}

	if s.DBParameterGroupFamily == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBParameterGroupFamily"))
	}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDBClusterParameterGroupOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon Neptune DB cluster parameter group.
	//
	// This data type is used as a response element in the DescribeDBClusterParameterGroups
	// action.
	DBClusterParameterGroup *DBClusterParameterGroup `type:"structure"`
}

// String returns the string representation
func (s CreateDBClusterParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDBClusterParameterGroup = "CreateDBClusterParameterGroup"

// CreateDBClusterParameterGroupRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Creates a new DB cluster parameter group.
//
// Parameters in a DB cluster parameter group apply to all of the instances
// in a DB cluster.
//
// A DB cluster parameter group is initially created with the default parameters
// for the database engine used by instances in the DB cluster. To provide custom
// values for any of the parameters, you must modify the group after creating
// it using ModifyDBClusterParameterGroup. Once you've created a DB cluster
// parameter group, you need to associate it with your DB cluster using ModifyDBCluster.
// When you associate a new DB cluster parameter group with a running DB cluster,
// you need to reboot the DB instances in the DB cluster without failover for
// the new DB cluster parameter group and associated settings to take effect.
//
// After you create a DB cluster parameter group, you should wait at least 5
// minutes before creating your first DB cluster that uses that DB cluster parameter
// group as the default parameter group. This allows Amazon Neptune to fully
// complete the create action before the DB cluster parameter group is used
// as the default for a new DB cluster. This is especially important for parameters
// that are critical when creating the default database for a DB cluster, such
// as the character set for the default database defined by the character_set_database
// parameter. You can use the Parameter Groups option of the Amazon Neptune
// console (https://console.aws.amazon.com/rds/) or the DescribeDBClusterParameters
// command to verify that your DB cluster parameter group has been created or
// modified.
//
//    // Example sending a request using CreateDBClusterParameterGroupRequest.
//    req := client.CreateDBClusterParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/CreateDBClusterParameterGroup
func (c *Client) CreateDBClusterParameterGroupRequest(input *CreateDBClusterParameterGroupInput) CreateDBClusterParameterGroupRequest {
	op := &aws.Operation{
		Name:       opCreateDBClusterParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDBClusterParameterGroupInput{}
	}

	req := c.newRequest(op, input, &CreateDBClusterParameterGroupOutput{})

	return CreateDBClusterParameterGroupRequest{Request: req, Input: input, Copy: c.CreateDBClusterParameterGroupRequest}
}

// CreateDBClusterParameterGroupRequest is the request type for the
// CreateDBClusterParameterGroup API operation.
type CreateDBClusterParameterGroupRequest struct {
	*aws.Request
	Input *CreateDBClusterParameterGroupInput
	Copy  func(*CreateDBClusterParameterGroupInput) CreateDBClusterParameterGroupRequest
}

// Send marshals and sends the CreateDBClusterParameterGroup API request.
func (r CreateDBClusterParameterGroupRequest) Send(ctx context.Context) (*CreateDBClusterParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDBClusterParameterGroupResponse{
		CreateDBClusterParameterGroupOutput: r.Request.Data.(*CreateDBClusterParameterGroupOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDBClusterParameterGroupResponse is the response type for the
// CreateDBClusterParameterGroup API operation.
type CreateDBClusterParameterGroupResponse struct {
	*CreateDBClusterParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDBClusterParameterGroup request.
func (r *CreateDBClusterParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
