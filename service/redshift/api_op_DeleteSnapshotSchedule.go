// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/query"
)

type DeleteSnapshotScheduleInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier of the snapshot schedule to delete.
	//
	// ScheduleIdentifier is a required field
	ScheduleIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSnapshotScheduleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSnapshotScheduleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSnapshotScheduleInput"}

	if s.ScheduleIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScheduleIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteSnapshotScheduleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteSnapshotScheduleOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteSnapshotSchedule = "DeleteSnapshotSchedule"

// DeleteSnapshotScheduleRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Deletes a snapshot schedule.
//
//    // Example sending a request using DeleteSnapshotScheduleRequest.
//    req := client.DeleteSnapshotScheduleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DeleteSnapshotSchedule
func (c *Client) DeleteSnapshotScheduleRequest(input *DeleteSnapshotScheduleInput) DeleteSnapshotScheduleRequest {
	op := &aws.Operation{
		Name:       opDeleteSnapshotSchedule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteSnapshotScheduleInput{}
	}

	req := c.newRequest(op, input, &DeleteSnapshotScheduleOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteSnapshotScheduleRequest{Request: req, Input: input, Copy: c.DeleteSnapshotScheduleRequest}
}

// DeleteSnapshotScheduleRequest is the request type for the
// DeleteSnapshotSchedule API operation.
type DeleteSnapshotScheduleRequest struct {
	*aws.Request
	Input *DeleteSnapshotScheduleInput
	Copy  func(*DeleteSnapshotScheduleInput) DeleteSnapshotScheduleRequest
}

// Send marshals and sends the DeleteSnapshotSchedule API request.
func (r DeleteSnapshotScheduleRequest) Send(ctx context.Context) (*DeleteSnapshotScheduleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSnapshotScheduleResponse{
		DeleteSnapshotScheduleOutput: r.Request.Data.(*DeleteSnapshotScheduleOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSnapshotScheduleResponse is the response type for the
// DeleteSnapshotSchedule API operation.
type DeleteSnapshotScheduleResponse struct {
	*DeleteSnapshotScheduleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSnapshotSchedule request.
func (r *DeleteSnapshotScheduleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
