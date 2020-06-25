// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type AssociateSigninDelegateGroupsWithAccountInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The sign-in delegate groups.
	//
	// SigninDelegateGroups is a required field
	SigninDelegateGroups []SigninDelegateGroup `type:"list" required:"true"`
}

// String returns the string representation
func (s AssociateSigninDelegateGroupsWithAccountInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateSigninDelegateGroupsWithAccountInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateSigninDelegateGroupsWithAccountInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.SigninDelegateGroups == nil {
		invalidParams.Add(aws.NewErrParamRequired("SigninDelegateGroups"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociateSigninDelegateGroupsWithAccountInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.SigninDelegateGroups != nil {
		v := s.SigninDelegateGroups

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "SigninDelegateGroups", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
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

type AssociateSigninDelegateGroupsWithAccountOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateSigninDelegateGroupsWithAccountOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociateSigninDelegateGroupsWithAccountOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opAssociateSigninDelegateGroupsWithAccount = "AssociateSigninDelegateGroupsWithAccount"

// AssociateSigninDelegateGroupsWithAccountRequest returns a request value for making API operation for
// Amazon Chime.
//
// Associates the specified sign-in delegate groups with the specified Amazon
// Chime account.
//
//    // Example sending a request using AssociateSigninDelegateGroupsWithAccountRequest.
//    req := client.AssociateSigninDelegateGroupsWithAccountRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/AssociateSigninDelegateGroupsWithAccount
func (c *Client) AssociateSigninDelegateGroupsWithAccountRequest(input *AssociateSigninDelegateGroupsWithAccountInput) AssociateSigninDelegateGroupsWithAccountRequest {
	op := &aws.Operation{
		Name:       opAssociateSigninDelegateGroupsWithAccount,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{accountId}?operation=associate-signin-delegate-groups",
	}

	if input == nil {
		input = &AssociateSigninDelegateGroupsWithAccountInput{}
	}

	req := c.newRequest(op, input, &AssociateSigninDelegateGroupsWithAccountOutput{})

	return AssociateSigninDelegateGroupsWithAccountRequest{Request: req, Input: input, Copy: c.AssociateSigninDelegateGroupsWithAccountRequest}
}

// AssociateSigninDelegateGroupsWithAccountRequest is the request type for the
// AssociateSigninDelegateGroupsWithAccount API operation.
type AssociateSigninDelegateGroupsWithAccountRequest struct {
	*aws.Request
	Input *AssociateSigninDelegateGroupsWithAccountInput
	Copy  func(*AssociateSigninDelegateGroupsWithAccountInput) AssociateSigninDelegateGroupsWithAccountRequest
}

// Send marshals and sends the AssociateSigninDelegateGroupsWithAccount API request.
func (r AssociateSigninDelegateGroupsWithAccountRequest) Send(ctx context.Context) (*AssociateSigninDelegateGroupsWithAccountResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateSigninDelegateGroupsWithAccountResponse{
		AssociateSigninDelegateGroupsWithAccountOutput: r.Request.Data.(*AssociateSigninDelegateGroupsWithAccountOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateSigninDelegateGroupsWithAccountResponse is the response type for the
// AssociateSigninDelegateGroupsWithAccount API operation.
type AssociateSigninDelegateGroupsWithAccountResponse struct {
	*AssociateSigninDelegateGroupsWithAccountOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateSigninDelegateGroupsWithAccount request.
func (r *AssociateSigninDelegateGroupsWithAccountResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
