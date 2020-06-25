// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListDataSetsInput struct {
	_ struct{} `type:"structure"`

	// The AWS account ID.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The maximum number of results to be returned per request.
	MaxResults *int64 `location:"querystring" locationName:"max-results" min:"1" type:"integer"`

	// The token for the next set of results, or null if there are no more results.
	NextToken *string `location:"querystring" locationName:"next-token" type:"string"`
}

// String returns the string representation
func (s ListDataSetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDataSetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDataSetsInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDataSetsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "max-results", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "next-token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListDataSetsOutput struct {
	_ struct{} `type:"structure"`

	// The list of dataset summaries.
	DataSetSummaries []DataSetSummary `type:"list"`

	// The token for the next set of results, or null if there are no more results.
	NextToken *string `type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s ListDataSetsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDataSetsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DataSetSummaries != nil {
		v := s.DataSetSummaries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "DataSetSummaries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
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
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opListDataSets = "ListDataSets"

// ListDataSetsRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Lists all of the datasets belonging to the current AWS account in an AWS
// Region.
//
// The permissions resource is arn:aws:quicksight:region:aws-account-id:dataset/*.
//
//    // Example sending a request using ListDataSetsRequest.
//    req := client.ListDataSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/ListDataSets
func (c *Client) ListDataSetsRequest(input *ListDataSetsInput) ListDataSetsRequest {
	op := &aws.Operation{
		Name:       opListDataSets,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/data-sets",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDataSetsInput{}
	}

	req := c.newRequest(op, input, &ListDataSetsOutput{})

	return ListDataSetsRequest{Request: req, Input: input, Copy: c.ListDataSetsRequest}
}

// ListDataSetsRequest is the request type for the
// ListDataSets API operation.
type ListDataSetsRequest struct {
	*aws.Request
	Input *ListDataSetsInput
	Copy  func(*ListDataSetsInput) ListDataSetsRequest
}

// Send marshals and sends the ListDataSets API request.
func (r ListDataSetsRequest) Send(ctx context.Context) (*ListDataSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDataSetsResponse{
		ListDataSetsOutput: r.Request.Data.(*ListDataSetsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDataSetsRequestPaginator returns a paginator for ListDataSets.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDataSetsRequest(input)
//   p := quicksight.NewListDataSetsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDataSetsPaginator(req ListDataSetsRequest) ListDataSetsPaginator {
	return ListDataSetsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDataSetsInput
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

// ListDataSetsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDataSetsPaginator struct {
	aws.Pager
}

func (p *ListDataSetsPaginator) CurrentPage() *ListDataSetsOutput {
	return p.Pager.CurrentPage().(*ListDataSetsOutput)
}

// ListDataSetsResponse is the response type for the
// ListDataSets API operation.
type ListDataSetsResponse struct {
	*ListDataSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDataSets request.
func (r *ListDataSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
