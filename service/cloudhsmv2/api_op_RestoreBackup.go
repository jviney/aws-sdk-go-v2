// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudhsmv2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type RestoreBackupInput struct {
	_ struct{} `type:"structure"`

	// The ID of the backup to be restored. To find the ID of a backup, use the
	// DescribeBackups operation.
	//
	// BackupId is a required field
	BackupId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RestoreBackupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreBackupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestoreBackupInput"}

	if s.BackupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RestoreBackupOutput struct {
	_ struct{} `type:"structure"`

	// Information on the Backup object created.
	Backup *Backup `type:"structure"`
}

// String returns the string representation
func (s RestoreBackupOutput) String() string {
	return awsutil.Prettify(s)
}

const opRestoreBackup = "RestoreBackup"

// RestoreBackupRequest returns a request value for making API operation for
// AWS CloudHSM V2.
//
// Restores a specified AWS CloudHSM backup that is in the PENDING_DELETION
// state. For mor information on deleting a backup, see DeleteBackup.
//
//    // Example sending a request using RestoreBackupRequest.
//    req := client.RestoreBackupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsmv2-2017-04-28/RestoreBackup
func (c *Client) RestoreBackupRequest(input *RestoreBackupInput) RestoreBackupRequest {
	op := &aws.Operation{
		Name:       opRestoreBackup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestoreBackupInput{}
	}

	req := c.newRequest(op, input, &RestoreBackupOutput{})

	return RestoreBackupRequest{Request: req, Input: input, Copy: c.RestoreBackupRequest}
}

// RestoreBackupRequest is the request type for the
// RestoreBackup API operation.
type RestoreBackupRequest struct {
	*aws.Request
	Input *RestoreBackupInput
	Copy  func(*RestoreBackupInput) RestoreBackupRequest
}

// Send marshals and sends the RestoreBackup API request.
func (r RestoreBackupRequest) Send(ctx context.Context) (*RestoreBackupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreBackupResponse{
		RestoreBackupOutput: r.Request.Data.(*RestoreBackupOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreBackupResponse is the response type for the
// RestoreBackup API operation.
type RestoreBackupResponse struct {
	*RestoreBackupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreBackup request.
func (r *RestoreBackupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
