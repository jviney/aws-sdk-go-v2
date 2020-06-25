// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetStagesInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	MaxResults *string `location:"querystring" locationName:"maxResults" type:"string"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetStagesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetStagesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetStagesInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetStagesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetStagesOutput struct {
	_ struct{} `type:"structure"`

	Items []Stage `locationName:"items" type:"list"`

	// The next page of elements from this collection. Not valid for the last element
	// of the collection.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetStagesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetStagesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Items != nil {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "items", metadata)
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

const opGetStages = "GetStages"

// GetStagesRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Gets the Stages for an API.
//
//    // Example sending a request using GetStagesRequest.
//    req := client.GetStagesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetStages
func (c *Client) GetStagesRequest(input *GetStagesInput) GetStagesRequest {
	op := &aws.Operation{
		Name:       opGetStages,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/apis/{apiId}/stages",
	}

	if input == nil {
		input = &GetStagesInput{}
	}

	req := c.newRequest(op, input, &GetStagesOutput{})

	return GetStagesRequest{Request: req, Input: input, Copy: c.GetStagesRequest}
}

// GetStagesRequest is the request type for the
// GetStages API operation.
type GetStagesRequest struct {
	*aws.Request
	Input *GetStagesInput
	Copy  func(*GetStagesInput) GetStagesRequest
}

// Send marshals and sends the GetStages API request.
func (r GetStagesRequest) Send(ctx context.Context) (*GetStagesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetStagesResponse{
		GetStagesOutput: r.Request.Data.(*GetStagesOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetStagesResponse is the response type for the
// GetStages API operation.
type GetStagesResponse struct {
	*GetStagesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetStages request.
func (r *GetStagesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
