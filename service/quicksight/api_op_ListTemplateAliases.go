// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListTemplateAliasesInput struct {
	_ struct{} `type:"structure"`

	// The ID of the AWS account that contains the template aliases that you're
	// listing.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The maximum number of results to be returned per request.
	MaxResults *int64 `location:"querystring" locationName:"max-result" min:"1" type:"integer"`

	// The token for the next set of results, or null if there are no more results.
	NextToken *string `location:"querystring" locationName:"next-token" type:"string"`

	// The ID for the template.
	//
	// TemplateId is a required field
	TemplateId *string `location:"uri" locationName:"TemplateId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListTemplateAliasesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTemplateAliasesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTemplateAliasesInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.TemplateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateId"))
	}
	if s.TemplateId != nil && len(*s.TemplateId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTemplateAliasesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateId != nil {
		v := *s.TemplateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "TemplateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "max-result", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "next-token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListTemplateAliasesOutput struct {
	_ struct{} `type:"structure"`

	// The token for the next set of results, or null if there are no more results.
	NextToken *string `type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`

	// A structure containing the list of the template's aliases.
	TemplateAliasList []TemplateAlias `type:"list"`
}

// String returns the string representation
func (s ListTemplateAliasesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTemplateAliasesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateAliasList != nil {
		v := s.TemplateAliasList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "TemplateAliasList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opListTemplateAliases = "ListTemplateAliases"

// ListTemplateAliasesRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Lists all the aliases of a template.
//
//    // Example sending a request using ListTemplateAliasesRequest.
//    req := client.ListTemplateAliasesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/ListTemplateAliases
func (c *Client) ListTemplateAliasesRequest(input *ListTemplateAliasesInput) ListTemplateAliasesRequest {
	op := &aws.Operation{
		Name:       opListTemplateAliases,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/templates/{TemplateId}/aliases",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTemplateAliasesInput{}
	}

	req := c.newRequest(op, input, &ListTemplateAliasesOutput{})

	return ListTemplateAliasesRequest{Request: req, Input: input, Copy: c.ListTemplateAliasesRequest}
}

// ListTemplateAliasesRequest is the request type for the
// ListTemplateAliases API operation.
type ListTemplateAliasesRequest struct {
	*aws.Request
	Input *ListTemplateAliasesInput
	Copy  func(*ListTemplateAliasesInput) ListTemplateAliasesRequest
}

// Send marshals and sends the ListTemplateAliases API request.
func (r ListTemplateAliasesRequest) Send(ctx context.Context) (*ListTemplateAliasesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTemplateAliasesResponse{
		ListTemplateAliasesOutput: r.Request.Data.(*ListTemplateAliasesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTemplateAliasesRequestPaginator returns a paginator for ListTemplateAliases.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTemplateAliasesRequest(input)
//   p := quicksight.NewListTemplateAliasesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTemplateAliasesPaginator(req ListTemplateAliasesRequest) ListTemplateAliasesPaginator {
	return ListTemplateAliasesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListTemplateAliasesInput
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

// ListTemplateAliasesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTemplateAliasesPaginator struct {
	aws.Pager
}

func (p *ListTemplateAliasesPaginator) CurrentPage() *ListTemplateAliasesOutput {
	return p.Pager.CurrentPage().(*ListTemplateAliasesOutput)
}

// ListTemplateAliasesResponse is the response type for the
// ListTemplateAliases API operation.
type ListTemplateAliasesResponse struct {
	*ListTemplateAliasesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTemplateAliases request.
func (r *ListTemplateAliasesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
