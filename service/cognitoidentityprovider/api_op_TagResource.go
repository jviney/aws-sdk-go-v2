// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type TagResourceInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the user pool to assign the tags to.
	//
	// ResourceArn is a required field
	ResourceArn *string `min:"20" type:"string" required:"true"`

	// The tags to assign to the user pool.
	//
	// Tags is a required field
	Tags map[string]string `type:"map" required:"true"`
}

// String returns the string representation
func (s TagResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagResourceInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}
	if s.ResourceArn != nil && len(*s.ResourceArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 20))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TagResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TagResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opTagResource = "TagResource"

// TagResourceRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Assigns a set of tags to an Amazon Cognito user pool. A tag is a label that
// you can use to categorize and manage user pools in different ways, such as
// by purpose, owner, environment, or other criteria.
//
// Each tag consists of a key and value, both of which you define. A key is
// a general category for more specific values. For example, if you have two
// versions of a user pool, one for testing and another for production, you
// might assign an Environment tag key to both user pools. The value of this
// key might be Test for one user pool and Production for the other.
//
// Tags are useful for cost tracking and access control. You can activate your
// tags so that they appear on the Billing and Cost Management console, where
// you can track the costs associated with your user pools. In an IAM policy,
// you can constrain permissions for user pools based on specific tags or tag
// values.
//
// You can use this action up to 5 times per second, per account. A user pool
// can have as many as 50 tags.
//
//    // Example sending a request using TagResourceRequest.
//    req := client.TagResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/TagResource
func (c *Client) TagResourceRequest(input *TagResourceInput) TagResourceRequest {
	op := &aws.Operation{
		Name:       opTagResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TagResourceInput{}
	}

	req := c.newRequest(op, input, &TagResourceOutput{})

	return TagResourceRequest{Request: req, Input: input, Copy: c.TagResourceRequest}
}

// TagResourceRequest is the request type for the
// TagResource API operation.
type TagResourceRequest struct {
	*aws.Request
	Input *TagResourceInput
	Copy  func(*TagResourceInput) TagResourceRequest
}

// Send marshals and sends the TagResource API request.
func (r TagResourceRequest) Send(ctx context.Context) (*TagResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TagResourceResponse{
		TagResourceOutput: r.Request.Data.(*TagResourceOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TagResourceResponse is the response type for the
// TagResource API operation.
type TagResourceResponse struct {
	*TagResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TagResource request.
func (r *TagResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
