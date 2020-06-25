// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type UpdateServiceInput struct {
	_ struct{} `type:"structure"`

	// The capacity provider strategy to update the service to use.
	//
	// If the service is using the default capacity provider strategy for the cluster,
	// the service can be updated to use one or more capacity providers as opposed
	// to the default capacity provider strategy. However, when a service is using
	// a capacity provider strategy that is not the default capacity provider strategy,
	// the service cannot be updated to use the cluster's default capacity provider
	// strategy.
	//
	// A capacity provider strategy consists of one or more capacity providers along
	// with the base and weight to assign to them. A capacity provider must be associated
	// with the cluster to be used in a capacity provider strategy. The PutClusterCapacityProviders
	// API is used to associate a capacity provider with a cluster. Only capacity
	// providers with an ACTIVE or UPDATING status can be used.
	//
	// If specifying a capacity provider that uses an Auto Scaling group, the capacity
	// provider must already be created. New capacity providers can be created with
	// the CreateCapacityProvider API operation.
	//
	// To use a AWS Fargate capacity provider, specify either the FARGATE or FARGATE_SPOT
	// capacity providers. The AWS Fargate capacity providers are available to all
	// accounts and only need to be associated with a cluster to be used.
	//
	// The PutClusterCapacityProviders API operation is used to update the list
	// of available capacity providers for a cluster after the cluster is created.
	CapacityProviderStrategy []CapacityProviderStrategyItem `locationName:"capacityProviderStrategy" type:"list"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that your
	// service is running on. If you do not specify a cluster, the default cluster
	// is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// Optional deployment parameters that control how many tasks run during the
	// deployment and the ordering of stopping and starting tasks.
	DeploymentConfiguration *DeploymentConfiguration `locationName:"deploymentConfiguration" type:"structure"`

	// The number of instantiations of the task to place and keep running in your
	// service.
	DesiredCount *int64 `locationName:"desiredCount" type:"integer"`

	// Whether to force a new deployment of the service. Deployments are not forced
	// by default. You can use this option to trigger a new deployment with no service
	// definition changes. For example, you can update a service's tasks to use
	// a newer Docker image with the same image/tag combination (my_image:latest)
	// or to roll Fargate tasks onto a newer platform version.
	ForceNewDeployment *bool `locationName:"forceNewDeployment" type:"boolean"`

	// The period of time, in seconds, that the Amazon ECS service scheduler should
	// ignore unhealthy Elastic Load Balancing target health checks after a task
	// has first started. This is only valid if your service is configured to use
	// a load balancer. If your service's tasks take a while to start and respond
	// to Elastic Load Balancing health checks, you can specify a health check grace
	// period of up to 2,147,483,647 seconds. During that time, the Amazon ECS service
	// scheduler ignores the Elastic Load Balancing health check status. This grace
	// period can prevent the ECS service scheduler from marking tasks as unhealthy
	// and stopping them before they have time to come up.
	HealthCheckGracePeriodSeconds *int64 `locationName:"healthCheckGracePeriodSeconds" type:"integer"`

	// An object representing the network configuration for a task or service.
	NetworkConfiguration *NetworkConfiguration `locationName:"networkConfiguration" type:"structure"`

	// An array of task placement constraint objects to update the service to use.
	// If no value is specified, the existing placement constraints for the service
	// will remain unchanged. If this value is specified, it will override any existing
	// placement constraints defined for the service. To remove all existing placement
	// constraints, specify an empty array.
	//
	// You can specify a maximum of 10 constraints per task (this limit includes
	// constraints in the task definition and those specified at runtime).
	PlacementConstraints []PlacementConstraint `locationName:"placementConstraints" type:"list"`

	// The task placement strategy objects to update the service to use. If no value
	// is specified, the existing placement strategy for the service will remain
	// unchanged. If this value is specified, it will override the existing placement
	// strategy defined for the service. To remove an existing placement strategy,
	// specify an empty object.
	//
	// You can specify a maximum of five strategy rules per service.
	PlacementStrategy []PlacementStrategy `locationName:"placementStrategy" type:"list"`

	// The platform version on which your tasks in the service are running. A platform
	// version is only specified for tasks using the Fargate launch type. If a platform
	// version is not specified, the LATEST platform version is used by default.
	// For more information, see AWS Fargate Platform Versions (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	PlatformVersion *string `locationName:"platformVersion" type:"string"`

	// The name of the service to update.
	//
	// Service is a required field
	Service *string `locationName:"service" type:"string" required:"true"`

	// The family and revision (family:revision) or full ARN of the task definition
	// to run in your service. If a revision is not specified, the latest ACTIVE
	// revision is used. If you modify the task definition with UpdateService, Amazon
	// ECS spawns a task with the new version of the task definition and then stops
	// an old task after the new version is running.
	TaskDefinition *string `locationName:"taskDefinition" type:"string"`
}

// String returns the string representation
func (s UpdateServiceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateServiceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateServiceInput"}

	if s.Service == nil {
		invalidParams.Add(aws.NewErrParamRequired("Service"))
	}
	if s.CapacityProviderStrategy != nil {
		for i, v := range s.CapacityProviderStrategy {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "CapacityProviderStrategy", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			invalidParams.AddNested("NetworkConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateServiceOutput struct {
	_ struct{} `type:"structure"`

	// The full description of your service following the update call.
	Service *Service `locationName:"service" type:"structure"`
}

// String returns the string representation
func (s UpdateServiceOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateService = "UpdateService"

// UpdateServiceRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
//
// Updating the task placement strategies and constraints on an Amazon ECS service
// remains in preview and is a Beta Service as defined by and subject to the
// Beta Service Participation Service Terms located at https://aws.amazon.com/service-terms
// (https://aws.amazon.com/service-terms) ("Beta Terms"). These Beta Terms apply
// to your participation in this preview.
//
// Modifies the parameters of a service.
//
// For services using the rolling update (ECS) deployment controller, the desired
// count, deployment configuration, network configuration, task placement constraints
// and strategies, or task definition used can be updated.
//
// For services using the blue/green (CODE_DEPLOY) deployment controller, only
// the desired count, deployment configuration, task placement constraints and
// strategies, and health check grace period can be updated using this API.
// If the network configuration, platform version, or task definition need to
// be updated, a new AWS CodeDeploy deployment should be created. For more information,
// see CreateDeployment (https://docs.aws.amazon.com/codedeploy/latest/APIReference/API_CreateDeployment.html)
// in the AWS CodeDeploy API Reference.
//
// For services using an external deployment controller, you can update only
// the desired count, task placement constraints and strategies, and health
// check grace period using this API. If the launch type, load balancer, network
// configuration, platform version, or task definition need to be updated, you
// should create a new task set. For more information, see CreateTaskSet.
//
// You can add to or subtract from the number of instantiations of a task definition
// in a service by specifying the cluster that the service is running in and
// a new desiredCount parameter.
//
// If you have updated the Docker image of your application, you can create
// a new task definition with that image and deploy it to your service. The
// service scheduler uses the minimum healthy percent and maximum percent parameters
// (in the service's deployment configuration) to determine the deployment strategy.
//
// If your updated Docker image uses the same tag as what is in the existing
// task definition for your service (for example, my_image:latest), you do not
// need to create a new revision of your task definition. You can update the
// service using the forceNewDeployment option. The new tasks launched by the
// deployment pull the current image/tag combination from your repository when
// they start.
//
// You can also update the deployment configuration of a service. When a deployment
// is triggered by updating the task definition of a service, the service scheduler
// uses the deployment configuration parameters, minimumHealthyPercent and maximumPercent,
// to determine the deployment strategy.
//
//    * If minimumHealthyPercent is below 100%, the scheduler can ignore desiredCount
//    temporarily during a deployment. For example, if desiredCount is four
//    tasks, a minimum of 50% allows the scheduler to stop two existing tasks
//    before starting two new tasks. Tasks for services that do not use a load
//    balancer are considered healthy if they are in the RUNNING state. Tasks
//    for services that use a load balancer are considered healthy if they are
//    in the RUNNING state and the container instance they are hosted on is
//    reported as healthy by the load balancer.
//
//    * The maximumPercent parameter represents an upper limit on the number
//    of running tasks during a deployment, which enables you to define the
//    deployment batch size. For example, if desiredCount is four tasks, a maximum
//    of 200% starts four new tasks before stopping the four older tasks (provided
//    that the cluster resources required to do this are available).
//
// When UpdateService stops a task during a deployment, the equivalent of docker
// stop is issued to the containers running in the task. This results in a SIGTERM
// and a 30-second timeout, after which SIGKILL is sent and the containers are
// forcibly stopped. If the container handles the SIGTERM gracefully and exits
// within 30 seconds from receiving it, no SIGKILL is sent.
//
// When the service scheduler launches new tasks, it determines task placement
// in your cluster with the following logic:
//
//    * Determine which of the container instances in your cluster can support
//    your service's task definition (for example, they have the required CPU,
//    memory, ports, and container instance attributes).
//
//    * By default, the service scheduler attempts to balance tasks across Availability
//    Zones in this manner (although you can choose a different placement strategy):
//    Sort the valid container instances by the fewest number of running tasks
//    for this service in the same Availability Zone as the instance. For example,
//    if zone A has one running service task and zones B and C each have zero,
//    valid container instances in either zone B or C are considered optimal
//    for placement. Place the new service task on a valid container instance
//    in an optimal Availability Zone (based on the previous steps), favoring
//    container instances with the fewest number of running tasks for this service.
//
// When the service scheduler stops running tasks, it attempts to maintain balance
// across the Availability Zones in your cluster using the following logic:
//
//    * Sort the container instances by the largest number of running tasks
//    for this service in the same Availability Zone as the instance. For example,
//    if zone A has one running service task and zones B and C each have two,
//    container instances in either zone B or C are considered optimal for termination.
//
//    * Stop the task on a container instance in an optimal Availability Zone
//    (based on the previous steps), favoring container instances with the largest
//    number of running tasks for this service.
//
//    // Example sending a request using UpdateServiceRequest.
//    req := client.UpdateServiceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/UpdateService
func (c *Client) UpdateServiceRequest(input *UpdateServiceInput) UpdateServiceRequest {
	op := &aws.Operation{
		Name:       opUpdateService,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateServiceInput{}
	}

	req := c.newRequest(op, input, &UpdateServiceOutput{})

	return UpdateServiceRequest{Request: req, Input: input, Copy: c.UpdateServiceRequest}
}

// UpdateServiceRequest is the request type for the
// UpdateService API operation.
type UpdateServiceRequest struct {
	*aws.Request
	Input *UpdateServiceInput
	Copy  func(*UpdateServiceInput) UpdateServiceRequest
}

// Send marshals and sends the UpdateService API request.
func (r UpdateServiceRequest) Send(ctx context.Context) (*UpdateServiceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateServiceResponse{
		UpdateServiceOutput: r.Request.Data.(*UpdateServiceOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateServiceResponse is the response type for the
// UpdateService API operation.
type UpdateServiceResponse struct {
	*UpdateServiceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateService request.
func (r *UpdateServiceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
