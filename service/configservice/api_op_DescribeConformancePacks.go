// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeConformancePacksInput struct {
	_ struct{} `type:"structure"`

	// Comma-separated list of conformance pack names for which you want details.
	// If you do not specify any names, AWS Config returns details for all your
	// conformance packs.
	ConformancePackNames []string `type:"list"`

	// The maximum number of conformance packs returned on each page.
	Limit *int64 `type:"integer"`

	// The nextToken string returned in a previous request that you use to request
	// the next page of results in a paginated response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeConformancePacksInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeConformancePacksOutput struct {
	_ struct{} `type:"structure"`

	// Returns a list of ConformancePackDetail objects.
	ConformancePackDetails []ConformancePackDetail `type:"list"`

	// The nextToken string returned in a previous request that you use to request
	// the next page of results in a paginated response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeConformancePacksOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeConformancePacks = "DescribeConformancePacks"

// DescribeConformancePacksRequest returns a request value for making API operation for
// AWS Config.
//
// Returns a list of one or more conformance packs.
//
//    // Example sending a request using DescribeConformancePacksRequest.
//    req := client.DescribeConformancePacksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DescribeConformancePacks
func (c *Client) DescribeConformancePacksRequest(input *DescribeConformancePacksInput) DescribeConformancePacksRequest {
	op := &aws.Operation{
		Name:       opDescribeConformancePacks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeConformancePacksInput{}
	}

	req := c.newRequest(op, input, &DescribeConformancePacksOutput{})

	return DescribeConformancePacksRequest{Request: req, Input: input, Copy: c.DescribeConformancePacksRequest}
}

// DescribeConformancePacksRequest is the request type for the
// DescribeConformancePacks API operation.
type DescribeConformancePacksRequest struct {
	*aws.Request
	Input *DescribeConformancePacksInput
	Copy  func(*DescribeConformancePacksInput) DescribeConformancePacksRequest
}

// Send marshals and sends the DescribeConformancePacks API request.
func (r DescribeConformancePacksRequest) Send(ctx context.Context) (*DescribeConformancePacksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeConformancePacksResponse{
		DescribeConformancePacksOutput: r.Request.Data.(*DescribeConformancePacksOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeConformancePacksResponse is the response type for the
// DescribeConformancePacks API operation.
type DescribeConformancePacksResponse struct {
	*DescribeConformancePacksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeConformancePacks request.
func (r *DescribeConformancePacksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
