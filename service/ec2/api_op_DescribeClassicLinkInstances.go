// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeClassicLinkInstancesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * group-id - The ID of a VPC security group that's associated with the
	//    instance.
	//
	//    * instance-id - The ID of the instance.
	//
	//    * tag:<key> - The key/value combination of a tag assigned to the resource.
	//    Use the tag key in the filter name and the tag value as the filter value.
	//    For example, to find all resources that have a tag with the key Owner
	//    and the value TeamA, specify tag:Owner for the filter name and TeamA for
	//    the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. Use this filter
	//    to find all resources assigned a tag with a specific key, regardless of
	//    the tag value.
	//
	//    * vpc-id - The ID of the VPC to which the instance is linked. vpc-id -
	//    The ID of the VPC that the instance is linked to.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// One or more instance IDs. Must be instances linked to a VPC through ClassicLink.
	InstanceIds []string `locationName:"InstanceId" locationNameList:"InstanceId" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	//
	// Constraint: If the value is greater than 1000, we return only 1000 items.
	MaxResults *int64 `locationName:"maxResults" min:"5" type:"integer"`

	// The token for the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeClassicLinkInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeClassicLinkInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeClassicLinkInstancesInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeClassicLinkInstancesOutput struct {
	_ struct{} `type:"structure"`

	// Information about one or more linked EC2-Classic instances.
	Instances []ClassicLinkInstance `locationName:"instancesSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeClassicLinkInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeClassicLinkInstances = "DescribeClassicLinkInstances"

// DescribeClassicLinkInstancesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more of your linked EC2-Classic instances. This request
// only returns information about EC2-Classic instances linked to a VPC through
// ClassicLink. You cannot use this request to return information about other
// instances.
//
//    // Example sending a request using DescribeClassicLinkInstancesRequest.
//    req := client.DescribeClassicLinkInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeClassicLinkInstances
func (c *Client) DescribeClassicLinkInstancesRequest(input *DescribeClassicLinkInstancesInput) DescribeClassicLinkInstancesRequest {
	op := &aws.Operation{
		Name:       opDescribeClassicLinkInstances,
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
		input = &DescribeClassicLinkInstancesInput{}
	}

	req := c.newRequest(op, input, &DescribeClassicLinkInstancesOutput{})

	return DescribeClassicLinkInstancesRequest{Request: req, Input: input, Copy: c.DescribeClassicLinkInstancesRequest}
}

// DescribeClassicLinkInstancesRequest is the request type for the
// DescribeClassicLinkInstances API operation.
type DescribeClassicLinkInstancesRequest struct {
	*aws.Request
	Input *DescribeClassicLinkInstancesInput
	Copy  func(*DescribeClassicLinkInstancesInput) DescribeClassicLinkInstancesRequest
}

// Send marshals and sends the DescribeClassicLinkInstances API request.
func (r DescribeClassicLinkInstancesRequest) Send(ctx context.Context) (*DescribeClassicLinkInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeClassicLinkInstancesResponse{
		DescribeClassicLinkInstancesOutput: r.Request.Data.(*DescribeClassicLinkInstancesOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeClassicLinkInstancesRequestPaginator returns a paginator for DescribeClassicLinkInstances.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeClassicLinkInstancesRequest(input)
//   p := ec2.NewDescribeClassicLinkInstancesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeClassicLinkInstancesPaginator(req DescribeClassicLinkInstancesRequest) DescribeClassicLinkInstancesPaginator {
	return DescribeClassicLinkInstancesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeClassicLinkInstancesInput
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

// DescribeClassicLinkInstancesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeClassicLinkInstancesPaginator struct {
	aws.Pager
}

func (p *DescribeClassicLinkInstancesPaginator) CurrentPage() *DescribeClassicLinkInstancesOutput {
	return p.Pager.CurrentPage().(*DescribeClassicLinkInstancesOutput)
}

// DescribeClassicLinkInstancesResponse is the response type for the
// DescribeClassicLinkInstances API operation.
type DescribeClassicLinkInstancesResponse struct {
	*DescribeClassicLinkInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeClassicLinkInstances request.
func (r *DescribeClassicLinkInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
