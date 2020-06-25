// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteMetricFilterInput struct {
	_ struct{} `type:"structure"`

	// The name of the metric filter.
	//
	// FilterName is a required field
	FilterName *string `locationName:"filterName" min:"1" type:"string" required:"true"`

	// The name of the log group.
	//
	// LogGroupName is a required field
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMetricFilterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMetricFilterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteMetricFilterInput"}

	if s.FilterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FilterName"))
	}
	if s.FilterName != nil && len(*s.FilterName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FilterName", 1))
	}

	if s.LogGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogGroupName"))
	}
	if s.LogGroupName != nil && len(*s.LogGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteMetricFilterOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteMetricFilterOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteMetricFilter = "DeleteMetricFilter"

// DeleteMetricFilterRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Deletes the specified metric filter.
//
//    // Example sending a request using DeleteMetricFilterRequest.
//    req := client.DeleteMetricFilterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DeleteMetricFilter
func (c *Client) DeleteMetricFilterRequest(input *DeleteMetricFilterInput) DeleteMetricFilterRequest {
	op := &aws.Operation{
		Name:       opDeleteMetricFilter,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteMetricFilterInput{}
	}

	req := c.newRequest(op, input, &DeleteMetricFilterOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteMetricFilterRequest{Request: req, Input: input, Copy: c.DeleteMetricFilterRequest}
}

// DeleteMetricFilterRequest is the request type for the
// DeleteMetricFilter API operation.
type DeleteMetricFilterRequest struct {
	*aws.Request
	Input *DeleteMetricFilterInput
	Copy  func(*DeleteMetricFilterInput) DeleteMetricFilterRequest
}

// Send marshals and sends the DeleteMetricFilter API request.
func (r DeleteMetricFilterRequest) Send(ctx context.Context) (*DeleteMetricFilterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteMetricFilterResponse{
		DeleteMetricFilterOutput: r.Request.Data.(*DeleteMetricFilterOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteMetricFilterResponse is the response type for the
// DeleteMetricFilter API operation.
type DeleteMetricFilterResponse struct {
	*DeleteMetricFilterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteMetricFilter request.
func (r *DeleteMetricFilterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
