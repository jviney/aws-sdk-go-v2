// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	"fmt"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateEnvironmentInput struct {
	_ struct{} `type:"structure"`

	// The name of the application that is associated with this environment.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// If specified, the environment attempts to use this value as the prefix for
	// the CNAME in your Elastic Beanstalk environment URL. If not specified, the
	// CNAME is generated automatically by appending a random alphanumeric string
	// to the environment name.
	CNAMEPrefix *string `min:"4" type:"string"`

	// Your description for this environment.
	Description *string `type:"string"`

	// A unique name for the environment.
	//
	// Constraint: Must be from 4 to 40 characters in length. The name can contain
	// only letters, numbers, and hyphens. It can't start or end with a hyphen.
	// This name must be unique within a region in your account. If the specified
	// name already exists in the region, Elastic Beanstalk returns an InvalidParameterValue
	// error.
	//
	// If you don't specify the CNAMEPrefix parameter, the environment name becomes
	// part of the CNAME, and therefore part of the visible URL for your application.
	EnvironmentName *string `min:"4" type:"string"`

	// The name of the group to which the target environment belongs. Specify a
	// group name only if the environment's name is specified in an environment
	// manifest and not with the environment name parameter. See Environment Manifest
	// (env.yaml) (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/environment-cfg-manifest.html)
	// for details.
	GroupName *string `min:"1" type:"string"`

	// If specified, AWS Elastic Beanstalk sets the specified configuration options
	// to the requested value in the configuration set for the new environment.
	// These override the values obtained from the solution stack or the configuration
	// template.
	OptionSettings []ConfigurationOptionSetting `type:"list"`

	// A list of custom user-defined configuration options to remove from the configuration
	// set for this new environment.
	OptionsToRemove []OptionSpecification `type:"list"`

	// The Amazon Resource Name (ARN) of the custom platform to use with the environment.
	// For more information, see Custom Platforms (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/custom-platforms.html)
	// in the AWS Elastic Beanstalk Developer Guide.
	//
	// If you specify PlatformArn, don't specify SolutionStackName.
	PlatformArn *string `type:"string"`

	// The name of an Elastic Beanstalk solution stack (platform version) to use
	// with the environment. If specified, Elastic Beanstalk sets the configuration
	// values to the default values associated with the specified solution stack.
	// For a list of current solution stacks, see Elastic Beanstalk Supported Platforms
	// (https://docs.aws.amazon.com/elasticbeanstalk/latest/platforms/platforms-supported.html)
	// in the AWS Elastic Beanstalk Platforms guide.
	//
	// If you specify SolutionStackName, don't specify PlatformArn or TemplateName.
	SolutionStackName *string `type:"string"`

	// Specifies the tags applied to resources in the environment.
	Tags []Tag `type:"list"`

	// The name of the Elastic Beanstalk configuration template to use with the
	// environment.
	//
	// If you specify TemplateName, then don't specify SolutionStackName.
	TemplateName *string `min:"1" type:"string"`

	// Specifies the tier to use in creating this environment. The environment tier
	// that you choose determines whether Elastic Beanstalk provisions resources
	// to support a web application that handles HTTP(S) requests or a web application
	// that handles background-processing tasks.
	Tier *EnvironmentTier `type:"structure"`

	// The name of the application version to deploy.
	//
	// Default: If not specified, Elastic Beanstalk attempts to deploy the sample
	// application.
	VersionLabel *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateEnvironmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateEnvironmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateEnvironmentInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}
	if s.CNAMEPrefix != nil && len(*s.CNAMEPrefix) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("CNAMEPrefix", 4))
	}
	if s.EnvironmentName != nil && len(*s.EnvironmentName) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("EnvironmentName", 4))
	}
	if s.GroupName != nil && len(*s.GroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupName", 1))
	}
	if s.TemplateName != nil && len(*s.TemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateName", 1))
	}
	if s.VersionLabel != nil && len(*s.VersionLabel) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VersionLabel", 1))
	}
	if s.OptionSettings != nil {
		for i, v := range s.OptionSettings {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "OptionSettings", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.OptionsToRemove != nil {
		for i, v := range s.OptionsToRemove {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "OptionsToRemove", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes the properties of an environment.
type CreateEnvironmentOutput struct {
	_ struct{} `type:"structure"`

	// Indicates if there is an in-progress environment configuration update or
	// application version deployment that you can cancel.
	//
	// true: There is an update in progress.
	//
	// false: There are no updates currently in progress.
	AbortableOperationInProgress *bool `type:"boolean"`

	// The name of the application associated with this environment.
	ApplicationName *string `min:"1" type:"string"`

	// The URL to the CNAME for this environment.
	CNAME *string `min:"1" type:"string"`

	// The creation date for this environment.
	DateCreated *time.Time `type:"timestamp"`

	// The last modified date for this environment.
	DateUpdated *time.Time `type:"timestamp"`

	// Describes this environment.
	Description *string `type:"string"`

	// For load-balanced, autoscaling environments, the URL to the LoadBalancer.
	// For single-instance environments, the IP address of the instance.
	EndpointURL *string `type:"string"`

	// The environment's Amazon Resource Name (ARN), which can be used in other
	// API requests that require an ARN.
	EnvironmentArn *string `type:"string"`

	// The ID of this environment.
	EnvironmentId *string `type:"string"`

	// A list of links to other environments in the same group.
	EnvironmentLinks []EnvironmentLink `type:"list"`

	// The name of this environment.
	EnvironmentName *string `min:"4" type:"string"`

	// Describes the health status of the environment. AWS Elastic Beanstalk indicates
	// the failure levels for a running environment:
	//
	//    * Red: Indicates the environment is not responsive. Occurs when three
	//    or more consecutive failures occur for an environment.
	//
	//    * Yellow: Indicates that something is wrong. Occurs when two consecutive
	//    failures occur for an environment.
	//
	//    * Green: Indicates the environment is healthy and fully functional.
	//
	//    * Grey: Default health for a new environment. The environment is not fully
	//    launched and health checks have not started or health checks are suspended
	//    during an UpdateEnvironment or RestartEnvironment request.
	//
	// Default: Grey
	Health EnvironmentHealth `type:"string" enum:"true"`

	// Returns the health status of the application running in your environment.
	// For more information, see Health Colors and Statuses (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/health-enhanced-status.html).
	HealthStatus EnvironmentHealthStatus `type:"string" enum:"true"`

	// The ARN of the platform version.
	PlatformArn *string `type:"string"`

	// The description of the AWS resources used by this environment.
	Resources *EnvironmentResourcesDescription `type:"structure"`

	// The name of the SolutionStack deployed with this environment.
	SolutionStackName *string `type:"string"`

	// The current operational status of the environment:
	//
	//    * Launching: Environment is in the process of initial deployment.
	//
	//    * Updating: Environment is in the process of updating its configuration
	//    settings or application version.
	//
	//    * Ready: Environment is available to have an action performed on it, such
	//    as update or terminate.
	//
	//    * Terminating: Environment is in the shut-down process.
	//
	//    * Terminated: Environment is not running.
	Status EnvironmentStatus `type:"string" enum:"true"`

	// The name of the configuration template used to originally launch this environment.
	TemplateName *string `min:"1" type:"string"`

	// Describes the current tier of this environment.
	Tier *EnvironmentTier `type:"structure"`

	// The application version deployed in this environment.
	VersionLabel *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateEnvironmentOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateEnvironment = "CreateEnvironment"

// CreateEnvironmentRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Launches an AWS Elastic Beanstalk environment for the specified application
// using the specified configuration.
//
//    // Example sending a request using CreateEnvironmentRequest.
//    req := client.CreateEnvironmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/CreateEnvironment
func (c *Client) CreateEnvironmentRequest(input *CreateEnvironmentInput) CreateEnvironmentRequest {
	op := &aws.Operation{
		Name:       opCreateEnvironment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateEnvironmentInput{}
	}

	req := c.newRequest(op, input, &CreateEnvironmentOutput{})

	return CreateEnvironmentRequest{Request: req, Input: input, Copy: c.CreateEnvironmentRequest}
}

// CreateEnvironmentRequest is the request type for the
// CreateEnvironment API operation.
type CreateEnvironmentRequest struct {
	*aws.Request
	Input *CreateEnvironmentInput
	Copy  func(*CreateEnvironmentInput) CreateEnvironmentRequest
}

// Send marshals and sends the CreateEnvironment API request.
func (r CreateEnvironmentRequest) Send(ctx context.Context) (*CreateEnvironmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateEnvironmentResponse{
		CreateEnvironmentOutput: r.Request.Data.(*CreateEnvironmentOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateEnvironmentResponse is the response type for the
// CreateEnvironment API operation.
type CreateEnvironmentResponse struct {
	*CreateEnvironmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateEnvironment request.
func (r *CreateEnvironmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
