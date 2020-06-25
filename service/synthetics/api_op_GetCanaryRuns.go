// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package synthetics

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetCanaryRunsInput struct {
	_ struct{} `type:"structure"`

	// Specify this parameter to limit how many runs are returned each time you
	// use the GetCanaryRuns operation. If you omit this parameter, the default
	// of 100 is used.
	MaxResults *int64 `min:"1" type:"integer"`

	// The name of the canary that you want to see runs for.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`

	// A token that indicates that there is more data available. You can use this
	// token in a subsequent GetCanaryRuns operation to retrieve the next set of
	// results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetCanaryRunsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCanaryRunsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCanaryRunsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCanaryRunsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

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
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetCanaryRunsOutput struct {
	_ struct{} `type:"structure"`

	// An array of structures. Each structure contains the details of one of the
	// retrieved canary runs.
	CanaryRuns []CanaryRun `type:"list"`

	// A token that indicates that there is more data available. You can use this
	// token in a subsequent GetCanaryRuns operation to retrieve the next set of
	// results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetCanaryRunsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCanaryRunsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CanaryRuns != nil {
		v := s.CanaryRuns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "CanaryRuns", metadata)
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

const opGetCanaryRuns = "GetCanaryRuns"

// GetCanaryRunsRequest returns a request value for making API operation for
// Synthetics.
//
// Retrieves a list of runs for a specified canary.
//
//    // Example sending a request using GetCanaryRunsRequest.
//    req := client.GetCanaryRunsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/synthetics-2017-10-11/GetCanaryRuns
func (c *Client) GetCanaryRunsRequest(input *GetCanaryRunsInput) GetCanaryRunsRequest {
	op := &aws.Operation{
		Name:       opGetCanaryRuns,
		HTTPMethod: "POST",
		HTTPPath:   "/canary/{name}/runs",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetCanaryRunsInput{}
	}

	req := c.newRequest(op, input, &GetCanaryRunsOutput{})

	return GetCanaryRunsRequest{Request: req, Input: input, Copy: c.GetCanaryRunsRequest}
}

// GetCanaryRunsRequest is the request type for the
// GetCanaryRuns API operation.
type GetCanaryRunsRequest struct {
	*aws.Request
	Input *GetCanaryRunsInput
	Copy  func(*GetCanaryRunsInput) GetCanaryRunsRequest
}

// Send marshals and sends the GetCanaryRuns API request.
func (r GetCanaryRunsRequest) Send(ctx context.Context) (*GetCanaryRunsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCanaryRunsResponse{
		GetCanaryRunsOutput: r.Request.Data.(*GetCanaryRunsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetCanaryRunsRequestPaginator returns a paginator for GetCanaryRuns.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetCanaryRunsRequest(input)
//   p := synthetics.NewGetCanaryRunsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetCanaryRunsPaginator(req GetCanaryRunsRequest) GetCanaryRunsPaginator {
	return GetCanaryRunsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetCanaryRunsInput
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

// GetCanaryRunsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetCanaryRunsPaginator struct {
	aws.Pager
}

func (p *GetCanaryRunsPaginator) CurrentPage() *GetCanaryRunsOutput {
	return p.Pager.CurrentPage().(*GetCanaryRunsOutput)
}

// GetCanaryRunsResponse is the response type for the
// GetCanaryRuns API operation.
type GetCanaryRunsResponse struct {
	*GetCanaryRunsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCanaryRuns request.
func (r *GetCanaryRunsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
