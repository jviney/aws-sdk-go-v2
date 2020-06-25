// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/query"
)

type DeleteServerCertificateInput struct {
	_ struct{} `type:"structure"`

	// The name of the server certificate you want to delete.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// ServerCertificateName is a required field
	ServerCertificateName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteServerCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteServerCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteServerCertificateInput"}

	if s.ServerCertificateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerCertificateName"))
	}
	if s.ServerCertificateName != nil && len(*s.ServerCertificateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServerCertificateName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteServerCertificateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteServerCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteServerCertificate = "DeleteServerCertificate"

// DeleteServerCertificateRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Deletes the specified server certificate.
//
// For more information about working with server certificates, see Working
// with Server Certificates (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html)
// in the IAM User Guide. This topic also includes a list of AWS services that
// can use the server certificates that you manage with IAM.
//
// If you are using a server certificate with Elastic Load Balancing, deleting
// the certificate could have implications for your application. If Elastic
// Load Balancing doesn't detect the deletion of bound certificates, it may
// continue to use the certificates. This could cause Elastic Load Balancing
// to stop accepting traffic. We recommend that you remove the reference to
// the certificate from Elastic Load Balancing before using this command to
// delete the certificate. For more information, go to DeleteLoadBalancerListeners
// (https://docs.aws.amazon.com/ElasticLoadBalancing/latest/APIReference/API_DeleteLoadBalancerListeners.html)
// in the Elastic Load Balancing API Reference.
//
//    // Example sending a request using DeleteServerCertificateRequest.
//    req := client.DeleteServerCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeleteServerCertificate
func (c *Client) DeleteServerCertificateRequest(input *DeleteServerCertificateInput) DeleteServerCertificateRequest {
	op := &aws.Operation{
		Name:       opDeleteServerCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteServerCertificateInput{}
	}

	req := c.newRequest(op, input, &DeleteServerCertificateOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteServerCertificateRequest{Request: req, Input: input, Copy: c.DeleteServerCertificateRequest}
}

// DeleteServerCertificateRequest is the request type for the
// DeleteServerCertificate API operation.
type DeleteServerCertificateRequest struct {
	*aws.Request
	Input *DeleteServerCertificateInput
	Copy  func(*DeleteServerCertificateInput) DeleteServerCertificateRequest
}

// Send marshals and sends the DeleteServerCertificate API request.
func (r DeleteServerCertificateRequest) Send(ctx context.Context) (*DeleteServerCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteServerCertificateResponse{
		DeleteServerCertificateOutput: r.Request.Data.(*DeleteServerCertificateOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteServerCertificateResponse is the response type for the
// DeleteServerCertificate API operation.
type DeleteServerCertificateResponse struct {
	*DeleteServerCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteServerCertificate request.
func (r *DeleteServerCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
