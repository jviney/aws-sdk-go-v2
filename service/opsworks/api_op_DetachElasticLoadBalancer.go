// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DetachElasticLoadBalancerInput struct {
	_ struct{} `type:"structure"`

	// The Elastic Load Balancing instance's name.
	//
	// ElasticLoadBalancerName is a required field
	ElasticLoadBalancerName *string `type:"string" required:"true"`

	// The ID of the layer that the Elastic Load Balancing instance is attached
	// to.
	//
	// LayerId is a required field
	LayerId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DetachElasticLoadBalancerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachElasticLoadBalancerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetachElasticLoadBalancerInput"}

	if s.ElasticLoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ElasticLoadBalancerName"))
	}

	if s.LayerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("LayerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DetachElasticLoadBalancerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DetachElasticLoadBalancerOutput) String() string {
	return awsutil.Prettify(s)
}

const opDetachElasticLoadBalancer = "DetachElasticLoadBalancer"

// DetachElasticLoadBalancerRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Detaches a specified Elastic Load Balancing instance from its layer.
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using DetachElasticLoadBalancerRequest.
//    req := client.DetachElasticLoadBalancerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DetachElasticLoadBalancer
func (c *Client) DetachElasticLoadBalancerRequest(input *DetachElasticLoadBalancerInput) DetachElasticLoadBalancerRequest {
	op := &aws.Operation{
		Name:       opDetachElasticLoadBalancer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachElasticLoadBalancerInput{}
	}

	req := c.newRequest(op, input, &DetachElasticLoadBalancerOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DetachElasticLoadBalancerRequest{Request: req, Input: input, Copy: c.DetachElasticLoadBalancerRequest}
}

// DetachElasticLoadBalancerRequest is the request type for the
// DetachElasticLoadBalancer API operation.
type DetachElasticLoadBalancerRequest struct {
	*aws.Request
	Input *DetachElasticLoadBalancerInput
	Copy  func(*DetachElasticLoadBalancerInput) DetachElasticLoadBalancerRequest
}

// Send marshals and sends the DetachElasticLoadBalancer API request.
func (r DetachElasticLoadBalancerRequest) Send(ctx context.Context) (*DetachElasticLoadBalancerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachElasticLoadBalancerResponse{
		DetachElasticLoadBalancerOutput: r.Request.Data.(*DetachElasticLoadBalancerOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachElasticLoadBalancerResponse is the response type for the
// DetachElasticLoadBalancer API operation.
type DetachElasticLoadBalancerResponse struct {
	*DetachElasticLoadBalancerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachElasticLoadBalancer request.
func (r *DetachElasticLoadBalancerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
