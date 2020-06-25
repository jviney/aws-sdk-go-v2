// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListPermissionsInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your AWS Lake
	// Formation environment.
	CatalogId *string `min:"1" type:"string"`

	// The maximum number of results to return.
	MaxResults *int64 `min:"1" type:"integer"`

	// A continuation token, if this is not the first call to retrieve this list.
	NextToken *string `type:"string"`

	// Specifies a principal to filter the permissions returned.
	Principal *DataLakePrincipal `type:"structure"`

	// A resource where you will get a list of the principal permissions.
	//
	// This operation does not support getting privileges on a table with columns.
	// Instead, call this operation on the table, and the operation returns the
	// table and the table w columns.
	Resource *Resource `type:"structure"`

	// Specifies a resource type to filter the permissions returned.
	ResourceType DataLakeResourceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListPermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPermissionsInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
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

type ListPermissionsOutput struct {
	_ struct{} `type:"structure"`

	// A continuation token, if this is not the first call to retrieve this list.
	NextToken *string `type:"string"`

	// A list of principals and their permissions on the resource for the specified
	// principal and resource types.
	PrincipalResourcePermissions []PrincipalResourcePermissions `type:"list"`
}

// String returns the string representation
func (s ListPermissionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListPermissions = "ListPermissions"

// ListPermissionsRequest returns a request value for making API operation for
// AWS Lake Formation.
//
// Returns a list of the principal permissions on the resource, filtered by
// the permissions of the caller. For example, if you are granted an ALTER permission,
// you are able to see only the principal permissions for ALTER.
//
// This operation returns only those permissions that have been explicitly granted.
//
// For information about permissions, see Security and Access Control to Metadata
// and Data (https://docs-aws.amazon.com/michigan/latest/dg/security-data-access.html).
//
//    // Example sending a request using ListPermissionsRequest.
//    req := client.ListPermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lakeformation-2017-03-31/ListPermissions
func (c *Client) ListPermissionsRequest(input *ListPermissionsInput) ListPermissionsRequest {
	op := &aws.Operation{
		Name:       opListPermissions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListPermissionsInput{}
	}

	req := c.newRequest(op, input, &ListPermissionsOutput{})

	return ListPermissionsRequest{Request: req, Input: input, Copy: c.ListPermissionsRequest}
}

// ListPermissionsRequest is the request type for the
// ListPermissions API operation.
type ListPermissionsRequest struct {
	*aws.Request
	Input *ListPermissionsInput
	Copy  func(*ListPermissionsInput) ListPermissionsRequest
}

// Send marshals and sends the ListPermissions API request.
func (r ListPermissionsRequest) Send(ctx context.Context) (*ListPermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPermissionsResponse{
		ListPermissionsOutput: r.Request.Data.(*ListPermissionsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPermissionsRequestPaginator returns a paginator for ListPermissions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPermissionsRequest(input)
//   p := lakeformation.NewListPermissionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPermissionsPaginator(req ListPermissionsRequest) ListPermissionsPaginator {
	return ListPermissionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListPermissionsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListPermissionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPermissionsPaginator struct {
	aws.Pager
}

func (p *ListPermissionsPaginator) CurrentPage() *ListPermissionsOutput {
	return p.Pager.CurrentPage().(*ListPermissionsOutput)
}

// ListPermissionsResponse is the response type for the
// ListPermissions API operation.
type ListPermissionsResponse struct {
	*ListPermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPermissions request.
func (r *ListPermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
