// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Initiates the creation of a conditional forwarder for your AWS Directory
// Service for Microsoft Active Directory. Conditional forwarders are required
// in order to set up a trust relationship with another domain.
type CreateConditionalForwarderInput struct {
	_ struct{} `type:"structure"`

	// The directory ID of the AWS directory for which you are creating the conditional
	// forwarder.
	//
	// DirectoryId is a required field
	DirectoryId *string `type:"string" required:"true"`

	// The IP addresses of the remote DNS server associated with RemoteDomainName.
	//
	// DnsIpAddrs is a required field
	DnsIpAddrs []string `type:"list" required:"true"`

	// The fully qualified domain name (FQDN) of the remote domain with which you
	// will set up a trust relationship.
	//
	// RemoteDomainName is a required field
	RemoteDomainName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateConditionalForwarderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateConditionalForwarderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateConditionalForwarderInput"}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}

	if s.DnsIpAddrs == nil {
		invalidParams.Add(aws.NewErrParamRequired("DnsIpAddrs"))
	}

	if s.RemoteDomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RemoteDomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of a CreateConditinalForwarder request.
type CreateConditionalForwarderOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateConditionalForwarderOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateConditionalForwarder = "CreateConditionalForwarder"

// CreateConditionalForwarderRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Creates a conditional forwarder associated with your AWS directory. Conditional
// forwarders are required in order to set up a trust relationship with another
// domain. The conditional forwarder points to the trusted domain.
//
//    // Example sending a request using CreateConditionalForwarderRequest.
//    req := client.CreateConditionalForwarderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/CreateConditionalForwarder
func (c *Client) CreateConditionalForwarderRequest(input *CreateConditionalForwarderInput) CreateConditionalForwarderRequest {
	op := &aws.Operation{
		Name:       opCreateConditionalForwarder,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateConditionalForwarderInput{}
	}

	req := c.newRequest(op, input, &CreateConditionalForwarderOutput{})

	return CreateConditionalForwarderRequest{Request: req, Input: input, Copy: c.CreateConditionalForwarderRequest}
}

// CreateConditionalForwarderRequest is the request type for the
// CreateConditionalForwarder API operation.
type CreateConditionalForwarderRequest struct {
	*aws.Request
	Input *CreateConditionalForwarderInput
	Copy  func(*CreateConditionalForwarderInput) CreateConditionalForwarderRequest
}

// Send marshals and sends the CreateConditionalForwarder API request.
func (r CreateConditionalForwarderRequest) Send(ctx context.Context) (*CreateConditionalForwarderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateConditionalForwarderResponse{
		CreateConditionalForwarderOutput: r.Request.Data.(*CreateConditionalForwarderOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateConditionalForwarderResponse is the response type for the
// CreateConditionalForwarder API operation.
type CreateConditionalForwarderResponse struct {
	*CreateConditionalForwarderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateConditionalForwarder request.
func (r *CreateConditionalForwarderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
