// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteApplicationSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The name of an existing application.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// The creation timestamp of the application snapshot to delete. You can retrieve
	// this value using or .
	//
	// SnapshotCreationTimestamp is a required field
	SnapshotCreationTimestamp *time.Time `type:"timestamp" required:"true"`

	// The identifier for the snapshot delete.
	//
	// SnapshotName is a required field
	SnapshotName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteApplicationSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteApplicationSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteApplicationSnapshotInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if s.SnapshotCreationTimestamp == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotCreationTimestamp"))
	}

	if s.SnapshotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotName"))
	}
	if s.SnapshotName != nil && len(*s.SnapshotName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SnapshotName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteApplicationSnapshotOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteApplicationSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteApplicationSnapshot = "DeleteApplicationSnapshot"

// DeleteApplicationSnapshotRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
// Deletes a snapshot of application state.
//
//    // Example sending a request using DeleteApplicationSnapshotRequest.
//    req := client.DeleteApplicationSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/DeleteApplicationSnapshot
func (c *Client) DeleteApplicationSnapshotRequest(input *DeleteApplicationSnapshotInput) DeleteApplicationSnapshotRequest {
	op := &aws.Operation{
		Name:       opDeleteApplicationSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteApplicationSnapshotInput{}
	}

	req := c.newRequest(op, input, &DeleteApplicationSnapshotOutput{})

	return DeleteApplicationSnapshotRequest{Request: req, Input: input, Copy: c.DeleteApplicationSnapshotRequest}
}

// DeleteApplicationSnapshotRequest is the request type for the
// DeleteApplicationSnapshot API operation.
type DeleteApplicationSnapshotRequest struct {
	*aws.Request
	Input *DeleteApplicationSnapshotInput
	Copy  func(*DeleteApplicationSnapshotInput) DeleteApplicationSnapshotRequest
}

// Send marshals and sends the DeleteApplicationSnapshot API request.
func (r DeleteApplicationSnapshotRequest) Send(ctx context.Context) (*DeleteApplicationSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteApplicationSnapshotResponse{
		DeleteApplicationSnapshotOutput: r.Request.Data.(*DeleteApplicationSnapshotOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteApplicationSnapshotResponse is the response type for the
// DeleteApplicationSnapshot API operation.
type DeleteApplicationSnapshotResponse struct {
	*DeleteApplicationSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteApplicationSnapshot request.
func (r *DeleteApplicationSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
