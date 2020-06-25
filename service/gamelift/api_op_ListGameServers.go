// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListGameServersInput struct {
	_ struct{} `type:"structure"`

	// An identifier for the game server group for the game server you want to list.
	// Use either the GameServerGroup name or ARN value.
	//
	// GameServerGroupName is a required field
	GameServerGroupName *string `min:"1" type:"string" required:"true"`

	// The maximum number of results to return. Use this parameter with NextToken
	// to get results as a set of sequential pages.
	Limit *int64 `min:"1" type:"integer"`

	// A token that indicates the start of the next sequential page of results.
	// Use the token that is returned with a previous call to this action. To start
	// at the beginning of the result set, do not specify a value.
	NextToken *string `min:"1" type:"string"`

	// Indicates how to sort the returned data based on the game servers' custom
	// key sort value. If this parameter is left empty, the list of game servers
	// is returned in no particular order.
	SortOrder SortOrder `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListGameServersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListGameServersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListGameServersInput"}

	if s.GameServerGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GameServerGroupName"))
	}
	if s.GameServerGroupName != nil && len(*s.GameServerGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameServerGroupName", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListGameServersOutput struct {
	_ struct{} `type:"structure"`

	// A collection of game server objects that match the request.
	GameServers []GameServer `type:"list"`

	// A token that indicates where to resume retrieving results on the next call
	// to this action. If no token is returned, these results represent the end
	// of the list.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListGameServersOutput) String() string {
	return awsutil.Prettify(s)
}

const opListGameServers = "ListGameServers"

// ListGameServersRequest returns a request value for making API operation for
// Amazon GameLift.
//
// This action is part of Amazon GameLift FleetIQ with game server groups, which
// is in preview release and is subject to change.
//
// Retrieves information on all game servers that are currently running in a
// specified game server group. If there are custom key sort values for your
// game servers, you can opt to have the returned list sorted based on these
// values. Use the pagination parameters to retrieve results in a set of sequential
// pages.
//
// Learn more
//
// GameLift FleetIQ Guide (https://docs.aws.amazon.com/gamelift/latest/developerguide/gsg-intro.html)
//
// Related operations
//
//    * RegisterGameServer
//
//    * ListGameServers
//
//    * ClaimGameServer
//
//    * DescribeGameServer
//
//    * UpdateGameServer
//
//    * DeregisterGameServer
//
//    // Example sending a request using ListGameServersRequest.
//    req := client.ListGameServersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/ListGameServers
func (c *Client) ListGameServersRequest(input *ListGameServersInput) ListGameServersRequest {
	op := &aws.Operation{
		Name:       opListGameServers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListGameServersInput{}
	}

	req := c.newRequest(op, input, &ListGameServersOutput{})

	return ListGameServersRequest{Request: req, Input: input, Copy: c.ListGameServersRequest}
}

// ListGameServersRequest is the request type for the
// ListGameServers API operation.
type ListGameServersRequest struct {
	*aws.Request
	Input *ListGameServersInput
	Copy  func(*ListGameServersInput) ListGameServersRequest
}

// Send marshals and sends the ListGameServers API request.
func (r ListGameServersRequest) Send(ctx context.Context) (*ListGameServersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGameServersResponse{
		ListGameServersOutput: r.Request.Data.(*ListGameServersOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListGameServersResponse is the response type for the
// ListGameServers API operation.
type ListGameServersResponse struct {
	*ListGameServersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGameServers request.
func (r *ListGameServersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
