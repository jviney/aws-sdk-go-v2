// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListDomainsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results that are returned per call. Use nextPageToken
	// to obtain further pages of results.
	MaximumPageSize *int64 `locationName:"maximumPageSize" type:"integer"`

	// If NextPageToken is returned there are more results available. The value
	// of NextPageToken is a unique pagination token for each page. Make the call
	// again using the returned token to retrieve the next page. Keep all other
	// arguments unchanged. Each pagination token expires after 60 seconds. Using
	// an expired pagination token will return a 400 error: "Specified token has
	// exceeded its maximum lifetime".
	//
	// The configured maximumPageSize determines how many results can be returned
	// in a single call.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`

	// Specifies the registration status of the domains to list.
	//
	// RegistrationStatus is a required field
	RegistrationStatus RegistrationStatus `locationName:"registrationStatus" type:"string" required:"true" enum:"true"`

	// When set to true, returns the results in reverse order. By default, the results
	// are returned in ascending alphabetical order by name of the domains.
	ReverseOrder *bool `locationName:"reverseOrder" type:"boolean"`
}

// String returns the string representation
func (s ListDomainsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDomainsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDomainsInput"}
	if len(s.RegistrationStatus) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("RegistrationStatus"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains a paginated collection of DomainInfo structures.
type ListDomainsOutput struct {
	_ struct{} `type:"structure"`

	// A list of DomainInfo structures.
	//
	// DomainInfos is a required field
	DomainInfos []DomainInfo `locationName:"domainInfos" type:"list" required:"true"`

	// If a NextPageToken was returned by a previous call, there are more results
	// available. To retrieve the next page of results, make the call again using
	// the returned token in nextPageToken. Keep all other arguments unchanged.
	//
	// The configured maximumPageSize determines how many results can be returned
	// in a single call.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`
}

// String returns the string representation
func (s ListDomainsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDomains = "ListDomains"

// ListDomainsRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Returns the list of domains registered in the account. The results may be
// split into multiple pages. To retrieve subsequent pages, make the call again
// using the nextPageToken returned by the initial call.
//
// This operation is eventually consistent. The results are best effort and
// may not exactly reflect recent updates and changes.
//
// Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF resources
// as follows:
//
//    * Use a Resource element with the domain name to limit the action to only
//    specified domains. The element must be set to arn:aws:swf::AccountID:domain/*,
//    where AccountID is the account ID, with no dashes.
//
//    * Use an Action element to allow or deny permission to call this action.
//
//    * You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using ListDomainsRequest.
//    req := client.ListDomainsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListDomainsRequest(input *ListDomainsInput) ListDomainsRequest {
	op := &aws.Operation{
		Name:       opListDomains,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextPageToken"},
			OutputTokens:    []string{"nextPageToken"},
			LimitToken:      "maximumPageSize",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDomainsInput{}
	}

	req := c.newRequest(op, input, &ListDomainsOutput{})

	return ListDomainsRequest{Request: req, Input: input, Copy: c.ListDomainsRequest}
}

// ListDomainsRequest is the request type for the
// ListDomains API operation.
type ListDomainsRequest struct {
	*aws.Request
	Input *ListDomainsInput
	Copy  func(*ListDomainsInput) ListDomainsRequest
}

// Send marshals and sends the ListDomains API request.
func (r ListDomainsRequest) Send(ctx context.Context) (*ListDomainsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDomainsResponse{
		ListDomainsOutput: r.Request.Data.(*ListDomainsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDomainsRequestPaginator returns a paginator for ListDomains.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDomainsRequest(input)
//   p := swf.NewListDomainsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDomainsPaginator(req ListDomainsRequest) ListDomainsPaginator {
	return ListDomainsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDomainsInput
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

// ListDomainsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDomainsPaginator struct {
	aws.Pager
}

func (p *ListDomainsPaginator) CurrentPage() *ListDomainsOutput {
	return p.Pager.CurrentPage().(*ListDomainsOutput)
}

// ListDomainsResponse is the response type for the
// ListDomains API operation.
type ListDomainsResponse struct {
	*ListDomainsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDomains request.
func (r *ListDomainsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
