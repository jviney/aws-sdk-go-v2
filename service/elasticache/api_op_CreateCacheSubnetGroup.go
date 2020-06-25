// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a CreateCacheSubnetGroup operation.
type CreateCacheSubnetGroupInput struct {
	_ struct{} `type:"structure"`

	// A description for the cache subnet group.
	//
	// CacheSubnetGroupDescription is a required field
	CacheSubnetGroupDescription *string `type:"string" required:"true"`

	// A name for the cache subnet group. This value is stored as a lowercase string.
	//
	// Constraints: Must contain no more than 255 alphanumeric characters or hyphens.
	//
	// Example: mysubnetgroup
	//
	// CacheSubnetGroupName is a required field
	CacheSubnetGroupName *string `type:"string" required:"true"`

	// A list of VPC subnet IDs for the cache subnet group.
	//
	// SubnetIds is a required field
	SubnetIds []string `locationNameList:"SubnetIdentifier" type:"list" required:"true"`
}

// String returns the string representation
func (s CreateCacheSubnetGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCacheSubnetGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCacheSubnetGroupInput"}

	if s.CacheSubnetGroupDescription == nil {
		invalidParams.Add(aws.NewErrParamRequired("CacheSubnetGroupDescription"))
	}

	if s.CacheSubnetGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CacheSubnetGroupName"))
	}

	if s.SubnetIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateCacheSubnetGroupOutput struct {
	_ struct{} `type:"structure"`

	// Represents the output of one of the following operations:
	//
	//    * CreateCacheSubnetGroup
	//
	//    * ModifyCacheSubnetGroup
	CacheSubnetGroup *CacheSubnetGroup `type:"structure"`
}

// String returns the string representation
func (s CreateCacheSubnetGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateCacheSubnetGroup = "CreateCacheSubnetGroup"

// CreateCacheSubnetGroupRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Creates a new cache subnet group.
//
// Use this parameter only when you are creating a cluster in an Amazon Virtual
// Private Cloud (Amazon VPC).
//
//    // Example sending a request using CreateCacheSubnetGroupRequest.
//    req := client.CreateCacheSubnetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/CreateCacheSubnetGroup
func (c *Client) CreateCacheSubnetGroupRequest(input *CreateCacheSubnetGroupInput) CreateCacheSubnetGroupRequest {
	op := &aws.Operation{
		Name:       opCreateCacheSubnetGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCacheSubnetGroupInput{}
	}

	req := c.newRequest(op, input, &CreateCacheSubnetGroupOutput{})

	return CreateCacheSubnetGroupRequest{Request: req, Input: input, Copy: c.CreateCacheSubnetGroupRequest}
}

// CreateCacheSubnetGroupRequest is the request type for the
// CreateCacheSubnetGroup API operation.
type CreateCacheSubnetGroupRequest struct {
	*aws.Request
	Input *CreateCacheSubnetGroupInput
	Copy  func(*CreateCacheSubnetGroupInput) CreateCacheSubnetGroupRequest
}

// Send marshals and sends the CreateCacheSubnetGroup API request.
func (r CreateCacheSubnetGroupRequest) Send(ctx context.Context) (*CreateCacheSubnetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCacheSubnetGroupResponse{
		CreateCacheSubnetGroupOutput: r.Request.Data.(*CreateCacheSubnetGroupOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCacheSubnetGroupResponse is the response type for the
// CreateCacheSubnetGroup API operation.
type CreateCacheSubnetGroupResponse struct {
	*CreateCacheSubnetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCacheSubnetGroup request.
func (r *CreateCacheSubnetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
