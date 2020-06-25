// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dax

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type UpdateSubnetGroupInput struct {
	_ struct{} `type:"structure"`

	// A description of the subnet group.
	Description *string `type:"string"`

	// The name of the subnet group.
	//
	// SubnetGroupName is a required field
	SubnetGroupName *string `type:"string" required:"true"`

	// A list of subnet IDs in the subnet group.
	SubnetIds []string `type:"list"`
}

// String returns the string representation
func (s UpdateSubnetGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSubnetGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSubnetGroupInput"}

	if s.SubnetGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateSubnetGroupOutput struct {
	_ struct{} `type:"structure"`

	// The subnet group that has been modified.
	SubnetGroup *SubnetGroup `type:"structure"`
}

// String returns the string representation
func (s UpdateSubnetGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateSubnetGroup = "UpdateSubnetGroup"

// UpdateSubnetGroupRequest returns a request value for making API operation for
// Amazon DynamoDB Accelerator (DAX).
//
// Modifies an existing subnet group.
//
//    // Example sending a request using UpdateSubnetGroupRequest.
//    req := client.UpdateSubnetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/UpdateSubnetGroup
func (c *Client) UpdateSubnetGroupRequest(input *UpdateSubnetGroupInput) UpdateSubnetGroupRequest {
	op := &aws.Operation{
		Name:       opUpdateSubnetGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateSubnetGroupInput{}
	}

	req := c.newRequest(op, input, &UpdateSubnetGroupOutput{})

	return UpdateSubnetGroupRequest{Request: req, Input: input, Copy: c.UpdateSubnetGroupRequest}
}

// UpdateSubnetGroupRequest is the request type for the
// UpdateSubnetGroup API operation.
type UpdateSubnetGroupRequest struct {
	*aws.Request
	Input *UpdateSubnetGroupInput
	Copy  func(*UpdateSubnetGroupInput) UpdateSubnetGroupRequest
}

// Send marshals and sends the UpdateSubnetGroup API request.
func (r UpdateSubnetGroupRequest) Send(ctx context.Context) (*UpdateSubnetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSubnetGroupResponse{
		UpdateSubnetGroupOutput: r.Request.Data.(*UpdateSubnetGroupOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSubnetGroupResponse is the response type for the
// UpdateSubnetGroup API operation.
type UpdateSubnetGroupResponse struct {
	*UpdateSubnetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSubnetGroup request.
func (r *UpdateSubnetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
