// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeClustersInput struct {
	_ struct{} `type:"structure"`

	// A list of up to 100 cluster names or full cluster Amazon Resource Name (ARN)
	// entries. If you do not specify a cluster, the default cluster is assumed.
	Clusters []string `locationName:"clusters" type:"list"`

	// Whether to include additional information about your clusters in the response.
	// If this field is omitted, the attachments, statistics, and tags are not included.
	//
	// If ATTACHMENTS is specified, the attachments for the container instances
	// or tasks within the cluster are included.
	//
	// If SETTINGS is specified, the settings for the cluster are included.
	//
	// If STATISTICS is specified, the following additional information, separated
	// by launch type, is included:
	//
	//    * runningEC2TasksCount
	//
	//    * runningFargateTasksCount
	//
	//    * pendingEC2TasksCount
	//
	//    * pendingFargateTasksCount
	//
	//    * activeEC2ServiceCount
	//
	//    * activeFargateServiceCount
	//
	//    * drainingEC2ServiceCount
	//
	//    * drainingFargateServiceCount
	//
	// If TAGS is specified, the metadata tags associated with the cluster are included.
	Include []ClusterField `locationName:"include" type:"list"`
}

// String returns the string representation
func (s DescribeClustersInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeClustersOutput struct {
	_ struct{} `type:"structure"`

	// The list of clusters.
	Clusters []Cluster `locationName:"clusters" type:"list"`

	// Any failures associated with the call.
	Failures []Failure `locationName:"failures" type:"list"`
}

// String returns the string representation
func (s DescribeClustersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeClusters = "DescribeClusters"

// DescribeClustersRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Describes one or more of your clusters.
//
//    // Example sending a request using DescribeClustersRequest.
//    req := client.DescribeClustersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/DescribeClusters
func (c *Client) DescribeClustersRequest(input *DescribeClustersInput) DescribeClustersRequest {
	op := &aws.Operation{
		Name:       opDescribeClusters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeClustersInput{}
	}

	req := c.newRequest(op, input, &DescribeClustersOutput{})

	return DescribeClustersRequest{Request: req, Input: input, Copy: c.DescribeClustersRequest}
}

// DescribeClustersRequest is the request type for the
// DescribeClusters API operation.
type DescribeClustersRequest struct {
	*aws.Request
	Input *DescribeClustersInput
	Copy  func(*DescribeClustersInput) DescribeClustersRequest
}

// Send marshals and sends the DescribeClusters API request.
func (r DescribeClustersRequest) Send(ctx context.Context) (*DescribeClustersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeClustersResponse{
		DescribeClustersOutput: r.Request.Data.(*DescribeClustersOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeClustersResponse is the response type for the
// DescribeClusters API operation.
type DescribeClustersResponse struct {
	*DescribeClustersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeClusters request.
func (r *DescribeClustersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
