// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acmpca

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeCertificateAuthorityAuditReportInput struct {
	_ struct{} `type:"structure"`

	// The report ID returned by calling the CreateCertificateAuthorityAuditReport
	// action.
	//
	// AuditReportId is a required field
	AuditReportId *string `min:"36" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the private CA. This must be of the form:
	//
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012 .
	//
	// CertificateAuthorityArn is a required field
	CertificateAuthorityArn *string `min:"5" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeCertificateAuthorityAuditReportInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCertificateAuthorityAuditReportInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCertificateAuthorityAuditReportInput"}

	if s.AuditReportId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AuditReportId"))
	}
	if s.AuditReportId != nil && len(*s.AuditReportId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("AuditReportId", 36))
	}

	if s.CertificateAuthorityArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if s.CertificateAuthorityArn != nil && len(*s.CertificateAuthorityArn) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateAuthorityArn", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeCertificateAuthorityAuditReportOutput struct {
	_ struct{} `type:"structure"`

	// Specifies whether report creation is in progress, has succeeded, or has failed.
	AuditReportStatus AuditReportStatus `type:"string" enum:"true"`

	// The date and time at which the report was created.
	CreatedAt *time.Time `type:"timestamp"`

	// Name of the S3 bucket that contains the report.
	S3BucketName *string `type:"string"`

	// S3 key that uniquely identifies the report file in your S3 bucket.
	S3Key *string `type:"string"`
}

// String returns the string representation
func (s DescribeCertificateAuthorityAuditReportOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCertificateAuthorityAuditReport = "DescribeCertificateAuthorityAuditReport"

// DescribeCertificateAuthorityAuditReportRequest returns a request value for making API operation for
// AWS Certificate Manager Private Certificate Authority.
//
// Lists information about a specific audit report created by calling the CreateCertificateAuthorityAuditReport
// action. Audit information is created every time the certificate authority
// (CA) private key is used. The private key is used when you call the IssueCertificate
// action or the RevokeCertificate action.
//
//    // Example sending a request using DescribeCertificateAuthorityAuditReportRequest.
//    req := client.DescribeCertificateAuthorityAuditReportRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/DescribeCertificateAuthorityAuditReport
func (c *Client) DescribeCertificateAuthorityAuditReportRequest(input *DescribeCertificateAuthorityAuditReportInput) DescribeCertificateAuthorityAuditReportRequest {
	op := &aws.Operation{
		Name:       opDescribeCertificateAuthorityAuditReport,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCertificateAuthorityAuditReportInput{}
	}

	req := c.newRequest(op, input, &DescribeCertificateAuthorityAuditReportOutput{})

	return DescribeCertificateAuthorityAuditReportRequest{Request: req, Input: input, Copy: c.DescribeCertificateAuthorityAuditReportRequest}
}

// DescribeCertificateAuthorityAuditReportRequest is the request type for the
// DescribeCertificateAuthorityAuditReport API operation.
type DescribeCertificateAuthorityAuditReportRequest struct {
	*aws.Request
	Input *DescribeCertificateAuthorityAuditReportInput
	Copy  func(*DescribeCertificateAuthorityAuditReportInput) DescribeCertificateAuthorityAuditReportRequest
}

// Send marshals and sends the DescribeCertificateAuthorityAuditReport API request.
func (r DescribeCertificateAuthorityAuditReportRequest) Send(ctx context.Context) (*DescribeCertificateAuthorityAuditReportResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCertificateAuthorityAuditReportResponse{
		DescribeCertificateAuthorityAuditReportOutput: r.Request.Data.(*DescribeCertificateAuthorityAuditReportOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCertificateAuthorityAuditReportResponse is the response type for the
// DescribeCertificateAuthorityAuditReport API operation.
type DescribeCertificateAuthorityAuditReportResponse struct {
	*DescribeCertificateAuthorityAuditReportOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCertificateAuthorityAuditReport request.
func (r *DescribeCertificateAuthorityAuditReportResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
