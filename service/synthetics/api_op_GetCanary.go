// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package synthetics

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetCanaryInput struct {
	_ struct{} `type:"structure"`

	// The name of the canary that you want details for.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetCanaryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCanaryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCanaryInput"}

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
func (s GetCanaryInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetCanaryOutput struct {
	_ struct{} `type:"structure"`

	// A strucure that contains the full information about the canary.
	Canary *Canary `type:"structure"`
}

// String returns the string representation
func (s GetCanaryOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCanaryOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Canary != nil {
		v := s.Canary

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Canary", v, metadata)
	}
	return nil
}

const opGetCanary = "GetCanary"

// GetCanaryRequest returns a request value for making API operation for
// Synthetics.
//
// Retrieves complete information about one canary. You must specify the name
// of the canary that you want. To get a list of canaries and their names, use
// DescribeCanaries (https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_DescribeCanaries.html).
//
//    // Example sending a request using GetCanaryRequest.
//    req := client.GetCanaryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/synthetics-2017-10-11/GetCanary
func (c *Client) GetCanaryRequest(input *GetCanaryInput) GetCanaryRequest {
	op := &aws.Operation{
		Name:       opGetCanary,
		HTTPMethod: "GET",
		HTTPPath:   "/canary/{name}",
	}

	if input == nil {
		input = &GetCanaryInput{}
	}

	req := c.newRequest(op, input, &GetCanaryOutput{})

	return GetCanaryRequest{Request: req, Input: input, Copy: c.GetCanaryRequest}
}

// GetCanaryRequest is the request type for the
// GetCanary API operation.
type GetCanaryRequest struct {
	*aws.Request
	Input *GetCanaryInput
	Copy  func(*GetCanaryInput) GetCanaryRequest
}

// Send marshals and sends the GetCanary API request.
func (r GetCanaryRequest) Send(ctx context.Context) (*GetCanaryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCanaryResponse{
		GetCanaryOutput: r.Request.Data.(*GetCanaryOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCanaryResponse is the response type for the
// GetCanary API operation.
type GetCanaryResponse struct {
	*GetCanaryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCanary request.
func (r *GetCanaryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
