// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ModifyVpcEndpointServicePermissionsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Names (ARN) of one or more principals. Permissions are
	// granted to the principals in this list. To grant permissions to all principals,
	// specify an asterisk (*).
	AddAllowedPrincipals []string `locationNameList:"item" type:"list"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The Amazon Resource Names (ARN) of one or more principals. Permissions are
	// revoked for principals in this list.
	RemoveAllowedPrincipals []string `locationNameList:"item" type:"list"`

	// The ID of the service.
	//
	// ServiceId is a required field
	ServiceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyVpcEndpointServicePermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVpcEndpointServicePermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyVpcEndpointServicePermissionsInput"}

	if s.ServiceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyVpcEndpointServicePermissionsOutput struct {
	_ struct{} `type:"structure"`

	// Returns true if the request succeeds; otherwise, it returns an error.
	ReturnValue *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s ModifyVpcEndpointServicePermissionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyVpcEndpointServicePermissions = "ModifyVpcEndpointServicePermissions"

// ModifyVpcEndpointServicePermissionsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the permissions for your VPC endpoint service (https://docs.aws.amazon.com/vpc/latest/userguide/endpoint-service.html).
// You can add or remove permissions for service consumers (IAM users, IAM roles,
// and AWS accounts) to connect to your endpoint service.
//
// If you grant permissions to all principals, the service is public. Any users
// who know the name of a public service can send a request to attach an endpoint.
// If the service does not require manual approval, attachments are automatically
// approved.
//
//    // Example sending a request using ModifyVpcEndpointServicePermissionsRequest.
//    req := client.ModifyVpcEndpointServicePermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyVpcEndpointServicePermissions
func (c *Client) ModifyVpcEndpointServicePermissionsRequest(input *ModifyVpcEndpointServicePermissionsInput) ModifyVpcEndpointServicePermissionsRequest {
	op := &aws.Operation{
		Name:       opModifyVpcEndpointServicePermissions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyVpcEndpointServicePermissionsInput{}
	}

	req := c.newRequest(op, input, &ModifyVpcEndpointServicePermissionsOutput{})

	return ModifyVpcEndpointServicePermissionsRequest{Request: req, Input: input, Copy: c.ModifyVpcEndpointServicePermissionsRequest}
}

// ModifyVpcEndpointServicePermissionsRequest is the request type for the
// ModifyVpcEndpointServicePermissions API operation.
type ModifyVpcEndpointServicePermissionsRequest struct {
	*aws.Request
	Input *ModifyVpcEndpointServicePermissionsInput
	Copy  func(*ModifyVpcEndpointServicePermissionsInput) ModifyVpcEndpointServicePermissionsRequest
}

// Send marshals and sends the ModifyVpcEndpointServicePermissions API request.
func (r ModifyVpcEndpointServicePermissionsRequest) Send(ctx context.Context) (*ModifyVpcEndpointServicePermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyVpcEndpointServicePermissionsResponse{
		ModifyVpcEndpointServicePermissionsOutput: r.Request.Data.(*ModifyVpcEndpointServicePermissionsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyVpcEndpointServicePermissionsResponse is the response type for the
// ModifyVpcEndpointServicePermissions API operation.
type ModifyVpcEndpointServicePermissionsResponse struct {
	*ModifyVpcEndpointServicePermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyVpcEndpointServicePermissions request.
func (r *ModifyVpcEndpointServicePermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
