// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type UpdateUserInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string" sensitive:"true"`

	// The given name of the user.
	GivenName *string `min:"1" type:"string"`

	// Boolean value to determine whether the user is granted Poweruser privileges.
	GrantPoweruserPrivileges BooleanEnumType `type:"string" enum:"true"`

	// The locale of the user.
	Locale LocaleType `type:"string" enum:"true"`

	// The amount of storage for the user.
	StorageRule *StorageRuleType `type:"structure"`

	// The surname of the user.
	Surname *string `min:"1" type:"string"`

	// The time zone ID of the user.
	TimeZoneId *string `min:"1" type:"string"`

	// The type of the user.
	Type UserType `type:"string" enum:"true"`

	// The ID of the user.
	//
	// UserId is a required field
	UserId *string `location:"uri" locationName:"UserId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateUserInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}
	if s.GivenName != nil && len(*s.GivenName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GivenName", 1))
	}
	if s.Surname != nil && len(*s.Surname) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Surname", 1))
	}
	if s.TimeZoneId != nil && len(*s.TimeZoneId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TimeZoneId", 1))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}
	if s.UserId != nil && len(*s.UserId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateUserInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GivenName != nil {
		v := *s.GivenName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GivenName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.GrantPoweruserPrivileges) > 0 {
		v := s.GrantPoweruserPrivileges

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GrantPoweruserPrivileges", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.Locale) > 0 {
		v := s.Locale

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Locale", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StorageRule != nil {
		v := s.StorageRule

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "StorageRule", v, metadata)
	}
	if s.Surname != nil {
		v := *s.Surname

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Surname", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TimeZoneId != nil {
		v := *s.TimeZoneId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TimeZoneId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.AuthenticationToken != nil {
		v := *s.AuthenticationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Authentication", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserId != nil {
		v := *s.UserId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "UserId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateUserOutput struct {
	_ struct{} `type:"structure"`

	// The user information.
	User *User `type:"structure"`
}

// String returns the string representation
func (s UpdateUserOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateUserOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.User != nil {
		v := s.User

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "User", v, metadata)
	}
	return nil
}

const opUpdateUser = "UpdateUser"

// UpdateUserRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Updates the specified attributes of the specified user, and grants or revokes
// administrative privileges to the Amazon WorkDocs site.
//
//    // Example sending a request using UpdateUserRequest.
//    req := client.UpdateUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/UpdateUser
func (c *Client) UpdateUserRequest(input *UpdateUserInput) UpdateUserRequest {
	op := &aws.Operation{
		Name:       opUpdateUser,
		HTTPMethod: "PATCH",
		HTTPPath:   "/api/v1/users/{UserId}",
	}

	if input == nil {
		input = &UpdateUserInput{}
	}

	req := c.newRequest(op, input, &UpdateUserOutput{})

	return UpdateUserRequest{Request: req, Input: input, Copy: c.UpdateUserRequest}
}

// UpdateUserRequest is the request type for the
// UpdateUser API operation.
type UpdateUserRequest struct {
	*aws.Request
	Input *UpdateUserInput
	Copy  func(*UpdateUserInput) UpdateUserRequest
}

// Send marshals and sends the UpdateUser API request.
func (r UpdateUserRequest) Send(ctx context.Context) (*UpdateUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateUserResponse{
		UpdateUserOutput: r.Request.Data.(*UpdateUserOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateUserResponse is the response type for the
// UpdateUser API operation.
type UpdateUserResponse struct {
	*UpdateUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateUser request.
func (r *UpdateUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
