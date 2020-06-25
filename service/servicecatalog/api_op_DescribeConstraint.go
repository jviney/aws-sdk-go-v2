// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeConstraintInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The identifier of the constraint.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeConstraintInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeConstraintInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeConstraintInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeConstraintOutput struct {
	_ struct{} `type:"structure"`

	// Information about the constraint.
	ConstraintDetail *ConstraintDetail `type:"structure"`

	// The constraint parameters.
	ConstraintParameters *string `type:"string"`

	// The status of the current request.
	Status Status `type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeConstraintOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeConstraint = "DescribeConstraint"

// DescribeConstraintRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Gets information about the specified constraint.
//
//    // Example sending a request using DescribeConstraintRequest.
//    req := client.DescribeConstraintRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribeConstraint
func (c *Client) DescribeConstraintRequest(input *DescribeConstraintInput) DescribeConstraintRequest {
	op := &aws.Operation{
		Name:       opDescribeConstraint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeConstraintInput{}
	}

	req := c.newRequest(op, input, &DescribeConstraintOutput{})

	return DescribeConstraintRequest{Request: req, Input: input, Copy: c.DescribeConstraintRequest}
}

// DescribeConstraintRequest is the request type for the
// DescribeConstraint API operation.
type DescribeConstraintRequest struct {
	*aws.Request
	Input *DescribeConstraintInput
	Copy  func(*DescribeConstraintInput) DescribeConstraintRequest
}

// Send marshals and sends the DescribeConstraint API request.
func (r DescribeConstraintRequest) Send(ctx context.Context) (*DescribeConstraintResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeConstraintResponse{
		DescribeConstraintOutput: r.Request.Data.(*DescribeConstraintOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeConstraintResponse is the response type for the
// DescribeConstraint API operation.
type DescribeConstraintResponse struct {
	*DescribeConstraintOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeConstraint request.
func (r *DescribeConstraintResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
