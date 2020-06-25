// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Request to create or update a group of environments.
type ComposeEnvironmentsInput struct {
	_ struct{} `type:"structure"`

	// The name of the application to which the specified source bundles belong.
	ApplicationName *string `min:"1" type:"string"`

	// The name of the group to which the target environments belong. Specify a
	// group name only if the environment name defined in each target environment's
	// manifest ends with a + (plus) character. See Environment Manifest (env.yaml)
	// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/environment-cfg-manifest.html)
	// for details.
	GroupName *string `min:"1" type:"string"`

	// A list of version labels, specifying one or more application source bundles
	// that belong to the target application. Each source bundle must include an
	// environment manifest that specifies the name of the environment and the name
	// of the solution stack to use, and optionally can specify environment links
	// to create.
	VersionLabels []string `type:"list"`
}

// String returns the string representation
func (s ComposeEnvironmentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ComposeEnvironmentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ComposeEnvironmentsInput"}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}
	if s.GroupName != nil && len(*s.GroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Result message containing a list of environment descriptions.
type ComposeEnvironmentsOutput struct {
	_ struct{} `type:"structure"`

	// Returns an EnvironmentDescription list.
	Environments []EnvironmentDescription `type:"list"`

	// In a paginated request, the token that you can pass in a subsequent request
	// to get the next response page.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ComposeEnvironmentsOutput) String() string {
	return awsutil.Prettify(s)
}

const opComposeEnvironments = "ComposeEnvironments"

// ComposeEnvironmentsRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Create or update a group of environments that each run a separate component
// of a single application. Takes a list of version labels that specify application
// source bundles for each of the environments to create or update. The name
// of each environment and other required information must be included in the
// source bundles in an environment manifest named env.yaml. See Compose Environments
// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/environment-mgmt-compose.html)
// for details.
//
//    // Example sending a request using ComposeEnvironmentsRequest.
//    req := client.ComposeEnvironmentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/ComposeEnvironments
func (c *Client) ComposeEnvironmentsRequest(input *ComposeEnvironmentsInput) ComposeEnvironmentsRequest {
	op := &aws.Operation{
		Name:       opComposeEnvironments,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ComposeEnvironmentsInput{}
	}

	req := c.newRequest(op, input, &ComposeEnvironmentsOutput{})

	return ComposeEnvironmentsRequest{Request: req, Input: input, Copy: c.ComposeEnvironmentsRequest}
}

// ComposeEnvironmentsRequest is the request type for the
// ComposeEnvironments API operation.
type ComposeEnvironmentsRequest struct {
	*aws.Request
	Input *ComposeEnvironmentsInput
	Copy  func(*ComposeEnvironmentsInput) ComposeEnvironmentsRequest
}

// Send marshals and sends the ComposeEnvironments API request.
func (r ComposeEnvironmentsRequest) Send(ctx context.Context) (*ComposeEnvironmentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ComposeEnvironmentsResponse{
		ComposeEnvironmentsOutput: r.Request.Data.(*ComposeEnvironmentsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ComposeEnvironmentsResponse is the response type for the
// ComposeEnvironments API operation.
type ComposeEnvironmentsResponse struct {
	*ComposeEnvironmentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ComposeEnvironments request.
func (r *ComposeEnvironmentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
