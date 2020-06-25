// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package networkmanager

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DeleteGlobalNetworkInput struct {
	_ struct{} `type:"structure"`

	// The ID of the global network.
	//
	// GlobalNetworkId is a required field
	GlobalNetworkId *string `location:"uri" locationName:"globalNetworkId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteGlobalNetworkInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteGlobalNetworkInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteGlobalNetworkInput"}

	if s.GlobalNetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GlobalNetworkId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteGlobalNetworkInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GlobalNetworkId != nil {
		v := *s.GlobalNetworkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "globalNetworkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteGlobalNetworkOutput struct {
	_ struct{} `type:"structure"`

	// Information about the global network.
	GlobalNetwork *GlobalNetwork `type:"structure"`
}

// String returns the string representation
func (s DeleteGlobalNetworkOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteGlobalNetworkOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.GlobalNetwork != nil {
		v := s.GlobalNetwork

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "GlobalNetwork", v, metadata)
	}
	return nil
}

const opDeleteGlobalNetwork = "DeleteGlobalNetwork"

// DeleteGlobalNetworkRequest returns a request value for making API operation for
// AWS Network Manager.
//
// Deletes an existing global network. You must first delete all global network
// objects (devices, links, and sites) and deregister all transit gateways.
//
//    // Example sending a request using DeleteGlobalNetworkRequest.
//    req := client.DeleteGlobalNetworkRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/networkmanager-2019-07-05/DeleteGlobalNetwork
func (c *Client) DeleteGlobalNetworkRequest(input *DeleteGlobalNetworkInput) DeleteGlobalNetworkRequest {
	op := &aws.Operation{
		Name:       opDeleteGlobalNetwork,
		HTTPMethod: "DELETE",
		HTTPPath:   "/global-networks/{globalNetworkId}",
	}

	if input == nil {
		input = &DeleteGlobalNetworkInput{}
	}

	req := c.newRequest(op, input, &DeleteGlobalNetworkOutput{})

	return DeleteGlobalNetworkRequest{Request: req, Input: input, Copy: c.DeleteGlobalNetworkRequest}
}

// DeleteGlobalNetworkRequest is the request type for the
// DeleteGlobalNetwork API operation.
type DeleteGlobalNetworkRequest struct {
	*aws.Request
	Input *DeleteGlobalNetworkInput
	Copy  func(*DeleteGlobalNetworkInput) DeleteGlobalNetworkRequest
}

// Send marshals and sends the DeleteGlobalNetwork API request.
func (r DeleteGlobalNetworkRequest) Send(ctx context.Context) (*DeleteGlobalNetworkResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteGlobalNetworkResponse{
		DeleteGlobalNetworkOutput: r.Request.Data.(*DeleteGlobalNetworkOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteGlobalNetworkResponse is the response type for the
// DeleteGlobalNetwork API operation.
type DeleteGlobalNetworkResponse struct {
	*DeleteGlobalNetworkOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteGlobalNetwork request.
func (r *DeleteGlobalNetworkResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
