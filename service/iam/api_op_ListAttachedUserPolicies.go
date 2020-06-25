// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListAttachedUserPoliciesInput struct {
	_ struct{} `type:"structure"`

	// Use this parameter only when paginating results and only after you receive
	// a response indicating that the results are truncated. Set it to the value
	// of the Marker element in the response that you received to indicate where
	// the next call should start.
	Marker *string `min:"1" type:"string"`

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true.
	//
	// If you do not include this parameter, the number of items defaults to 100.
	// Note that IAM might return fewer results, even when there are more results
	// available. In that case, the IsTruncated response element returns true, and
	// Marker contains a value to include in the subsequent call that tells the
	// service where to continue from.
	MaxItems *int64 `min:"1" type:"integer"`

	// The path prefix for filtering the results. This parameter is optional. If
	// it is not included, it defaults to a slash (/), listing all policies.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of either a forward slash (/) by itself
	// or a string that must begin and end with forward slashes. In addition, it
	// can contain any ASCII character from the ! (\u0021) through the DEL character
	// (\u007F), including most punctuation characters, digits, and upper and lowercased
	// letters.
	PathPrefix *string `min:"1" type:"string"`

	// The name (friendly name, not ARN) of the user to list attached policies for.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListAttachedUserPoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAttachedUserPoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAttachedUserPoliciesInput"}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}
	if s.PathPrefix != nil && len(*s.PathPrefix) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PathPrefix", 1))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful ListAttachedUserPolicies request.
type ListAttachedUserPoliciesOutput struct {
	_ struct{} `type:"structure"`

	// A list of the attached policies.
	AttachedPolicies []AttachedPolicy `type:"list"`

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer
	// than the MaxItems number of results even when there are more results available.
	// We recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated *bool `type:"boolean"`

	// When IsTruncated is true, this element is present and contains the value
	// to use for the Marker parameter in a subsequent pagination request.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s ListAttachedUserPoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListAttachedUserPolicies = "ListAttachedUserPolicies"

// ListAttachedUserPoliciesRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Lists all managed policies that are attached to the specified IAM user.
//
// An IAM user can also have inline policies embedded with it. To list the inline
// policies for a user, use the ListUserPolicies API. For information about
// policies, see Managed Policies and Inline Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
//
// You can paginate the results using the MaxItems and Marker parameters. You
// can use the PathPrefix parameter to limit the list of policies to only those
// matching the specified path prefix. If there are no policies attached to
// the specified group (or none that match the specified path prefix), the operation
// returns an empty list.
//
//    // Example sending a request using ListAttachedUserPoliciesRequest.
//    req := client.ListAttachedUserPoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListAttachedUserPolicies
func (c *Client) ListAttachedUserPoliciesRequest(input *ListAttachedUserPoliciesInput) ListAttachedUserPoliciesRequest {
	op := &aws.Operation{
		Name:       opListAttachedUserPolicies,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxItems",
			TruncationToken: "IsTruncated",
		},
	}

	if input == nil {
		input = &ListAttachedUserPoliciesInput{}
	}

	req := c.newRequest(op, input, &ListAttachedUserPoliciesOutput{})

	return ListAttachedUserPoliciesRequest{Request: req, Input: input, Copy: c.ListAttachedUserPoliciesRequest}
}

// ListAttachedUserPoliciesRequest is the request type for the
// ListAttachedUserPolicies API operation.
type ListAttachedUserPoliciesRequest struct {
	*aws.Request
	Input *ListAttachedUserPoliciesInput
	Copy  func(*ListAttachedUserPoliciesInput) ListAttachedUserPoliciesRequest
}

// Send marshals and sends the ListAttachedUserPolicies API request.
func (r ListAttachedUserPoliciesRequest) Send(ctx context.Context) (*ListAttachedUserPoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAttachedUserPoliciesResponse{
		ListAttachedUserPoliciesOutput: r.Request.Data.(*ListAttachedUserPoliciesOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListAttachedUserPoliciesRequestPaginator returns a paginator for ListAttachedUserPolicies.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListAttachedUserPoliciesRequest(input)
//   p := iam.NewListAttachedUserPoliciesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListAttachedUserPoliciesPaginator(req ListAttachedUserPoliciesRequest) ListAttachedUserPoliciesPaginator {
	return ListAttachedUserPoliciesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListAttachedUserPoliciesInput
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

// ListAttachedUserPoliciesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListAttachedUserPoliciesPaginator struct {
	aws.Pager
}

func (p *ListAttachedUserPoliciesPaginator) CurrentPage() *ListAttachedUserPoliciesOutput {
	return p.Pager.CurrentPage().(*ListAttachedUserPoliciesOutput)
}

// ListAttachedUserPoliciesResponse is the response type for the
// ListAttachedUserPolicies API operation.
type ListAttachedUserPoliciesResponse struct {
	*ListAttachedUserPoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAttachedUserPolicies request.
func (r *ListAttachedUserPoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
