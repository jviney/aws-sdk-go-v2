// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetAppsInput struct {
	_ struct{} `type:"structure"`

	PageSize *string `location:"querystring" locationName:"page-size" type:"string"`

	Token *string `location:"querystring" locationName:"token" type:"string"`
}

// String returns the string representation
func (s GetAppsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAppsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.PageSize != nil {
		v := *s.PageSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "page-size", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Token != nil {
		v := *s.Token

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetAppsOutput struct {
	_ struct{} `type:"structure" payload:"ApplicationsResponse"`

	// Provides information about all of your applications.
	//
	// ApplicationsResponse is a required field
	ApplicationsResponse *ApplicationsResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetAppsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAppsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApplicationsResponse != nil {
		v := s.ApplicationsResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "ApplicationsResponse", v, metadata)
	}
	return nil
}

const opGetApps = "GetApps"

// GetAppsRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about all the applications that are associated with
// your Amazon Pinpoint account.
//
//    // Example sending a request using GetAppsRequest.
//    req := client.GetAppsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetApps
func (c *Client) GetAppsRequest(input *GetAppsInput) GetAppsRequest {
	op := &aws.Operation{
		Name:       opGetApps,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps",
	}

	if input == nil {
		input = &GetAppsInput{}
	}

	req := c.newRequest(op, input, &GetAppsOutput{})

	return GetAppsRequest{Request: req, Input: input, Copy: c.GetAppsRequest}
}

// GetAppsRequest is the request type for the
// GetApps API operation.
type GetAppsRequest struct {
	*aws.Request
	Input *GetAppsInput
	Copy  func(*GetAppsInput) GetAppsRequest
}

// Send marshals and sends the GetApps API request.
func (r GetAppsRequest) Send(ctx context.Context) (*GetAppsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAppsResponse{
		GetAppsOutput: r.Request.Data.(*GetAppsOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAppsResponse is the response type for the
// GetApps API operation.
type GetAppsResponse struct {
	*GetAppsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetApps request.
func (r *GetAppsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
