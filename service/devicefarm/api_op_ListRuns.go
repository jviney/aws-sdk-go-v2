// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to the list runs operation.
type ListRunsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the project for which you want to list
	// runs.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`
}

// String returns the string representation
func (s ListRunsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRunsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRunsInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}
	if s.NextToken != nil && len(*s.NextToken) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the result of a list runs request.
type ListRunsOutput struct {
	_ struct{} `type:"structure"`

	// If the number of items that are returned is significantly large, this is
	// an identifier that is also returned. It can be used in a subsequent call
	// to this operation to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`

	// Information about the runs.
	Runs []Run `locationName:"runs" type:"list"`
}

// String returns the string representation
func (s ListRunsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListRuns = "ListRuns"

// ListRunsRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Gets information about runs, given an AWS Device Farm project ARN.
//
//    // Example sending a request using ListRunsRequest.
//    req := client.ListRunsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListRuns
func (c *Client) ListRunsRequest(input *ListRunsInput) ListRunsRequest {
	op := &aws.Operation{
		Name:       opListRuns,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListRunsInput{}
	}

	req := c.newRequest(op, input, &ListRunsOutput{})

	return ListRunsRequest{Request: req, Input: input, Copy: c.ListRunsRequest}
}

// ListRunsRequest is the request type for the
// ListRuns API operation.
type ListRunsRequest struct {
	*aws.Request
	Input *ListRunsInput
	Copy  func(*ListRunsInput) ListRunsRequest
}

// Send marshals and sends the ListRuns API request.
func (r ListRunsRequest) Send(ctx context.Context) (*ListRunsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRunsResponse{
		ListRunsOutput: r.Request.Data.(*ListRunsOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListRunsRequestPaginator returns a paginator for ListRuns.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListRunsRequest(input)
//   p := devicefarm.NewListRunsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListRunsPaginator(req ListRunsRequest) ListRunsPaginator {
	return ListRunsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListRunsInput
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

// ListRunsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListRunsPaginator struct {
	aws.Pager
}

func (p *ListRunsPaginator) CurrentPage() *ListRunsOutput {
	return p.Pager.CurrentPage().(*ListRunsOutput)
}

// ListRunsResponse is the response type for the
// ListRuns API operation.
type ListRunsResponse struct {
	*ListRunsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRuns request.
func (r *ListRunsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
