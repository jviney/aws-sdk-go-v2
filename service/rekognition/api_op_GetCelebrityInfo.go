// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetCelebrityInfoInput struct {
	_ struct{} `type:"structure"`

	// The ID for the celebrity. You get the celebrity ID from a call to the RecognizeCelebrities
	// operation, which recognizes celebrities in an image.
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetCelebrityInfoInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCelebrityInfoInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCelebrityInfoInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetCelebrityInfoOutput struct {
	_ struct{} `type:"structure"`

	// The name of the celebrity.
	Name *string `type:"string"`

	// An array of URLs pointing to additional celebrity information.
	Urls []string `type:"list"`
}

// String returns the string representation
func (s GetCelebrityInfoOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetCelebrityInfo = "GetCelebrityInfo"

// GetCelebrityInfoRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Gets the name and additional information about a celebrity based on his or
// her Amazon Rekognition ID. The additional information is returned as an array
// of URLs. If there is no additional information about the celebrity, this
// list is empty.
//
// For more information, see Recognizing Celebrities in an Image in the Amazon
// Rekognition Developer Guide.
//
// This operation requires permissions to perform the rekognition:GetCelebrityInfo
// action.
//
//    // Example sending a request using GetCelebrityInfoRequest.
//    req := client.GetCelebrityInfoRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetCelebrityInfoRequest(input *GetCelebrityInfoInput) GetCelebrityInfoRequest {
	op := &aws.Operation{
		Name:       opGetCelebrityInfo,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetCelebrityInfoInput{}
	}

	req := c.newRequest(op, input, &GetCelebrityInfoOutput{})

	return GetCelebrityInfoRequest{Request: req, Input: input, Copy: c.GetCelebrityInfoRequest}
}

// GetCelebrityInfoRequest is the request type for the
// GetCelebrityInfo API operation.
type GetCelebrityInfoRequest struct {
	*aws.Request
	Input *GetCelebrityInfoInput
	Copy  func(*GetCelebrityInfoInput) GetCelebrityInfoRequest
}

// Send marshals and sends the GetCelebrityInfo API request.
func (r GetCelebrityInfoRequest) Send(ctx context.Context) (*GetCelebrityInfoResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCelebrityInfoResponse{
		GetCelebrityInfoOutput: r.Request.Data.(*GetCelebrityInfoOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCelebrityInfoResponse is the response type for the
// GetCelebrityInfo API operation.
type GetCelebrityInfoResponse struct {
	*GetCelebrityInfoOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCelebrityInfo request.
func (r *GetCelebrityInfoResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
