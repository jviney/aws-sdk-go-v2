// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for CreateNetworkInterfacePermission.
type CreateNetworkInterfacePermissionInput struct {
	_ struct{} `type:"structure"`

	// The AWS account ID.
	AwsAccountId *string `type:"string"`

	// The AWS service. Currently not supported.
	AwsService *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the network interface.
	//
	// NetworkInterfaceId is a required field
	NetworkInterfaceId *string `type:"string" required:"true"`

	// The type of permission to grant.
	//
	// Permission is a required field
	Permission InterfacePermissionType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreateNetworkInterfacePermissionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateNetworkInterfacePermissionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateNetworkInterfacePermissionInput"}

	if s.NetworkInterfaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkInterfaceId"))
	}
	if len(s.Permission) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Permission"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of CreateNetworkInterfacePermission.
type CreateNetworkInterfacePermissionOutput struct {
	_ struct{} `type:"structure"`

	// Information about the permission for the network interface.
	InterfacePermission *NetworkInterfacePermission `locationName:"interfacePermission" type:"structure"`
}

// String returns the string representation
func (s CreateNetworkInterfacePermissionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateNetworkInterfacePermission = "CreateNetworkInterfacePermission"

// CreateNetworkInterfacePermissionRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Grants an AWS-authorized account permission to attach the specified network
// interface to an instance in their account.
//
// You can grant permission to a single AWS account only, and only one account
// at a time.
//
//    // Example sending a request using CreateNetworkInterfacePermissionRequest.
//    req := client.CreateNetworkInterfacePermissionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateNetworkInterfacePermission
func (c *Client) CreateNetworkInterfacePermissionRequest(input *CreateNetworkInterfacePermissionInput) CreateNetworkInterfacePermissionRequest {
	op := &aws.Operation{
		Name:       opCreateNetworkInterfacePermission,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateNetworkInterfacePermissionInput{}
	}

	req := c.newRequest(op, input, &CreateNetworkInterfacePermissionOutput{})

	return CreateNetworkInterfacePermissionRequest{Request: req, Input: input, Copy: c.CreateNetworkInterfacePermissionRequest}
}

// CreateNetworkInterfacePermissionRequest is the request type for the
// CreateNetworkInterfacePermission API operation.
type CreateNetworkInterfacePermissionRequest struct {
	*aws.Request
	Input *CreateNetworkInterfacePermissionInput
	Copy  func(*CreateNetworkInterfacePermissionInput) CreateNetworkInterfacePermissionRequest
}

// Send marshals and sends the CreateNetworkInterfacePermission API request.
func (r CreateNetworkInterfacePermissionRequest) Send(ctx context.Context) (*CreateNetworkInterfacePermissionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateNetworkInterfacePermissionResponse{
		CreateNetworkInterfacePermissionOutput: r.Request.Data.(*CreateNetworkInterfacePermissionOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateNetworkInterfacePermissionResponse is the response type for the
// CreateNetworkInterfacePermission API operation.
type CreateNetworkInterfacePermissionResponse struct {
	*CreateNetworkInterfacePermissionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateNetworkInterfacePermission request.
func (r *CreateNetworkInterfacePermissionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
