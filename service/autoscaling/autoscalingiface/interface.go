// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package autoscalingiface provides an interface to enable mocking the Auto Scaling service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package autoscalingiface

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/service/autoscaling"
)

// ClientAPI provides an interface to enable mocking the
// autoscaling.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Auto Scaling.
//    func myFunc(svc autoscalingiface.ClientAPI) bool {
//        // Make svc.AttachInstances request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := autoscaling.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        autoscalingiface.ClientPI
//    }
//    func (m *mockClientClient) AttachInstances(input *autoscaling.AttachInstancesInput) (*autoscaling.AttachInstancesOutput, error) {
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
	AttachInstancesRequest(*autoscaling.AttachInstancesInput) autoscaling.AttachInstancesRequest

	AttachLoadBalancerTargetGroupsRequest(*autoscaling.AttachLoadBalancerTargetGroupsInput) autoscaling.AttachLoadBalancerTargetGroupsRequest

	AttachLoadBalancersRequest(*autoscaling.AttachLoadBalancersInput) autoscaling.AttachLoadBalancersRequest

	BatchDeleteScheduledActionRequest(*autoscaling.BatchDeleteScheduledActionInput) autoscaling.BatchDeleteScheduledActionRequest

	BatchPutScheduledUpdateGroupActionRequest(*autoscaling.BatchPutScheduledUpdateGroupActionInput) autoscaling.BatchPutScheduledUpdateGroupActionRequest

	CompleteLifecycleActionRequest(*autoscaling.CompleteLifecycleActionInput) autoscaling.CompleteLifecycleActionRequest

	CreateAutoScalingGroupRequest(*autoscaling.CreateAutoScalingGroupInput) autoscaling.CreateAutoScalingGroupRequest

	CreateLaunchConfigurationRequest(*autoscaling.CreateLaunchConfigurationInput) autoscaling.CreateLaunchConfigurationRequest

	CreateOrUpdateTagsRequest(*autoscaling.CreateOrUpdateTagsInput) autoscaling.CreateOrUpdateTagsRequest

	DeleteAutoScalingGroupRequest(*autoscaling.DeleteAutoScalingGroupInput) autoscaling.DeleteAutoScalingGroupRequest

	DeleteLaunchConfigurationRequest(*autoscaling.DeleteLaunchConfigurationInput) autoscaling.DeleteLaunchConfigurationRequest

	DeleteLifecycleHookRequest(*autoscaling.DeleteLifecycleHookInput) autoscaling.DeleteLifecycleHookRequest

	DeleteNotificationConfigurationRequest(*autoscaling.DeleteNotificationConfigurationInput) autoscaling.DeleteNotificationConfigurationRequest

	DeletePolicyRequest(*autoscaling.DeletePolicyInput) autoscaling.DeletePolicyRequest

	DeleteScheduledActionRequest(*autoscaling.DeleteScheduledActionInput) autoscaling.DeleteScheduledActionRequest

	DeleteTagsRequest(*autoscaling.DeleteTagsInput) autoscaling.DeleteTagsRequest

	DescribeAccountLimitsRequest(*autoscaling.DescribeAccountLimitsInput) autoscaling.DescribeAccountLimitsRequest

	DescribeAdjustmentTypesRequest(*autoscaling.DescribeAdjustmentTypesInput) autoscaling.DescribeAdjustmentTypesRequest

	DescribeAutoScalingGroupsRequest(*autoscaling.DescribeAutoScalingGroupsInput) autoscaling.DescribeAutoScalingGroupsRequest

	DescribeAutoScalingInstancesRequest(*autoscaling.DescribeAutoScalingInstancesInput) autoscaling.DescribeAutoScalingInstancesRequest

	DescribeAutoScalingNotificationTypesRequest(*autoscaling.DescribeAutoScalingNotificationTypesInput) autoscaling.DescribeAutoScalingNotificationTypesRequest

	DescribeLaunchConfigurationsRequest(*autoscaling.DescribeLaunchConfigurationsInput) autoscaling.DescribeLaunchConfigurationsRequest

	DescribeLifecycleHookTypesRequest(*autoscaling.DescribeLifecycleHookTypesInput) autoscaling.DescribeLifecycleHookTypesRequest

	DescribeLifecycleHooksRequest(*autoscaling.DescribeLifecycleHooksInput) autoscaling.DescribeLifecycleHooksRequest

	DescribeLoadBalancerTargetGroupsRequest(*autoscaling.DescribeLoadBalancerTargetGroupsInput) autoscaling.DescribeLoadBalancerTargetGroupsRequest

	DescribeLoadBalancersRequest(*autoscaling.DescribeLoadBalancersInput) autoscaling.DescribeLoadBalancersRequest

	DescribeMetricCollectionTypesRequest(*autoscaling.DescribeMetricCollectionTypesInput) autoscaling.DescribeMetricCollectionTypesRequest

	DescribeNotificationConfigurationsRequest(*autoscaling.DescribeNotificationConfigurationsInput) autoscaling.DescribeNotificationConfigurationsRequest

	DescribePoliciesRequest(*autoscaling.DescribePoliciesInput) autoscaling.DescribePoliciesRequest

	DescribeScalingActivitiesRequest(*autoscaling.DescribeScalingActivitiesInput) autoscaling.DescribeScalingActivitiesRequest

	DescribeScalingProcessTypesRequest(*autoscaling.DescribeScalingProcessTypesInput) autoscaling.DescribeScalingProcessTypesRequest

	DescribeScheduledActionsRequest(*autoscaling.DescribeScheduledActionsInput) autoscaling.DescribeScheduledActionsRequest

	DescribeTagsRequest(*autoscaling.DescribeTagsInput) autoscaling.DescribeTagsRequest

	DescribeTerminationPolicyTypesRequest(*autoscaling.DescribeTerminationPolicyTypesInput) autoscaling.DescribeTerminationPolicyTypesRequest

	DetachInstancesRequest(*autoscaling.DetachInstancesInput) autoscaling.DetachInstancesRequest

	DetachLoadBalancerTargetGroupsRequest(*autoscaling.DetachLoadBalancerTargetGroupsInput) autoscaling.DetachLoadBalancerTargetGroupsRequest

	DetachLoadBalancersRequest(*autoscaling.DetachLoadBalancersInput) autoscaling.DetachLoadBalancersRequest

	DisableMetricsCollectionRequest(*autoscaling.DisableMetricsCollectionInput) autoscaling.DisableMetricsCollectionRequest

	EnableMetricsCollectionRequest(*autoscaling.EnableMetricsCollectionInput) autoscaling.EnableMetricsCollectionRequest

	EnterStandbyRequest(*autoscaling.EnterStandbyInput) autoscaling.EnterStandbyRequest

	ExecutePolicyRequest(*autoscaling.ExecutePolicyInput) autoscaling.ExecutePolicyRequest

	ExitStandbyRequest(*autoscaling.ExitStandbyInput) autoscaling.ExitStandbyRequest

	PutLifecycleHookRequest(*autoscaling.PutLifecycleHookInput) autoscaling.PutLifecycleHookRequest

	PutNotificationConfigurationRequest(*autoscaling.PutNotificationConfigurationInput) autoscaling.PutNotificationConfigurationRequest

	PutScalingPolicyRequest(*autoscaling.PutScalingPolicyInput) autoscaling.PutScalingPolicyRequest

	PutScheduledUpdateGroupActionRequest(*autoscaling.PutScheduledUpdateGroupActionInput) autoscaling.PutScheduledUpdateGroupActionRequest

	RecordLifecycleActionHeartbeatRequest(*autoscaling.RecordLifecycleActionHeartbeatInput) autoscaling.RecordLifecycleActionHeartbeatRequest

	ResumeProcessesRequest(*autoscaling.ResumeProcessesInput) autoscaling.ResumeProcessesRequest

	SetDesiredCapacityRequest(*autoscaling.SetDesiredCapacityInput) autoscaling.SetDesiredCapacityRequest

	SetInstanceHealthRequest(*autoscaling.SetInstanceHealthInput) autoscaling.SetInstanceHealthRequest

	SetInstanceProtectionRequest(*autoscaling.SetInstanceProtectionInput) autoscaling.SetInstanceProtectionRequest

	SuspendProcessesRequest(*autoscaling.SuspendProcessesInput) autoscaling.SuspendProcessesRequest

	TerminateInstanceInAutoScalingGroupRequest(*autoscaling.TerminateInstanceInAutoScalingGroupInput) autoscaling.TerminateInstanceInAutoScalingGroupRequest

	UpdateAutoScalingGroupRequest(*autoscaling.UpdateAutoScalingGroupInput) autoscaling.UpdateAutoScalingGroupRequest

	WaitUntilGroupExists(context.Context, *autoscaling.DescribeAutoScalingGroupsInput, ...aws.WaiterOption) error

	WaitUntilGroupInService(context.Context, *autoscaling.DescribeAutoScalingGroupsInput, ...aws.WaiterOption) error

	WaitUntilGroupNotExists(context.Context, *autoscaling.DescribeAutoScalingGroupsInput, ...aws.WaiterOption) error
}

var _ ClientAPI = (*autoscaling.Client)(nil)
