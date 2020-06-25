// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ram

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListResourceSharePermissionsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The Amazon Resource Name (ARN) of the resource share.
	//
	// ResourceShareArn is a required field
	ResourceShareArn *string `locationName:"resourceShareArn" type:"string" required:"true"`
}

// String returns the string representation
func (s ListResourceSharePermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListResourceSharePermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListResourceSharePermissionsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ResourceShareArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceShareArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListResourceSharePermissionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceShareArn != nil {
		v := *s.ResourceShareArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceShareArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListResourceSharePermissionsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The permissions associated with the resource share.
	Permissions []ResourceSharePermissionSummary `locationName:"permissions" type:"list"`
}

// String returns the string representation
func (s ListResourceSharePermissionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListResourceSharePermissionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Permissions != nil {
		v := s.Permissions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "permissions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListResourceSharePermissions = "ListResourceSharePermissions"

// ListResourceSharePermissionsRequest returns a request value for making API operation for
// AWS Resource Access Manager.
//
// Lists the AWS RAM permissions that are associated with a resource share.
//
//    // Example sending a request using ListResourceSharePermissionsRequest.
//    req := client.ListResourceSharePermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/ListResourceSharePermissions
func (c *Client) ListResourceSharePermissionsRequest(input *ListResourceSharePermissionsInput) ListResourceSharePermissionsRequest {
	op := &aws.Operation{
		Name:       opListResourceSharePermissions,
		HTTPMethod: "POST",
		HTTPPath:   "/listresourcesharepermissions",
	}

	if input == nil {
		input = &ListResourceSharePermissionsInput{}
	}

	req := c.newRequest(op, input, &ListResourceSharePermissionsOutput{})

	return ListResourceSharePermissionsRequest{Request: req, Input: input, Copy: c.ListResourceSharePermissionsRequest}
}

// ListResourceSharePermissionsRequest is the request type for the
// ListResourceSharePermissions API operation.
type ListResourceSharePermissionsRequest struct {
	*aws.Request
	Input *ListResourceSharePermissionsInput
	Copy  func(*ListResourceSharePermissionsInput) ListResourceSharePermissionsRequest
}

// Send marshals and sends the ListResourceSharePermissions API request.
func (r ListResourceSharePermissionsRequest) Send(ctx context.Context) (*ListResourceSharePermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListResourceSharePermissionsResponse{
		ListResourceSharePermissionsOutput: r.Request.Data.(*ListResourceSharePermissionsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListResourceSharePermissionsResponse is the response type for the
// ListResourceSharePermissions API operation.
type ListResourceSharePermissionsResponse struct {
	*ListResourceSharePermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListResourceSharePermissions request.
func (r *ListResourceSharePermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
