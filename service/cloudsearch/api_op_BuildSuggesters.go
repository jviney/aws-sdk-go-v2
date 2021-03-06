// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudsearch

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Container for the parameters to the BuildSuggester operation. Specifies the
// name of the domain you want to update.
type BuildSuggestersInput struct {
	_ struct{} `type:"structure"`

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start
	// with a letter or number and can contain the following characters: a-z (lowercase),
	// 0-9, and - (hyphen).
	//
	// DomainName is a required field
	DomainName *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s BuildSuggestersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BuildSuggestersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BuildSuggestersInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of a BuildSuggester request. Contains a list of the fields used
// for suggestions.
type BuildSuggestersOutput struct {
	_ struct{} `type:"structure"`

	// A list of field names.
	FieldNames []string `type:"list"`
}

// String returns the string representation
func (s BuildSuggestersOutput) String() string {
	return awsutil.Prettify(s)
}

const opBuildSuggesters = "BuildSuggesters"

// BuildSuggestersRequest returns a request value for making API operation for
// Amazon CloudSearch.
//
// Indexes the search suggestions. For more information, see Configuring Suggesters
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-suggestions.html#configuring-suggesters)
// in the Amazon CloudSearch Developer Guide.
//
//    // Example sending a request using BuildSuggestersRequest.
//    req := client.BuildSuggestersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) BuildSuggestersRequest(input *BuildSuggestersInput) BuildSuggestersRequest {
	op := &aws.Operation{
		Name:       opBuildSuggesters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BuildSuggestersInput{}
	}

	req := c.newRequest(op, input, &BuildSuggestersOutput{})

	return BuildSuggestersRequest{Request: req, Input: input, Copy: c.BuildSuggestersRequest}
}

// BuildSuggestersRequest is the request type for the
// BuildSuggesters API operation.
type BuildSuggestersRequest struct {
	*aws.Request
	Input *BuildSuggestersInput
	Copy  func(*BuildSuggestersInput) BuildSuggestersRequest
}

// Send marshals and sends the BuildSuggesters API request.
func (r BuildSuggestersRequest) Send(ctx context.Context) (*BuildSuggestersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BuildSuggestersResponse{
		BuildSuggestersOutput: r.Request.Data.(*BuildSuggestersOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BuildSuggestersResponse is the response type for the
// BuildSuggesters API operation.
type BuildSuggestersResponse struct {
	*BuildSuggestersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BuildSuggesters request.
func (r *BuildSuggestersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
