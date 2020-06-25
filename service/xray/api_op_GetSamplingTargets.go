// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package xray

import (
	"context"
	"fmt"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetSamplingTargetsInput struct {
	_ struct{} `type:"structure"`

	// Information about rules that the service is using to sample requests.
	//
	// SamplingStatisticsDocuments is a required field
	SamplingStatisticsDocuments []SamplingStatisticsDocument `type:"list" required:"true"`
}

// String returns the string representation
func (s GetSamplingTargetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSamplingTargetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSamplingTargetsInput"}

	if s.SamplingStatisticsDocuments == nil {
		invalidParams.Add(aws.NewErrParamRequired("SamplingStatisticsDocuments"))
	}
	if s.SamplingStatisticsDocuments != nil {
		for i, v := range s.SamplingStatisticsDocuments {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SamplingStatisticsDocuments", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSamplingTargetsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.SamplingStatisticsDocuments != nil {
		v := s.SamplingStatisticsDocuments

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "SamplingStatisticsDocuments", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

type GetSamplingTargetsOutput struct {
	_ struct{} `type:"structure"`

	// The last time a user changed the sampling rule configuration. If the sampling
	// rule configuration changed since the service last retrieved it, the service
	// should call GetSamplingRules to get the latest version.
	LastRuleModification *time.Time `type:"timestamp"`

	// Updated rules that the service should use to sample requests.
	SamplingTargetDocuments []SamplingTargetDocument `type:"list"`

	// Information about SamplingStatisticsDocument that X-Ray could not process.
	UnprocessedStatistics []UnprocessedStatistics `type:"list"`
}

// String returns the string representation
func (s GetSamplingTargetsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSamplingTargetsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.LastRuleModification != nil {
		v := *s.LastRuleModification

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastRuleModification",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.SamplingTargetDocuments != nil {
		v := s.SamplingTargetDocuments

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "SamplingTargetDocuments", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.UnprocessedStatistics != nil {
		v := s.UnprocessedStatistics

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "UnprocessedStatistics", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetSamplingTargets = "GetSamplingTargets"

// GetSamplingTargetsRequest returns a request value for making API operation for
// AWS X-Ray.
//
// Requests a sampling quota for rules that the service is using to sample requests.
//
//    // Example sending a request using GetSamplingTargetsRequest.
//    req := client.GetSamplingTargetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/xray-2016-04-12/GetSamplingTargets
func (c *Client) GetSamplingTargetsRequest(input *GetSamplingTargetsInput) GetSamplingTargetsRequest {
	op := &aws.Operation{
		Name:       opGetSamplingTargets,
		HTTPMethod: "POST",
		HTTPPath:   "/SamplingTargets",
	}

	if input == nil {
		input = &GetSamplingTargetsInput{}
	}

	req := c.newRequest(op, input, &GetSamplingTargetsOutput{})

	return GetSamplingTargetsRequest{Request: req, Input: input, Copy: c.GetSamplingTargetsRequest}
}

// GetSamplingTargetsRequest is the request type for the
// GetSamplingTargets API operation.
type GetSamplingTargetsRequest struct {
	*aws.Request
	Input *GetSamplingTargetsInput
	Copy  func(*GetSamplingTargetsInput) GetSamplingTargetsRequest
}

// Send marshals and sends the GetSamplingTargets API request.
func (r GetSamplingTargetsRequest) Send(ctx context.Context) (*GetSamplingTargetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSamplingTargetsResponse{
		GetSamplingTargetsOutput: r.Request.Data.(*GetSamplingTargetsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSamplingTargetsResponse is the response type for the
// GetSamplingTargets API operation.
type GetSamplingTargetsResponse struct {
	*GetSamplingTargetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSamplingTargets request.
func (r *GetSamplingTargetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
