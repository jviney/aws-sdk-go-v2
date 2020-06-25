// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pricing

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeServicesInput struct {
	_ struct{} `type:"structure"`

	// The format version that you want the response to be in.
	//
	// Valid values are: aws_v1
	FormatVersion *string `type:"string"`

	// The maximum number of results that you want returned in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token that indicates the next set of results that you want
	// to retrieve.
	NextToken *string `type:"string"`

	// The code for the service whose information you want to retrieve, such as
	// AmazonEC2. You can use the ServiceCode to filter the results in a GetProducts
	// call. To retrieve a list of all services, leave this blank.
	ServiceCode *string `type:"string"`
}

// String returns the string representation
func (s DescribeServicesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeServicesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeServicesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeServicesOutput struct {
	_ struct{} `type:"structure"`

	// The format version of the response. For example, aws_v1.
	FormatVersion *string `type:"string"`

	// The pagination token for the next set of retreivable results.
	NextToken *string `type:"string"`

	// The service metadata for the service or services in the response.
	Services []Service `type:"list"`
}

// String returns the string representation
func (s DescribeServicesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeServices = "DescribeServices"

// DescribeServicesRequest returns a request value for making API operation for
// AWS Price List Service.
//
// Returns the metadata for one service or a list of the metadata for all services.
// Use this without a service code to get the service codes for all services.
// Use it with a service code, such as AmazonEC2, to get information specific
// to that service, such as the attribute names available for that service.
// For example, some of the attribute names available for EC2 are volumeType,
// maxIopsVolume, operation, locationType, and instanceCapacity10xlarge.
//
//    // Example sending a request using DescribeServicesRequest.
//    req := client.DescribeServicesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15/DescribeServices
func (c *Client) DescribeServicesRequest(input *DescribeServicesInput) DescribeServicesRequest {
	op := &aws.Operation{
		Name:       opDescribeServices,
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
		input = &DescribeServicesInput{}
	}

	req := c.newRequest(op, input, &DescribeServicesOutput{})

	return DescribeServicesRequest{Request: req, Input: input, Copy: c.DescribeServicesRequest}
}

// DescribeServicesRequest is the request type for the
// DescribeServices API operation.
type DescribeServicesRequest struct {
	*aws.Request
	Input *DescribeServicesInput
	Copy  func(*DescribeServicesInput) DescribeServicesRequest
}

// Send marshals and sends the DescribeServices API request.
func (r DescribeServicesRequest) Send(ctx context.Context) (*DescribeServicesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeServicesResponse{
		DescribeServicesOutput: r.Request.Data.(*DescribeServicesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeServicesRequestPaginator returns a paginator for DescribeServices.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeServicesRequest(input)
//   p := pricing.NewDescribeServicesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeServicesPaginator(req DescribeServicesRequest) DescribeServicesPaginator {
	return DescribeServicesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeServicesInput
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

// DescribeServicesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeServicesPaginator struct {
	aws.Pager
}

func (p *DescribeServicesPaginator) CurrentPage() *DescribeServicesOutput {
	return p.Pager.CurrentPage().(*DescribeServicesOutput)
}

// DescribeServicesResponse is the response type for the
// DescribeServices API operation.
type DescribeServicesResponse struct {
	*DescribeServicesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeServices request.
func (r *DescribeServicesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
