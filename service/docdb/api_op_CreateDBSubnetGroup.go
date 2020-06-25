// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package docdb

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input to CreateDBSubnetGroup.
type CreateDBSubnetGroupInput struct {
	_ struct{} `type:"structure"`

	// The description for the subnet group.
	//
	// DBSubnetGroupDescription is a required field
	DBSubnetGroupDescription *string `type:"string" required:"true"`

	// The name for the subnet group. This value is stored as a lowercase string.
	//
	// Constraints: Must contain no more than 255 letters, numbers, periods, underscores,
	// spaces, or hyphens. Must not be default.
	//
	// Example: mySubnetgroup
	//
	// DBSubnetGroupName is a required field
	DBSubnetGroupName *string `type:"string" required:"true"`

	// The Amazon EC2 subnet IDs for the subnet group.
	//
	// SubnetIds is a required field
	SubnetIds []string `locationNameList:"SubnetIdentifier" type:"list" required:"true"`

	// The tags to be assigned to the subnet group.
	Tags []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s CreateDBSubnetGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDBSubnetGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDBSubnetGroupInput"}

	if s.DBSubnetGroupDescription == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBSubnetGroupDescription"))
	}

	if s.DBSubnetGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBSubnetGroupName"))
	}

	if s.SubnetIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDBSubnetGroupOutput struct {
	_ struct{} `type:"structure"`

	// Detailed information about a subnet group.
	DBSubnetGroup *DBSubnetGroup `type:"structure"`
}

// String returns the string representation
func (s CreateDBSubnetGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDBSubnetGroup = "CreateDBSubnetGroup"

// CreateDBSubnetGroupRequest returns a request value for making API operation for
// Amazon DocumentDB with MongoDB compatibility.
//
// Creates a new subnet group. subnet groups must contain at least one subnet
// in at least two Availability Zones in the AWS Region.
//
//    // Example sending a request using CreateDBSubnetGroupRequest.
//    req := client.CreateDBSubnetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/CreateDBSubnetGroup
func (c *Client) CreateDBSubnetGroupRequest(input *CreateDBSubnetGroupInput) CreateDBSubnetGroupRequest {
	op := &aws.Operation{
		Name:       opCreateDBSubnetGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDBSubnetGroupInput{}
	}

	req := c.newRequest(op, input, &CreateDBSubnetGroupOutput{})

	return CreateDBSubnetGroupRequest{Request: req, Input: input, Copy: c.CreateDBSubnetGroupRequest}
}

// CreateDBSubnetGroupRequest is the request type for the
// CreateDBSubnetGroup API operation.
type CreateDBSubnetGroupRequest struct {
	*aws.Request
	Input *CreateDBSubnetGroupInput
	Copy  func(*CreateDBSubnetGroupInput) CreateDBSubnetGroupRequest
}

// Send marshals and sends the CreateDBSubnetGroup API request.
func (r CreateDBSubnetGroupRequest) Send(ctx context.Context) (*CreateDBSubnetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDBSubnetGroupResponse{
		CreateDBSubnetGroupOutput: r.Request.Data.(*CreateDBSubnetGroupOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDBSubnetGroupResponse is the response type for the
// CreateDBSubnetGroup API operation.
type CreateDBSubnetGroupResponse struct {
	*CreateDBSubnetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDBSubnetGroup request.
func (r *CreateDBSubnetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
