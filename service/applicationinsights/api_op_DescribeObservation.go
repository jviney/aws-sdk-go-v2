// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationinsights

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeObservationInput struct {
	_ struct{} `type:"structure"`

	// The ID of the observation.
	//
	// ObservationId is a required field
	ObservationId *string `min:"38" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeObservationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeObservationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeObservationInput"}

	if s.ObservationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObservationId"))
	}
	if s.ObservationId != nil && len(*s.ObservationId) < 38 {
		invalidParams.Add(aws.NewErrParamMinLen("ObservationId", 38))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeObservationOutput struct {
	_ struct{} `type:"structure"`

	// Information about the observation.
	Observation *Observation `type:"structure"`
}

// String returns the string representation
func (s DescribeObservationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeObservation = "DescribeObservation"

// DescribeObservationRequest returns a request value for making API operation for
// Amazon CloudWatch Application Insights.
//
// Describes an anomaly or error with the application.
//
//    // Example sending a request using DescribeObservationRequest.
//    req := client.DescribeObservationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-insights-2018-11-25/DescribeObservation
func (c *Client) DescribeObservationRequest(input *DescribeObservationInput) DescribeObservationRequest {
	op := &aws.Operation{
		Name:       opDescribeObservation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeObservationInput{}
	}

	req := c.newRequest(op, input, &DescribeObservationOutput{})

	return DescribeObservationRequest{Request: req, Input: input, Copy: c.DescribeObservationRequest}
}

// DescribeObservationRequest is the request type for the
// DescribeObservation API operation.
type DescribeObservationRequest struct {
	*aws.Request
	Input *DescribeObservationInput
	Copy  func(*DescribeObservationInput) DescribeObservationRequest
}

// Send marshals and sends the DescribeObservation API request.
func (r DescribeObservationRequest) Send(ctx context.Context) (*DescribeObservationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeObservationResponse{
		DescribeObservationOutput: r.Request.Data.(*DescribeObservationOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeObservationResponse is the response type for the
// DescribeObservation API operation.
type DescribeObservationResponse struct {
	*DescribeObservationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeObservation request.
func (r *DescribeObservationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
