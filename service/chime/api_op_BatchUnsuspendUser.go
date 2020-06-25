// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type BatchUnsuspendUserInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The request containing the user IDs to unsuspend.
	//
	// UserIdList is a required field
	UserIdList []string `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchUnsuspendUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchUnsuspendUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchUnsuspendUserInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.UserIdList == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserIdList"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchUnsuspendUserInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.UserIdList != nil {
		v := s.UserIdList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "UserIdList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type BatchUnsuspendUserOutput struct {
	_ struct{} `type:"structure"`

	// If the BatchUnsuspendUser action fails for one or more of the user IDs in
	// the request, a list of the user IDs is returned, along with error codes and
	// error messages.
	UserErrors []UserError `type:"list"`
}

// String returns the string representation
func (s BatchUnsuspendUserOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchUnsuspendUserOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.UserErrors != nil {
		v := s.UserErrors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "UserErrors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opBatchUnsuspendUser = "BatchUnsuspendUser"

// BatchUnsuspendUserRequest returns a request value for making API operation for
// Amazon Chime.
//
// Removes the suspension from up to 50 previously suspended users for the specified
// Amazon Chime EnterpriseLWA account. Only users on EnterpriseLWA accounts
// can be unsuspended using this action. For more information about different
// account types, see Managing Your Amazon Chime Accounts (https://docs.aws.amazon.com/chime/latest/ag/manage-chime-account.html)
// in the Amazon Chime Administration Guide.
//
// Previously suspended users who are unsuspended using this action are returned
// to Registered status. Users who are not previously suspended are ignored.
//
//    // Example sending a request using BatchUnsuspendUserRequest.
//    req := client.BatchUnsuspendUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/BatchUnsuspendUser
func (c *Client) BatchUnsuspendUserRequest(input *BatchUnsuspendUserInput) BatchUnsuspendUserRequest {
	op := &aws.Operation{
		Name:       opBatchUnsuspendUser,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{accountId}/users?operation=unsuspend",
	}

	if input == nil {
		input = &BatchUnsuspendUserInput{}
	}

	req := c.newRequest(op, input, &BatchUnsuspendUserOutput{})

	return BatchUnsuspendUserRequest{Request: req, Input: input, Copy: c.BatchUnsuspendUserRequest}
}

// BatchUnsuspendUserRequest is the request type for the
// BatchUnsuspendUser API operation.
type BatchUnsuspendUserRequest struct {
	*aws.Request
	Input *BatchUnsuspendUserInput
	Copy  func(*BatchUnsuspendUserInput) BatchUnsuspendUserRequest
}

// Send marshals and sends the BatchUnsuspendUser API request.
func (r BatchUnsuspendUserRequest) Send(ctx context.Context) (*BatchUnsuspendUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchUnsuspendUserResponse{
		BatchUnsuspendUserOutput: r.Request.Data.(*BatchUnsuspendUserOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchUnsuspendUserResponse is the response type for the
// BatchUnsuspendUser API operation.
type BatchUnsuspendUserResponse struct {
	*BatchUnsuspendUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchUnsuspendUser request.
func (r *BatchUnsuspendUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
