// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type StopGameSessionPlacementInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for a game session placement to cancel.
	//
	// PlacementId is a required field
	PlacementId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopGameSessionPlacementInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopGameSessionPlacementInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopGameSessionPlacementInput"}

	if s.PlacementId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlacementId"))
	}
	if s.PlacementId != nil && len(*s.PlacementId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PlacementId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type StopGameSessionPlacementOutput struct {
	_ struct{} `type:"structure"`

	// Object that describes the canceled game session placement, with CANCELLED
	// status and an end time stamp.
	GameSessionPlacement *GameSessionPlacement `type:"structure"`
}

// String returns the string representation
func (s StopGameSessionPlacementOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopGameSessionPlacement = "StopGameSessionPlacement"

// StopGameSessionPlacementRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Cancels a game session placement that is in PENDING status. To stop a placement,
// provide the placement ID values. If successful, the placement is moved to
// CANCELLED status.
//
//    * CreateGameSession
//
//    * DescribeGameSessions
//
//    * DescribeGameSessionDetails
//
//    * SearchGameSessions
//
//    * UpdateGameSession
//
//    * GetGameSessionLogUrl
//
//    * Game session placements StartGameSessionPlacement DescribeGameSessionPlacement
//    StopGameSessionPlacement
//
//    // Example sending a request using StopGameSessionPlacementRequest.
//    req := client.StopGameSessionPlacementRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/StopGameSessionPlacement
func (c *Client) StopGameSessionPlacementRequest(input *StopGameSessionPlacementInput) StopGameSessionPlacementRequest {
	op := &aws.Operation{
		Name:       opStopGameSessionPlacement,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopGameSessionPlacementInput{}
	}

	req := c.newRequest(op, input, &StopGameSessionPlacementOutput{})

	return StopGameSessionPlacementRequest{Request: req, Input: input, Copy: c.StopGameSessionPlacementRequest}
}

// StopGameSessionPlacementRequest is the request type for the
// StopGameSessionPlacement API operation.
type StopGameSessionPlacementRequest struct {
	*aws.Request
	Input *StopGameSessionPlacementInput
	Copy  func(*StopGameSessionPlacementInput) StopGameSessionPlacementRequest
}

// Send marshals and sends the StopGameSessionPlacement API request.
func (r StopGameSessionPlacementRequest) Send(ctx context.Context) (*StopGameSessionPlacementResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopGameSessionPlacementResponse{
		StopGameSessionPlacementOutput: r.Request.Data.(*StopGameSessionPlacementOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopGameSessionPlacementResponse is the response type for the
// StopGameSessionPlacement API operation.
type StopGameSessionPlacementResponse struct {
	*StopGameSessionPlacementOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopGameSessionPlacement request.
func (r *StopGameSessionPlacementResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
