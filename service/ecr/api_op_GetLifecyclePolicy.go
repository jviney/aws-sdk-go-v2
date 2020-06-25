// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetLifecyclePolicyInput struct {
	_ struct{} `type:"structure"`

	// The AWS account ID associated with the registry that contains the repository.
	// If you do not specify a registry, the default registry is assumed.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The name of the repository.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s GetLifecyclePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLifecyclePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLifecyclePolicyInput"}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetLifecyclePolicyOutput struct {
	_ struct{} `type:"structure"`

	// The time stamp of the last time that the lifecycle policy was run.
	LastEvaluatedAt *time.Time `locationName:"lastEvaluatedAt" type:"timestamp"`

	// The JSON lifecycle policy text.
	LifecyclePolicyText *string `locationName:"lifecyclePolicyText" min:"100" type:"string"`

	// The registry ID associated with the request.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The repository name associated with the request.
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string"`
}

// String returns the string representation
func (s GetLifecyclePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetLifecyclePolicy = "GetLifecyclePolicy"

// GetLifecyclePolicyRequest returns a request value for making API operation for
// Amazon EC2 Container Registry.
//
// Retrieves the lifecycle policy for the specified repository.
//
//    // Example sending a request using GetLifecyclePolicyRequest.
//    req := client.GetLifecyclePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/GetLifecyclePolicy
func (c *Client) GetLifecyclePolicyRequest(input *GetLifecyclePolicyInput) GetLifecyclePolicyRequest {
	op := &aws.Operation{
		Name:       opGetLifecyclePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetLifecyclePolicyInput{}
	}

	req := c.newRequest(op, input, &GetLifecyclePolicyOutput{})

	return GetLifecyclePolicyRequest{Request: req, Input: input, Copy: c.GetLifecyclePolicyRequest}
}

// GetLifecyclePolicyRequest is the request type for the
// GetLifecyclePolicy API operation.
type GetLifecyclePolicyRequest struct {
	*aws.Request
	Input *GetLifecyclePolicyInput
	Copy  func(*GetLifecyclePolicyInput) GetLifecyclePolicyRequest
}

// Send marshals and sends the GetLifecyclePolicy API request.
func (r GetLifecyclePolicyRequest) Send(ctx context.Context) (*GetLifecyclePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLifecyclePolicyResponse{
		GetLifecyclePolicyOutput: r.Request.Data.(*GetLifecyclePolicyOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLifecyclePolicyResponse is the response type for the
// GetLifecyclePolicy API operation.
type GetLifecyclePolicyResponse struct {
	*GetLifecyclePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLifecyclePolicy request.
func (r *GetLifecyclePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
