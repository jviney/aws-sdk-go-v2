// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeProvisionedProductPlanInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The maximum number of items to return with this call.
	PageSize *int64 `type:"integer"`

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string `type:"string"`

	// The plan identifier.
	//
	// PlanId is a required field
	PlanId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeProvisionedProductPlanInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeProvisionedProductPlanInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeProvisionedProductPlanInput"}

	if s.PlanId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlanId"))
	}
	if s.PlanId != nil && len(*s.PlanId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PlanId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeProvisionedProductPlanOutput struct {
	_ struct{} `type:"structure"`

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string `type:"string"`

	// Information about the plan.
	ProvisionedProductPlanDetails *ProvisionedProductPlanDetails `type:"structure"`

	// Information about the resource changes that will occur when the plan is executed.
	ResourceChanges []ResourceChange `type:"list"`
}

// String returns the string representation
func (s DescribeProvisionedProductPlanOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeProvisionedProductPlan = "DescribeProvisionedProductPlan"

// DescribeProvisionedProductPlanRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Gets information about the resource changes for the specified plan.
//
//    // Example sending a request using DescribeProvisionedProductPlanRequest.
//    req := client.DescribeProvisionedProductPlanRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribeProvisionedProductPlan
func (c *Client) DescribeProvisionedProductPlanRequest(input *DescribeProvisionedProductPlanInput) DescribeProvisionedProductPlanRequest {
	op := &aws.Operation{
		Name:       opDescribeProvisionedProductPlan,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeProvisionedProductPlanInput{}
	}

	req := c.newRequest(op, input, &DescribeProvisionedProductPlanOutput{})

	return DescribeProvisionedProductPlanRequest{Request: req, Input: input, Copy: c.DescribeProvisionedProductPlanRequest}
}

// DescribeProvisionedProductPlanRequest is the request type for the
// DescribeProvisionedProductPlan API operation.
type DescribeProvisionedProductPlanRequest struct {
	*aws.Request
	Input *DescribeProvisionedProductPlanInput
	Copy  func(*DescribeProvisionedProductPlanInput) DescribeProvisionedProductPlanRequest
}

// Send marshals and sends the DescribeProvisionedProductPlan API request.
func (r DescribeProvisionedProductPlanRequest) Send(ctx context.Context) (*DescribeProvisionedProductPlanResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeProvisionedProductPlanResponse{
		DescribeProvisionedProductPlanOutput: r.Request.Data.(*DescribeProvisionedProductPlanOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeProvisionedProductPlanResponse is the response type for the
// DescribeProvisionedProductPlan API operation.
type DescribeProvisionedProductPlanResponse struct {
	*DescribeProvisionedProductPlanOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeProvisionedProductPlan request.
func (r *DescribeProvisionedProductPlanResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
