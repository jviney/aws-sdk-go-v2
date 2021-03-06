// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeEmergencyContactSettingsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeEmergencyContactSettingsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeEmergencyContactSettingsOutput struct {
	_ struct{} `type:"structure"`

	// A list of email addresses that the DRT can use to contact you during a suspected
	// attack.
	EmergencyContactList []EmergencyContact `type:"list"`
}

// String returns the string representation
func (s DescribeEmergencyContactSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEmergencyContactSettings = "DescribeEmergencyContactSettings"

// DescribeEmergencyContactSettingsRequest returns a request value for making API operation for
// AWS Shield.
//
// Lists the email addresses that the DRT can use to contact you during a suspected
// attack.
//
//    // Example sending a request using DescribeEmergencyContactSettingsRequest.
//    req := client.DescribeEmergencyContactSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/DescribeEmergencyContactSettings
func (c *Client) DescribeEmergencyContactSettingsRequest(input *DescribeEmergencyContactSettingsInput) DescribeEmergencyContactSettingsRequest {
	op := &aws.Operation{
		Name:       opDescribeEmergencyContactSettings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEmergencyContactSettingsInput{}
	}

	req := c.newRequest(op, input, &DescribeEmergencyContactSettingsOutput{})

	return DescribeEmergencyContactSettingsRequest{Request: req, Input: input, Copy: c.DescribeEmergencyContactSettingsRequest}
}

// DescribeEmergencyContactSettingsRequest is the request type for the
// DescribeEmergencyContactSettings API operation.
type DescribeEmergencyContactSettingsRequest struct {
	*aws.Request
	Input *DescribeEmergencyContactSettingsInput
	Copy  func(*DescribeEmergencyContactSettingsInput) DescribeEmergencyContactSettingsRequest
}

// Send marshals and sends the DescribeEmergencyContactSettings API request.
func (r DescribeEmergencyContactSettingsRequest) Send(ctx context.Context) (*DescribeEmergencyContactSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEmergencyContactSettingsResponse{
		DescribeEmergencyContactSettingsOutput: r.Request.Data.(*DescribeEmergencyContactSettingsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEmergencyContactSettingsResponse is the response type for the
// DescribeEmergencyContactSettings API operation.
type DescribeEmergencyContactSettingsResponse struct {
	*DescribeEmergencyContactSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEmergencyContactSettings request.
func (r *DescribeEmergencyContactSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
