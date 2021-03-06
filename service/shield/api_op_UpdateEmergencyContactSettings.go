// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type UpdateEmergencyContactSettingsInput struct {
	_ struct{} `type:"structure"`

	// A list of email addresses that the DRT can use to contact you during a suspected
	// attack.
	EmergencyContactList []EmergencyContact `type:"list"`
}

// String returns the string representation
func (s UpdateEmergencyContactSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateEmergencyContactSettingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateEmergencyContactSettingsInput"}
	if s.EmergencyContactList != nil {
		for i, v := range s.EmergencyContactList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "EmergencyContactList", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateEmergencyContactSettingsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateEmergencyContactSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateEmergencyContactSettings = "UpdateEmergencyContactSettings"

// UpdateEmergencyContactSettingsRequest returns a request value for making API operation for
// AWS Shield.
//
// Updates the details of the list of email addresses that the DRT can use to
// contact you during a suspected attack.
//
//    // Example sending a request using UpdateEmergencyContactSettingsRequest.
//    req := client.UpdateEmergencyContactSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/UpdateEmergencyContactSettings
func (c *Client) UpdateEmergencyContactSettingsRequest(input *UpdateEmergencyContactSettingsInput) UpdateEmergencyContactSettingsRequest {
	op := &aws.Operation{
		Name:       opUpdateEmergencyContactSettings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateEmergencyContactSettingsInput{}
	}

	req := c.newRequest(op, input, &UpdateEmergencyContactSettingsOutput{})

	return UpdateEmergencyContactSettingsRequest{Request: req, Input: input, Copy: c.UpdateEmergencyContactSettingsRequest}
}

// UpdateEmergencyContactSettingsRequest is the request type for the
// UpdateEmergencyContactSettings API operation.
type UpdateEmergencyContactSettingsRequest struct {
	*aws.Request
	Input *UpdateEmergencyContactSettingsInput
	Copy  func(*UpdateEmergencyContactSettingsInput) UpdateEmergencyContactSettingsRequest
}

// Send marshals and sends the UpdateEmergencyContactSettings API request.
func (r UpdateEmergencyContactSettingsRequest) Send(ctx context.Context) (*UpdateEmergencyContactSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateEmergencyContactSettingsResponse{
		UpdateEmergencyContactSettingsOutput: r.Request.Data.(*UpdateEmergencyContactSettingsOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateEmergencyContactSettingsResponse is the response type for the
// UpdateEmergencyContactSettings API operation.
type UpdateEmergencyContactSettingsResponse struct {
	*UpdateEmergencyContactSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateEmergencyContactSettings request.
func (r *UpdateEmergencyContactSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
