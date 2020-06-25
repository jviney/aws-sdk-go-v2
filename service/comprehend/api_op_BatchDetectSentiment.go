// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type BatchDetectSentimentInput struct {
	_ struct{} `type:"structure"`

	// The language of the input documents. You can specify any of the primary languages
	// supported by Amazon Comprehend. All documents must be in the same language.
	//
	// LanguageCode is a required field
	LanguageCode LanguageCode `type:"string" required:"true" enum:"true"`

	// A list containing the text of the input documents. The list can contain a
	// maximum of 25 documents. Each document must contain fewer that 5,000 bytes
	// of UTF-8 encoded characters.
	//
	// TextList is a required field
	TextList []string `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDetectSentimentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchDetectSentimentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchDetectSentimentInput"}
	if len(s.LanguageCode) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("LanguageCode"))
	}

	if s.TextList == nil {
		invalidParams.Add(aws.NewErrParamRequired("TextList"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchDetectSentimentOutput struct {
	_ struct{} `type:"structure"`

	// A list containing one object for each document that contained an error. The
	// results are sorted in ascending order by the Index field and match the order
	// of the documents in the input list. If there are no errors in the batch,
	// the ErrorList is empty.
	//
	// ErrorList is a required field
	ErrorList []BatchItemError `type:"list" required:"true"`

	// A list of objects containing the results of the operation. The results are
	// sorted in ascending order by the Index field and match the order of the documents
	// in the input list. If all of the documents contain an error, the ResultList
	// is empty.
	//
	// ResultList is a required field
	ResultList []BatchDetectSentimentItemResult `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDetectSentimentOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchDetectSentiment = "BatchDetectSentiment"

// BatchDetectSentimentRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Inspects a batch of documents and returns an inference of the prevailing
// sentiment, POSITIVE, NEUTRAL, MIXED, or NEGATIVE, in each one.
//
//    // Example sending a request using BatchDetectSentimentRequest.
//    req := client.BatchDetectSentimentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/BatchDetectSentiment
func (c *Client) BatchDetectSentimentRequest(input *BatchDetectSentimentInput) BatchDetectSentimentRequest {
	op := &aws.Operation{
		Name:       opBatchDetectSentiment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchDetectSentimentInput{}
	}

	req := c.newRequest(op, input, &BatchDetectSentimentOutput{})

	return BatchDetectSentimentRequest{Request: req, Input: input, Copy: c.BatchDetectSentimentRequest}
}

// BatchDetectSentimentRequest is the request type for the
// BatchDetectSentiment API operation.
type BatchDetectSentimentRequest struct {
	*aws.Request
	Input *BatchDetectSentimentInput
	Copy  func(*BatchDetectSentimentInput) BatchDetectSentimentRequest
}

// Send marshals and sends the BatchDetectSentiment API request.
func (r BatchDetectSentimentRequest) Send(ctx context.Context) (*BatchDetectSentimentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDetectSentimentResponse{
		BatchDetectSentimentOutput: r.Request.Data.(*BatchDetectSentimentOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDetectSentimentResponse is the response type for the
// BatchDetectSentiment API operation.
type BatchDetectSentimentResponse struct {
	*BatchDetectSentimentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDetectSentiment request.
func (r *BatchDetectSentimentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
