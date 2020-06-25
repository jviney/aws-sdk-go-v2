// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package licensemanager

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteLicenseConfigurationInput struct {
	_ struct{} `type:"structure"`

	// ID of the license configuration.
	//
	// LicenseConfigurationArn is a required field
	LicenseConfigurationArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteLicenseConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLicenseConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLicenseConfigurationInput"}

	if s.LicenseConfigurationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("LicenseConfigurationArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteLicenseConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteLicenseConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteLicenseConfiguration = "DeleteLicenseConfiguration"

// DeleteLicenseConfigurationRequest returns a request value for making API operation for
// AWS License Manager.
//
// Deletes the specified license configuration.
//
// You cannot delete a license configuration that is in use.
//
//    // Example sending a request using DeleteLicenseConfigurationRequest.
//    req := client.DeleteLicenseConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/license-manager-2018-08-01/DeleteLicenseConfiguration
func (c *Client) DeleteLicenseConfigurationRequest(input *DeleteLicenseConfigurationInput) DeleteLicenseConfigurationRequest {
	op := &aws.Operation{
		Name:       opDeleteLicenseConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteLicenseConfigurationInput{}
	}

	req := c.newRequest(op, input, &DeleteLicenseConfigurationOutput{})

	return DeleteLicenseConfigurationRequest{Request: req, Input: input, Copy: c.DeleteLicenseConfigurationRequest}
}

// DeleteLicenseConfigurationRequest is the request type for the
// DeleteLicenseConfiguration API operation.
type DeleteLicenseConfigurationRequest struct {
	*aws.Request
	Input *DeleteLicenseConfigurationInput
	Copy  func(*DeleteLicenseConfigurationInput) DeleteLicenseConfigurationRequest
}

// Send marshals and sends the DeleteLicenseConfiguration API request.
func (r DeleteLicenseConfigurationRequest) Send(ctx context.Context) (*DeleteLicenseConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteLicenseConfigurationResponse{
		DeleteLicenseConfigurationOutput: r.Request.Data.(*DeleteLicenseConfigurationOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteLicenseConfigurationResponse is the response type for the
// DeleteLicenseConfiguration API operation.
type DeleteLicenseConfigurationResponse struct {
	*DeleteLicenseConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteLicenseConfiguration request.
func (r *DeleteLicenseConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
