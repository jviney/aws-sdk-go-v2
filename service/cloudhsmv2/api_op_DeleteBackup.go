// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudhsmv2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteBackupInput struct {
	_ struct{} `type:"structure"`

	// The ID of the backup to be deleted. To find the ID of a backup, use the DescribeBackups
	// operation.
	//
	// BackupId is a required field
	BackupId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBackupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBackupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBackupInput"}

	if s.BackupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteBackupOutput struct {
	_ struct{} `type:"structure"`

	// Information on the Backup object deleted.
	Backup *Backup `type:"structure"`
}

// String returns the string representation
func (s DeleteBackupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteBackup = "DeleteBackup"

// DeleteBackupRequest returns a request value for making API operation for
// AWS CloudHSM V2.
//
// Deletes a specified AWS CloudHSM backup. A backup can be restored up to 7
// days after the DeleteBackup request is made. For more information on restoring
// a backup, see RestoreBackup.
//
//    // Example sending a request using DeleteBackupRequest.
//    req := client.DeleteBackupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsmv2-2017-04-28/DeleteBackup
func (c *Client) DeleteBackupRequest(input *DeleteBackupInput) DeleteBackupRequest {
	op := &aws.Operation{
		Name:       opDeleteBackup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteBackupInput{}
	}

	req := c.newRequest(op, input, &DeleteBackupOutput{})

	return DeleteBackupRequest{Request: req, Input: input, Copy: c.DeleteBackupRequest}
}

// DeleteBackupRequest is the request type for the
// DeleteBackup API operation.
type DeleteBackupRequest struct {
	*aws.Request
	Input *DeleteBackupInput
	Copy  func(*DeleteBackupInput) DeleteBackupRequest
}

// Send marshals and sends the DeleteBackup API request.
func (r DeleteBackupRequest) Send(ctx context.Context) (*DeleteBackupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBackupResponse{
		DeleteBackupOutput: r.Request.Data.(*DeleteBackupOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBackupResponse is the response type for the
// DeleteBackup API operation.
type DeleteBackupResponse struct {
	*DeleteBackupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBackup request.
func (r *DeleteBackupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
