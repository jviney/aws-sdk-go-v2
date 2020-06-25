// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type CreateMultiplexProgramInput struct {
	_ struct{} `type:"structure"`

	// MultiplexId is a required field
	MultiplexId *string `location:"uri" locationName:"multiplexId" type:"string" required:"true"`

	// Multiplex Program settings configuration.
	//
	// MultiplexProgramSettings is a required field
	MultiplexProgramSettings *MultiplexProgramSettings `locationName:"multiplexProgramSettings" type:"structure" required:"true"`

	// ProgramName is a required field
	ProgramName *string `locationName:"programName" type:"string" required:"true"`

	// RequestId is a required field
	RequestId *string `locationName:"requestId" type:"string" required:"true" idempotencyToken:"true"`
}

// String returns the string representation
func (s CreateMultiplexProgramInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMultiplexProgramInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateMultiplexProgramInput"}

	if s.MultiplexId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MultiplexId"))
	}

	if s.MultiplexProgramSettings == nil {
		invalidParams.Add(aws.NewErrParamRequired("MultiplexProgramSettings"))
	}

	if s.ProgramName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProgramName"))
	}

	if s.RequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RequestId"))
	}
	if s.MultiplexProgramSettings != nil {
		if err := s.MultiplexProgramSettings.Validate(); err != nil {
			invalidParams.AddNested("MultiplexProgramSettings", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMultiplexProgramInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MultiplexProgramSettings != nil {
		v := s.MultiplexProgramSettings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "multiplexProgramSettings", v, metadata)
	}
	if s.ProgramName != nil {
		v := *s.ProgramName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "programName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	var RequestId string
	if s.RequestId != nil {
		RequestId = *s.RequestId
	} else {
		RequestId = protocol.GetIdempotencyToken()
	}
	{
		v := RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MultiplexId != nil {
		v := *s.MultiplexId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "multiplexId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateMultiplexProgramOutput struct {
	_ struct{} `type:"structure"`

	// The multiplex program object.
	MultiplexProgram *MultiplexProgram `locationName:"multiplexProgram" type:"structure"`
}

// String returns the string representation
func (s CreateMultiplexProgramOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMultiplexProgramOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.MultiplexProgram != nil {
		v := s.MultiplexProgram

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "multiplexProgram", v, metadata)
	}
	return nil
}

const opCreateMultiplexProgram = "CreateMultiplexProgram"

// CreateMultiplexProgramRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// Create a new program in the multiplex.
//
//    // Example sending a request using CreateMultiplexProgramRequest.
//    req := client.CreateMultiplexProgramRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/CreateMultiplexProgram
func (c *Client) CreateMultiplexProgramRequest(input *CreateMultiplexProgramInput) CreateMultiplexProgramRequest {
	op := &aws.Operation{
		Name:       opCreateMultiplexProgram,
		HTTPMethod: "POST",
		HTTPPath:   "/prod/multiplexes/{multiplexId}/programs",
	}

	if input == nil {
		input = &CreateMultiplexProgramInput{}
	}

	req := c.newRequest(op, input, &CreateMultiplexProgramOutput{})

	return CreateMultiplexProgramRequest{Request: req, Input: input, Copy: c.CreateMultiplexProgramRequest}
}

// CreateMultiplexProgramRequest is the request type for the
// CreateMultiplexProgram API operation.
type CreateMultiplexProgramRequest struct {
	*aws.Request
	Input *CreateMultiplexProgramInput
	Copy  func(*CreateMultiplexProgramInput) CreateMultiplexProgramRequest
}

// Send marshals and sends the CreateMultiplexProgram API request.
func (r CreateMultiplexProgramRequest) Send(ctx context.Context) (*CreateMultiplexProgramResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMultiplexProgramResponse{
		CreateMultiplexProgramOutput: r.Request.Data.(*CreateMultiplexProgramOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMultiplexProgramResponse is the response type for the
// CreateMultiplexProgram API operation.
type CreateMultiplexProgramResponse struct {
	*CreateMultiplexProgramOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMultiplexProgram request.
func (r *CreateMultiplexProgramResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
