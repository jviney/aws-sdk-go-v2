// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListUsersInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in a single call.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token to use to retrieve the next page of results. The first call does
	// not contain any tokens.
	NextToken *string `min:"1" type:"string"`

	// The identifier for the organization under which the users exist.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListUsersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListUsersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListUsersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListUsersOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is `null`
	// when there are no more results to return.
	NextToken *string `min:"1" type:"string"`

	// The overview of users for an organization.
	Users []User `type:"list"`
}

// String returns the string representation
func (s ListUsersOutput) String() string {
	return awsutil.Prettify(s)
}

const opListUsers = "ListUsers"

// ListUsersRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Returns summaries of the organization's users.
//
//    // Example sending a request using ListUsersRequest.
//    req := client.ListUsersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/ListUsers
func (c *Client) ListUsersRequest(input *ListUsersInput) ListUsersRequest {
	op := &aws.Operation{
		Name:       opListUsers,
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
		input = &ListUsersInput{}
	}

	req := c.newRequest(op, input, &ListUsersOutput{})

	return ListUsersRequest{Request: req, Input: input, Copy: c.ListUsersRequest}
}

// ListUsersRequest is the request type for the
// ListUsers API operation.
type ListUsersRequest struct {
	*aws.Request
	Input *ListUsersInput
	Copy  func(*ListUsersInput) ListUsersRequest
}

// Send marshals and sends the ListUsers API request.
func (r ListUsersRequest) Send(ctx context.Context) (*ListUsersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListUsersResponse{
		ListUsersOutput: r.Request.Data.(*ListUsersOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListUsersRequestPaginator returns a paginator for ListUsers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListUsersRequest(input)
//   p := workmail.NewListUsersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListUsersPaginator(req ListUsersRequest) ListUsersPaginator {
	return ListUsersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListUsersInput
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

// ListUsersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListUsersPaginator struct {
	aws.Pager
}

func (p *ListUsersPaginator) CurrentPage() *ListUsersOutput {
	return p.Pager.CurrentPage().(*ListUsersOutput)
}

// ListUsersResponse is the response type for the
// ListUsers API operation.
type ListUsersResponse struct {
	*ListUsersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListUsers request.
func (r *ListUsersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
