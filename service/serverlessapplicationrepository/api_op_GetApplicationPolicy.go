// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetApplicationPolicyInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"applicationId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetApplicationPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetApplicationPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetApplicationPolicyInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetApplicationPolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "applicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetApplicationPolicyOutput struct {
	_ struct{} `type:"structure"`

	Statements []ApplicationPolicyStatement `locationName:"statements" type:"list"`
}

// String returns the string representation
func (s GetApplicationPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetApplicationPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Statements != nil {
		v := s.Statements

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "statements", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetApplicationPolicy = "GetApplicationPolicy"

// GetApplicationPolicyRequest returns a request value for making API operation for
// AWSServerlessApplicationRepository.
//
// Retrieves the policy for the application.
//
//    // Example sending a request using GetApplicationPolicyRequest.
//    req := client.GetApplicationPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/GetApplicationPolicy
func (c *Client) GetApplicationPolicyRequest(input *GetApplicationPolicyInput) GetApplicationPolicyRequest {
	op := &aws.Operation{
		Name:       opGetApplicationPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/applications/{applicationId}/policy",
	}

	if input == nil {
		input = &GetApplicationPolicyInput{}
	}

	req := c.newRequest(op, input, &GetApplicationPolicyOutput{})

	return GetApplicationPolicyRequest{Request: req, Input: input, Copy: c.GetApplicationPolicyRequest}
}

// GetApplicationPolicyRequest is the request type for the
// GetApplicationPolicy API operation.
type GetApplicationPolicyRequest struct {
	*aws.Request
	Input *GetApplicationPolicyInput
	Copy  func(*GetApplicationPolicyInput) GetApplicationPolicyRequest
}

// Send marshals and sends the GetApplicationPolicy API request.
func (r GetApplicationPolicyRequest) Send(ctx context.Context) (*GetApplicationPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetApplicationPolicyResponse{
		GetApplicationPolicyOutput: r.Request.Data.(*GetApplicationPolicyOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetApplicationPolicyResponse is the response type for the
// GetApplicationPolicy API operation.
type GetApplicationPolicyResponse struct {
	*GetApplicationPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetApplicationPolicy request.
func (r *GetApplicationPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
