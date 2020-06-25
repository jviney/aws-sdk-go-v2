// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListKeyPoliciesInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the customer master key (CMK).
	//
	// Specify the key ID or the Amazon Resource Name (ARN) of the CMK.
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// KeyId is a required field
	KeyId *string `min:"1" type:"string" required:"true"`

	// Use this parameter to specify the maximum number of items to return. When
	// this value is present, AWS KMS does not return more than the specified number
	// of items, but it might return fewer.
	//
	// This value is optional. If you include a value, it must be between 1 and
	// 1000, inclusive. If you do not include a value, it defaults to 100.
	//
	// Only one policy can be attached to a key.
	Limit *int64 `min:"1" type:"integer"`

	// Use this parameter in a subsequent request after you receive a response with
	// truncated results. Set it to the value of NextMarker from the truncated response
	// you just received.
	Marker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListKeyPoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListKeyPoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListKeyPoliciesInput"}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListKeyPoliciesOutput struct {
	_ struct{} `type:"structure"`

	// When Truncated is true, this element is present and contains the value to
	// use for the Marker parameter in a subsequent request.
	NextMarker *string `min:"1" type:"string"`

	// A list of key policy names. The only valid value is default.
	PolicyNames []string `type:"list"`

	// A flag that indicates whether there are more items in the list. When this
	// value is true, the list in this response is truncated. To get more items,
	// pass the value of the NextMarker element in thisresponse to the Marker parameter
	// in a subsequent request.
	Truncated *bool `type:"boolean"`
}

// String returns the string representation
func (s ListKeyPoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListKeyPolicies = "ListKeyPolicies"

// ListKeyPoliciesRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Gets the names of the key policies that are attached to a customer master
// key (CMK). This operation is designed to get policy names that you can use
// in a GetKeyPolicy operation. However, the only valid policy name is default.
// You cannot perform this operation on a CMK in a different AWS account.
//
//    // Example sending a request using ListKeyPoliciesRequest.
//    req := client.ListKeyPoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/ListKeyPolicies
func (c *Client) ListKeyPoliciesRequest(input *ListKeyPoliciesInput) ListKeyPoliciesRequest {
	op := &aws.Operation{
		Name:       opListKeyPolicies,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"NextMarker"},
			LimitToken:      "Limit",
			TruncationToken: "Truncated",
		},
	}

	if input == nil {
		input = &ListKeyPoliciesInput{}
	}

	req := c.newRequest(op, input, &ListKeyPoliciesOutput{})

	return ListKeyPoliciesRequest{Request: req, Input: input, Copy: c.ListKeyPoliciesRequest}
}

// ListKeyPoliciesRequest is the request type for the
// ListKeyPolicies API operation.
type ListKeyPoliciesRequest struct {
	*aws.Request
	Input *ListKeyPoliciesInput
	Copy  func(*ListKeyPoliciesInput) ListKeyPoliciesRequest
}

// Send marshals and sends the ListKeyPolicies API request.
func (r ListKeyPoliciesRequest) Send(ctx context.Context) (*ListKeyPoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListKeyPoliciesResponse{
		ListKeyPoliciesOutput: r.Request.Data.(*ListKeyPoliciesOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListKeyPoliciesRequestPaginator returns a paginator for ListKeyPolicies.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListKeyPoliciesRequest(input)
//   p := kms.NewListKeyPoliciesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListKeyPoliciesPaginator(req ListKeyPoliciesRequest) ListKeyPoliciesPaginator {
	return ListKeyPoliciesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListKeyPoliciesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListKeyPoliciesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListKeyPoliciesPaginator struct {
	aws.Pager
}

func (p *ListKeyPoliciesPaginator) CurrentPage() *ListKeyPoliciesOutput {
	return p.Pager.CurrentPage().(*ListKeyPoliciesOutput)
}

// ListKeyPoliciesResponse is the response type for the
// ListKeyPolicies API operation.
type ListKeyPoliciesResponse struct {
	*ListKeyPoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListKeyPolicies request.
func (r *ListKeyPoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
