// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudhsm

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListHapgsInput struct {
	_ struct{} `type:"structure"`

	// The NextToken value from a previous call to ListHapgs. Pass null if this
	// is the first call.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListHapgsInput) String() string {
	return awsutil.Prettify(s)
}

type ListHapgsOutput struct {
	_ struct{} `type:"structure"`

	// The list of high-availability partition groups.
	//
	// HapgList is a required field
	HapgList []string `type:"list" required:"true"`

	// If not null, more results are available. Pass this value to ListHapgs to
	// retrieve the next set of items.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListHapgsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListHapgs = "ListHapgs"

// ListHapgsRequest returns a request value for making API operation for
// Amazon CloudHSM.
//
// This is documentation for AWS CloudHSM Classic. For more information, see
// AWS CloudHSM Classic FAQs (http://aws.amazon.com/cloudhsm/faqs-classic/),
// the AWS CloudHSM Classic User Guide (http://docs.aws.amazon.com/cloudhsm/classic/userguide/),
// and the AWS CloudHSM Classic API Reference (http://docs.aws.amazon.com/cloudhsm/classic/APIReference/).
//
// For information about the current version of AWS CloudHSM, see AWS CloudHSM
// (http://aws.amazon.com/cloudhsm/), the AWS CloudHSM User Guide (http://docs.aws.amazon.com/cloudhsm/latest/userguide/),
// and the AWS CloudHSM API Reference (http://docs.aws.amazon.com/cloudhsm/latest/APIReference/).
//
// Lists the high-availability partition groups for the account.
//
// This operation supports pagination with the use of the NextToken member.
// If more results are available, the NextToken member of the response contains
// a token that you pass in the next call to ListHapgs to retrieve the next
// set of items.
//
//    // Example sending a request using ListHapgsRequest.
//    req := client.ListHapgsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsm-2014-05-30/ListHapgs
func (c *Client) ListHapgsRequest(input *ListHapgsInput) ListHapgsRequest {
	op := &aws.Operation{
		Name:       opListHapgs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListHapgsInput{}
	}

	req := c.newRequest(op, input, &ListHapgsOutput{})

	return ListHapgsRequest{Request: req, Input: input, Copy: c.ListHapgsRequest}
}

// ListHapgsRequest is the request type for the
// ListHapgs API operation.
type ListHapgsRequest struct {
	*aws.Request
	Input *ListHapgsInput
	Copy  func(*ListHapgsInput) ListHapgsRequest
}

// Send marshals and sends the ListHapgs API request.
func (r ListHapgsRequest) Send(ctx context.Context) (*ListHapgsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListHapgsResponse{
		ListHapgsOutput: r.Request.Data.(*ListHapgsOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListHapgsResponse is the response type for the
// ListHapgs API operation.
type ListHapgsResponse struct {
	*ListHapgsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListHapgs request.
func (r *ListHapgsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
