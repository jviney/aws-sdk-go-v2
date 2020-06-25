// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type CreateGroupCertificateAuthorityInput struct {
	_ struct{} `type:"structure"`

	AmznClientToken *string `location:"header" locationName:"X-Amzn-Client-Token" type:"string"`

	// GroupId is a required field
	GroupId *string `location:"uri" locationName:"GroupId" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateGroupCertificateAuthorityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateGroupCertificateAuthorityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateGroupCertificateAuthorityInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateGroupCertificateAuthorityInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AmznClientToken != nil {
		v := *s.AmznClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Amzn-Client-Token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GroupId != nil {
		v := *s.GroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateGroupCertificateAuthorityOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the group certificate authority.
	GroupCertificateAuthorityArn *string `type:"string"`
}

// String returns the string representation
func (s CreateGroupCertificateAuthorityOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateGroupCertificateAuthorityOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.GroupCertificateAuthorityArn != nil {
		v := *s.GroupCertificateAuthorityArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GroupCertificateAuthorityArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateGroupCertificateAuthority = "CreateGroupCertificateAuthority"

// CreateGroupCertificateAuthorityRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Creates a CA for the group. If a CA already exists, it will rotate the existing
// CA.
//
//    // Example sending a request using CreateGroupCertificateAuthorityRequest.
//    req := client.CreateGroupCertificateAuthorityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateGroupCertificateAuthority
func (c *Client) CreateGroupCertificateAuthorityRequest(input *CreateGroupCertificateAuthorityInput) CreateGroupCertificateAuthorityRequest {
	op := &aws.Operation{
		Name:       opCreateGroupCertificateAuthority,
		HTTPMethod: "POST",
		HTTPPath:   "/greengrass/groups/{GroupId}/certificateauthorities",
	}

	if input == nil {
		input = &CreateGroupCertificateAuthorityInput{}
	}

	req := c.newRequest(op, input, &CreateGroupCertificateAuthorityOutput{})

	return CreateGroupCertificateAuthorityRequest{Request: req, Input: input, Copy: c.CreateGroupCertificateAuthorityRequest}
}

// CreateGroupCertificateAuthorityRequest is the request type for the
// CreateGroupCertificateAuthority API operation.
type CreateGroupCertificateAuthorityRequest struct {
	*aws.Request
	Input *CreateGroupCertificateAuthorityInput
	Copy  func(*CreateGroupCertificateAuthorityInput) CreateGroupCertificateAuthorityRequest
}

// Send marshals and sends the CreateGroupCertificateAuthority API request.
func (r CreateGroupCertificateAuthorityRequest) Send(ctx context.Context) (*CreateGroupCertificateAuthorityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateGroupCertificateAuthorityResponse{
		CreateGroupCertificateAuthorityOutput: r.Request.Data.(*CreateGroupCertificateAuthorityOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateGroupCertificateAuthorityResponse is the response type for the
// CreateGroupCertificateAuthority API operation.
type CreateGroupCertificateAuthorityResponse struct {
	*CreateGroupCertificateAuthorityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateGroupCertificateAuthority request.
func (r *CreateGroupCertificateAuthorityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
