// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeRecordInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The record identifier of the provisioned product. This identifier is returned
	// by the request operation.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// The maximum number of items to return with this call.
	PageSize *int64 `type:"integer"`

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeRecordInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRecordInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeRecordInput"}

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

type DescribeRecordOutput struct {
	_ struct{} `type:"structure"`

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string `type:"string"`

	// Information about the product.
	RecordDetail *RecordDetail `type:"structure"`

	// Information about the product created as the result of a request. For example,
	// the output for a CloudFormation-backed product that creates an S3 bucket
	// would include the S3 bucket URL.
	RecordOutputs []RecordOutput `type:"list"`
}

// String returns the string representation
func (s DescribeRecordOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeRecord = "DescribeRecord"

// DescribeRecordRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Gets information about the specified request operation.
//
// Use this operation after calling a request operation (for example, ProvisionProduct,
// TerminateProvisionedProduct, or UpdateProvisionedProduct).
//
// If a provisioned product was transferred to a new owner using UpdateProvisionedProductProperties,
// the new owner will be able to describe all past records for that product.
// The previous owner will no longer be able to describe the records, but will
// be able to use ListRecordHistory to see the product's history from when he
// was the owner.
//
//    // Example sending a request using DescribeRecordRequest.
//    req := client.DescribeRecordRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribeRecord
func (c *Client) DescribeRecordRequest(input *DescribeRecordInput) DescribeRecordRequest {
	op := &aws.Operation{
		Name:       opDescribeRecord,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeRecordInput{}
	}

	req := c.newRequest(op, input, &DescribeRecordOutput{})

	return DescribeRecordRequest{Request: req, Input: input, Copy: c.DescribeRecordRequest}
}

// DescribeRecordRequest is the request type for the
// DescribeRecord API operation.
type DescribeRecordRequest struct {
	*aws.Request
	Input *DescribeRecordInput
	Copy  func(*DescribeRecordInput) DescribeRecordRequest
}

// Send marshals and sends the DescribeRecord API request.
func (r DescribeRecordRequest) Send(ctx context.Context) (*DescribeRecordResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRecordResponse{
		DescribeRecordOutput: r.Request.Data.(*DescribeRecordOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeRecordResponse is the response type for the
// DescribeRecord API operation.
type DescribeRecordResponse struct {
	*DescribeRecordOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRecord request.
func (r *DescribeRecordResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
