// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package frauddetector

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetRulesInput struct {
	_ struct{} `type:"structure"`

	// The detector ID.
	//
	// DetectorId is a required field
	DetectorId *string `locationName:"detectorId" min:"1" type:"string" required:"true"`

	// The maximum number of rules to return for the request.
	MaxResults *int64 `locationName:"maxResults" min:"50" type:"integer"`

	// The next page token.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The rule ID.
	RuleId *string `locationName:"ruleId" min:"1" type:"string"`

	// The rule version.
	RuleVersion *string `locationName:"ruleVersion" min:"1" type:"string"`
}

// String returns the string representation
func (s GetRulesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRulesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRulesInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 50 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 50))
	}
	if s.RuleId != nil && len(*s.RuleId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleId", 1))
	}
	if s.RuleVersion != nil && len(*s.RuleVersion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleVersion", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetRulesOutput struct {
	_ struct{} `type:"structure"`

	// The next page token to be used in subsequent requests.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The details of the requested rule.
	RuleDetails []RuleDetail `locationName:"ruleDetails" type:"list"`
}

// String returns the string representation
func (s GetRulesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRules = "GetRules"

// GetRulesRequest returns a request value for making API operation for
// Amazon Fraud Detector.
//
// Gets all rules available for the specified detector.
//
//    // Example sending a request using GetRulesRequest.
//    req := client.GetRulesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/frauddetector-2019-11-15/GetRules
func (c *Client) GetRulesRequest(input *GetRulesInput) GetRulesRequest {
	op := &aws.Operation{
		Name:       opGetRules,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetRulesInput{}
	}

	req := c.newRequest(op, input, &GetRulesOutput{})

	return GetRulesRequest{Request: req, Input: input, Copy: c.GetRulesRequest}
}

// GetRulesRequest is the request type for the
// GetRules API operation.
type GetRulesRequest struct {
	*aws.Request
	Input *GetRulesInput
	Copy  func(*GetRulesInput) GetRulesRequest
}

// Send marshals and sends the GetRules API request.
func (r GetRulesRequest) Send(ctx context.Context) (*GetRulesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRulesResponse{
		GetRulesOutput: r.Request.Data.(*GetRulesOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetRulesRequestPaginator returns a paginator for GetRules.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetRulesRequest(input)
//   p := frauddetector.NewGetRulesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetRulesPaginator(req GetRulesRequest) GetRulesPaginator {
	return GetRulesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetRulesInput
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

// GetRulesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetRulesPaginator struct {
	aws.Pager
}

func (p *GetRulesPaginator) CurrentPage() *GetRulesOutput {
	return p.Pager.CurrentPage().(*GetRulesOutput)
}

// GetRulesResponse is the response type for the
// GetRules API operation.
type GetRulesResponse struct {
	*GetRulesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRules request.
func (r *GetRulesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
