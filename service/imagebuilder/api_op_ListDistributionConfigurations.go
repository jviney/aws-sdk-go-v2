// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListDistributionConfigurationsInput struct {
	_ struct{} `type:"structure"`

	// The filters.
	Filters []Filter `locationName:"filters" min:"1" type:"list"`

	// The maximum items to return in a request.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// A token to specify where to start paginating. This is the NextToken from
	// a previously truncated response.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListDistributionConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDistributionConfigurationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDistributionConfigurationsInput"}
	if s.Filters != nil && len(s.Filters) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Filters", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
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

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDistributionConfigurationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Filters != nil {
		v := s.Filters

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "filters", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListDistributionConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	// The list of distributions.
	DistributionConfigurationSummaryList []DistributionConfigurationSummary `locationName:"distributionConfigurationSummaryList" type:"list"`

	// The next token used for paginated responses. When this is not empty, there
	// are additional elements that the service has not included in this request.
	// Use this token with the next request to retrieve additional objects.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// The request ID that uniquely identifies this request.
	RequestId *string `locationName:"requestId" min:"1" type:"string"`
}

// String returns the string representation
func (s ListDistributionConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDistributionConfigurationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DistributionConfigurationSummaryList != nil {
		v := s.DistributionConfigurationSummaryList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "distributionConfigurationSummaryList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListDistributionConfigurations = "ListDistributionConfigurations"

// ListDistributionConfigurationsRequest returns a request value for making API operation for
// EC2 Image Builder.
//
// Returns a list of distribution configurations.
//
//    // Example sending a request using ListDistributionConfigurationsRequest.
//    req := client.ListDistributionConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/imagebuilder-2019-12-02/ListDistributionConfigurations
func (c *Client) ListDistributionConfigurationsRequest(input *ListDistributionConfigurationsInput) ListDistributionConfigurationsRequest {
	op := &aws.Operation{
		Name:       opListDistributionConfigurations,
		HTTPMethod: "POST",
		HTTPPath:   "/ListDistributionConfigurations",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDistributionConfigurationsInput{}
	}

	req := c.newRequest(op, input, &ListDistributionConfigurationsOutput{})

	return ListDistributionConfigurationsRequest{Request: req, Input: input, Copy: c.ListDistributionConfigurationsRequest}
}

// ListDistributionConfigurationsRequest is the request type for the
// ListDistributionConfigurations API operation.
type ListDistributionConfigurationsRequest struct {
	*aws.Request
	Input *ListDistributionConfigurationsInput
	Copy  func(*ListDistributionConfigurationsInput) ListDistributionConfigurationsRequest
}

// Send marshals and sends the ListDistributionConfigurations API request.
func (r ListDistributionConfigurationsRequest) Send(ctx context.Context) (*ListDistributionConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDistributionConfigurationsResponse{
		ListDistributionConfigurationsOutput: r.Request.Data.(*ListDistributionConfigurationsOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDistributionConfigurationsRequestPaginator returns a paginator for ListDistributionConfigurations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDistributionConfigurationsRequest(input)
//   p := imagebuilder.NewListDistributionConfigurationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDistributionConfigurationsPaginator(req ListDistributionConfigurationsRequest) ListDistributionConfigurationsPaginator {
	return ListDistributionConfigurationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDistributionConfigurationsInput
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

// ListDistributionConfigurationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDistributionConfigurationsPaginator struct {
	aws.Pager
}

func (p *ListDistributionConfigurationsPaginator) CurrentPage() *ListDistributionConfigurationsOutput {
	return p.Pager.CurrentPage().(*ListDistributionConfigurationsOutput)
}

// ListDistributionConfigurationsResponse is the response type for the
// ListDistributionConfigurations API operation.
type ListDistributionConfigurationsResponse struct {
	*ListDistributionConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDistributionConfigurations request.
func (r *ListDistributionConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
