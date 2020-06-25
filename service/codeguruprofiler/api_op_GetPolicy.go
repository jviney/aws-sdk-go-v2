// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codeguruprofiler

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// The structure representing the getPolicyRequest.
type GetPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name of the profiling group.
	//
	// ProfilingGroupName is a required field
	ProfilingGroupName *string `location:"uri" locationName:"profilingGroupName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPolicyInput"}

	if s.ProfilingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProfilingGroupName"))
	}
	if s.ProfilingGroupName != nil && len(*s.ProfilingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProfilingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ProfilingGroupName != nil {
		v := *s.ProfilingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "profilingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The structure representing the getPolicyResponse.
type GetPolicyOutput struct {
	_ struct{} `type:"structure"`

	// The resource-based policy attached to the ProfilingGroup.
	//
	// Policy is a required field
	Policy *string `locationName:"policy" type:"string" required:"true"`

	// A unique identifier for the current revision of the policy.
	//
	// RevisionId is a required field
	RevisionId *string `locationName:"revisionId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Policy != nil {
		v := *s.Policy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "policy", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "revisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetPolicy = "GetPolicy"

// GetPolicyRequest returns a request value for making API operation for
// Amazon CodeGuru Profiler.
//
// Gets the profiling group policy.
//
//    // Example sending a request using GetPolicyRequest.
//    req := client.GetPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codeguruprofiler-2019-07-18/GetPolicy
func (c *Client) GetPolicyRequest(input *GetPolicyInput) GetPolicyRequest {
	op := &aws.Operation{
		Name:       opGetPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/profilingGroups/{profilingGroupName}/policy",
	}

	if input == nil {
		input = &GetPolicyInput{}
	}

	req := c.newRequest(op, input, &GetPolicyOutput{})

	return GetPolicyRequest{Request: req, Input: input, Copy: c.GetPolicyRequest}
}

// GetPolicyRequest is the request type for the
// GetPolicy API operation.
type GetPolicyRequest struct {
	*aws.Request
	Input *GetPolicyInput
	Copy  func(*GetPolicyInput) GetPolicyRequest
}

// Send marshals and sends the GetPolicy API request.
func (r GetPolicyRequest) Send(ctx context.Context) (*GetPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPolicyResponse{
		GetPolicyOutput: r.Request.Data.(*GetPolicyOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPolicyResponse is the response type for the
// GetPolicy API operation.
type GetPolicyResponse struct {
	*GetPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPolicy request.
func (r *GetPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
