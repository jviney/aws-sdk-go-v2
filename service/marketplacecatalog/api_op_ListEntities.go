// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplacecatalog

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListEntitiesInput struct {
	_ struct{} `type:"structure"`

	// The catalog related to the request. Fixed value: AWSMarketplace
	//
	// Catalog is a required field
	Catalog *string `min:"1" type:"string" required:"true"`

	// The type of entities to retrieve.
	//
	// EntityType is a required field
	EntityType *string `min:"1" type:"string" required:"true"`

	// An array of filter objects. Each filter object contains two attributes, filterName
	// and filterValues.
	FilterList []Filter `min:"1" type:"list"`

	// Specifies the upper limit of the elements on a single page. If a value isn't
	// provided, the default value is 20.
	MaxResults *int64 `min:"1" type:"integer"`

	// The value of the next token, if it exists. Null if there are no more results.
	NextToken *string `min:"1" type:"string"`

	// An object that contains two attributes, SortBy and SortOrder.
	Sort *Sort `type:"structure"`
}

// String returns the string representation
func (s ListEntitiesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListEntitiesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListEntitiesInput"}

	if s.Catalog == nil {
		invalidParams.Add(aws.NewErrParamRequired("Catalog"))
	}
	if s.Catalog != nil && len(*s.Catalog) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Catalog", 1))
	}

	if s.EntityType == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntityType"))
	}
	if s.EntityType != nil && len(*s.EntityType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityType", 1))
	}
	if s.FilterList != nil && len(s.FilterList) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FilterList", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.FilterList != nil {
		for i, v := range s.FilterList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "FilterList", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Sort != nil {
		if err := s.Sort.Validate(); err != nil {
			invalidParams.AddNested("Sort", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListEntitiesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Catalog != nil {
		v := *s.Catalog

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Catalog", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EntityType != nil {
		v := *s.EntityType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EntityType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FilterList != nil {
		v := s.FilterList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "FilterList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Sort != nil {
		v := s.Sort

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Sort", v, metadata)
	}
	return nil
}

type ListEntitiesOutput struct {
	_ struct{} `type:"structure"`

	// Array of EntitySummary object.
	EntitySummaryList []EntitySummary `type:"list"`

	// The value of the next token if it exists. Null if there is no more result.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListEntitiesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListEntitiesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.EntitySummaryList != nil {
		v := s.EntitySummaryList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "EntitySummaryList", metadata)
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
	return nil
}

const opListEntities = "ListEntities"

// ListEntitiesRequest returns a request value for making API operation for
// AWS Marketplace Catalog Service.
//
// Provides the list of entities of a given type.
//
//    // Example sending a request using ListEntitiesRequest.
//    req := client.ListEntitiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/marketplace-catalog-2018-09-17/ListEntities
func (c *Client) ListEntitiesRequest(input *ListEntitiesInput) ListEntitiesRequest {
	op := &aws.Operation{
		Name:       opListEntities,
		HTTPMethod: "POST",
		HTTPPath:   "/ListEntities",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListEntitiesInput{}
	}

	req := c.newRequest(op, input, &ListEntitiesOutput{})

	return ListEntitiesRequest{Request: req, Input: input, Copy: c.ListEntitiesRequest}
}

// ListEntitiesRequest is the request type for the
// ListEntities API operation.
type ListEntitiesRequest struct {
	*aws.Request
	Input *ListEntitiesInput
	Copy  func(*ListEntitiesInput) ListEntitiesRequest
}

// Send marshals and sends the ListEntities API request.
func (r ListEntitiesRequest) Send(ctx context.Context) (*ListEntitiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListEntitiesResponse{
		ListEntitiesOutput: r.Request.Data.(*ListEntitiesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListEntitiesRequestPaginator returns a paginator for ListEntities.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListEntitiesRequest(input)
//   p := marketplacecatalog.NewListEntitiesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListEntitiesPaginator(req ListEntitiesRequest) ListEntitiesPaginator {
	return ListEntitiesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListEntitiesInput
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

// ListEntitiesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListEntitiesPaginator struct {
	aws.Pager
}

func (p *ListEntitiesPaginator) CurrentPage() *ListEntitiesOutput {
	return p.Pager.CurrentPage().(*ListEntitiesOutput)
}

// ListEntitiesResponse is the response type for the
// ListEntities API operation.
type ListEntitiesResponse struct {
	*ListEntitiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListEntities request.
func (r *ListEntitiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
