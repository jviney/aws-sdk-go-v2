// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeSessionsInput struct {
	_ struct{} `type:"structure"`

	// The authentication method. Specify API for a user authenticated using a streaming
	// URL or SAML for a SAML federated user. The default is to authenticate users
	// using a streaming URL.
	AuthenticationType AuthenticationType `type:"string" enum:"true"`

	// The name of the fleet. This value is case-sensitive.
	//
	// FleetName is a required field
	FleetName *string `min:"1" type:"string" required:"true"`

	// The size of each page of results. The default value is 20 and the maximum
	// value is 50.
	Limit *int64 `type:"integer"`

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string `min:"1" type:"string"`

	// The name of the stack. This value is case-sensitive.
	//
	// StackName is a required field
	StackName *string `min:"1" type:"string" required:"true"`

	// The user identifier.
	UserId *string `min:"2" type:"string"`
}

// String returns the string representation
func (s DescribeSessionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSessionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSessionsInput"}

	if s.FleetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetName"))
	}
	if s.FleetName != nil && len(*s.FleetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetName", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.StackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackName"))
	}
	if s.StackName != nil && len(*s.StackName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StackName", 1))
	}
	if s.UserId != nil && len(*s.UserId) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("UserId", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeSessionsOutput struct {
	_ struct{} `type:"structure"`

	// The pagination token to use to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string `min:"1" type:"string"`

	// Information about the streaming sessions.
	Sessions []Session `type:"list"`
}

// String returns the string representation
func (s DescribeSessionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSessions = "DescribeSessions"

// DescribeSessionsRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Retrieves a list that describes the streaming sessions for a specified stack
// and fleet. If a UserId is provided for the stack and fleet, only streaming
// sessions for that user are described. If an authentication type is not provided,
// the default is to authenticate users using a streaming URL.
//
//    // Example sending a request using DescribeSessionsRequest.
//    req := client.DescribeSessionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DescribeSessions
func (c *Client) DescribeSessionsRequest(input *DescribeSessionsInput) DescribeSessionsRequest {
	op := &aws.Operation{
		Name:       opDescribeSessions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSessionsInput{}
	}

	req := c.newRequest(op, input, &DescribeSessionsOutput{})

	return DescribeSessionsRequest{Request: req, Input: input, Copy: c.DescribeSessionsRequest}
}

// DescribeSessionsRequest is the request type for the
// DescribeSessions API operation.
type DescribeSessionsRequest struct {
	*aws.Request
	Input *DescribeSessionsInput
	Copy  func(*DescribeSessionsInput) DescribeSessionsRequest
}

// Send marshals and sends the DescribeSessions API request.
func (r DescribeSessionsRequest) Send(ctx context.Context) (*DescribeSessionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSessionsResponse{
		DescribeSessionsOutput: r.Request.Data.(*DescribeSessionsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSessionsResponse is the response type for the
// DescribeSessions API operation.
type DescribeSessionsResponse struct {
	*DescribeSessionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSessions request.
func (r *DescribeSessionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
