// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type AdvertiseByoipCidrInput struct {
	_ struct{} `type:"structure"`

	// The address range, in CIDR notation. This must be the exact range that you
	// provisioned. You can't advertise only a portion of the provisioned range.
	//
	// Cidr is a required field
	Cidr *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`
}

// String returns the string representation
func (s AdvertiseByoipCidrInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AdvertiseByoipCidrInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AdvertiseByoipCidrInput"}

	if s.Cidr == nil {
		invalidParams.Add(aws.NewErrParamRequired("Cidr"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AdvertiseByoipCidrOutput struct {
	_ struct{} `type:"structure"`

	// Information about the address range.
	ByoipCidr *ByoipCidr `locationName:"byoipCidr" type:"structure"`
}

// String returns the string representation
func (s AdvertiseByoipCidrOutput) String() string {
	return awsutil.Prettify(s)
}

const opAdvertiseByoipCidr = "AdvertiseByoipCidr"

// AdvertiseByoipCidrRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Advertises an IPv4 or IPv6 address range that is provisioned for use with
// your AWS resources through bring your own IP addresses (BYOIP).
//
// You can perform this operation at most once every 10 seconds, even if you
// specify different address ranges each time.
//
// We recommend that you stop advertising the BYOIP CIDR from other locations
// when you advertise it from AWS. To minimize down time, you can configure
// your AWS resources to use an address from a BYOIP CIDR before it is advertised,
// and then simultaneously stop advertising it from the current location and
// start advertising it through AWS.
//
// It can take a few minutes before traffic to the specified addresses starts
// routing to AWS because of BGP propagation delays.
//
// To stop advertising the BYOIP CIDR, use WithdrawByoipCidr.
//
//    // Example sending a request using AdvertiseByoipCidrRequest.
//    req := client.AdvertiseByoipCidrRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AdvertiseByoipCidr
func (c *Client) AdvertiseByoipCidrRequest(input *AdvertiseByoipCidrInput) AdvertiseByoipCidrRequest {
	op := &aws.Operation{
		Name:       opAdvertiseByoipCidr,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AdvertiseByoipCidrInput{}
	}

	req := c.newRequest(op, input, &AdvertiseByoipCidrOutput{})

	return AdvertiseByoipCidrRequest{Request: req, Input: input, Copy: c.AdvertiseByoipCidrRequest}
}

// AdvertiseByoipCidrRequest is the request type for the
// AdvertiseByoipCidr API operation.
type AdvertiseByoipCidrRequest struct {
	*aws.Request
	Input *AdvertiseByoipCidrInput
	Copy  func(*AdvertiseByoipCidrInput) AdvertiseByoipCidrRequest
}

// Send marshals and sends the AdvertiseByoipCidr API request.
func (r AdvertiseByoipCidrRequest) Send(ctx context.Context) (*AdvertiseByoipCidrResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AdvertiseByoipCidrResponse{
		AdvertiseByoipCidrOutput: r.Request.Data.(*AdvertiseByoipCidrOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AdvertiseByoipCidrResponse is the response type for the
// AdvertiseByoipCidr API operation.
type AdvertiseByoipCidrResponse struct {
	*AdvertiseByoipCidrOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AdvertiseByoipCidr request.
func (r *AdvertiseByoipCidrResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
