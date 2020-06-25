// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetThreatIntelSetInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the detector that the threatIntelSet is associated with.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`

	// The unique ID of the threatIntelSet that you want to get.
	//
	// ThreatIntelSetId is a required field
	ThreatIntelSetId *string `location:"uri" locationName:"threatIntelSetId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetThreatIntelSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetThreatIntelSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetThreatIntelSetInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if s.ThreatIntelSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThreatIntelSetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetThreatIntelSetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThreatIntelSetId != nil {
		v := *s.ThreatIntelSetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "threatIntelSetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetThreatIntelSetOutput struct {
	_ struct{} `type:"structure"`

	// The format of the threatIntelSet.
	//
	// Format is a required field
	Format ThreatIntelSetFormat `locationName:"format" min:"1" type:"string" required:"true" enum:"true"`

	// The URI of the file that contains the ThreatIntelSet.
	//
	// Location is a required field
	Location *string `locationName:"location" min:"1" type:"string" required:"true"`

	// A user-friendly ThreatIntelSet name displayed in all findings that are generated
	// by activity that involves IP addresses included in this ThreatIntelSet.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// The status of threatIntelSet file uploaded.
	//
	// Status is a required field
	Status ThreatIntelSetStatus `locationName:"status" min:"1" type:"string" required:"true" enum:"true"`

	// The tags of the threat list resource.
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`
}

// String returns the string representation
func (s GetThreatIntelSetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetThreatIntelSetOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Format) > 0 {
		v := s.Format

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "format", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Location != nil {
		v := *s.Location

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "location", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opGetThreatIntelSet = "GetThreatIntelSet"

// GetThreatIntelSetRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Retrieves the ThreatIntelSet that is specified by the ThreatIntelSet ID.
//
//    // Example sending a request using GetThreatIntelSetRequest.
//    req := client.GetThreatIntelSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/GetThreatIntelSet
func (c *Client) GetThreatIntelSetRequest(input *GetThreatIntelSetInput) GetThreatIntelSetRequest {
	op := &aws.Operation{
		Name:       opGetThreatIntelSet,
		HTTPMethod: "GET",
		HTTPPath:   "/detector/{detectorId}/threatintelset/{threatIntelSetId}",
	}

	if input == nil {
		input = &GetThreatIntelSetInput{}
	}

	req := c.newRequest(op, input, &GetThreatIntelSetOutput{})

	return GetThreatIntelSetRequest{Request: req, Input: input, Copy: c.GetThreatIntelSetRequest}
}

// GetThreatIntelSetRequest is the request type for the
// GetThreatIntelSet API operation.
type GetThreatIntelSetRequest struct {
	*aws.Request
	Input *GetThreatIntelSetInput
	Copy  func(*GetThreatIntelSetInput) GetThreatIntelSetRequest
}

// Send marshals and sends the GetThreatIntelSet API request.
func (r GetThreatIntelSetRequest) Send(ctx context.Context) (*GetThreatIntelSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetThreatIntelSetResponse{
		GetThreatIntelSetOutput: r.Request.Data.(*GetThreatIntelSetOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetThreatIntelSetResponse is the response type for the
// GetThreatIntelSet API operation.
type GetThreatIntelSetResponse struct {
	*GetThreatIntelSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetThreatIntelSet request.
func (r *GetThreatIntelSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
