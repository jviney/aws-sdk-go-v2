// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type AssociateTransitGatewayRouteTableInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the attachment.
	//
	// TransitGatewayAttachmentId is a required field
	TransitGatewayAttachmentId *string `type:"string" required:"true"`

	// The ID of the transit gateway route table.
	//
	// TransitGatewayRouteTableId is a required field
	TransitGatewayRouteTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateTransitGatewayRouteTableInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateTransitGatewayRouteTableInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateTransitGatewayRouteTableInput"}

	if s.TransitGatewayAttachmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayAttachmentId"))
	}

	if s.TransitGatewayRouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayRouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateTransitGatewayRouteTableOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the association.
	Association *TransitGatewayAssociation `locationName:"association" type:"structure"`
}

// String returns the string representation
func (s AssociateTransitGatewayRouteTableOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateTransitGatewayRouteTable = "AssociateTransitGatewayRouteTable"

// AssociateTransitGatewayRouteTableRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Associates the specified attachment with the specified transit gateway route
// table. You can associate only one route table with an attachment.
//
//    // Example sending a request using AssociateTransitGatewayRouteTableRequest.
//    req := client.AssociateTransitGatewayRouteTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateTransitGatewayRouteTable
func (c *Client) AssociateTransitGatewayRouteTableRequest(input *AssociateTransitGatewayRouteTableInput) AssociateTransitGatewayRouteTableRequest {
	op := &aws.Operation{
		Name:       opAssociateTransitGatewayRouteTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateTransitGatewayRouteTableInput{}
	}

	req := c.newRequest(op, input, &AssociateTransitGatewayRouteTableOutput{})

	return AssociateTransitGatewayRouteTableRequest{Request: req, Input: input, Copy: c.AssociateTransitGatewayRouteTableRequest}
}

// AssociateTransitGatewayRouteTableRequest is the request type for the
// AssociateTransitGatewayRouteTable API operation.
type AssociateTransitGatewayRouteTableRequest struct {
	*aws.Request
	Input *AssociateTransitGatewayRouteTableInput
	Copy  func(*AssociateTransitGatewayRouteTableInput) AssociateTransitGatewayRouteTableRequest
}

// Send marshals and sends the AssociateTransitGatewayRouteTable API request.
func (r AssociateTransitGatewayRouteTableRequest) Send(ctx context.Context) (*AssociateTransitGatewayRouteTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateTransitGatewayRouteTableResponse{
		AssociateTransitGatewayRouteTableOutput: r.Request.Data.(*AssociateTransitGatewayRouteTableOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateTransitGatewayRouteTableResponse is the response type for the
// AssociateTransitGatewayRouteTable API operation.
type AssociateTransitGatewayRouteTableResponse struct {
	*AssociateTransitGatewayRouteTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateTransitGatewayRouteTable request.
func (r *AssociateTransitGatewayRouteTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
