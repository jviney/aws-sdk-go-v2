// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ListMembersInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	OnlyAssociated *string `location:"querystring" locationName:"onlyAssociated" type:"string"`
}

// String returns the string representation
func (s ListMembersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListMembersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListMembersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListMembersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

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
	if s.OnlyAssociated != nil {
		v := *s.OnlyAssociated

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "onlyAssociated", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Provides information about the accounts that are associated with an Amazon
// Macie master account.
type ListMembersOutput struct {
	_ struct{} `type:"structure"`

	Members []Member `locationName:"members" type:"list"`

	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListMembersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListMembersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Members != nil {
		v := s.Members

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "members", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListMembers = "ListMembers"

// ListMembersRequest returns a request value for making API operation for
// Amazon Macie 2.
//
// Retrieves information about the accounts that are associated with an Amazon
// Macie master account.
//
//    // Example sending a request using ListMembersRequest.
//    req := client.ListMembersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie2-2020-01-01/ListMembers
func (c *Client) ListMembersRequest(input *ListMembersInput) ListMembersRequest {
	op := &aws.Operation{
		Name:       opListMembers,
		HTTPMethod: "GET",
		HTTPPath:   "/members",
	}

	if input == nil {
		input = &ListMembersInput{}
	}

	req := c.newRequest(op, input, &ListMembersOutput{})

	return ListMembersRequest{Request: req, Input: input, Copy: c.ListMembersRequest}
}

// ListMembersRequest is the request type for the
// ListMembers API operation.
type ListMembersRequest struct {
	*aws.Request
	Input *ListMembersInput
	Copy  func(*ListMembersInput) ListMembersRequest
}

// Send marshals and sends the ListMembers API request.
func (r ListMembersRequest) Send(ctx context.Context) (*ListMembersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListMembersResponse{
		ListMembersOutput: r.Request.Data.(*ListMembersOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListMembersResponse is the response type for the
// ListMembers API operation.
type ListMembersResponse struct {
	*ListMembersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListMembers request.
func (r *ListMembersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
