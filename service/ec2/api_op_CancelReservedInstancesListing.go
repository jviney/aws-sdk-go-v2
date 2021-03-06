// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for CancelReservedInstancesListing.
type CancelReservedInstancesListingInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Reserved Instance listing.
	//
	// ReservedInstancesListingId is a required field
	ReservedInstancesListingId *string `locationName:"reservedInstancesListingId" type:"string" required:"true"`
}

// String returns the string representation
func (s CancelReservedInstancesListingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelReservedInstancesListingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelReservedInstancesListingInput"}

	if s.ReservedInstancesListingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReservedInstancesListingId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of CancelReservedInstancesListing.
type CancelReservedInstancesListingOutput struct {
	_ struct{} `type:"structure"`

	// The Reserved Instance listing.
	ReservedInstancesListings []ReservedInstancesListing `locationName:"reservedInstancesListingsSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s CancelReservedInstancesListingOutput) String() string {
	return awsutil.Prettify(s)
}

const opCancelReservedInstancesListing = "CancelReservedInstancesListing"

// CancelReservedInstancesListingRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Cancels the specified Reserved Instance listing in the Reserved Instance
// Marketplace.
//
// For more information, see Reserved Instance Marketplace (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-market-general.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using CancelReservedInstancesListingRequest.
//    req := client.CancelReservedInstancesListingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CancelReservedInstancesListing
func (c *Client) CancelReservedInstancesListingRequest(input *CancelReservedInstancesListingInput) CancelReservedInstancesListingRequest {
	op := &aws.Operation{
		Name:       opCancelReservedInstancesListing,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CancelReservedInstancesListingInput{}
	}

	req := c.newRequest(op, input, &CancelReservedInstancesListingOutput{})

	return CancelReservedInstancesListingRequest{Request: req, Input: input, Copy: c.CancelReservedInstancesListingRequest}
}

// CancelReservedInstancesListingRequest is the request type for the
// CancelReservedInstancesListing API operation.
type CancelReservedInstancesListingRequest struct {
	*aws.Request
	Input *CancelReservedInstancesListingInput
	Copy  func(*CancelReservedInstancesListingInput) CancelReservedInstancesListingRequest
}

// Send marshals and sends the CancelReservedInstancesListing API request.
func (r CancelReservedInstancesListingRequest) Send(ctx context.Context) (*CancelReservedInstancesListingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelReservedInstancesListingResponse{
		CancelReservedInstancesListingOutput: r.Request.Data.(*CancelReservedInstancesListingOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelReservedInstancesListingResponse is the response type for the
// CancelReservedInstancesListing API operation.
type CancelReservedInstancesListingResponse struct {
	*CancelReservedInstancesListingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelReservedInstancesListing request.
func (r *CancelReservedInstancesListingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
