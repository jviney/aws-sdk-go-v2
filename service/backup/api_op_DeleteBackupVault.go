// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteBackupVaultInput struct {
	_ struct{} `type:"structure"`

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and
	// the AWS Region where they are created. They consist of lowercase letters,
	// numbers, and hyphens.
	//
	// BackupVaultName is a required field
	BackupVaultName *string `location:"uri" locationName:"backupVaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBackupVaultInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBackupVaultInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBackupVaultInput"}

	if s.BackupVaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupVaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBackupVaultInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BackupVaultName != nil {
		v := *s.BackupVaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "backupVaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteBackupVaultOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteBackupVaultOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBackupVaultOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteBackupVault = "DeleteBackupVault"

// DeleteBackupVaultRequest returns a request value for making API operation for
// AWS Backup.
//
// Deletes the backup vault identified by its name. A vault can be deleted only
// if it is empty.
//
//    // Example sending a request using DeleteBackupVaultRequest.
//    req := client.DeleteBackupVaultRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/DeleteBackupVault
func (c *Client) DeleteBackupVaultRequest(input *DeleteBackupVaultInput) DeleteBackupVaultRequest {
	op := &aws.Operation{
		Name:       opDeleteBackupVault,
		HTTPMethod: "DELETE",
		HTTPPath:   "/backup-vaults/{backupVaultName}",
	}

	if input == nil {
		input = &DeleteBackupVaultInput{}
	}

	req := c.newRequest(op, input, &DeleteBackupVaultOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteBackupVaultRequest{Request: req, Input: input, Copy: c.DeleteBackupVaultRequest}
}

// DeleteBackupVaultRequest is the request type for the
// DeleteBackupVault API operation.
type DeleteBackupVaultRequest struct {
	*aws.Request
	Input *DeleteBackupVaultInput
	Copy  func(*DeleteBackupVaultInput) DeleteBackupVaultRequest
}

// Send marshals and sends the DeleteBackupVault API request.
func (r DeleteBackupVaultRequest) Send(ctx context.Context) (*DeleteBackupVaultResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBackupVaultResponse{
		DeleteBackupVaultOutput: r.Request.Data.(*DeleteBackupVaultOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBackupVaultResponse is the response type for the
// DeleteBackupVault API operation.
type DeleteBackupVaultResponse struct {
	*DeleteBackupVaultOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBackupVault request.
func (r *DeleteBackupVaultResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
