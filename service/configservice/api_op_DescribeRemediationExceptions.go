// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeRemediationExceptionsInput struct {
	_ struct{} `type:"structure"`

	// The name of the AWS Config rule.
	//
	// ConfigRuleName is a required field
	ConfigRuleName *string `min:"1" type:"string" required:"true"`

	// The maximum number of RemediationExceptionResourceKey returned on each page.
	// The default is 25. If you specify 0, AWS Config uses the default.
	Limit *int64 `type:"integer"`

	// The nextToken string returned in a previous request that you use to request
	// the next page of results in a paginated response.
	NextToken *string `type:"string"`

	// An exception list of resource exception keys to be processed with the current
	// request. AWS Config adds exception for each resource key. For example, AWS
	// Config adds 3 exceptions for 3 resource keys.
	ResourceKeys []RemediationExceptionResourceKey `min:"1" type:"list"`
}

// String returns the string representation
func (s DescribeRemediationExceptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRemediationExceptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeRemediationExceptionsInput"}

	if s.ConfigRuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigRuleName"))
	}
	if s.ConfigRuleName != nil && len(*s.ConfigRuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConfigRuleName", 1))
	}
	if s.ResourceKeys != nil && len(s.ResourceKeys) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceKeys", 1))
	}
	if s.ResourceKeys != nil {
		for i, v := range s.ResourceKeys {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ResourceKeys", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeRemediationExceptionsOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken string returned in a previous request that you use to request
	// the next page of results in a paginated response.
	NextToken *string `type:"string"`

	// Returns a list of remediation exception objects.
	RemediationExceptions []RemediationException `type:"list"`
}

// String returns the string representation
func (s DescribeRemediationExceptionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeRemediationExceptions = "DescribeRemediationExceptions"

// DescribeRemediationExceptionsRequest returns a request value for making API operation for
// AWS Config.
//
// Returns the details of one or more remediation exceptions. A detailed view
// of a remediation exception for a set of resources that includes an explanation
// of an exception and the time when the exception will be deleted. When you
// specify the limit and the next token, you receive a paginated response.
//
// When you specify the limit and the next token, you receive a paginated response.
//
// Limit and next token are not applicable if you request resources in batch.
// It is only applicable, when you request all resources.
//
//    // Example sending a request using DescribeRemediationExceptionsRequest.
//    req := client.DescribeRemediationExceptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DescribeRemediationExceptions
func (c *Client) DescribeRemediationExceptionsRequest(input *DescribeRemediationExceptionsInput) DescribeRemediationExceptionsRequest {
	op := &aws.Operation{
		Name:       opDescribeRemediationExceptions,
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
		input = &DescribeRemediationExceptionsInput{}
	}

	req := c.newRequest(op, input, &DescribeRemediationExceptionsOutput{})

	return DescribeRemediationExceptionsRequest{Request: req, Input: input, Copy: c.DescribeRemediationExceptionsRequest}
}

// DescribeRemediationExceptionsRequest is the request type for the
// DescribeRemediationExceptions API operation.
type DescribeRemediationExceptionsRequest struct {
	*aws.Request
	Input *DescribeRemediationExceptionsInput
	Copy  func(*DescribeRemediationExceptionsInput) DescribeRemediationExceptionsRequest
}

// Send marshals and sends the DescribeRemediationExceptions API request.
func (r DescribeRemediationExceptionsRequest) Send(ctx context.Context) (*DescribeRemediationExceptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRemediationExceptionsResponse{
		DescribeRemediationExceptionsOutput: r.Request.Data.(*DescribeRemediationExceptionsOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeRemediationExceptionsRequestPaginator returns a paginator for DescribeRemediationExceptions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeRemediationExceptionsRequest(input)
//   p := configservice.NewDescribeRemediationExceptionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeRemediationExceptionsPaginator(req DescribeRemediationExceptionsRequest) DescribeRemediationExceptionsPaginator {
	return DescribeRemediationExceptionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeRemediationExceptionsInput
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

// DescribeRemediationExceptionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeRemediationExceptionsPaginator struct {
	aws.Pager
}

func (p *DescribeRemediationExceptionsPaginator) CurrentPage() *DescribeRemediationExceptionsOutput {
	return p.Pager.CurrentPage().(*DescribeRemediationExceptionsOutput)
}

// DescribeRemediationExceptionsResponse is the response type for the
// DescribeRemediationExceptions API operation.
type DescribeRemediationExceptionsResponse struct {
	*DescribeRemediationExceptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRemediationExceptions request.
func (r *DescribeRemediationExceptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
