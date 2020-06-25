// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type CreateCampaignInput struct {
	_ struct{} `type:"structure" payload:"WriteCampaignRequest"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// Specifies the configuration and other settings for a campaign.
	//
	// WriteCampaignRequest is a required field
	WriteCampaignRequest *WriteCampaignRequest `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateCampaignInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCampaignInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCampaignInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.WriteCampaignRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("WriteCampaignRequest"))
	}
	if s.WriteCampaignRequest != nil {
		if err := s.WriteCampaignRequest.Validate(); err != nil {
			invalidParams.AddNested("WriteCampaignRequest", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateCampaignInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.WriteCampaignRequest != nil {
		v := s.WriteCampaignRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "WriteCampaignRequest", v, metadata)
	}
	return nil
}

type CreateCampaignOutput struct {
	_ struct{} `type:"structure" payload:"CampaignResponse"`

	// Provides information about the status, configuration, and other settings
	// for a campaign.
	//
	// CampaignResponse is a required field
	CampaignResponse *CampaignResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateCampaignOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateCampaignOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CampaignResponse != nil {
		v := s.CampaignResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "CampaignResponse", v, metadata)
	}
	return nil
}

const opCreateCampaign = "CreateCampaign"

// CreateCampaignRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Creates a new campaign for an application or updates the settings of an existing
// campaign for an application.
//
//    // Example sending a request using CreateCampaignRequest.
//    req := client.CreateCampaignRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/CreateCampaign
func (c *Client) CreateCampaignRequest(input *CreateCampaignInput) CreateCampaignRequest {
	op := &aws.Operation{
		Name:       opCreateCampaign,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apps/{application-id}/campaigns",
	}

	if input == nil {
		input = &CreateCampaignInput{}
	}

	req := c.newRequest(op, input, &CreateCampaignOutput{})

	return CreateCampaignRequest{Request: req, Input: input, Copy: c.CreateCampaignRequest}
}

// CreateCampaignRequest is the request type for the
// CreateCampaign API operation.
type CreateCampaignRequest struct {
	*aws.Request
	Input *CreateCampaignInput
	Copy  func(*CreateCampaignInput) CreateCampaignRequest
}

// Send marshals and sends the CreateCampaign API request.
func (r CreateCampaignRequest) Send(ctx context.Context) (*CreateCampaignResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCampaignResponse{
		CreateCampaignOutput: r.Request.Data.(*CreateCampaignOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCampaignResponse is the response type for the
// CreateCampaign API operation.
type CreateCampaignResponse struct {
	*CreateCampaignOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCampaign request.
func (r *CreateCampaignResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
