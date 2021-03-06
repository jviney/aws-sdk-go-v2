// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type CreateTypedLinkFacetInput struct {
	_ struct{} `type:"structure"`

	// Facet structure that is associated with the typed link facet.
	//
	// Facet is a required field
	Facet *TypedLinkFacet `type:"structure" required:"true"`

	// The Amazon Resource Name (ARN) that is associated with the schema. For more
	// information, see arns.
	//
	// SchemaArn is a required field
	SchemaArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateTypedLinkFacetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTypedLinkFacetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTypedLinkFacetInput"}

	if s.Facet == nil {
		invalidParams.Add(aws.NewErrParamRequired("Facet"))
	}

	if s.SchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaArn"))
	}
	if s.Facet != nil {
		if err := s.Facet.Validate(); err != nil {
			invalidParams.AddNested("Facet", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateTypedLinkFacetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Facet != nil {
		v := s.Facet

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Facet", v, metadata)
	}
	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateTypedLinkFacetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateTypedLinkFacetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateTypedLinkFacetOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCreateTypedLinkFacet = "CreateTypedLinkFacet"

// CreateTypedLinkFacetRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Creates a TypedLinkFacet. For more information, see Typed Links (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
//
//    // Example sending a request using CreateTypedLinkFacetRequest.
//    req := client.CreateTypedLinkFacetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/CreateTypedLinkFacet
func (c *Client) CreateTypedLinkFacetRequest(input *CreateTypedLinkFacetInput) CreateTypedLinkFacetRequest {
	op := &aws.Operation{
		Name:       opCreateTypedLinkFacet,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/typedlink/facet/create",
	}

	if input == nil {
		input = &CreateTypedLinkFacetInput{}
	}

	req := c.newRequest(op, input, &CreateTypedLinkFacetOutput{})

	return CreateTypedLinkFacetRequest{Request: req, Input: input, Copy: c.CreateTypedLinkFacetRequest}
}

// CreateTypedLinkFacetRequest is the request type for the
// CreateTypedLinkFacet API operation.
type CreateTypedLinkFacetRequest struct {
	*aws.Request
	Input *CreateTypedLinkFacetInput
	Copy  func(*CreateTypedLinkFacetInput) CreateTypedLinkFacetRequest
}

// Send marshals and sends the CreateTypedLinkFacet API request.
func (r CreateTypedLinkFacetRequest) Send(ctx context.Context) (*CreateTypedLinkFacetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTypedLinkFacetResponse{
		CreateTypedLinkFacetOutput: r.Request.Data.(*CreateTypedLinkFacetOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTypedLinkFacetResponse is the response type for the
// CreateTypedLinkFacet API operation.
type CreateTypedLinkFacetResponse struct {
	*CreateTypedLinkFacetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTypedLinkFacet request.
func (r *CreateTypedLinkFacetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
