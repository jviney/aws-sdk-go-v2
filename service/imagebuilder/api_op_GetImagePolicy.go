// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetImagePolicyInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the image whose policy you want to retrieve.
	//
	// ImageArn is a required field
	ImageArn *string `location:"querystring" locationName:"imageArn" type:"string" required:"true"`
}

// String returns the string representation
func (s GetImagePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetImagePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetImagePolicyInput"}

	if s.ImageArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImageArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetImagePolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ImageArn != nil {
		v := *s.ImageArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "imageArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetImagePolicyOutput struct {
	_ struct{} `type:"structure"`

	// The image policy object.
	Policy *string `locationName:"policy" min:"1" type:"string"`

	// The request ID that uniquely identifies this request.
	RequestId *string `locationName:"requestId" min:"1" type:"string"`
}

// String returns the string representation
func (s GetImagePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetImagePolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Policy != nil {
		v := *s.Policy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "policy", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetImagePolicy = "GetImagePolicy"

// GetImagePolicyRequest returns a request value for making API operation for
// EC2 Image Builder.
//
// Gets an image policy.
//
//    // Example sending a request using GetImagePolicyRequest.
//    req := client.GetImagePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/imagebuilder-2019-12-02/GetImagePolicy
func (c *Client) GetImagePolicyRequest(input *GetImagePolicyInput) GetImagePolicyRequest {
	op := &aws.Operation{
		Name:       opGetImagePolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/GetImagePolicy",
	}

	if input == nil {
		input = &GetImagePolicyInput{}
	}

	req := c.newRequest(op, input, &GetImagePolicyOutput{})

	return GetImagePolicyRequest{Request: req, Input: input, Copy: c.GetImagePolicyRequest}
}

// GetImagePolicyRequest is the request type for the
// GetImagePolicy API operation.
type GetImagePolicyRequest struct {
	*aws.Request
	Input *GetImagePolicyInput
	Copy  func(*GetImagePolicyInput) GetImagePolicyRequest
}

// Send marshals and sends the GetImagePolicy API request.
func (r GetImagePolicyRequest) Send(ctx context.Context) (*GetImagePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetImagePolicyResponse{
		GetImagePolicyOutput: r.Request.Data.(*GetImagePolicyOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetImagePolicyResponse is the response type for the
// GetImagePolicy API operation.
type GetImagePolicyResponse struct {
	*GetImagePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetImagePolicy request.
func (r *GetImagePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
