// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestar

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeUserProfileInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the user.
	//
	// UserArn is a required field
	UserArn *string `locationName:"userArn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeUserProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeUserProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeUserProfileInput"}

	if s.UserArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserArn"))
	}
	if s.UserArn != nil && len(*s.UserArn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("UserArn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeUserProfileOutput struct {
	_ struct{} `type:"structure"`

	// The date and time when the user profile was created in AWS CodeStar, in timestamp
	// format.
	//
	// CreatedTimestamp is a required field
	CreatedTimestamp *time.Time `locationName:"createdTimestamp" type:"timestamp" required:"true"`

	// The display name shown for the user in AWS CodeStar projects. For example,
	// this could be set to both first and last name ("Mary Major") or a single
	// name ("Mary"). The display name is also used to generate the initial icon
	// associated with the user in AWS CodeStar projects. If spaces are included
	// in the display name, the first character that appears after the space will
	// be used as the second character in the user initial icon. The initial icon
	// displays a maximum of two characters, so a display name with more than one
	// space (for example "Mary Jane Major") would generate an initial icon using
	// the first character and the first character after the space ("MJ", not "MM").
	DisplayName *string `locationName:"displayName" min:"1" type:"string" sensitive:"true"`

	// The email address for the user. Optional.
	EmailAddress *string `locationName:"emailAddress" min:"3" type:"string" sensitive:"true"`

	// The date and time when the user profile was last modified, in timestamp format.
	//
	// LastModifiedTimestamp is a required field
	LastModifiedTimestamp *time.Time `locationName:"lastModifiedTimestamp" type:"timestamp" required:"true"`

	// The SSH public key associated with the user. This SSH public key is associated
	// with the user profile, and can be used in conjunction with the associated
	// private key for access to project resources, such as Amazon EC2 instances,
	// if a project owner grants remote access to those resources.
	SshPublicKey *string `locationName:"sshPublicKey" type:"string"`

	// The Amazon Resource Name (ARN) of the user.
	//
	// UserArn is a required field
	UserArn *string `locationName:"userArn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeUserProfileOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeUserProfile = "DescribeUserProfile"

// DescribeUserProfileRequest returns a request value for making API operation for
// AWS CodeStar.
//
// Describes a user in AWS CodeStar and the user attributes across all projects.
//
//    // Example sending a request using DescribeUserProfileRequest.
//    req := client.DescribeUserProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/DescribeUserProfile
func (c *Client) DescribeUserProfileRequest(input *DescribeUserProfileInput) DescribeUserProfileRequest {
	op := &aws.Operation{
		Name:       opDescribeUserProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeUserProfileInput{}
	}

	req := c.newRequest(op, input, &DescribeUserProfileOutput{})

	return DescribeUserProfileRequest{Request: req, Input: input, Copy: c.DescribeUserProfileRequest}
}

// DescribeUserProfileRequest is the request type for the
// DescribeUserProfile API operation.
type DescribeUserProfileRequest struct {
	*aws.Request
	Input *DescribeUserProfileInput
	Copy  func(*DescribeUserProfileInput) DescribeUserProfileRequest
}

// Send marshals and sends the DescribeUserProfile API request.
func (r DescribeUserProfileRequest) Send(ctx context.Context) (*DescribeUserProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeUserProfileResponse{
		DescribeUserProfileOutput: r.Request.Data.(*DescribeUserProfileOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeUserProfileResponse is the response type for the
// DescribeUserProfile API operation.
type DescribeUserProfileResponse struct {
	*DescribeUserProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeUserProfile request.
func (r *DescribeUserProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
