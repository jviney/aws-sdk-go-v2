// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetSlotTypeInput struct {
	_ struct{} `type:"structure"`

	// The name of the slot type. The name is case sensitive.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`

	// The version of the slot type.
	//
	// Version is a required field
	Version *string `location:"uri" locationName:"version" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetSlotTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSlotTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSlotTypeInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.Version == nil {
		invalidParams.Add(aws.NewErrParamRequired("Version"))
	}
	if s.Version != nil && len(*s.Version) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Version", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSlotTypeInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetSlotTypeOutput struct {
	_ struct{} `type:"structure"`

	// Checksum of the $LATEST version of the slot type.
	Checksum *string `locationName:"checksum" type:"string"`

	// The date that the slot type was created.
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp"`

	// A description of the slot type.
	Description *string `locationName:"description" type:"string"`

	// A list of EnumerationValue objects that defines the values that the slot
	// type can take.
	EnumerationValues []EnumerationValue `locationName:"enumerationValues" type:"list"`

	// The date that the slot type was updated. When you create a resource, the
	// creation date and last update date are the same.
	LastUpdatedDate *time.Time `locationName:"lastUpdatedDate" type:"timestamp"`

	// The name of the slot type.
	Name *string `locationName:"name" min:"1" type:"string"`

	// The built-in slot type used as a parent for the slot type.
	ParentSlotTypeSignature *string `locationName:"parentSlotTypeSignature" min:"1" type:"string"`

	// Configuration information that extends the parent built-in slot type.
	SlotTypeConfigurations []SlotTypeConfiguration `locationName:"slotTypeConfigurations" type:"list"`

	// The strategy that Amazon Lex uses to determine the value of the slot. For
	// more information, see PutSlotType.
	ValueSelectionStrategy SlotValueSelectionStrategy `locationName:"valueSelectionStrategy" type:"string" enum:"true"`

	// The version of the slot type.
	Version *string `locationName:"version" min:"1" type:"string"`
}

// String returns the string representation
func (s GetSlotTypeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSlotTypeOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Checksum != nil {
		v := *s.Checksum

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "checksum", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EnumerationValues != nil {
		v := s.EnumerationValues

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "enumerationValues", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.LastUpdatedDate != nil {
		v := *s.LastUpdatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ParentSlotTypeSignature != nil {
		v := *s.ParentSlotTypeSignature

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "parentSlotTypeSignature", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SlotTypeConfigurations != nil {
		v := s.SlotTypeConfigurations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "slotTypeConfigurations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.ValueSelectionStrategy) > 0 {
		v := s.ValueSelectionStrategy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "valueSelectionStrategy", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetSlotType = "GetSlotType"

// GetSlotTypeRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Returns information about a specific version of a slot type. In addition
// to specifying the slot type name, you must specify the slot type version.
//
// This operation requires permissions for the lex:GetSlotType action.
//
//    // Example sending a request using GetSlotTypeRequest.
//    req := client.GetSlotTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetSlotType
func (c *Client) GetSlotTypeRequest(input *GetSlotTypeInput) GetSlotTypeRequest {
	op := &aws.Operation{
		Name:       opGetSlotType,
		HTTPMethod: "GET",
		HTTPPath:   "/slottypes/{name}/versions/{version}",
	}

	if input == nil {
		input = &GetSlotTypeInput{}
	}

	req := c.newRequest(op, input, &GetSlotTypeOutput{})

	return GetSlotTypeRequest{Request: req, Input: input, Copy: c.GetSlotTypeRequest}
}

// GetSlotTypeRequest is the request type for the
// GetSlotType API operation.
type GetSlotTypeRequest struct {
	*aws.Request
	Input *GetSlotTypeInput
	Copy  func(*GetSlotTypeInput) GetSlotTypeRequest
}

// Send marshals and sends the GetSlotType API request.
func (r GetSlotTypeRequest) Send(ctx context.Context) (*GetSlotTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSlotTypeResponse{
		GetSlotTypeOutput: r.Request.Data.(*GetSlotTypeOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSlotTypeResponse is the response type for the
// GetSlotType API operation.
type GetSlotTypeResponse struct {
	*GetSlotTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSlotType request.
func (r *GetSlotTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
