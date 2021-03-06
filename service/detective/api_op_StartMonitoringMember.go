// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package detective

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/restjson"
)

type StartMonitoringMemberInput struct {
	_ struct{} `type:"structure"`

	// The account ID of the member account to try to enable.
	//
	// The account must be an invited member account with a status of ACCEPTED_BUT_DISABLED.
	//
	// AccountId is a required field
	AccountId *string `min:"12" type:"string" required:"true"`

	// The ARN of the behavior graph.
	//
	// GraphArn is a required field
	GraphArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StartMonitoringMemberInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartMonitoringMemberInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartMonitoringMemberInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}
	if s.AccountId != nil && len(*s.AccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountId", 12))
	}

	if s.GraphArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("GraphArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartMonitoringMemberInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GraphArn != nil {
		v := *s.GraphArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GraphArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type StartMonitoringMemberOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartMonitoringMemberOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartMonitoringMemberOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opStartMonitoringMember = "StartMonitoringMember"

// StartMonitoringMemberRequest returns a request value for making API operation for
// Amazon Detective.
//
// Sends a request to enable data ingest for a member account that has a status
// of ACCEPTED_BUT_DISABLED.
//
// For valid member accounts, the status is updated as follows.
//
//    * If Detective enabled the member account, then the new status is ENABLED.
//
//    * If Detective cannot enable the member account, the status remains ACCEPTED_BUT_DISABLED.
//
//    // Example sending a request using StartMonitoringMemberRequest.
//    req := client.StartMonitoringMemberRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/detective-2018-10-26/StartMonitoringMember
func (c *Client) StartMonitoringMemberRequest(input *StartMonitoringMemberInput) StartMonitoringMemberRequest {
	op := &aws.Operation{
		Name:       opStartMonitoringMember,
		HTTPMethod: "POST",
		HTTPPath:   "/graph/member/monitoringstate",
	}

	if input == nil {
		input = &StartMonitoringMemberInput{}
	}

	req := c.newRequest(op, input, &StartMonitoringMemberOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return StartMonitoringMemberRequest{Request: req, Input: input, Copy: c.StartMonitoringMemberRequest}
}

// StartMonitoringMemberRequest is the request type for the
// StartMonitoringMember API operation.
type StartMonitoringMemberRequest struct {
	*aws.Request
	Input *StartMonitoringMemberInput
	Copy  func(*StartMonitoringMemberInput) StartMonitoringMemberRequest
}

// Send marshals and sends the StartMonitoringMember API request.
func (r StartMonitoringMemberRequest) Send(ctx context.Context) (*StartMonitoringMemberResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartMonitoringMemberResponse{
		StartMonitoringMemberOutput: r.Request.Data.(*StartMonitoringMemberOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartMonitoringMemberResponse is the response type for the
// StartMonitoringMember API operation.
type StartMonitoringMemberResponse struct {
	*StartMonitoringMemberOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartMonitoringMember request.
func (r *StartMonitoringMemberResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
