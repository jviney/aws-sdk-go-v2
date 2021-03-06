// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for ModifyLoadBalancerAttributes.
type ModifyLoadBalancerAttributesInput struct {
	_ struct{} `type:"structure"`

	// The attributes for the load balancer.
	//
	// LoadBalancerAttributes is a required field
	LoadBalancerAttributes *LoadBalancerAttributes `type:"structure" required:"true"`

	// The name of the load balancer.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyLoadBalancerAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyLoadBalancerAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyLoadBalancerAttributesInput"}

	if s.LoadBalancerAttributes == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerAttributes"))
	}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}
	if s.LoadBalancerAttributes != nil {
		if err := s.LoadBalancerAttributes.Validate(); err != nil {
			invalidParams.AddNested("LoadBalancerAttributes", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of ModifyLoadBalancerAttributes.
type ModifyLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the load balancer attributes.
	LoadBalancerAttributes *LoadBalancerAttributes `type:"structure"`

	// The name of the load balancer.
	LoadBalancerName *string `type:"string"`
}

// String returns the string representation
func (s ModifyLoadBalancerAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyLoadBalancerAttributes = "ModifyLoadBalancerAttributes"

// ModifyLoadBalancerAttributesRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Modifies the attributes of the specified load balancer.
//
// You can modify the load balancer attributes, such as AccessLogs, ConnectionDraining,
// and CrossZoneLoadBalancing by either enabling or disabling them. Or, you
// can modify the load balancer attribute ConnectionSettings by specifying an
// idle connection timeout value for your load balancer.
//
// For more information, see the following in the Classic Load Balancers Guide:
//
//    * Cross-Zone Load Balancing (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-disable-crosszone-lb.html)
//
//    * Connection Draining (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-conn-drain.html)
//
//    * Access Logs (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/access-log-collection.html)
//
//    * Idle Connection Timeout (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-idle-timeout.html)
//
//    // Example sending a request using ModifyLoadBalancerAttributesRequest.
//    req := client.ModifyLoadBalancerAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/ModifyLoadBalancerAttributes
func (c *Client) ModifyLoadBalancerAttributesRequest(input *ModifyLoadBalancerAttributesInput) ModifyLoadBalancerAttributesRequest {
	op := &aws.Operation{
		Name:       opModifyLoadBalancerAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyLoadBalancerAttributesInput{}
	}

	req := c.newRequest(op, input, &ModifyLoadBalancerAttributesOutput{})

	return ModifyLoadBalancerAttributesRequest{Request: req, Input: input, Copy: c.ModifyLoadBalancerAttributesRequest}
}

// ModifyLoadBalancerAttributesRequest is the request type for the
// ModifyLoadBalancerAttributes API operation.
type ModifyLoadBalancerAttributesRequest struct {
	*aws.Request
	Input *ModifyLoadBalancerAttributesInput
	Copy  func(*ModifyLoadBalancerAttributesInput) ModifyLoadBalancerAttributesRequest
}

// Send marshals and sends the ModifyLoadBalancerAttributes API request.
func (r ModifyLoadBalancerAttributesRequest) Send(ctx context.Context) (*ModifyLoadBalancerAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyLoadBalancerAttributesResponse{
		ModifyLoadBalancerAttributesOutput: r.Request.Data.(*ModifyLoadBalancerAttributesOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyLoadBalancerAttributesResponse is the response type for the
// ModifyLoadBalancerAttributes API operation.
type ModifyLoadBalancerAttributesResponse struct {
	*ModifyLoadBalancerAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyLoadBalancerAttributes request.
func (r *ModifyLoadBalancerAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
