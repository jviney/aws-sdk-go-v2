// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListFunctionsInput struct {
	_ struct{} `type:"structure"`

	// The GraphQL API ID.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// The maximum number of results you want the request to return.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `location:"querystring" locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListFunctionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFunctionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFunctionsInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFunctionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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

type ListFunctionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of Function objects.
	Functions []FunctionConfiguration `locationName:"functions" type:"list"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListFunctionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFunctionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Functions != nil {
		v := s.Functions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "functions", metadata)
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
	return nil
}

const opListFunctions = "ListFunctions"

// ListFunctionsRequest returns a request value for making API operation for
// AWS AppSync.
//
// List multiple functions.
//
//    // Example sending a request using ListFunctionsRequest.
//    req := client.ListFunctionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/ListFunctions
func (c *Client) ListFunctionsRequest(input *ListFunctionsInput) ListFunctionsRequest {
	op := &aws.Operation{
		Name:       opListFunctions,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apis/{apiId}/functions",
	}

	if input == nil {
		input = &ListFunctionsInput{}
	}

	req := c.newRequest(op, input, &ListFunctionsOutput{})

	return ListFunctionsRequest{Request: req, Input: input, Copy: c.ListFunctionsRequest}
}

// ListFunctionsRequest is the request type for the
// ListFunctions API operation.
type ListFunctionsRequest struct {
	*aws.Request
	Input *ListFunctionsInput
	Copy  func(*ListFunctionsInput) ListFunctionsRequest
}

// Send marshals and sends the ListFunctions API request.
func (r ListFunctionsRequest) Send(ctx context.Context) (*ListFunctionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFunctionsResponse{
		ListFunctionsOutput: r.Request.Data.(*ListFunctionsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListFunctionsResponse is the response type for the
// ListFunctions API operation.
type ListFunctionsResponse struct {
	*ListFunctionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFunctions request.
func (r *ListFunctionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
