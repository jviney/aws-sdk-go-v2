// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DeleteDynamicThingGroupInput struct {
	_ struct{} `type:"structure"`

	// The expected version of the dynamic thing group to delete.
	ExpectedVersion *int64 `location:"querystring" locationName:"expectedVersion" type:"long"`

	// The name of the dynamic thing group to delete.
	//
	// ThingGroupName is a required field
	ThingGroupName *string `location:"uri" locationName:"thingGroupName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDynamicThingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDynamicThingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDynamicThingGroupInput"}

	if s.ThingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingGroupName"))
	}
	if s.ThingGroupName != nil && len(*s.ThingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDynamicThingGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ThingGroupName != nil {
		v := *s.ThingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ExpectedVersion != nil {
		v := *s.ExpectedVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "expectedVersion", protocol.Int64Value(v), metadata)
	}
	return nil
}

type DeleteDynamicThingGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDynamicThingGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDynamicThingGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteDynamicThingGroup = "DeleteDynamicThingGroup"

// DeleteDynamicThingGroupRequest returns a request value for making API operation for
// AWS IoT.
//
// Deletes a dynamic thing group.
//
//    // Example sending a request using DeleteDynamicThingGroupRequest.
//    req := client.DeleteDynamicThingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteDynamicThingGroupRequest(input *DeleteDynamicThingGroupInput) DeleteDynamicThingGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteDynamicThingGroup,
		HTTPMethod: "DELETE",
		HTTPPath:   "/dynamic-thing-groups/{thingGroupName}",
	}

	if input == nil {
		input = &DeleteDynamicThingGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteDynamicThingGroupOutput{})

	return DeleteDynamicThingGroupRequest{Request: req, Input: input, Copy: c.DeleteDynamicThingGroupRequest}
}

// DeleteDynamicThingGroupRequest is the request type for the
// DeleteDynamicThingGroup API operation.
type DeleteDynamicThingGroupRequest struct {
	*aws.Request
	Input *DeleteDynamicThingGroupInput
	Copy  func(*DeleteDynamicThingGroupInput) DeleteDynamicThingGroupRequest
}

// Send marshals and sends the DeleteDynamicThingGroup API request.
func (r DeleteDynamicThingGroupRequest) Send(ctx context.Context) (*DeleteDynamicThingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDynamicThingGroupResponse{
		DeleteDynamicThingGroupOutput: r.Request.Data.(*DeleteDynamicThingGroupOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDynamicThingGroupResponse is the response type for the
// DeleteDynamicThingGroup API operation.
type DeleteDynamicThingGroupResponse struct {
	*DeleteDynamicThingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDynamicThingGroup request.
func (r *DeleteDynamicThingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
