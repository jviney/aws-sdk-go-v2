// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package resourcegroupstaggingapi

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type TagResourcesInput struct {
	_ struct{} `type:"structure"`

	// A list of ARNs. An ARN (Amazon Resource Name) uniquely identifies a resource.
	// For more information, see Amazon Resource Names (ARNs) and AWS Service Namespaces
	// (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	//
	// ResourceARNList is a required field
	ResourceARNList []string `min:"1" type:"list" required:"true"`

	// The tags that you want to add to the specified resources. A tag consists
	// of a key and a value that you define.
	//
	// Tags is a required field
	Tags map[string]string `min:"1" type:"map" required:"true"`
}

// String returns the string representation
func (s TagResourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagResourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagResourcesInput"}

	if s.ResourceARNList == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceARNList"))
	}
	if s.ResourceARNList != nil && len(s.ResourceARNList) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceARNList", 1))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TagResourcesOutput struct {
	_ struct{} `type:"structure"`

	// A map containing a key-value pair for each failed item that couldn't be tagged.
	// The key is the ARN of the failed resource. The value is a FailureInfo object
	// that contains an error code, a status code, and an error message. If there
	// are no errors, the FailedResourcesMap is empty.
	FailedResourcesMap map[string]FailureInfo `type:"map"`
}

// String returns the string representation
func (s TagResourcesOutput) String() string {
	return awsutil.Prettify(s)
}

const opTagResources = "TagResources"

// TagResourcesRequest returns a request value for making API operation for
// AWS Resource Groups Tagging API.
//
// Applies one or more tags to the specified resources. Note the following:
//
//    * Not all resources can have tags. For a list of services that support
//    tagging, see this list (http://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/Welcome.html).
//
//    * Each resource can have up to 50 tags. For other limits, see Tag Naming
//    and Usage Conventions (http://docs.aws.amazon.com/general/latest/gr/aws_tagging.html#tag-conventions)
//    in the AWS General Reference.
//
//    * You can only tag resources that are located in the specified Region
//    for the AWS account.
//
//    * To add tags to a resource, you need the necessary permissions for the
//    service that the resource belongs to as well as permissions for adding
//    tags. For more information, see this list (http://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/Welcome.html).
//
//    // Example sending a request using TagResourcesRequest.
//    req := client.TagResourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/resourcegroupstaggingapi-2017-01-26/TagResources
func (c *Client) TagResourcesRequest(input *TagResourcesInput) TagResourcesRequest {
	op := &aws.Operation{
		Name:       opTagResources,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TagResourcesInput{}
	}

	req := c.newRequest(op, input, &TagResourcesOutput{})

	return TagResourcesRequest{Request: req, Input: input, Copy: c.TagResourcesRequest}
}

// TagResourcesRequest is the request type for the
// TagResources API operation.
type TagResourcesRequest struct {
	*aws.Request
	Input *TagResourcesInput
	Copy  func(*TagResourcesInput) TagResourcesRequest
}

// Send marshals and sends the TagResources API request.
func (r TagResourcesRequest) Send(ctx context.Context) (*TagResourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TagResourcesResponse{
		TagResourcesOutput: r.Request.Data.(*TagResourcesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TagResourcesResponse is the response type for the
// TagResources API operation.
type TagResourcesResponse struct {
	*TagResourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TagResources request.
func (r *TagResourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
