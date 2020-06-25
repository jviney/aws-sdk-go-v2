// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteFunctionEventInvokeConfigInput struct {
	_ struct{} `type:"structure"`

	// The name of the Lambda function, version, or alias.
	//
	// Name formats
	//
	//    * Function name - my-function (name-only), my-function:v1 (with alias).
	//
	//    * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//    * Partial ARN - 123456789012:function:my-function.
	//
	// You can append a version number or alias to any of the formats. The length
	// constraint applies only to the full ARN. If you specify only the function
	// name, it is limited to 64 characters in length.
	//
	// FunctionName is a required field
	FunctionName *string `location:"uri" locationName:"FunctionName" min:"1" type:"string" required:"true"`

	// A version number or alias name.
	Qualifier *string `location:"querystring" locationName:"Qualifier" min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteFunctionEventInvokeConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFunctionEventInvokeConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFunctionEventInvokeConfigInput"}

	if s.FunctionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionName"))
	}
	if s.FunctionName != nil && len(*s.FunctionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionName", 1))
	}
	if s.Qualifier != nil && len(*s.Qualifier) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Qualifier", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteFunctionEventInvokeConfigInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FunctionName != nil {
		v := *s.FunctionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FunctionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Qualifier != nil {
		v := *s.Qualifier

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Qualifier", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteFunctionEventInvokeConfigOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteFunctionEventInvokeConfigOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteFunctionEventInvokeConfigOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteFunctionEventInvokeConfig = "DeleteFunctionEventInvokeConfig"

// DeleteFunctionEventInvokeConfigRequest returns a request value for making API operation for
// AWS Lambda.
//
// Deletes the configuration for asynchronous invocation for a function, version,
// or alias.
//
// To configure options for asynchronous invocation, use PutFunctionEventInvokeConfig.
//
//    // Example sending a request using DeleteFunctionEventInvokeConfigRequest.
//    req := client.DeleteFunctionEventInvokeConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/DeleteFunctionEventInvokeConfig
func (c *Client) DeleteFunctionEventInvokeConfigRequest(input *DeleteFunctionEventInvokeConfigInput) DeleteFunctionEventInvokeConfigRequest {
	op := &aws.Operation{
		Name:       opDeleteFunctionEventInvokeConfig,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2019-09-25/functions/{FunctionName}/event-invoke-config",
	}

	if input == nil {
		input = &DeleteFunctionEventInvokeConfigInput{}
	}

	req := c.newRequest(op, input, &DeleteFunctionEventInvokeConfigOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteFunctionEventInvokeConfigRequest{Request: req, Input: input, Copy: c.DeleteFunctionEventInvokeConfigRequest}
}

// DeleteFunctionEventInvokeConfigRequest is the request type for the
// DeleteFunctionEventInvokeConfig API operation.
type DeleteFunctionEventInvokeConfigRequest struct {
	*aws.Request
	Input *DeleteFunctionEventInvokeConfigInput
	Copy  func(*DeleteFunctionEventInvokeConfigInput) DeleteFunctionEventInvokeConfigRequest
}

// Send marshals and sends the DeleteFunctionEventInvokeConfig API request.
func (r DeleteFunctionEventInvokeConfigRequest) Send(ctx context.Context) (*DeleteFunctionEventInvokeConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteFunctionEventInvokeConfigResponse{
		DeleteFunctionEventInvokeConfigOutput: r.Request.Data.(*DeleteFunctionEventInvokeConfigOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteFunctionEventInvokeConfigResponse is the response type for the
// DeleteFunctionEventInvokeConfig API operation.
type DeleteFunctionEventInvokeConfigResponse struct {
	*DeleteFunctionEventInvokeConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteFunctionEventInvokeConfig request.
func (r *DeleteFunctionEventInvokeConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
