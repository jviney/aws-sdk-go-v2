// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteTrafficMirrorSessionInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the Traffic Mirror session.
	//
	// TrafficMirrorSessionId is a required field
	TrafficMirrorSessionId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTrafficMirrorSessionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTrafficMirrorSessionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTrafficMirrorSessionInput"}

	if s.TrafficMirrorSessionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrafficMirrorSessionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteTrafficMirrorSessionOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the deleted Traffic Mirror session.
	TrafficMirrorSessionId *string `locationName:"trafficMirrorSessionId" type:"string"`
}

// String returns the string representation
func (s DeleteTrafficMirrorSessionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTrafficMirrorSession = "DeleteTrafficMirrorSession"

// DeleteTrafficMirrorSessionRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified Traffic Mirror session.
//
//    // Example sending a request using DeleteTrafficMirrorSessionRequest.
//    req := client.DeleteTrafficMirrorSessionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTrafficMirrorSession
func (c *Client) DeleteTrafficMirrorSessionRequest(input *DeleteTrafficMirrorSessionInput) DeleteTrafficMirrorSessionRequest {
	op := &aws.Operation{
		Name:       opDeleteTrafficMirrorSession,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTrafficMirrorSessionInput{}
	}

	req := c.newRequest(op, input, &DeleteTrafficMirrorSessionOutput{})

	return DeleteTrafficMirrorSessionRequest{Request: req, Input: input, Copy: c.DeleteTrafficMirrorSessionRequest}
}

// DeleteTrafficMirrorSessionRequest is the request type for the
// DeleteTrafficMirrorSession API operation.
type DeleteTrafficMirrorSessionRequest struct {
	*aws.Request
	Input *DeleteTrafficMirrorSessionInput
	Copy  func(*DeleteTrafficMirrorSessionInput) DeleteTrafficMirrorSessionRequest
}

// Send marshals and sends the DeleteTrafficMirrorSession API request.
func (r DeleteTrafficMirrorSessionRequest) Send(ctx context.Context) (*DeleteTrafficMirrorSessionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTrafficMirrorSessionResponse{
		DeleteTrafficMirrorSessionOutput: r.Request.Data.(*DeleteTrafficMirrorSessionOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTrafficMirrorSessionResponse is the response type for the
// DeleteTrafficMirrorSession API operation.
type DeleteTrafficMirrorSessionResponse struct {
	*DeleteTrafficMirrorSessionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTrafficMirrorSession request.
func (r *DeleteTrafficMirrorSessionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
