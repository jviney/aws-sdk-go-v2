// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalytics

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteApplicationOutputInput struct {
	_ struct{} `type:"structure"`

	// Amazon Kinesis Analytics application name.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// Amazon Kinesis Analytics application version. You can use the DescribeApplication
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
	// operation to get the current application version. If the version specified
	// is not the current version, the ConcurrentModificationException is returned.
	//
	// CurrentApplicationVersionId is a required field
	CurrentApplicationVersionId *int64 `min:"1" type:"long" required:"true"`

	// The ID of the configuration to delete. Each output configuration that is
	// added to the application, either when the application is created or later
	// using the AddApplicationOutput (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_AddApplicationOutput.html)
	// operation, has a unique ID. You need to provide the ID to uniquely identify
	// the output configuration that you want to delete from the application configuration.
	// You can use the DescribeApplication (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
	// operation to get the specific OutputId.
	//
	// OutputId is a required field
	OutputId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteApplicationOutputInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteApplicationOutputInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteApplicationOutputInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if s.CurrentApplicationVersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CurrentApplicationVersionId"))
	}
	if s.CurrentApplicationVersionId != nil && *s.CurrentApplicationVersionId < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("CurrentApplicationVersionId", 1))
	}

	if s.OutputId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OutputId"))
	}
	if s.OutputId != nil && len(*s.OutputId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OutputId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteApplicationOutputOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteApplicationOutputOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteApplicationOutput = "DeleteApplicationOutput"

// DeleteApplicationOutputRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
//
// This documentation is for version 1 of the Amazon Kinesis Data Analytics
// API, which only supports SQL applications. Version 2 of the API supports
// SQL and Java applications. For more information about version 2, see Amazon
// Kinesis Data Analytics API V2 Documentation (/kinesisanalytics/latest/apiv2/Welcome.html).
//
// Deletes output destination configuration from your application configuration.
// Amazon Kinesis Analytics will no longer write data from the corresponding
// in-application stream to the external output destination.
//
// This operation requires permissions to perform the kinesisanalytics:DeleteApplicationOutput
// action.
//
//    // Example sending a request using DeleteApplicationOutputRequest.
//    req := client.DeleteApplicationOutputRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalytics-2015-08-14/DeleteApplicationOutput
func (c *Client) DeleteApplicationOutputRequest(input *DeleteApplicationOutputInput) DeleteApplicationOutputRequest {
	op := &aws.Operation{
		Name:       opDeleteApplicationOutput,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteApplicationOutputInput{}
	}

	req := c.newRequest(op, input, &DeleteApplicationOutputOutput{})

	return DeleteApplicationOutputRequest{Request: req, Input: input, Copy: c.DeleteApplicationOutputRequest}
}

// DeleteApplicationOutputRequest is the request type for the
// DeleteApplicationOutput API operation.
type DeleteApplicationOutputRequest struct {
	*aws.Request
	Input *DeleteApplicationOutputInput
	Copy  func(*DeleteApplicationOutputInput) DeleteApplicationOutputRequest
}

// Send marshals and sends the DeleteApplicationOutput API request.
func (r DeleteApplicationOutputRequest) Send(ctx context.Context) (*DeleteApplicationOutputResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteApplicationOutputResponse{
		DeleteApplicationOutputOutput: r.Request.Data.(*DeleteApplicationOutputOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteApplicationOutputResponse is the response type for the
// DeleteApplicationOutput API operation.
type DeleteApplicationOutputResponse struct {
	*DeleteApplicationOutputOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteApplicationOutput request.
func (r *DeleteApplicationOutputResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
