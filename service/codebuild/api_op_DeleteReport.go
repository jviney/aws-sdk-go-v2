// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteReportInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the report to delete.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteReportInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteReportInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteReportInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteReportOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteReportOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteReport = "DeleteReport"

// DeleteReportRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// Deletes a report.
//
//    // Example sending a request using DeleteReportRequest.
//    req := client.DeleteReportRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/DeleteReport
func (c *Client) DeleteReportRequest(input *DeleteReportInput) DeleteReportRequest {
	op := &aws.Operation{
		Name:       opDeleteReport,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteReportInput{}
	}

	req := c.newRequest(op, input, &DeleteReportOutput{})

	return DeleteReportRequest{Request: req, Input: input, Copy: c.DeleteReportRequest}
}

// DeleteReportRequest is the request type for the
// DeleteReport API operation.
type DeleteReportRequest struct {
	*aws.Request
	Input *DeleteReportInput
	Copy  func(*DeleteReportInput) DeleteReportRequest
}

// Send marshals and sends the DeleteReport API request.
func (r DeleteReportRequest) Send(ctx context.Context) (*DeleteReportResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteReportResponse{
		DeleteReportOutput: r.Request.Data.(*DeleteReportOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteReportResponse is the response type for the
// DeleteReport API operation.
type DeleteReportResponse struct {
	*DeleteReportOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteReport request.
func (r *DeleteReportResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
