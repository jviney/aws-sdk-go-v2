// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DeleteTypedLinkFacetInput struct {
	_ struct{} `type:"structure"`

	// The unique name of the typed link facet.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The Amazon Resource Name (ARN) that is associated with the schema. For more
	// information, see arns.
	//
	// SchemaArn is a required field
	SchemaArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTypedLinkFacetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTypedLinkFacetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTypedLinkFacetInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.SchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteTypedLinkFacetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteTypedLinkFacetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteTypedLinkFacetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteTypedLinkFacetOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteTypedLinkFacet = "DeleteTypedLinkFacet"

// DeleteTypedLinkFacetRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Deletes a TypedLinkFacet. For more information, see Typed Links (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
//
//    // Example sending a request using DeleteTypedLinkFacetRequest.
//    req := client.DeleteTypedLinkFacetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/DeleteTypedLinkFacet
func (c *Client) DeleteTypedLinkFacetRequest(input *DeleteTypedLinkFacetInput) DeleteTypedLinkFacetRequest {
	op := &aws.Operation{
		Name:       opDeleteTypedLinkFacet,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/typedlink/facet/delete",
	}

	if input == nil {
		input = &DeleteTypedLinkFacetInput{}
	}

	req := c.newRequest(op, input, &DeleteTypedLinkFacetOutput{})

	return DeleteTypedLinkFacetRequest{Request: req, Input: input, Copy: c.DeleteTypedLinkFacetRequest}
}

// DeleteTypedLinkFacetRequest is the request type for the
// DeleteTypedLinkFacet API operation.
type DeleteTypedLinkFacetRequest struct {
	*aws.Request
	Input *DeleteTypedLinkFacetInput
	Copy  func(*DeleteTypedLinkFacetInput) DeleteTypedLinkFacetRequest
}

// Send marshals and sends the DeleteTypedLinkFacet API request.
func (r DeleteTypedLinkFacetRequest) Send(ctx context.Context) (*DeleteTypedLinkFacetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTypedLinkFacetResponse{
		DeleteTypedLinkFacetOutput: r.Request.Data.(*DeleteTypedLinkFacetOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTypedLinkFacetResponse is the response type for the
// DeleteTypedLinkFacet API operation.
type DeleteTypedLinkFacetResponse struct {
	*DeleteTypedLinkFacetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTypedLinkFacet request.
func (r *DeleteTypedLinkFacetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
