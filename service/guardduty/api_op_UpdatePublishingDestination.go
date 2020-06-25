// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type UpdatePublishingDestinationInput struct {
	_ struct{} `type:"structure"`

	// The ID of the publishing destination to update.
	//
	// DestinationId is a required field
	DestinationId *string `location:"uri" locationName:"destinationId" type:"string" required:"true"`

	// A DestinationProperties object that includes the DestinationArn and KmsKeyArn
	// of the publishing destination.
	DestinationProperties *DestinationProperties `locationName:"destinationProperties" type:"structure"`

	// The ID of the detector associated with the publishing destinations to update.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdatePublishingDestinationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdatePublishingDestinationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdatePublishingDestinationInput"}

	if s.DestinationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationId"))
	}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdatePublishingDestinationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DestinationProperties != nil {
		v := s.DestinationProperties

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "destinationProperties", v, metadata)
	}
	if s.DestinationId != nil {
		v := *s.DestinationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "destinationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdatePublishingDestinationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdatePublishingDestinationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdatePublishingDestinationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdatePublishingDestination = "UpdatePublishingDestination"

// UpdatePublishingDestinationRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Updates information about the publishing destination specified by the destinationId.
//
//    // Example sending a request using UpdatePublishingDestinationRequest.
//    req := client.UpdatePublishingDestinationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/UpdatePublishingDestination
func (c *Client) UpdatePublishingDestinationRequest(input *UpdatePublishingDestinationInput) UpdatePublishingDestinationRequest {
	op := &aws.Operation{
		Name:       opUpdatePublishingDestination,
		HTTPMethod: "POST",
		HTTPPath:   "/detector/{detectorId}/publishingDestination/{destinationId}",
	}

	if input == nil {
		input = &UpdatePublishingDestinationInput{}
	}

	req := c.newRequest(op, input, &UpdatePublishingDestinationOutput{})

	return UpdatePublishingDestinationRequest{Request: req, Input: input, Copy: c.UpdatePublishingDestinationRequest}
}

// UpdatePublishingDestinationRequest is the request type for the
// UpdatePublishingDestination API operation.
type UpdatePublishingDestinationRequest struct {
	*aws.Request
	Input *UpdatePublishingDestinationInput
	Copy  func(*UpdatePublishingDestinationInput) UpdatePublishingDestinationRequest
}

// Send marshals and sends the UpdatePublishingDestination API request.
func (r UpdatePublishingDestinationRequest) Send(ctx context.Context) (*UpdatePublishingDestinationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdatePublishingDestinationResponse{
		UpdatePublishingDestinationOutput: r.Request.Data.(*UpdatePublishingDestinationOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdatePublishingDestinationResponse is the response type for the
// UpdatePublishingDestination API operation.
type UpdatePublishingDestinationResponse struct {
	*UpdatePublishingDestinationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdatePublishingDestination request.
func (r *UpdatePublishingDestinationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
