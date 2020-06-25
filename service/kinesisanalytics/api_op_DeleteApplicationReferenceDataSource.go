// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalytics

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteApplicationReferenceDataSourceInput struct {
	_ struct{} `type:"structure"`

	// Name of an existing application.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// Version of the application. You can use the DescribeApplication (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
	// operation to get the current application version. If the version specified
	// is not the current version, the ConcurrentModificationException is returned.
	//
	// CurrentApplicationVersionId is a required field
	CurrentApplicationVersionId *int64 `min:"1" type:"long" required:"true"`

	// ID of the reference data source. When you add a reference data source to
	// your application using the AddApplicationReferenceDataSource (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_AddApplicationReferenceDataSource.html),
	// Amazon Kinesis Analytics assigns an ID. You can use the DescribeApplication
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
	// operation to get the reference ID.
	//
	// ReferenceId is a required field
	ReferenceId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteApplicationReferenceDataSourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteApplicationReferenceDataSourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteApplicationReferenceDataSourceInput"}

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

	if s.ReferenceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReferenceId"))
	}
	if s.ReferenceId != nil && len(*s.ReferenceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ReferenceId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteApplicationReferenceDataSourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteApplicationReferenceDataSourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteApplicationReferenceDataSource = "DeleteApplicationReferenceDataSource"

// DeleteApplicationReferenceDataSourceRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
//
// This documentation is for version 1 of the Amazon Kinesis Data Analytics
// API, which only supports SQL applications. Version 2 of the API supports
// SQL and Java applications. For more information about version 2, see Amazon
// Kinesis Data Analytics API V2 Documentation (/kinesisanalytics/latest/apiv2/Welcome.html).
//
// Deletes a reference data source configuration from the specified application
// configuration.
//
// If the application is running, Amazon Kinesis Analytics immediately removes
// the in-application table that you created using the AddApplicationReferenceDataSource
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_AddApplicationReferenceDataSource.html)
// operation.
//
// This operation requires permissions to perform the kinesisanalytics.DeleteApplicationReferenceDataSource
// action.
//
//    // Example sending a request using DeleteApplicationReferenceDataSourceRequest.
//    req := client.DeleteApplicationReferenceDataSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalytics-2015-08-14/DeleteApplicationReferenceDataSource
func (c *Client) DeleteApplicationReferenceDataSourceRequest(input *DeleteApplicationReferenceDataSourceInput) DeleteApplicationReferenceDataSourceRequest {
	op := &aws.Operation{
		Name:       opDeleteApplicationReferenceDataSource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteApplicationReferenceDataSourceInput{}
	}

	req := c.newRequest(op, input, &DeleteApplicationReferenceDataSourceOutput{})

	return DeleteApplicationReferenceDataSourceRequest{Request: req, Input: input, Copy: c.DeleteApplicationReferenceDataSourceRequest}
}

// DeleteApplicationReferenceDataSourceRequest is the request type for the
// DeleteApplicationReferenceDataSource API operation.
type DeleteApplicationReferenceDataSourceRequest struct {
	*aws.Request
	Input *DeleteApplicationReferenceDataSourceInput
	Copy  func(*DeleteApplicationReferenceDataSourceInput) DeleteApplicationReferenceDataSourceRequest
}

// Send marshals and sends the DeleteApplicationReferenceDataSource API request.
func (r DeleteApplicationReferenceDataSourceRequest) Send(ctx context.Context) (*DeleteApplicationReferenceDataSourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteApplicationReferenceDataSourceResponse{
		DeleteApplicationReferenceDataSourceOutput: r.Request.Data.(*DeleteApplicationReferenceDataSourceOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteApplicationReferenceDataSourceResponse is the response type for the
// DeleteApplicationReferenceDataSource API operation.
type DeleteApplicationReferenceDataSourceResponse struct {
	*DeleteApplicationReferenceDataSourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteApplicationReferenceDataSource request.
func (r *DeleteApplicationReferenceDataSourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
