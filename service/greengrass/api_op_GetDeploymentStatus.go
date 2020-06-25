// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetDeploymentStatusInput struct {
	_ struct{} `type:"structure"`

	// DeploymentId is a required field
	DeploymentId *string `location:"uri" locationName:"DeploymentId" type:"string" required:"true"`

	// GroupId is a required field
	GroupId *string `location:"uri" locationName:"GroupId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDeploymentStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDeploymentStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDeploymentStatusInput"}

	if s.DeploymentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentId"))
	}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDeploymentStatusInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DeploymentId != nil {
		v := *s.DeploymentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DeploymentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GroupId != nil {
		v := *s.GroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about the status of a deployment for a group.
type GetDeploymentStatusOutput struct {
	_ struct{} `type:"structure"`

	// The status of the deployment: ''InProgress'', ''Building'', ''Success'',
	// or ''Failure''.
	DeploymentStatus *string `type:"string"`

	// The type of the deployment.
	DeploymentType DeploymentType `type:"string" enum:"true"`

	// Error details
	ErrorDetails []ErrorDetail `type:"list"`

	// Error message
	ErrorMessage *string `type:"string"`

	// The time, in milliseconds since the epoch, when the deployment status was
	// updated.
	UpdatedAt *string `type:"string"`
}

// String returns the string representation
func (s GetDeploymentStatusOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDeploymentStatusOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DeploymentStatus != nil {
		v := *s.DeploymentStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeploymentStatus", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.DeploymentType) > 0 {
		v := s.DeploymentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeploymentType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ErrorDetails != nil {
		v := s.ErrorDetails

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ErrorDetails", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.ErrorMessage != nil {
		v := *s.ErrorMessage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ErrorMessage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UpdatedAt != nil {
		v := *s.UpdatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "UpdatedAt", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetDeploymentStatus = "GetDeploymentStatus"

// GetDeploymentStatusRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Returns the status of a deployment.
//
//    // Example sending a request using GetDeploymentStatusRequest.
//    req := client.GetDeploymentStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetDeploymentStatus
func (c *Client) GetDeploymentStatusRequest(input *GetDeploymentStatusInput) GetDeploymentStatusRequest {
	op := &aws.Operation{
		Name:       opGetDeploymentStatus,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/groups/{GroupId}/deployments/{DeploymentId}/status",
	}

	if input == nil {
		input = &GetDeploymentStatusInput{}
	}

	req := c.newRequest(op, input, &GetDeploymentStatusOutput{})

	return GetDeploymentStatusRequest{Request: req, Input: input, Copy: c.GetDeploymentStatusRequest}
}

// GetDeploymentStatusRequest is the request type for the
// GetDeploymentStatus API operation.
type GetDeploymentStatusRequest struct {
	*aws.Request
	Input *GetDeploymentStatusInput
	Copy  func(*GetDeploymentStatusInput) GetDeploymentStatusRequest
}

// Send marshals and sends the GetDeploymentStatus API request.
func (r GetDeploymentStatusRequest) Send(ctx context.Context) (*GetDeploymentStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDeploymentStatusResponse{
		GetDeploymentStatusOutput: r.Request.Data.(*GetDeploymentStatusOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDeploymentStatusResponse is the response type for the
// GetDeploymentStatus API operation.
type GetDeploymentStatusResponse struct {
	*GetDeploymentStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDeploymentStatus request.
func (r *GetDeploymentStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
