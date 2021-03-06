// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListProvisioningTemplateVersionsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return at one time.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// A token to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The name of the fleet provisioning template.
	//
	// TemplateName is a required field
	TemplateName *string `location:"uri" locationName:"templateName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListProvisioningTemplateVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListProvisioningTemplateVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListProvisioningTemplateVersionsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.TemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateName"))
	}
	if s.TemplateName != nil && len(*s.TemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListProvisioningTemplateVersionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.TemplateName != nil {
		v := *s.TemplateName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "templateName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListProvisioningTemplateVersionsOutput struct {
	_ struct{} `type:"structure"`

	// A token to retrieve the next set of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of fleet provisioning template versions.
	Versions []ProvisioningTemplateVersionSummary `locationName:"versions" type:"list"`
}

// String returns the string representation
func (s ListProvisioningTemplateVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListProvisioningTemplateVersionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Versions != nil {
		v := s.Versions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "versions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListProvisioningTemplateVersions = "ListProvisioningTemplateVersions"

// ListProvisioningTemplateVersionsRequest returns a request value for making API operation for
// AWS IoT.
//
// A list of fleet provisioning template versions.
//
//    // Example sending a request using ListProvisioningTemplateVersionsRequest.
//    req := client.ListProvisioningTemplateVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListProvisioningTemplateVersionsRequest(input *ListProvisioningTemplateVersionsInput) ListProvisioningTemplateVersionsRequest {
	op := &aws.Operation{
		Name:       opListProvisioningTemplateVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/provisioning-templates/{templateName}/versions",
	}

	if input == nil {
		input = &ListProvisioningTemplateVersionsInput{}
	}

	req := c.newRequest(op, input, &ListProvisioningTemplateVersionsOutput{})

	return ListProvisioningTemplateVersionsRequest{Request: req, Input: input, Copy: c.ListProvisioningTemplateVersionsRequest}
}

// ListProvisioningTemplateVersionsRequest is the request type for the
// ListProvisioningTemplateVersions API operation.
type ListProvisioningTemplateVersionsRequest struct {
	*aws.Request
	Input *ListProvisioningTemplateVersionsInput
	Copy  func(*ListProvisioningTemplateVersionsInput) ListProvisioningTemplateVersionsRequest
}

// Send marshals and sends the ListProvisioningTemplateVersions API request.
func (r ListProvisioningTemplateVersionsRequest) Send(ctx context.Context) (*ListProvisioningTemplateVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListProvisioningTemplateVersionsResponse{
		ListProvisioningTemplateVersionsOutput: r.Request.Data.(*ListProvisioningTemplateVersionsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListProvisioningTemplateVersionsResponse is the response type for the
// ListProvisioningTemplateVersions API operation.
type ListProvisioningTemplateVersionsResponse struct {
	*ListProvisioningTemplateVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListProvisioningTemplateVersions request.
func (r *ListProvisioningTemplateVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
