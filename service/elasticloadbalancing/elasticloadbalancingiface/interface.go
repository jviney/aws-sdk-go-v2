// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package elasticloadbalancingiface provides an interface to enable mocking the Elastic Load Balancing service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package elasticloadbalancingiface

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/service/elasticloadbalancing"
)

// ClientAPI provides an interface to enable mocking the
// elasticloadbalancing.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Elastic Load Balancing.
//    func myFunc(svc elasticloadbalancingiface.ClientAPI) bool {
//        // Make svc.AddTags request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := elasticloadbalancing.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        elasticloadbalancingiface.ClientPI
//    }
//    func (m *mockClientClient) AddTags(input *elasticloadbalancing.AddTagsInput) (*elasticloadbalancing.AddTagsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	AddTagsRequest(*elasticloadbalancing.AddTagsInput) elasticloadbalancing.AddTagsRequest

	ApplySecurityGroupsToLoadBalancerRequest(*elasticloadbalancing.ApplySecurityGroupsToLoadBalancerInput) elasticloadbalancing.ApplySecurityGroupsToLoadBalancerRequest

	AttachLoadBalancerToSubnetsRequest(*elasticloadbalancing.AttachLoadBalancerToSubnetsInput) elasticloadbalancing.AttachLoadBalancerToSubnetsRequest

	ConfigureHealthCheckRequest(*elasticloadbalancing.ConfigureHealthCheckInput) elasticloadbalancing.ConfigureHealthCheckRequest

	CreateAppCookieStickinessPolicyRequest(*elasticloadbalancing.CreateAppCookieStickinessPolicyInput) elasticloadbalancing.CreateAppCookieStickinessPolicyRequest

	CreateLBCookieStickinessPolicyRequest(*elasticloadbalancing.CreateLBCookieStickinessPolicyInput) elasticloadbalancing.CreateLBCookieStickinessPolicyRequest

	CreateLoadBalancerRequest(*elasticloadbalancing.CreateLoadBalancerInput) elasticloadbalancing.CreateLoadBalancerRequest

	CreateLoadBalancerListenersRequest(*elasticloadbalancing.CreateLoadBalancerListenersInput) elasticloadbalancing.CreateLoadBalancerListenersRequest

	CreateLoadBalancerPolicyRequest(*elasticloadbalancing.CreateLoadBalancerPolicyInput) elasticloadbalancing.CreateLoadBalancerPolicyRequest

	DeleteLoadBalancerRequest(*elasticloadbalancing.DeleteLoadBalancerInput) elasticloadbalancing.DeleteLoadBalancerRequest

	DeleteLoadBalancerListenersRequest(*elasticloadbalancing.DeleteLoadBalancerListenersInput) elasticloadbalancing.DeleteLoadBalancerListenersRequest

	DeleteLoadBalancerPolicyRequest(*elasticloadbalancing.DeleteLoadBalancerPolicyInput) elasticloadbalancing.DeleteLoadBalancerPolicyRequest

	DeregisterInstancesFromLoadBalancerRequest(*elasticloadbalancing.DeregisterInstancesFromLoadBalancerInput) elasticloadbalancing.DeregisterInstancesFromLoadBalancerRequest

	DescribeAccountLimitsRequest(*elasticloadbalancing.DescribeAccountLimitsInput) elasticloadbalancing.DescribeAccountLimitsRequest

	DescribeInstanceHealthRequest(*elasticloadbalancing.DescribeInstanceHealthInput) elasticloadbalancing.DescribeInstanceHealthRequest

	DescribeLoadBalancerAttributesRequest(*elasticloadbalancing.DescribeLoadBalancerAttributesInput) elasticloadbalancing.DescribeLoadBalancerAttributesRequest

	DescribeLoadBalancerPoliciesRequest(*elasticloadbalancing.DescribeLoadBalancerPoliciesInput) elasticloadbalancing.DescribeLoadBalancerPoliciesRequest

	DescribeLoadBalancerPolicyTypesRequest(*elasticloadbalancing.DescribeLoadBalancerPolicyTypesInput) elasticloadbalancing.DescribeLoadBalancerPolicyTypesRequest

	DescribeLoadBalancersRequest(*elasticloadbalancing.DescribeLoadBalancersInput) elasticloadbalancing.DescribeLoadBalancersRequest

	DescribeTagsRequest(*elasticloadbalancing.DescribeTagsInput) elasticloadbalancing.DescribeTagsRequest

	DetachLoadBalancerFromSubnetsRequest(*elasticloadbalancing.DetachLoadBalancerFromSubnetsInput) elasticloadbalancing.DetachLoadBalancerFromSubnetsRequest

	DisableAvailabilityZonesForLoadBalancerRequest(*elasticloadbalancing.DisableAvailabilityZonesForLoadBalancerInput) elasticloadbalancing.DisableAvailabilityZonesForLoadBalancerRequest

	EnableAvailabilityZonesForLoadBalancerRequest(*elasticloadbalancing.EnableAvailabilityZonesForLoadBalancerInput) elasticloadbalancing.EnableAvailabilityZonesForLoadBalancerRequest

	ModifyLoadBalancerAttributesRequest(*elasticloadbalancing.ModifyLoadBalancerAttributesInput) elasticloadbalancing.ModifyLoadBalancerAttributesRequest

	RegisterInstancesWithLoadBalancerRequest(*elasticloadbalancing.RegisterInstancesWithLoadBalancerInput) elasticloadbalancing.RegisterInstancesWithLoadBalancerRequest

	RemoveTagsRequest(*elasticloadbalancing.RemoveTagsInput) elasticloadbalancing.RemoveTagsRequest

	SetLoadBalancerListenerSSLCertificateRequest(*elasticloadbalancing.SetLoadBalancerListenerSSLCertificateInput) elasticloadbalancing.SetLoadBalancerListenerSSLCertificateRequest

	SetLoadBalancerPoliciesForBackendServerRequest(*elasticloadbalancing.SetLoadBalancerPoliciesForBackendServerInput) elasticloadbalancing.SetLoadBalancerPoliciesForBackendServerRequest

	SetLoadBalancerPoliciesOfListenerRequest(*elasticloadbalancing.SetLoadBalancerPoliciesOfListenerInput) elasticloadbalancing.SetLoadBalancerPoliciesOfListenerRequest

	WaitUntilAnyInstanceInService(context.Context, *elasticloadbalancing.DescribeInstanceHealthInput, ...aws.WaiterOption) error

	WaitUntilInstanceDeregistered(context.Context, *elasticloadbalancing.DescribeInstanceHealthInput, ...aws.WaiterOption) error

	WaitUntilInstanceInService(context.Context, *elasticloadbalancing.DescribeInstanceHealthInput, ...aws.WaiterOption) error
}

var _ ClientAPI = (*elasticloadbalancing.Client)(nil)
