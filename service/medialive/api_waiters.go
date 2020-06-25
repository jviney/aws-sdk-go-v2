// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
)

// WaitUntilChannelCreated uses the MediaLive API operation
// DescribeChannel to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilChannelCreated(ctx context.Context, input *DescribeChannelInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilChannelCreated",
		MaxAttempts: 5,
		Delay:       aws.ConstantWaiterDelay(3 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "IDLE",
			},
			{
				State:   aws.RetryWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "CREATING",
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.StatusWaiterMatch,
				Expected: 500,
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "CREATE_FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeChannelInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeChannelRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilChannelDeleted uses the MediaLive API operation
// DescribeChannel to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilChannelDeleted(ctx context.Context, input *DescribeChannelInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilChannelDeleted",
		MaxAttempts: 84,
		Delay:       aws.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "DELETED",
			},
			{
				State:   aws.RetryWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "DELETING",
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.StatusWaiterMatch,
				Expected: 500,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeChannelInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeChannelRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilChannelRunning uses the MediaLive API operation
// DescribeChannel to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilChannelRunning(ctx context.Context, input *DescribeChannelInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilChannelRunning",
		MaxAttempts: 120,
		Delay:       aws.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "RUNNING",
			},
			{
				State:   aws.RetryWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "STARTING",
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.StatusWaiterMatch,
				Expected: 500,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeChannelInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeChannelRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilChannelStopped uses the MediaLive API operation
// DescribeChannel to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilChannelStopped(ctx context.Context, input *DescribeChannelInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilChannelStopped",
		MaxAttempts: 60,
		Delay:       aws.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "IDLE",
			},
			{
				State:   aws.RetryWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "STOPPING",
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.StatusWaiterMatch,
				Expected: 500,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeChannelInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeChannelRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilInputAttached uses the MediaLive API operation
// DescribeInput to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilInputAttached(ctx context.Context, input *DescribeInputInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilInputAttached",
		MaxAttempts: 20,
		Delay:       aws.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "ATTACHED",
			},
			{
				State:   aws.RetryWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "DETACHED",
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.StatusWaiterMatch,
				Expected: 500,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeInputInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeInputRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilInputDeleted uses the MediaLive API operation
// DescribeInput to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilInputDeleted(ctx context.Context, input *DescribeInputInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilInputDeleted",
		MaxAttempts: 20,
		Delay:       aws.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "DELETED",
			},
			{
				State:   aws.RetryWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "DELETING",
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.StatusWaiterMatch,
				Expected: 500,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeInputInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeInputRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilInputDetached uses the MediaLive API operation
// DescribeInput to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilInputDetached(ctx context.Context, input *DescribeInputInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilInputDetached",
		MaxAttempts: 84,
		Delay:       aws.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "DETACHED",
			},
			{
				State:   aws.RetryWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "CREATING",
			},
			{
				State:   aws.RetryWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "ATTACHED",
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.StatusWaiterMatch,
				Expected: 500,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeInputInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeInputRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilMultiplexCreated uses the MediaLive API operation
// DescribeMultiplex to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilMultiplexCreated(ctx context.Context, input *DescribeMultiplexInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilMultiplexCreated",
		MaxAttempts: 5,
		Delay:       aws.ConstantWaiterDelay(3 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "IDLE",
			},
			{
				State:   aws.RetryWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "CREATING",
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.StatusWaiterMatch,
				Expected: 500,
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "CREATE_FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeMultiplexInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeMultiplexRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilMultiplexDeleted uses the MediaLive API operation
// DescribeMultiplex to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilMultiplexDeleted(ctx context.Context, input *DescribeMultiplexInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilMultiplexDeleted",
		MaxAttempts: 20,
		Delay:       aws.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "DELETED",
			},
			{
				State:   aws.RetryWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "DELETING",
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.StatusWaiterMatch,
				Expected: 500,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeMultiplexInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeMultiplexRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilMultiplexRunning uses the MediaLive API operation
// DescribeMultiplex to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilMultiplexRunning(ctx context.Context, input *DescribeMultiplexInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilMultiplexRunning",
		MaxAttempts: 120,
		Delay:       aws.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "RUNNING",
			},
			{
				State:   aws.RetryWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "STARTING",
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.StatusWaiterMatch,
				Expected: 500,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeMultiplexInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeMultiplexRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilMultiplexStopped uses the MediaLive API operation
// DescribeMultiplex to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilMultiplexStopped(ctx context.Context, input *DescribeMultiplexInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilMultiplexStopped",
		MaxAttempts: 28,
		Delay:       aws.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "IDLE",
			},
			{
				State:   aws.RetryWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "State",
				Expected: "STOPPING",
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.StatusWaiterMatch,
				Expected: 500,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeMultiplexInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeMultiplexRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}
