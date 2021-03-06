// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// A complex type that contains information about the request to remove authorization
// to associate a VPC that was created by one AWS account with a hosted zone
// that was created with a different AWS account.
type DeleteVPCAssociationAuthorizationInput struct {
	_ struct{} `locationName:"DeleteVPCAssociationAuthorizationRequest" type:"structure" xmlURI:"https://route53.amazonaws.com/doc/2013-04-01/"`

	// When removing authorization to associate a VPC that was created by one AWS
	// account with a hosted zone that was created with a different AWS account,
	// the ID of the hosted zone.
	//
	// HostedZoneId is a required field
	HostedZoneId *string `location:"uri" locationName:"Id" type:"string" required:"true"`

	// When removing authorization to associate a VPC that was created by one AWS
	// account with a hosted zone that was created with a different AWS account,
	// a complex type that includes the ID and region of the VPC.
	//
	// VPC is a required field
	VPC *VPC `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteVPCAssociationAuthorizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVPCAssociationAuthorizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVPCAssociationAuthorizationInput"}

	if s.HostedZoneId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HostedZoneId"))
	}

	if s.VPC == nil {
		invalidParams.Add(aws.NewErrParamRequired("VPC"))
	}
	if s.VPC != nil {
		if err := s.VPC.Validate(); err != nil {
			invalidParams.AddNested("VPC", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVPCAssociationAuthorizationInput) MarshalFields(e protocol.FieldEncoder) error {

	e.SetFields(protocol.BodyTarget, "DeleteVPCAssociationAuthorizationRequest", protocol.FieldMarshalerFunc(func(e protocol.FieldEncoder) error {
		if s.VPC != nil {
			v := s.VPC

			metadata := protocol.Metadata{}
			e.SetFields(protocol.BodyTarget, "VPC", v, metadata)
		}
		return nil
	}), protocol.Metadata{XMLNamespaceURI: "https://route53.amazonaws.com/doc/2013-04-01/"})
	if s.HostedZoneId != nil {
		v := *s.HostedZoneId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	return nil
}

// Empty response for the request.
type DeleteVPCAssociationAuthorizationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteVPCAssociationAuthorizationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVPCAssociationAuthorizationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteVPCAssociationAuthorization = "DeleteVPCAssociationAuthorization"

// DeleteVPCAssociationAuthorizationRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Removes authorization to submit an AssociateVPCWithHostedZone request to
// associate a specified VPC with a hosted zone that was created by a different
// account. You must use the account that created the hosted zone to submit
// a DeleteVPCAssociationAuthorization request.
//
// Sending this request only prevents the AWS account that created the VPC from
// associating the VPC with the Amazon Route 53 hosted zone in the future. If
// the VPC is already associated with the hosted zone, DeleteVPCAssociationAuthorization
// won't disassociate the VPC from the hosted zone. If you want to delete an
// existing association, use DisassociateVPCFromHostedZone.
//
//    // Example sending a request using DeleteVPCAssociationAuthorizationRequest.
//    req := client.DeleteVPCAssociationAuthorizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/DeleteVPCAssociationAuthorization
func (c *Client) DeleteVPCAssociationAuthorizationRequest(input *DeleteVPCAssociationAuthorizationInput) DeleteVPCAssociationAuthorizationRequest {
	op := &aws.Operation{
		Name:       opDeleteVPCAssociationAuthorization,
		HTTPMethod: "POST",
		HTTPPath:   "/2013-04-01/hostedzone/{Id}/deauthorizevpcassociation",
	}

	if input == nil {
		input = &DeleteVPCAssociationAuthorizationInput{}
	}

	req := c.newRequest(op, input, &DeleteVPCAssociationAuthorizationOutput{})

	return DeleteVPCAssociationAuthorizationRequest{Request: req, Input: input, Copy: c.DeleteVPCAssociationAuthorizationRequest}
}

// DeleteVPCAssociationAuthorizationRequest is the request type for the
// DeleteVPCAssociationAuthorization API operation.
type DeleteVPCAssociationAuthorizationRequest struct {
	*aws.Request
	Input *DeleteVPCAssociationAuthorizationInput
	Copy  func(*DeleteVPCAssociationAuthorizationInput) DeleteVPCAssociationAuthorizationRequest
}

// Send marshals and sends the DeleteVPCAssociationAuthorization API request.
func (r DeleteVPCAssociationAuthorizationRequest) Send(ctx context.Context) (*DeleteVPCAssociationAuthorizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVPCAssociationAuthorizationResponse{
		DeleteVPCAssociationAuthorizationOutput: r.Request.Data.(*DeleteVPCAssociationAuthorizationOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVPCAssociationAuthorizationResponse is the response type for the
// DeleteVPCAssociationAuthorization API operation.
type DeleteVPCAssociationAuthorizationResponse struct {
	*DeleteVPCAssociationAuthorizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVPCAssociationAuthorization request.
func (r *DeleteVPCAssociationAuthorizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
