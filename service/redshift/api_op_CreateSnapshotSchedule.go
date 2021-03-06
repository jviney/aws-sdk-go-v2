// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateSnapshotScheduleInput struct {
	_ struct{} `type:"structure"`

	DryRun *bool `type:"boolean"`

	NextInvocations *int64 `type:"integer"`

	// The definition of the snapshot schedule. The definition is made up of schedule
	// expressions, for example "cron(30 12 *)" or "rate(12 hours)".
	ScheduleDefinitions []string `locationNameList:"ScheduleDefinition" type:"list"`

	// The description of the snapshot schedule.
	ScheduleDescription *string `type:"string"`

	// A unique identifier for a snapshot schedule. Only alphanumeric characters
	// are allowed for the identifier.
	ScheduleIdentifier *string `type:"string"`

	// An optional set of tags you can use to search for the schedule.
	Tags []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s CreateSnapshotScheduleInput) String() string {
	return awsutil.Prettify(s)
}

// Describes a snapshot schedule. You can set a regular interval for creating
// snapshots of a cluster. You can also schedule snapshots for specific dates.
type CreateSnapshotScheduleOutput struct {
	_ struct{} `type:"structure"`

	// The number of clusters associated with the schedule.
	AssociatedClusterCount *int64 `type:"integer"`

	// A list of clusters associated with the schedule. A maximum of 100 clusters
	// is returned.
	AssociatedClusters []ClusterAssociatedToSchedule `locationNameList:"ClusterAssociatedToSchedule" type:"list"`

	NextInvocations []time.Time `locationNameList:"SnapshotTime" type:"list"`

	// A list of ScheduleDefinitions.
	ScheduleDefinitions []string `locationNameList:"ScheduleDefinition" type:"list"`

	// The description of the schedule.
	ScheduleDescription *string `type:"string"`

	// A unique identifier for the schedule.
	ScheduleIdentifier *string `type:"string"`

	// An optional set of tags describing the schedule.
	Tags []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s CreateSnapshotScheduleOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateSnapshotSchedule = "CreateSnapshotSchedule"

// CreateSnapshotScheduleRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Create a snapshot schedule that can be associated to a cluster and which
// overrides the default system backup schedule.
//
//    // Example sending a request using CreateSnapshotScheduleRequest.
//    req := client.CreateSnapshotScheduleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/CreateSnapshotSchedule
func (c *Client) CreateSnapshotScheduleRequest(input *CreateSnapshotScheduleInput) CreateSnapshotScheduleRequest {
	op := &aws.Operation{
		Name:       opCreateSnapshotSchedule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSnapshotScheduleInput{}
	}

	req := c.newRequest(op, input, &CreateSnapshotScheduleOutput{})

	return CreateSnapshotScheduleRequest{Request: req, Input: input, Copy: c.CreateSnapshotScheduleRequest}
}

// CreateSnapshotScheduleRequest is the request type for the
// CreateSnapshotSchedule API operation.
type CreateSnapshotScheduleRequest struct {
	*aws.Request
	Input *CreateSnapshotScheduleInput
	Copy  func(*CreateSnapshotScheduleInput) CreateSnapshotScheduleRequest
}

// Send marshals and sends the CreateSnapshotSchedule API request.
func (r CreateSnapshotScheduleRequest) Send(ctx context.Context) (*CreateSnapshotScheduleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSnapshotScheduleResponse{
		CreateSnapshotScheduleOutput: r.Request.Data.(*CreateSnapshotScheduleOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSnapshotScheduleResponse is the response type for the
// CreateSnapshotSchedule API operation.
type CreateSnapshotScheduleResponse struct {
	*CreateSnapshotScheduleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSnapshotSchedule request.
func (r *CreateSnapshotScheduleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
