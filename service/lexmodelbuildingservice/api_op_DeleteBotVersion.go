// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteBotVersionInput struct {
	_ struct{} `type:"structure"`

	// The name of the bot.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"2" type:"string" required:"true"`

	// The version of the bot to delete. You cannot delete the $LATEST version of
	// the bot. To delete the $LATEST version, use the DeleteBot operation.
	//
	// Version is a required field
	Version *string `location:"uri" locationName:"version" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBotVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBotVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBotVersionInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 2))
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
func (s DeleteBotVersionInput) MarshalFields(e protocol.FieldEncoder) error {
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

type DeleteBotVersionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteBotVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBotVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteBotVersion = "DeleteBotVersion"

// DeleteBotVersionRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Deletes a specific version of a bot. To delete all versions of a bot, use
// the DeleteBot operation.
//
// This operation requires permissions for the lex:DeleteBotVersion action.
//
//    // Example sending a request using DeleteBotVersionRequest.
//    req := client.DeleteBotVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/DeleteBotVersion
func (c *Client) DeleteBotVersionRequest(input *DeleteBotVersionInput) DeleteBotVersionRequest {
	op := &aws.Operation{
		Name:       opDeleteBotVersion,
		HTTPMethod: "DELETE",
		HTTPPath:   "/bots/{name}/versions/{version}",
	}

	if input == nil {
		input = &DeleteBotVersionInput{}
	}

	req := c.newRequest(op, input, &DeleteBotVersionOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteBotVersionRequest{Request: req, Input: input, Copy: c.DeleteBotVersionRequest}
}

// DeleteBotVersionRequest is the request type for the
// DeleteBotVersion API operation.
type DeleteBotVersionRequest struct {
	*aws.Request
	Input *DeleteBotVersionInput
	Copy  func(*DeleteBotVersionInput) DeleteBotVersionRequest
}

// Send marshals and sends the DeleteBotVersion API request.
func (r DeleteBotVersionRequest) Send(ctx context.Context) (*DeleteBotVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBotVersionResponse{
		DeleteBotVersionOutput: r.Request.Data.(*DeleteBotVersionOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBotVersionResponse is the response type for the
// DeleteBotVersion API operation.
type DeleteBotVersionResponse struct {
	*DeleteBotVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBotVersion request.
func (r *DeleteBotVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
