// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListThreatIntelSetsInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the detector that the threatIntelSet is associated with.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`

	// You can use this parameter to indicate the maximum number of items that you
	// want in the response. The default value is 50. The maximum value is 50.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// You can use this parameter to paginate results in the response. Set the value
	// of this parameter to null on your first call to the list action. For subsequent
	// calls to the action, fill nextToken in the request with the value of NextToken
	// from the previous response to continue listing data.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListThreatIntelSetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListThreatIntelSetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListThreatIntelSetsInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
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
func (s ListThreatIntelSetsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListThreatIntelSetsOutput struct {
	_ struct{} `type:"structure"`

	// The pagination parameter to be used on the next list operation to retrieve
	// more items.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The IDs of the ThreatIntelSet resources.
	//
	// ThreatIntelSetIds is a required field
	ThreatIntelSetIds []string `locationName:"threatIntelSetIds" type:"list" required:"true"`
}

// String returns the string representation
func (s ListThreatIntelSetsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListThreatIntelSetsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThreatIntelSetIds != nil {
		v := s.ThreatIntelSetIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "threatIntelSetIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opListThreatIntelSets = "ListThreatIntelSets"

// ListThreatIntelSetsRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Lists the ThreatIntelSets of the GuardDuty service specified by the detector
// ID. If you use this operation from a member account, the ThreatIntelSets
// associated with the master account are returned.
//
//    // Example sending a request using ListThreatIntelSetsRequest.
//    req := client.ListThreatIntelSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/ListThreatIntelSets
func (c *Client) ListThreatIntelSetsRequest(input *ListThreatIntelSetsInput) ListThreatIntelSetsRequest {
	op := &aws.Operation{
		Name:       opListThreatIntelSets,
		HTTPMethod: "GET",
		HTTPPath:   "/detector/{detectorId}/threatintelset",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListThreatIntelSetsInput{}
	}

	req := c.newRequest(op, input, &ListThreatIntelSetsOutput{})

	return ListThreatIntelSetsRequest{Request: req, Input: input, Copy: c.ListThreatIntelSetsRequest}
}

// ListThreatIntelSetsRequest is the request type for the
// ListThreatIntelSets API operation.
type ListThreatIntelSetsRequest struct {
	*aws.Request
	Input *ListThreatIntelSetsInput
	Copy  func(*ListThreatIntelSetsInput) ListThreatIntelSetsRequest
}

// Send marshals and sends the ListThreatIntelSets API request.
func (r ListThreatIntelSetsRequest) Send(ctx context.Context) (*ListThreatIntelSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListThreatIntelSetsResponse{
		ListThreatIntelSetsOutput: r.Request.Data.(*ListThreatIntelSetsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListThreatIntelSetsRequestPaginator returns a paginator for ListThreatIntelSets.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListThreatIntelSetsRequest(input)
//   p := guardduty.NewListThreatIntelSetsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListThreatIntelSetsPaginator(req ListThreatIntelSetsRequest) ListThreatIntelSetsPaginator {
	return ListThreatIntelSetsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListThreatIntelSetsInput
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

// ListThreatIntelSetsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListThreatIntelSetsPaginator struct {
	aws.Pager
}

func (p *ListThreatIntelSetsPaginator) CurrentPage() *ListThreatIntelSetsOutput {
	return p.Pager.CurrentPage().(*ListThreatIntelSetsOutput)
}

// ListThreatIntelSetsResponse is the response type for the
// ListThreatIntelSets API operation.
type ListThreatIntelSetsResponse struct {
	*ListThreatIntelSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListThreatIntelSets request.
func (r *ListThreatIntelSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
