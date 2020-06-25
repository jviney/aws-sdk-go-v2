// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a GetApplication operation.
type GetApplicationInput struct {
	_ struct{} `type:"structure"`

	// The name of an AWS CodeDeploy application associated with the IAM user or
	// AWS account.
	//
	// ApplicationName is a required field
	ApplicationName *string `locationName:"applicationName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetApplicationInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a GetApplication operation.
type GetApplicationOutput struct {
	_ struct{} `type:"structure"`

	// Information about the application.
	Application *ApplicationInfo `locationName:"application" type:"structure"`
}

// String returns the string representation
func (s GetApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetApplication = "GetApplication"

// GetApplicationRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Gets information about an application.
//
//    // Example sending a request using GetApplicationRequest.
//    req := client.GetApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/GetApplication
func (c *Client) GetApplicationRequest(input *GetApplicationInput) GetApplicationRequest {
	op := &aws.Operation{
		Name:       opGetApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetApplicationInput{}
	}

	req := c.newRequest(op, input, &GetApplicationOutput{})

	return GetApplicationRequest{Request: req, Input: input, Copy: c.GetApplicationRequest}
}

// GetApplicationRequest is the request type for the
// GetApplication API operation.
type GetApplicationRequest struct {
	*aws.Request
	Input *GetApplicationInput
	Copy  func(*GetApplicationInput) GetApplicationRequest
}

// Send marshals and sends the GetApplication API request.
func (r GetApplicationRequest) Send(ctx context.Context) (*GetApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetApplicationResponse{
		GetApplicationOutput: r.Request.Data.(*GetApplicationOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetApplicationResponse is the response type for the
// GetApplication API operation.
type GetApplicationResponse struct {
	*GetApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetApplication request.
func (r *GetApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
