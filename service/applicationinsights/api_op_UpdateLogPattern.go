// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationinsights

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type UpdateLogPatternInput struct {
	_ struct{} `type:"structure"`

	// The log pattern.
	Pattern *string `min:"1" type:"string"`

	// The name of the log pattern.
	//
	// PatternName is a required field
	PatternName *string `min:"1" type:"string" required:"true"`

	// The name of the log pattern set.
	//
	// PatternSetName is a required field
	PatternSetName *string `min:"1" type:"string" required:"true"`

	// Rank of the log pattern.
	Rank *int64 `type:"integer"`

	// The name of the resource group.
	//
	// ResourceGroupName is a required field
	ResourceGroupName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateLogPatternInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateLogPatternInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateLogPatternInput"}
	if s.Pattern != nil && len(*s.Pattern) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Pattern", 1))
	}

	if s.PatternName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PatternName"))
	}
	if s.PatternName != nil && len(*s.PatternName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PatternName", 1))
	}

	if s.PatternSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PatternSetName"))
	}
	if s.PatternSetName != nil && len(*s.PatternSetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PatternSetName", 1))
	}

	if s.ResourceGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceGroupName"))
	}
	if s.ResourceGroupName != nil && len(*s.ResourceGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateLogPatternOutput struct {
	_ struct{} `type:"structure"`

	// The successfully created log pattern.
	LogPattern *LogPattern `type:"structure"`

	// The name of the resource group.
	ResourceGroupName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateLogPatternOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateLogPattern = "UpdateLogPattern"

// UpdateLogPatternRequest returns a request value for making API operation for
// Amazon CloudWatch Application Insights.
//
// Adds a log pattern to a LogPatternSet.
//
//    // Example sending a request using UpdateLogPatternRequest.
//    req := client.UpdateLogPatternRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-insights-2018-11-25/UpdateLogPattern
func (c *Client) UpdateLogPatternRequest(input *UpdateLogPatternInput) UpdateLogPatternRequest {
	op := &aws.Operation{
		Name:       opUpdateLogPattern,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateLogPatternInput{}
	}

	req := c.newRequest(op, input, &UpdateLogPatternOutput{})

	return UpdateLogPatternRequest{Request: req, Input: input, Copy: c.UpdateLogPatternRequest}
}

// UpdateLogPatternRequest is the request type for the
// UpdateLogPattern API operation.
type UpdateLogPatternRequest struct {
	*aws.Request
	Input *UpdateLogPatternInput
	Copy  func(*UpdateLogPatternInput) UpdateLogPatternRequest
}

// Send marshals and sends the UpdateLogPattern API request.
func (r UpdateLogPatternRequest) Send(ctx context.Context) (*UpdateLogPatternResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateLogPatternResponse{
		UpdateLogPatternOutput: r.Request.Data.(*UpdateLogPatternOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateLogPatternResponse is the response type for the
// UpdateLogPattern API operation.
type UpdateLogPatternResponse struct {
	*UpdateLogPatternOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateLogPattern request.
func (r *UpdateLogPatternResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
