// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListClustersInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of cluster results returned by ListClusters in paginated
	// output. When you use this parameter, ListClusters returns only maxResults
	// results in a single page along with a nextToken response element. You can
	// see the remaining results of the initial request by sending another ListClusters
	// request with the returned nextToken value. This value can be between 1 and
	// 100. If you don't use this parameter, ListClusters returns up to 100 results
	// and a nextToken value if applicable.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The nextToken value returned from a previous paginated ListClusters request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value.
	//
	// This token should be treated as an opaque identifier that is used only to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListClustersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListClustersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListClustersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListClustersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

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

type ListClustersOutput struct {
	_ struct{} `type:"structure"`

	// A list of all of the clusters for your account in the specified Region.
	Clusters []string `locationName:"clusters" type:"list"`

	// The nextToken value to include in a future ListClusters request. When the
	// results of a ListClusters request exceed maxResults, you can use this value
	// to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListClustersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListClustersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Clusters != nil {
		v := s.Clusters

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "clusters", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListClusters = "ListClusters"

// ListClustersRequest returns a request value for making API operation for
// Amazon Elastic Kubernetes Service.
//
// Lists the Amazon EKS clusters in your AWS account in the specified Region.
//
//    // Example sending a request using ListClustersRequest.
//    req := client.ListClustersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/ListClusters
func (c *Client) ListClustersRequest(input *ListClustersInput) ListClustersRequest {
	op := &aws.Operation{
		Name:       opListClusters,
		HTTPMethod: "GET",
		HTTPPath:   "/clusters",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListClustersInput{}
	}

	req := c.newRequest(op, input, &ListClustersOutput{})

	return ListClustersRequest{Request: req, Input: input, Copy: c.ListClustersRequest}
}

// ListClustersRequest is the request type for the
// ListClusters API operation.
type ListClustersRequest struct {
	*aws.Request
	Input *ListClustersInput
	Copy  func(*ListClustersInput) ListClustersRequest
}

// Send marshals and sends the ListClusters API request.
func (r ListClustersRequest) Send(ctx context.Context) (*ListClustersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListClustersResponse{
		ListClustersOutput: r.Request.Data.(*ListClustersOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListClustersRequestPaginator returns a paginator for ListClusters.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListClustersRequest(input)
//   p := eks.NewListClustersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListClustersPaginator(req ListClustersRequest) ListClustersPaginator {
	return ListClustersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListClustersInput
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

// ListClustersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListClustersPaginator struct {
	aws.Pager
}

func (p *ListClustersPaginator) CurrentPage() *ListClustersOutput {
	return p.Pager.CurrentPage().(*ListClustersOutput)
}

// ListClustersResponse is the response type for the
// ListClusters API operation.
type ListClustersResponse struct {
	*ListClustersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListClusters request.
func (r *ListClustersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
