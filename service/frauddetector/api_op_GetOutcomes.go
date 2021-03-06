// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package frauddetector

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetOutcomesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of objects to return for the request.
	MaxResults *int64 `locationName:"maxResults" min:"50" type:"integer"`

	// The name of the outcome or outcomes to get.
	Name *string `locationName:"name" min:"1" type:"string"`

	// The next page token for the request.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetOutcomesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetOutcomesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetOutcomesInput"}
	if s.MaxResults != nil && *s.MaxResults < 50 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 50))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetOutcomesOutput struct {
	_ struct{} `type:"structure"`

	// The next page token for subsequent requests.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The outcomes.
	Outcomes []Outcome `locationName:"outcomes" type:"list"`
}

// String returns the string representation
func (s GetOutcomesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetOutcomes = "GetOutcomes"

// GetOutcomesRequest returns a request value for making API operation for
// Amazon Fraud Detector.
//
// Gets one or more outcomes. This is a paginated API. If you provide a null
// maxSizePerPage, this actions retrieves a maximum of 10 records per page.
// If you provide a maxSizePerPage, the value must be between 50 and 100. To
// get the next page results, provide the pagination token from the GetOutcomesResult
// as part of your request. A null pagination token fetches the records from
// the beginning.
//
//    // Example sending a request using GetOutcomesRequest.
//    req := client.GetOutcomesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/frauddetector-2019-11-15/GetOutcomes
func (c *Client) GetOutcomesRequest(input *GetOutcomesInput) GetOutcomesRequest {
	op := &aws.Operation{
		Name:       opGetOutcomes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetOutcomesInput{}
	}

	req := c.newRequest(op, input, &GetOutcomesOutput{})

	return GetOutcomesRequest{Request: req, Input: input, Copy: c.GetOutcomesRequest}
}

// GetOutcomesRequest is the request type for the
// GetOutcomes API operation.
type GetOutcomesRequest struct {
	*aws.Request
	Input *GetOutcomesInput
	Copy  func(*GetOutcomesInput) GetOutcomesRequest
}

// Send marshals and sends the GetOutcomes API request.
func (r GetOutcomesRequest) Send(ctx context.Context) (*GetOutcomesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetOutcomesResponse{
		GetOutcomesOutput: r.Request.Data.(*GetOutcomesOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetOutcomesRequestPaginator returns a paginator for GetOutcomes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetOutcomesRequest(input)
//   p := frauddetector.NewGetOutcomesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetOutcomesPaginator(req GetOutcomesRequest) GetOutcomesPaginator {
	return GetOutcomesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetOutcomesInput
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

// GetOutcomesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetOutcomesPaginator struct {
	aws.Pager
}

func (p *GetOutcomesPaginator) CurrentPage() *GetOutcomesOutput {
	return p.Pager.CurrentPage().(*GetOutcomesOutput)
}

// GetOutcomesResponse is the response type for the
// GetOutcomes API operation.
type GetOutcomesResponse struct {
	*GetOutcomesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetOutcomes request.
func (r *GetOutcomesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
