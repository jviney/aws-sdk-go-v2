// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package machinelearning

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeDataSourcesInput struct {
	_ struct{} `type:"structure"`

	// The equal to operator. The DataSource results will have FilterVariable values
	// that exactly match the value specified with EQ.
	EQ *string `type:"string"`

	// Use one of the following variables to filter a list of DataSource:
	//
	//    * CreatedAt - Sets the search criteria to DataSource creation dates.
	//
	//    * Status - Sets the search criteria to DataSource statuses.
	//
	//    * Name - Sets the search criteria to the contents of DataSource Name.
	//
	//    * DataUri - Sets the search criteria to the URI of data files used to
	//    create the DataSource. The URI can identify either a file or an Amazon
	//    Simple Storage Service (Amazon S3) bucket or directory.
	//
	//    * IAMUser - Sets the search criteria to the user account that invoked
	//    the DataSource creation.
	FilterVariable DataSourceFilterVariable `type:"string" enum:"true"`

	// The greater than or equal to operator. The DataSource results will have FilterVariable
	// values that are greater than or equal to the value specified with GE.
	GE *string `type:"string"`

	// The greater than operator. The DataSource results will have FilterVariable
	// values that are greater than the value specified with GT.
	GT *string `type:"string"`

	// The less than or equal to operator. The DataSource results will have FilterVariable
	// values that are less than or equal to the value specified with LE.
	LE *string `type:"string"`

	// The less than operator. The DataSource results will have FilterVariable values
	// that are less than the value specified with LT.
	LT *string `type:"string"`

	// The maximum number of DataSource to include in the result.
	Limit *int64 `min:"1" type:"integer"`

	// The not equal to operator. The DataSource results will have FilterVariable
	// values not equal to the value specified with NE.
	NE *string `type:"string"`

	// The ID of the page in the paginated results.
	NextToken *string `type:"string"`

	// A string that is found at the beginning of a variable, such as Name or Id.
	//
	// For example, a DataSource could have the Name 2014-09-09-HolidayGiftMailer.
	// To search for this DataSource, select Name for the FilterVariable and any
	// of the following strings for the Prefix:
	//
	//    * 2014-09
	//
	//    * 2014-09-09
	//
	//    * 2014-09-09-Holiday
	Prefix *string `type:"string"`

	// A two-value parameter that determines the sequence of the resulting list
	// of DataSource.
	//
	//    * asc - Arranges the list in ascending order (A-Z, 0-9).
	//
	//    * dsc - Arranges the list in descending order (Z-A, 9-0).
	//
	// Results are sorted by FilterVariable.
	SortOrder SortOrder `type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeDataSourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDataSourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDataSourcesInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the query results from a DescribeDataSources operation. The content
// is essentially a list of DataSource.
type DescribeDataSourcesOutput struct {
	_ struct{} `type:"structure"`

	// An ID of the next page in the paginated results that indicates at least one
	// more page follows.
	NextToken *string `type:"string"`

	// A list of DataSource that meet the search criteria.
	Results []DataSource `type:"list"`
}

// String returns the string representation
func (s DescribeDataSourcesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDataSources = "DescribeDataSources"

// DescribeDataSourcesRequest returns a request value for making API operation for
// Amazon Machine Learning.
//
// Returns a list of DataSource that match the search criteria in the request.
//
//    // Example sending a request using DescribeDataSourcesRequest.
//    req := client.DescribeDataSourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeDataSourcesRequest(input *DescribeDataSourcesInput) DescribeDataSourcesRequest {
	op := &aws.Operation{
		Name:       opDescribeDataSources,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeDataSourcesInput{}
	}

	req := c.newRequest(op, input, &DescribeDataSourcesOutput{})

	return DescribeDataSourcesRequest{Request: req, Input: input, Copy: c.DescribeDataSourcesRequest}
}

// DescribeDataSourcesRequest is the request type for the
// DescribeDataSources API operation.
type DescribeDataSourcesRequest struct {
	*aws.Request
	Input *DescribeDataSourcesInput
	Copy  func(*DescribeDataSourcesInput) DescribeDataSourcesRequest
}

// Send marshals and sends the DescribeDataSources API request.
func (r DescribeDataSourcesRequest) Send(ctx context.Context) (*DescribeDataSourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDataSourcesResponse{
		DescribeDataSourcesOutput: r.Request.Data.(*DescribeDataSourcesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeDataSourcesRequestPaginator returns a paginator for DescribeDataSources.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeDataSourcesRequest(input)
//   p := machinelearning.NewDescribeDataSourcesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeDataSourcesPaginator(req DescribeDataSourcesRequest) DescribeDataSourcesPaginator {
	return DescribeDataSourcesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeDataSourcesInput
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

// DescribeDataSourcesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeDataSourcesPaginator struct {
	aws.Pager
}

func (p *DescribeDataSourcesPaginator) CurrentPage() *DescribeDataSourcesOutput {
	return p.Pager.CurrentPage().(*DescribeDataSourcesOutput)
}

// DescribeDataSourcesResponse is the response type for the
// DescribeDataSources API operation.
type DescribeDataSourcesResponse struct {
	*DescribeDataSourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDataSources request.
func (r *DescribeDataSourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
