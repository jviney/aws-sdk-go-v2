// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticinference

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DescribeAcceleratorsInput struct {
	_ struct{} `type:"structure"`

	// The IDs of the accelerators to describe.
	AcceleratorIds []string `locationName:"acceleratorIds" type:"list"`

	// One or more filters. Filter names and values are case-sensitive. Valid filter
	// names are: accelerator-types: can provide a list of accelerator type names
	// to filter for. instance-id: can provide a list of EC2 instance ids to filter
	// for.
	Filters []Filter `locationName:"filters" type:"list"`

	// The total number of items to return in the command's output. If the total
	// number of items available is more than the value specified, a NextToken is
	// provided in the command's output. To resume pagination, provide the NextToken
	// value in the starting-token argument of a subsequent command. Do not use
	// the NextToken response element directly outside of the AWS CLI.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// A token to specify where to start paginating. This is the NextToken from
	// a previously truncated response.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeAcceleratorsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAcceleratorsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAcceleratorsInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeAcceleratorsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AcceleratorIds != nil {
		v := s.AcceleratorIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "acceleratorIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Filters != nil {
		v := s.Filters

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "filters", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeAcceleratorsOutput struct {
	_ struct{} `type:"structure"`

	// The details of the Elastic Inference Accelerators.
	AcceleratorSet []ElasticInferenceAccelerator `locationName:"acceleratorSet" type:"list"`

	// A token to specify where to start paginating. This is the NextToken from
	// a previously truncated response.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeAcceleratorsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeAcceleratorsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AcceleratorSet != nil {
		v := s.AcceleratorSet

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "acceleratorSet", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeAccelerators = "DescribeAccelerators"

// DescribeAcceleratorsRequest returns a request value for making API operation for
// Amazon Elastic  Inference.
//
// Describes information over a provided set of accelerators belonging to an
// account.
//
//    // Example sending a request using DescribeAcceleratorsRequest.
//    req := client.DescribeAcceleratorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elastic-inference-2017-07-25/DescribeAccelerators
func (c *Client) DescribeAcceleratorsRequest(input *DescribeAcceleratorsInput) DescribeAcceleratorsRequest {
	op := &aws.Operation{
		Name:       opDescribeAccelerators,
		HTTPMethod: "POST",
		HTTPPath:   "/describe-accelerators",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeAcceleratorsInput{}
	}

	req := c.newRequest(op, input, &DescribeAcceleratorsOutput{})

	return DescribeAcceleratorsRequest{Request: req, Input: input, Copy: c.DescribeAcceleratorsRequest}
}

// DescribeAcceleratorsRequest is the request type for the
// DescribeAccelerators API operation.
type DescribeAcceleratorsRequest struct {
	*aws.Request
	Input *DescribeAcceleratorsInput
	Copy  func(*DescribeAcceleratorsInput) DescribeAcceleratorsRequest
}

// Send marshals and sends the DescribeAccelerators API request.
func (r DescribeAcceleratorsRequest) Send(ctx context.Context) (*DescribeAcceleratorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAcceleratorsResponse{
		DescribeAcceleratorsOutput: r.Request.Data.(*DescribeAcceleratorsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeAcceleratorsRequestPaginator returns a paginator for DescribeAccelerators.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeAcceleratorsRequest(input)
//   p := elasticinference.NewDescribeAcceleratorsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeAcceleratorsPaginator(req DescribeAcceleratorsRequest) DescribeAcceleratorsPaginator {
	return DescribeAcceleratorsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeAcceleratorsInput
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

// DescribeAcceleratorsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeAcceleratorsPaginator struct {
	aws.Pager
}

func (p *DescribeAcceleratorsPaginator) CurrentPage() *DescribeAcceleratorsOutput {
	return p.Pager.CurrentPage().(*DescribeAcceleratorsOutput)
}

// DescribeAcceleratorsResponse is the response type for the
// DescribeAccelerators API operation.
type DescribeAcceleratorsResponse struct {
	*DescribeAcceleratorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAccelerators request.
func (r *DescribeAcceleratorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
