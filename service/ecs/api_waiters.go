// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
)

// WaitUntilServicesInactive uses the Amazon ECS API operation
// DescribeServices to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilServicesInactive(ctx context.Context, input *DescribeServicesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilServicesInactive",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "failures[].reason",
				Expected: "MISSING",
			},
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "services[].status",
				Expected: "INACTIVE",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeServicesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeServicesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilServicesStable uses the Amazon ECS API operation
// DescribeServices to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilServicesStable(ctx context.Context, input *DescribeServicesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilServicesStable",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "failures[].reason",
				Expected: "MISSING",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "services[].status",
				Expected: "DRAINING",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "services[].status",
				Expected: "INACTIVE",
			},
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "length(services[?!(length(deployments) == `1` && runningCount == desiredCount)]) == `0`",
				Expected: true,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeServicesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeServicesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilTasksRunning uses the Amazon ECS API operation
// DescribeTasks to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilTasksRunning(ctx context.Context, input *DescribeTasksInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilTasksRunning",
		MaxAttempts: 100,
		Delay:       aws.ConstantWaiterDelay(6 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "tasks[].lastStatus",
				Expected: "STOPPED",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "failures[].reason",
				Expected: "MISSING",
			},
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "tasks[].lastStatus",
				Expected: "RUNNING",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeTasksInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeTasksRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilTasksStopped uses the Amazon ECS API operation
// DescribeTasks to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilTasksStopped(ctx context.Context, input *DescribeTasksInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilTasksStopped",
		MaxAttempts: 100,
		Delay:       aws.ConstantWaiterDelay(6 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "tasks[].lastStatus",
				Expected: "STOPPED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeTasksInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeTasksRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}
