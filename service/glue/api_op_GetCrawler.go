// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetCrawlerInput struct {
	_ struct{} `type:"structure"`

	// The name of the crawler to retrieve metadata for.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetCrawlerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCrawlerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCrawlerInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetCrawlerOutput struct {
	_ struct{} `type:"structure"`

	// The metadata for the specified crawler.
	Crawler *Crawler `type:"structure"`
}

// String returns the string representation
func (s GetCrawlerOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetCrawler = "GetCrawler"

// GetCrawlerRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves metadata for a specified crawler.
//
//    // Example sending a request using GetCrawlerRequest.
//    req := client.GetCrawlerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetCrawler
func (c *Client) GetCrawlerRequest(input *GetCrawlerInput) GetCrawlerRequest {
	op := &aws.Operation{
		Name:       opGetCrawler,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetCrawlerInput{}
	}

	req := c.newRequest(op, input, &GetCrawlerOutput{})

	return GetCrawlerRequest{Request: req, Input: input, Copy: c.GetCrawlerRequest}
}

// GetCrawlerRequest is the request type for the
// GetCrawler API operation.
type GetCrawlerRequest struct {
	*aws.Request
	Input *GetCrawlerInput
	Copy  func(*GetCrawlerInput) GetCrawlerRequest
}

// Send marshals and sends the GetCrawler API request.
func (r GetCrawlerRequest) Send(ctx context.Context) (*GetCrawlerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCrawlerResponse{
		GetCrawlerOutput: r.Request.Data.(*GetCrawlerOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCrawlerResponse is the response type for the
// GetCrawler API operation.
type GetCrawlerResponse struct {
	*GetCrawlerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCrawler request.
func (r *GetCrawlerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
