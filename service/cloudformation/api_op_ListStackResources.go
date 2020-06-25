// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// The input for the ListStackResource action.
type ListStackResourcesInput struct {
	_ struct{} `type:"structure"`

	// A string that identifies the next page of stack resources that you want to
	// retrieve.
	NextToken *string `min:"1" type:"string"`

	// The name or the unique stack ID that is associated with the stack, which
	// are not always interchangeable:
	//
	//    * Running stacks: You can specify either the stack's name or its unique
	//    stack ID.
	//
	//    * Deleted stacks: You must specify the unique stack ID.
	//
	// Default: There is no default value.
	//
	// StackName is a required field
	StackName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListStackResourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListStackResourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListStackResourcesInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.StackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output for a ListStackResources action.
type ListStackResourcesOutput struct {
	_ struct{} `type:"structure"`

	// If the output exceeds 1 MB, a string that identifies the next page of stack
	// resources. If no additional page exists, this value is null.
	NextToken *string `min:"1" type:"string"`

	// A list of StackResourceSummary structures.
	StackResourceSummaries []StackResourceSummary `type:"list"`
}

// String returns the string representation
func (s ListStackResourcesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListStackResources = "ListStackResources"

// ListStackResourcesRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns descriptions of all resources of the specified stack.
//
// For deleted stacks, ListStackResources returns resource information for up
// to 90 days after the stack has been deleted.
//
//    // Example sending a request using ListStackResourcesRequest.
//    req := client.ListStackResourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ListStackResources
func (c *Client) ListStackResourcesRequest(input *ListStackResourcesInput) ListStackResourcesRequest {
	op := &aws.Operation{
		Name:       opListStackResources,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListStackResourcesInput{}
	}

	req := c.newRequest(op, input, &ListStackResourcesOutput{})

	return ListStackResourcesRequest{Request: req, Input: input, Copy: c.ListStackResourcesRequest}
}

// ListStackResourcesRequest is the request type for the
// ListStackResources API operation.
type ListStackResourcesRequest struct {
	*aws.Request
	Input *ListStackResourcesInput
	Copy  func(*ListStackResourcesInput) ListStackResourcesRequest
}

// Send marshals and sends the ListStackResources API request.
func (r ListStackResourcesRequest) Send(ctx context.Context) (*ListStackResourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListStackResourcesResponse{
		ListStackResourcesOutput: r.Request.Data.(*ListStackResourcesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListStackResourcesRequestPaginator returns a paginator for ListStackResources.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListStackResourcesRequest(input)
//   p := cloudformation.NewListStackResourcesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListStackResourcesPaginator(req ListStackResourcesRequest) ListStackResourcesPaginator {
	return ListStackResourcesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListStackResourcesInput
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

// ListStackResourcesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListStackResourcesPaginator struct {
	aws.Pager
}

func (p *ListStackResourcesPaginator) CurrentPage() *ListStackResourcesOutput {
	return p.Pager.CurrentPage().(*ListStackResourcesOutput)
}

// ListStackResourcesResponse is the response type for the
// ListStackResources API operation.
type ListStackResourcesResponse struct {
	*ListStackResourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListStackResources request.
func (r *ListStackResourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
