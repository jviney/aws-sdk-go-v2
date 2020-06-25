// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DeregisterRobotInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the fleet.
	//
	// Fleet is a required field
	Fleet *string `locationName:"fleet" min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the robot.
	//
	// Robot is a required field
	Robot *string `locationName:"robot" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeregisterRobotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterRobotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeregisterRobotInput"}

	if s.Fleet == nil {
		invalidParams.Add(aws.NewErrParamRequired("Fleet"))
	}
	if s.Fleet != nil && len(*s.Fleet) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Fleet", 1))
	}

	if s.Robot == nil {
		invalidParams.Add(aws.NewErrParamRequired("Robot"))
	}
	if s.Robot != nil && len(*s.Robot) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Robot", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeregisterRobotInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Fleet != nil {
		v := *s.Fleet

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "fleet", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Robot != nil {
		v := *s.Robot

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "robot", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeregisterRobotOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the fleet.
	Fleet *string `locationName:"fleet" min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the robot.
	Robot *string `locationName:"robot" min:"1" type:"string"`
}

// String returns the string representation
func (s DeregisterRobotOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeregisterRobotOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Fleet != nil {
		v := *s.Fleet

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "fleet", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Robot != nil {
		v := *s.Robot

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "robot", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDeregisterRobot = "DeregisterRobot"

// DeregisterRobotRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Deregisters a robot.
//
//    // Example sending a request using DeregisterRobotRequest.
//    req := client.DeregisterRobotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/DeregisterRobot
func (c *Client) DeregisterRobotRequest(input *DeregisterRobotInput) DeregisterRobotRequest {
	op := &aws.Operation{
		Name:       opDeregisterRobot,
		HTTPMethod: "POST",
		HTTPPath:   "/deregisterRobot",
	}

	if input == nil {
		input = &DeregisterRobotInput{}
	}

	req := c.newRequest(op, input, &DeregisterRobotOutput{})

	return DeregisterRobotRequest{Request: req, Input: input, Copy: c.DeregisterRobotRequest}
}

// DeregisterRobotRequest is the request type for the
// DeregisterRobot API operation.
type DeregisterRobotRequest struct {
	*aws.Request
	Input *DeregisterRobotInput
	Copy  func(*DeregisterRobotInput) DeregisterRobotRequest
}

// Send marshals and sends the DeregisterRobot API request.
func (r DeregisterRobotRequest) Send(ctx context.Context) (*DeregisterRobotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterRobotResponse{
		DeregisterRobotOutput: r.Request.Data.(*DeregisterRobotOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterRobotResponse is the response type for the
// DeregisterRobot API operation.
type DeregisterRobotResponse struct {
	*DeregisterRobotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterRobot request.
func (r *DeregisterRobotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
