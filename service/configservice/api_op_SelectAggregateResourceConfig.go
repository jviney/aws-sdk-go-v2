// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type SelectAggregateResourceConfigInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration aggregator.
	//
	// ConfigurationAggregatorName is a required field
	ConfigurationAggregatorName *string `min:"1" type:"string" required:"true"`

	// The SQL query SELECT command.
	//
	// Expression is a required field
	Expression *string `min:"1" type:"string" required:"true"`

	// The maximum number of query results returned on each page.
	Limit *int64 `type:"integer"`

	MaxResults *int64 `type:"integer"`

	// The nextToken string returned in a previous request that you use to request
	// the next page of results in a paginated response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s SelectAggregateResourceConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SelectAggregateResourceConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SelectAggregateResourceConfigInput"}

	if s.ConfigurationAggregatorName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationAggregatorName"))
	}
	if s.ConfigurationAggregatorName != nil && len(*s.ConfigurationAggregatorName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConfigurationAggregatorName", 1))
	}

	if s.Expression == nil {
		invalidParams.Add(aws.NewErrParamRequired("Expression"))
	}
	if s.Expression != nil && len(*s.Expression) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Expression", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SelectAggregateResourceConfigOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken string returned in a previous request that you use to request
	// the next page of results in a paginated response.
	NextToken *string `type:"string"`

	// Details about the query.
	QueryInfo *QueryInfo `type:"structure"`

	// Returns the results for the SQL query.
	Results []string `type:"list"`
}

// String returns the string representation
func (s SelectAggregateResourceConfigOutput) String() string {
	return awsutil.Prettify(s)
}

const opSelectAggregateResourceConfig = "SelectAggregateResourceConfig"

// SelectAggregateResourceConfigRequest returns a request value for making API operation for
// AWS Config.
//
// Accepts a structured query language (SQL) SELECT command and an aggregator
// to query configuration state of AWS resources across multiple accounts and
// regions, performs the corresponding search, and returns resource configurations
// matching the properties.
//
// For more information about query components, see the Query Components (https://docs.aws.amazon.com/config/latest/developerguide/query-components.html)
// section in the AWS Config Developer Guide.
//
//    // Example sending a request using SelectAggregateResourceConfigRequest.
//    req := client.SelectAggregateResourceConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/SelectAggregateResourceConfig
func (c *Client) SelectAggregateResourceConfigRequest(input *SelectAggregateResourceConfigInput) SelectAggregateResourceConfigRequest {
	op := &aws.Operation{
		Name:       opSelectAggregateResourceConfig,
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
		input = &SelectAggregateResourceConfigInput{}
	}

	req := c.newRequest(op, input, &SelectAggregateResourceConfigOutput{})

	return SelectAggregateResourceConfigRequest{Request: req, Input: input, Copy: c.SelectAggregateResourceConfigRequest}
}

// SelectAggregateResourceConfigRequest is the request type for the
// SelectAggregateResourceConfig API operation.
type SelectAggregateResourceConfigRequest struct {
	*aws.Request
	Input *SelectAggregateResourceConfigInput
	Copy  func(*SelectAggregateResourceConfigInput) SelectAggregateResourceConfigRequest
}

// Send marshals and sends the SelectAggregateResourceConfig API request.
func (r SelectAggregateResourceConfigRequest) Send(ctx context.Context) (*SelectAggregateResourceConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SelectAggregateResourceConfigResponse{
		SelectAggregateResourceConfigOutput: r.Request.Data.(*SelectAggregateResourceConfigOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewSelectAggregateResourceConfigRequestPaginator returns a paginator for SelectAggregateResourceConfig.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.SelectAggregateResourceConfigRequest(input)
//   p := configservice.NewSelectAggregateResourceConfigRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewSelectAggregateResourceConfigPaginator(req SelectAggregateResourceConfigRequest) SelectAggregateResourceConfigPaginator {
	return SelectAggregateResourceConfigPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *SelectAggregateResourceConfigInput
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

// SelectAggregateResourceConfigPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type SelectAggregateResourceConfigPaginator struct {
	aws.Pager
}

func (p *SelectAggregateResourceConfigPaginator) CurrentPage() *SelectAggregateResourceConfigOutput {
	return p.Pager.CurrentPage().(*SelectAggregateResourceConfigOutput)
}

// SelectAggregateResourceConfigResponse is the response type for the
// SelectAggregateResourceConfig API operation.
type SelectAggregateResourceConfigResponse struct {
	*SelectAggregateResourceConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SelectAggregateResourceConfig request.
func (r *SelectAggregateResourceConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
