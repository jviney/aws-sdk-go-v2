// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeVolumeStatusInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The filters.
	//
	//    * action.code - The action code for the event (for example, enable-volume-io).
	//
	//    * action.description - A description of the action.
	//
	//    * action.event-id - The event ID associated with the action.
	//
	//    * availability-zone - The Availability Zone of the instance.
	//
	//    * event.description - A description of the event.
	//
	//    * event.event-id - The event ID.
	//
	//    * event.event-type - The event type (for io-enabled: passed | failed;
	//    for io-performance: io-performance:degraded | io-performance:severely-degraded
	//    | io-performance:stalled).
	//
	//    * event.not-after - The latest end time for the event.
	//
	//    * event.not-before - The earliest start time for the event.
	//
	//    * volume-status.details-name - The cause for volume-status.status (io-enabled
	//    | io-performance).
	//
	//    * volume-status.details-status - The status of volume-status.details-name
	//    (for io-enabled: passed | failed; for io-performance: normal | degraded
	//    | severely-degraded | stalled).
	//
	//    * volume-status.status - The status of the volume (ok | impaired | warning
	//    | insufficient-data).
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of volume results returned by DescribeVolumeStatus in
	// paginated output. When this parameter is used, the request only returns MaxResults
	// results in a single page along with a NextToken response element. The remaining
	// results of the initial request can be seen by sending another request with
	// the returned NextToken value. This value can be between 5 and 1000; if MaxResults
	// is given a value larger than 1000, only 1000 results are returned. If this
	// parameter is not used, then DescribeVolumeStatus returns all results. You
	// cannot specify this parameter and the volume IDs parameter in the same request.
	MaxResults *int64 `type:"integer"`

	// The NextToken value to include in a future DescribeVolumeStatus request.
	// When the results of the request exceed MaxResults, this value can be used
	// to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string `type:"string"`

	// The IDs of the volumes.
	//
	// Default: Describes all your volumes.
	VolumeIds []string `locationName:"VolumeId" locationNameList:"VolumeId" type:"list"`
}

// String returns the string representation
func (s DescribeVolumeStatusInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeVolumeStatusOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the status of the volumes.
	VolumeStatuses []VolumeStatusItem `locationName:"volumeStatusSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeVolumeStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeVolumeStatus = "DescribeVolumeStatus"

// DescribeVolumeStatusRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the status of the specified volumes. Volume status provides the
// result of the checks performed on your volumes to determine events that can
// impair the performance of your volumes. The performance of a volume can be
// affected if an issue occurs on the volume's underlying host. If the volume's
// underlying host experiences a power outage or system issue, after the system
// is restored, there could be data inconsistencies on the volume. Volume events
// notify you if this occurs. Volume actions notify you if any action needs
// to be taken in response to the event.
//
// The DescribeVolumeStatus operation provides the following information about
// the specified volumes:
//
// Status: Reflects the current status of the volume. The possible values are
// ok, impaired , warning, or insufficient-data. If all checks pass, the overall
// status of the volume is ok. If the check fails, the overall status is impaired.
// If the status is insufficient-data, then the checks may still be taking place
// on your volume at the time. We recommend that you retry the request. For
// more information about volume status, see Monitoring the Status of Your Volumes
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-volume-status.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// Events: Reflect the cause of a volume status and may require you to take
// action. For example, if your volume returns an impaired status, then the
// volume event might be potential-data-inconsistency. This means that your
// volume has been affected by an issue with the underlying host, has all I/O
// operations disabled, and may have inconsistent data.
//
// Actions: Reflect the actions you may have to take in response to an event.
// For example, if the status of the volume is impaired and the volume event
// shows potential-data-inconsistency, then the action shows enable-volume-io.
// This means that you may want to enable the I/O operations for the volume
// by calling the EnableVolumeIO action and then check the volume for data consistency.
//
// Volume status is based on the volume status checks, and does not reflect
// the volume state. Therefore, volume status does not indicate volumes in the
// error state (for example, when a volume is incapable of accepting I/O.)
//
//    // Example sending a request using DescribeVolumeStatusRequest.
//    req := client.DescribeVolumeStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeVolumeStatus
func (c *Client) DescribeVolumeStatusRequest(input *DescribeVolumeStatusInput) DescribeVolumeStatusRequest {
	op := &aws.Operation{
		Name:       opDescribeVolumeStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeVolumeStatusInput{}
	}

	req := c.newRequest(op, input, &DescribeVolumeStatusOutput{})

	return DescribeVolumeStatusRequest{Request: req, Input: input, Copy: c.DescribeVolumeStatusRequest}
}

// DescribeVolumeStatusRequest is the request type for the
// DescribeVolumeStatus API operation.
type DescribeVolumeStatusRequest struct {
	*aws.Request
	Input *DescribeVolumeStatusInput
	Copy  func(*DescribeVolumeStatusInput) DescribeVolumeStatusRequest
}

// Send marshals and sends the DescribeVolumeStatus API request.
func (r DescribeVolumeStatusRequest) Send(ctx context.Context) (*DescribeVolumeStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVolumeStatusResponse{
		DescribeVolumeStatusOutput: r.Request.Data.(*DescribeVolumeStatusOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeVolumeStatusRequestPaginator returns a paginator for DescribeVolumeStatus.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeVolumeStatusRequest(input)
//   p := ec2.NewDescribeVolumeStatusRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeVolumeStatusPaginator(req DescribeVolumeStatusRequest) DescribeVolumeStatusPaginator {
	return DescribeVolumeStatusPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeVolumeStatusInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeVolumeStatusPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeVolumeStatusPaginator struct {
	aws.Pager
}

func (p *DescribeVolumeStatusPaginator) CurrentPage() *DescribeVolumeStatusOutput {
	return p.Pager.CurrentPage().(*DescribeVolumeStatusOutput)
}

// DescribeVolumeStatusResponse is the response type for the
// DescribeVolumeStatus API operation.
type DescribeVolumeStatusResponse struct {
	*DescribeVolumeStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVolumeStatus request.
func (r *DescribeVolumeStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
