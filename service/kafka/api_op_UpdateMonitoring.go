// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// Request body for UpdateMonitoring.
type UpdateMonitoringInput struct {
	_ struct{} `type:"structure"`

	// ClusterArn is a required field
	ClusterArn *string `location:"uri" locationName:"clusterArn" type:"string" required:"true"`

	// The version of cluster to update from. A successful operation will then generate
	// a new version.
	//
	// CurrentVersion is a required field
	CurrentVersion *string `locationName:"currentVersion" type:"string" required:"true"`

	// Specifies which Apache Kafka metrics Amazon MSK gathers and sends to Amazon
	// CloudWatch for this cluster.
	EnhancedMonitoring EnhancedMonitoring `locationName:"enhancedMonitoring" type:"string" enum:"true"`

	// LoggingInfo details.
	LoggingInfo *LoggingInfo `locationName:"loggingInfo" type:"structure"`

	// The settings for open monitoring.
	OpenMonitoring *OpenMonitoringInfo `locationName:"openMonitoring" type:"structure"`
}

// String returns the string representation
func (s UpdateMonitoringInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateMonitoringInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateMonitoringInput"}

	if s.ClusterArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterArn"))
	}

	if s.CurrentVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("CurrentVersion"))
	}
	if s.LoggingInfo != nil {
		if err := s.LoggingInfo.Validate(); err != nil {
			invalidParams.AddNested("LoggingInfo", err.(aws.ErrInvalidParams))
		}
	}
	if s.OpenMonitoring != nil {
		if err := s.OpenMonitoring.Validate(); err != nil {
			invalidParams.AddNested("OpenMonitoring", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateMonitoringInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CurrentVersion != nil {
		v := *s.CurrentVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "currentVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.EnhancedMonitoring) > 0 {
		v := s.EnhancedMonitoring

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enhancedMonitoring", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.LoggingInfo != nil {
		v := s.LoggingInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "loggingInfo", v, metadata)
	}
	if s.OpenMonitoring != nil {
		v := s.OpenMonitoring

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "openMonitoring", v, metadata)
	}
	if s.ClusterArn != nil {
		v := *s.ClusterArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "clusterArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Response body for UpdateMonitoring.
type UpdateMonitoringOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string `locationName:"clusterArn" type:"string"`

	// The Amazon Resource Name (ARN) of the cluster operation.
	ClusterOperationArn *string `locationName:"clusterOperationArn" type:"string"`
}

// String returns the string representation
func (s UpdateMonitoringOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateMonitoringOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClusterArn != nil {
		v := *s.ClusterArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clusterArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ClusterOperationArn != nil {
		v := *s.ClusterOperationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clusterOperationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateMonitoring = "UpdateMonitoring"

// UpdateMonitoringRequest returns a request value for making API operation for
// Managed Streaming for Kafka.
//
// Updates the monitoring settings for the cluster. You can use this operation
// to specify which Apache Kafka metrics you want Amazon MSK to send to Amazon
// CloudWatch. You can also specify settings for open monitoring with Prometheus.
//
//    // Example sending a request using UpdateMonitoringRequest.
//    req := client.UpdateMonitoringRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/UpdateMonitoring
func (c *Client) UpdateMonitoringRequest(input *UpdateMonitoringInput) UpdateMonitoringRequest {
	op := &aws.Operation{
		Name:       opUpdateMonitoring,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/clusters/{clusterArn}/monitoring",
	}

	if input == nil {
		input = &UpdateMonitoringInput{}
	}

	req := c.newRequest(op, input, &UpdateMonitoringOutput{})

	return UpdateMonitoringRequest{Request: req, Input: input, Copy: c.UpdateMonitoringRequest}
}

// UpdateMonitoringRequest is the request type for the
// UpdateMonitoring API operation.
type UpdateMonitoringRequest struct {
	*aws.Request
	Input *UpdateMonitoringInput
	Copy  func(*UpdateMonitoringInput) UpdateMonitoringRequest
}

// Send marshals and sends the UpdateMonitoring API request.
func (r UpdateMonitoringRequest) Send(ctx context.Context) (*UpdateMonitoringResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateMonitoringResponse{
		UpdateMonitoringOutput: r.Request.Data.(*UpdateMonitoringOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateMonitoringResponse is the response type for the
// UpdateMonitoring API operation.
type UpdateMonitoringResponse struct {
	*UpdateMonitoringOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateMonitoring request.
func (r *UpdateMonitoringResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
