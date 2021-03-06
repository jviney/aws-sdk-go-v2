// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeDBProxyTargetsInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the DBProxyTarget to describe.
	//
	// DBProxyName is a required field
	DBProxyName *string `type:"string" required:"true"`

	// This parameter is not currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `min:"20" type:"integer"`

	// The identifier of the DBProxyTargetGroup to describe.
	TargetGroupName *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBProxyTargetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBProxyTargetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDBProxyTargetsInput"}

	if s.DBProxyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBProxyName"))
	}
	if s.MaxRecords != nil && *s.MaxRecords < 20 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxRecords", 20))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeDBProxyTargetsOutput struct {
	_ struct{} `type:"structure"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// An arbitrary number of DBProxyTarget objects, containing details of the corresponding
	// targets.
	Targets []DBProxyTarget `type:"list"`
}

// String returns the string representation
func (s DescribeDBProxyTargetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDBProxyTargets = "DescribeDBProxyTargets"

// DescribeDBProxyTargetsRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
//
// This is prerelease documentation for the RDS Database Proxy feature in preview
// release. It is subject to change.
//
// Returns information about DBProxyTarget objects. This API supports pagination.
//
//    // Example sending a request using DescribeDBProxyTargetsRequest.
//    req := client.DescribeDBProxyTargetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBProxyTargets
func (c *Client) DescribeDBProxyTargetsRequest(input *DescribeDBProxyTargetsInput) DescribeDBProxyTargetsRequest {
	op := &aws.Operation{
		Name:       opDescribeDBProxyTargets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeDBProxyTargetsInput{}
	}

	req := c.newRequest(op, input, &DescribeDBProxyTargetsOutput{})

	return DescribeDBProxyTargetsRequest{Request: req, Input: input, Copy: c.DescribeDBProxyTargetsRequest}
}

// DescribeDBProxyTargetsRequest is the request type for the
// DescribeDBProxyTargets API operation.
type DescribeDBProxyTargetsRequest struct {
	*aws.Request
	Input *DescribeDBProxyTargetsInput
	Copy  func(*DescribeDBProxyTargetsInput) DescribeDBProxyTargetsRequest
}

// Send marshals and sends the DescribeDBProxyTargets API request.
func (r DescribeDBProxyTargetsRequest) Send(ctx context.Context) (*DescribeDBProxyTargetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBProxyTargetsResponse{
		DescribeDBProxyTargetsOutput: r.Request.Data.(*DescribeDBProxyTargetsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeDBProxyTargetsRequestPaginator returns a paginator for DescribeDBProxyTargets.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeDBProxyTargetsRequest(input)
//   p := rds.NewDescribeDBProxyTargetsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeDBProxyTargetsPaginator(req DescribeDBProxyTargetsRequest) DescribeDBProxyTargetsPaginator {
	return DescribeDBProxyTargetsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeDBProxyTargetsInput
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

// DescribeDBProxyTargetsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeDBProxyTargetsPaginator struct {
	aws.Pager
}

func (p *DescribeDBProxyTargetsPaginator) CurrentPage() *DescribeDBProxyTargetsOutput {
	return p.Pager.CurrentPage().(*DescribeDBProxyTargetsOutput)
}

// DescribeDBProxyTargetsResponse is the response type for the
// DescribeDBProxyTargets API operation.
type DescribeDBProxyTargetsResponse struct {
	*DescribeDBProxyTargetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBProxyTargets request.
func (r *DescribeDBProxyTargetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
