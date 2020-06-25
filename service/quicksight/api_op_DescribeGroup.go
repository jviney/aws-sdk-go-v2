// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DescribeGroupInput struct {
	_ struct{} `type:"structure"`

	// The ID for the AWS account that the group is in. Currently, you use the ID
	// for the AWS account that contains your Amazon QuickSight account.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The name of the group that you want to describe.
	//
	// GroupName is a required field
	GroupName *string `location:"uri" locationName:"GroupName" min:"1" type:"string" required:"true"`

	// The namespace. Currently, you should set this to default.
	//
	// Namespace is a required field
	Namespace *string `location:"uri" locationName:"Namespace" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeGroupInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.GroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupName"))
	}
	if s.GroupName != nil && len(*s.GroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupName", 1))
	}

	if s.Namespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("Namespace"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GroupName != nil {
		v := *s.GroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Namespace != nil {
		v := *s.Namespace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Namespace", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeGroupOutput struct {
	_ struct{} `type:"structure"`

	// The name of the group.
	Group *Group `type:"structure"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s DescribeGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Group != nil {
		v := s.Group

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Group", v, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opDescribeGroup = "DescribeGroup"

// DescribeGroupRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Returns an Amazon QuickSight group's description and Amazon Resource Name
// (ARN).
//
//    // Example sending a request using DescribeGroupRequest.
//    req := client.DescribeGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DescribeGroup
func (c *Client) DescribeGroupRequest(input *DescribeGroupInput) DescribeGroupRequest {
	op := &aws.Operation{
		Name:       opDescribeGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/groups/{GroupName}",
	}

	if input == nil {
		input = &DescribeGroupInput{}
	}

	req := c.newRequest(op, input, &DescribeGroupOutput{})

	return DescribeGroupRequest{Request: req, Input: input, Copy: c.DescribeGroupRequest}
}

// DescribeGroupRequest is the request type for the
// DescribeGroup API operation.
type DescribeGroupRequest struct {
	*aws.Request
	Input *DescribeGroupInput
	Copy  func(*DescribeGroupInput) DescribeGroupRequest
}

// Send marshals and sends the DescribeGroup API request.
func (r DescribeGroupRequest) Send(ctx context.Context) (*DescribeGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeGroupResponse{
		DescribeGroupOutput: r.Request.Data.(*DescribeGroupOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeGroupResponse is the response type for the
// DescribeGroup API operation.
type DescribeGroupResponse struct {
	*DescribeGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeGroup request.
func (r *DescribeGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
