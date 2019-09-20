// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Specifies an account to associate with an Amazon Macie master account.
type CreateMemberInput struct {
	_ struct{} `type:"structure"`

	// Specifies details for an account to associate with an Amazon Macie master
	// account.
	//
	// Account is a required field
	Account *AccountDetail `locationName:"account" type:"structure" required:"true"`

	// A string-to-string map of key-value pairs that specifies the tags (keys and
	// values) for a classification job, custom data identifier, findings filter,
	// or member account.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateMemberInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMemberInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateMemberInput"}

	if s.Account == nil {
		invalidParams.Add(aws.NewErrParamRequired("Account"))
	}
	if s.Account != nil {
		if err := s.Account.Validate(); err != nil {
			invalidParams.AddNested("Account", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMemberInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Account != nil {
		v := s.Account

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "account", v, metadata)
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

// Provides information about a request to associate an account with an Amazon
// Macie master account.
type CreateMemberOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `locationName:"arn" type:"string"`
}

// String returns the string representation
func (s CreateMemberOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMemberOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateMember = "CreateMember"

// CreateMemberRequest returns a request value for making API operation for
// Amazon Macie 2.
//
// Associates an account with an Amazon Macie master account.
//
//    // Example sending a request using CreateMemberRequest.
//    req := client.CreateMemberRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie2-2020-01-01/CreateMember
func (c *Client) CreateMemberRequest(input *CreateMemberInput) CreateMemberRequest {
	op := &aws.Operation{
		Name:       opCreateMember,
		HTTPMethod: "POST",
		HTTPPath:   "/members",
	}

	if input == nil {
		input = &CreateMemberInput{}
	}

	req := c.newRequest(op, input, &CreateMemberOutput{})

	return CreateMemberRequest{Request: req, Input: input, Copy: c.CreateMemberRequest}
}

// CreateMemberRequest is the request type for the
// CreateMember API operation.
type CreateMemberRequest struct {
	*aws.Request
	Input *CreateMemberInput
	Copy  func(*CreateMemberInput) CreateMemberRequest
}

// Send marshals and sends the CreateMember API request.
func (r CreateMemberRequest) Send(ctx context.Context) (*CreateMemberResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMemberResponse{
		CreateMemberOutput: r.Request.Data.(*CreateMemberOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMemberResponse is the response type for the
// CreateMember API operation.
type CreateMemberResponse struct {
	*CreateMemberOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMember request.
func (r *CreateMemberResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}