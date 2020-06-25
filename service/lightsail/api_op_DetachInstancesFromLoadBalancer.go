// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DetachInstancesFromLoadBalancerInput struct {
	_ struct{} `type:"structure"`

	// An array of strings containing the names of the instances you want to detach
	// from the load balancer.
	//
	// InstanceNames is a required field
	InstanceNames []string `locationName:"instanceNames" type:"list" required:"true"`

	// The name of the Lightsail load balancer.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `locationName:"loadBalancerName" type:"string" required:"true"`
}

// String returns the string representation
func (s DetachInstancesFromLoadBalancerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachInstancesFromLoadBalancerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetachInstancesFromLoadBalancerInput"}

	if s.InstanceNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceNames"))
	}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DetachInstancesFromLoadBalancerOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected
	// by the request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s DetachInstancesFromLoadBalancerOutput) String() string {
	return awsutil.Prettify(s)
}

const opDetachInstancesFromLoadBalancer = "DetachInstancesFromLoadBalancer"

// DetachInstancesFromLoadBalancerRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Detaches the specified instances from a Lightsail load balancer.
//
// This operation waits until the instances are no longer needed before they
// are detached from the load balancer.
//
// The detach instances from load balancer operation supports tag-based access
// control via resource tags applied to the resource identified by load balancer
// name. For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using DetachInstancesFromLoadBalancerRequest.
//    req := client.DetachInstancesFromLoadBalancerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DetachInstancesFromLoadBalancer
func (c *Client) DetachInstancesFromLoadBalancerRequest(input *DetachInstancesFromLoadBalancerInput) DetachInstancesFromLoadBalancerRequest {
	op := &aws.Operation{
		Name:       opDetachInstancesFromLoadBalancer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachInstancesFromLoadBalancerInput{}
	}

	req := c.newRequest(op, input, &DetachInstancesFromLoadBalancerOutput{})

	return DetachInstancesFromLoadBalancerRequest{Request: req, Input: input, Copy: c.DetachInstancesFromLoadBalancerRequest}
}

// DetachInstancesFromLoadBalancerRequest is the request type for the
// DetachInstancesFromLoadBalancer API operation.
type DetachInstancesFromLoadBalancerRequest struct {
	*aws.Request
	Input *DetachInstancesFromLoadBalancerInput
	Copy  func(*DetachInstancesFromLoadBalancerInput) DetachInstancesFromLoadBalancerRequest
}

// Send marshals and sends the DetachInstancesFromLoadBalancer API request.
func (r DetachInstancesFromLoadBalancerRequest) Send(ctx context.Context) (*DetachInstancesFromLoadBalancerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachInstancesFromLoadBalancerResponse{
		DetachInstancesFromLoadBalancerOutput: r.Request.Data.(*DetachInstancesFromLoadBalancerOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachInstancesFromLoadBalancerResponse is the response type for the
// DetachInstancesFromLoadBalancer API operation.
type DetachInstancesFromLoadBalancerResponse struct {
	*DetachInstancesFromLoadBalancerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachInstancesFromLoadBalancer request.
func (r *DetachInstancesFromLoadBalancerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
