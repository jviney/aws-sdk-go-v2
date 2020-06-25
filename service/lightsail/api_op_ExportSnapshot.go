// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ExportSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The name of the instance or disk snapshot to be exported to Amazon EC2.
	//
	// SourceSnapshotName is a required field
	SourceSnapshotName *string `locationName:"sourceSnapshotName" type:"string" required:"true"`
}

// String returns the string representation
func (s ExportSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ExportSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ExportSnapshotInput"}

	if s.SourceSnapshotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceSnapshotName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ExportSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected
	// by the request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s ExportSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opExportSnapshot = "ExportSnapshot"

// ExportSnapshotRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Exports an Amazon Lightsail instance or block storage disk snapshot to Amazon
// Elastic Compute Cloud (Amazon EC2). This operation results in an export snapshot
// record that can be used with the create cloud formation stack operation to
// create new Amazon EC2 instances.
//
// Exported instance snapshots appear in Amazon EC2 as Amazon Machine Images
// (AMIs), and the instance system disk appears as an Amazon Elastic Block Store
// (Amazon EBS) volume. Exported disk snapshots appear in Amazon EC2 as Amazon
// EBS volumes. Snapshots are exported to the same Amazon Web Services Region
// in Amazon EC2 as the source Lightsail snapshot.
//
// The export snapshot operation supports tag-based access control via resource
// tags applied to the resource identified by source snapshot name. For more
// information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
// Use the get instance snapshots or get disk snapshots operations to get a
// list of snapshots that you can export to Amazon EC2.
//
//    // Example sending a request using ExportSnapshotRequest.
//    req := client.ExportSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/ExportSnapshot
func (c *Client) ExportSnapshotRequest(input *ExportSnapshotInput) ExportSnapshotRequest {
	op := &aws.Operation{
		Name:       opExportSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ExportSnapshotInput{}
	}

	req := c.newRequest(op, input, &ExportSnapshotOutput{})

	return ExportSnapshotRequest{Request: req, Input: input, Copy: c.ExportSnapshotRequest}
}

// ExportSnapshotRequest is the request type for the
// ExportSnapshot API operation.
type ExportSnapshotRequest struct {
	*aws.Request
	Input *ExportSnapshotInput
	Copy  func(*ExportSnapshotInput) ExportSnapshotRequest
}

// Send marshals and sends the ExportSnapshot API request.
func (r ExportSnapshotRequest) Send(ctx context.Context) (*ExportSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ExportSnapshotResponse{
		ExportSnapshotOutput: r.Request.Data.(*ExportSnapshotOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ExportSnapshotResponse is the response type for the
// ExportSnapshot API operation.
type ExportSnapshotResponse struct {
	*ExportSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ExportSnapshot request.
func (r *ExportSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
