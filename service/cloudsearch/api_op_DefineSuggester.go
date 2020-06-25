// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudsearch

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Container for the parameters to the DefineSuggester operation. Specifies
// the name of the domain you want to update and the suggester configuration.
type DefineSuggesterInput struct {
	_ struct{} `type:"structure"`

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start
	// with a letter or number and can contain the following characters: a-z (lowercase),
	// 0-9, and - (hyphen).
	//
	// DomainName is a required field
	DomainName *string `min:"3" type:"string" required:"true"`

	// Configuration information for a search suggester. Each suggester has a unique
	// name and specifies the text field you want to use for suggestions. The following
	// options can be configured for a suggester: FuzzyMatching, SortExpression.
	//
	// Suggester is a required field
	Suggester *Suggester `type:"structure" required:"true"`
}

// String returns the string representation
func (s DefineSuggesterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DefineSuggesterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DefineSuggesterInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 3))
	}

	if s.Suggester == nil {
		invalidParams.Add(aws.NewErrParamRequired("Suggester"))
	}
	if s.Suggester != nil {
		if err := s.Suggester.Validate(); err != nil {
			invalidParams.AddNested("Suggester", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of a DefineSuggester request. Contains the status of the newly-configured
// suggester.
type DefineSuggesterOutput struct {
	_ struct{} `type:"structure"`

	// The value of a Suggester and its current status.
	//
	// Suggester is a required field
	Suggester *SuggesterStatus `type:"structure" required:"true"`
}

// String returns the string representation
func (s DefineSuggesterOutput) String() string {
	return awsutil.Prettify(s)
}

const opDefineSuggester = "DefineSuggester"

// DefineSuggesterRequest returns a request value for making API operation for
// Amazon CloudSearch.
//
// Configures a suggester for a domain. A suggester enables you to display possible
// matches before users finish typing their queries. When you configure a suggester,
// you must specify the name of the text field you want to search for possible
// matches and a unique name for the suggester. For more information, see Getting
// Search Suggestions (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-suggestions.html)
// in the Amazon CloudSearch Developer Guide.
//
//    // Example sending a request using DefineSuggesterRequest.
//    req := client.DefineSuggesterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DefineSuggesterRequest(input *DefineSuggesterInput) DefineSuggesterRequest {
	op := &aws.Operation{
		Name:       opDefineSuggester,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DefineSuggesterInput{}
	}

	req := c.newRequest(op, input, &DefineSuggesterOutput{})

	return DefineSuggesterRequest{Request: req, Input: input, Copy: c.DefineSuggesterRequest}
}

// DefineSuggesterRequest is the request type for the
// DefineSuggester API operation.
type DefineSuggesterRequest struct {
	*aws.Request
	Input *DefineSuggesterInput
	Copy  func(*DefineSuggesterInput) DefineSuggesterRequest
}

// Send marshals and sends the DefineSuggester API request.
func (r DefineSuggesterRequest) Send(ctx context.Context) (*DefineSuggesterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DefineSuggesterResponse{
		DefineSuggesterOutput: r.Request.Data.(*DefineSuggesterOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DefineSuggesterResponse is the response type for the
// DefineSuggester API operation.
type DefineSuggesterResponse struct {
	*DefineSuggesterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DefineSuggester request.
func (r *DefineSuggesterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
