// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteServiceLinkedRoleInput struct {
	_ struct{} `type:"structure"`

	// The name of the service-linked role to be deleted.
	//
	// RoleName is a required field
	RoleName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteServiceLinkedRoleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteServiceLinkedRoleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteServiceLinkedRoleInput"}

	if s.RoleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleName"))
	}
	if s.RoleName != nil && len(*s.RoleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteServiceLinkedRoleOutput struct {
	_ struct{} `type:"structure"`

	// The deletion task identifier that you can use to check the status of the
	// deletion. This identifier is returned in the format task/aws-service-role/<service-principal-name>/<role-name>/<task-uuid>.
	//
	// DeletionTaskId is a required field
	DeletionTaskId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteServiceLinkedRoleOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteServiceLinkedRole = "DeleteServiceLinkedRole"

// DeleteServiceLinkedRoleRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Submits a service-linked role deletion request and returns a DeletionTaskId,
// which you can use to check the status of the deletion. Before you call this
// operation, confirm that the role has no active sessions and that any resources
// used by the role in the linked service are deleted. If you call this operation
// more than once for the same service-linked role and an earlier deletion task
// is not complete, then the DeletionTaskId of the earlier request is returned.
//
// If you submit a deletion request for a service-linked role whose linked service
// is still accessing a resource, then the deletion task fails. If it fails,
// the GetServiceLinkedRoleDeletionStatus API operation returns the reason for
// the failure, usually including the resources that must be deleted. To delete
// the service-linked role, you must first remove those resources from the linked
// service and then submit the deletion request again. Resources are specific
// to the service that is linked to the role. For more information about removing
// resources from a service, see the AWS documentation (http://docs.aws.amazon.com/)
// for your service.
//
// For more information about service-linked roles, see Roles Terms and Concepts:
// AWS Service-Linked Role (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html#iam-term-service-linked-role)
// in the IAM User Guide.
//
//    // Example sending a request using DeleteServiceLinkedRoleRequest.
//    req := client.DeleteServiceLinkedRoleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeleteServiceLinkedRole
func (c *Client) DeleteServiceLinkedRoleRequest(input *DeleteServiceLinkedRoleInput) DeleteServiceLinkedRoleRequest {
	op := &aws.Operation{
		Name:       opDeleteServiceLinkedRole,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteServiceLinkedRoleInput{}
	}

	req := c.newRequest(op, input, &DeleteServiceLinkedRoleOutput{})

	return DeleteServiceLinkedRoleRequest{Request: req, Input: input, Copy: c.DeleteServiceLinkedRoleRequest}
}

// DeleteServiceLinkedRoleRequest is the request type for the
// DeleteServiceLinkedRole API operation.
type DeleteServiceLinkedRoleRequest struct {
	*aws.Request
	Input *DeleteServiceLinkedRoleInput
	Copy  func(*DeleteServiceLinkedRoleInput) DeleteServiceLinkedRoleRequest
}

// Send marshals and sends the DeleteServiceLinkedRole API request.
func (r DeleteServiceLinkedRoleRequest) Send(ctx context.Context) (*DeleteServiceLinkedRoleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteServiceLinkedRoleResponse{
		DeleteServiceLinkedRoleOutput: r.Request.Data.(*DeleteServiceLinkedRoleOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteServiceLinkedRoleResponse is the response type for the
// DeleteServiceLinkedRole API operation.
type DeleteServiceLinkedRoleResponse struct {
	*DeleteServiceLinkedRoleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteServiceLinkedRole request.
func (r *DeleteServiceLinkedRoleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
