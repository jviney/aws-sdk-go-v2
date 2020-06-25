// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package health

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeAffectedEntitiesForOrganizationInput struct {
	_ struct{} `type:"structure"`

	// The locale (language) to return information in. English (en) is the default
	// and the only supported value at this time.
	Locale *string `locationName:"locale" min:"2" type:"string"`

	// The maximum number of items to return in one batch, between 10 and 100, inclusive.
	MaxResults *int64 `locationName:"maxResults" min:"10" type:"integer"`

	// If the results of a search are large, only a portion of the results are returned,
	// and a nextToken pagination token is returned in the response. To retrieve
	// the next batch of results, reissue the search request and include the returned
	// token. When all results have been returned, the response does not contain
	// a pagination token value.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`

	// A JSON set of elements including the awsAccountId and the eventArn.
	//
	// OrganizationEntityFilters is a required field
	OrganizationEntityFilters []EventAccountFilter `locationName:"organizationEntityFilters" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeAffectedEntitiesForOrganizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAffectedEntitiesForOrganizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAffectedEntitiesForOrganizationInput"}
	if s.Locale != nil && len(*s.Locale) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Locale", 2))
	}
	if s.MaxResults != nil && *s.MaxResults < 10 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 10))
	}
	if s.NextToken != nil && len(*s.NextToken) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 4))
	}

	if s.OrganizationEntityFilters == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationEntityFilters"))
	}
	if s.OrganizationEntityFilters != nil && len(s.OrganizationEntityFilters) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OrganizationEntityFilters", 1))
	}
	if s.OrganizationEntityFilters != nil {
		for i, v := range s.OrganizationEntityFilters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "OrganizationEntityFilters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeAffectedEntitiesForOrganizationOutput struct {
	_ struct{} `type:"structure"`

	// A JSON set of elements including the awsAccountId and its entityArn, entityValue
	// and its entityArn, lastUpdatedTime, statusCode, and tags.
	Entities []AffectedEntity `locationName:"entities" type:"list"`

	// A JSON set of elements of the failed response, including the awsAccountId,
	// errorMessage, errorName, and eventArn.
	FailedSet []OrganizationAffectedEntitiesErrorItem `locationName:"failedSet" type:"list"`

	// If the results of a search are large, only a portion of the results are returned,
	// and a nextToken pagination token is returned in the response. To retrieve
	// the next batch of results, reissue the search request and include the returned
	// token. When all results have been returned, the response does not contain
	// a pagination token value.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`
}

// String returns the string representation
func (s DescribeAffectedEntitiesForOrganizationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAffectedEntitiesForOrganization = "DescribeAffectedEntitiesForOrganization"

// DescribeAffectedEntitiesForOrganizationRequest returns a request value for making API operation for
// AWS Health APIs and Notifications.
//
// Returns a list of entities that have been affected by one or more events
// for one or more accounts in your organization in AWS Organizations, based
// on the filter criteria. Entities can refer to individual customer resources,
// groups of customer resources, or any other construct, depending on the AWS
// service.
//
// At least one event ARN and account ID are required. Results are sorted by
// the lastUpdatedTime of the entity, starting with the most recent.
//
// Before you can call this operation, you must first enable AWS Health to work
// with AWS Organizations. To do this, call the EnableHealthServiceAccessForOrganization
// operation from your organization's master account.
//
//    // Example sending a request using DescribeAffectedEntitiesForOrganizationRequest.
//    req := client.DescribeAffectedEntitiesForOrganizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/health-2016-08-04/DescribeAffectedEntitiesForOrganization
func (c *Client) DescribeAffectedEntitiesForOrganizationRequest(input *DescribeAffectedEntitiesForOrganizationInput) DescribeAffectedEntitiesForOrganizationRequest {
	op := &aws.Operation{
		Name:       opDescribeAffectedEntitiesForOrganization,
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
		input = &DescribeAffectedEntitiesForOrganizationInput{}
	}

	req := c.newRequest(op, input, &DescribeAffectedEntitiesForOrganizationOutput{})

	return DescribeAffectedEntitiesForOrganizationRequest{Request: req, Input: input, Copy: c.DescribeAffectedEntitiesForOrganizationRequest}
}

// DescribeAffectedEntitiesForOrganizationRequest is the request type for the
// DescribeAffectedEntitiesForOrganization API operation.
type DescribeAffectedEntitiesForOrganizationRequest struct {
	*aws.Request
	Input *DescribeAffectedEntitiesForOrganizationInput
	Copy  func(*DescribeAffectedEntitiesForOrganizationInput) DescribeAffectedEntitiesForOrganizationRequest
}

// Send marshals and sends the DescribeAffectedEntitiesForOrganization API request.
func (r DescribeAffectedEntitiesForOrganizationRequest) Send(ctx context.Context) (*DescribeAffectedEntitiesForOrganizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAffectedEntitiesForOrganizationResponse{
		DescribeAffectedEntitiesForOrganizationOutput: r.Request.Data.(*DescribeAffectedEntitiesForOrganizationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeAffectedEntitiesForOrganizationRequestPaginator returns a paginator for DescribeAffectedEntitiesForOrganization.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeAffectedEntitiesForOrganizationRequest(input)
//   p := health.NewDescribeAffectedEntitiesForOrganizationRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeAffectedEntitiesForOrganizationPaginator(req DescribeAffectedEntitiesForOrganizationRequest) DescribeAffectedEntitiesForOrganizationPaginator {
	return DescribeAffectedEntitiesForOrganizationPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeAffectedEntitiesForOrganizationInput
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

// DescribeAffectedEntitiesForOrganizationPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeAffectedEntitiesForOrganizationPaginator struct {
	aws.Pager
}

func (p *DescribeAffectedEntitiesForOrganizationPaginator) CurrentPage() *DescribeAffectedEntitiesForOrganizationOutput {
	return p.Pager.CurrentPage().(*DescribeAffectedEntitiesForOrganizationOutput)
}

// DescribeAffectedEntitiesForOrganizationResponse is the response type for the
// DescribeAffectedEntitiesForOrganization API operation.
type DescribeAffectedEntitiesForOrganizationResponse struct {
	*DescribeAffectedEntitiesForOrganizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAffectedEntitiesForOrganization request.
func (r *DescribeAffectedEntitiesForOrganizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
