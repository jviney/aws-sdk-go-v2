// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconvert

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// Query a preset by sending a request with the preset name.
type GetPresetInput struct {
	_ struct{} `type:"structure"`

	// The name of the preset.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPresetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPresetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPresetInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPresetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Successful get preset requests will return an OK message and the preset JSON.
type GetPresetOutput struct {
	_ struct{} `type:"structure"`

	// A preset is a collection of preconfigured media conversion settings that
	// you want MediaConvert to apply to the output during the conversion process.
	Preset *Preset `locationName:"preset" type:"structure"`
}

// String returns the string representation
func (s GetPresetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPresetOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Preset != nil {
		v := s.Preset

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "preset", v, metadata)
	}
	return nil
}

const opGetPreset = "GetPreset"

// GetPresetRequest returns a request value for making API operation for
// AWS Elemental MediaConvert.
//
// Retrieve the JSON for a specific preset.
//
//    // Example sending a request using GetPresetRequest.
//    req := client.GetPresetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/GetPreset
func (c *Client) GetPresetRequest(input *GetPresetInput) GetPresetRequest {
	op := &aws.Operation{
		Name:       opGetPreset,
		HTTPMethod: "GET",
		HTTPPath:   "/2017-08-29/presets/{name}",
	}

	if input == nil {
		input = &GetPresetInput{}
	}

	req := c.newRequest(op, input, &GetPresetOutput{})

	return GetPresetRequest{Request: req, Input: input, Copy: c.GetPresetRequest}
}

// GetPresetRequest is the request type for the
// GetPreset API operation.
type GetPresetRequest struct {
	*aws.Request
	Input *GetPresetInput
	Copy  func(*GetPresetInput) GetPresetRequest
}

// Send marshals and sends the GetPreset API request.
func (r GetPresetRequest) Send(ctx context.Context) (*GetPresetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPresetResponse{
		GetPresetOutput: r.Request.Data.(*GetPresetOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPresetResponse is the response type for the
// GetPreset API operation.
type GetPresetResponse struct {
	*GetPresetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPreset request.
func (r *GetPresetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
