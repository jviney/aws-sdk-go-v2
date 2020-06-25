// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListMigrationTasksInput struct {
	_ struct{} `type:"structure"`

	// Value to specify how many results are returned per page.
	MaxResults *int64 `min:"1" type:"integer"`

	// If a NextToken was returned by a previous call, there are more results available.
	// To retrieve the next page of results, make the call again using the returned
	// token in NextToken.
	NextToken *string `type:"string"`

	// Filter migration tasks by discovered resource name.
	ResourceName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListMigrationTasksInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListMigrationTasksInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListMigrationTasksInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.ResourceName != nil && len(*s.ResourceName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListMigrationTasksOutput struct {
	_ struct{} `type:"structure"`

	// Lists the migration task's summary which includes: MigrationTaskName, ProgressPercent,
	// ProgressUpdateStream, Status, and the UpdateDateTime for each task.
	MigrationTaskSummaryList []MigrationTaskSummary `type:"list"`

	// If there are more migration tasks than the max result, return the next token
	// to be passed to the next call as a bookmark of where to start from.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListMigrationTasksOutput) String() string {
	return awsutil.Prettify(s)
}

const opListMigrationTasks = "ListMigrationTasks"

// ListMigrationTasksRequest returns a request value for making API operation for
// AWS Migration Hub.
//
// Lists all, or filtered by resource name, migration tasks associated with
// the user account making this call. This API has the following traits:
//
//    * Can show a summary list of the most recent migration tasks.
//
//    * Can show a summary list of migration tasks associated with a given discovered
//    resource.
//
//    * Lists migration tasks in a paginated interface.
//
//    // Example sending a request using ListMigrationTasksRequest.
//    req := client.ListMigrationTasksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/ListMigrationTasks
func (c *Client) ListMigrationTasksRequest(input *ListMigrationTasksInput) ListMigrationTasksRequest {
	op := &aws.Operation{
		Name:       opListMigrationTasks,
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
		input = &ListMigrationTasksInput{}
	}

	req := c.newRequest(op, input, &ListMigrationTasksOutput{})

	return ListMigrationTasksRequest{Request: req, Input: input, Copy: c.ListMigrationTasksRequest}
}

// ListMigrationTasksRequest is the request type for the
// ListMigrationTasks API operation.
type ListMigrationTasksRequest struct {
	*aws.Request
	Input *ListMigrationTasksInput
	Copy  func(*ListMigrationTasksInput) ListMigrationTasksRequest
}

// Send marshals and sends the ListMigrationTasks API request.
func (r ListMigrationTasksRequest) Send(ctx context.Context) (*ListMigrationTasksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListMigrationTasksResponse{
		ListMigrationTasksOutput: r.Request.Data.(*ListMigrationTasksOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListMigrationTasksRequestPaginator returns a paginator for ListMigrationTasks.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListMigrationTasksRequest(input)
//   p := migrationhub.NewListMigrationTasksRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListMigrationTasksPaginator(req ListMigrationTasksRequest) ListMigrationTasksPaginator {
	return ListMigrationTasksPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListMigrationTasksInput
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

// ListMigrationTasksPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListMigrationTasksPaginator struct {
	aws.Pager
}

func (p *ListMigrationTasksPaginator) CurrentPage() *ListMigrationTasksOutput {
	return p.Pager.CurrentPage().(*ListMigrationTasksOutput)
}

// ListMigrationTasksResponse is the response type for the
// ListMigrationTasks API operation.
type ListMigrationTasksResponse struct {
	*ListMigrationTasksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListMigrationTasks request.
func (r *ListMigrationTasksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
