// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteMaintenanceWindowInput struct {
	_ struct{} `type:"structure"`

	// The ID of the maintenance window to delete.
	//
	// WindowId is a required field
	WindowId *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMaintenanceWindowInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMaintenanceWindowInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteMaintenanceWindowInput"}

	if s.WindowId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WindowId"))
	}
	if s.WindowId != nil && len(*s.WindowId) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("WindowId", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteMaintenanceWindowOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the deleted maintenance window.
	WindowId *string `min:"20" type:"string"`
}

// String returns the string representation
func (s DeleteMaintenanceWindowOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteMaintenanceWindow = "DeleteMaintenanceWindow"

// DeleteMaintenanceWindowRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Deletes a maintenance window.
//
//    // Example sending a request using DeleteMaintenanceWindowRequest.
//    req := client.DeleteMaintenanceWindowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeleteMaintenanceWindow
func (c *Client) DeleteMaintenanceWindowRequest(input *DeleteMaintenanceWindowInput) DeleteMaintenanceWindowRequest {
	op := &aws.Operation{
		Name:       opDeleteMaintenanceWindow,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteMaintenanceWindowInput{}
	}

	req := c.newRequest(op, input, &DeleteMaintenanceWindowOutput{})

	return DeleteMaintenanceWindowRequest{Request: req, Input: input, Copy: c.DeleteMaintenanceWindowRequest}
}

// DeleteMaintenanceWindowRequest is the request type for the
// DeleteMaintenanceWindow API operation.
type DeleteMaintenanceWindowRequest struct {
	*aws.Request
	Input *DeleteMaintenanceWindowInput
	Copy  func(*DeleteMaintenanceWindowInput) DeleteMaintenanceWindowRequest
}

// Send marshals and sends the DeleteMaintenanceWindow API request.
func (r DeleteMaintenanceWindowRequest) Send(ctx context.Context) (*DeleteMaintenanceWindowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteMaintenanceWindowResponse{
		DeleteMaintenanceWindowOutput: r.Request.Data.(*DeleteMaintenanceWindowOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteMaintenanceWindowResponse is the response type for the
// DeleteMaintenanceWindow API operation.
type DeleteMaintenanceWindowResponse struct {
	*DeleteMaintenanceWindowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteMaintenanceWindow request.
func (r *DeleteMaintenanceWindowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
