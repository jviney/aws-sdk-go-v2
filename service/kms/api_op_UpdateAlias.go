// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type UpdateAliasInput struct {
	_ struct{} `type:"structure"`

	// Identifies the alias that is changing its CMK. This value must begin with
	// alias/ followed by the alias name, such as alias/ExampleAlias. You cannot
	// use UpdateAlias to change the alias name.
	//
	// AliasName is a required field
	AliasName *string `min:"1" type:"string" required:"true"`

	// Identifies the CMK to associate with the alias. When the update operation
	// completes, the alias will point to this CMK.
	//
	// The CMK must be in the same AWS account and Region as the alias. Also, the
	// new target CMK must be the same type as the current target CMK (both symmetric
	// or both asymmetric) and they must have the same key usage.
	//
	// Specify the key ID or the Amazon Resource Name (ARN) of the CMK.
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// To verify that the alias is mapped to the correct CMK, use ListAliases.
	//
	// TargetKeyId is a required field
	TargetKeyId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateAliasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAliasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAliasInput"}

	if s.AliasName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AliasName"))
	}
	if s.AliasName != nil && len(*s.AliasName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AliasName", 1))
	}

	if s.TargetKeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetKeyId"))
	}
	if s.TargetKeyId != nil && len(*s.TargetKeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetKeyId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateAliasOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateAliasOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateAlias = "UpdateAlias"

// UpdateAliasRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Associates an existing AWS KMS alias with a different customer master key
// (CMK). Each alias is associated with only one CMK at a time, although a CMK
// can have multiple aliases. The alias and the CMK must be in the same AWS
// account and region. You cannot perform this operation on an alias in a different
// AWS account.
//
// The current and new CMK must be the same type (both symmetric or both asymmetric),
// and they must have the same key usage (ENCRYPT_DECRYPT or SIGN_VERIFY). This
// restriction prevents errors in code that uses aliases. If you must assign
// an alias to a different type of CMK, use DeleteAlias to delete the old alias
// and CreateAlias to create a new alias.
//
// You cannot use UpdateAlias to change an alias name. To change an alias name,
// use DeleteAlias to delete the old alias and CreateAlias to create a new alias.
//
// Because an alias is not a property of a CMK, you can create, update, and
// delete the aliases of a CMK without affecting the CMK. Also, aliases do not
// appear in the response from the DescribeKey operation. To get the aliases
// of all CMKs in the account, use the ListAliases operation.
//
// The CMK that you use for this operation must be in a compatible key state.
// For details, see How Key State Affects Use of a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the AWS Key Management Service Developer Guide.
//
//    // Example sending a request using UpdateAliasRequest.
//    req := client.UpdateAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/UpdateAlias
func (c *Client) UpdateAliasRequest(input *UpdateAliasInput) UpdateAliasRequest {
	op := &aws.Operation{
		Name:       opUpdateAlias,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateAliasInput{}
	}

	req := c.newRequest(op, input, &UpdateAliasOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return UpdateAliasRequest{Request: req, Input: input, Copy: c.UpdateAliasRequest}
}

// UpdateAliasRequest is the request type for the
// UpdateAlias API operation.
type UpdateAliasRequest struct {
	*aws.Request
	Input *UpdateAliasInput
	Copy  func(*UpdateAliasInput) UpdateAliasRequest
}

// Send marshals and sends the UpdateAlias API request.
func (r UpdateAliasRequest) Send(ctx context.Context) (*UpdateAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAliasResponse{
		UpdateAliasOutput: r.Request.Data.(*UpdateAliasOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAliasResponse is the response type for the
// UpdateAlias API operation.
type UpdateAliasResponse struct {
	*UpdateAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAlias request.
func (r *UpdateAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
