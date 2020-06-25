// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeScriptInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for a Realtime script to retrieve properties for. You
	// can use either the script ID or ARN value.
	//
	// ScriptId is a required field
	ScriptId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeScriptInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeScriptInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeScriptInput"}

	if s.ScriptId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScriptId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeScriptOutput struct {
	_ struct{} `type:"structure"`

	// A set of properties describing the requested script.
	Script *Script `type:"structure"`
}

// String returns the string representation
func (s DescribeScriptOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeScript = "DescribeScript"

// DescribeScriptRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves properties for a Realtime script.
//
// To request a script record, specify the script ID. If successful, an object
// containing the script properties is returned.
//
// Learn more
//
// Amazon GameLift Realtime Servers (https://docs.aws.amazon.com/gamelift/latest/developerguide/realtime-intro.html)
//
// Related operations
//
//    * CreateScript
//
//    * ListScripts
//
//    * DescribeScript
//
//    * UpdateScript
//
//    * DeleteScript
//
//    // Example sending a request using DescribeScriptRequest.
//    req := client.DescribeScriptRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeScript
func (c *Client) DescribeScriptRequest(input *DescribeScriptInput) DescribeScriptRequest {
	op := &aws.Operation{
		Name:       opDescribeScript,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeScriptInput{}
	}

	req := c.newRequest(op, input, &DescribeScriptOutput{})

	return DescribeScriptRequest{Request: req, Input: input, Copy: c.DescribeScriptRequest}
}

// DescribeScriptRequest is the request type for the
// DescribeScript API operation.
type DescribeScriptRequest struct {
	*aws.Request
	Input *DescribeScriptInput
	Copy  func(*DescribeScriptInput) DescribeScriptRequest
}

// Send marshals and sends the DescribeScript API request.
func (r DescribeScriptRequest) Send(ctx context.Context) (*DescribeScriptResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeScriptResponse{
		DescribeScriptOutput: r.Request.Data.(*DescribeScriptOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeScriptResponse is the response type for the
// DescribeScript API operation.
type DescribeScriptResponse struct {
	*DescribeScriptOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeScript request.
func (r *DescribeScriptResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
