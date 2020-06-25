// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeCampaignInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the campaign.
	//
	// CampaignArn is a required field
	CampaignArn *string `locationName:"campaignArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeCampaignInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCampaignInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCampaignInput"}

	if s.CampaignArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CampaignArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeCampaignOutput struct {
	_ struct{} `type:"structure"`

	// The properties of the campaign.
	Campaign *Campaign `locationName:"campaign" type:"structure"`
}

// String returns the string representation
func (s DescribeCampaignOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCampaign = "DescribeCampaign"

// DescribeCampaignRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Describes the given campaign, including its status.
//
// A campaign can be in one of the following states:
//
//    * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
//    * DELETE PENDING > DELETE IN_PROGRESS
//
// When the status is CREATE FAILED, the response includes the failureReason
// key, which describes why.
//
// For more information on campaigns, see CreateCampaign.
//
//    // Example sending a request using DescribeCampaignRequest.
//    req := client.DescribeCampaignRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/DescribeCampaign
func (c *Client) DescribeCampaignRequest(input *DescribeCampaignInput) DescribeCampaignRequest {
	op := &aws.Operation{
		Name:       opDescribeCampaign,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCampaignInput{}
	}

	req := c.newRequest(op, input, &DescribeCampaignOutput{})

	return DescribeCampaignRequest{Request: req, Input: input, Copy: c.DescribeCampaignRequest}
}

// DescribeCampaignRequest is the request type for the
// DescribeCampaign API operation.
type DescribeCampaignRequest struct {
	*aws.Request
	Input *DescribeCampaignInput
	Copy  func(*DescribeCampaignInput) DescribeCampaignRequest
}

// Send marshals and sends the DescribeCampaign API request.
func (r DescribeCampaignRequest) Send(ctx context.Context) (*DescribeCampaignResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCampaignResponse{
		DescribeCampaignOutput: r.Request.Data.(*DescribeCampaignOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCampaignResponse is the response type for the
// DescribeCampaign API operation.
type DescribeCampaignResponse struct {
	*DescribeCampaignOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCampaign request.
func (r *DescribeCampaignResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
