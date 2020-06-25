// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a CreateDeploymentConfig operation.
type CreateDeploymentConfigInput struct {
	_ struct{} `type:"structure"`

	// The destination platform type for the deployment (Lambda, Server, or ECS).
	ComputePlatform ComputePlatform `locationName:"computePlatform" type:"string" enum:"true"`

	// The name of the deployment configuration to create.
	//
	// DeploymentConfigName is a required field
	DeploymentConfigName *string `locationName:"deploymentConfigName" min:"1" type:"string" required:"true"`

	// The minimum number of healthy instances that should be available at any time
	// during the deployment. There are two parameters expected in the input: type
	// and value.
	//
	// The type parameter takes either of the following values:
	//
	//    * HOST_COUNT: The value parameter represents the minimum number of healthy
	//    instances as an absolute value.
	//
	//    * FLEET_PERCENT: The value parameter represents the minimum number of
	//    healthy instances as a percentage of the total number of instances in
	//    the deployment. If you specify FLEET_PERCENT, at the start of the deployment,
	//    AWS CodeDeploy converts the percentage to the equivalent number of instances
	//    and rounds up fractional instances.
	//
	// The value parameter takes an integer.
	//
	// For example, to set a minimum of 95% healthy instance, specify a type of
	// FLEET_PERCENT and a value of 95.
	MinimumHealthyHosts *MinimumHealthyHosts `locationName:"minimumHealthyHosts" type:"structure"`

	// The configuration that specifies how the deployment traffic is routed.
	TrafficRoutingConfig *TrafficRoutingConfig `locationName:"trafficRoutingConfig" type:"structure"`
}

// String returns the string representation
func (s CreateDeploymentConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDeploymentConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDeploymentConfigInput"}

	if s.DeploymentConfigName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentConfigName"))
	}
	if s.DeploymentConfigName != nil && len(*s.DeploymentConfigName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeploymentConfigName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a CreateDeploymentConfig operation.
type CreateDeploymentConfigOutput struct {
	_ struct{} `type:"structure"`

	// A unique deployment configuration ID.
	DeploymentConfigId *string `locationName:"deploymentConfigId" type:"string"`
}

// String returns the string representation
func (s CreateDeploymentConfigOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDeploymentConfig = "CreateDeploymentConfig"

// CreateDeploymentConfigRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Creates a deployment configuration.
//
//    // Example sending a request using CreateDeploymentConfigRequest.
//    req := client.CreateDeploymentConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/CreateDeploymentConfig
func (c *Client) CreateDeploymentConfigRequest(input *CreateDeploymentConfigInput) CreateDeploymentConfigRequest {
	op := &aws.Operation{
		Name:       opCreateDeploymentConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDeploymentConfigInput{}
	}

	req := c.newRequest(op, input, &CreateDeploymentConfigOutput{})

	return CreateDeploymentConfigRequest{Request: req, Input: input, Copy: c.CreateDeploymentConfigRequest}
}

// CreateDeploymentConfigRequest is the request type for the
// CreateDeploymentConfig API operation.
type CreateDeploymentConfigRequest struct {
	*aws.Request
	Input *CreateDeploymentConfigInput
	Copy  func(*CreateDeploymentConfigInput) CreateDeploymentConfigRequest
}

// Send marshals and sends the CreateDeploymentConfig API request.
func (r CreateDeploymentConfigRequest) Send(ctx context.Context) (*CreateDeploymentConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDeploymentConfigResponse{
		CreateDeploymentConfigOutput: r.Request.Data.(*CreateDeploymentConfigOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDeploymentConfigResponse is the response type for the
// CreateDeploymentConfig API operation.
type CreateDeploymentConfigResponse struct {
	*CreateDeploymentConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDeploymentConfig request.
func (r *CreateDeploymentConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
