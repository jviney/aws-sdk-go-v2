// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListDiscoveredResourcesInput struct {
	_ struct{} `type:"structure"`

	// Specifies whether AWS Config includes deleted resources in the results. By
	// default, deleted resources are not included.
	IncludeDeletedResources *bool `locationName:"includeDeletedResources" type:"boolean"`

	// The maximum number of resource identifiers returned on each page. The default
	// is 100. You cannot specify a number greater than 100. If you specify 0, AWS
	// Config uses the default.
	Limit *int64 `locationName:"limit" type:"integer"`

	// The nextToken string returned on a previous page that you use to get the
	// next page of results in a paginated response.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The IDs of only those resources that you want AWS Config to list in the response.
	// If you do not specify this parameter, AWS Config lists all resources of the
	// specified type that it has discovered.
	ResourceIds []string `locationName:"resourceIds" type:"list"`

	// The custom name of only those resources that you want AWS Config to list
	// in the response. If you do not specify this parameter, AWS Config lists all
	// resources of the specified type that it has discovered.
	ResourceName *string `locationName:"resourceName" type:"string"`

	// The type of resources that you want AWS Config to list in the response.
	//
	// ResourceType is a required field
	ResourceType ResourceType `locationName:"resourceType" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s ListDiscoveredResourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDiscoveredResourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDiscoveredResourcesInput"}
	if len(s.ResourceType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ResourceType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListDiscoveredResourcesOutput struct {
	_ struct{} `type:"structure"`

	// The string that you use in a subsequent request to get the next page of results
	// in a paginated response.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The details that identify a resource that is discovered by AWS Config, including
	// the resource type, ID, and (if available) the custom resource name.
	ResourceIdentifiers []ResourceIdentifier `locationName:"resourceIdentifiers" type:"list"`
}

// String returns the string representation
func (s ListDiscoveredResourcesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDiscoveredResources = "ListDiscoveredResources"

// ListDiscoveredResourcesRequest returns a request value for making API operation for
// AWS Config.
//
// Accepts a resource type and returns a list of resource identifiers for the
// resources of that type. A resource identifier includes the resource type,
// ID, and (if available) the custom resource name. The results consist of resources
// that AWS Config has discovered, including those that AWS Config is not currently
// recording. You can narrow the results to include only resources that have
// specific resource IDs or a resource name.
//
// You can specify either resource IDs or a resource name, but not both, in
// the same request.
//
// The response is paginated. By default, AWS Config lists 100 resource identifiers
// on each page. You can customize this number with the limit parameter. The
// response includes a nextToken string. To get the next page of results, run
// the request again and specify the string for the nextToken parameter.
//
//    // Example sending a request using ListDiscoveredResourcesRequest.
//    req := client.ListDiscoveredResourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/ListDiscoveredResources
func (c *Client) ListDiscoveredResourcesRequest(input *ListDiscoveredResourcesInput) ListDiscoveredResourcesRequest {
	op := &aws.Operation{
		Name:       opListDiscoveredResources,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListDiscoveredResourcesInput{}
	}

	req := c.newRequest(op, input, &ListDiscoveredResourcesOutput{})

	return ListDiscoveredResourcesRequest{Request: req, Input: input, Copy: c.ListDiscoveredResourcesRequest}
}

// ListDiscoveredResourcesRequest is the request type for the
// ListDiscoveredResources API operation.
type ListDiscoveredResourcesRequest struct {
	*aws.Request
	Input *ListDiscoveredResourcesInput
	Copy  func(*ListDiscoveredResourcesInput) ListDiscoveredResourcesRequest
}

// Send marshals and sends the ListDiscoveredResources API request.
func (r ListDiscoveredResourcesRequest) Send(ctx context.Context) (*ListDiscoveredResourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDiscoveredResourcesResponse{
		ListDiscoveredResourcesOutput: r.Request.Data.(*ListDiscoveredResourcesOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDiscoveredResourcesResponse is the response type for the
// ListDiscoveredResources API operation.
type ListDiscoveredResourcesResponse struct {
	*ListDiscoveredResourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDiscoveredResources request.
func (r *ListDiscoveredResourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
