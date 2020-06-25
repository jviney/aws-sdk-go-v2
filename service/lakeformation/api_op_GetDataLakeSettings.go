// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetDataLakeSettingsInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your AWS Lake
	// Formation environment.
	CatalogId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetDataLakeSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDataLakeSettingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDataLakeSettingsInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetDataLakeSettingsOutput struct {
	_ struct{} `type:"structure"`

	// A list of AWS Lake Formation principals.
	DataLakeSettings *DataLakeSettings `type:"structure"`
}

// String returns the string representation
func (s GetDataLakeSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDataLakeSettings = "GetDataLakeSettings"

// GetDataLakeSettingsRequest returns a request value for making API operation for
// AWS Lake Formation.
//
// The AWS Lake Formation principal.
//
//    // Example sending a request using GetDataLakeSettingsRequest.
//    req := client.GetDataLakeSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lakeformation-2017-03-31/GetDataLakeSettings
func (c *Client) GetDataLakeSettingsRequest(input *GetDataLakeSettingsInput) GetDataLakeSettingsRequest {
	op := &aws.Operation{
		Name:       opGetDataLakeSettings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDataLakeSettingsInput{}
	}

	req := c.newRequest(op, input, &GetDataLakeSettingsOutput{})

	return GetDataLakeSettingsRequest{Request: req, Input: input, Copy: c.GetDataLakeSettingsRequest}
}

// GetDataLakeSettingsRequest is the request type for the
// GetDataLakeSettings API operation.
type GetDataLakeSettingsRequest struct {
	*aws.Request
	Input *GetDataLakeSettingsInput
	Copy  func(*GetDataLakeSettingsInput) GetDataLakeSettingsRequest
}

// Send marshals and sends the GetDataLakeSettings API request.
func (r GetDataLakeSettingsRequest) Send(ctx context.Context) (*GetDataLakeSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDataLakeSettingsResponse{
		GetDataLakeSettingsOutput: r.Request.Data.(*GetDataLakeSettingsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDataLakeSettingsResponse is the response type for the
// GetDataLakeSettings API operation.
type GetDataLakeSettingsResponse struct {
	*GetDataLakeSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDataLakeSettings request.
func (r *GetDataLakeSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
