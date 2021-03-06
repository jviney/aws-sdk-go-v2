// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteBackupInput struct {
	_ struct{} `type:"structure"`

	// The ARN associated with the backup.
	//
	// BackupArn is a required field
	BackupArn *string `min:"37" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBackupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBackupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBackupInput"}

	if s.BackupArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupArn"))
	}
	if s.BackupArn != nil && len(*s.BackupArn) < 37 {
		invalidParams.Add(aws.NewErrParamMinLen("BackupArn", 37))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteBackupOutput struct {
	_ struct{} `type:"structure"`

	// Contains the description of the backup created for the table.
	BackupDescription *BackupDescription `type:"structure"`
}

// String returns the string representation
func (s DeleteBackupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteBackup = "DeleteBackup"

// DeleteBackupRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// Deletes an existing backup of a table.
//
// You can call DeleteBackup at a maximum rate of 10 times per second.
//
//    // Example sending a request using DeleteBackupRequest.
//    req := client.DeleteBackupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/DeleteBackup
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

	if req.Config.EnableEndpointDiscovery {
		de := discovererDescribeEndpoints{
			Client:        c,
			Required:      false,
			EndpointCache: c.endpointCache,
			Params: map[string]*string{
				"op": &req.Operation.Name,
			},
		}

		for k, v := range de.Params {
			if v == nil {
				delete(de.Params, k)
			}
		}

		req.Handlers.Build.PushFrontNamed(aws.NamedHandler{
			Name: "crr.endpointdiscovery",
			Fn:   de.Handler,
		})
	}

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
