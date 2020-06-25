// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetEffectivePermissionsForPathInput struct {
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

	// The Amazon Resource Name (ARN) of the resource for which you want to get
	// permissions.
	//
	// ResourceArn is a required field
	ResourceArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetEffectivePermissionsForPathInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetEffectivePermissionsForPathInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetEffectivePermissionsForPathInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetEffectivePermissionsForPathOutput struct {
	_ struct{} `type:"structure"`

	// A continuation token, if this is not the first call to retrieve this list.
	NextToken *string `type:"string"`

	// A list of the permissions for the specified table or database resource located
	// at the path in Amazon S3.
	Permissions []PrincipalResourcePermissions `type:"list"`
}

// String returns the string representation
func (s GetEffectivePermissionsForPathOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetEffectivePermissionsForPath = "GetEffectivePermissionsForPath"

// GetEffectivePermissionsForPathRequest returns a request value for making API operation for
// AWS Lake Formation.
//
// Returns the permissions for a specified table or database resource located
// at a path in Amazon S3.
//
//    // Example sending a request using GetEffectivePermissionsForPathRequest.
//    req := client.GetEffectivePermissionsForPathRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lakeformation-2017-03-31/GetEffectivePermissionsForPath
func (c *Client) GetEffectivePermissionsForPathRequest(input *GetEffectivePermissionsForPathInput) GetEffectivePermissionsForPathRequest {
	op := &aws.Operation{
		Name:       opGetEffectivePermissionsForPath,
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
		input = &GetEffectivePermissionsForPathInput{}
	}

	req := c.newRequest(op, input, &GetEffectivePermissionsForPathOutput{})

	return GetEffectivePermissionsForPathRequest{Request: req, Input: input, Copy: c.GetEffectivePermissionsForPathRequest}
}

// GetEffectivePermissionsForPathRequest is the request type for the
// GetEffectivePermissionsForPath API operation.
type GetEffectivePermissionsForPathRequest struct {
	*aws.Request
	Input *GetEffectivePermissionsForPathInput
	Copy  func(*GetEffectivePermissionsForPathInput) GetEffectivePermissionsForPathRequest
}

// Send marshals and sends the GetEffectivePermissionsForPath API request.
func (r GetEffectivePermissionsForPathRequest) Send(ctx context.Context) (*GetEffectivePermissionsForPathResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetEffectivePermissionsForPathResponse{
		GetEffectivePermissionsForPathOutput: r.Request.Data.(*GetEffectivePermissionsForPathOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetEffectivePermissionsForPathRequestPaginator returns a paginator for GetEffectivePermissionsForPath.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetEffectivePermissionsForPathRequest(input)
//   p := lakeformation.NewGetEffectivePermissionsForPathRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetEffectivePermissionsForPathPaginator(req GetEffectivePermissionsForPathRequest) GetEffectivePermissionsForPathPaginator {
	return GetEffectivePermissionsForPathPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetEffectivePermissionsForPathInput
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

// GetEffectivePermissionsForPathPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetEffectivePermissionsForPathPaginator struct {
	aws.Pager
}

func (p *GetEffectivePermissionsForPathPaginator) CurrentPage() *GetEffectivePermissionsForPathOutput {
	return p.Pager.CurrentPage().(*GetEffectivePermissionsForPathOutput)
}

// GetEffectivePermissionsForPathResponse is the response type for the
// GetEffectivePermissionsForPath API operation.
type GetEffectivePermissionsForPathResponse struct {
	*GetEffectivePermissionsForPathOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetEffectivePermissionsForPath request.
func (r *GetEffectivePermissionsForPathResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
