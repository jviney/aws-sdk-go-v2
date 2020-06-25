// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// Specifies the types of findings to include in a set of sample findings that
// Amazon Macie creates.
type CreateSampleFindingsInput struct {
	_ struct{} `type:"structure"`

	FindingTypes []FindingType `locationName:"findingTypes" type:"list"`
}

// String returns the string representation
func (s CreateSampleFindingsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateSampleFindingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FindingTypes != nil {
		v := s.FindingTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "findingTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

type CreateSampleFindingsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateSampleFindingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateSampleFindingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCreateSampleFindings = "CreateSampleFindings"

// CreateSampleFindingsRequest returns a request value for making API operation for
// Amazon Macie 2.
//
// Creates sample findings.
//
//    // Example sending a request using CreateSampleFindingsRequest.
//    req := client.CreateSampleFindingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie2-2020-01-01/CreateSampleFindings
func (c *Client) CreateSampleFindingsRequest(input *CreateSampleFindingsInput) CreateSampleFindingsRequest {
	op := &aws.Operation{
		Name:       opCreateSampleFindings,
		HTTPMethod: "POST",
		HTTPPath:   "/findings/sample",
	}

	if input == nil {
		input = &CreateSampleFindingsInput{}
	}

	req := c.newRequest(op, input, &CreateSampleFindingsOutput{})

	return CreateSampleFindingsRequest{Request: req, Input: input, Copy: c.CreateSampleFindingsRequest}
}

// CreateSampleFindingsRequest is the request type for the
// CreateSampleFindings API operation.
type CreateSampleFindingsRequest struct {
	*aws.Request
	Input *CreateSampleFindingsInput
	Copy  func(*CreateSampleFindingsInput) CreateSampleFindingsRequest
}

// Send marshals and sends the CreateSampleFindings API request.
func (r CreateSampleFindingsRequest) Send(ctx context.Context) (*CreateSampleFindingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSampleFindingsResponse{
		CreateSampleFindingsOutput: r.Request.Data.(*CreateSampleFindingsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSampleFindingsResponse is the response type for the
// CreateSampleFindings API operation.
type CreateSampleFindingsResponse struct {
	*CreateSampleFindingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSampleFindings request.
func (r *CreateSampleFindingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
