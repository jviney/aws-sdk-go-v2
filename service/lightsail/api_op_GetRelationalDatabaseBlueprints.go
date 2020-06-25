// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetRelationalDatabaseBlueprintsInput struct {
	_ struct{} `type:"structure"`

	// The token to advance to the next page of results from your request.
	//
	// To get a page token, perform an initial GetRelationalDatabaseBlueprints request.
	// If your results are paginated, the response will return a next page token
	// that you can specify as the page token in a subsequent request.
	PageToken *string `locationName:"pageToken" type:"string"`
}

// String returns the string representation
func (s GetRelationalDatabaseBlueprintsInput) String() string {
	return awsutil.Prettify(s)
}

type GetRelationalDatabaseBlueprintsOutput struct {
	_ struct{} `type:"structure"`

	// An object describing the result of your get relational database blueprints
	// request.
	Blueprints []RelationalDatabaseBlueprint `locationName:"blueprints" type:"list"`

	// The token to advance to the next page of resutls from your request.
	//
	// A next page token is not returned if there are no more results to display.
	//
	// To get the next page of results, perform another GetRelationalDatabaseBlueprints
	// request and specify the next page token using the pageToken parameter.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`
}

// String returns the string representation
func (s GetRelationalDatabaseBlueprintsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRelationalDatabaseBlueprints = "GetRelationalDatabaseBlueprints"

// GetRelationalDatabaseBlueprintsRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns a list of available database blueprints in Amazon Lightsail. A blueprint
// describes the major engine version of a database.
//
// You can use a blueprint ID to create a new database that runs a specific
// database engine.
//
//    // Example sending a request using GetRelationalDatabaseBlueprintsRequest.
//    req := client.GetRelationalDatabaseBlueprintsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetRelationalDatabaseBlueprints
func (c *Client) GetRelationalDatabaseBlueprintsRequest(input *GetRelationalDatabaseBlueprintsInput) GetRelationalDatabaseBlueprintsRequest {
	op := &aws.Operation{
		Name:       opGetRelationalDatabaseBlueprints,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRelationalDatabaseBlueprintsInput{}
	}

	req := c.newRequest(op, input, &GetRelationalDatabaseBlueprintsOutput{})

	return GetRelationalDatabaseBlueprintsRequest{Request: req, Input: input, Copy: c.GetRelationalDatabaseBlueprintsRequest}
}

// GetRelationalDatabaseBlueprintsRequest is the request type for the
// GetRelationalDatabaseBlueprints API operation.
type GetRelationalDatabaseBlueprintsRequest struct {
	*aws.Request
	Input *GetRelationalDatabaseBlueprintsInput
	Copy  func(*GetRelationalDatabaseBlueprintsInput) GetRelationalDatabaseBlueprintsRequest
}

// Send marshals and sends the GetRelationalDatabaseBlueprints API request.
func (r GetRelationalDatabaseBlueprintsRequest) Send(ctx context.Context) (*GetRelationalDatabaseBlueprintsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRelationalDatabaseBlueprintsResponse{
		GetRelationalDatabaseBlueprintsOutput: r.Request.Data.(*GetRelationalDatabaseBlueprintsOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRelationalDatabaseBlueprintsResponse is the response type for the
// GetRelationalDatabaseBlueprints API operation.
type GetRelationalDatabaseBlueprintsResponse struct {
	*GetRelationalDatabaseBlueprintsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRelationalDatabaseBlueprints request.
func (r *GetRelationalDatabaseBlueprintsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
