// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costexplorer

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetTagsInput struct {
	_ struct{} `type:"structure"`

	// The token to retrieve the next set of results. AWS provides the token when
	// the response from a previous call has more results than the maximum page
	// size.
	NextPageToken *string `type:"string"`

	// The value that you want to search for.
	SearchString *string `type:"string"`

	// The key of the tag that you want to return values for.
	TagKey *string `type:"string"`

	// The start and end dates for retrieving the dimension values. The start date
	// is inclusive, but the end date is exclusive. For example, if start is 2017-01-01
	// and end is 2017-05-01, then the cost and usage data is retrieved from 2017-01-01
	// up to and including 2017-04-30 but not including 2017-05-01.
	//
	// TimePeriod is a required field
	TimePeriod *DateInterval `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTagsInput"}

	if s.TimePeriod == nil {
		invalidParams.Add(aws.NewErrParamRequired("TimePeriod"))
	}
	if s.TimePeriod != nil {
		if err := s.TimePeriod.Validate(); err != nil {
			invalidParams.AddNested("TimePeriod", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetTagsOutput struct {
	_ struct{} `type:"structure"`

	// The token for the next set of retrievable results. AWS provides the token
	// when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string `type:"string"`

	// The number of query results that AWS returns at a time.
	//
	// ReturnSize is a required field
	ReturnSize *int64 `type:"integer" required:"true"`

	// The tags that match your request.
	//
	// Tags is a required field
	Tags []string `type:"list" required:"true"`

	// The total number of query results.
	//
	// TotalSize is a required field
	TotalSize *int64 `type:"integer" required:"true"`
}

// String returns the string representation
func (s GetTagsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetTags = "GetTags"

// GetTagsRequest returns a request value for making API operation for
// AWS Cost Explorer Service.
//
// Queries for available tag keys and tag values for a specified period. You
// can search the tag values for an arbitrary string.
//
//    // Example sending a request using GetTagsRequest.
//    req := client.GetTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetTags
func (c *Client) GetTagsRequest(input *GetTagsInput) GetTagsRequest {
	op := &aws.Operation{
		Name:       opGetTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetTagsInput{}
	}

	req := c.newRequest(op, input, &GetTagsOutput{})

	return GetTagsRequest{Request: req, Input: input, Copy: c.GetTagsRequest}
}

// GetTagsRequest is the request type for the
// GetTags API operation.
type GetTagsRequest struct {
	*aws.Request
	Input *GetTagsInput
	Copy  func(*GetTagsInput) GetTagsRequest
}

// Send marshals and sends the GetTags API request.
func (r GetTagsRequest) Send(ctx context.Context) (*GetTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTagsResponse{
		GetTagsOutput: r.Request.Data.(*GetTagsOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetTagsResponse is the response type for the
// GetTags API operation.
type GetTagsResponse struct {
	*GetTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTags request.
func (r *GetTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
