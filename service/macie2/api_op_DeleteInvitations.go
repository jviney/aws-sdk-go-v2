// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// Specifies one or more accounts that sent Amazon Macie membership invitations
// to delete.
type DeleteInvitationsInput struct {
	_ struct{} `type:"structure"`

	// AccountIds is a required field
	AccountIds []string `locationName:"accountIds" type:"list" required:"true"`
}

// String returns the string representation
func (s DeleteInvitationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteInvitationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteInvitationsInput"}

	if s.AccountIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteInvitationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountIds != nil {
		v := s.AccountIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "accountIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// Provides information about unprocessed requests to delete Amazon Macie membership
// invitations that were received from specific accounts.
type DeleteInvitationsOutput struct {
	_ struct{} `type:"structure"`

	UnprocessedAccounts []UnprocessedAccount `locationName:"unprocessedAccounts" type:"list"`
}

// String returns the string representation
func (s DeleteInvitationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteInvitationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.UnprocessedAccounts != nil {
		v := s.UnprocessedAccounts

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "unprocessedAccounts", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opDeleteInvitations = "DeleteInvitations"

// DeleteInvitationsRequest returns a request value for making API operation for
// Amazon Macie 2.
//
// Deletes Amazon Macie membership invitations that were received from specific
// accounts.
//
//    // Example sending a request using DeleteInvitationsRequest.
//    req := client.DeleteInvitationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie2-2020-01-01/DeleteInvitations
func (c *Client) DeleteInvitationsRequest(input *DeleteInvitationsInput) DeleteInvitationsRequest {
	op := &aws.Operation{
		Name:       opDeleteInvitations,
		HTTPMethod: "POST",
		HTTPPath:   "/invitations/delete",
	}

	if input == nil {
		input = &DeleteInvitationsInput{}
	}

	req := c.newRequest(op, input, &DeleteInvitationsOutput{})

	return DeleteInvitationsRequest{Request: req, Input: input, Copy: c.DeleteInvitationsRequest}
}

// DeleteInvitationsRequest is the request type for the
// DeleteInvitations API operation.
type DeleteInvitationsRequest struct {
	*aws.Request
	Input *DeleteInvitationsInput
	Copy  func(*DeleteInvitationsInput) DeleteInvitationsRequest
}

// Send marshals and sends the DeleteInvitations API request.
func (r DeleteInvitationsRequest) Send(ctx context.Context) (*DeleteInvitationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteInvitationsResponse{
		DeleteInvitationsOutput: r.Request.Data.(*DeleteInvitationsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteInvitationsResponse is the response type for the
// DeleteInvitations API operation.
type DeleteInvitationsResponse struct {
	*DeleteInvitationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteInvitations request.
func (r *DeleteInvitationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
