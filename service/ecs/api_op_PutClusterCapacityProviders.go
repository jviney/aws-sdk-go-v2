// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type PutClusterCapacityProvidersInput struct {
	_ struct{} `type:"structure"`

	// The name of one or more capacity providers to associate with the cluster.
	//
	// If specifying a capacity provider that uses an Auto Scaling group, the capacity
	// provider must already be created. New capacity providers can be created with
	// the CreateCapacityProvider API operation.
	//
	// To use a AWS Fargate capacity provider, specify either the FARGATE or FARGATE_SPOT
	// capacity providers. The AWS Fargate capacity providers are available to all
	// accounts and only need to be associated with a cluster to be used.
	//
	// CapacityProviders is a required field
	CapacityProviders []string `locationName:"capacityProviders" type:"list" required:"true"`

	// The short name or full Amazon Resource Name (ARN) of the cluster to modify
	// the capacity provider settings for. If you do not specify a cluster, the
	// default cluster is assumed.
	//
	// Cluster is a required field
	Cluster *string `locationName:"cluster" type:"string" required:"true"`

	// The capacity provider strategy to use by default for the cluster.
	//
	// When creating a service or running a task on a cluster, if no capacity provider
	// or launch type is specified then the default capacity provider strategy for
	// the cluster is used.
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
	// DefaultCapacityProviderStrategy is a required field
	DefaultCapacityProviderStrategy []CapacityProviderStrategyItem `locationName:"defaultCapacityProviderStrategy" type:"list" required:"true"`
}

// String returns the string representation
func (s PutClusterCapacityProvidersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutClusterCapacityProvidersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutClusterCapacityProvidersInput"}

	if s.CapacityProviders == nil {
		invalidParams.Add(aws.NewErrParamRequired("CapacityProviders"))
	}

	if s.Cluster == nil {
		invalidParams.Add(aws.NewErrParamRequired("Cluster"))
	}

	if s.DefaultCapacityProviderStrategy == nil {
		invalidParams.Add(aws.NewErrParamRequired("DefaultCapacityProviderStrategy"))
	}
	if s.DefaultCapacityProviderStrategy != nil {
		for i, v := range s.DefaultCapacityProviderStrategy {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "DefaultCapacityProviderStrategy", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutClusterCapacityProvidersOutput struct {
	_ struct{} `type:"structure"`

	// A regional grouping of one or more container instances on which you can run
	// task requests. Each account receives a default cluster the first time you
	// use the Amazon ECS service, but you may also create other clusters. Clusters
	// may contain more than one instance type simultaneously.
	Cluster *Cluster `locationName:"cluster" type:"structure"`
}

// String returns the string representation
func (s PutClusterCapacityProvidersOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutClusterCapacityProviders = "PutClusterCapacityProviders"

// PutClusterCapacityProvidersRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Modifies the available capacity providers and the default capacity provider
// strategy for a cluster.
//
// You must specify both the available capacity providers and a default capacity
// provider strategy for the cluster. If the specified cluster has existing
// capacity providers associated with it, you must specify all existing capacity
// providers in addition to any new ones you want to add. Any existing capacity
// providers associated with a cluster that are omitted from a PutClusterCapacityProviders
// API call will be disassociated with the cluster. You can only disassociate
// an existing capacity provider from a cluster if it's not being used by any
// existing tasks.
//
// When creating a service or running a task on a cluster, if no capacity provider
// or launch type is specified, then the cluster's default capacity provider
// strategy is used. It is recommended to define a default capacity provider
// strategy for your cluster, however you may specify an empty array ([]) to
// bypass defining a default strategy.
//
//    // Example sending a request using PutClusterCapacityProvidersRequest.
//    req := client.PutClusterCapacityProvidersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/PutClusterCapacityProviders
func (c *Client) PutClusterCapacityProvidersRequest(input *PutClusterCapacityProvidersInput) PutClusterCapacityProvidersRequest {
	op := &aws.Operation{
		Name:       opPutClusterCapacityProviders,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutClusterCapacityProvidersInput{}
	}

	req := c.newRequest(op, input, &PutClusterCapacityProvidersOutput{})

	return PutClusterCapacityProvidersRequest{Request: req, Input: input, Copy: c.PutClusterCapacityProvidersRequest}
}

// PutClusterCapacityProvidersRequest is the request type for the
// PutClusterCapacityProviders API operation.
type PutClusterCapacityProvidersRequest struct {
	*aws.Request
	Input *PutClusterCapacityProvidersInput
	Copy  func(*PutClusterCapacityProvidersInput) PutClusterCapacityProvidersRequest
}

// Send marshals and sends the PutClusterCapacityProviders API request.
func (r PutClusterCapacityProvidersRequest) Send(ctx context.Context) (*PutClusterCapacityProvidersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutClusterCapacityProvidersResponse{
		PutClusterCapacityProvidersOutput: r.Request.Data.(*PutClusterCapacityProvidersOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutClusterCapacityProvidersResponse is the response type for the
// PutClusterCapacityProviders API operation.
type PutClusterCapacityProvidersResponse struct {
	*PutClusterCapacityProvidersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutClusterCapacityProviders request.
func (r *PutClusterCapacityProvidersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
