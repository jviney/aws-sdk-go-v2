// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type EnableSecurityHubInput struct {
	_ struct{} `type:"structure"`

	// Whether to enable the security standards that Security Hub has designated
	// as automatically enabled. If you do not provide a value for EnableDefaultStandards,
	// it is set to true. To not enable the automatically enabled standards, set
	// EnableDefaultStandards to false.
	EnableDefaultStandards *bool `type:"boolean"`

	// The tags to add to the hub resource when you enable Security Hub.
	Tags map[string]string `min:"1" type:"map"`
}

// String returns the string representation
func (s EnableSecurityHubInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableSecurityHubInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EnableSecurityHubInput"}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s EnableSecurityHubInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.EnableDefaultStandards != nil {
		v := *s.EnableDefaultStandards

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EnableDefaultStandards", protocol.BoolValue(v), metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

type EnableSecurityHubOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EnableSecurityHubOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s EnableSecurityHubOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opEnableSecurityHub = "EnableSecurityHub"

// EnableSecurityHubRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Enables Security Hub for your account in the current Region or the Region
// you specify in the request.
//
// When you enable Security Hub, you grant to Security Hub the permissions necessary
// to gather findings from other services that are integrated with Security
// Hub.
//
// When you use the EnableSecurityHub operation to enable Security Hub, you
// also automatically enable the following standards.
//
//    * CIS AWS Foundations
//
//    * AWS Foundational Security Best Practices
//
// You do not enable the Payment Card Industry Data Security Standard (PCI DSS)
// standard.
//
// To not enable the automatically enabled standards, set EnableDefaultStandards
// to false.
//
// After you enable Security Hub, to enable a standard, use the BatchEnableStandards
// operation. To disable a standard, use the BatchDisableStandards operation.
//
// To learn more, see Setting Up AWS Security Hub (https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-settingup.html)
// in the AWS Security Hub User Guide.
//
//    // Example sending a request using EnableSecurityHubRequest.
//    req := client.EnableSecurityHubRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/EnableSecurityHub
func (c *Client) EnableSecurityHubRequest(input *EnableSecurityHubInput) EnableSecurityHubRequest {
	op := &aws.Operation{
		Name:       opEnableSecurityHub,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts",
	}

	if input == nil {
		input = &EnableSecurityHubInput{}
	}

	req := c.newRequest(op, input, &EnableSecurityHubOutput{})

	return EnableSecurityHubRequest{Request: req, Input: input, Copy: c.EnableSecurityHubRequest}
}

// EnableSecurityHubRequest is the request type for the
// EnableSecurityHub API operation.
type EnableSecurityHubRequest struct {
	*aws.Request
	Input *EnableSecurityHubInput
	Copy  func(*EnableSecurityHubInput) EnableSecurityHubRequest
}

// Send marshals and sends the EnableSecurityHub API request.
func (r EnableSecurityHubRequest) Send(ctx context.Context) (*EnableSecurityHubResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableSecurityHubResponse{
		EnableSecurityHubOutput: r.Request.Data.(*EnableSecurityHubOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableSecurityHubResponse is the response type for the
// EnableSecurityHub API operation.
type EnableSecurityHubResponse struct {
	*EnableSecurityHubOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableSecurityHub request.
func (r *EnableSecurityHubResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
