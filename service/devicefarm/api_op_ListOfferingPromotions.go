// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListOfferingPromotionsInput struct {
	_ struct{} `type:"structure"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`
}

// String returns the string representation
func (s ListOfferingPromotionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListOfferingPromotionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListOfferingPromotionsInput"}
	if s.NextToken != nil && len(*s.NextToken) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListOfferingPromotionsOutput struct {
	_ struct{} `type:"structure"`

	// An identifier to be used in the next call to this operation, to return the
	// next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`

	// Information about the offering promotions.
	OfferingPromotions []OfferingPromotion `locationName:"offeringPromotions" type:"list"`
}

// String returns the string representation
func (s ListOfferingPromotionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListOfferingPromotions = "ListOfferingPromotions"

// ListOfferingPromotionsRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Returns a list of offering promotions. Each offering promotion record contains
// the ID and description of the promotion. The API returns a NotEligible error
// if the caller is not permitted to invoke the operation. Contact aws-devicefarm-support@amazon.com
// (mailto:aws-devicefarm-support@amazon.com) if you must be able to invoke
// this operation.
//
//    // Example sending a request using ListOfferingPromotionsRequest.
//    req := client.ListOfferingPromotionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListOfferingPromotions
func (c *Client) ListOfferingPromotionsRequest(input *ListOfferingPromotionsInput) ListOfferingPromotionsRequest {
	op := &aws.Operation{
		Name:       opListOfferingPromotions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListOfferingPromotionsInput{}
	}

	req := c.newRequest(op, input, &ListOfferingPromotionsOutput{})

	return ListOfferingPromotionsRequest{Request: req, Input: input, Copy: c.ListOfferingPromotionsRequest}
}

// ListOfferingPromotionsRequest is the request type for the
// ListOfferingPromotions API operation.
type ListOfferingPromotionsRequest struct {
	*aws.Request
	Input *ListOfferingPromotionsInput
	Copy  func(*ListOfferingPromotionsInput) ListOfferingPromotionsRequest
}

// Send marshals and sends the ListOfferingPromotions API request.
func (r ListOfferingPromotionsRequest) Send(ctx context.Context) (*ListOfferingPromotionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListOfferingPromotionsResponse{
		ListOfferingPromotionsOutput: r.Request.Data.(*ListOfferingPromotionsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListOfferingPromotionsResponse is the response type for the
// ListOfferingPromotions API operation.
type ListOfferingPromotionsResponse struct {
	*ListOfferingPromotionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListOfferingPromotions request.
func (r *ListOfferingPromotionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
