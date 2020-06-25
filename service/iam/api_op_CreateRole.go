// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateRoleInput struct {
	_ struct{} `type:"structure"`

	// The trust relationship policy document that grants an entity permission to
	// assume the role.
	//
	// In IAM, you must provide a JSON policy that has been converted to a string.
	// However, for AWS CloudFormation templates formatted in YAML, you can provide
	// the policy in JSON or YAML format. AWS CloudFormation always converts a YAML
	// policy to JSON format before submitting it to IAM.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) used to validate this
	// parameter is a string of characters consisting of the following:
	//
	//    * Any printable ASCII character ranging from the space character (\u0020)
	//    through the end of the ASCII character range
	//
	//    * The printable characters in the Basic Latin and Latin-1 Supplement character
	//    set (through \u00FF)
	//
	//    * The special characters tab (\u0009), line feed (\u000A), and carriage
	//    return (\u000D)
	//
	// Upon success, the response includes the same trust policy in JSON format.
	//
	// AssumeRolePolicyDocument is a required field
	AssumeRolePolicyDocument *string `min:"1" type:"string" required:"true"`

	// A description of the role.
	Description *string `type:"string"`

	// The maximum session duration (in seconds) that you want to set for the specified
	// role. If you do not specify a value for this setting, the default maximum
	// of one hour is applied. This setting can have a value from 1 hour to 12 hours.
	//
	// Anyone who assumes the role from the AWS CLI or API can use the DurationSeconds
	// API parameter or the duration-seconds CLI parameter to request a longer session.
	// The MaxSessionDuration setting determines the maximum duration that can be
	// requested using the DurationSeconds parameter. If users don't specify a value
	// for the DurationSeconds parameter, their security credentials are valid for
	// one hour by default. This applies when you use the AssumeRole* API operations
	// or the assume-role* CLI operations but does not apply when you use those
	// operations to create a console URL. For more information, see Using IAM Roles
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html) in the
	// IAM User Guide.
	MaxSessionDuration *int64 `min:"3600" type:"integer"`

	// The path to the role. For more information about paths, see IAM Identifiers
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html)
	// in the IAM User Guide.
	//
	// This parameter is optional. If it is not included, it defaults to a slash
	// (/).
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of either a forward slash (/) by itself
	// or a string that must begin and end with forward slashes. In addition, it
	// can contain any ASCII character from the ! (\u0021) through the DEL character
	// (\u007F), including most punctuation characters, digits, and upper and lowercased
	// letters.
	Path *string `min:"1" type:"string"`

	// The ARN of the policy that is used to set the permissions boundary for the
	// role.
	PermissionsBoundary *string `min:"20" type:"string"`

	// The name of the role to create.
	//
	// IAM user, group, role, and policy names must be unique within the account.
	// Names are not distinguished by case. For example, you cannot create resources
	// named both "MyResource" and "myresource".
	//
	// RoleName is a required field
	RoleName *string `min:"1" type:"string" required:"true"`

	// A list of tags that you want to attach to the newly created role. Each tag
	// consists of a key name and an associated value. For more information about
	// tagging, see Tagging IAM Identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html)
	// in the IAM User Guide.
	//
	// If any one of the tags is invalid or if you exceed the allowed number of
	// tags per role, then the entire request fails and the role is not created.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateRoleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRoleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRoleInput"}

	if s.AssumeRolePolicyDocument == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssumeRolePolicyDocument"))
	}
	if s.AssumeRolePolicyDocument != nil && len(*s.AssumeRolePolicyDocument) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssumeRolePolicyDocument", 1))
	}
	if s.MaxSessionDuration != nil && *s.MaxSessionDuration < 3600 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxSessionDuration", 3600))
	}
	if s.Path != nil && len(*s.Path) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Path", 1))
	}
	if s.PermissionsBoundary != nil && len(*s.PermissionsBoundary) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("PermissionsBoundary", 20))
	}

	if s.RoleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleName"))
	}
	if s.RoleName != nil && len(*s.RoleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleName", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful CreateRole request.
type CreateRoleOutput struct {
	_ struct{} `type:"structure"`

	// A structure containing details about the new role.
	//
	// Role is a required field
	Role *Role `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateRoleOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateRole = "CreateRole"

// CreateRoleRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Creates a new role for your AWS account. For more information about roles,
// go to IAM Roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/WorkingWithRoles.html).
// For information about limitations on role names and the number of roles you
// can create, go to Limitations on IAM Entities (https://docs.aws.amazon.com/IAM/latest/UserGuide/LimitationsOnEntities.html)
// in the IAM User Guide.
//
//    // Example sending a request using CreateRoleRequest.
//    req := client.CreateRoleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/CreateRole
func (c *Client) CreateRoleRequest(input *CreateRoleInput) CreateRoleRequest {
	op := &aws.Operation{
		Name:       opCreateRole,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateRoleInput{}
	}

	req := c.newRequest(op, input, &CreateRoleOutput{})

	return CreateRoleRequest{Request: req, Input: input, Copy: c.CreateRoleRequest}
}

// CreateRoleRequest is the request type for the
// CreateRole API operation.
type CreateRoleRequest struct {
	*aws.Request
	Input *CreateRoleInput
	Copy  func(*CreateRoleInput) CreateRoleRequest
}

// Send marshals and sends the CreateRole API request.
func (r CreateRoleRequest) Send(ctx context.Context) (*CreateRoleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRoleResponse{
		CreateRoleOutput: r.Request.Data.(*CreateRoleOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRoleResponse is the response type for the
// CreateRole API operation.
type CreateRoleResponse struct {
	*CreateRoleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRole request.
func (r *CreateRoleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
