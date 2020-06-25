// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DeleteLoadBalancerListeners.
type DeleteLoadBalancerListenersInput struct {
	_ struct{} `type:"structure"`

	// The name of the load balancer.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `type:"string" required:"true"`

	// The client port numbers of the listeners.
	//
	// LoadBalancerPorts is a required field
	LoadBalancerPorts []int64 `type:"list" required:"true"`
}

// String returns the string representation
func (s DeleteLoadBalancerListenersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLoadBalancerListenersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLoadBalancerListenersInput"}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}

	if s.LoadBalancerPorts == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerPorts"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of DeleteLoadBalancerListeners.
type DeleteLoadBalancerListenersOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteLoadBalancerListenersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteLoadBalancerListeners = "DeleteLoadBalancerListeners"

// DeleteLoadBalancerListenersRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Deletes the specified listeners from the specified load balancer.
//
//    // Example sending a request using DeleteLoadBalancerListenersRequest.
//    req := client.DeleteLoadBalancerListenersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/DeleteLoadBalancerListeners
func (c *Client) DeleteLoadBalancerListenersRequest(input *DeleteLoadBalancerListenersInput) DeleteLoadBalancerListenersRequest {
	op := &aws.Operation{
		Name:       opDeleteLoadBalancerListeners,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteLoadBalancerListenersInput{}
	}

	req := c.newRequest(op, input, &DeleteLoadBalancerListenersOutput{})

	return DeleteLoadBalancerListenersRequest{Request: req, Input: input, Copy: c.DeleteLoadBalancerListenersRequest}
}

// DeleteLoadBalancerListenersRequest is the request type for the
// DeleteLoadBalancerListeners API operation.
type DeleteLoadBalancerListenersRequest struct {
	*aws.Request
	Input *DeleteLoadBalancerListenersInput
	Copy  func(*DeleteLoadBalancerListenersInput) DeleteLoadBalancerListenersRequest
}

// Send marshals and sends the DeleteLoadBalancerListeners API request.
func (r DeleteLoadBalancerListenersRequest) Send(ctx context.Context) (*DeleteLoadBalancerListenersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteLoadBalancerListenersResponse{
		DeleteLoadBalancerListenersOutput: r.Request.Data.(*DeleteLoadBalancerListenersOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteLoadBalancerListenersResponse is the response type for the
// DeleteLoadBalancerListeners API operation.
type DeleteLoadBalancerListenersResponse struct {
	*DeleteLoadBalancerListenersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteLoadBalancerListeners request.
func (r *DeleteLoadBalancerListenersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
