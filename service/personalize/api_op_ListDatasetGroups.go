// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListDatasetGroupsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of dataset groups to return.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// A token returned from the previous call to ListDatasetGroups for getting
	// the next set of dataset groups (if they exist).
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDatasetGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDatasetGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDatasetGroupsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListDatasetGroupsOutput struct {
	_ struct{} `type:"structure"`

	// The list of your dataset groups.
	DatasetGroups []DatasetGroupSummary `locationName:"datasetGroups" type:"list"`

	// A token for getting the next set of dataset groups (if they exist).
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDatasetGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDatasetGroups = "ListDatasetGroups"

// ListDatasetGroupsRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Returns a list of dataset groups. The response provides the properties for
// each dataset group, including the Amazon Resource Name (ARN). For more information
// on dataset groups, see CreateDatasetGroup.
//
//    // Example sending a request using ListDatasetGroupsRequest.
//    req := client.ListDatasetGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/ListDatasetGroups
func (c *Client) ListDatasetGroupsRequest(input *ListDatasetGroupsInput) ListDatasetGroupsRequest {
	op := &aws.Operation{
		Name:       opListDatasetGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDatasetGroupsInput{}
	}

	req := c.newRequest(op, input, &ListDatasetGroupsOutput{})

	return ListDatasetGroupsRequest{Request: req, Input: input, Copy: c.ListDatasetGroupsRequest}
}

// ListDatasetGroupsRequest is the request type for the
// ListDatasetGroups API operation.
type ListDatasetGroupsRequest struct {
	*aws.Request
	Input *ListDatasetGroupsInput
	Copy  func(*ListDatasetGroupsInput) ListDatasetGroupsRequest
}

// Send marshals and sends the ListDatasetGroups API request.
func (r ListDatasetGroupsRequest) Send(ctx context.Context) (*ListDatasetGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDatasetGroupsResponse{
		ListDatasetGroupsOutput: r.Request.Data.(*ListDatasetGroupsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDatasetGroupsRequestPaginator returns a paginator for ListDatasetGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDatasetGroupsRequest(input)
//   p := personalize.NewListDatasetGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDatasetGroupsPaginator(req ListDatasetGroupsRequest) ListDatasetGroupsPaginator {
	return ListDatasetGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDatasetGroupsInput
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

// ListDatasetGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDatasetGroupsPaginator struct {
	aws.Pager
}

func (p *ListDatasetGroupsPaginator) CurrentPage() *ListDatasetGroupsOutput {
	return p.Pager.CurrentPage().(*ListDatasetGroupsOutput)
}

// ListDatasetGroupsResponse is the response type for the
// ListDatasetGroups API operation.
type ListDatasetGroupsResponse struct {
	*ListDatasetGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDatasetGroups request.
func (r *ListDatasetGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
