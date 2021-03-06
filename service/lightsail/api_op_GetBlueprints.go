// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetBlueprintsInput struct {
	_ struct{} `type:"structure"`

	// A Boolean value indicating whether to include inactive results in your request.
	IncludeInactive *bool `locationName:"includeInactive" type:"boolean"`

	// The token to advance to the next page of results from your request.
	//
	// To get a page token, perform an initial GetBlueprints request. If your results
	// are paginated, the response will return a next page token that you can specify
	// as the page token in a subsequent request.
	PageToken *string `locationName:"pageToken" type:"string"`
}

// String returns the string representation
func (s GetBlueprintsInput) String() string {
	return awsutil.Prettify(s)
}

type GetBlueprintsOutput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs that contains information about the available
	// blueprints.
	Blueprints []Blueprint `locationName:"blueprints" type:"list"`

	// The token to advance to the next page of resutls from your request.
	//
	// A next page token is not returned if there are no more results to display.
	//
	// To get the next page of results, perform another GetBlueprints request and
	// specify the next page token using the pageToken parameter.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`
}

// String returns the string representation
func (s GetBlueprintsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetBlueprints = "GetBlueprints"

// GetBlueprintsRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns the list of available instance images, or blueprints. You can use
// a blueprint to create a new instance already running a specific operating
// system, as well as a preinstalled app or development stack. The software
// each instance is running depends on the blueprint image you choose.
//
// Use active blueprints when creating new instances. Inactive blueprints are
// listed to support customers with existing instances and are not necessarily
// available to create new instances. Blueprints are marked inactive when they
// become outdated due to operating system updates or new application releases.
//
//    // Example sending a request using GetBlueprintsRequest.
//    req := client.GetBlueprintsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetBlueprints
func (c *Client) GetBlueprintsRequest(input *GetBlueprintsInput) GetBlueprintsRequest {
	op := &aws.Operation{
		Name:       opGetBlueprints,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetBlueprintsInput{}
	}

	req := c.newRequest(op, input, &GetBlueprintsOutput{})

	return GetBlueprintsRequest{Request: req, Input: input, Copy: c.GetBlueprintsRequest}
}

// GetBlueprintsRequest is the request type for the
// GetBlueprints API operation.
type GetBlueprintsRequest struct {
	*aws.Request
	Input *GetBlueprintsInput
	Copy  func(*GetBlueprintsInput) GetBlueprintsRequest
}

// Send marshals and sends the GetBlueprints API request.
func (r GetBlueprintsRequest) Send(ctx context.Context) (*GetBlueprintsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBlueprintsResponse{
		GetBlueprintsOutput: r.Request.Data.(*GetBlueprintsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBlueprintsResponse is the response type for the
// GetBlueprints API operation.
type GetBlueprintsResponse struct {
	*GetBlueprintsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBlueprints request.
func (r *GetBlueprintsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
