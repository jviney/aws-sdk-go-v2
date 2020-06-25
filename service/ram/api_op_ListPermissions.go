// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ram

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListPermissionsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Specifies the resource type for which to list permissions. For example, to
	// list only permissions that apply to EC2 subnets, specify ec2:Subnet.
	ResourceType *string `locationName:"resourceType" type:"string"`
}

// String returns the string representation
func (s ListPermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPermissionsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPermissionsInput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.ResourceType != nil {
		v := *s.ResourceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListPermissionsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the permissions.
	Permissions []ResourceSharePermissionSummary `locationName:"permissions" type:"list"`
}

// String returns the string representation
func (s ListPermissionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPermissionsOutput) MarshalFields(e protocol.FieldEncoder) error {
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

const opListPermissions = "ListPermissions"

// ListPermissionsRequest returns a request value for making API operation for
// AWS Resource Access Manager.
//
// Lists the AWS RAM permissions.
//
//    // Example sending a request using ListPermissionsRequest.
//    req := client.ListPermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/ListPermissions
func (c *Client) ListPermissionsRequest(input *ListPermissionsInput) ListPermissionsRequest {
	op := &aws.Operation{
		Name:       opListPermissions,
		HTTPMethod: "POST",
		HTTPPath:   "/listpermissions",
	}

	if input == nil {
		input = &ListPermissionsInput{}
	}

	req := c.newRequest(op, input, &ListPermissionsOutput{})

	return ListPermissionsRequest{Request: req, Input: input, Copy: c.ListPermissionsRequest}
}

// ListPermissionsRequest is the request type for the
// ListPermissions API operation.
type ListPermissionsRequest struct {
	*aws.Request
	Input *ListPermissionsInput
	Copy  func(*ListPermissionsInput) ListPermissionsRequest
}

// Send marshals and sends the ListPermissions API request.
func (r ListPermissionsRequest) Send(ctx context.Context) (*ListPermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPermissionsResponse{
		ListPermissionsOutput: r.Request.Data.(*ListPermissionsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListPermissionsResponse is the response type for the
// ListPermissions API operation.
type ListPermissionsResponse struct {
	*ListPermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPermissions request.
func (r *ListPermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
