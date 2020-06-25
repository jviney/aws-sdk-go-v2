// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package polly

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetLexiconInput struct {
	_ struct{} `type:"structure"`

	// Name of the lexicon.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"LexiconName" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s GetLexiconInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLexiconInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLexiconInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetLexiconInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "LexiconName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetLexiconOutput struct {
	_ struct{} `type:"structure"`

	// Lexicon object that provides name and the string content of the lexicon.
	Lexicon *Lexicon `type:"structure"`

	// Metadata of the lexicon, including phonetic alphabetic used, language code,
	// lexicon ARN, number of lexemes defined in the lexicon, and size of lexicon
	// in bytes.
	LexiconAttributes *LexiconAttributes `type:"structure"`
}

// String returns the string representation
func (s GetLexiconOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetLexiconOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Lexicon != nil {
		v := s.Lexicon

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Lexicon", v, metadata)
	}
	if s.LexiconAttributes != nil {
		v := s.LexiconAttributes

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "LexiconAttributes", v, metadata)
	}
	return nil
}

const opGetLexicon = "GetLexicon"

// GetLexiconRequest returns a request value for making API operation for
// Amazon Polly.
//
// Returns the content of the specified pronunciation lexicon stored in an AWS
// Region. For more information, see Managing Lexicons (https://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html).
//
//    // Example sending a request using GetLexiconRequest.
//    req := client.GetLexiconRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/GetLexicon
func (c *Client) GetLexiconRequest(input *GetLexiconInput) GetLexiconRequest {
	op := &aws.Operation{
		Name:       opGetLexicon,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/lexicons/{LexiconName}",
	}

	if input == nil {
		input = &GetLexiconInput{}
	}

	req := c.newRequest(op, input, &GetLexiconOutput{})

	return GetLexiconRequest{Request: req, Input: input, Copy: c.GetLexiconRequest}
}

// GetLexiconRequest is the request type for the
// GetLexicon API operation.
type GetLexiconRequest struct {
	*aws.Request
	Input *GetLexiconInput
	Copy  func(*GetLexiconInput) GetLexiconRequest
}

// Send marshals and sends the GetLexicon API request.
func (r GetLexiconRequest) Send(ctx context.Context) (*GetLexiconResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLexiconResponse{
		GetLexiconOutput: r.Request.Data.(*GetLexiconOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLexiconResponse is the response type for the
// GetLexicon API operation.
type GetLexiconResponse struct {
	*GetLexiconOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLexicon request.
func (r *GetLexiconResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
