// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type UntagLogGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the log group.
	//
	// LogGroupName is a required field
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string" required:"true"`

	// The tag keys. The corresponding tags are removed from the log group.
	//
	// Tags is a required field
	Tags []string `locationName:"tags" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s UntagLogGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UntagLogGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UntagLogGroupInput"}

	if s.LogGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogGroupName"))
	}
	if s.LogGroupName != nil && len(*s.LogGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogGroupName", 1))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UntagLogGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UntagLogGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opUntagLogGroup = "UntagLogGroup"

// UntagLogGroupRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Removes the specified tags from the specified log group.
//
// To list the tags for a log group, use ListTagsLogGroup (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ListTagsLogGroup.html).
// To add tags, use TagLogGroup (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_TagLogGroup.html).
//
//    // Example sending a request using UntagLogGroupRequest.
//    req := client.UntagLogGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/UntagLogGroup
func (c *Client) UntagLogGroupRequest(input *UntagLogGroupInput) UntagLogGroupRequest {
	op := &aws.Operation{
		Name:       opUntagLogGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UntagLogGroupInput{}
	}

	req := c.newRequest(op, input, &UntagLogGroupOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return UntagLogGroupRequest{Request: req, Input: input, Copy: c.UntagLogGroupRequest}
}

// UntagLogGroupRequest is the request type for the
// UntagLogGroup API operation.
type UntagLogGroupRequest struct {
	*aws.Request
	Input *UntagLogGroupInput
	Copy  func(*UntagLogGroupInput) UntagLogGroupRequest
}

// Send marshals and sends the UntagLogGroup API request.
func (r UntagLogGroupRequest) Send(ctx context.Context) (*UntagLogGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UntagLogGroupResponse{
		UntagLogGroupOutput: r.Request.Data.(*UntagLogGroupOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UntagLogGroupResponse is the response type for the
// UntagLogGroup API operation.
type UntagLogGroupResponse struct {
	*UntagLogGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UntagLogGroup request.
func (r *UntagLogGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
