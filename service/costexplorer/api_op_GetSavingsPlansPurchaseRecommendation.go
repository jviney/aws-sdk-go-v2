// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costexplorer

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetSavingsPlansPurchaseRecommendationInput struct {
	_ struct{} `type:"structure"`

	// The account scope that you want your recommendations for. Amazon Web Services
	// calculates recommendations including the payer account and linked accounts
	// if the value is set to PAYER. If the value is LINKED, recommendations are
	// calculated for individual linked accounts only.
	AccountScope AccountScope `type:"string" enum:"true"`

	// You can filter your recommendations by Account ID with the LINKED_ACCOUNT
	// dimension. To filter your recommendations by Account ID, specify Key as LINKED_ACCOUNT
	// and Value as the comma-separated Acount ID(s) for which you want to see Savings
	// Plans purchase recommendations.
	//
	// For GetSavingsPlansPurchaseRecommendation, the Filter does not include CostCategories
	// or Tags. It only includes Dimensions. With Dimensions, Key must be LINKED_ACCOUNT
	// and Value can be a single Account ID or multiple comma-separated Account
	// IDs for which you want to see Savings Plans Purchase Recommendations. AND
	// and OR operators are not supported.
	Filter *Expression `type:"structure"`

	// The lookback period used to generate the recommendation.
	//
	// LookbackPeriodInDays is a required field
	LookbackPeriodInDays LookbackPeriodInDays `type:"string" required:"true" enum:"true"`

	// The token to retrieve the next set of results. Amazon Web Services provides
	// the token when the response from a previous call has more results than the
	// maximum page size.
	NextPageToken *string `type:"string"`

	// The number of recommendations that you want returned in a single response
	// object.
	PageSize *int64 `type:"integer"`

	// The payment option used to generate these recommendations.
	//
	// PaymentOption is a required field
	PaymentOption PaymentOption `type:"string" required:"true" enum:"true"`

	// The Savings Plans recommendation type requested.
	//
	// SavingsPlansType is a required field
	SavingsPlansType SupportedSavingsPlansType `type:"string" required:"true" enum:"true"`

	// The savings plan recommendation term used to generated these recommendations.
	//
	// TermInYears is a required field
	TermInYears TermInYears `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s GetSavingsPlansPurchaseRecommendationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSavingsPlansPurchaseRecommendationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSavingsPlansPurchaseRecommendationInput"}
	if len(s.LookbackPeriodInDays) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("LookbackPeriodInDays"))
	}
	if len(s.PaymentOption) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("PaymentOption"))
	}
	if len(s.SavingsPlansType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("SavingsPlansType"))
	}
	if len(s.TermInYears) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("TermInYears"))
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

type GetSavingsPlansPurchaseRecommendationOutput struct {
	_ struct{} `type:"structure"`

	// Information regarding this specific recommendation set.
	Metadata *SavingsPlansPurchaseRecommendationMetadata `type:"structure"`

	// The token for the next set of retrievable results. AWS provides the token
	// when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string `type:"string"`

	// Contains your request parameters, Savings Plan Recommendations Summary, and
	// Details.
	SavingsPlansPurchaseRecommendation *SavingsPlansPurchaseRecommendation `type:"structure"`
}

// String returns the string representation
func (s GetSavingsPlansPurchaseRecommendationOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetSavingsPlansPurchaseRecommendation = "GetSavingsPlansPurchaseRecommendation"

// GetSavingsPlansPurchaseRecommendationRequest returns a request value for making API operation for
// AWS Cost Explorer Service.
//
// Retrieves your request parameters, Savings Plan Recommendations Summary and
// Details.
//
//    // Example sending a request using GetSavingsPlansPurchaseRecommendationRequest.
//    req := client.GetSavingsPlansPurchaseRecommendationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetSavingsPlansPurchaseRecommendation
func (c *Client) GetSavingsPlansPurchaseRecommendationRequest(input *GetSavingsPlansPurchaseRecommendationInput) GetSavingsPlansPurchaseRecommendationRequest {
	op := &aws.Operation{
		Name:       opGetSavingsPlansPurchaseRecommendation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSavingsPlansPurchaseRecommendationInput{}
	}

	req := c.newRequest(op, input, &GetSavingsPlansPurchaseRecommendationOutput{})

	return GetSavingsPlansPurchaseRecommendationRequest{Request: req, Input: input, Copy: c.GetSavingsPlansPurchaseRecommendationRequest}
}

// GetSavingsPlansPurchaseRecommendationRequest is the request type for the
// GetSavingsPlansPurchaseRecommendation API operation.
type GetSavingsPlansPurchaseRecommendationRequest struct {
	*aws.Request
	Input *GetSavingsPlansPurchaseRecommendationInput
	Copy  func(*GetSavingsPlansPurchaseRecommendationInput) GetSavingsPlansPurchaseRecommendationRequest
}

// Send marshals and sends the GetSavingsPlansPurchaseRecommendation API request.
func (r GetSavingsPlansPurchaseRecommendationRequest) Send(ctx context.Context) (*GetSavingsPlansPurchaseRecommendationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSavingsPlansPurchaseRecommendationResponse{
		GetSavingsPlansPurchaseRecommendationOutput: r.Request.Data.(*GetSavingsPlansPurchaseRecommendationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSavingsPlansPurchaseRecommendationResponse is the response type for the
// GetSavingsPlansPurchaseRecommendation API operation.
type GetSavingsPlansPurchaseRecommendationResponse struct {
	*GetSavingsPlansPurchaseRecommendationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSavingsPlansPurchaseRecommendation request.
func (r *GetSavingsPlansPurchaseRecommendationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
