// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetOpsItemInput struct {
	_ struct{} `type:"structure"`

	// The ID of the OpsItem that you want to get.
	//
	// OpsItemId is a required field
	OpsItemId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetOpsItemInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetOpsItemInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetOpsItemInput"}

	if s.OpsItemId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OpsItemId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetOpsItemOutput struct {
	_ struct{} `type:"structure"`

	// The OpsItem.
	OpsItem *OpsItem `type:"structure"`
}

// String returns the string representation
func (s GetOpsItemOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetOpsItem = "GetOpsItem"

// GetOpsItemRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Get information about an OpsItem by using the ID. You must have permission
// in AWS Identity and Access Management (IAM) to view information about an
// OpsItem. For more information, see Getting started with OpsCenter (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-getting-started.html)
// in the AWS Systems Manager User Guide.
//
// Operations engineers and IT professionals use OpsCenter to view, investigate,
// and remediate operational issues impacting the performance and health of
// their AWS resources. For more information, see AWS Systems Manager OpsCenter
// (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter.html)
// in the AWS Systems Manager User Guide.
//
//    // Example sending a request using GetOpsItemRequest.
//    req := client.GetOpsItemRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetOpsItem
func (c *Client) GetOpsItemRequest(input *GetOpsItemInput) GetOpsItemRequest {
	op := &aws.Operation{
		Name:       opGetOpsItem,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetOpsItemInput{}
	}

	req := c.newRequest(op, input, &GetOpsItemOutput{})

	return GetOpsItemRequest{Request: req, Input: input, Copy: c.GetOpsItemRequest}
}

// GetOpsItemRequest is the request type for the
// GetOpsItem API operation.
type GetOpsItemRequest struct {
	*aws.Request
	Input *GetOpsItemInput
	Copy  func(*GetOpsItemInput) GetOpsItemRequest
}

// Send marshals and sends the GetOpsItem API request.
func (r GetOpsItemRequest) Send(ctx context.Context) (*GetOpsItemResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetOpsItemResponse{
		GetOpsItemOutput: r.Request.Data.(*GetOpsItemOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetOpsItemResponse is the response type for the
// GetOpsItem API operation.
type GetOpsItemResponse struct {
	*GetOpsItemOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetOpsItem request.
func (r *GetOpsItemResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
