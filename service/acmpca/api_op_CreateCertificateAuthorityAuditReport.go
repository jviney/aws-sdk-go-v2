// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acmpca

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateCertificateAuthorityAuditReportInput struct {
	_ struct{} `type:"structure"`

	// The format in which to create the report. This can be either JSON or CSV.
	//
	// AuditReportResponseFormat is a required field
	AuditReportResponseFormat AuditReportResponseFormat `type:"string" required:"true" enum:"true"`

	// The Amazon Resource Name (ARN) of the CA to be audited. This is of the form:
	//
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012 .
	//
	// CertificateAuthorityArn is a required field
	CertificateAuthorityArn *string `min:"5" type:"string" required:"true"`

	// The name of the S3 bucket that will contain the audit report.
	//
	// S3BucketName is a required field
	S3BucketName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateCertificateAuthorityAuditReportInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCertificateAuthorityAuditReportInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCertificateAuthorityAuditReportInput"}
	if len(s.AuditReportResponseFormat) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("AuditReportResponseFormat"))
	}

	if s.CertificateAuthorityArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if s.CertificateAuthorityArn != nil && len(*s.CertificateAuthorityArn) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateAuthorityArn", 5))
	}

	if s.S3BucketName == nil {
		invalidParams.Add(aws.NewErrParamRequired("S3BucketName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateCertificateAuthorityAuditReportOutput struct {
	_ struct{} `type:"structure"`

	// An alphanumeric string that contains a report identifier.
	AuditReportId *string `min:"36" type:"string"`

	// The key that uniquely identifies the report file in your S3 bucket.
	S3Key *string `type:"string"`
}

// String returns the string representation
func (s CreateCertificateAuthorityAuditReportOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateCertificateAuthorityAuditReport = "CreateCertificateAuthorityAuditReport"

// CreateCertificateAuthorityAuditReportRequest returns a request value for making API operation for
// AWS Certificate Manager Private Certificate Authority.
//
// Creates an audit report that lists every time that your CA private key is
// used. The report is saved in the Amazon S3 bucket that you specify on input.
// The IssueCertificate and RevokeCertificate actions use the private key.
//
//    // Example sending a request using CreateCertificateAuthorityAuditReportRequest.
//    req := client.CreateCertificateAuthorityAuditReportRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/CreateCertificateAuthorityAuditReport
func (c *Client) CreateCertificateAuthorityAuditReportRequest(input *CreateCertificateAuthorityAuditReportInput) CreateCertificateAuthorityAuditReportRequest {
	op := &aws.Operation{
		Name:       opCreateCertificateAuthorityAuditReport,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCertificateAuthorityAuditReportInput{}
	}

	req := c.newRequest(op, input, &CreateCertificateAuthorityAuditReportOutput{})

	return CreateCertificateAuthorityAuditReportRequest{Request: req, Input: input, Copy: c.CreateCertificateAuthorityAuditReportRequest}
}

// CreateCertificateAuthorityAuditReportRequest is the request type for the
// CreateCertificateAuthorityAuditReport API operation.
type CreateCertificateAuthorityAuditReportRequest struct {
	*aws.Request
	Input *CreateCertificateAuthorityAuditReportInput
	Copy  func(*CreateCertificateAuthorityAuditReportInput) CreateCertificateAuthorityAuditReportRequest
}

// Send marshals and sends the CreateCertificateAuthorityAuditReport API request.
func (r CreateCertificateAuthorityAuditReportRequest) Send(ctx context.Context) (*CreateCertificateAuthorityAuditReportResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCertificateAuthorityAuditReportResponse{
		CreateCertificateAuthorityAuditReportOutput: r.Request.Data.(*CreateCertificateAuthorityAuditReportOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCertificateAuthorityAuditReportResponse is the response type for the
// CreateCertificateAuthorityAuditReport API operation.
type CreateCertificateAuthorityAuditReportResponse struct {
	*CreateCertificateAuthorityAuditReportOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCertificateAuthorityAuditReport request.
func (r *CreateCertificateAuthorityAuditReportResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
