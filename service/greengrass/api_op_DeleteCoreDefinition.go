// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DeleteCoreDefinitionInput struct {
	_ struct{} `type:"structure"`

	// CoreDefinitionId is a required field
	CoreDefinitionId *string `location:"uri" locationName:"CoreDefinitionId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCoreDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCoreDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCoreDefinitionInput"}

	if s.CoreDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CoreDefinitionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteCoreDefinitionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CoreDefinitionId != nil {
		v := *s.CoreDefinitionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "CoreDefinitionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteCoreDefinitionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteCoreDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteCoreDefinitionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteCoreDefinition = "DeleteCoreDefinition"

// DeleteCoreDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Deletes a core definition.
//
//    // Example sending a request using DeleteCoreDefinitionRequest.
//    req := client.DeleteCoreDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/DeleteCoreDefinition
func (c *Client) DeleteCoreDefinitionRequest(input *DeleteCoreDefinitionInput) DeleteCoreDefinitionRequest {
	op := &aws.Operation{
		Name:       opDeleteCoreDefinition,
		HTTPMethod: "DELETE",
		HTTPPath:   "/greengrass/definition/cores/{CoreDefinitionId}",
	}

	if input == nil {
		input = &DeleteCoreDefinitionInput{}
	}

	req := c.newRequest(op, input, &DeleteCoreDefinitionOutput{})

	return DeleteCoreDefinitionRequest{Request: req, Input: input, Copy: c.DeleteCoreDefinitionRequest}
}

// DeleteCoreDefinitionRequest is the request type for the
// DeleteCoreDefinition API operation.
type DeleteCoreDefinitionRequest struct {
	*aws.Request
	Input *DeleteCoreDefinitionInput
	Copy  func(*DeleteCoreDefinitionInput) DeleteCoreDefinitionRequest
}

// Send marshals and sends the DeleteCoreDefinition API request.
func (r DeleteCoreDefinitionRequest) Send(ctx context.Context) (*DeleteCoreDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCoreDefinitionResponse{
		DeleteCoreDefinitionOutput: r.Request.Data.(*DeleteCoreDefinitionOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCoreDefinitionResponse is the response type for the
// DeleteCoreDefinition API operation.
type DeleteCoreDefinitionResponse struct {
	*DeleteCoreDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCoreDefinition request.
func (r *DeleteCoreDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
