// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetJourneyInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// JourneyId is a required field
	JourneyId *string `location:"uri" locationName:"journey-id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetJourneyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetJourneyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetJourneyInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.JourneyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JourneyId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetJourneyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JourneyId != nil {
		v := *s.JourneyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "journey-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetJourneyOutput struct {
	_ struct{} `type:"structure" payload:"JourneyResponse"`

	// Provides information about the status, configuration, and other settings
	// for a journey.
	//
	// JourneyResponse is a required field
	JourneyResponse *JourneyResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetJourneyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetJourneyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.JourneyResponse != nil {
		v := s.JourneyResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "JourneyResponse", v, metadata)
	}
	return nil
}

const opGetJourney = "GetJourney"

// GetJourneyRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the status, configuration, and other settings
// for a journey.
//
//    // Example sending a request using GetJourneyRequest.
//    req := client.GetJourneyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetJourney
func (c *Client) GetJourneyRequest(input *GetJourneyInput) GetJourneyRequest {
	op := &aws.Operation{
		Name:       opGetJourney,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/journeys/{journey-id}",
	}

	if input == nil {
		input = &GetJourneyInput{}
	}

	req := c.newRequest(op, input, &GetJourneyOutput{})

	return GetJourneyRequest{Request: req, Input: input, Copy: c.GetJourneyRequest}
}

// GetJourneyRequest is the request type for the
// GetJourney API operation.
type GetJourneyRequest struct {
	*aws.Request
	Input *GetJourneyInput
	Copy  func(*GetJourneyInput) GetJourneyRequest
}

// Send marshals and sends the GetJourney API request.
func (r GetJourneyRequest) Send(ctx context.Context) (*GetJourneyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetJourneyResponse{
		GetJourneyOutput: r.Request.Data.(*GetJourneyOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetJourneyResponse is the response type for the
// GetJourney API operation.
type GetJourneyResponse struct {
	*GetJourneyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetJourney request.
func (r *GetJourneyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
