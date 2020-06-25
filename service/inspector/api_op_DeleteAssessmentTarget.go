// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteAssessmentTargetInput struct {
	_ struct{} `type:"structure"`

	// The ARN that specifies the assessment target that you want to delete.
	//
	// AssessmentTargetArn is a required field
	AssessmentTargetArn *string `locationName:"assessmentTargetArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAssessmentTargetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAssessmentTargetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAssessmentTargetInput"}

	if s.AssessmentTargetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssessmentTargetArn"))
	}
	if s.AssessmentTargetArn != nil && len(*s.AssessmentTargetArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssessmentTargetArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteAssessmentTargetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAssessmentTargetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteAssessmentTarget = "DeleteAssessmentTarget"

// DeleteAssessmentTargetRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Deletes the assessment target that is specified by the ARN of the assessment
// target.
//
//    // Example sending a request using DeleteAssessmentTargetRequest.
//    req := client.DeleteAssessmentTargetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/DeleteAssessmentTarget
func (c *Client) DeleteAssessmentTargetRequest(input *DeleteAssessmentTargetInput) DeleteAssessmentTargetRequest {
	op := &aws.Operation{
		Name:       opDeleteAssessmentTarget,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteAssessmentTargetInput{}
	}

	req := c.newRequest(op, input, &DeleteAssessmentTargetOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteAssessmentTargetRequest{Request: req, Input: input, Copy: c.DeleteAssessmentTargetRequest}
}

// DeleteAssessmentTargetRequest is the request type for the
// DeleteAssessmentTarget API operation.
type DeleteAssessmentTargetRequest struct {
	*aws.Request
	Input *DeleteAssessmentTargetInput
	Copy  func(*DeleteAssessmentTargetInput) DeleteAssessmentTargetRequest
}

// Send marshals and sends the DeleteAssessmentTarget API request.
func (r DeleteAssessmentTargetRequest) Send(ctx context.Context) (*DeleteAssessmentTargetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAssessmentTargetResponse{
		DeleteAssessmentTargetOutput: r.Request.Data.(*DeleteAssessmentTargetOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAssessmentTargetResponse is the response type for the
// DeleteAssessmentTarget API operation.
type DeleteAssessmentTargetResponse struct {
	*DeleteAssessmentTargetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAssessmentTarget request.
func (r *DeleteAssessmentTargetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
