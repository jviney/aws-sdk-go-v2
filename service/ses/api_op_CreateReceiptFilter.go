// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to create a new IP address filter. You use IP address
// filters when you receive email with Amazon SES. For more information, see
// the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
type CreateReceiptFilterInput struct {
	_ struct{} `type:"structure"`

	// A data structure that describes the IP address filter to create, which consists
	// of a name, an IP address range, and whether to allow or block mail from it.
	//
	// Filter is a required field
	Filter *ReceiptFilter `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateReceiptFilterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateReceiptFilterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateReceiptFilterInput"}

	if s.Filter == nil {
		invalidParams.Add(aws.NewErrParamRequired("Filter"))
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			invalidParams.AddNested("Filter", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An empty element returned on a successful request.
type CreateReceiptFilterOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateReceiptFilterOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateReceiptFilter = "CreateReceiptFilter"

// CreateReceiptFilterRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Creates a new IP address filter.
//
// For information about setting up IP address filters, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-ip-filters.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using CreateReceiptFilterRequest.
//    req := client.CreateReceiptFilterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/CreateReceiptFilter
func (c *Client) CreateReceiptFilterRequest(input *CreateReceiptFilterInput) CreateReceiptFilterRequest {
	op := &aws.Operation{
		Name:       opCreateReceiptFilter,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateReceiptFilterInput{}
	}

	req := c.newRequest(op, input, &CreateReceiptFilterOutput{})

	return CreateReceiptFilterRequest{Request: req, Input: input, Copy: c.CreateReceiptFilterRequest}
}

// CreateReceiptFilterRequest is the request type for the
// CreateReceiptFilter API operation.
type CreateReceiptFilterRequest struct {
	*aws.Request
	Input *CreateReceiptFilterInput
	Copy  func(*CreateReceiptFilterInput) CreateReceiptFilterRequest
}

// Send marshals and sends the CreateReceiptFilter API request.
func (r CreateReceiptFilterRequest) Send(ctx context.Context) (*CreateReceiptFilterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateReceiptFilterResponse{
		CreateReceiptFilterOutput: r.Request.Data.(*CreateReceiptFilterOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateReceiptFilterResponse is the response type for the
// CreateReceiptFilter API operation.
type CreateReceiptFilterResponse struct {
	*CreateReceiptFilterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateReceiptFilter request.
func (r *CreateReceiptFilterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
