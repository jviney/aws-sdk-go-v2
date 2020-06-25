// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// A request for the reason that a health check failed most recently.
type GetHealthCheckLastFailureReasonInput struct {
	_ struct{} `type:"structure"`

	// The ID for the health check for which you want the last failure reason. When
	// you created the health check, CreateHealthCheck returned the ID in the response,
	// in the HealthCheckId element.
	//
	// If you want to get the last failure reason for a calculated health check,
	// you must use the Amazon Route 53 console or the CloudWatch console. You can't
	// use GetHealthCheckLastFailureReason for a calculated health check.
	//
	// HealthCheckId is a required field
	HealthCheckId *string `location:"uri" locationName:"HealthCheckId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetHealthCheckLastFailureReasonInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetHealthCheckLastFailureReasonInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetHealthCheckLastFailureReasonInput"}

	if s.HealthCheckId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HealthCheckId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetHealthCheckLastFailureReasonInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.HealthCheckId != nil {
		v := *s.HealthCheckId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "HealthCheckId", protocol.StringValue(v), metadata)
	}
	return nil
}

// A complex type that contains the response to a GetHealthCheckLastFailureReason
// request.
type GetHealthCheckLastFailureReasonOutput struct {
	_ struct{} `type:"structure"`

	// A list that contains one Observation element for each Amazon Route 53 health
	// checker that is reporting a last failure reason.
	//
	// HealthCheckObservations is a required field
	HealthCheckObservations []HealthCheckObservation `locationNameList:"HealthCheckObservation" type:"list" required:"true"`
}

// String returns the string representation
func (s GetHealthCheckLastFailureReasonOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetHealthCheckLastFailureReasonOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.HealthCheckObservations != nil {
		v := s.HealthCheckObservations

		metadata := protocol.Metadata{ListLocationName: "HealthCheckObservation"}
		ls0 := e.List(protocol.BodyTarget, "HealthCheckObservations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetHealthCheckLastFailureReason = "GetHealthCheckLastFailureReason"

// GetHealthCheckLastFailureReasonRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Gets the reason that a specified health check failed most recently.
//
//    // Example sending a request using GetHealthCheckLastFailureReasonRequest.
//    req := client.GetHealthCheckLastFailureReasonRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/GetHealthCheckLastFailureReason
func (c *Client) GetHealthCheckLastFailureReasonRequest(input *GetHealthCheckLastFailureReasonInput) GetHealthCheckLastFailureReasonRequest {
	op := &aws.Operation{
		Name:       opGetHealthCheckLastFailureReason,
		HTTPMethod: "GET",
		HTTPPath:   "/2013-04-01/healthcheck/{HealthCheckId}/lastfailurereason",
	}

	if input == nil {
		input = &GetHealthCheckLastFailureReasonInput{}
	}

	req := c.newRequest(op, input, &GetHealthCheckLastFailureReasonOutput{})

	return GetHealthCheckLastFailureReasonRequest{Request: req, Input: input, Copy: c.GetHealthCheckLastFailureReasonRequest}
}

// GetHealthCheckLastFailureReasonRequest is the request type for the
// GetHealthCheckLastFailureReason API operation.
type GetHealthCheckLastFailureReasonRequest struct {
	*aws.Request
	Input *GetHealthCheckLastFailureReasonInput
	Copy  func(*GetHealthCheckLastFailureReasonInput) GetHealthCheckLastFailureReasonRequest
}

// Send marshals and sends the GetHealthCheckLastFailureReason API request.
func (r GetHealthCheckLastFailureReasonRequest) Send(ctx context.Context) (*GetHealthCheckLastFailureReasonResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetHealthCheckLastFailureReasonResponse{
		GetHealthCheckLastFailureReasonOutput: r.Request.Data.(*GetHealthCheckLastFailureReasonOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetHealthCheckLastFailureReasonResponse is the response type for the
// GetHealthCheckLastFailureReason API operation.
type GetHealthCheckLastFailureReasonResponse struct {
	*GetHealthCheckLastFailureReasonOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetHealthCheckLastFailureReason request.
func (r *GetHealthCheckLastFailureReasonResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
