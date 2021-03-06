// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type PutAggregationAuthorizationInput struct {
	_ struct{} `type:"structure"`

	// The 12-digit account ID of the account authorized to aggregate data.
	//
	// AuthorizedAccountId is a required field
	AuthorizedAccountId *string `type:"string" required:"true"`

	// The region authorized to collect aggregated data.
	//
	// AuthorizedAwsRegion is a required field
	AuthorizedAwsRegion *string `min:"1" type:"string" required:"true"`

	// An array of tag object.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s PutAggregationAuthorizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutAggregationAuthorizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutAggregationAuthorizationInput"}

	if s.AuthorizedAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AuthorizedAccountId"))
	}

	if s.AuthorizedAwsRegion == nil {
		invalidParams.Add(aws.NewErrParamRequired("AuthorizedAwsRegion"))
	}
	if s.AuthorizedAwsRegion != nil && len(*s.AuthorizedAwsRegion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthorizedAwsRegion", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutAggregationAuthorizationOutput struct {
	_ struct{} `type:"structure"`

	// Returns an AggregationAuthorization object.
	AggregationAuthorization *AggregationAuthorization `type:"structure"`
}

// String returns the string representation
func (s PutAggregationAuthorizationOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutAggregationAuthorization = "PutAggregationAuthorization"

// PutAggregationAuthorizationRequest returns a request value for making API operation for
// AWS Config.
//
// Authorizes the aggregator account and region to collect data from the source
// account and region.
//
//    // Example sending a request using PutAggregationAuthorizationRequest.
//    req := client.PutAggregationAuthorizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/PutAggregationAuthorization
func (c *Client) PutAggregationAuthorizationRequest(input *PutAggregationAuthorizationInput) PutAggregationAuthorizationRequest {
	op := &aws.Operation{
		Name:       opPutAggregationAuthorization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutAggregationAuthorizationInput{}
	}

	req := c.newRequest(op, input, &PutAggregationAuthorizationOutput{})

	return PutAggregationAuthorizationRequest{Request: req, Input: input, Copy: c.PutAggregationAuthorizationRequest}
}

// PutAggregationAuthorizationRequest is the request type for the
// PutAggregationAuthorization API operation.
type PutAggregationAuthorizationRequest struct {
	*aws.Request
	Input *PutAggregationAuthorizationInput
	Copy  func(*PutAggregationAuthorizationInput) PutAggregationAuthorizationRequest
}

// Send marshals and sends the PutAggregationAuthorization API request.
func (r PutAggregationAuthorizationRequest) Send(ctx context.Context) (*PutAggregationAuthorizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutAggregationAuthorizationResponse{
		PutAggregationAuthorizationOutput: r.Request.Data.(*PutAggregationAuthorizationOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutAggregationAuthorizationResponse is the response type for the
// PutAggregationAuthorization API operation.
type PutAggregationAuthorizationResponse struct {
	*PutAggregationAuthorizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutAggregationAuthorization request.
func (r *PutAggregationAuthorizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
