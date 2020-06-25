// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
)

// WaitUntilVaultExists uses the Amazon Glacier API operation
// DescribeVault to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilVaultExists(ctx context.Context, input *DescribeVaultInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilVaultExists",
		MaxAttempts: 15,
		Delay:       aws.ConstantWaiterDelay(3 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.StatusWaiterMatch,
				Expected: 200,
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "ResourceNotFoundException",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeVaultInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeVaultRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilVaultNotExists uses the Amazon Glacier API operation
// DescribeVault to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilVaultNotExists(ctx context.Context, input *DescribeVaultInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilVaultNotExists",
		MaxAttempts: 15,
		Delay:       aws.ConstantWaiterDelay(3 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.StatusWaiterMatch,
				Expected: 200,
			},
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "ResourceNotFoundException",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeVaultInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeVaultRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}
