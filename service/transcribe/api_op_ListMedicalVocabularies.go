// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transcribe

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListMedicalVocabulariesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of vocabularies to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// Returns vocabularies in the list whose name contains the specified string.
	// The search is case-insensitive, ListMedicalVocabularies returns both "vocabularyname"
	// and "VocabularyName" in the response list.
	NameContains *string `min:"1" type:"string"`

	// If the result of your previous request to ListMedicalVocabularies was truncated,
	// include the NextToken to fetch the next set of jobs.
	NextToken *string `type:"string"`

	// When specified, only returns vocabularies with the VocabularyState equal
	// to the specified vocabulary state.
	StateEquals VocabularyState `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListMedicalVocabulariesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListMedicalVocabulariesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListMedicalVocabulariesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NameContains != nil && len(*s.NameContains) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NameContains", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListMedicalVocabulariesOutput struct {
	_ struct{} `type:"structure"`

	// The ListMedicalVocabularies operation returns a page of vocabularies at a
	// time. The maximum size of the page is set by the MaxResults parameter. If
	// there are more jobs in the list than the page size, Amazon Transcribe Medical
	// returns the NextPage token. Include the token in the next request to the
	// ListMedicalVocabularies operation to return the next page of jobs.
	NextToken *string `type:"string"`

	// The requested vocabulary state.
	Status VocabularyState `type:"string" enum:"true"`

	// A list of objects that describe the vocabularies that match the search criteria
	// in the request.
	Vocabularies []VocabularyInfo `type:"list"`
}

// String returns the string representation
func (s ListMedicalVocabulariesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListMedicalVocabularies = "ListMedicalVocabularies"

// ListMedicalVocabulariesRequest returns a request value for making API operation for
// Amazon Transcribe Service.
//
// Returns a list of vocabularies that match the specified criteria. You get
// the entire list of vocabularies if you don't enter a value in any of the
// request parameters.
//
//    // Example sending a request using ListMedicalVocabulariesRequest.
//    req := client.ListMedicalVocabulariesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/ListMedicalVocabularies
func (c *Client) ListMedicalVocabulariesRequest(input *ListMedicalVocabulariesInput) ListMedicalVocabulariesRequest {
	op := &aws.Operation{
		Name:       opListMedicalVocabularies,
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
		input = &ListMedicalVocabulariesInput{}
	}

	req := c.newRequest(op, input, &ListMedicalVocabulariesOutput{})

	return ListMedicalVocabulariesRequest{Request: req, Input: input, Copy: c.ListMedicalVocabulariesRequest}
}

// ListMedicalVocabulariesRequest is the request type for the
// ListMedicalVocabularies API operation.
type ListMedicalVocabulariesRequest struct {
	*aws.Request
	Input *ListMedicalVocabulariesInput
	Copy  func(*ListMedicalVocabulariesInput) ListMedicalVocabulariesRequest
}

// Send marshals and sends the ListMedicalVocabularies API request.
func (r ListMedicalVocabulariesRequest) Send(ctx context.Context) (*ListMedicalVocabulariesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListMedicalVocabulariesResponse{
		ListMedicalVocabulariesOutput: r.Request.Data.(*ListMedicalVocabulariesOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListMedicalVocabulariesRequestPaginator returns a paginator for ListMedicalVocabularies.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListMedicalVocabulariesRequest(input)
//   p := transcribe.NewListMedicalVocabulariesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListMedicalVocabulariesPaginator(req ListMedicalVocabulariesRequest) ListMedicalVocabulariesPaginator {
	return ListMedicalVocabulariesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListMedicalVocabulariesInput
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

// ListMedicalVocabulariesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListMedicalVocabulariesPaginator struct {
	aws.Pager
}

func (p *ListMedicalVocabulariesPaginator) CurrentPage() *ListMedicalVocabulariesOutput {
	return p.Pager.CurrentPage().(*ListMedicalVocabulariesOutput)
}

// ListMedicalVocabulariesResponse is the response type for the
// ListMedicalVocabularies API operation.
type ListMedicalVocabulariesResponse struct {
	*ListMedicalVocabulariesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListMedicalVocabularies request.
func (r *ListMedicalVocabulariesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
