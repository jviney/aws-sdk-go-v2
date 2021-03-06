// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/query"
)

type CreateAccountAliasInput struct {
	_ struct{} `type:"structure"`

	// The account alias to create.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of lowercase letters, digits, and dashes.
	// You cannot start or finish with a dash, nor can you have two dashes in a
	// row.
	//
	// AccountAlias is a required field
	AccountAlias *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateAccountAliasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAccountAliasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAccountAliasInput"}

	if s.AccountAlias == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountAlias"))
	}
	if s.AccountAlias != nil && len(*s.AccountAlias) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountAlias", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateAccountAliasOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateAccountAliasOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateAccountAlias = "CreateAccountAlias"

// CreateAccountAliasRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Creates an alias for your AWS account. For information about using an AWS
// account alias, see Using an Alias for Your AWS Account ID (https://docs.aws.amazon.com/IAM/latest/UserGuide/AccountAlias.html)
// in the IAM User Guide.
//
//    // Example sending a request using CreateAccountAliasRequest.
//    req := client.CreateAccountAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/CreateAccountAlias
func (c *Client) CreateAccountAliasRequest(input *CreateAccountAliasInput) CreateAccountAliasRequest {
	op := &aws.Operation{
		Name:       opCreateAccountAlias,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateAccountAliasInput{}
	}

	req := c.newRequest(op, input, &CreateAccountAliasOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return CreateAccountAliasRequest{Request: req, Input: input, Copy: c.CreateAccountAliasRequest}
}

// CreateAccountAliasRequest is the request type for the
// CreateAccountAlias API operation.
type CreateAccountAliasRequest struct {
	*aws.Request
	Input *CreateAccountAliasInput
	Copy  func(*CreateAccountAliasInput) CreateAccountAliasRequest
}

// Send marshals and sends the CreateAccountAlias API request.
func (r CreateAccountAliasRequest) Send(ctx context.Context) (*CreateAccountAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAccountAliasResponse{
		CreateAccountAliasOutput: r.Request.Data.(*CreateAccountAliasOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAccountAliasResponse is the response type for the
// CreateAccountAlias API operation.
type CreateAccountAliasResponse struct {
	*CreateAccountAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateAccountAlias request.
func (r *CreateAccountAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
