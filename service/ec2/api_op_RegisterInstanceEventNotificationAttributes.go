// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type RegisterInstanceEventNotificationAttributesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// Information about the tag keys to register.
	InstanceTagAttribute *RegisterInstanceTagAttributeRequest `type:"structure"`
}

// String returns the string representation
func (s RegisterInstanceEventNotificationAttributesInput) String() string {
	return awsutil.Prettify(s)
}

type RegisterInstanceEventNotificationAttributesOutput struct {
	_ struct{} `type:"structure"`

	// The resulting set of tag keys.
	InstanceTagAttribute *InstanceTagNotificationAttribute `locationName:"instanceTagAttribute" type:"structure"`
}

// String returns the string representation
func (s RegisterInstanceEventNotificationAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opRegisterInstanceEventNotificationAttributes = "RegisterInstanceEventNotificationAttributes"

// RegisterInstanceEventNotificationAttributesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Registers a set of tag keys to include in scheduled event notifications for
// your resources.
//
// To remove tags, use .
//
//    // Example sending a request using RegisterInstanceEventNotificationAttributesRequest.
//    req := client.RegisterInstanceEventNotificationAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RegisterInstanceEventNotificationAttributes
func (c *Client) RegisterInstanceEventNotificationAttributesRequest(input *RegisterInstanceEventNotificationAttributesInput) RegisterInstanceEventNotificationAttributesRequest {
	op := &aws.Operation{
		Name:       opRegisterInstanceEventNotificationAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterInstanceEventNotificationAttributesInput{}
	}

	req := c.newRequest(op, input, &RegisterInstanceEventNotificationAttributesOutput{})

	return RegisterInstanceEventNotificationAttributesRequest{Request: req, Input: input, Copy: c.RegisterInstanceEventNotificationAttributesRequest}
}

// RegisterInstanceEventNotificationAttributesRequest is the request type for the
// RegisterInstanceEventNotificationAttributes API operation.
type RegisterInstanceEventNotificationAttributesRequest struct {
	*aws.Request
	Input *RegisterInstanceEventNotificationAttributesInput
	Copy  func(*RegisterInstanceEventNotificationAttributesInput) RegisterInstanceEventNotificationAttributesRequest
}

// Send marshals and sends the RegisterInstanceEventNotificationAttributes API request.
func (r RegisterInstanceEventNotificationAttributesRequest) Send(ctx context.Context) (*RegisterInstanceEventNotificationAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterInstanceEventNotificationAttributesResponse{
		RegisterInstanceEventNotificationAttributesOutput: r.Request.Data.(*RegisterInstanceEventNotificationAttributesOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterInstanceEventNotificationAttributesResponse is the response type for the
// RegisterInstanceEventNotificationAttributes API operation.
type RegisterInstanceEventNotificationAttributesResponse struct {
	*RegisterInstanceEventNotificationAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterInstanceEventNotificationAttributes request.
func (r *RegisterInstanceEventNotificationAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
