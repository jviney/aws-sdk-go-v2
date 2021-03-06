// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetAutoSnapshotsInput struct {
	_ struct{} `type:"structure"`

	// The name of the source instance or disk from which to get automatic snapshot
	// information.
	//
	// ResourceName is a required field
	ResourceName *string `locationName:"resourceName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetAutoSnapshotsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAutoSnapshotsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAutoSnapshotsInput"}

	if s.ResourceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetAutoSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that describe the automatic snapshots that are available
	// for the specified source instance or disk.
	AutoSnapshots []AutoSnapshotDetails `locationName:"autoSnapshots" type:"list"`

	// The name of the source instance or disk for the automatic snapshots.
	ResourceName *string `locationName:"resourceName" type:"string"`

	// The resource type (e.g., Instance or Disk).
	ResourceType ResourceType `locationName:"resourceType" type:"string" enum:"true"`
}

// String returns the string representation
func (s GetAutoSnapshotsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetAutoSnapshots = "GetAutoSnapshots"

// GetAutoSnapshotsRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns the available automatic snapshots for an instance or disk. For more
// information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-configuring-automatic-snapshots).
//
//    // Example sending a request using GetAutoSnapshotsRequest.
//    req := client.GetAutoSnapshotsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetAutoSnapshots
func (c *Client) GetAutoSnapshotsRequest(input *GetAutoSnapshotsInput) GetAutoSnapshotsRequest {
	op := &aws.Operation{
		Name:       opGetAutoSnapshots,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAutoSnapshotsInput{}
	}

	req := c.newRequest(op, input, &GetAutoSnapshotsOutput{})

	return GetAutoSnapshotsRequest{Request: req, Input: input, Copy: c.GetAutoSnapshotsRequest}
}

// GetAutoSnapshotsRequest is the request type for the
// GetAutoSnapshots API operation.
type GetAutoSnapshotsRequest struct {
	*aws.Request
	Input *GetAutoSnapshotsInput
	Copy  func(*GetAutoSnapshotsInput) GetAutoSnapshotsRequest
}

// Send marshals and sends the GetAutoSnapshots API request.
func (r GetAutoSnapshotsRequest) Send(ctx context.Context) (*GetAutoSnapshotsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAutoSnapshotsResponse{
		GetAutoSnapshotsOutput: r.Request.Data.(*GetAutoSnapshotsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAutoSnapshotsResponse is the response type for the
// GetAutoSnapshots API operation.
type GetAutoSnapshotsResponse struct {
	*GetAutoSnapshotsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAutoSnapshots request.
func (r *GetAutoSnapshotsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
