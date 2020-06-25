// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeEventCategoriesInput struct {
	_ struct{} `type:"structure"`

	// The source type, such as cluster or parameter group, to which the described
	// event categories apply.
	//
	// Valid values: cluster, cluster-snapshot, cluster-parameter-group, cluster-security-group,
	// and scheduled-action.
	SourceType *string `type:"string"`
}

// String returns the string representation
func (s DescribeEventCategoriesInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeEventCategoriesOutput struct {
	_ struct{} `type:"structure"`

	// A list of event categories descriptions.
	EventCategoriesMapList []EventCategoriesMap `locationNameList:"EventCategoriesMap" type:"list"`
}

// String returns the string representation
func (s DescribeEventCategoriesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEventCategories = "DescribeEventCategories"

// DescribeEventCategoriesRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Displays a list of event categories for all event source types, or for a
// specified source type. For a list of the event categories and source types,
// go to Amazon Redshift Event Notifications (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-event-notifications.html).
//
//    // Example sending a request using DescribeEventCategoriesRequest.
//    req := client.DescribeEventCategoriesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeEventCategories
func (c *Client) DescribeEventCategoriesRequest(input *DescribeEventCategoriesInput) DescribeEventCategoriesRequest {
	op := &aws.Operation{
		Name:       opDescribeEventCategories,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEventCategoriesInput{}
	}

	req := c.newRequest(op, input, &DescribeEventCategoriesOutput{})

	return DescribeEventCategoriesRequest{Request: req, Input: input, Copy: c.DescribeEventCategoriesRequest}
}

// DescribeEventCategoriesRequest is the request type for the
// DescribeEventCategories API operation.
type DescribeEventCategoriesRequest struct {
	*aws.Request
	Input *DescribeEventCategoriesInput
	Copy  func(*DescribeEventCategoriesInput) DescribeEventCategoriesRequest
}

// Send marshals and sends the DescribeEventCategories API request.
func (r DescribeEventCategoriesRequest) Send(ctx context.Context) (*DescribeEventCategoriesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEventCategoriesResponse{
		DescribeEventCategoriesOutput: r.Request.Data.(*DescribeEventCategoriesOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEventCategoriesResponse is the response type for the
// DescribeEventCategories API operation.
type DescribeEventCategoriesResponse struct {
	*DescribeEventCategoriesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEventCategories request.
func (r *DescribeEventCategoriesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
