// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type CreateBotInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The bot display name.
	//
	// DisplayName is a required field
	DisplayName *string `type:"string" required:"true" sensitive:"true"`

	// The domain of the Amazon Chime Enterprise account.
	Domain *string `type:"string"`
}

// String returns the string representation
func (s CreateBotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateBotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateBotInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.DisplayName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DisplayName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateBotInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DisplayName != nil {
		v := *s.DisplayName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DisplayName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Domain != nil {
		v := *s.Domain

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Domain", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateBotOutput struct {
	_ struct{} `type:"structure"`

	// The bot details.
	Bot *Bot `type:"structure"`
}

// String returns the string representation
func (s CreateBotOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateBotOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Bot != nil {
		v := s.Bot

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Bot", v, metadata)
	}
	return nil
}

const opCreateBot = "CreateBot"

// CreateBotRequest returns a request value for making API operation for
// Amazon Chime.
//
// Creates a bot for an Amazon Chime Enterprise account.
//
//    // Example sending a request using CreateBotRequest.
//    req := client.CreateBotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/CreateBot
func (c *Client) CreateBotRequest(input *CreateBotInput) CreateBotRequest {
	op := &aws.Operation{
		Name:       opCreateBot,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{accountId}/bots",
	}

	if input == nil {
		input = &CreateBotInput{}
	}

	req := c.newRequest(op, input, &CreateBotOutput{})

	return CreateBotRequest{Request: req, Input: input, Copy: c.CreateBotRequest}
}

// CreateBotRequest is the request type for the
// CreateBot API operation.
type CreateBotRequest struct {
	*aws.Request
	Input *CreateBotInput
	Copy  func(*CreateBotInput) CreateBotRequest
}

// Send marshals and sends the CreateBot API request.
func (r CreateBotRequest) Send(ctx context.Context) (*CreateBotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateBotResponse{
		CreateBotOutput: r.Request.Data.(*CreateBotOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateBotResponse is the response type for the
// CreateBot API operation.
type CreateBotResponse struct {
	*CreateBotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateBot request.
func (r *CreateBotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
