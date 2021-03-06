// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ModifyClusterSnapshotInput struct {
	_ struct{} `type:"structure"`

	// A Boolean option to override an exception if the retention period has already
	// passed.
	Force *bool `type:"boolean"`

	// The number of days that a manual snapshot is retained. If the value is -1,
	// the manual snapshot is retained indefinitely.
	//
	// If the manual snapshot falls outside of the new retention period, you can
	// specify the force option to immediately delete the snapshot.
	//
	// The value must be either -1 or an integer between 1 and 3,653.
	ManualSnapshotRetentionPeriod *int64 `type:"integer"`

	// The identifier of the snapshot whose setting you want to modify.
	//
	// SnapshotIdentifier is a required field
	SnapshotIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyClusterSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyClusterSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyClusterSnapshotInput"}

	if s.SnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyClusterSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// Describes a snapshot.
	Snapshot *Snapshot `type:"structure"`
}

// String returns the string representation
func (s ModifyClusterSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyClusterSnapshot = "ModifyClusterSnapshot"

// ModifyClusterSnapshotRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Modifies the settings for a snapshot.
//
// This exanmple modifies the manual retention period setting for a cluster
// snapshot.
//
//    // Example sending a request using ModifyClusterSnapshotRequest.
//    req := client.ModifyClusterSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/ModifyClusterSnapshot
func (c *Client) ModifyClusterSnapshotRequest(input *ModifyClusterSnapshotInput) ModifyClusterSnapshotRequest {
	op := &aws.Operation{
		Name:       opModifyClusterSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyClusterSnapshotInput{}
	}

	req := c.newRequest(op, input, &ModifyClusterSnapshotOutput{})

	return ModifyClusterSnapshotRequest{Request: req, Input: input, Copy: c.ModifyClusterSnapshotRequest}
}

// ModifyClusterSnapshotRequest is the request type for the
// ModifyClusterSnapshot API operation.
type ModifyClusterSnapshotRequest struct {
	*aws.Request
	Input *ModifyClusterSnapshotInput
	Copy  func(*ModifyClusterSnapshotInput) ModifyClusterSnapshotRequest
}

// Send marshals and sends the ModifyClusterSnapshot API request.
func (r ModifyClusterSnapshotRequest) Send(ctx context.Context) (*ModifyClusterSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyClusterSnapshotResponse{
		ModifyClusterSnapshotOutput: r.Request.Data.(*ModifyClusterSnapshotOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyClusterSnapshotResponse is the response type for the
// ModifyClusterSnapshot API operation.
type ModifyClusterSnapshotResponse struct {
	*ModifyClusterSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyClusterSnapshot request.
func (r *ModifyClusterSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
