// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListProvisionedProductPlansInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The access level to use to obtain results. The default is User.
	AccessLevelFilter *AccessLevelFilter `type:"structure"`

	// The maximum number of items to return with this call.
	PageSize *int64 `type:"integer"`

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string `type:"string"`

	// The product identifier.
	ProvisionProductId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListProvisionedProductPlansInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListProvisionedProductPlansInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListProvisionedProductPlansInput"}
	if s.ProvisionProductId != nil && len(*s.ProvisionProductId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProvisionProductId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListProvisionedProductPlansOutput struct {
	_ struct{} `type:"structure"`

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string `type:"string"`

	// Information about the plans.
	ProvisionedProductPlans []ProvisionedProductPlanSummary `type:"list"`
}

// String returns the string representation
func (s ListProvisionedProductPlansOutput) String() string {
	return awsutil.Prettify(s)
}

const opListProvisionedProductPlans = "ListProvisionedProductPlans"

// ListProvisionedProductPlansRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Lists the plans for the specified provisioned product or all plans to which
// the user has access.
//
//    // Example sending a request using ListProvisionedProductPlansRequest.
//    req := client.ListProvisionedProductPlansRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/ListProvisionedProductPlans
func (c *Client) ListProvisionedProductPlansRequest(input *ListProvisionedProductPlansInput) ListProvisionedProductPlansRequest {
	op := &aws.Operation{
		Name:       opListProvisionedProductPlans,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListProvisionedProductPlansInput{}
	}

	req := c.newRequest(op, input, &ListProvisionedProductPlansOutput{})

	return ListProvisionedProductPlansRequest{Request: req, Input: input, Copy: c.ListProvisionedProductPlansRequest}
}

// ListProvisionedProductPlansRequest is the request type for the
// ListProvisionedProductPlans API operation.
type ListProvisionedProductPlansRequest struct {
	*aws.Request
	Input *ListProvisionedProductPlansInput
	Copy  func(*ListProvisionedProductPlansInput) ListProvisionedProductPlansRequest
}

// Send marshals and sends the ListProvisionedProductPlans API request.
func (r ListProvisionedProductPlansRequest) Send(ctx context.Context) (*ListProvisionedProductPlansResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListProvisionedProductPlansResponse{
		ListProvisionedProductPlansOutput: r.Request.Data.(*ListProvisionedProductPlansOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListProvisionedProductPlansResponse is the response type for the
// ListProvisionedProductPlans API operation.
type ListProvisionedProductPlansResponse struct {
	*ListProvisionedProductPlansOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListProvisionedProductPlans request.
func (r *ListProvisionedProductPlansResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
