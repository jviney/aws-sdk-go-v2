// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeCertificateInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the certificate.
	//
	// CertificateId is a required field
	CertificateId *string `type:"string" required:"true"`

	// The identifier of the directory.
	//
	// DirectoryId is a required field
	DirectoryId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCertificateInput"}

	if s.CertificateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateId"))
	}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeCertificateOutput struct {
	_ struct{} `type:"structure"`

	// Information about the certificate, including registered date time, certificate
	// state, the reason for the state, expiration date time, and certificate common
	// name.
	Certificate *Certificate `type:"structure"`
}

// String returns the string representation
func (s DescribeCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCertificate = "DescribeCertificate"

// DescribeCertificateRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Displays information about the certificate registered for a secured LDAP
// connection.
//
//    // Example sending a request using DescribeCertificateRequest.
//    req := client.DescribeCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/DescribeCertificate
func (c *Client) DescribeCertificateRequest(input *DescribeCertificateInput) DescribeCertificateRequest {
	op := &aws.Operation{
		Name:       opDescribeCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCertificateInput{}
	}

	req := c.newRequest(op, input, &DescribeCertificateOutput{})

	return DescribeCertificateRequest{Request: req, Input: input, Copy: c.DescribeCertificateRequest}
}

// DescribeCertificateRequest is the request type for the
// DescribeCertificate API operation.
type DescribeCertificateRequest struct {
	*aws.Request
	Input *DescribeCertificateInput
	Copy  func(*DescribeCertificateInput) DescribeCertificateRequest
}

// Send marshals and sends the DescribeCertificate API request.
func (r DescribeCertificateRequest) Send(ctx context.Context) (*DescribeCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCertificateResponse{
		DescribeCertificateOutput: r.Request.Data.(*DescribeCertificateOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCertificateResponse is the response type for the
// DescribeCertificate API operation.
type DescribeCertificateResponse struct {
	*DescribeCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCertificate request.
func (r *DescribeCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
