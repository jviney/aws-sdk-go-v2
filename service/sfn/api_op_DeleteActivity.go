// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sfn

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteActivityInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the activity to delete.
	//
	// ActivityArn is a required field
	ActivityArn *string `locationName:"activityArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteActivityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteActivityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteActivityInput"}

	if s.ActivityArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActivityArn"))
	}
	if s.ActivityArn != nil && len(*s.ActivityArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ActivityArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteActivityOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteActivityOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteActivity = "DeleteActivity"

// DeleteActivityRequest returns a request value for making API operation for
// AWS Step Functions.
//
// Deletes an activity.
//
//    // Example sending a request using DeleteActivityRequest.
//    req := client.DeleteActivityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23/DeleteActivity
func (c *Client) DeleteActivityRequest(input *DeleteActivityInput) DeleteActivityRequest {
	op := &aws.Operation{
		Name:       opDeleteActivity,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteActivityInput{}
	}

	req := c.newRequest(op, input, &DeleteActivityOutput{})

	return DeleteActivityRequest{Request: req, Input: input, Copy: c.DeleteActivityRequest}
}

// DeleteActivityRequest is the request type for the
// DeleteActivity API operation.
type DeleteActivityRequest struct {
	*aws.Request
	Input *DeleteActivityInput
	Copy  func(*DeleteActivityInput) DeleteActivityRequest
}

// Send marshals and sends the DeleteActivity API request.
func (r DeleteActivityRequest) Send(ctx context.Context) (*DeleteActivityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteActivityResponse{
		DeleteActivityOutput: r.Request.Data.(*DeleteActivityOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteActivityResponse is the response type for the
// DeleteActivity API operation.
type DeleteActivityResponse struct {
	*DeleteActivityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteActivity request.
func (r *DeleteActivityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
