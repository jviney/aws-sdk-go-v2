// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GrantPermissionsInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your AWS Lake
	// Formation environment.
	CatalogId *string `min:"1" type:"string"`

	// The permissions granted to the principal on the resource. AWS Lake Formation
	// defines privileges to grant and revoke access to metadata in the Data Catalog
	// and data organized in underlying data storage such as Amazon S3. AWS Lake
	// Formation requires that each principal be authorized to perform a specific
	// task on AWS Lake Formation resources.
	//
	// Permissions is a required field
	Permissions []Permission `type:"list" required:"true"`

	// Indicates a list of the granted permissions that the principal may pass to
	// other users. These permissions may only be a subset of the permissions granted
	// in the Privileges.
	PermissionsWithGrantOption []Permission `type:"list"`

	// The principal to be granted the permissions on the resource. Supported principals
	// are IAM users or IAM roles, and they are defined by their principal type
	// and their ARN.
	//
	// Note that if you define a resource with a particular ARN, then later delete,
	// and recreate a resource with that same ARN, the resource maintains the permissions
	// already granted.
	//
	// Principal is a required field
	Principal *DataLakePrincipal `type:"structure" required:"true"`

	// The resource to which permissions are to be granted. Resources in AWS Lake
	// Formation are the Data Catalog, databases, and tables.
	//
	// Resource is a required field
	Resource *Resource `type:"structure" required:"true"`
}

// String returns the string representation
func (s GrantPermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GrantPermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GrantPermissionsInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.Permissions == nil {
		invalidParams.Add(aws.NewErrParamRequired("Permissions"))
	}

	if s.Principal == nil {
		invalidParams.Add(aws.NewErrParamRequired("Principal"))
	}

	if s.Resource == nil {
		invalidParams.Add(aws.NewErrParamRequired("Resource"))
	}
	if s.Principal != nil {
		if err := s.Principal.Validate(); err != nil {
			invalidParams.AddNested("Principal", err.(aws.ErrInvalidParams))
		}
	}
	if s.Resource != nil {
		if err := s.Resource.Validate(); err != nil {
			invalidParams.AddNested("Resource", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GrantPermissionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GrantPermissionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGrantPermissions = "GrantPermissions"

// GrantPermissionsRequest returns a request value for making API operation for
// AWS Lake Formation.
//
// Grants permissions to the principal to access metadata in the Data Catalog
// and data organized in underlying data storage such as Amazon S3.
//
// For information about permissions, see Security and Access Control to Metadata
// and Data (https://docs-aws.amazon.com/michigan/latest/dg/security-data-access.html).
//
//    // Example sending a request using GrantPermissionsRequest.
//    req := client.GrantPermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lakeformation-2017-03-31/GrantPermissions
func (c *Client) GrantPermissionsRequest(input *GrantPermissionsInput) GrantPermissionsRequest {
	op := &aws.Operation{
		Name:       opGrantPermissions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GrantPermissionsInput{}
	}

	req := c.newRequest(op, input, &GrantPermissionsOutput{})

	return GrantPermissionsRequest{Request: req, Input: input, Copy: c.GrantPermissionsRequest}
}

// GrantPermissionsRequest is the request type for the
// GrantPermissions API operation.
type GrantPermissionsRequest struct {
	*aws.Request
	Input *GrantPermissionsInput
	Copy  func(*GrantPermissionsInput) GrantPermissionsRequest
}

// Send marshals and sends the GrantPermissions API request.
func (r GrantPermissionsRequest) Send(ctx context.Context) (*GrantPermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GrantPermissionsResponse{
		GrantPermissionsOutput: r.Request.Data.(*GrantPermissionsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GrantPermissionsResponse is the response type for the
// GrantPermissions API operation.
type GrantPermissionsResponse struct {
	*GrantPermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GrantPermissions request.
func (r *GrantPermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
