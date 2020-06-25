// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// The request body of the UpdateServiceSetting API action.
type UpdateServiceSettingInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the service setting to reset. For example,
	// arn:aws:ssm:us-east-1:111122223333:servicesetting/ssm/parameter-store/high-throughput-enabled.
	// The setting ID can be one of the following.
	//
	//    * /ssm/parameter-store/default-parameter-tier
	//
	//    * /ssm/parameter-store/high-throughput-enabled
	//
	//    * /ssm/managed-instance/activation-tier
	//
	// SettingId is a required field
	SettingId *string `min:"1" type:"string" required:"true"`

	// The new value to specify for the service setting. For the /ssm/parameter-store/default-parameter-tier
	// setting ID, the setting value can be one of the following.
	//
	//    * Standard
	//
	//    * Advanced
	//
	//    * Intelligent-Tiering
	//
	// For the /ssm/parameter-store/high-throughput-enabled, and /ssm/managed-instance/activation-tier
	// setting IDs, the setting value can be true or false.
	//
	// SettingValue is a required field
	SettingValue *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateServiceSettingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateServiceSettingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateServiceSettingInput"}

	if s.SettingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SettingId"))
	}
	if s.SettingId != nil && len(*s.SettingId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SettingId", 1))
	}

	if s.SettingValue == nil {
		invalidParams.Add(aws.NewErrParamRequired("SettingValue"))
	}
	if s.SettingValue != nil && len(*s.SettingValue) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SettingValue", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result body of the UpdateServiceSetting API action.
type UpdateServiceSettingOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateServiceSettingOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateServiceSetting = "UpdateServiceSetting"

// UpdateServiceSettingRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// ServiceSetting is an account-level setting for an AWS service. This setting
// defines how a user interacts with or uses a service or a feature of a service.
// For example, if an AWS service charges money to the account based on feature
// or service usage, then the AWS service team might create a default setting
// of "false". This means the user can't use this feature unless they change
// the setting to "true" and intentionally opt in for a paid feature.
//
// Services map a SettingId object to a setting value. AWS services teams define
// the default value for a SettingId. You can't create a new SettingId, but
// you can overwrite the default value if you have the ssm:UpdateServiceSetting
// permission for the setting. Use the GetServiceSetting API action to view
// the current value. Or, use the ResetServiceSetting to change the value back
// to the original value defined by the AWS service team.
//
// Update the service setting for the account.
//
//    // Example sending a request using UpdateServiceSettingRequest.
//    req := client.UpdateServiceSettingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/UpdateServiceSetting
func (c *Client) UpdateServiceSettingRequest(input *UpdateServiceSettingInput) UpdateServiceSettingRequest {
	op := &aws.Operation{
		Name:       opUpdateServiceSetting,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateServiceSettingInput{}
	}

	req := c.newRequest(op, input, &UpdateServiceSettingOutput{})

	return UpdateServiceSettingRequest{Request: req, Input: input, Copy: c.UpdateServiceSettingRequest}
}

// UpdateServiceSettingRequest is the request type for the
// UpdateServiceSetting API operation.
type UpdateServiceSettingRequest struct {
	*aws.Request
	Input *UpdateServiceSettingInput
	Copy  func(*UpdateServiceSettingInput) UpdateServiceSettingRequest
}

// Send marshals and sends the UpdateServiceSetting API request.
func (r UpdateServiceSettingRequest) Send(ctx context.Context) (*UpdateServiceSettingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateServiceSettingResponse{
		UpdateServiceSettingOutput: r.Request.Data.(*UpdateServiceSettingOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateServiceSettingResponse is the response type for the
// UpdateServiceSetting API operation.
type UpdateServiceSettingResponse struct {
	*UpdateServiceSettingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateServiceSetting request.
func (r *UpdateServiceSettingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
