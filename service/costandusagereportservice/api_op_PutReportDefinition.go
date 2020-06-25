// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costandusagereportservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Creates a Cost and Usage Report.
type PutReportDefinitionInput struct {
	_ struct{} `type:"structure"`

	// Represents the output of the PutReportDefinition operation. The content consists
	// of the detailed metadata and data file information.
	//
	// ReportDefinition is a required field
	ReportDefinition *ReportDefinition `type:"structure" required:"true"`
}

// String returns the string representation
func (s PutReportDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutReportDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutReportDefinitionInput"}

	if s.ReportDefinition == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReportDefinition"))
	}
	if s.ReportDefinition != nil {
		if err := s.ReportDefinition.Validate(); err != nil {
			invalidParams.AddNested("ReportDefinition", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// If the action is successful, the service sends back an HTTP 200 response
// with an empty HTTP body.
type PutReportDefinitionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutReportDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutReportDefinition = "PutReportDefinition"

// PutReportDefinitionRequest returns a request value for making API operation for
// AWS Cost and Usage Report Service.
//
// Creates a new report using the description that you provide.
//
//    // Example sending a request using PutReportDefinitionRequest.
//    req := client.PutReportDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cur-2017-01-06/PutReportDefinition
func (c *Client) PutReportDefinitionRequest(input *PutReportDefinitionInput) PutReportDefinitionRequest {
	op := &aws.Operation{
		Name:       opPutReportDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutReportDefinitionInput{}
	}

	req := c.newRequest(op, input, &PutReportDefinitionOutput{})

	return PutReportDefinitionRequest{Request: req, Input: input, Copy: c.PutReportDefinitionRequest}
}

// PutReportDefinitionRequest is the request type for the
// PutReportDefinition API operation.
type PutReportDefinitionRequest struct {
	*aws.Request
	Input *PutReportDefinitionInput
	Copy  func(*PutReportDefinitionInput) PutReportDefinitionRequest
}

// Send marshals and sends the PutReportDefinition API request.
func (r PutReportDefinitionRequest) Send(ctx context.Context) (*PutReportDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutReportDefinitionResponse{
		PutReportDefinitionOutput: r.Request.Data.(*PutReportDefinitionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutReportDefinitionResponse is the response type for the
// PutReportDefinition API operation.
type PutReportDefinitionResponse struct {
	*PutReportDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutReportDefinition request.
func (r *PutReportDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
