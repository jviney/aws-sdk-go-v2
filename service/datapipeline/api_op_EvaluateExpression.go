// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datapipeline

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for EvaluateExpression.
type EvaluateExpressionInput struct {
	_ struct{} `type:"structure"`

	// The expression to evaluate.
	//
	// Expression is a required field
	Expression *string `locationName:"expression" type:"string" required:"true"`

	// The ID of the object.
	//
	// ObjectId is a required field
	ObjectId *string `locationName:"objectId" min:"1" type:"string" required:"true"`

	// The ID of the pipeline.
	//
	// PipelineId is a required field
	PipelineId *string `locationName:"pipelineId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s EvaluateExpressionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EvaluateExpressionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EvaluateExpressionInput"}

	if s.Expression == nil {
		invalidParams.Add(aws.NewErrParamRequired("Expression"))
	}

	if s.ObjectId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObjectId"))
	}
	if s.ObjectId != nil && len(*s.ObjectId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ObjectId", 1))
	}

	if s.PipelineId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineId"))
	}
	if s.PipelineId != nil && len(*s.PipelineId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PipelineId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of EvaluateExpression.
type EvaluateExpressionOutput struct {
	_ struct{} `type:"structure"`

	// The evaluated expression.
	//
	// EvaluatedExpression is a required field
	EvaluatedExpression *string `locationName:"evaluatedExpression" type:"string" required:"true"`
}

// String returns the string representation
func (s EvaluateExpressionOutput) String() string {
	return awsutil.Prettify(s)
}

const opEvaluateExpression = "EvaluateExpression"

// EvaluateExpressionRequest returns a request value for making API operation for
// AWS Data Pipeline.
//
// Task runners call EvaluateExpression to evaluate a string in the context
// of the specified object. For example, a task runner can evaluate SQL queries
// stored in Amazon S3.
//
//    // Example sending a request using EvaluateExpressionRequest.
//    req := client.EvaluateExpressionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/EvaluateExpression
func (c *Client) EvaluateExpressionRequest(input *EvaluateExpressionInput) EvaluateExpressionRequest {
	op := &aws.Operation{
		Name:       opEvaluateExpression,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EvaluateExpressionInput{}
	}

	req := c.newRequest(op, input, &EvaluateExpressionOutput{})

	return EvaluateExpressionRequest{Request: req, Input: input, Copy: c.EvaluateExpressionRequest}
}

// EvaluateExpressionRequest is the request type for the
// EvaluateExpression API operation.
type EvaluateExpressionRequest struct {
	*aws.Request
	Input *EvaluateExpressionInput
	Copy  func(*EvaluateExpressionInput) EvaluateExpressionRequest
}

// Send marshals and sends the EvaluateExpression API request.
func (r EvaluateExpressionRequest) Send(ctx context.Context) (*EvaluateExpressionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EvaluateExpressionResponse{
		EvaluateExpressionOutput: r.Request.Data.(*EvaluateExpressionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EvaluateExpressionResponse is the response type for the
// EvaluateExpression API operation.
type EvaluateExpressionResponse struct {
	*EvaluateExpressionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EvaluateExpression request.
func (r *EvaluateExpressionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
