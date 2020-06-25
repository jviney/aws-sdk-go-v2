// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetRegionsInput struct {
	_ struct{} `type:"structure"`

	// A Boolean value indicating whether to also include Availability Zones in
	// your get regions request. Availability Zones are indicated with a letter:
	// e.g., us-east-2a.
	IncludeAvailabilityZones *bool `locationName:"includeAvailabilityZones" type:"boolean"`

	// >A Boolean value indicating whether to also include Availability Zones for
	// databases in your get regions request. Availability Zones are indicated with
	// a letter (e.g., us-east-2a).
	IncludeRelationalDatabaseAvailabilityZones *bool `locationName:"includeRelationalDatabaseAvailabilityZones" type:"boolean"`
}

// String returns the string representation
func (s GetRegionsInput) String() string {
	return awsutil.Prettify(s)
}

type GetRegionsOutput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs containing information about your get regions
	// request.
	Regions []Region `locationName:"regions" type:"list"`
}

// String returns the string representation
func (s GetRegionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRegions = "GetRegions"

// GetRegionsRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns a list of all valid regions for Amazon Lightsail. Use the include
// availability zones parameter to also return the Availability Zones in a region.
//
//    // Example sending a request using GetRegionsRequest.
//    req := client.GetRegionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetRegions
func (c *Client) GetRegionsRequest(input *GetRegionsInput) GetRegionsRequest {
	op := &aws.Operation{
		Name:       opGetRegions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRegionsInput{}
	}

	req := c.newRequest(op, input, &GetRegionsOutput{})

	return GetRegionsRequest{Request: req, Input: input, Copy: c.GetRegionsRequest}
}

// GetRegionsRequest is the request type for the
// GetRegions API operation.
type GetRegionsRequest struct {
	*aws.Request
	Input *GetRegionsInput
	Copy  func(*GetRegionsInput) GetRegionsRequest
}

// Send marshals and sends the GetRegions API request.
func (r GetRegionsRequest) Send(ctx context.Context) (*GetRegionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRegionsResponse{
		GetRegionsOutput: r.Request.Data.(*GetRegionsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRegionsResponse is the response type for the
// GetRegions API operation.
type GetRegionsResponse struct {
	*GetRegionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRegions request.
func (r *GetRegionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
