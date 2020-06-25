// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeClientVpnAuthorizationRulesInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Client VPN endpoint.
	//
	// ClientVpnEndpointId is a required field
	ClientVpnEndpointId *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters. Filter names and values are case-sensitive.
	//
	//    * description - The description of the authorization rule.
	//
	//    * destination-cidr - The CIDR of the network to which the authorization
	//    rule applies.
	//
	//    * group-id - The ID of the Active Directory group to which the authorization
	//    rule grants access.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return for the request in a single page.
	// The remaining results can be seen by sending another request with the nextToken
	// value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token to retrieve the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeClientVpnAuthorizationRulesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeClientVpnAuthorizationRulesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeClientVpnAuthorizationRulesInput"}

	if s.ClientVpnEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientVpnEndpointId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeClientVpnAuthorizationRulesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the authorization rules.
	AuthorizationRules []AuthorizationRule `locationName:"authorizationRule" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeClientVpnAuthorizationRulesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeClientVpnAuthorizationRules = "DescribeClientVpnAuthorizationRules"

// DescribeClientVpnAuthorizationRulesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the authorization rules for a specified Client VPN endpoint.
//
//    // Example sending a request using DescribeClientVpnAuthorizationRulesRequest.
//    req := client.DescribeClientVpnAuthorizationRulesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeClientVpnAuthorizationRules
func (c *Client) DescribeClientVpnAuthorizationRulesRequest(input *DescribeClientVpnAuthorizationRulesInput) DescribeClientVpnAuthorizationRulesRequest {
	op := &aws.Operation{
		Name:       opDescribeClientVpnAuthorizationRules,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeClientVpnAuthorizationRulesInput{}
	}

	req := c.newRequest(op, input, &DescribeClientVpnAuthorizationRulesOutput{})

	return DescribeClientVpnAuthorizationRulesRequest{Request: req, Input: input, Copy: c.DescribeClientVpnAuthorizationRulesRequest}
}

// DescribeClientVpnAuthorizationRulesRequest is the request type for the
// DescribeClientVpnAuthorizationRules API operation.
type DescribeClientVpnAuthorizationRulesRequest struct {
	*aws.Request
	Input *DescribeClientVpnAuthorizationRulesInput
	Copy  func(*DescribeClientVpnAuthorizationRulesInput) DescribeClientVpnAuthorizationRulesRequest
}

// Send marshals and sends the DescribeClientVpnAuthorizationRules API request.
func (r DescribeClientVpnAuthorizationRulesRequest) Send(ctx context.Context) (*DescribeClientVpnAuthorizationRulesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeClientVpnAuthorizationRulesResponse{
		DescribeClientVpnAuthorizationRulesOutput: r.Request.Data.(*DescribeClientVpnAuthorizationRulesOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeClientVpnAuthorizationRulesRequestPaginator returns a paginator for DescribeClientVpnAuthorizationRules.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeClientVpnAuthorizationRulesRequest(input)
//   p := ec2.NewDescribeClientVpnAuthorizationRulesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeClientVpnAuthorizationRulesPaginator(req DescribeClientVpnAuthorizationRulesRequest) DescribeClientVpnAuthorizationRulesPaginator {
	return DescribeClientVpnAuthorizationRulesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeClientVpnAuthorizationRulesInput
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

// DescribeClientVpnAuthorizationRulesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeClientVpnAuthorizationRulesPaginator struct {
	aws.Pager
}

func (p *DescribeClientVpnAuthorizationRulesPaginator) CurrentPage() *DescribeClientVpnAuthorizationRulesOutput {
	return p.Pager.CurrentPage().(*DescribeClientVpnAuthorizationRulesOutput)
}

// DescribeClientVpnAuthorizationRulesResponse is the response type for the
// DescribeClientVpnAuthorizationRules API operation.
type DescribeClientVpnAuthorizationRulesResponse struct {
	*DescribeClientVpnAuthorizationRulesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeClientVpnAuthorizationRules request.
func (r *DescribeClientVpnAuthorizationRulesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
