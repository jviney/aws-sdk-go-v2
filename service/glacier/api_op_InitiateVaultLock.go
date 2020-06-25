// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// The input values for InitiateVaultLock.
type InitiateVaultLockInput struct {
	_ struct{} `type:"structure" payload:"Policy"`

	// The AccountId value is the AWS account ID. This value must match the AWS
	// account ID associated with the credentials used to sign the request. You
	// can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you specify your account ID, do
	// not include any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The vault lock policy as a JSON string, which uses "\" as an escape character.
	Policy *VaultLockPolicy `locationName:"policy" type:"structure"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s InitiateVaultLockInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InitiateVaultLockInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "InitiateVaultLockInput"}

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
func (s InitiateVaultLockInput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.Policy != nil {
		v := s.Policy

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "policy", v, metadata)
	}
	return nil
}

// Contains the Amazon S3 Glacier response to your request.
type InitiateVaultLockOutput struct {
	_ struct{} `type:"structure"`

	// The lock ID, which is used to complete the vault locking process.
	LockId *string `location:"header" locationName:"x-amz-lock-id" type:"string"`
}

// String returns the string representation
func (s InitiateVaultLockOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s InitiateVaultLockOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.LockId != nil {
		v := *s.LockId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-lock-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opInitiateVaultLock = "InitiateVaultLock"

// InitiateVaultLockRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation initiates the vault locking process by doing the following:
//
//    * Installing a vault lock policy on the specified vault.
//
//    * Setting the lock state of vault lock to InProgress.
//
//    * Returning a lock ID, which is used to complete the vault locking process.
//
// You can set one vault lock policy for each vault and this policy can be up
// to 20 KB in size. For more information about vault lock policies, see Amazon
// Glacier Access Control with Vault Lock Policies (https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock-policy.html).
//
// You must complete the vault locking process within 24 hours after the vault
// lock enters the InProgress state. After the 24 hour window ends, the lock
// ID expires, the vault automatically exits the InProgress state, and the vault
// lock policy is removed from the vault. You call CompleteVaultLock to complete
// the vault locking process by setting the state of the vault lock to Locked.
//
// After a vault lock is in the Locked state, you cannot initiate a new vault
// lock for the vault.
//
// You can abort the vault locking process by calling AbortVaultLock. You can
// get the state of the vault lock by calling GetVaultLock. For more information
// about the vault locking process, Amazon Glacier Vault Lock (https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock.html).
//
// If this operation is called when the vault lock is in the InProgress state,
// the operation returns an AccessDeniedException error. When the vault lock
// is in the InProgress state you must call AbortVaultLock before you can initiate
// a new vault lock policy.
//
//    // Example sending a request using InitiateVaultLockRequest.
//    req := client.InitiateVaultLockRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) InitiateVaultLockRequest(input *InitiateVaultLockInput) InitiateVaultLockRequest {
	op := &aws.Operation{
		Name:       opInitiateVaultLock,
		HTTPMethod: "POST",
		HTTPPath:   "/{accountId}/vaults/{vaultName}/lock-policy",
	}

	if input == nil {
		input = &InitiateVaultLockInput{}
	}

	req := c.newRequest(op, input, &InitiateVaultLockOutput{})

	return InitiateVaultLockRequest{Request: req, Input: input, Copy: c.InitiateVaultLockRequest}
}

// InitiateVaultLockRequest is the request type for the
// InitiateVaultLock API operation.
type InitiateVaultLockRequest struct {
	*aws.Request
	Input *InitiateVaultLockInput
	Copy  func(*InitiateVaultLockInput) InitiateVaultLockRequest
}

// Send marshals and sends the InitiateVaultLock API request.
func (r InitiateVaultLockRequest) Send(ctx context.Context) (*InitiateVaultLockResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &InitiateVaultLockResponse{
		InitiateVaultLockOutput: r.Request.Data.(*InitiateVaultLockOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// InitiateVaultLockResponse is the response type for the
// InitiateVaultLock API operation.
type InitiateVaultLockResponse struct {
	*InitiateVaultLockOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// InitiateVaultLock request.
func (r *InitiateVaultLockResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
