// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeFleetCapacityInput
type DescribeFleetCapacityInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for a fleet(s) to retrieve capacity information for. To
	// request capacity information for all fleets, leave this parameter empty.
	FleetIds []string `min:"1" type:"list"`

	// Maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. This parameter is ignored when
	// the request specifies one or a list of fleet IDs.
	Limit *int64 `min:"1" type:"integer"`

	// Token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this action. To start
	// at the beginning of the result set, do not specify a value. This parameter
	// is ignored when the request specifies one or a list of fleet IDs.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeFleetCapacityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeFleetCapacityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeFleetCapacityInput"}
	if s.FleetIds != nil && len(s.FleetIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetIds", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeFleetCapacityOutput
type DescribeFleetCapacityOutput struct {
	_ struct{} `type:"structure"`

	// Collection of objects containing capacity information for each requested
	// fleet ID. Leave this parameter empty to retrieve capacity information for
	// all fleets.
	FleetCapacity []FleetCapacity `type:"list"`

	// Token that indicates where to resume retrieving results on the next call
	// to this action. If no token is returned, these results represent the end
	// of the list.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeFleetCapacityOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeFleetCapacity = "DescribeFleetCapacity"

// DescribeFleetCapacityRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves the current status of fleet capacity for one or more fleets. This
// information includes the number of instances that have been requested for
// the fleet and the number currently active. You can request capacity for all
// fleets, or specify a list of one or more fleet IDs. When requesting multiple
// fleets, use the pagination parameters to retrieve results as a set of sequential
// pages. If successful, a FleetCapacity object is returned for each requested
// fleet ID. When specifying a list of fleet IDs, attribute objects are returned
// only for fleets that currently exist.
//
// Some API actions may limit the number of fleet IDs allowed in one request.
// If a request exceeds this limit, the request fails and the error message
// includes the maximum allowed.
//
// Learn more
//
// Working with Fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html).
//
// Related operations
//
//    * CreateFleet
//
//    * ListFleets
//
//    * DeleteFleet
//
//    * Describe fleets:
//
// DescribeFleetAttributes
//
// DescribeFleetCapacity
//
// DescribeFleetPortSettings
//
// DescribeFleetUtilization
//
// DescribeRuntimeConfiguration
//
// DescribeEC2InstanceLimits
//
// DescribeFleetEvents
//
//    * Update fleets:
//
// UpdateFleetAttributes
//
// UpdateFleetCapacity
//
// UpdateFleetPortSettings
//
// UpdateRuntimeConfiguration
//
//    * Manage fleet actions:
//
// StartFleetActions
//
// StopFleetActions
//
//    // Example sending a request using DescribeFleetCapacityRequest.
//    req := client.DescribeFleetCapacityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeFleetCapacity
func (c *Client) DescribeFleetCapacityRequest(input *DescribeFleetCapacityInput) DescribeFleetCapacityRequest {
	op := &aws.Operation{
		Name:       opDescribeFleetCapacity,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeFleetCapacityInput{}
	}

	req := c.newRequest(op, input, &DescribeFleetCapacityOutput{})
	return DescribeFleetCapacityRequest{Request: req, Input: input, Copy: c.DescribeFleetCapacityRequest}
}

// DescribeFleetCapacityRequest is the request type for the
// DescribeFleetCapacity API operation.
type DescribeFleetCapacityRequest struct {
	*aws.Request
	Input *DescribeFleetCapacityInput
	Copy  func(*DescribeFleetCapacityInput) DescribeFleetCapacityRequest
}

// Send marshals and sends the DescribeFleetCapacity API request.
func (r DescribeFleetCapacityRequest) Send(ctx context.Context) (*DescribeFleetCapacityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFleetCapacityResponse{
		DescribeFleetCapacityOutput: r.Request.Data.(*DescribeFleetCapacityOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeFleetCapacityResponse is the response type for the
// DescribeFleetCapacity API operation.
type DescribeFleetCapacityResponse struct {
	*DescribeFleetCapacityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFleetCapacity request.
func (r *DescribeFleetCapacityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}