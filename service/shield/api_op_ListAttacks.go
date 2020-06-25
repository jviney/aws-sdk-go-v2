// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListAttacksInput struct {
	_ struct{} `type:"structure"`

	// The end of the time period for the attacks. This is a timestamp type. The
	// sample request above indicates a number type because the default used by
	// WAF is Unix time in seconds. However any valid timestamp format (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types)
	// is allowed.
	EndTime *TimeRange `type:"structure"`

	// The maximum number of AttackSummary objects to be returned. If this is left
	// blank, the first 20 results will be returned.
	//
	// This is a maximum value; it is possible that AWS WAF will return the results
	// in smaller batches. That is, the number of AttackSummary objects returned
	// could be less than MaxResults, even if there are still more AttackSummary
	// objects yet to return. If there are more AttackSummary objects to return,
	// AWS WAF will always also return a NextToken.
	MaxResults *int64 `type:"integer"`

	// The ListAttacksRequest.NextMarker value from a previous call to ListAttacksRequest.
	// Pass null if this is the first call.
	NextToken *string `min:"1" type:"string"`

	// The ARN (Amazon Resource Name) of the resource that was attacked. If this
	// is left blank, all applicable resources for this account will be included.
	ResourceArns []string `type:"list"`

	// The start of the time period for the attacks. This is a timestamp type. The
	// sample request above indicates a number type because the default used by
	// WAF is Unix time in seconds. However any valid timestamp format (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types)
	// is allowed.
	StartTime *TimeRange `type:"structure"`
}

// String returns the string representation
func (s ListAttacksInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAttacksInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAttacksInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListAttacksOutput struct {
	_ struct{} `type:"structure"`

	// The attack information for the specified time range.
	AttackSummaries []AttackSummary `type:"list"`

	// The token returned by a previous call to indicate that there is more data
	// available. If not null, more results are available. Pass this value for the
	// NextMarker parameter in a subsequent call to ListAttacks to retrieve the
	// next set of items.
	//
	// AWS WAF might return the list of AttackSummary objects in batches smaller
	// than the number specified by MaxResults. If there are more AttackSummary
	// objects to return, AWS WAF will always also return a NextToken.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListAttacksOutput) String() string {
	return awsutil.Prettify(s)
}

const opListAttacks = "ListAttacks"

// ListAttacksRequest returns a request value for making API operation for
// AWS Shield.
//
// Returns all ongoing DDoS attacks or all DDoS attacks during a specified time
// period.
//
//    // Example sending a request using ListAttacksRequest.
//    req := client.ListAttacksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/ListAttacks
func (c *Client) ListAttacksRequest(input *ListAttacksInput) ListAttacksRequest {
	op := &aws.Operation{
		Name:       opListAttacks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAttacksInput{}
	}

	req := c.newRequest(op, input, &ListAttacksOutput{})

	return ListAttacksRequest{Request: req, Input: input, Copy: c.ListAttacksRequest}
}

// ListAttacksRequest is the request type for the
// ListAttacks API operation.
type ListAttacksRequest struct {
	*aws.Request
	Input *ListAttacksInput
	Copy  func(*ListAttacksInput) ListAttacksRequest
}

// Send marshals and sends the ListAttacks API request.
func (r ListAttacksRequest) Send(ctx context.Context) (*ListAttacksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAttacksResponse{
		ListAttacksOutput: r.Request.Data.(*ListAttacksOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListAttacksResponse is the response type for the
// ListAttacks API operation.
type ListAttacksResponse struct {
	*ListAttacksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAttacks request.
func (r *ListAttacksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
