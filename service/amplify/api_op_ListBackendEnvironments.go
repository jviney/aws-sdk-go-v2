// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// Request structure for list backend environments request.
type ListBackendEnvironmentsInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for an amplify App.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// Name of the backend environment
	EnvironmentName *string `locationName:"environmentName" min:"1" type:"string"`

	// Maximum number of records to list in a single response.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// Pagination token. Set to null to start listing backen environments from start.
	// If a non-null pagination token is returned in a result, then pass its value
	// in here to list more backend environments.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListBackendEnvironmentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBackendEnvironmentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListBackendEnvironmentsInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}
	if s.EnvironmentName != nil && len(*s.EnvironmentName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EnvironmentName", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListBackendEnvironmentsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.EnvironmentName != nil {
		v := *s.EnvironmentName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "environmentName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AppId != nil {
		v := *s.AppId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "appId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure for list backend environments result.
type ListBackendEnvironmentsOutput struct {
	_ struct{} `type:"structure"`

	// List of backend environments for an Amplify App.
	//
	// BackendEnvironments is a required field
	BackendEnvironments []BackendEnvironment `locationName:"backendEnvironments" type:"list" required:"true"`

	// Pagination token. If non-null pagination token is returned in a result, then
	// pass its value in another request to fetch more entries.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListBackendEnvironmentsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListBackendEnvironmentsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BackendEnvironments != nil {
		v := s.BackendEnvironments

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "backendEnvironments", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListBackendEnvironments = "ListBackendEnvironments"

// ListBackendEnvironmentsRequest returns a request value for making API operation for
// AWS Amplify.
//
// Lists backend environments for an Amplify App.
//
//    // Example sending a request using ListBackendEnvironmentsRequest.
//    req := client.ListBackendEnvironmentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/ListBackendEnvironments
func (c *Client) ListBackendEnvironmentsRequest(input *ListBackendEnvironmentsInput) ListBackendEnvironmentsRequest {
	op := &aws.Operation{
		Name:       opListBackendEnvironments,
		HTTPMethod: "GET",
		HTTPPath:   "/apps/{appId}/backendenvironments",
	}

	if input == nil {
		input = &ListBackendEnvironmentsInput{}
	}

	req := c.newRequest(op, input, &ListBackendEnvironmentsOutput{})

	return ListBackendEnvironmentsRequest{Request: req, Input: input, Copy: c.ListBackendEnvironmentsRequest}
}

// ListBackendEnvironmentsRequest is the request type for the
// ListBackendEnvironments API operation.
type ListBackendEnvironmentsRequest struct {
	*aws.Request
	Input *ListBackendEnvironmentsInput
	Copy  func(*ListBackendEnvironmentsInput) ListBackendEnvironmentsRequest
}

// Send marshals and sends the ListBackendEnvironments API request.
func (r ListBackendEnvironmentsRequest) Send(ctx context.Context) (*ListBackendEnvironmentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListBackendEnvironmentsResponse{
		ListBackendEnvironmentsOutput: r.Request.Data.(*ListBackendEnvironmentsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListBackendEnvironmentsResponse is the response type for the
// ListBackendEnvironments API operation.
type ListBackendEnvironmentsResponse struct {
	*ListBackendEnvironmentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListBackendEnvironments request.
func (r *ListBackendEnvironmentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
