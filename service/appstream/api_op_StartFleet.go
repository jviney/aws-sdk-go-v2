// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type StartFleetInput struct {
	_ struct{} `type:"structure"`

	// The name of the fleet.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartFleetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartFleetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartFleetInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartFleetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartFleetOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartFleet = "StartFleet"

// StartFleetRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Starts the specified fleet.
//
//    // Example sending a request using StartFleetRequest.
//    req := client.StartFleetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/StartFleet
func (c *Client) StartFleetRequest(input *StartFleetInput) StartFleetRequest {
	op := &aws.Operation{
		Name:       opStartFleet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartFleetInput{}
	}

	req := c.newRequest(op, input, &StartFleetOutput{})

	return StartFleetRequest{Request: req, Input: input, Copy: c.StartFleetRequest}
}

// StartFleetRequest is the request type for the
// StartFleet API operation.
type StartFleetRequest struct {
	*aws.Request
	Input *StartFleetInput
	Copy  func(*StartFleetInput) StartFleetRequest
}

// Send marshals and sends the StartFleet API request.
func (r StartFleetRequest) Send(ctx context.Context) (*StartFleetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartFleetResponse{
		StartFleetOutput: r.Request.Data.(*StartFleetOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartFleetResponse is the response type for the
// StartFleet API operation.
type StartFleetResponse struct {
	*StartFleetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartFleet request.
func (r *StartFleetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
