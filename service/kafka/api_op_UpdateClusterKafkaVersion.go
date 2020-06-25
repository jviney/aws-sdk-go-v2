// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// Request body for UpdateClusterKafkaVersion.
type UpdateClusterKafkaVersionInput struct {
	_ struct{} `type:"structure"`

	// ClusterArn is a required field
	ClusterArn *string `location:"uri" locationName:"clusterArn" type:"string" required:"true"`

	// Specifies the configuration to use for the brokers.
	ConfigurationInfo *ConfigurationInfo `locationName:"configurationInfo" type:"structure"`

	// Current cluster version.
	//
	// CurrentVersion is a required field
	CurrentVersion *string `locationName:"currentVersion" type:"string" required:"true"`

	// Target Kafka version.
	//
	// TargetKafkaVersion is a required field
	TargetKafkaVersion *string `locationName:"targetKafkaVersion" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateClusterKafkaVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateClusterKafkaVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateClusterKafkaVersionInput"}

	if s.ClusterArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterArn"))
	}

	if s.CurrentVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("CurrentVersion"))
	}

	if s.TargetKafkaVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetKafkaVersion"))
	}
	if s.ConfigurationInfo != nil {
		if err := s.ConfigurationInfo.Validate(); err != nil {
			invalidParams.AddNested("ConfigurationInfo", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateClusterKafkaVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ConfigurationInfo != nil {
		v := s.ConfigurationInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "configurationInfo", v, metadata)
	}
	if s.CurrentVersion != nil {
		v := *s.CurrentVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "currentVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TargetKafkaVersion != nil {
		v := *s.TargetKafkaVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "targetKafkaVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ClusterArn != nil {
		v := *s.ClusterArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "clusterArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Response body for UpdateClusterKafkaVersion.
type UpdateClusterKafkaVersionOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string `locationName:"clusterArn" type:"string"`

	// The Amazon Resource Name (ARN) of the cluster operation.
	ClusterOperationArn *string `locationName:"clusterOperationArn" type:"string"`
}

// String returns the string representation
func (s UpdateClusterKafkaVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateClusterKafkaVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
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

const opUpdateClusterKafkaVersion = "UpdateClusterKafkaVersion"

// UpdateClusterKafkaVersionRequest returns a request value for making API operation for
// Managed Streaming for Kafka.
//
// Updates the Apache Kafka version for the cluster.
//
//    // Example sending a request using UpdateClusterKafkaVersionRequest.
//    req := client.UpdateClusterKafkaVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/UpdateClusterKafkaVersion
func (c *Client) UpdateClusterKafkaVersionRequest(input *UpdateClusterKafkaVersionInput) UpdateClusterKafkaVersionRequest {
	op := &aws.Operation{
		Name:       opUpdateClusterKafkaVersion,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/clusters/{clusterArn}/version",
	}

	if input == nil {
		input = &UpdateClusterKafkaVersionInput{}
	}

	req := c.newRequest(op, input, &UpdateClusterKafkaVersionOutput{})

	return UpdateClusterKafkaVersionRequest{Request: req, Input: input, Copy: c.UpdateClusterKafkaVersionRequest}
}

// UpdateClusterKafkaVersionRequest is the request type for the
// UpdateClusterKafkaVersion API operation.
type UpdateClusterKafkaVersionRequest struct {
	*aws.Request
	Input *UpdateClusterKafkaVersionInput
	Copy  func(*UpdateClusterKafkaVersionInput) UpdateClusterKafkaVersionRequest
}

// Send marshals and sends the UpdateClusterKafkaVersion API request.
func (r UpdateClusterKafkaVersionRequest) Send(ctx context.Context) (*UpdateClusterKafkaVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateClusterKafkaVersionResponse{
		UpdateClusterKafkaVersionOutput: r.Request.Data.(*UpdateClusterKafkaVersionOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateClusterKafkaVersionResponse is the response type for the
// UpdateClusterKafkaVersion API operation.
type UpdateClusterKafkaVersionResponse struct {
	*UpdateClusterKafkaVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateClusterKafkaVersion request.
func (r *UpdateClusterKafkaVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
