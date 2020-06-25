// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type UpdateReservationInput struct {
	_ struct{} `type:"structure"`

	Name *string `locationName:"name" type:"string"`

	// ReservationId is a required field
	ReservationId *string `location:"uri" locationName:"reservationId" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateReservationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateReservationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateReservationInput"}

	if s.ReservationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReservationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateReservationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ReservationId != nil {
		v := *s.ReservationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "reservationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateReservationOutput struct {
	_ struct{} `type:"structure"`

	// Reserved resources available to use
	Reservation *Reservation `locationName:"reservation" type:"structure"`
}

// String returns the string representation
func (s UpdateReservationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateReservationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Reservation != nil {
		v := s.Reservation

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "reservation", v, metadata)
	}
	return nil
}

const opUpdateReservation = "UpdateReservation"

// UpdateReservationRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// Update reservation.
//
//    // Example sending a request using UpdateReservationRequest.
//    req := client.UpdateReservationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/UpdateReservation
func (c *Client) UpdateReservationRequest(input *UpdateReservationInput) UpdateReservationRequest {
	op := &aws.Operation{
		Name:       opUpdateReservation,
		HTTPMethod: "PUT",
		HTTPPath:   "/prod/reservations/{reservationId}",
	}

	if input == nil {
		input = &UpdateReservationInput{}
	}

	req := c.newRequest(op, input, &UpdateReservationOutput{})

	return UpdateReservationRequest{Request: req, Input: input, Copy: c.UpdateReservationRequest}
}

// UpdateReservationRequest is the request type for the
// UpdateReservation API operation.
type UpdateReservationRequest struct {
	*aws.Request
	Input *UpdateReservationInput
	Copy  func(*UpdateReservationInput) UpdateReservationRequest
}

// Send marshals and sends the UpdateReservation API request.
func (r UpdateReservationRequest) Send(ctx context.Context) (*UpdateReservationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateReservationResponse{
		UpdateReservationOutput: r.Request.Data.(*UpdateReservationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateReservationResponse is the response type for the
// UpdateReservation API operation.
type UpdateReservationResponse struct {
	*UpdateReservationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateReservation request.
func (r *UpdateReservationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
