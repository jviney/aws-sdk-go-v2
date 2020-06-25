// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeWorkspaceSnapshotsInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the WorkSpace.
	//
	// WorkspaceId is a required field
	WorkspaceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeWorkspaceSnapshotsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeWorkspaceSnapshotsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeWorkspaceSnapshotsInput"}

	if s.WorkspaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkspaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeWorkspaceSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the snapshots that can be used to rebuild a WorkSpace.
	// These snapshots include the user volume.
	RebuildSnapshots []Snapshot `type:"list"`

	// Information about the snapshots that can be used to restore a WorkSpace.
	// These snapshots include both the root volume and the user volume.
	RestoreSnapshots []Snapshot `type:"list"`
}

// String returns the string representation
func (s DescribeWorkspaceSnapshotsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeWorkspaceSnapshots = "DescribeWorkspaceSnapshots"

// DescribeWorkspaceSnapshotsRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Describes the snapshots for the specified WorkSpace.
//
//    // Example sending a request using DescribeWorkspaceSnapshotsRequest.
//    req := client.DescribeWorkspaceSnapshotsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/DescribeWorkspaceSnapshots
func (c *Client) DescribeWorkspaceSnapshotsRequest(input *DescribeWorkspaceSnapshotsInput) DescribeWorkspaceSnapshotsRequest {
	op := &aws.Operation{
		Name:       opDescribeWorkspaceSnapshots,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeWorkspaceSnapshotsInput{}
	}

	req := c.newRequest(op, input, &DescribeWorkspaceSnapshotsOutput{})

	return DescribeWorkspaceSnapshotsRequest{Request: req, Input: input, Copy: c.DescribeWorkspaceSnapshotsRequest}
}

// DescribeWorkspaceSnapshotsRequest is the request type for the
// DescribeWorkspaceSnapshots API operation.
type DescribeWorkspaceSnapshotsRequest struct {
	*aws.Request
	Input *DescribeWorkspaceSnapshotsInput
	Copy  func(*DescribeWorkspaceSnapshotsInput) DescribeWorkspaceSnapshotsRequest
}

// Send marshals and sends the DescribeWorkspaceSnapshots API request.
func (r DescribeWorkspaceSnapshotsRequest) Send(ctx context.Context) (*DescribeWorkspaceSnapshotsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeWorkspaceSnapshotsResponse{
		DescribeWorkspaceSnapshotsOutput: r.Request.Data.(*DescribeWorkspaceSnapshotsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeWorkspaceSnapshotsResponse is the response type for the
// DescribeWorkspaceSnapshots API operation.
type DescribeWorkspaceSnapshotsResponse struct {
	*DescribeWorkspaceSnapshotsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeWorkspaceSnapshots request.
func (r *DescribeWorkspaceSnapshotsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
