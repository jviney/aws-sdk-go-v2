// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/restjson"
)

// The input values for AbortVaultLock.
type AbortVaultLockInput struct {
	_ struct{} `type:"structure"`

	// The AccountId value is the AWS account ID. This value must match the AWS
	// account ID associated with the credentials used to sign the request. You
	// can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you specify your account ID, do
	// not include any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s AbortVaultLockInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AbortVaultLockInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AbortVaultLockInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.VaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AbortVaultLockInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VaultName != nil {
		v := *s.VaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "vaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type AbortVaultLockOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AbortVaultLockOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AbortVaultLockOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opAbortVaultLock = "AbortVaultLock"

// AbortVaultLockRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation aborts the vault locking process if the vault lock is not
// in the Locked state. If the vault lock is in the Locked state when this operation
// is requested, the operation returns an AccessDeniedException error. Aborting
// the vault locking process removes the vault lock policy from the specified
// vault.
//
// A vault lock is put into the InProgress state by calling InitiateVaultLock.
// A vault lock is put into the Locked state by calling CompleteVaultLock. You
// can get the state of a vault lock by calling GetVaultLock. For more information
// about the vault locking process, see Amazon Glacier Vault Lock (https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock.html).
// For more information about vault lock policies, see Amazon Glacier Access
// Control with Vault Lock Policies (https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock-policy.html).
//
// This operation is idempotent. You can successfully invoke this operation
// multiple times, if the vault lock is in the InProgress state or if there
// is no policy associated with the vault.
//
//    // Example sending a request using AbortVaultLockRequest.
//    req := client.AbortVaultLockRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) AbortVaultLockRequest(input *AbortVaultLockInput) AbortVaultLockRequest {
	op := &aws.Operation{
		Name:       opAbortVaultLock,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{accountId}/vaults/{vaultName}/lock-policy",
	}

	if input == nil {
		input = &AbortVaultLockInput{}
	}

	req := c.newRequest(op, input, &AbortVaultLockOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return AbortVaultLockRequest{Request: req, Input: input, Copy: c.AbortVaultLockRequest}
}

// AbortVaultLockRequest is the request type for the
// AbortVaultLock API operation.
type AbortVaultLockRequest struct {
	*aws.Request
	Input *AbortVaultLockInput
	Copy  func(*AbortVaultLockInput) AbortVaultLockRequest
}

// Send marshals and sends the AbortVaultLock API request.
func (r AbortVaultLockRequest) Send(ctx context.Context) (*AbortVaultLockResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AbortVaultLockResponse{
		AbortVaultLockOutput: r.Request.Data.(*AbortVaultLockOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AbortVaultLockResponse is the response type for the
// AbortVaultLock API operation.
type AbortVaultLockResponse struct {
	*AbortVaultLockOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AbortVaultLock request.
func (r *AbortVaultLockResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
