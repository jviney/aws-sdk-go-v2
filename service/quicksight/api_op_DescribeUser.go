// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DescribeUserInput struct {
	_ struct{} `type:"structure"`

	// The ID for the AWS account that the user is in. Currently, you use the ID
	// for the AWS account that contains your Amazon QuickSight account.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The namespace. Currently, you should set this to default.
	//
	// Namespace is a required field
	Namespace *string `location:"uri" locationName:"Namespace" type:"string" required:"true"`

	// The name of the user that you want to describe.
	//
	// UserName is a required field
	UserName *string `location:"uri" locationName:"UserName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeUserInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.Namespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("Namespace"))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeUserInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Namespace != nil {
		v := *s.Namespace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Namespace", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserName != nil {
		v := *s.UserName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "UserName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeUserOutput struct {
	_ struct{} `type:"structure"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`

	// The user name.
	User *User `type:"structure"`
}

// String returns the string representation
func (s DescribeUserOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeUserOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.User != nil {
		v := s.User

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "User", v, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opDescribeUser = "DescribeUser"

// DescribeUserRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Returns information about a user, given the user name.
//
//    // Example sending a request using DescribeUserRequest.
//    req := client.DescribeUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DescribeUser
func (c *Client) DescribeUserRequest(input *DescribeUserInput) DescribeUserRequest {
	op := &aws.Operation{
		Name:       opDescribeUser,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/users/{UserName}",
	}

	if input == nil {
		input = &DescribeUserInput{}
	}

	req := c.newRequest(op, input, &DescribeUserOutput{})

	return DescribeUserRequest{Request: req, Input: input, Copy: c.DescribeUserRequest}
}

// DescribeUserRequest is the request type for the
// DescribeUser API operation.
type DescribeUserRequest struct {
	*aws.Request
	Input *DescribeUserInput
	Copy  func(*DescribeUserInput) DescribeUserRequest
}

// Send marshals and sends the DescribeUser API request.
func (r DescribeUserRequest) Send(ctx context.Context) (*DescribeUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeUserResponse{
		DescribeUserOutput: r.Request.Data.(*DescribeUserOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeUserResponse is the response type for the
// DescribeUser API operation.
type DescribeUserResponse struct {
	*DescribeUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeUser request.
func (r *DescribeUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
