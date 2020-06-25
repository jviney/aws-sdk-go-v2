// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a StopDeployment operation.
type StopDeploymentInput struct {
	_ struct{} `type:"structure"`

	// Indicates, when a deployment is stopped, whether instances that have been
	// updated should be rolled back to the previous version of the application
	// revision.
	AutoRollbackEnabled *bool `locationName:"autoRollbackEnabled" type:"boolean"`

	// The unique ID of a deployment.
	//
	// DeploymentId is a required field
	DeploymentId *string `locationName:"deploymentId" type:"string" required:"true"`
}

// String returns the string representation
func (s StopDeploymentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopDeploymentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopDeploymentInput"}

	if s.DeploymentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a StopDeployment operation.
type StopDeploymentOutput struct {
	_ struct{} `type:"structure"`

	// The status of the stop deployment operation:
	//
	//    * Pending: The stop operation is pending.
	//
	//    * Succeeded: The stop operation was successful.
	Status StopStatus `locationName:"status" type:"string" enum:"true"`

	// An accompanying status message.
	StatusMessage *string `locationName:"statusMessage" type:"string"`
}

// String returns the string representation
func (s StopDeploymentOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopDeployment = "StopDeployment"

// StopDeploymentRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Attempts to stop an ongoing deployment.
//
//    // Example sending a request using StopDeploymentRequest.
//    req := client.StopDeploymentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/StopDeployment
func (c *Client) StopDeploymentRequest(input *StopDeploymentInput) StopDeploymentRequest {
	op := &aws.Operation{
		Name:       opStopDeployment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopDeploymentInput{}
	}

	req := c.newRequest(op, input, &StopDeploymentOutput{})

	return StopDeploymentRequest{Request: req, Input: input, Copy: c.StopDeploymentRequest}
}

// StopDeploymentRequest is the request type for the
// StopDeployment API operation.
type StopDeploymentRequest struct {
	*aws.Request
	Input *StopDeploymentInput
	Copy  func(*StopDeploymentInput) StopDeploymentRequest
}

// Send marshals and sends the StopDeployment API request.
func (r StopDeploymentRequest) Send(ctx context.Context) (*StopDeploymentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopDeploymentResponse{
		StopDeploymentOutput: r.Request.Data.(*StopDeploymentOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopDeploymentResponse is the response type for the
// StopDeployment API operation.
type StopDeploymentResponse struct {
	*StopDeploymentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopDeployment request.
func (r *StopDeploymentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
