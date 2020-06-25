// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a PurchaseReservedCacheNodesOffering operation.
type PurchaseReservedCacheNodesOfferingInput struct {
	_ struct{} `type:"structure"`

	// The number of cache node instances to reserve.
	//
	// Default: 1
	CacheNodeCount *int64 `type:"integer"`

	// A customer-specified identifier to track this reservation.
	//
	// The Reserved Cache Node ID is an unique customer-specified identifier to
	// track this reservation. If this parameter is not specified, ElastiCache automatically
	// generates an identifier for the reservation.
	//
	// Example: myreservationID
	ReservedCacheNodeId *string `type:"string"`

	// The ID of the reserved cache node offering to purchase.
	//
	// Example: 438012d3-4052-4cc7-b2e3-8d3372e0e706
	//
	// ReservedCacheNodesOfferingId is a required field
	ReservedCacheNodesOfferingId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PurchaseReservedCacheNodesOfferingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PurchaseReservedCacheNodesOfferingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PurchaseReservedCacheNodesOfferingInput"}

	if s.ReservedCacheNodesOfferingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReservedCacheNodesOfferingId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PurchaseReservedCacheNodesOfferingOutput struct {
	_ struct{} `type:"structure"`

	// Represents the output of a PurchaseReservedCacheNodesOffering operation.
	ReservedCacheNode *ReservedCacheNode `type:"structure"`
}

// String returns the string representation
func (s PurchaseReservedCacheNodesOfferingOutput) String() string {
	return awsutil.Prettify(s)
}

const opPurchaseReservedCacheNodesOffering = "PurchaseReservedCacheNodesOffering"

// PurchaseReservedCacheNodesOfferingRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Allows you to purchase a reserved cache node offering.
//
//    // Example sending a request using PurchaseReservedCacheNodesOfferingRequest.
//    req := client.PurchaseReservedCacheNodesOfferingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/PurchaseReservedCacheNodesOffering
func (c *Client) PurchaseReservedCacheNodesOfferingRequest(input *PurchaseReservedCacheNodesOfferingInput) PurchaseReservedCacheNodesOfferingRequest {
	op := &aws.Operation{
		Name:       opPurchaseReservedCacheNodesOffering,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PurchaseReservedCacheNodesOfferingInput{}
	}

	req := c.newRequest(op, input, &PurchaseReservedCacheNodesOfferingOutput{})

	return PurchaseReservedCacheNodesOfferingRequest{Request: req, Input: input, Copy: c.PurchaseReservedCacheNodesOfferingRequest}
}

// PurchaseReservedCacheNodesOfferingRequest is the request type for the
// PurchaseReservedCacheNodesOffering API operation.
type PurchaseReservedCacheNodesOfferingRequest struct {
	*aws.Request
	Input *PurchaseReservedCacheNodesOfferingInput
	Copy  func(*PurchaseReservedCacheNodesOfferingInput) PurchaseReservedCacheNodesOfferingRequest
}

// Send marshals and sends the PurchaseReservedCacheNodesOffering API request.
func (r PurchaseReservedCacheNodesOfferingRequest) Send(ctx context.Context) (*PurchaseReservedCacheNodesOfferingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PurchaseReservedCacheNodesOfferingResponse{
		PurchaseReservedCacheNodesOfferingOutput: r.Request.Data.(*PurchaseReservedCacheNodesOfferingOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PurchaseReservedCacheNodesOfferingResponse is the response type for the
// PurchaseReservedCacheNodesOffering API operation.
type PurchaseReservedCacheNodesOfferingResponse struct {
	*PurchaseReservedCacheNodesOfferingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PurchaseReservedCacheNodesOffering request.
func (r *PurchaseReservedCacheNodesOfferingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
